// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/account/aggregates/permission/grpc/template_organizational_role_get.proto

package protopermissiongrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	grpc "vl-account/proto_generated/common/grpc"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TemplateOrganizationalRoleGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCRequest                         `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *TemplateOrganizationalRoleGetRequestData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TemplateOrganizationalRoleGetRequest) Reset() {
	*x = TemplateOrganizationalRoleGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateOrganizationalRoleGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateOrganizationalRoleGetRequest) ProtoMessage() {}

func (x *TemplateOrganizationalRoleGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateOrganizationalRoleGetRequest.ProtoReflect.Descriptor instead.
func (*TemplateOrganizationalRoleGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_rawDescGZIP(), []int{0}
}

func (x *TemplateOrganizationalRoleGetRequest) GetBase() *grpc.GRPCRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *TemplateOrganizationalRoleGetRequest) GetData() *TemplateOrganizationalRoleGetRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type TemplateOrganizationalRoleGetRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TemplateOrganizationalRoleGetRequestData) Reset() {
	*x = TemplateOrganizationalRoleGetRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateOrganizationalRoleGetRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateOrganizationalRoleGetRequestData) ProtoMessage() {}

func (x *TemplateOrganizationalRoleGetRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateOrganizationalRoleGetRequestData.ProtoReflect.Descriptor instead.
func (*TemplateOrganizationalRoleGetRequestData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_rawDescGZIP(), []int{1}
}

func (x *TemplateOrganizationalRoleGetRequestData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TemplateOrganizationalRoleGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCResponse                         `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *TemplateOrganizationalRoleGetResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TemplateOrganizationalRoleGetResponse) Reset() {
	*x = TemplateOrganizationalRoleGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateOrganizationalRoleGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateOrganizationalRoleGetResponse) ProtoMessage() {}

func (x *TemplateOrganizationalRoleGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateOrganizationalRoleGetResponse.ProtoReflect.Descriptor instead.
func (*TemplateOrganizationalRoleGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_rawDescGZIP(), []int{2}
}

func (x *TemplateOrganizationalRoleGetResponse) GetBase() *grpc.GRPCResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *TemplateOrganizationalRoleGetResponse) GetData() *TemplateOrganizationalRoleGetResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type TemplateOrganizationalRoleGetResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateOrganizationalRole *TemplateOrganizationalRole `protobuf:"bytes,1,opt,name=template_organizational_role,json=templateOrganizationalRole,proto3" json:"template_organizational_role,omitempty"`
}

func (x *TemplateOrganizationalRoleGetResponseData) Reset() {
	*x = TemplateOrganizationalRoleGetResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateOrganizationalRoleGetResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateOrganizationalRoleGetResponseData) ProtoMessage() {}

func (x *TemplateOrganizationalRoleGetResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateOrganizationalRoleGetResponseData.ProtoReflect.Descriptor instead.
func (*TemplateOrganizationalRoleGetResponseData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_rawDescGZIP(), []int{3}
}

func (x *TemplateOrganizationalRoleGetResponseData) GetTemplateOrganizationalRole() *TemplateOrganizationalRole {
	if x != nil {
		return x.TemplateOrganizationalRole
	}
	return nil
}

type TemplateOrganizationalRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TemplateOrganizationalRole) Reset() {
	*x = TemplateOrganizationalRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateOrganizationalRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateOrganizationalRole) ProtoMessage() {}

func (x *TemplateOrganizationalRole) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateOrganizationalRole.ProtoReflect.Descriptor instead.
func (*TemplateOrganizationalRole) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_rawDescGZIP(), []int{4}
}

func (x *TemplateOrganizationalRole) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TemplateOrganizationalRole) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto protoreflect.FileDescriptor

var file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x22, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x24, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f,
	0x6c, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04,
	0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x28,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb9, 0x01, 0x0a, 0x25, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x61, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x4d, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xae, 0x01, 0x0a, 0x29, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52,
	0x6f, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x80, 0x01, 0x0a, 0x1c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x1a, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x40, 0x0a, 0x1a, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x53, 0x5a, 0x51, 0x76, 0x6c, 0x2d, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_rawDescOnce sync.Once
	file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_rawDescData = file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_rawDesc
)

func file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_rawDescGZIP() []byte {
	file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_rawDescOnce.Do(func() {
		file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_rawDescData)
	})
	return file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_rawDescData
}

var file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_goTypes = []interface{}{
	(*TemplateOrganizationalRoleGetRequest)(nil),      // 0: account.aggregates.permission.grpc.TemplateOrganizationalRoleGetRequest
	(*TemplateOrganizationalRoleGetRequestData)(nil),  // 1: account.aggregates.permission.grpc.TemplateOrganizationalRoleGetRequestData
	(*TemplateOrganizationalRoleGetResponse)(nil),     // 2: account.aggregates.permission.grpc.TemplateOrganizationalRoleGetResponse
	(*TemplateOrganizationalRoleGetResponseData)(nil), // 3: account.aggregates.permission.grpc.TemplateOrganizationalRoleGetResponseData
	(*TemplateOrganizationalRole)(nil),                // 4: account.aggregates.permission.grpc.TemplateOrganizationalRole
	(*grpc.GRPCRequest)(nil),                          // 5: common.grpc.GRPCRequest
	(*grpc.GRPCResponse)(nil),                         // 6: common.grpc.GRPCResponse
}
var file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_depIdxs = []int32{
	5, // 0: account.aggregates.permission.grpc.TemplateOrganizationalRoleGetRequest.base:type_name -> common.grpc.GRPCRequest
	1, // 1: account.aggregates.permission.grpc.TemplateOrganizationalRoleGetRequest.data:type_name -> account.aggregates.permission.grpc.TemplateOrganizationalRoleGetRequestData
	6, // 2: account.aggregates.permission.grpc.TemplateOrganizationalRoleGetResponse.base:type_name -> common.grpc.GRPCResponse
	3, // 3: account.aggregates.permission.grpc.TemplateOrganizationalRoleGetResponse.data:type_name -> account.aggregates.permission.grpc.TemplateOrganizationalRoleGetResponseData
	4, // 4: account.aggregates.permission.grpc.TemplateOrganizationalRoleGetResponseData.template_organizational_role:type_name -> account.aggregates.permission.grpc.TemplateOrganizationalRole
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() {
	file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_init()
}
func file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_init() {
	if File_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateOrganizationalRoleGetRequest); i {
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
		file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateOrganizationalRoleGetRequestData); i {
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
		file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateOrganizationalRoleGetResponse); i {
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
		file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateOrganizationalRoleGetResponseData); i {
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
		file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateOrganizationalRole); i {
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
			RawDescriptor: file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_goTypes,
		DependencyIndexes: file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_depIdxs,
		MessageInfos:      file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_msgTypes,
	}.Build()
	File_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto = out.File
	file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_rawDesc = nil
	file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_goTypes = nil
	file_proto_account_aggregates_permission_grpc_template_organizational_role_get_proto_depIdxs = nil
}
