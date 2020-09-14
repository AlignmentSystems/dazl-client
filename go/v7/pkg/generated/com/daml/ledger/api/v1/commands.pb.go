// Copyright (c) 2020 Digital Asset (Switzerland) GmbH and/or its affiliates. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: com/daml/ledger/api/v1/commands.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// A composite command that groups multiple commands together.
type Commands struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Must correspond to the ledger ID reported by the Ledger Identification Service.
	// Must be a valid LedgerString (as described in ``value.proto``).
	// Required
	LedgerId string `protobuf:"bytes,1,opt,name=ledger_id,json=ledgerId,proto3" json:"ledger_id,omitempty"`
	// Identifier of the on-ledger workflow that this command is a part of.
	// Must be a valid LedgerString (as described in ``value.proto``).
	// Optional
	WorkflowId string `protobuf:"bytes,2,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	// Uniquely identifies the application (or its part) that issued the command. This is used in tracing
	// across different components and to let applications subscribe to their own submissions only.
	// Must be a valid LedgerString (as described in ``value.proto``).
	// Required
	ApplicationId string `protobuf:"bytes,3,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	// Uniquely identified the command. This identifier should be unique for each new command within an
	// application domain, i.e., the triple (application_id, party, command_id) must be unique.
	// It can be used for matching the requests with their respective completions.
	// Must be a valid LedgerString (as described in ``value.proto``).
	// Required
	CommandId string `protobuf:"bytes,4,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	// Party on whose behalf the command should be executed. It is up to the server to verify that the
	// authorisation can be granted and that the connection has been authenticated for that party.
	// Must be a valid PartyIdString (as described in ``value.proto``).
	// Required
	Party string `protobuf:"bytes,5,opt,name=party,proto3" json:"party,omitempty"`
	// Individual elements of this atomic command. Must be non-empty.
	// Required
	Commands []*Command `protobuf:"bytes,8,rep,name=commands,proto3" json:"commands,omitempty"`
	// The length of the time window during which all commands with the same party and command ID will be deduplicated.
	// Duplicate commands submitted before the end of this window return an ALREADY_EXISTS error.
	// Optional
	DeduplicationTime *duration.Duration `protobuf:"bytes,9,opt,name=deduplication_time,json=deduplicationTime,proto3" json:"deduplication_time,omitempty"`
	// Lower bound for the ledger time assigned to the resulting transaction.
	// Note: The ledger time of a transaction is assigned as part of command interpretation.
	// Use this property if you expect that command interpretation will take a considerate amount of time, such that by
	// the time the resulting transaction is sequenced, its assigned ledger time is not valid anymore.
	// Must not be set at the same time as min_ledger_time_rel.
	// Optional
	MinLedgerTimeAbs *timestamp.Timestamp `protobuf:"bytes,10,opt,name=min_ledger_time_abs,json=minLedgerTimeAbs,proto3" json:"min_ledger_time_abs,omitempty"`
	// Same as min_ledger_time_abs, but specified as a duration, starting from the time the command is received by the server.
	// Must not be set at the same time as min_ledger_time_abs.
	// Optional
	MinLedgerTimeRel *duration.Duration `protobuf:"bytes,11,opt,name=min_ledger_time_rel,json=minLedgerTimeRel,proto3" json:"min_ledger_time_rel,omitempty"`
}

func (x *Commands) Reset() {
	*x = Commands{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_commands_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commands) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commands) ProtoMessage() {}

func (x *Commands) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_commands_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commands.ProtoReflect.Descriptor instead.
func (*Commands) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_commands_proto_rawDescGZIP(), []int{0}
}

func (x *Commands) GetLedgerId() string {
	if x != nil {
		return x.LedgerId
	}
	return ""
}

func (x *Commands) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *Commands) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *Commands) GetCommandId() string {
	if x != nil {
		return x.CommandId
	}
	return ""
}

func (x *Commands) GetParty() string {
	if x != nil {
		return x.Party
	}
	return ""
}

func (x *Commands) GetCommands() []*Command {
	if x != nil {
		return x.Commands
	}
	return nil
}

func (x *Commands) GetDeduplicationTime() *duration.Duration {
	if x != nil {
		return x.DeduplicationTime
	}
	return nil
}

func (x *Commands) GetMinLedgerTimeAbs() *timestamp.Timestamp {
	if x != nil {
		return x.MinLedgerTimeAbs
	}
	return nil
}

func (x *Commands) GetMinLedgerTimeRel() *duration.Duration {
	if x != nil {
		return x.MinLedgerTimeRel
	}
	return nil
}

// A command can either create a new contract or exercise a choice on an existing contract.
type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Command:
	//	*Command_Create
	//	*Command_Exercise
	//	*Command_ExerciseByKey
	//	*Command_CreateAndExercise
	Command isCommand_Command `protobuf_oneof:"command"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_commands_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_commands_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_commands_proto_rawDescGZIP(), []int{1}
}

