// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/account/aggregates/permission/grpc/organizational_role_permission_policy_get.proto

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

type OrganizationalRolePermissionPolicyGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCRequest                                 `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *OrganizationalRolePermissionPolicyGetRequestData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OrganizationalRolePermissionPolicyGetRequest) Reset() {
	*x = OrganizationalRolePermissionPolicyGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationalRolePermissionPolicyGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationalRolePermissionPolicyGetRequest) ProtoMessage() {}

func (x *OrganizationalRolePermissionPolicyGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationalRolePermissionPolicyGetRequest.ProtoReflect.Descriptor instead.
func (*OrganizationalRolePermissionPolicyGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_rawDescGZIP(), []int{0}
}

func (x *OrganizationalRolePermissionPolicyGetRequest) GetBase() *grpc.GRPCRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *OrganizationalRolePermissionPolicyGetRequest) GetData() *OrganizationalRolePermissionPolicyGetRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type OrganizationalRolePermissionPolicyGetRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrganizationalRolePermissionPolicyGetRequestData) Reset() {
	*x = OrganizationalRolePermissionPolicyGetRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationalRolePermissionPolicyGetRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationalRolePermissionPolicyGetRequestData) ProtoMessage() {}

func (x *OrganizationalRolePermissionPolicyGetRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationalRolePermissionPolicyGetRequestData.ProtoReflect.Descriptor instead.
func (*OrganizationalRolePermissionPolicyGetRequestData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_rawDescGZIP(), []int{1}
}

func (x *OrganizationalRolePermissionPolicyGetRequestData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OrganizationalRolePermissionPolicyGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCResponse                                 `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *OrganizationalRolePermissionPolicyGetResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OrganizationalRolePermissionPolicyGetResponse) Reset() {
	*x = OrganizationalRolePermissionPolicyGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationalRolePermissionPolicyGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationalRolePermissionPolicyGetResponse) ProtoMessage() {}

func (x *OrganizationalRolePermissionPolicyGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationalRolePermissionPolicyGetResponse.ProtoReflect.Descriptor instead.
func (*OrganizationalRolePermissionPolicyGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_rawDescGZIP(), []int{2}
}

func (x *OrganizationalRolePermissionPolicyGetResponse) GetBase() *grpc.GRPCResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *OrganizationalRolePermissionPolicyGetResponse) GetData() *OrganizationalRolePermissionPolicyGetResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type OrganizationalRolePermissionPolicyGetResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationalRolePermissionPolicy *OrganizationalRolePermissionPolicy `protobuf:"bytes,1,opt,name=organizational_role_permission_policy,json=organizationalRolePermissionPolicy,proto3" json:"organizational_role_permission_policy,omitempty"`
}

func (x *OrganizationalRolePermissionPolicyGetResponseData) Reset() {
	*x = OrganizationalRolePermissionPolicyGetResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationalRolePermissionPolicyGetResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationalRolePermissionPolicyGetResponseData) ProtoMessage() {}

func (x *OrganizationalRolePermissionPolicyGetResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationalRolePermissionPolicyGetResponseData.ProtoReflect.Descriptor instead.
func (*OrganizationalRolePermissionPolicyGetResponseData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_rawDescGZIP(), []int{3}
}

func (x *OrganizationalRolePermissionPolicyGetResponseData) GetOrganizationalRolePermissionPolicy() *OrganizationalRolePermissionPolicy {
	if x != nil {
		return x.OrganizationalRolePermissionPolicy
	}
	return nil
}

type OrganizationalRolePermissionPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PermissionPolicyId   uint64 `protobuf:"varint,2,opt,name=permission_policy_id,json=permissionPolicyId,proto3" json:"permission_policy_id,omitempty"`
	OrganizationalRoleId uint64 `protobuf:"varint,3,opt,name=organizational_role_id,json=organizationalRoleId,proto3" json:"organizational_role_id,omitempty"`
}

func (x *OrganizationalRolePermissionPolicy) Reset() {
	*x = OrganizationalRolePermissionPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationalRolePermissionPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationalRolePermissionPolicy) ProtoMessage() {}

func (x *OrganizationalRolePermissionPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationalRolePermissionPolicy.ProtoReflect.Descriptor instead.
func (*OrganizationalRolePermissionPolicy) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_rawDescGZIP(), []int{4}
}

