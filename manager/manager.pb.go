// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.0
// source: manager/manager.proto

package manager

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ForwardedPort string                 `protobuf:"bytes,1,opt,name=forwarded_port,json=forwardedPort,proto3" json:"forwarded_port,omitempty"`
	SvmId         string                 `protobuf:"bytes,2,opt,name=svm_id,json=svmId,proto3" json:"svm_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRes) Reset() {
	*x = CreateRes{}
	mi := &file_manager_manager_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRes) ProtoMessage() {}

func (x *CreateRes) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRes.ProtoReflect.Descriptor instead.
func (*CreateRes) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRes) GetForwardedPort() string {
	if x != nil {
		return x.ForwardedPort
	}
	return ""
}

func (x *CreateRes) GetSvmId() string {
	if x != nil {
		return x.SvmId
	}
	return ""
}

type RemoveReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SvmId         string                 `protobuf:"bytes,1,opt,name=svm_id,json=svmId,proto3" json:"svm_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveReq) Reset() {
	*x = RemoveReq{}
	mi := &file_manager_manager_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveReq) ProtoMessage() {}

func (x *RemoveReq) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveReq.ProtoReflect.Descriptor instead.
func (*RemoveReq) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{1}
}

func (x *RemoveReq) GetSvmId() string {
	if x != nil {
		return x.SvmId
	}
	return ""
}

type AttestationPolicyRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Info          []byte                 `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Id            string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AttestationPolicyRes) Reset() {
	*x = AttestationPolicyRes{}
	mi := &file_manager_manager_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AttestationPolicyRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestationPolicyRes) ProtoMessage() {}

func (x *AttestationPolicyRes) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestationPolicyRes.ProtoReflect.Descriptor instead.
func (*AttestationPolicyRes) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{2}
}

func (x *AttestationPolicyRes) GetInfo() []byte {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *AttestationPolicyRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SVMInfoRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OvmfVersion   string                 `protobuf:"bytes,2,opt,name=ovmf_version,json=ovmfVersion,proto3" json:"ovmf_version,omitempty"`
	CpuNum        int32                  `protobuf:"varint,3,opt,name=cpu_num,json=cpuNum,proto3" json:"cpu_num,omitempty"`
	CpuType       string                 `protobuf:"bytes,4,opt,name=cpu_type,json=cpuType,proto3" json:"cpu_type,omitempty"`
	KernelCmd     string                 `protobuf:"bytes,5,opt,name=kernel_cmd,json=kernelCmd,proto3" json:"kernel_cmd,omitempty"`
	EosVersion    string                 `protobuf:"bytes,6,opt,name=eos_version,json=eosVersion,proto3" json:"eos_version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SVMInfoRes) Reset() {
	*x = SVMInfoRes{}
	mi := &file_manager_manager_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SVMInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SVMInfoRes) ProtoMessage() {}

func (x *SVMInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SVMInfoRes.ProtoReflect.Descriptor instead.
func (*SVMInfoRes) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{3}
}

func (x *SVMInfoRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SVMInfoRes) GetOvmfVersion() string {
	if x != nil {
		return x.OvmfVersion
	}
	return ""
}

func (x *SVMInfoRes) GetCpuNum() int32 {
	if x != nil {
		return x.CpuNum
	}
	return 0
}

func (x *SVMInfoRes) GetCpuType() string {
	if x != nil {
		return x.CpuType
	}
	return ""
}

func (x *SVMInfoRes) GetKernelCmd() string {
	if x != nil {
		return x.KernelCmd
	}
	return ""
}

func (x *SVMInfoRes) GetEosVersion() string {
	if x != nil {
		return x.EosVersion
	}
	return ""
}

type AttestationPolicyReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AttestationPolicyReq) Reset() {
	*x = AttestationPolicyReq{}
	mi := &file_manager_manager_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AttestationPolicyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestationPolicyReq) ProtoMessage() {}

func (x *AttestationPolicyReq) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestationPolicyReq.ProtoReflect.Descriptor instead.
func (*AttestationPolicyReq) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{4}
}

func (x *AttestationPolicyReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SVMInfoReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SVMInfoReq) Reset() {
	*x = SVMInfoReq{}
	mi := &file_manager_manager_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SVMInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SVMInfoReq) ProtoMessage() {}

func (x *SVMInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_manager_manager_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SVMInfoReq.ProtoReflect.Descriptor instead.
func (*SVMInfoReq) Descriptor() ([]byte, []int) {
	return file_manager_manager_proto_rawDescGZIP(), []int{5}
}

func (x *SVMInfoReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_manager_manager_proto protoreflect.FileDescriptor

var file_manager_manager_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x76, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x76, 0x6d, 0x49, 0x64, 0x22, 0x22, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x76, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x76, 0x6d, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x14,
	0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb3, 0x01, 0x0a, 0x0a, 0x53, 0x56, 0x4d,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x76, 0x6d, 0x66, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f,
	0x76, 0x6d, 0x66, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x70,
	0x75, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x70, 0x75,
	0x4e, 0x75, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x70, 0x75, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x70, 0x75, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x63, 0x6d, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x43, 0x6d, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x65, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x65, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x26,
	0x0a, 0x14, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1c, 0x0a, 0x0a, 0x53, 0x56, 0x4d, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x32, 0x90, 0x02, 0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x56, 0x6d, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x38, 0x0a, 0x08, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x56, 0x6d, 0x12, 0x12, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x07, 0x53,
	0x56, 0x4d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x53, 0x56, 0x4d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x56, 0x4d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x53, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x6d, 0x61, 0x6e,
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

var file_manager_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_manager_manager_proto_goTypes = []any{
	(*CreateRes)(nil),            // 0: manager.CreateRes
	(*RemoveReq)(nil),            // 1: manager.RemoveReq
	(*AttestationPolicyRes)(nil), // 2: manager.AttestationPolicyRes
	(*SVMInfoRes)(nil),           // 3: manager.SVMInfoRes
	(*AttestationPolicyReq)(nil), // 4: manager.AttestationPolicyReq
	(*SVMInfoReq)(nil),           // 5: manager.SVMInfoReq
	(*emptypb.Empty)(nil),        // 6: google.protobuf.Empty
}
var file_manager_manager_proto_depIdxs = []int32{
	6, // 0: manager.ManagerService.CreateVm:input_type -> google.protobuf.Empty
	1, // 1: manager.ManagerService.RemoveVm:input_type -> manager.RemoveReq
	5, // 2: manager.ManagerService.SVMInfo:input_type -> manager.SVMInfoReq
	4, // 3: manager.ManagerService.AttestationPolicy:input_type -> manager.AttestationPolicyReq
	0, // 4: manager.ManagerService.CreateVm:output_type -> manager.CreateRes
	6, // 5: manager.ManagerService.RemoveVm:output_type -> google.protobuf.Empty
	3, // 6: manager.ManagerService.SVMInfo:output_type -> manager.SVMInfoRes
	2, // 7: manager.ManagerService.AttestationPolicy:output_type -> manager.AttestationPolicyRes
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_manager_manager_proto_init() }
func file_manager_manager_proto_init() {
	if File_manager_manager_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_manager_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
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