func (m *Command) GetCommand() isCommand_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (x *Command) GetCreate() *CreateCommand {
	if x, ok := x.GetCommand().(*Command_Create); ok {
		return x.Create
	}
	return nil
}

func (x *Command) GetExercise() *ExerciseCommand {
	if x, ok := x.GetCommand().(*Command_Exercise); ok {
		return x.Exercise
	}
	return nil
}

func (x *Command) GetExerciseByKey() *ExerciseByKeyCommand {
	if x, ok := x.GetCommand().(*Command_ExerciseByKey); ok {
		return x.ExerciseByKey
	}
	return nil
}

func (x *Command) GetCreateAndExercise() *CreateAndExerciseCommand {
	if x, ok := x.GetCommand().(*Command_CreateAndExercise); ok {
		return x.CreateAndExercise
	}
	return nil
}

type isCommand_Command interface {
	isCommand_Command()
}

type Command_Create struct {
	Create *CreateCommand `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type Command_Exercise struct {
	Exercise *ExerciseCommand `protobuf:"bytes,2,opt,name=exercise,proto3,oneof"`
}

type Command_ExerciseByKey struct {
	ExerciseByKey *ExerciseByKeyCommand `protobuf:"bytes,4,opt,name=exerciseByKey,proto3,oneof"`
}

type Command_CreateAndExercise struct {
	CreateAndExercise *CreateAndExerciseCommand `protobuf:"bytes,3,opt,name=createAndExercise,proto3,oneof"`
}

func (*Command_Create) isCommand_Command() {}

func (*Command_Exercise) isCommand_Command() {}

func (*Command_ExerciseByKey) isCommand_Command() {}

func (*Command_CreateAndExercise) isCommand_Command() {}

// Create a new contract instance based on a template.
type CreateCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The template of contract the client wants to create.
	// Required
	TemplateId *Identifier `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	// The arguments required for creating a contract from this template.
	// Required
	CreateArguments *Record `protobuf:"bytes,2,opt,name=create_arguments,json=createArguments,proto3" json:"create_arguments,omitempty"`
}

func (x *CreateCommand) Reset() {
	*x = CreateCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_commands_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommand) ProtoMessage() {}

func (x *CreateCommand) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_commands_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommand.ProtoReflect.Descriptor instead.
func (*CreateCommand) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_commands_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCommand) GetTemplateId() *Identifier {
	if x != nil {
		return x.TemplateId
	}
	return nil
}

func (x *CreateCommand) GetCreateArguments() *Record {
	if x != nil {
		return x.CreateArguments
	}
	return nil
}

// Exercise a choice on an existing contract.
type ExerciseCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The template of contract the client wants to exercise.
	// Required
	TemplateId *Identifier `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	// The ID of the contract the client wants to exercise upon.
	// Must be a valid LedgerString (as described in ``value.proto``).
	// Required
	ContractId string `protobuf:"bytes,2,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`
	// The name of the choice the client wants to exercise.
	// Must be a valid NameString (as described in ``value.proto``)
	// Required
	Choice string `protobuf:"bytes,3,opt,name=choice,proto3" json:"choice,omitempty"`
	// The argument for this choice.
	// Required
	ChoiceArgument *Value `protobuf:"bytes,4,opt,name=choice_argument,json=choiceArgument,proto3" json:"choice_argument,omitempty"`
}

func (x *ExerciseCommand) Reset() {
	*x = ExerciseCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_commands_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExerciseCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExerciseCommand) ProtoMessage() {}

