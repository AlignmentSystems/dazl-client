# Copyright (c) 2019 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: com/digitalasset/ledger/api/v1/admin/config_management_service.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='com/digitalasset/ledger/api/v1/admin/config_management_service.proto',
  package='com.digitalasset.ledger.api.v1.admin',
  syntax='proto3',
  serialized_options=_b('\n$com.digitalasset.ledger.api.v1.adminB!ConfigManagementServiceOuterClass\252\002$Com.DigitalAsset.Ledger.Api.V1.Admin'),
  serialized_pb=_b('\nDcom/digitalasset/ledger/api/v1/admin/config_management_service.proto\x12$com.digitalasset.ledger.api.v1.admin\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x15\n\x13GetTimeModelRequest\"}\n\x14GetTimeModelResponse\x12 \n\x18\x63onfiguration_generation\x18\x01 \x01(\x03\x12\x43\n\ntime_model\x18\x02 \x01(\x0b\x32/.com.digitalasset.ledger.api.v1.admin.TimeModel\"\xd0\x01\n\x13SetTimeModelRequest\x12\x15\n\rsubmission_id\x18\x01 \x01(\t\x12\x37\n\x13maximum_record_time\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12 \n\x18\x63onfiguration_generation\x18\x03 \x01(\x03\x12G\n\x0enew_time_model\x18\x04 \x01(\x0b\x32/.com.digitalasset.ledger.api.v1.admin.TimeModel\"8\n\x14SetTimeModelResponse\x12 \n\x18\x63onfiguration_generation\x18\x01 \x01(\x03\"\xa6\x01\n\tTimeModel\x12:\n\x17min_transaction_latency\x18\x01 \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x31\n\x0emax_clock_skew\x18\x02 \x01(\x0b\x32\x19.google.protobuf.Duration\x12*\n\x07max_ttl\x18\x03 \x01(\x0b\x32\x19.google.protobuf.Duration2\xa9\x02\n\x17\x43onfigManagementService\x12\x85\x01\n\x0cGetTimeModel\x12\x39.com.digitalasset.ledger.api.v1.admin.GetTimeModelRequest\x1a:.com.digitalasset.ledger.api.v1.admin.GetTimeModelResponse\x12\x85\x01\n\x0cSetTimeModel\x12\x39.com.digitalasset.ledger.api.v1.admin.SetTimeModelRequest\x1a:.com.digitalasset.ledger.api.v1.admin.SetTimeModelResponseBp\n$com.digitalasset.ledger.api.v1.adminB!ConfigManagementServiceOuterClass\xaa\x02$Com.DigitalAsset.Ledger.Api.V1.Adminb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_duration__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_GETTIMEMODELREQUEST = _descriptor.Descriptor(
  name='GetTimeModelRequest',
  full_name='com.digitalasset.ledger.api.v1.admin.GetTimeModelRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
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
  serialized_start=175,
  serialized_end=196,
)


_GETTIMEMODELRESPONSE = _descriptor.Descriptor(
  name='GetTimeModelResponse',
  full_name='com.digitalasset.ledger.api.v1.admin.GetTimeModelResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='configuration_generation', full_name='com.digitalasset.ledger.api.v1.admin.GetTimeModelResponse.configuration_generation', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='time_model', full_name='com.digitalasset.ledger.api.v1.admin.GetTimeModelResponse.time_model', index=1,
      number=2, type=11, cpp_type=10, label=1,
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
  serialized_start=198,
  serialized_end=323,
)


_SETTIMEMODELREQUEST = _descriptor.Descriptor(
  name='SetTimeModelRequest',
  full_name='com.digitalasset.ledger.api.v1.admin.SetTimeModelRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='submission_id', full_name='com.digitalasset.ledger.api.v1.admin.SetTimeModelRequest.submission_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='maximum_record_time', full_name='com.digitalasset.ledger.api.v1.admin.SetTimeModelRequest.maximum_record_time', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='configuration_generation', full_name='com.digitalasset.ledger.api.v1.admin.SetTimeModelRequest.configuration_generation', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='new_time_model', full_name='com.digitalasset.ledger.api.v1.admin.SetTimeModelRequest.new_time_model', index=3,
      number=4, type=11, cpp_type=10, label=1,
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
  serialized_start=326,
  serialized_end=534,
)


_SETTIMEMODELRESPONSE = _descriptor.Descriptor(
  name='SetTimeModelResponse',
  full_name='com.digitalasset.ledger.api.v1.admin.SetTimeModelResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='configuration_generation', full_name='com.digitalasset.ledger.api.v1.admin.SetTimeModelResponse.configuration_generation', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=536,
  serialized_end=592,
)


