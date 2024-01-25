// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/account/aggregates/template/template_role_permission_policy_create.proto

package prototemplate

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	grpc "vl-account/proto_generated/common/grpc"
	messaging "vl-account/proto_generated/common/messaging"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TemplateRolePermissionPolicyCreateRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateRoleId     uint64 `protobuf:"varint,1,opt,name=template_role_id,json=templateRoleId,proto3" json:"template_role_id,omitempty"`
	PermissionPolicyId uint64 `protobuf:"varint,2,opt,name=permission_policy_id,json=permissionPolicyId,proto3" json:"permission_policy_id,omitempty"`
}

func (x *TemplateRolePermissionPolicyCreateRequestData) Reset() {
	*x = TemplateRolePermissionPolicyCreateRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRolePermissionPolicyCreateRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRolePermissionPolicyCreateRequestData) ProtoMessage() {}

func (x *TemplateRolePermissionPolicyCreateRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRolePermissionPolicyCreateRequestData.ProtoReflect.Descriptor instead.
func (*TemplateRolePermissionPolicyCreateRequestData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_template_template_role_permission_policy_create_proto_rawDescGZIP(), []int{0}
}

func (x *TemplateRolePermissionPolicyCreateRequestData) GetTemplateRoleId() uint64 {
	if x != nil {
		return x.TemplateRoleId
	}
	return 0
}

func (x *TemplateRolePermissionPolicyCreateRequestData) GetPermissionPolicyId() uint64 {
	if x != nil {
		return x.PermissionPolicyId
	}
	return 0
}

type TemplateRolePermissionPolicyCreateResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TemplateRolePermissionPolicyCreateResponseData) Reset() {
	*x = TemplateRolePermissionPolicyCreateResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRolePermissionPolicyCreateResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRolePermissionPolicyCreateResponseData) ProtoMessage() {}

func (x *TemplateRolePermissionPolicyCreateResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRolePermissionPolicyCreateResponseData.ProtoReflect.Descriptor instead.
func (*TemplateRolePermissionPolicyCreateResponseData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_template_template_role_permission_policy_create_proto_rawDescGZIP(), []int{1}
}

func (x *TemplateRolePermissionPolicyCreateResponseData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TemplateRolePermissionPolicyCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCRequest                              `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *TemplateRolePermissionPolicyCreateRequestData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TemplateRolePermissionPolicyCreateRequest) Reset() {
	*x = TemplateRolePermissionPolicyCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRolePermissionPolicyCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRolePermissionPolicyCreateRequest) ProtoMessage() {}

func (x *TemplateRolePermissionPolicyCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRolePermissionPolicyCreateRequest.ProtoReflect.Descriptor instead.
func (*TemplateRolePermissionPolicyCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_template_template_role_permission_policy_create_proto_rawDescGZIP(), []int{2}
}

func (x *TemplateRolePermissionPolicyCreateRequest) GetBase() *grpc.GRPCRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *TemplateRolePermissionPolicyCreateRequest) GetData() *TemplateRolePermissionPolicyCreateRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type TemplateRolePermissionPolicyCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCResponse                              `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *TemplateRolePermissionPolicyCreateResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TemplateRolePermissionPolicyCreateResponse) Reset() {
	*x = TemplateRolePermissionPolicyCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRolePermissionPolicyCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRolePermissionPolicyCreateResponse) ProtoMessage() {}

func (x *TemplateRolePermissionPolicyCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRolePermissionPolicyCreateResponse.ProtoReflect.Descriptor instead.
func (*TemplateRolePermissionPolicyCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_template_template_role_permission_policy_create_proto_rawDescGZIP(), []int{3}
}

func (x *TemplateRolePermissionPolicyCreateResponse) GetBase() *grpc.GRPCResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *TemplateRolePermissionPolicyCreateResponse) GetData() *TemplateRolePermissionPolicyCreateResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type TemplateRolePermissionPolicyCreateUseCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction                         `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Client      *messaging.Client                              `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Data        *TemplateRolePermissionPolicyCreateRequestData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TemplateRolePermissionPolicyCreateUseCaseRequest) Reset() {
	*x = TemplateRolePermissionPolicyCreateUseCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRolePermissionPolicyCreateUseCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRolePermissionPolicyCreateUseCaseRequest) ProtoMessage() {}

