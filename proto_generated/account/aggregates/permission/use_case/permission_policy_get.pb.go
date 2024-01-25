// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/account/aggregates/permission/use_case/permission_policy_get.proto

package protopermissionusecase

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	messaging "vl-account/proto_generated/common/messaging"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PermissionPolicyGetUseCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction                 `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Client      *messaging.Client                      `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Data        *PermissionPolicyGetUseCaseRequestData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PermissionPolicyGetUseCaseRequest) Reset() {
	*x = PermissionPolicyGetUseCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionPolicyGetUseCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionPolicyGetUseCaseRequest) ProtoMessage() {}

func (x *PermissionPolicyGetUseCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionPolicyGetUseCaseRequest.ProtoReflect.Descriptor instead.
func (*PermissionPolicyGetUseCaseRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_rawDescGZIP(), []int{0}
}

func (x *PermissionPolicyGetUseCaseRequest) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *PermissionPolicyGetUseCaseRequest) GetClient() *messaging.Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *PermissionPolicyGetUseCaseRequest) GetData() *PermissionPolicyGetUseCaseRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type PermissionPolicyGetUseCaseRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PermissionPolicyGetUseCaseRequestData) Reset() {
	*x = PermissionPolicyGetUseCaseRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionPolicyGetUseCaseRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionPolicyGetUseCaseRequestData) ProtoMessage() {}

func (x *PermissionPolicyGetUseCaseRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionPolicyGetUseCaseRequestData.ProtoReflect.Descriptor instead.
func (*PermissionPolicyGetUseCaseRequestData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_rawDescGZIP(), []int{1}
}

func (x *PermissionPolicyGetUseCaseRequestData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PermissionPolicyGetUseCaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction                  `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Error       *messaging.Error                        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data        *PermissionPolicyGetUseCaseResponseData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PermissionPolicyGetUseCaseResponse) Reset() {
	*x = PermissionPolicyGetUseCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionPolicyGetUseCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionPolicyGetUseCaseResponse) ProtoMessage() {}

func (x *PermissionPolicyGetUseCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionPolicyGetUseCaseResponse.ProtoReflect.Descriptor instead.
func (*PermissionPolicyGetUseCaseResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_rawDescGZIP(), []int{2}
}

func (x *PermissionPolicyGetUseCaseResponse) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *PermissionPolicyGetUseCaseResponse) GetError() *messaging.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *PermissionPolicyGetUseCaseResponse) GetData() *PermissionPolicyGetUseCaseResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type PermissionPolicyGetUseCaseResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PermissionPolicy *PermissionPolicy `protobuf:"bytes,1,opt,name=permission_policy,json=permissionPolicy,proto3" json:"permission_policy,omitempty"`
}

func (x *PermissionPolicyGetUseCaseResponseData) Reset() {
	*x = PermissionPolicyGetUseCaseResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionPolicyGetUseCaseResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionPolicyGetUseCaseResponseData) ProtoMessage() {}

func (x *PermissionPolicyGetUseCaseResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionPolicyGetUseCaseResponseData.ProtoReflect.Descriptor instead.
func (*PermissionPolicyGetUseCaseResponseData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_rawDescGZIP(), []int{3}
}

func (x *PermissionPolicyGetUseCaseResponseData) GetPermissionPolicy() *PermissionPolicy {
	if x != nil {
		return x.PermissionPolicy
	}
	return nil
}

var File_proto_account_aggregates_permission_use_case_permission_policy_get_proto protoreflect.FileDescriptor

var file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_rawDesc = []byte{
	0x0a, 0x48, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61,
	0x73, 0x65, 0x1a, 0x44, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65,
	0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf9, 0x01, 0x0a, 0x21, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x61, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65,
	0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x37, 0x0a, 0x25,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0xf8, 0x01, 0x0a, 0x22, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x62, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63,
	0x61, 0x73, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x8f, 0x01, 0x0a, 0x26, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x65, 0x0a, 0x11, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2e,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x10, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x42, 0x5a, 0x5a, 0x58, 0x76, 0x6c, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x75,
	0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x75, 0x73, 0x65, 0x63, 0x61, 0x73, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_rawDescOnce sync.Once
	file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_rawDescData = file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_rawDesc
)

func file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_rawDescGZIP() []byte {
	file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_rawDescOnce.Do(func() {
		file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_rawDescData)
	})
	return file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_rawDescData
}

var file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_goTypes = []interface{}{
	(*PermissionPolicyGetUseCaseRequest)(nil),      // 0: account.aggregates.permission.use_case.PermissionPolicyGetUseCaseRequest
	(*PermissionPolicyGetUseCaseRequestData)(nil),  // 1: account.aggregates.permission.use_case.PermissionPolicyGetUseCaseRequestData
	(*PermissionPolicyGetUseCaseResponse)(nil),     // 2: account.aggregates.permission.use_case.PermissionPolicyGetUseCaseResponse
	(*PermissionPolicyGetUseCaseResponseData)(nil), // 3: account.aggregates.permission.use_case.PermissionPolicyGetUseCaseResponseData
	(*messaging.Transaction)(nil),                  // 4: common.messaging.Transaction
	(*messaging.Client)(nil),                       // 5: common.messaging.Client
	(*messaging.Error)(nil),                        // 6: common.messaging.Error
	(*PermissionPolicy)(nil),                       // 7: account.aggregates.permission.use_case.PermissionPolicy
}
var file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_depIdxs = []int32{
	4, // 0: account.aggregates.permission.use_case.PermissionPolicyGetUseCaseRequest.transaction:type_name -> common.messaging.Transaction
	5, // 1: account.aggregates.permission.use_case.PermissionPolicyGetUseCaseRequest.client:type_name -> common.messaging.Client
	1, // 2: account.aggregates.permission.use_case.PermissionPolicyGetUseCaseRequest.data:type_name -> account.aggregates.permission.use_case.PermissionPolicyGetUseCaseRequestData
	4, // 3: account.aggregates.permission.use_case.PermissionPolicyGetUseCaseResponse.transaction:type_name -> common.messaging.Transaction
	6, // 4: account.aggregates.permission.use_case.PermissionPolicyGetUseCaseResponse.error:type_name -> common.messaging.Error
	3, // 5: account.aggregates.permission.use_case.PermissionPolicyGetUseCaseResponse.data:type_name -> account.aggregates.permission.use_case.PermissionPolicyGetUseCaseResponseData
	7, // 6: account.aggregates.permission.use_case.PermissionPolicyGetUseCaseResponseData.permission_policy:type_name -> account.aggregates.permission.use_case.PermissionPolicy
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_init() }
func file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_init() {
	if File_proto_account_aggregates_permission_use_case_permission_policy_get_proto != nil {
		return
	}
	file_proto_account_aggregates_permission_use_case_permission_policy_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionPolicyGetUseCaseRequest); i {
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
		file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionPolicyGetUseCaseRequestData); i {
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
		file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionPolicyGetUseCaseResponse); i {
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
		file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionPolicyGetUseCaseResponseData); i {
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
			RawDescriptor: file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_goTypes,
		DependencyIndexes: file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_depIdxs,
		MessageInfos:      file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_msgTypes,
	}.Build()
	File_proto_account_aggregates_permission_use_case_permission_policy_get_proto = out.File
	file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_rawDesc = nil
	file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_goTypes = nil
	file_proto_account_aggregates_permission_use_case_permission_policy_get_proto_depIdxs = nil
}