_TIMEMODEL = _descriptor.Descriptor(
  name='TimeModel',
  full_name='com.digitalasset.ledger.api.v1.admin.TimeModel',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='min_transaction_latency', full_name='com.digitalasset.ledger.api.v1.admin.TimeModel.min_transaction_latency', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='max_clock_skew', full_name='com.digitalasset.ledger.api.v1.admin.TimeModel.max_clock_skew', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='max_ttl', full_name='com.digitalasset.ledger.api.v1.admin.TimeModel.max_ttl', index=2,
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
  serialized_start=595,
  serialized_end=761,
)

_GETTIMEMODELRESPONSE.fields_by_name['time_model'].message_type = _TIMEMODEL
_SETTIMEMODELREQUEST.fields_by_name['maximum_record_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_SETTIMEMODELREQUEST.fields_by_name['new_time_model'].message_type = _TIMEMODEL
_TIMEMODEL.fields_by_name['min_transaction_latency'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_TIMEMODEL.fields_by_name['max_clock_skew'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_TIMEMODEL.fields_by_name['max_ttl'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
DESCRIPTOR.message_types_by_name['GetTimeModelRequest'] = _GETTIMEMODELREQUEST
DESCRIPTOR.message_types_by_name['GetTimeModelResponse'] = _GETTIMEMODELRESPONSE
DESCRIPTOR.message_types_by_name['SetTimeModelRequest'] = _SETTIMEMODELREQUEST
DESCRIPTOR.message_types_by_name['SetTimeModelResponse'] = _SETTIMEMODELRESPONSE
DESCRIPTOR.message_types_by_name['TimeModel'] = _TIMEMODEL
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

GetTimeModelRequest = _reflection.GeneratedProtocolMessageType('GetTimeModelRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETTIMEMODELREQUEST,
  '__module__' : 'com.digitalasset.ledger.api.v1.admin.config_management_service_pb2'
  # @@protoc_insertion_point(class_scope:com.digitalasset.ledger.api.v1.admin.GetTimeModelRequest)
  })
_sym_db.RegisterMessage(GetTimeModelRequest)

GetTimeModelResponse = _reflection.GeneratedProtocolMessageType('GetTimeModelResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETTIMEMODELRESPONSE,
  '__module__' : 'com.digitalasset.ledger.api.v1.admin.config_management_service_pb2'
  # @@protoc_insertion_point(class_scope:com.digitalasset.ledger.api.v1.admin.GetTimeModelResponse)
  })
_sym_db.RegisterMessage(GetTimeModelResponse)

SetTimeModelRequest = _reflection.GeneratedProtocolMessageType('SetTimeModelRequest', (_message.Message,), {
  'DESCRIPTOR' : _SETTIMEMODELREQUEST,
  '__module__' : 'com.digitalasset.ledger.api.v1.admin.config_management_service_pb2'
  # @@protoc_insertion_point(class_scope:com.digitalasset.ledger.api.v1.admin.SetTimeModelRequest)
  })
_sym_db.RegisterMessage(SetTimeModelRequest)

SetTimeModelResponse = _reflection.GeneratedProtocolMessageType('SetTimeModelResponse', (_message.Message,), {
  'DESCRIPTOR' : _SETTIMEMODELRESPONSE,
  '__module__' : 'com.digitalasset.ledger.api.v1.admin.config_management_service_pb2'
  # @@protoc_insertion_point(class_scope:com.digitalasset.ledger.api.v1.admin.SetTimeModelResponse)
  })
_sym_db.RegisterMessage(SetTimeModelResponse)

TimeModel = _reflection.GeneratedProtocolMessageType('TimeModel', (_message.Message,), {
  'DESCRIPTOR' : _TIMEMODEL,
  '__module__' : 'com.digitalasset.ledger.api.v1.admin.config_management_service_pb2'
  # @@protoc_insertion_point(class_scope:com.digitalasset.ledger.api.v1.admin.TimeModel)
  })
_sym_db.RegisterMessage(TimeModel)


DESCRIPTOR._options = None

_CONFIGMANAGEMENTSERVICE = _descriptor.ServiceDescriptor(
  name='ConfigManagementService',
  full_name='com.digitalasset.ledger.api.v1.admin.ConfigManagementService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=764,
  serialized_end=1061,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetTimeModel',
    full_name='com.digitalasset.ledger.api.v1.admin.ConfigManagementService.GetTimeModel',
    index=0,
    containing_service=None,
    input_type=_GETTIMEMODELREQUEST,
    output_type=_GETTIMEMODELRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SetTimeModel',
    full_name='com.digitalasset.ledger.api.v1.admin.ConfigManagementService.SetTimeModel',
    index=1,
    containing_service=None,
    input_type=_SETTIMEMODELREQUEST,
    output_type=_SETTIMEMODELRESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_CONFIGMANAGEMENTSERVICE)

DESCRIPTOR.services_by_name['ConfigManagementService'] = _CONFIGMANAGEMENTSERVICE

# @@protoc_insertion_point(module_scope)
