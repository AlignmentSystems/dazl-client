# Copyright (c) 2017-2020 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

"""
Support for the gRPC-based Ledger API.
"""
from asyncio import gather, get_event_loop
from datetime import datetime
from threading import Event
from typing import Awaitable, Iterable, Optional, Sequence

# noinspection PyPackageRequirements
from grpc import Channel, secure_channel, insecure_channel, ssl_channel_credentials, RpcError, \
    StatusCode

from ... import LOG
from .._base import LedgerClient, _LedgerConnection, LedgerConnectionOptions
from .pb_parse_event import serialize_acs_request, serialize_transactions_request, \
    to_acs_events, to_transaction_events, BaseEventDeserializationContext
from .pb_parse_metadata import parse_daml_metadata_pb, parse_archive_payload, find_dependencies
from ...model.core import Party, UserTerminateRequest, ConnectionTimeoutError
from ...model.ledger import LedgerMetadata
from ...model.network import HTTPConnectionSettings
from ...model.reading import BaseEvent, ContractFilter, TransactionFilter
from ...model.types import PackageId, PackageIdSet
from ...model.types_store import PackageStore, PackageProvider
from ...model.writing import CommandPayload
from ...scheduler import Invoker, RunLevel
from ...util.io import read_file_bytes
from ...util.typing import safe_cast


class GRPCv1LedgerClient(LedgerClient):
    def __init__(self, connection: 'GRPCv1Connection', ledger: LedgerMetadata, party: Party):
        self.connection = safe_cast(GRPCv1Connection, connection)
        self.ledger = safe_cast(LedgerMetadata, ledger)
        self.party = Party(safe_cast(str, party))

    def commands(self, commands: CommandPayload) -> None:
        serializer = self.ledger.serializer
        request = serializer.serialize_command_request(commands)
        return self.connection.invoker.run_in_executor(
            lambda: self.connection.command_service.SubmitAndWait(request))

    async def active_contracts(self, contract_filter: 'ContractFilter') -> 'Sequence[BaseEvent]':
        request = serialize_acs_request(contract_filter, self.ledger.ledger_id, self.party)
        context = BaseEventDeserializationContext(
            None, self.ledger.store, self.party, self.ledger.ledger_id)
        acs_stream = await self.connection.invoker.run_in_executor(
            lambda: self.connection.active_contract_set_service.GetActiveContracts(request))
        acs_events = to_acs_events(context, acs_stream)
        return acs_events

    def events(self, transaction_filter: 'TransactionFilter') \
            -> Awaitable[Sequence[BaseEvent]]:
        results_future = get_event_loop().create_future()
        request = serialize_transactions_request(
            transaction_filter, self.ledger.ledger_id, self.party)
        context = BaseEventDeserializationContext(
            None, self.ledger.store, self.party, self.ledger.ledger_id)

        def on_events_done(fut):
            try:
                tx_stream, tt_stream = fut.result()
                events = to_transaction_events(
                    context, tx_stream, tt_stream, transaction_filter.destination_offset)
            except Exception as ex:
                LOG.exception('Failed to parse data coming back from an event!')
                results_future.set_exception(ex)
                return
            results_future.set_result(events)

        ts_future = self.connection.invoker.run_in_executor(
            lambda: self.connection.transaction_service.GetTransactions(request))

        # Filtering by package must disable the ability to handle exercise nodes; we may want to
        # consider dropping client-side support for exercise events anyway because they are not
        # widely used
        if transaction_filter.templates is None:
            tst_future = self.connection.invoker.run_in_executor(
                lambda: self.connection.transaction_service.GetTransactionTrees(request))
        else:
            tst_future = get_event_loop().create_future()
            tst_future.set_result(None)

        t_fut = gather(ts_future, tst_future)
        t_fut.add_done_callback(on_events_done)

        return results_future

    async def events_end(self) -> str:
        from . import model as G
        request = G.GetLedgerEndRequest(ledger_id=self.ledger.ledger_id)
        return await self.connection.invoker.run_in_executor(
            lambda: self.connection.transaction_service.GetLedgerEnd(request).offset.absolute)


def grpc_set_time(connection: 'GRPCv1Connection', ledger_id: str, new_datetime: datetime) -> None:
    from . import model as G
    from .pb_ser_command import as_api_timestamp

    request = G.GetTimeRequest(ledger_id=ledger_id)
    response = connection.time_service.GetTime(request)
    ts = next(iter(response))

    request = G.SetTimeRequest(
        ledger_id=ledger_id,
        current_time=ts.current_time,
        new_time=as_api_timestamp(new_datetime))
    connection.time_service.SetTime(request)
    LOG.info('Time on the server changed by the local client to %s.', new_datetime)


