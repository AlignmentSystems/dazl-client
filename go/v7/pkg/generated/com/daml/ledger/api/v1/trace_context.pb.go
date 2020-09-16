// Copyright (c) 2020 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: com/daml/ledger/api/v1/trace_context.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Data structure to propagate Zipkin trace information.
// See https://github.com/openzipkin/b3-propagation
// Trace identifiers are 64 or 128-bit, but all span identifiers within a trace are 64-bit. All identifiers are opaque.
type TraceContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If present, this is the high 64 bits of the 128-bit identifier. Otherwise the trace ID is 64 bits long.
	TraceIdHigh uint64 `protobuf:"varint,1,opt,name=trace_id_high,json=traceIdHigh,proto3" json:"trace_id_high,omitempty"`
	// The TraceId is 64 or 128-bit in length and indicates the overall ID of the trace. Every span in a trace shares this ID.
	TraceId uint64 `protobuf:"varint,2,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// The SpanId is 64-bit in length and indicates the position of the current operation in the trace tree.
	// The value should not be interpreted: it may or may not be derived from the value of the TraceId.
	SpanId uint64 `protobuf:"varint,3,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	// The ParentSpanId is 64-bit in length and indicates the position of the parent operation in the trace tree.
	// When the span is the root of the trace tree, the ParentSpanId is absent.
	ParentSpanId *wrappers.UInt64Value `protobuf:"bytes,4,opt,name=parent_span_id,json=parentSpanId,proto3" json:"parent_span_id,omitempty"`
	// When the sampled decision is accept, report this span to the tracing system. When it is reject, do not.
	// When B3 attributes are sent without a sampled decision, the receiver should make one.
	// Once the sampling decision is made, the same value should be consistently sent downstream.
	Sampled bool `protobuf:"varint,5,opt,name=sampled,proto3" json:"sampled,omitempty"`
}

func (x *TraceContext) Reset() {
	*x = TraceContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_trace_context_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceContext) ProtoMessage() {}

func (x *TraceContext) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_trace_context_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceContext.ProtoReflect.Descriptor instead.
func (*TraceContext) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_trace_context_proto_rawDescGZIP(), []int{0}
}

func (x *TraceContext) GetTraceIdHigh() uint64 {
	if x != nil {
		return x.TraceIdHigh
	}
	return 0
}

func (x *TraceContext) GetTraceId() uint64 {
	if x != nil {
		return x.TraceId
	}
	return 0
}

func (x *TraceContext) GetSpanId() uint64 {
	if x != nil {
		return x.SpanId
	}
	return 0
}

func (x *TraceContext) GetParentSpanId() *wrappers.UInt64Value {
	if x != nil {
		return x.ParentSpanId
	}
	return nil
}

func (x *TraceContext) GetSampled() bool {
	if x != nil {
		return x.Sampled
	}
	return false
}

var File_com_daml_ledger_api_v1_trace_context_proto protoreflect.FileDescriptor

var file_com_daml_ledger_api_v1_trace_context_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x01, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x26, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x5f, 0x68, 0x69, 0x67, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01,
	0x52, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x48, 0x69, 0x67, 0x68, 0x12, 0x1d, 0x0a,
	0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x02, 0x30, 0x01, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x07,
	0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30,
	0x01, 0x52, 0x06, 0x73, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x0e, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x64, 0x42, 0x82, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x64, 0x61, 0x7a,
	0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x37, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0xaa, 0x02, 0x16, 0x43, 0x6f, 0x6d, 0x2e, 0x44, 0x61, 0x6d, 0x6c, 0x2e, 0x4c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_daml_ledger_api_v1_trace_context_proto_rawDescOnce sync.Once
	file_com_daml_ledger_api_v1_trace_context_proto_rawDescData = file_com_daml_ledger_api_v1_trace_context_proto_rawDesc
)

func file_com_daml_ledger_api_v1_trace_context_proto_rawDescGZIP() []byte {
	file_com_daml_ledger_api_v1_trace_context_proto_rawDescOnce.Do(func() {
		file_com_daml_ledger_api_v1_trace_context_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_daml_ledger_api_v1_trace_context_proto_rawDescData)
	})
	return file_com_daml_ledger_api_v1_trace_context_proto_rawDescData
}

var file_com_daml_ledger_api_v1_trace_context_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_daml_ledger_api_v1_trace_context_proto_goTypes = []interface{}{
	(*TraceContext)(nil),         // 0: com.daml.ledger.api.v1.TraceContext
	(*wrappers.UInt64Value)(nil), // 1: google.protobuf.UInt64Value
}
var file_com_daml_ledger_api_v1_trace_context_proto_depIdxs = []int32{
	1, // 0: com.daml.ledger.api.v1.TraceContext.parent_span_id:type_name -> google.protobuf.UInt64Value
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_daml_ledger_api_v1_trace_context_proto_init() }
func file_com_daml_ledger_api_v1_trace_context_proto_init() {
	if File_com_daml_ledger_api_v1_trace_context_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_daml_ledger_api_v1_trace_context_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraceContext); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_daml_ledger_api_v1_trace_context_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_daml_ledger_api_v1_trace_context_proto_goTypes,
		DependencyIndexes: file_com_daml_ledger_api_v1_trace_context_proto_depIdxs,
		MessageInfos:      file_com_daml_ledger_api_v1_trace_context_proto_msgTypes,
	}.Build()
	File_com_daml_ledger_api_v1_trace_context_proto = out.File
	file_com_daml_ledger_api_v1_trace_context_proto_rawDesc = nil
	file_com_daml_ledger_api_v1_trace_context_proto_goTypes = nil
	file_com_daml_ledger_api_v1_trace_context_proto_depIdxs = nil
}