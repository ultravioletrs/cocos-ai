// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: manager/manager.proto

package manager

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WhoAmI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WhoAmI) Reset() {
	*x = WhoAmI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhoAmI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoAmI) ProtoMessage() {}

func (x *WhoAmI) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoAmI.ProtoReflect.Descriptor instead.
func (*WhoAmI) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{0}
}

type RunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentPort     string `protobuf:"bytes,1,opt,name=agent_port,json=agentPort,proto3" json:"agent_port,omitempty"`
	ComputationId string `protobuf:"bytes,2,opt,name=computation_id,json=computationId,proto3" json:"computation_id,omitempty"`
}

func (x *RunResponse) Reset() {
	*x = RunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunResponse) ProtoMessage() {}

func (x *RunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunResponse.ProtoReflect.Descriptor instead.
func (*RunResponse) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{1}
}

func (x *RunResponse) GetAgentPort() string {
	if x != nil {
		return x.AgentPort
	}
	return ""
}

func (x *RunResponse) GetComputationId() string {
	if x != nil {
		return x.ComputationId
	}
	return ""
}

type AgentEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType     string                 `protobuf:"bytes,1,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ComputationId string                 `protobuf:"bytes,3,opt,name=computation_id,json=computationId,proto3" json:"computation_id,omitempty"`
	Details       []byte                 `protobuf:"bytes,4,opt,name=details,proto3" json:"details,omitempty"`
	Originator    string                 `protobuf:"bytes,5,opt,name=originator,proto3" json:"originator,omitempty"`
	Status        string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AgentEvent) Reset() {
	*x = AgentEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentEvent) ProtoMessage() {}

func (x *AgentEvent) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentEvent.ProtoReflect.Descriptor instead.
func (*AgentEvent) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{2}
}

func (x *AgentEvent) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *AgentEvent) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *AgentEvent) GetComputationId() string {
	if x != nil {
		return x.ComputationId
	}
	return ""
}

func (x *AgentEvent) GetDetails() []byte {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *AgentEvent) GetOriginator() string {
	if x != nil {
		return x.Originator
	}
	return ""
}

func (x *AgentEvent) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type AgentLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogMessage    string `protobuf:"bytes,1,opt,name=log_message,json=logMessage,proto3" json:"log_message,omitempty"`
	ComputationId string `protobuf:"bytes,2,opt,name=computation_id,json=computationId,proto3" json:"computation_id,omitempty"`
}

func (x *AgentLog) Reset() {
	*x = AgentLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentLog) ProtoMessage() {}

func (x *AgentLog) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentLog.ProtoReflect.Descriptor instead.
func (*AgentLog) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{3}
}

func (x *AgentLog) GetLogMessage() string {
	if x != nil {
		return x.LogMessage
	}
	return ""
}

func (x *AgentLog) GetComputationId() string {
	if x != nil {
		return x.ComputationId
	}
	return ""
}

type ClientStreamMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*ClientStreamMessage_Whoami
	//	*ClientStreamMessage_AgentLog
	//	*ClientStreamMessage_AgentEvent
	//	*ClientStreamMessage_RunRes
	Message isClientStreamMessage_Message `protobuf_oneof:"message"`
}

func (x *ClientStreamMessage) Reset() {
	*x = ClientStreamMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamMessage) ProtoMessage() {}

func (x *ClientStreamMessage) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStreamMessage.ProtoReflect.Descriptor instead.
func (*ClientStreamMessage) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{4}
}

func (m *ClientStreamMessage) GetMessage() isClientStreamMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *ClientStreamMessage) GetWhoami() *WhoAmI {
	if x, ok := x.GetMessage().(*ClientStreamMessage_Whoami); ok {
		return x.Whoami
	}
	return nil
}

func (x *ClientStreamMessage) GetAgentLog() *AgentLog {
	if x, ok := x.GetMessage().(*ClientStreamMessage_AgentLog); ok {
		return x.AgentLog
	}
	return nil
}

func (x *ClientStreamMessage) GetAgentEvent() *AgentEvent {
	if x, ok := x.GetMessage().(*ClientStreamMessage_AgentEvent); ok {
		return x.AgentEvent
	}
	return nil
}

func (x *ClientStreamMessage) GetRunRes() *RunResponse {
	if x, ok := x.GetMessage().(*ClientStreamMessage_RunRes); ok {
		return x.RunRes
	}
	return nil
}

type isClientStreamMessage_Message interface {
	isClientStreamMessage_Message()
}

type ClientStreamMessage_Whoami struct {
	Whoami *WhoAmI `protobuf:"bytes,1,opt,name=whoami,proto3,oneof"`
}

type ClientStreamMessage_AgentLog struct {
	AgentLog *AgentLog `protobuf:"bytes,2,opt,name=agent_log,json=agentLog,proto3,oneof"`
}

type ClientStreamMessage_AgentEvent struct {
	AgentEvent *AgentEvent `protobuf:"bytes,3,opt,name=agent_event,json=agentEvent,proto3,oneof"`
}

type ClientStreamMessage_RunRes struct {
	RunRes *RunResponse `protobuf:"bytes,4,opt,name=run_res,json=runRes,proto3,oneof"`
}

func (*ClientStreamMessage_Whoami) isClientStreamMessage_Message() {}

func (*ClientStreamMessage_AgentLog) isClientStreamMessage_Message() {}

func (*ClientStreamMessage_AgentEvent) isClientStreamMessage_Message() {}

func (*ClientStreamMessage_RunRes) isClientStreamMessage_Message() {}

type ComputationRunReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description     string       `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Datasets        []*Dataset   `protobuf:"bytes,4,rep,name=datasets,proto3" json:"datasets,omitempty"`
	Algorithms      []*Algorithm `protobuf:"bytes,5,rep,name=algorithms,proto3" json:"algorithms,omitempty"`
	ResultConsumers []string     `protobuf:"bytes,6,rep,name=result_consumers,json=resultConsumers,proto3" json:"result_consumers,omitempty"`
	AgentConfig     *AgentConfig `protobuf:"bytes,7,opt,name=agent_config,json=agentConfig,proto3" json:"agent_config,omitempty"`
}

