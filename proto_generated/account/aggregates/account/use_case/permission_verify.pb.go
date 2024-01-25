// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/account/aggregates/account/use_case/permission_verify.proto

package protoaccountusecase

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

type PermissionVerifyUseCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction              `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Client      *messaging.Client                   `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Data        *PermissionVerifyUseCaseRequestData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PermissionVerifyUseCaseRequest) Reset() {
	*x = PermissionVerifyUseCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_use_case_permission_verify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionVerifyUseCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionVerifyUseCaseRequest) ProtoMessage() {}

func (x *PermissionVerifyUseCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_use_case_permission_verify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionVerifyUseCaseRequest.ProtoReflect.Descriptor instead.
func (*PermissionVerifyUseCaseRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_use_case_permission_verify_proto_rawDescGZIP(), []int{0}
}

func (x *PermissionVerifyUseCaseRequest) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *PermissionVerifyUseCaseRequest) GetClient() *messaging.Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *PermissionVerifyUseCaseRequest) GetData() *PermissionVerifyUseCaseRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type PermissionVerifyUseCaseRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FunctionId             int64 `protobuf:"varint,1,opt,name=function_id,json=functionId,proto3" json:"function_id,omitempty"`
	PageId                 int64 `protobuf:"varint,2,opt,name=page_id,json=pageId,proto3" json:"page_id,omitempty"`
	OrganizationId         int64 `protobuf:"varint,3,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	UserId                 int64 `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IsAdmin                bool  `protobuf:"varint,5,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	OrchestratorFunctionId int64 `protobuf:"varint,6,opt,name=orchestrator_function_id,json=orchestratorFunctionId,proto3" json:"orchestrator_function_id,omitempty"`
	Create                 bool  `protobuf:"varint,7,opt,name=create,proto3" json:"create,omitempty"`
	Read                   bool  `protobuf:"varint,8,opt,name=read,proto3" json:"read,omitempty"`
	Update                 bool  `protobuf:"varint,9,opt,name=update,proto3" json:"update,omitempty"`
	Delete                 bool  `protobuf:"varint,10,opt,name=delete,proto3" json:"delete,omitempty"`
}

func (x *PermissionVerifyUseCaseRequestData) Reset() {
	*x = PermissionVerifyUseCaseRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_use_case_permission_verify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionVerifyUseCaseRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionVerifyUseCaseRequestData) ProtoMessage() {}

func (x *PermissionVerifyUseCaseRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_use_case_permission_verify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionVerifyUseCaseRequestData.ProtoReflect.Descriptor instead.
func (*PermissionVerifyUseCaseRequestData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_use_case_permission_verify_proto_rawDescGZIP(), []int{1}
}

func (x *PermissionVerifyUseCaseRequestData) GetFunctionId() int64 {
	if x != nil {
		return x.FunctionId
	}
	return 0
}

func (x *PermissionVerifyUseCaseRequestData) GetPageId() int64 {
	if x != nil {
		return x.PageId
	}
	return 0
}

func (x *PermissionVerifyUseCaseRequestData) GetOrganizationId() int64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *PermissionVerifyUseCaseRequestData) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PermissionVerifyUseCaseRequestData) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *PermissionVerifyUseCaseRequestData) GetOrchestratorFunctionId() int64 {
	if x != nil {
		return x.OrchestratorFunctionId
	}
	return 0
}

func (x *PermissionVerifyUseCaseRequestData) GetCreate() bool {
	if x != nil {
		return x.Create
	}
	return false
}

func (x *PermissionVerifyUseCaseRequestData) GetRead() bool {
	if x != nil {
		return x.Read
	}
	return false
}

func (x *PermissionVerifyUseCaseRequestData) GetUpdate() bool {
	if x != nil {
		return x.Update
	}
	return false
}

func (x *PermissionVerifyUseCaseRequestData) GetDelete() bool {
	if x != nil {
		return x.Delete
	}
	return false
}

type PermissionVerifyUseCaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Error       *messaging.Error       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data        bool                   `protobuf:"varint,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PermissionVerifyUseCaseResponse) Reset() {
	*x = PermissionVerifyUseCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_use_case_permission_verify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionVerifyUseCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionVerifyUseCaseResponse) ProtoMessage() {}

func (x *PermissionVerifyUseCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_use_case_permission_verify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionVerifyUseCaseResponse.ProtoReflect.Descriptor instead.
func (*PermissionVerifyUseCaseResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_use_case_permission_verify_proto_rawDescGZIP(), []int{2}
}

func (x *PermissionVerifyUseCaseResponse) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *PermissionVerifyUseCaseResponse) GetError() *messaging.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *PermissionVerifyUseCaseResponse) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