def grpc_upload_package(connection: 'GRPCv1Connection', dar_contents: bytes) -> None:
    from . import model as G

    request = G.UploadDarFileRequest(dar_file=dar_contents)
    connection.package_management_service.UploadDarFile(request)


GRPC_KNOWN_RETRYABLE_ERRORS = ('DNS resolution failed', 'failed to connect to all addresses', 'no healthy upstream')


def grpc_detect_ledger_id(connection: 'GRPCv1Connection') -> str:
    """
    Return the ledger ID from the remote server when it becomes available. This method blocks until
    a ledger ID has been successfully retrieved, or the timeout is reached (in which case an
    exception is thrown).
    """
    from . import model as G
    from time import sleep
    LOG.debug("Starting a monitor thread for connection: %s", connection)

    start_time = datetime.utcnow()
    connect_timeout = connection.options.connect_timeout

    while connect_timeout is None or (datetime.utcnow() - start_time) < connect_timeout:
        if connection.invoker.level >= RunLevel.TERMINATE_GRACEFULLY:
            raise UserTerminateRequest()
        if connection.closed:
            raise Exception('connection closed')

        try:
            response = connection.ledger_identity_service.GetLedgerIdentity(
                G.GetLedgerIdentityRequest())
        except RpcError as ex:
            details_str = ex.details()

            # suppress some warning strings because they're not particularly useful and just clutter
            # up the logs
            if details_str not in GRPC_KNOWN_RETRYABLE_ERRORS:
                LOG.exception('An unexpected error occurred when trying to fetch the '
                              'ledger identity; this will be retried.')
            sleep(1)
            continue

        return response.ledger_id

    raise ConnectionTimeoutError(
        f'connection timeout exceeded: {connect_timeout.total_seconds()} seconds')


def grpc_main_thread(connection: 'GRPCv1Connection', ledger_id: str) -> 'Iterable[LedgerMetadata]':
    from .pb_ser_command import ProtobufSerializer

    store = PackageStore.empty()

    package_provider = GRPCPackageProvider(connection, ledger_id)

    # TODO: We could probably disable package syncing if expected packages are provided AND we have
    #  fetched metadata for them all.
    grpc_package_sync(package_provider, store)

    yield LedgerMetadata(
        ledger_id=ledger_id,
        store=store,
        serializer=ProtobufSerializer(store),
        protocol_version="v1")

    # poll for package updates once a second
    while not connection._closed.wait(1):
        try:
            grpc_package_sync(package_provider, store)
        except Exception as ex:
            if not isinstance(ex, RpcError) or ex.code() != StatusCode.CANCELLED:
                LOG.warning('Package syncing raised an exception.', exc_info=True)

    LOG.debug('The gRPC monitor thread is now shutting down.')
    yield None


class GRPCPackageProvider(PackageProvider):
    def __init__(self, connection: 'GRPCv1Connection', ledger_id: str):
        self.connection = connection
        self.ledger_id = ledger_id

    def get_package_ids(self) -> 'PackageIdSet':
        from . import model as G
        request = G.ListPackagesRequest(ledger_id=self.ledger_id)
        response = self.connection.package_service.ListPackages(request)
        return frozenset(response.package_ids)

    def fetch_package(self, package_id: 'PackageId') -> bytes:
        from . import model as G
        request = G.GetPackageRequest(ledger_id=self.ledger_id, package_id=package_id)
        package_info = self.connection.package_service.GetPackage(request)
        return package_info.archive_payload


def grpc_package_sync(package_provider: 'PackageProvider', store: 'PackageStore') -> 'None':
    all_package_ids = package_provider.get_package_ids()
    loaded_package_ids = {a.hash for a in store.archives()}
    expected_package_ids = store.expected_package_ids()

    missing_package_ids = all_package_ids - loaded_package_ids

    def should_load(package_id: str) -> bool:
        # TODO: Filtering by expected package IDs may cause packages to never fully load due to
        #  transitive dependencies; here we put the onus on the app writer to ensure that they
        #  supply the complete graph, and we don't even warn them if there is an issue (but
        #  we could only actually warn them if we parse the archive, which is the expensive
        #  operation we're trying to avoid).
        return (expected_package_ids is None or package_id in expected_package_ids) and \
            package_id not in loaded_package_ids

    metadatas_pb = {}
    for package_id in all_package_ids:
        if should_load(package_id):
            archive_payload = package_provider.fetch_package(package_id)
            metadatas_pb[package_id] = parse_archive_payload(archive_payload)

    metadatas_pb = find_dependencies(metadatas_pb, loaded_package_ids)
    for package_id, archive_payload in metadatas_pb.sorted_archives.items():
        store.register_all(parse_daml_metadata_pb(package_id, archive_payload))