func (x *OrganizationalRolePermissionPolicy) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrganizationalRolePermissionPolicy) GetPermissionPolicyId() uint64 {
	if x != nil {
		return x.PermissionPolicyId
	}
	return 0
}

func (x *OrganizationalRolePermissionPolicy) GetOrganizationalRoleId() uint64 {
	if x != nil {
		return x.OrganizationalRoleId
	}
	return 0
}

var File_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto protoreflect.FileDescriptor

var file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_rawDesc = []byte{
	0x0a, 0x58, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1c,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a,
	0x2c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52,
	0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x68, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x54, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x42, 0x0a, 0x30, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0xc9, 0x01, 0x0a, 0x2d, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x55, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xcf, 0x01, 0x0a, 0x31, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x99, 0x01, 0x0a, 0x25,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73,
	0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52,
	0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x22, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x9c, 0x01, 0x0a, 0x22, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30,
	0x0a, 0x14, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64,
	0x12, 0x34, 0x0a, 0x16, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x14, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x42, 0x53, 0x5a, 0x51, 0x76, 0x6c, 0x2d, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_rawDescOnce sync.Once
	file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_rawDescData = file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_rawDesc
)

func file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_rawDescGZIP() []byte {
	file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_rawDescOnce.Do(func() {
		file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_rawDescData)
	})
	return file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_rawDescData
}

var file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_goTypes = []interface{}{
	(*OrganizationalRolePermissionPolicyGetRequest)(nil),      // 0: account.aggregates.permission.grpc.OrganizationalRolePermissionPolicyGetRequest
	(*OrganizationalRolePermissionPolicyGetRequestData)(nil),  // 1: account.aggregates.permission.grpc.OrganizationalRolePermissionPolicyGetRequestData
	(*OrganizationalRolePermissionPolicyGetResponse)(nil),     // 2: account.aggregates.permission.grpc.OrganizationalRolePermissionPolicyGetResponse
	(*OrganizationalRolePermissionPolicyGetResponseData)(nil), // 3: account.aggregates.permission.grpc.OrganizationalRolePermissionPolicyGetResponseData
	(*OrganizationalRolePermissionPolicy)(nil),                // 4: account.aggregates.permission.grpc.OrganizationalRolePermissionPolicy
	(*grpc.GRPCRequest)(nil),                                  // 5: common.grpc.GRPCRequest
	(*grpc.GRPCResponse)(nil),                                 // 6: common.grpc.GRPCResponse
}
var file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_depIdxs = []int32{
	5, // 0: account.aggregates.permission.grpc.OrganizationalRolePermissionPolicyGetRequest.base:type_name -> common.grpc.GRPCRequest
	1, // 1: account.aggregates.permission.grpc.OrganizationalRolePermissionPolicyGetRequest.data:type_name -> account.aggregates.permission.grpc.OrganizationalRolePermissionPolicyGetRequestData
	6, // 2: account.aggregates.permission.grpc.OrganizationalRolePermissionPolicyGetResponse.base:type_name -> common.grpc.GRPCResponse
	3, // 3: account.aggregates.permission.grpc.OrganizationalRolePermissionPolicyGetResponse.data:type_name -> account.aggregates.permission.grpc.OrganizationalRolePermissionPolicyGetResponseData
	4, // 4: account.aggregates.permission.grpc.OrganizationalRolePermissionPolicyGetResponseData.organizational_role_permission_policy:type_name -> account.aggregates.permission.grpc.OrganizationalRolePermissionPolicy
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() {
	file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_init()
}
func file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_init() {
	if File_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationalRolePermissionPolicyGetRequest); i {
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
		file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationalRolePermissionPolicyGetRequestData); i {
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
		file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationalRolePermissionPolicyGetResponse); i {
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
		file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationalRolePermissionPolicyGetResponseData); i {
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
		file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationalRolePermissionPolicy); i {
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
			RawDescriptor: file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_goTypes,
		DependencyIndexes: file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_depIdxs,
		MessageInfos:      file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_msgTypes,
	}.Build()
	File_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto = out.File
	file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_rawDesc = nil
	file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_goTypes = nil
	file_proto_account_aggregates_permission_grpc_organizational_role_permission_policy_get_proto_depIdxs = nil
}