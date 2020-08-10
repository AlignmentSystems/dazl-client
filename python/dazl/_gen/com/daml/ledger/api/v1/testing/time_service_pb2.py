# Copyright (c) 2020 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/daml/ledger/api/v1/testing/time_service.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='com/daml/ledger/api/v1/testing/time_service.proto',
  package='com.daml.ledger.api.v1.testing',
  syntax='proto3',
  serialized_options=b'\n\036com.daml.ledger.api.v1.testingB\025TimeServiceOuterClass\252\002\036Com.Daml.Ledger.Api.V1.Testing',
  serialized_pb=b'\n1com/daml/ledger/api/v1/testing/time_service.proto\x12\x1e\x63om.daml.ledger.api.v1.testing\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"#\n\x0eGetTimeRequest\x12\x11\n\tledger_id\x18\x01 \x01(\t\"C\n\x0fGetTimeResponse\x12\x30\n\x0c\x63urrent_time\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\x83\x01\n\x0eSetTimeRequest\x12\x11\n\tledger_id\x18\x01 \x01(\t\x12\x30\n\x0c\x63urrent_time\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12,\n\x08new_time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp2\xce\x01\n\x0bTimeService\x12l\n\x07GetTime\x12..com.daml.ledger.api.v1.testing.GetTimeRequest\x1a/.com.daml.ledger.api.v1.testing.GetTimeResponse0\x01\x12Q\n\x07SetTime\x12..com.daml.ledger.api.v1.testing.SetTimeRequest\x1a\x16.google.protobuf.EmptyBX\n\x1e\x63om.daml.ledger.api.v1.testingB\x15TimeServiceOuterClass\xaa\x02\x1e\x43om.Daml.Ledger.Api.V1.Testingb\x06proto3'
  ,
  dependencies=[google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_GETTIMEREQUEST = _descriptor.Descriptor(
  name='GetTimeRequest',
  full_name='com.daml.ledger.api.v1.testing.GetTimeRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ledger_id', full_name='com.daml.ledger.api.v1.testing.GetTimeRequest.ledger_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=147,
  serialized_end=182,
)


_GETTIMERESPONSE = _descriptor.Descriptor(
  name='GetTimeResponse',
  full_name='com.daml.ledger.api.v1.testing.GetTimeResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='current_time', full_name='com.daml.ledger.api.v1.testing.GetTimeResponse.current_time', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=184,
  serialized_end=251,
)


_SETTIMEREQUEST = _descriptor.Descriptor(
  name='SetTimeRequest',
  full_name='com.daml.ledger.api.v1.testing.SetTimeRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ledger_id', full_name='com.daml.ledger.api.v1.testing.SetTimeRequest.ledger_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='current_time', full_name='com.daml.ledger.api.v1.testing.SetTimeRequest.current_time', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='new_time', full_name='com.daml.ledger.api.v1.testing.SetTimeRequest.new_time', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=254,
  serialized_end=385,
)

_GETTIMERESPONSE.fields_by_name['current_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_SETTIMEREQUEST.fields_by_name['current_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_SETTIMEREQUEST.fields_by_name['new_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
DESCRIPTOR.message_types_by_name['GetTimeRequest'] = _GETTIMEREQUEST
DESCRIPTOR.message_types_by_name['GetTimeResponse'] = _GETTIMERESPONSE
DESCRIPTOR.message_types_by_name['SetTimeRequest'] = _SETTIMEREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

GetTimeRequest = _reflection.GeneratedProtocolMessageType('GetTimeRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETTIMEREQUEST,
  '__module__' : 'com.daml.ledger.api.v1.testing.time_service_pb2'
  # @@protoc_insertion_point(class_scope:com.daml.ledger.api.v1.testing.GetTimeRequest)
  })
_sym_db.RegisterMessage(GetTimeRequest)

GetTimeResponse = _reflection.GeneratedProtocolMessageType('GetTimeResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETTIMERESPONSE,
  '__module__' : 'com.daml.ledger.api.v1.testing.time_service_pb2'
  # @@protoc_insertion_point(class_scope:com.daml.ledger.api.v1.testing.GetTimeResponse)
  })
_sym_db.RegisterMessage(GetTimeResponse)

SetTimeRequest = _reflection.GeneratedProtocolMessageType('SetTimeRequest', (_message.Message,), {
  'DESCRIPTOR' : _SETTIMEREQUEST,
  '__module__' : 'com.daml.ledger.api.v1.testing.time_service_pb2'
  # @@protoc_insertion_point(class_scope:com.daml.ledger.api.v1.testing.SetTimeRequest)
  })
_sym_db.RegisterMessage(SetTimeRequest)


DESCRIPTOR._options = None

_TIMESERVICE = _descriptor.ServiceDescriptor(
  name='TimeService',
  full_name='com.daml.ledger.api.v1.testing.TimeService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=388,
  serialized_end=594,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetTime',
    full_name='com.daml.ledger.api.v1.testing.TimeService.GetTime',
    index=0,
    containing_service=None,
    input_type=_GETTIMEREQUEST,
    output_type=_GETTIMERESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SetTime',
    full_name='com.daml.ledger.api.v1.testing.TimeService.SetTime',
    index=1,
    containing_service=None,
    input_type=_SETTIMEREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_TIMESERVICE)

DESCRIPTOR.services_by_name['TimeService'] = _TIMESERVICE

# @@protoc_insertion_point(module_scope)