def grpc_create_channel(settings: HTTPConnectionSettings) -> Channel:
    target = f'{settings.host}:{settings.port}'
    options = [('grpc.max_send_message_length', -1),
               ('grpc.max_receive_message_length', -1)]

    if settings.oauth:
        # noinspection PyPackageRequirements
        from google.auth.transport.grpc import secure_authorized_channel
        # noinspection PyPackageRequirements
        from google.auth.transport.requests import Request as RefreshRequester
        # noinspection PyPackageRequirements
        from google.oauth2.credentials import Credentials as OAuthCredentials

        LOG.debug('Using a secure gRPC connection over OAuth:')

        credentials = OAuthCredentials(
            token=settings.oauth.token,
            refresh_token=settings.oauth.refresh_token,
            id_token=settings.oauth.id_token,
            token_uri=settings.oauth.token_uri,
            client_id=settings.oauth.client_id,
            client_secret=settings.oauth.client_secret)

        ssl_credentials=None
        if settings.ssl_settings:
            cert_chain = read_file_bytes(settings.ssl_settings.cert_file)
            cert = read_file_bytes(settings.ssl_settings.cert_key_file)
            ca_cert = read_file_bytes(settings.ssl_settings.ca_file)

            LOG.debug('Using a secure gRPC connection:')
            LOG.debug('    target: %s', target)
            LOG.debug('    root_certificates: contents of %s', settings.ssl_settings.ca_file)
            LOG.debug('    private_key: contents of %s', settings.ssl_settings.cert_key_file)
            LOG.debug('    certificate_chain: contents of %s', settings.ssl_settings.cert_file)

            ssl_credentials = ssl_channel_credentials(
                root_certificates=ca_cert,
                private_key=cert,
                certificate_chain=cert_chain)
        return secure_authorized_channel(credentials, RefreshRequester(), target, ssl_credentials=ssl_credentials, options=options)

    if settings.ssl_settings:
        cert_chain = read_file_bytes(settings.ssl_settings.cert_file)
        cert = read_file_bytes(settings.ssl_settings.cert_key_file)
        ca_cert = read_file_bytes(settings.ssl_settings.ca_file)

        LOG.debug('Using a secure gRPC connection:')
        LOG.debug('    target: %s', target)
        LOG.debug('    root_certificates: contents of %s', settings.ssl_settings.ca_file)
        LOG.debug('    private_key: contents of %s', settings.ssl_settings.cert_key_file)
        LOG.debug('    certificate_chain: contents of %s', settings.ssl_settings.cert_file)

        credentials = ssl_channel_credentials(
            root_certificates=ca_cert,
            private_key=cert,
            certificate_chain=cert_chain)
        return secure_channel(target, credentials, options)
    else:
        LOG.debug('Using an insecure gRPC connection...')
        return insecure_channel(target, options)


class GRPCv1Connection(_LedgerConnection):
    def __init__(self,
                 invoker: 'Invoker',
                 options: 'LedgerConnectionOptions',
                 settings: 'HTTPConnectionSettings',
                 context_path: 'Optional[str]'):
        super(GRPCv1Connection, self).__init__(invoker, options, settings, context_path)
        from . import model as G

        self._closed = Event()
        self._channel = grpc_create_channel(settings)
        self.active_contract_set_service = G.ActiveContractsServiceStub(self._channel)
        self.command_service = G.CommandServiceStub(self._channel)
        self.transaction_service = G.TransactionServiceStub(self._channel)
        self.package_service = G.PackageServiceStub(self._channel)
        self.package_management_service = G.PackageManagementServiceStub(self._channel)
        self.ledger_identity_service = G.LedgerIdentityServiceStub(self._channel)
        self.time_service = G.TimeServiceStub(self._channel)

    def close(self):
        try:
            self._closed.set()
            self._channel.close()
        except:
            LOG.warning('An exception was thrown when trying to close down connections.')
        finally:
            super().close()

    @property
    def closed(self) -> bool:
        return self._closed.is_set()
