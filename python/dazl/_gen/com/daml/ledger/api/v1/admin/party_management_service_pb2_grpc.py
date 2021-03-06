# Copyright (c) 2017-2020 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from . import party_management_service_pb2 as com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_party__management__service__pb2


class PartyManagementServiceStub(object):
  """Status: experimental interface, will change before it is deemed production
  ready

  Inspect the party management state of a ledger participant and modify the
  parts that are modifiable. We use 'backing participant' to refer to this
  specific participant in the methods of this API.
  When the participant is run in mode requiring authentication, all the calls
  in this interface will respond with UNAUTHENTICATED, if the caller fails
  to provide a valid access token, and will respond with PERMISSION_DENIED, if
  the claims in the token are insufficient to perform a given operation.
  Subsequently, only specific errors of individual calls not related to
  authorization will be described.
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetParticipantId = channel.unary_unary(
        '/com.daml.ledger.api.v1.admin.PartyManagementService/GetParticipantId',
        request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_party__management__service__pb2.GetParticipantIdRequest.SerializeToString,
        response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_party__management__service__pb2.GetParticipantIdResponse.FromString,
        )
    self.GetParties = channel.unary_unary(
        '/com.daml.ledger.api.v1.admin.PartyManagementService/GetParties',
        request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_party__management__service__pb2.GetPartiesRequest.SerializeToString,
        response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_party__management__service__pb2.GetPartiesResponse.FromString,
        )
    self.ListKnownParties = channel.unary_unary(
        '/com.daml.ledger.api.v1.admin.PartyManagementService/ListKnownParties',
        request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_party__management__service__pb2.ListKnownPartiesRequest.SerializeToString,
        response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_party__management__service__pb2.ListKnownPartiesResponse.FromString,
        )
    self.AllocateParty = channel.unary_unary(
        '/com.daml.ledger.api.v1.admin.PartyManagementService/AllocateParty',
        request_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_party__management__service__pb2.AllocatePartyRequest.SerializeToString,
        response_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_party__management__service__pb2.AllocatePartyResponse.FromString,
        )


class PartyManagementServiceServicer(object):
  """Status: experimental interface, will change before it is deemed production
  ready

  Inspect the party management state of a ledger participant and modify the
  parts that are modifiable. We use 'backing participant' to refer to this
  specific participant in the methods of this API.
  When the participant is run in mode requiring authentication, all the calls
  in this interface will respond with UNAUTHENTICATED, if the caller fails
  to provide a valid access token, and will respond with PERMISSION_DENIED, if
  the claims in the token are insufficient to perform a given operation.
  Subsequently, only specific errors of individual calls not related to
  authorization will be described.
  """

  def GetParticipantId(self, request, context):
    """Return the identifier of the backing participant.
    All horizontally scaled replicas should return the same id.
    This method is expected to succeed provided the backing participant is
    healthy, otherwise it responds with INTERNAL grpc error.
    daml-on-sql: returns an identifier supplied on command line at launch time
    daml-on-kv-ledger: as above
    canton: returns globally unique identifier of the backing participant
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetParties(self, request, context):
    """Get the party details of the given parties. Only known parties will be
    returned in the list.
    This request will always succeed.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListKnownParties(self, request, context):
    """List the parties known by the backing participant.
    The list returned contains parties whose ledger access is facilitated by
    backing participant and the ones maintained elsewhere.
    This request will always succeed.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def AllocateParty(self, request, context):
    """Adds a new party to the set managed by the backing participant.
    Caller specifies a party identifier suggestion, the actual identifier
    allocated might be different and is implementation specific.
    This call may:
    - Succeed, in which case the actual allocated identifier is visible in
    the response.
    - Respond with UNIMPLEMENTED if synchronous party allocation is not
    supported by the backing participant.
    - Respond with INVALID_ARGUMENT if the provided hint and/or display name
    is invalid on the given ledger (see below).
    daml-on-sql: suggestion's uniqueness is checked and call rejected if the
    identifier is already present
    daml-on-kv-ledger: suggestion's uniqueness is checked by the validators in
    the consensus layer and call rejected if the identifier is already present.
    canton: completely different globally unique identifier is allocated.
    Behind the scenes calls to an internal protocol are made. As that protocol
    is richer than the the surface protocol, the arguments take implicit values
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_PartyManagementServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetParticipantId': grpc.unary_unary_rpc_method_handler(
          servicer.GetParticipantId,
          request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_party__management__service__pb2.GetParticipantIdRequest.FromString,
          response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_party__management__service__pb2.GetParticipantIdResponse.SerializeToString,
      ),
      'GetParties': grpc.unary_unary_rpc_method_handler(
          servicer.GetParties,
          request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_party__management__service__pb2.GetPartiesRequest.FromString,
          response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_party__management__service__pb2.GetPartiesResponse.SerializeToString,
      ),
      'ListKnownParties': grpc.unary_unary_rpc_method_handler(
          servicer.ListKnownParties,
          request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_party__management__service__pb2.ListKnownPartiesRequest.FromString,
          response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_party__management__service__pb2.ListKnownPartiesResponse.SerializeToString,
      ),
      'AllocateParty': grpc.unary_unary_rpc_method_handler(
          servicer.AllocateParty,
          request_deserializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_party__management__service__pb2.AllocatePartyRequest.FromString,
          response_serializer=com_dot_daml_dot_ledger_dot_api_dot_v1_dot_admin_dot_party__management__service__pb2.AllocatePartyResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'com.daml.ledger.api.v1.admin.PartyManagementService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
