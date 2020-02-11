# Copyright (c) 2019 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from . import command_submission_service_pb2 as com_dot_digitalasset_dot_ledger_dot_api_dot_v1_dot_command__submission__service__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


class CommandSubmissionServiceStub(object):
  """Allows clients to attempt advancing the ledger's state by submitting commands.
  The final states of their submissions are disclosed by the Command Completion Service.
  The on-ledger effects of their submissions are disclosed by the Transaction Service.
  Commands may fail in 4 distinct manners:

  1) ``INVALID_PARAMETER`` gRPC error on malformed payloads and missing required fields.
  2) Failure communicated in the gRPC error.
  3) Failure communicated in a Completion.
  4) A Checkpoint with ``record_time`` > command ``mrt`` arrives through the Completion Stream, and the command's Completion was not visible before. In this case the command is lost.

  Clients that do not receive a successful completion about their submission MUST NOT assume that it was successful.
  Clients SHOULD subscribe to the CompletionStream before starting to submit commands to prevent race conditions.

  Interprocess tracing of command submissions may be achieved via Zipkin by filling out the ``trace_context`` field.
  The server will return a child context of the submitted one, (or a new one if the context was missing) on both the Completion and Transaction streams.
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Submit = channel.unary_unary(
        '/com.digitalasset.ledger.api.v1.CommandSubmissionService/Submit',
        request_serializer=com_dot_digitalasset_dot_ledger_dot_api_dot_v1_dot_command__submission__service__pb2.SubmitRequest.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )


class CommandSubmissionServiceServicer(object):
  """Allows clients to attempt advancing the ledger's state by submitting commands.
  The final states of their submissions are disclosed by the Command Completion Service.
  The on-ledger effects of their submissions are disclosed by the Transaction Service.
  Commands may fail in 4 distinct manners:

  1) ``INVALID_PARAMETER`` gRPC error on malformed payloads and missing required fields.
  2) Failure communicated in the gRPC error.
  3) Failure communicated in a Completion.
  4) A Checkpoint with ``record_time`` > command ``mrt`` arrives through the Completion Stream, and the command's Completion was not visible before. In this case the command is lost.

  Clients that do not receive a successful completion about their submission MUST NOT assume that it was successful.
  Clients SHOULD subscribe to the CompletionStream before starting to submit commands to prevent race conditions.

  Interprocess tracing of command submissions may be achieved via Zipkin by filling out the ``trace_context`` field.
  The server will return a child context of the submitted one, (or a new one if the context was missing) on both the Completion and Transaction streams.
  """

  def Submit(self, request, context):
    """Submit a single composite command.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_CommandSubmissionServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Submit': grpc.unary_unary_rpc_method_handler(
          servicer.Submit,
          request_deserializer=com_dot_digitalasset_dot_ledger_dot_api_dot_v1_dot_command__submission__service__pb2.SubmitRequest.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'com.digitalasset.ledger.api.v1.CommandSubmissionService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))