func (x *ExerciseCommand) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_commands_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExerciseCommand.ProtoReflect.Descriptor instead.
func (*ExerciseCommand) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_commands_proto_rawDescGZIP(), []int{3}
}

func (x *ExerciseCommand) GetTemplateId() *Identifier {
	if x != nil {
		return x.TemplateId
	}
	return nil
}

func (x *ExerciseCommand) GetContractId() string {
	if x != nil {
		return x.ContractId
	}
	return ""
}

func (x *ExerciseCommand) GetChoice() string {
	if x != nil {
		return x.Choice
	}
	return ""
}

func (x *ExerciseCommand) GetChoiceArgument() *Value {
	if x != nil {
		return x.ChoiceArgument
	}
	return nil
}

// Exercise a choice on an existing contract specified by its key.
type ExerciseByKeyCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The template of contract the client wants to exercise.
	// Required
	TemplateId *Identifier `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	// The key of the contract the client wants to exercise upon.
	// Required
	ContractKey *Value `protobuf:"bytes,2,opt,name=contract_key,json=contractKey,proto3" json:"contract_key,omitempty"`
	// The name of the choice the client wants to exercise.
	// Must be a valid NameString (as described in ``value.proto``)
	// Required
	Choice string `protobuf:"bytes,3,opt,name=choice,proto3" json:"choice,omitempty"`
	// The argument for this choice.
	// Required
	ChoiceArgument *Value `protobuf:"bytes,4,opt,name=choice_argument,json=choiceArgument,proto3" json:"choice_argument,omitempty"`
}

func (x *ExerciseByKeyCommand) Reset() {
	*x = ExerciseByKeyCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_commands_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExerciseByKeyCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExerciseByKeyCommand) ProtoMessage() {}

func (x *ExerciseByKeyCommand) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_commands_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExerciseByKeyCommand.ProtoReflect.Descriptor instead.
func (*ExerciseByKeyCommand) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_commands_proto_rawDescGZIP(), []int{4}
}

func (x *ExerciseByKeyCommand) GetTemplateId() *Identifier {
	if x != nil {
		return x.TemplateId
	}
	return nil
}

func (x *ExerciseByKeyCommand) GetContractKey() *Value {
	if x != nil {
		return x.ContractKey
	}
	return nil
}

func (x *ExerciseByKeyCommand) GetChoice() string {
	if x != nil {
		return x.Choice
	}
	return ""
}

func (x *ExerciseByKeyCommand) GetChoiceArgument() *Value {
	if x != nil {
		return x.ChoiceArgument
	}
	return nil
}

// Create a contract and exercise a choice on it in the same transaction.
type CreateAndExerciseCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The template of the contract the client wants to create.
	// Required
	TemplateId *Identifier `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	// The arguments required for creating a contract from this template.
	// Required
	CreateArguments *Record `protobuf:"bytes,2,opt,name=create_arguments,json=createArguments,proto3" json:"create_arguments,omitempty"`
	// The name of the choice the client wants to exercise.
	// Must be a valid NameString (as described in ``value.proto``).
	// Required
	Choice string `protobuf:"bytes,3,opt,name=choice,proto3" json:"choice,omitempty"`
	// The argument for this choice.
	// Required
	ChoiceArgument *Value `protobuf:"bytes,4,opt,name=choice_argument,json=choiceArgument,proto3" json:"choice_argument,omitempty"`
}

func (x *CreateAndExerciseCommand) Reset() {
	*x = CreateAndExerciseCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_daml_ledger_api_v1_commands_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAndExerciseCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAndExerciseCommand) ProtoMessage() {}