func (x *ComputationRunReq) Reset() {
	*x = ComputationRunReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputationRunReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputationRunReq) ProtoMessage() {}

func (x *ComputationRunReq) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputationRunReq.ProtoReflect.Descriptor instead.
func (*ComputationRunReq) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{5}
}

func (x *ComputationRunReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ComputationRunReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ComputationRunReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ComputationRunReq) GetDatasets() []*Dataset {
	if x != nil {
		return x.Datasets
	}
	return nil
}

func (x *ComputationRunReq) GetAlgorithms() []*Algorithm {
	if x != nil {
		return x.Algorithms
	}
	return nil
}

func (x *ComputationRunReq) GetResultConsumers() []string {
	if x != nil {
		return x.ResultConsumers
	}
	return nil
}

func (x *ComputationRunReq) GetAgentConfig() *AgentConfig {
	if x != nil {
		return x.AgentConfig
	}
	return nil
}

type Dataset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Dataset) Reset() {
	*x = Dataset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dataset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dataset) ProtoMessage() {}

func (x *Dataset) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dataset.ProtoReflect.Descriptor instead.
func (*Dataset) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{6}
}

func (x *Dataset) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *Dataset) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Algorithm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Algorithm) Reset() {
	*x = Algorithm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Algorithm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Algorithm) ProtoMessage() {}

func (x *Algorithm) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Algorithm.ProtoReflect.Descriptor instead.
func (*Algorithm) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{7}
}

func (x *Algorithm) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *Algorithm) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AgentConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port         string `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	Host         string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	CertFile     string `protobuf:"bytes,3,opt,name=cert_file,json=certFile,proto3" json:"cert_file,omitempty"`
	KeyFile      string `protobuf:"bytes,4,opt,name=key_file,json=keyFile,proto3" json:"key_file,omitempty"`
	ClientCaFile string `protobuf:"bytes,5,opt,name=client_ca_file,json=clientCaFile,proto3" json:"client_ca_file,omitempty"`
	ServerCaFile string `protobuf:"bytes,6,opt,name=server_ca_file,json=serverCaFile,proto3" json:"server_ca_file,omitempty"`
	LogLevel     string `protobuf:"bytes,7,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	InstanceId   string `protobuf:"bytes,8,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
}

func (x *AgentConfig) Reset() {
	*x = AgentConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_manager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentConfig) ProtoMessage() {}

func (x *AgentConfig) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentConfig.ProtoReflect.Descriptor instead.
func (*AgentConfig) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{8}
}

func (x *AgentConfig) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *AgentConfig) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *AgentConfig) GetCertFile() string {
	if x != nil {
		return x.CertFile
	}
	return ""
}

func (x *AgentConfig) GetKeyFile() string {
	if x != nil {
		return x.KeyFile
	}
	return ""
}

func (x *AgentConfig) GetClientCaFile() string {
	if x != nil {
		return x.ClientCaFile
	}
	return ""
}

func (x *AgentConfig) GetServerCaFile() string {
	if x != nil {
		return x.ServerCaFile
	}
	return ""
}

func (x *AgentConfig) GetLogLevel() string {
	if x != nil {
		return x.LogLevel
	}
	return ""
}