func (x *TemplateRolePermissionPolicyCreateUseCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRolePermissionPolicyCreateUseCaseRequest.ProtoReflect.Descriptor instead.
func (*TemplateRolePermissionPolicyCreateUseCaseRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_template_template_role_permission_policy_create_proto_rawDescGZIP(), []int{4}
}

func (x *TemplateRolePermissionPolicyCreateUseCaseRequest) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TemplateRolePermissionPolicyCreateUseCaseRequest) GetClient() *messaging.Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *TemplateRolePermissionPolicyCreateUseCaseRequest) GetData() *TemplateRolePermissionPolicyCreateRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type TemplateRolePermissionPolicyCreateUseCaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction                          `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Error       *messaging.Error                                `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data        *TemplateRolePermissionPolicyCreateResponseData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TemplateRolePermissionPolicyCreateUseCaseResponse) Reset() {
	*x = TemplateRolePermissionPolicyCreateUseCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRolePermissionPolicyCreateUseCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRolePermissionPolicyCreateUseCaseResponse) ProtoMessage() {}

func (x *TemplateRolePermissionPolicyCreateUseCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRolePermissionPolicyCreateUseCaseResponse.ProtoReflect.Descriptor instead.
func (*TemplateRolePermissionPolicyCreateUseCaseResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_template_template_role_permission_policy_create_proto_rawDescGZIP(), []int{5}
}

func (x *TemplateRolePermissionPolicyCreateUseCaseResponse) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TemplateRolePermissionPolicyCreateUseCaseResponse) GetError() *messaging.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *TemplateRolePermissionPolicyCreateUseCaseResponse) GetData() *TemplateRolePermissionPolicyCreateResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type TemplateRolePermissionPolicyCreateCQRSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction                         `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Client      *messaging.Client                              `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Data        *TemplateRolePermissionPolicyCreateRequestData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TemplateRolePermissionPolicyCreateCQRSRequest) Reset() {
	*x = TemplateRolePermissionPolicyCreateCQRSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRolePermissionPolicyCreateCQRSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRolePermissionPolicyCreateCQRSRequest) ProtoMessage() {}

