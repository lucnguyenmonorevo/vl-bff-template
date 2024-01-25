// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/account/aggregates/permission/use_case/template_organizational_role_create.proto

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

type TemplateOrganizationalRoleCreateUseCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction                              `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Client      *messaging.Client                                   `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Data        *TemplateOrganizationalRoleCreateUseCaseRequestData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TemplateOrganizationalRoleCreateUseCaseRequest) Reset() {
	*x = TemplateOrganizationalRoleCreateUseCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateOrganizationalRoleCreateUseCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateOrganizationalRoleCreateUseCaseRequest) ProtoMessage() {}

func (x *TemplateOrganizationalRoleCreateUseCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateOrganizationalRoleCreateUseCaseRequest.ProtoReflect.Descriptor instead.
func (*TemplateOrganizationalRoleCreateUseCaseRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_rawDescGZIP(), []int{0}
}

func (x *TemplateOrganizationalRoleCreateUseCaseRequest) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TemplateOrganizationalRoleCreateUseCaseRequest) GetClient() *messaging.Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *TemplateOrganizationalRoleCreateUseCaseRequest) GetData() *TemplateOrganizationalRoleCreateUseCaseRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type TemplateOrganizationalRoleCreateUseCaseRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TemplateOrganizationalRoleCreateUseCaseRequestData) Reset() {
	*x = TemplateOrganizationalRoleCreateUseCaseRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateOrganizationalRoleCreateUseCaseRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateOrganizationalRoleCreateUseCaseRequestData) ProtoMessage() {}

func (x *TemplateOrganizationalRoleCreateUseCaseRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateOrganizationalRoleCreateUseCaseRequestData.ProtoReflect.Descriptor instead.
func (*TemplateOrganizationalRoleCreateUseCaseRequestData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_rawDescGZIP(), []int{1}
}

func (x *TemplateOrganizationalRoleCreateUseCaseRequestData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TemplateOrganizationalRoleCreateUseCaseRequestData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TemplateOrganizationalRoleCreateUseCaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction      `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Error       *messaging.Error            `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data        *TemplateOrganizationalRole `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TemplateOrganizationalRoleCreateUseCaseResponse) Reset() {
	*x = TemplateOrganizationalRoleCreateUseCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateOrganizationalRoleCreateUseCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateOrganizationalRoleCreateUseCaseResponse) ProtoMessage() {}

func (x *TemplateOrganizationalRoleCreateUseCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateOrganizationalRoleCreateUseCaseResponse.ProtoReflect.Descriptor instead.
func (*TemplateOrganizationalRoleCreateUseCaseResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_rawDescGZIP(), []int{2}
}

func (x *TemplateOrganizationalRoleCreateUseCaseResponse) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TemplateOrganizationalRoleCreateUseCaseResponse) GetError() *messaging.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *TemplateOrganizationalRoleCreateUseCaseResponse) GetData() *TemplateOrganizationalRole {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto protoreflect.FileDescriptor

var file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_rawDesc = []byte{
	0x0a, 0x56, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65,
	0x1a, 0x4f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x02, 0x0a, 0x2e, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a,
	0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x6e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5a, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x75, 0x73,
	0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x58, 0x0a, 0x32, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xf9, 0x01, 0x0a, 0x2f, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x56, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73,
	0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x5f,
	0x63, 0x61, 0x73, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x5a, 0x5a, 0x58, 0x76, 0x6c, 0x2d, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x75, 0x73, 0x65, 0x63, 0x61, 0x73,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_rawDescOnce sync.Once
	file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_rawDescData = file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_rawDesc
)

func file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_rawDescGZIP() []byte {
	file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_rawDescOnce.Do(func() {
		file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_rawDescData)
	})
	return file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_rawDescData
}

var file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_goTypes = []interface{}{
	(*TemplateOrganizationalRoleCreateUseCaseRequest)(nil),     // 0: account.aggregates.permission.use_case.TemplateOrganizationalRoleCreateUseCaseRequest
	(*TemplateOrganizationalRoleCreateUseCaseRequestData)(nil), // 1: account.aggregates.permission.use_case.TemplateOrganizationalRoleCreateUseCaseRequestData
	(*TemplateOrganizationalRoleCreateUseCaseResponse)(nil),    // 2: account.aggregates.permission.use_case.TemplateOrganizationalRoleCreateUseCaseResponse
	(*messaging.Transaction)(nil),                              // 3: common.messaging.Transaction
	(*messaging.Client)(nil),                                   // 4: common.messaging.Client
	(*messaging.Error)(nil),                                    // 5: common.messaging.Error
	(*TemplateOrganizationalRole)(nil),                         // 6: account.aggregates.permission.use_case.TemplateOrganizationalRole
}
var file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_depIdxs = []int32{
	3, // 0: account.aggregates.permission.use_case.TemplateOrganizationalRoleCreateUseCaseRequest.transaction:type_name -> common.messaging.Transaction
	4, // 1: account.aggregates.permission.use_case.TemplateOrganizationalRoleCreateUseCaseRequest.client:type_name -> common.messaging.Client
	1, // 2: account.aggregates.permission.use_case.TemplateOrganizationalRoleCreateUseCaseRequest.data:type_name -> account.aggregates.permission.use_case.TemplateOrganizationalRoleCreateUseCaseRequestData
	3, // 3: account.aggregates.permission.use_case.TemplateOrganizationalRoleCreateUseCaseResponse.transaction:type_name -> common.messaging.Transaction
	5, // 4: account.aggregates.permission.use_case.TemplateOrganizationalRoleCreateUseCaseResponse.error:type_name -> common.messaging.Error
	6, // 5: account.aggregates.permission.use_case.TemplateOrganizationalRoleCreateUseCaseResponse.data:type_name -> account.aggregates.permission.use_case.TemplateOrganizationalRole
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_init()
}
func file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_init() {
	if File_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto != nil {
		return
	}
	file_proto_account_aggregates_permission_use_case_template_organizational_role_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateOrganizationalRoleCreateUseCaseRequest); i {
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
		file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateOrganizationalRoleCreateUseCaseRequestData); i {
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
		file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateOrganizationalRoleCreateUseCaseResponse); i {
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
			RawDescriptor: file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_goTypes,
		DependencyIndexes: file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_depIdxs,
		MessageInfos:      file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_msgTypes,
	}.Build()
	File_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto = out.File
	file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_rawDesc = nil
	file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_goTypes = nil
	file_proto_account_aggregates_permission_use_case_template_organizational_role_create_proto_depIdxs = nil
}