var File_proto_account_aggregates_account_use_case_permission_verify_proto protoreflect.FileDescriptor

var file_proto_account_aggregates_account_use_case_permission_verify_proto_rawDesc = []byte{
	0x0a, 0x41, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x23, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf0, 0x01, 0x0a, 0x1e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x5b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x43, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0xd1, 0x02, 0x0a, 0x22, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x12, 0x38, 0x0a, 0x18, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x16, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x72, 0x65, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x22, 0xa5, 0x01, 0x0a, 0x1f, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x43,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42,
	0x54, 0x5a, 0x52, 0x76, 0x6c, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73,
	0x65, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x75, 0x73,
	0x65, 0x63, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_aggregates_account_use_case_permission_verify_proto_rawDescOnce sync.Once
	file_proto_account_aggregates_account_use_case_permission_verify_proto_rawDescData = file_proto_account_aggregates_account_use_case_permission_verify_proto_rawDesc
)

func file_proto_account_aggregates_account_use_case_permission_verify_proto_rawDescGZIP() []byte {
	file_proto_account_aggregates_account_use_case_permission_verify_proto_rawDescOnce.Do(func() {
		file_proto_account_aggregates_account_use_case_permission_verify_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_aggregates_account_use_case_permission_verify_proto_rawDescData)
	})
	return file_proto_account_aggregates_account_use_case_permission_verify_proto_rawDescData
}

var file_proto_account_aggregates_account_use_case_permission_verify_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_account_aggregates_account_use_case_permission_verify_proto_goTypes = []interface{}{
	(*PermissionVerifyUseCaseRequest)(nil),     // 0: account.aggregates.account.use_case.PermissionVerifyUseCaseRequest
	(*PermissionVerifyUseCaseRequestData)(nil), // 1: account.aggregates.account.use_case.PermissionVerifyUseCaseRequestData
	(*PermissionVerifyUseCaseResponse)(nil),    // 2: account.aggregates.account.use_case.PermissionVerifyUseCaseResponse
	(*messaging.Transaction)(nil),              // 3: common.messaging.Transaction
	(*messaging.Client)(nil),                   // 4: common.messaging.Client
	(*messaging.Error)(nil),                    // 5: common.messaging.Error
}
var file_proto_account_aggregates_account_use_case_permission_verify_proto_depIdxs = []int32{
	3, // 0: account.aggregates.account.use_case.PermissionVerifyUseCaseRequest.transaction:type_name -> common.messaging.Transaction
	4, // 1: account.aggregates.account.use_case.PermissionVerifyUseCaseRequest.client:type_name -> common.messaging.Client
	1, // 2: account.aggregates.account.use_case.PermissionVerifyUseCaseRequest.data:type_name -> account.aggregates.account.use_case.PermissionVerifyUseCaseRequestData
	3, // 3: account.aggregates.account.use_case.PermissionVerifyUseCaseResponse.transaction:type_name -> common.messaging.Transaction
	5, // 4: account.aggregates.account.use_case.PermissionVerifyUseCaseResponse.error:type_name -> common.messaging.Error
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_account_aggregates_account_use_case_permission_verify_proto_init() }
func file_proto_account_aggregates_account_use_case_permission_verify_proto_init() {
	if File_proto_account_aggregates_account_use_case_permission_verify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_account_aggregates_account_use_case_permission_verify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionVerifyUseCaseRequest); i {
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
		file_proto_account_aggregates_account_use_case_permission_verify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionVerifyUseCaseRequestData); i {
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
		file_proto_account_aggregates_account_use_case_permission_verify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionVerifyUseCaseResponse); i {
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
			RawDescriptor: file_proto_account_aggregates_account_use_case_permission_verify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_aggregates_account_use_case_permission_verify_proto_goTypes,
		DependencyIndexes: file_proto_account_aggregates_account_use_case_permission_verify_proto_depIdxs,
		MessageInfos:      file_proto_account_aggregates_account_use_case_permission_verify_proto_msgTypes,
	}.Build()
	File_proto_account_aggregates_account_use_case_permission_verify_proto = out.File
	file_proto_account_aggregates_account_use_case_permission_verify_proto_rawDesc = nil
	file_proto_account_aggregates_account_use_case_permission_verify_proto_goTypes = nil
	file_proto_account_aggregates_account_use_case_permission_verify_proto_depIdxs = nil
}