func (x *TemplateRolePermissionPolicyCreateCQRSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRolePermissionPolicyCreateCQRSRequest.ProtoReflect.Descriptor instead.
func (*TemplateRolePermissionPolicyCreateCQRSRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_template_template_role_permission_policy_create_proto_rawDescGZIP(), []int{6}
}

func (x *TemplateRolePermissionPolicyCreateCQRSRequest) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TemplateRolePermissionPolicyCreateCQRSRequest) GetClient() *messaging.Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *TemplateRolePermissionPolicyCreateCQRSRequest) GetData() *TemplateRolePermissionPolicyCreateRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type TemplateRolePermissionPolicyCreateCQRSResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction        `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Error       *messaging.Error              `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data        *TemplateRolePermissionPolicy `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TemplateRolePermissionPolicyCreateCQRSResponse) Reset() {
	*x = TemplateRolePermissionPolicyCreateCQRSResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRolePermissionPolicyCreateCQRSResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRolePermissionPolicyCreateCQRSResponse) ProtoMessage() {}

func (x *TemplateRolePermissionPolicyCreateCQRSResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRolePermissionPolicyCreateCQRSResponse.ProtoReflect.Descriptor instead.
func (*TemplateRolePermissionPolicyCreateCQRSResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_template_template_role_permission_policy_create_proto_rawDescGZIP(), []int{7}
}

func (x *TemplateRolePermissionPolicyCreateCQRSResponse) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TemplateRolePermissionPolicyCreateCQRSResponse) GetError() *messaging.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *TemplateRolePermissionPolicyCreateCQRSResponse) GetData() *TemplateRolePermissionPolicy {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_account_aggregates_template_template_role_permission_policy_create_proto protoreflect.FileDescriptor

var file_proto_account_aggregates_template_template_role_permission_policy_create_proto_rawDesc = []byte{
	0x0a, 0x4e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6c,
	0x65, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x47, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a,
	0x2d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x28,
	0x0a, 0x10, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x2e, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb9, 0x01, 0x0a,
	0x29, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xbc, 0x01, 0x0a, 0x2a, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x85, 0x02, 0x0a, 0x30, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a,
	0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x5e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4a, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x84, 0x02, 0x0a, 0x31, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x5f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x82, 0x02, 0x0a, 0x2d, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x51, 0x52,
	0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x06, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x5e, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xef, 0x01, 0x0a, 0x2e,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x51, 0x52, 0x53, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x4d,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x47, 0x5a,
	0x45, 0x76, 0x6c, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_aggregates_template_template_role_permission_policy_create_proto_rawDescOnce sync.Once
	file_proto_account_aggregates_template_template_role_permission_policy_create_proto_rawDescData = file_proto_account_aggregates_template_template_role_permission_policy_create_proto_rawDesc
)

func file_proto_account_aggregates_template_template_role_permission_policy_create_proto_rawDescGZIP() []byte {
	file_proto_account_aggregates_template_template_role_permission_policy_create_proto_rawDescOnce.Do(func() {
		file_proto_account_aggregates_template_template_role_permission_policy_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_aggregates_template_template_role_permission_policy_create_proto_rawDescData)
	})
	return file_proto_account_aggregates_template_template_role_permission_policy_create_proto_rawDescData
}

var file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_account_aggregates_template_template_role_permission_policy_create_proto_goTypes = []interface{}{
	(*TemplateRolePermissionPolicyCreateRequestData)(nil),     // 0: account.aggregates.template.TemplateRolePermissionPolicyCreateRequestData
	(*TemplateRolePermissionPolicyCreateResponseData)(nil),    // 1: account.aggregates.template.TemplateRolePermissionPolicyCreateResponseData
	(*TemplateRolePermissionPolicyCreateRequest)(nil),         // 2: account.aggregates.template.TemplateRolePermissionPolicyCreateRequest
	(*TemplateRolePermissionPolicyCreateResponse)(nil),        // 3: account.aggregates.template.TemplateRolePermissionPolicyCreateResponse
	(*TemplateRolePermissionPolicyCreateUseCaseRequest)(nil),  // 4: account.aggregates.template.TemplateRolePermissionPolicyCreateUseCaseRequest
	(*TemplateRolePermissionPolicyCreateUseCaseResponse)(nil), // 5: account.aggregates.template.TemplateRolePermissionPolicyCreateUseCaseResponse
	(*TemplateRolePermissionPolicyCreateCQRSRequest)(nil),     // 6: account.aggregates.template.TemplateRolePermissionPolicyCreateCQRSRequest
	(*TemplateRolePermissionPolicyCreateCQRSResponse)(nil),    // 7: account.aggregates.template.TemplateRolePermissionPolicyCreateCQRSResponse
	(*grpc.GRPCRequest)(nil),                                  // 8: common.grpc.GRPCRequest
	(*grpc.GRPCResponse)(nil),                                 // 9: common.grpc.GRPCResponse
	(*messaging.Transaction)(nil),                             // 10: common.messaging.Transaction
	(*messaging.Client)(nil),                                  // 11: common.messaging.Client
	(*messaging.Error)(nil),                                   // 12: common.messaging.Error
	(*TemplateRolePermissionPolicy)(nil),                      // 13: account.aggregates.template.TemplateRolePermissionPolicy
}
var file_proto_account_aggregates_template_template_role_permission_policy_create_proto_depIdxs = []int32{
	8,  // 0: account.aggregates.template.TemplateRolePermissionPolicyCreateRequest.base:type_name -> common.grpc.GRPCRequest
	0,  // 1: account.aggregates.template.TemplateRolePermissionPolicyCreateRequest.data:type_name -> account.aggregates.template.TemplateRolePermissionPolicyCreateRequestData
	9,  // 2: account.aggregates.template.TemplateRolePermissionPolicyCreateResponse.base:type_name -> common.grpc.GRPCResponse
	1,  // 3: account.aggregates.template.TemplateRolePermissionPolicyCreateResponse.data:type_name -> account.aggregates.template.TemplateRolePermissionPolicyCreateResponseData
	10, // 4: account.aggregates.template.TemplateRolePermissionPolicyCreateUseCaseRequest.transaction:type_name -> common.messaging.Transaction
	11, // 5: account.aggregates.template.TemplateRolePermissionPolicyCreateUseCaseRequest.client:type_name -> common.messaging.Client
	0,  // 6: account.aggregates.template.TemplateRolePermissionPolicyCreateUseCaseRequest.data:type_name -> account.aggregates.template.TemplateRolePermissionPolicyCreateRequestData
	10, // 7: account.aggregates.template.TemplateRolePermissionPolicyCreateUseCaseResponse.transaction:type_name -> common.messaging.Transaction
	12, // 8: account.aggregates.template.TemplateRolePermissionPolicyCreateUseCaseResponse.error:type_name -> common.messaging.Error
	1,  // 9: account.aggregates.template.TemplateRolePermissionPolicyCreateUseCaseResponse.data:type_name -> account.aggregates.template.TemplateRolePermissionPolicyCreateResponseData
	10, // 10: account.aggregates.template.TemplateRolePermissionPolicyCreateCQRSRequest.transaction:type_name -> common.messaging.Transaction
	11, // 11: account.aggregates.template.TemplateRolePermissionPolicyCreateCQRSRequest.client:type_name -> common.messaging.Client
	0,  // 12: account.aggregates.template.TemplateRolePermissionPolicyCreateCQRSRequest.data:type_name -> account.aggregates.template.TemplateRolePermissionPolicyCreateRequestData
	10, // 13: account.aggregates.template.TemplateRolePermissionPolicyCreateCQRSResponse.transaction:type_name -> common.messaging.Transaction
	12, // 14: account.aggregates.template.TemplateRolePermissionPolicyCreateCQRSResponse.error:type_name -> common.messaging.Error
	13, // 15: account.aggregates.template.TemplateRolePermissionPolicyCreateCQRSResponse.data:type_name -> account.aggregates.template.TemplateRolePermissionPolicy
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() {
	file_proto_account_aggregates_template_template_role_permission_policy_create_proto_init()
}
func file_proto_account_aggregates_template_template_role_permission_policy_create_proto_init() {
	if File_proto_account_aggregates_template_template_role_permission_policy_create_proto != nil {
		return
	}
	file_proto_account_aggregates_template_template_role_permission_policy_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRolePermissionPolicyCreateRequestData); i {
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
		file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRolePermissionPolicyCreateResponseData); i {
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
		file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRolePermissionPolicyCreateRequest); i {
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
		file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRolePermissionPolicyCreateResponse); i {
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
		file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRolePermissionPolicyCreateUseCaseRequest); i {
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
		file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRolePermissionPolicyCreateUseCaseResponse); i {
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
		file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRolePermissionPolicyCreateCQRSRequest); i {
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
		file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRolePermissionPolicyCreateCQRSResponse); i {
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
			RawDescriptor: file_proto_account_aggregates_template_template_role_permission_policy_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_aggregates_template_template_role_permission_policy_create_proto_goTypes,
		DependencyIndexes: file_proto_account_aggregates_template_template_role_permission_policy_create_proto_depIdxs,
		MessageInfos:      file_proto_account_aggregates_template_template_role_permission_policy_create_proto_msgTypes,
	}.Build()
	File_proto_account_aggregates_template_template_role_permission_policy_create_proto = out.File
	file_proto_account_aggregates_template_template_role_permission_policy_create_proto_rawDesc = nil
	file_proto_account_aggregates_template_template_role_permission_policy_create_proto_goTypes = nil
	file_proto_account_aggregates_template_template_role_permission_policy_create_proto_depIdxs = nil
}