func (x *CreateAndExerciseCommand) ProtoReflect() protoreflect.Message {
	mi := &file_com_daml_ledger_api_v1_commands_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAndExerciseCommand.ProtoReflect.Descriptor instead.
func (*CreateAndExerciseCommand) Descriptor() ([]byte, []int) {
	return file_com_daml_ledger_api_v1_commands_proto_rawDescGZIP(), []int{5}
}

func (x *CreateAndExerciseCommand) GetTemplateId() *Identifier {
	if x != nil {
		return x.TemplateId
	}
	return nil
}

func (x *CreateAndExerciseCommand) GetCreateArguments() *Record {
	if x != nil {
		return x.CreateArguments
	}
	return nil
}

func (x *CreateAndExerciseCommand) GetChoice() string {
	if x != nil {
		return x.Choice
	}
	return ""
}

func (x *CreateAndExerciseCommand) GetChoiceArgument() *Value {
	if x != nil {
		return x.ChoiceArgument
	}
	return nil
}

var File_com_daml_ledger_api_v1_commands_proto protoreflect.FileDescriptor

var file_com_daml_ledger_api_v1_commands_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d,
	0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a,
	0x22, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x03, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x08,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x48, 0x0a, 0x12, 0x64, 0x65, 0x64, 0x75,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x11, 0x64, 0x65, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x49, 0x0a, 0x13, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x61, 0x62, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x6d, 0x69, 0x6e,
	0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x41, 0x62, 0x73, 0x12, 0x48, 0x0a,
	0x13, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x72, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x6d, 0x69, 0x6e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6c, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x4a, 0x04, 0x08,
	0x07, 0x10, 0x08, 0x22, 0xd4, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x3f, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x45, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x72,
	0x63, 0x69, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x08, 0x65,
	0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x65, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65,
	0x42, 0x79, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x0d,
	0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x60, 0x0a,
	0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69,
	0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x45, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x11, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x42,
	0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x9f, 0x01, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x43, 0x0a, 0x0b,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x49, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x72, 0x67, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xd7, 0x01, 0x0a,
	0x0f, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x43, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c,
	0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x46,
	0x0a, 0x0f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61,
	0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x41, 0x72,
	0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xfd, 0x01, 0x0a, 0x14, 0x45, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x43, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x46,
	0x0a, 0x0f, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61,
	0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x41, 0x72,
	0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x8a, 0x02, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x6e, 0x64, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x43, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x63,
	0x68, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c, 0x2e,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0e, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x41, 0x72, 0x67, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x82, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x6d, 0x6c,
	0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x4f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x61, 0x6c, 0x2d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x64, 0x61, 0x7a, 0x6c, 0x2d, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x37, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6d,
	0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0xaa,
	0x02, 0x16, 0x43, 0x6f, 0x6d, 0x2e, 0x44, 0x61, 0x6d, 0x6c, 0x2e, 0x4c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_daml_ledger_api_v1_commands_proto_rawDescOnce sync.Once
	file_com_daml_ledger_api_v1_commands_proto_rawDescData = file_com_daml_ledger_api_v1_commands_proto_rawDesc
)

func file_com_daml_ledger_api_v1_commands_proto_rawDescGZIP() []byte {
	file_com_daml_ledger_api_v1_commands_proto_rawDescOnce.Do(func() {
		file_com_daml_ledger_api_v1_commands_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_daml_ledger_api_v1_commands_proto_rawDescData)
	})
	return file_com_daml_ledger_api_v1_commands_proto_rawDescData
}