func (x *AgentConfig) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

var File_manager_manager_proto protoreflect.FileDescriptor

var file_manager_manager_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x08, 0x0a, 0x06, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x22, 0x53, 0x0a, 0x0b, 0x52,
	0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x22, 0xde, 0x01, 0x0a, 0x0a, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x52, 0x0a, 0x08, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x1f, 0x0a,
	0x0b, 0x6c, 0x6f, 0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xe6, 0x01, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a,
	0x06, 0x77, 0x68, 0x6f, 0x61, 0x6d, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x48, 0x00,
	0x52, 0x06, 0x77, 0x68, 0x6f, 0x61, 0x6d, 0x69, 0x12, 0x30, 0x0a, 0x09, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x48, 0x00,
	0x52, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x36, 0x0a, 0x0b, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x75,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x06, 0x72, 0x75, 0x6e,
	0x52, 0x65, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x9f,
	0x02, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75,
	0x6e, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x61, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x52, 0x0a, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x12, 0x29, 0x0a, 0x10,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x12, 0x37, 0x0a, 0x0c, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0x35, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x09, 0x41, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xf7, 0x01, 0x0a, 0x0b, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x65, 0x72, 0x74,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x65, 0x72,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x43, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x63, 0x61, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x32, 0x5b, 0x0a, 0x0e, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x07,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6e, 0x52, 0x65,
	0x71, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_manager_manager_proto_rawDescOnce sync.Once
	file_manager_manager_proto_rawDescData = file_manager_manager_proto_rawDesc
)

func file_manager_manager_proto_rawDescGZIP() []byte {
	file_manager_manager_proto_rawDescOnce.Do(func() {
		file_manager_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_manager_manager_proto_rawDescData)
	})
	return file_manager_manager_proto_rawDescData
}

var file_manager_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_manager_manager_proto_goTypes = []interface{}{
	(*WhoAmI)(nil),                // 0: manager.WhoAmI
	(*RunResponse)(nil),           // 1: manager.RunResponse
	(*AgentEvent)(nil),            // 2: manager.AgentEvent
	(*AgentLog)(nil),              // 3: manager.AgentLog
	(*ClientStreamMessage)(nil),   // 4: manager.ClientStreamMessage
	(*ComputationRunReq)(nil),     // 5: manager.ComputationRunReq
	(*Dataset)(nil),               // 6: manager.Dataset
	(*Algorithm)(nil),             // 7: manager.Algorithm
	(*AgentConfig)(nil),           // 8: manager.AgentConfig
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_manager_manager_proto_depIdxs = []int32{
	9, // 0: manager.AgentEvent.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: manager.ClientStreamMessage.whoami:type_name -> manager.WhoAmI
	3, // 2: manager.ClientStreamMessage.agent_log:type_name -> manager.AgentLog
	2, // 3: manager.ClientStreamMessage.agent_event:type_name -> manager.AgentEvent
	1, // 4: manager.ClientStreamMessage.run_res:type_name -> manager.RunResponse
	6, // 5: manager.ComputationRunReq.datasets:type_name -> manager.Dataset
	7, // 6: manager.ComputationRunReq.algorithms:type_name -> manager.Algorithm
	8, // 7: manager.ComputationRunReq.agent_config:type_name -> manager.AgentConfig
	4, // 8: manager.ManagerService.Process:input_type -> manager.ClientStreamMessage
	5, // 9: manager.ManagerService.Process:output_type -> manager.ComputationRunReq
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_manager_manager_proto_init() }
func file_manager_manager_proto_init() {
	if File_manager_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manager_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhoAmI); i {
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
		file_manager_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunResponse); i {
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
		file_manager_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentEvent); i {
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
		file_manager_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentLog); i {
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
		file_manager_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStreamMessage); i {
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
		file_manager_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputationRunReq); i {
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
		file_manager_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dataset); i {
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
		file_manager_manager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Algorithm); i {
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
		file_manager_manager_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentConfig); i {
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
	file_manager_manager_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ClientStreamMessage_Whoami)(nil),
		(*ClientStreamMessage_AgentLog)(nil),
		(*ClientStreamMessage_AgentEvent)(nil),
		(*ClientStreamMessage_RunRes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_manager_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_manager_manager_proto_goTypes,
		DependencyIndexes: file_manager_manager_proto_depIdxs,
		MessageInfos:      file_manager_manager_proto_msgTypes,
	}.Build()
	File_manager_manager_proto = out.File
	file_manager_manager_proto_rawDesc = nil
	file_manager_manager_proto_goTypes = nil
	file_manager_manager_proto_depIdxs = nil
}