var file_com_daml_ledger_api_v1_commands_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_com_daml_ledger_api_v1_commands_proto_goTypes = []interface{}{
	(*Commands)(nil),                 // 0: com.daml.ledger.api.v1.Commands
	(*Command)(nil),                  // 1: com.daml.ledger.api.v1.Command
	(*CreateCommand)(nil),            // 2: com.daml.ledger.api.v1.CreateCommand
	(*ExerciseCommand)(nil),          // 3: com.daml.ledger.api.v1.ExerciseCommand
	(*ExerciseByKeyCommand)(nil),     // 4: com.daml.ledger.api.v1.ExerciseByKeyCommand
	(*CreateAndExerciseCommand)(nil), // 5: com.daml.ledger.api.v1.CreateAndExerciseCommand
	(*duration.Duration)(nil),        // 6: google.protobuf.Duration
	(*timestamp.Timestamp)(nil),      // 7: google.protobuf.Timestamp
	(*Identifier)(nil),               // 8: com.daml.ledger.api.v1.Identifier
	(*Record)(nil),                   // 9: com.daml.ledger.api.v1.Record
	(*Value)(nil),                    // 10: com.daml.ledger.api.v1.Value
}
var file_com_daml_ledger_api_v1_commands_proto_depIdxs = []int32{
	1,  // 0: com.daml.ledger.api.v1.Commands.commands:type_name -> com.daml.ledger.api.v1.Command
	6,  // 1: com.daml.ledger.api.v1.Commands.deduplication_time:type_name -> google.protobuf.Duration
	7,  // 2: com.daml.ledger.api.v1.Commands.min_ledger_time_abs:type_name -> google.protobuf.Timestamp
	6,  // 3: com.daml.ledger.api.v1.Commands.min_ledger_time_rel:type_name -> google.protobuf.Duration
	2,  // 4: com.daml.ledger.api.v1.Command.create:type_name -> com.daml.ledger.api.v1.CreateCommand
	3,  // 5: com.daml.ledger.api.v1.Command.exercise:type_name -> com.daml.ledger.api.v1.ExerciseCommand
	4,  // 6: com.daml.ledger.api.v1.Command.exerciseByKey:type_name -> com.daml.ledger.api.v1.ExerciseByKeyCommand
	5,  // 7: com.daml.ledger.api.v1.Command.createAndExercise:type_name -> com.daml.ledger.api.v1.CreateAndExerciseCommand
	8,  // 8: com.daml.ledger.api.v1.CreateCommand.template_id:type_name -> com.daml.ledger.api.v1.Identifier
	9,  // 9: com.daml.ledger.api.v1.CreateCommand.create_arguments:type_name -> com.daml.ledger.api.v1.Record
	8,  // 10: com.daml.ledger.api.v1.ExerciseCommand.template_id:type_name -> com.daml.ledger.api.v1.Identifier
	10, // 11: com.daml.ledger.api.v1.ExerciseCommand.choice_argument:type_name -> com.daml.ledger.api.v1.Value
	8,  // 12: com.daml.ledger.api.v1.ExerciseByKeyCommand.template_id:type_name -> com.daml.ledger.api.v1.Identifier
	10, // 13: com.daml.ledger.api.v1.ExerciseByKeyCommand.contract_key:type_name -> com.daml.ledger.api.v1.Value
	10, // 14: com.daml.ledger.api.v1.ExerciseByKeyCommand.choice_argument:type_name -> com.daml.ledger.api.v1.Value
	8,  // 15: com.daml.ledger.api.v1.CreateAndExerciseCommand.template_id:type_name -> com.daml.ledger.api.v1.Identifier
	9,  // 16: com.daml.ledger.api.v1.CreateAndExerciseCommand.create_arguments:type_name -> com.daml.ledger.api.v1.Record
	10, // 17: com.daml.ledger.api.v1.CreateAndExerciseCommand.choice_argument:type_name -> com.daml.ledger.api.v1.Value
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_com_daml_ledger_api_v1_commands_proto_init() }
func file_com_daml_ledger_api_v1_commands_proto_init() {
	if File_com_daml_ledger_api_v1_commands_proto != nil {
		return
	}
	file_com_daml_ledger_api_v1_value_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_daml_ledger_api_v1_commands_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commands); i {
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
		file_com_daml_ledger_api_v1_commands_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_com_daml_ledger_api_v1_commands_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommand); i {
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
		file_com_daml_ledger_api_v1_commands_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExerciseCommand); i {
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
		file_com_daml_ledger_api_v1_commands_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExerciseByKeyCommand); i {
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
		file_com_daml_ledger_api_v1_commands_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAndExerciseCommand); i {
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
	file_com_daml_ledger_api_v1_commands_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Command_Create)(nil),
		(*Command_Exercise)(nil),
		(*Command_ExerciseByKey)(nil),
		(*Command_CreateAndExercise)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_daml_ledger_api_v1_commands_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_daml_ledger_api_v1_commands_proto_goTypes,
		DependencyIndexes: file_com_daml_ledger_api_v1_commands_proto_depIdxs,
		MessageInfos:      file_com_daml_ledger_api_v1_commands_proto_msgTypes,
	}.Build()
	File_com_daml_ledger_api_v1_commands_proto = out.File
	file_com_daml_ledger_api_v1_commands_proto_rawDesc = nil
	file_com_daml_ledger_api_v1_commands_proto_goTypes = nil
	file_com_daml_ledger_api_v1_commands_proto_depIdxs = nil
}
