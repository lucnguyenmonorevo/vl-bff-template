// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/account/aggregates/permission/grpc/permission_user_list.proto

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

type PermissionUserListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCRequest              `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *PermissionUserListRequestData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PermissionUserListRequest) Reset() {
	*x = PermissionUserListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_permission_user_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionUserListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionUserListRequest) ProtoMessage() {}

func (x *PermissionUserListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_permission_user_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionUserListRequest.ProtoReflect.Descriptor instead.
func (*PermissionUserListRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_permission_user_list_proto_rawDescGZIP(), []int{0}
}

func (x *PermissionUserListRequest) GetBase() *grpc.GRPCRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *PermissionUserListRequest) GetData() *PermissionUserListRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type PermissionUserListRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit          int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset         int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Ids            []uint64 `protobuf:"varint,3,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	SearchText     string   `protobuf:"bytes,4,opt,name=search_text,json=searchText,proto3" json:"search_text,omitempty"`
	UserId         uint64   `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrganizationId uint64   `protobuf:"varint,6,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
}

func (x *PermissionUserListRequestData) Reset() {
	*x = PermissionUserListRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_permission_user_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionUserListRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionUserListRequestData) ProtoMessage() {}

func (x *PermissionUserListRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_permission_user_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionUserListRequestData.ProtoReflect.Descriptor instead.
func (*PermissionUserListRequestData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_permission_user_list_proto_rawDescGZIP(), []int{1}
}

func (x *PermissionUserListRequestData) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PermissionUserListRequestData) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *PermissionUserListRequestData) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *PermissionUserListRequestData) GetSearchText() string {
	if x != nil {
		return x.SearchText
	}
	return ""
}

func (x *PermissionUserListRequestData) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PermissionUserListRequestData) GetOrganizationId() uint64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

type PermissionUserListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCResponse              `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *PermissionUserListResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PermissionUserListResponse) Reset() {
	*x = PermissionUserListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_permission_user_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionUserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionUserListResponse) ProtoMessage() {}

func (x *PermissionUserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_permission_user_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionUserListResponse.ProtoReflect.Descriptor instead.
func (*PermissionUserListResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_permission_user_list_proto_rawDescGZIP(), []int{2}
}

func (x *PermissionUserListResponse) GetBase() *grpc.GRPCResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *PermissionUserListResponse) GetData() *PermissionUserListResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type PermissionUserListResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PermissionUsers []*PermissionUserGetResponseData `protobuf:"bytes,1,rep,name=permission_users,json=permissionUsers,proto3" json:"permission_users,omitempty"`
}

func (x *PermissionUserListResponseData) Reset() {
	*x = PermissionUserListResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_permission_user_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionUserListResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionUserListResponseData) ProtoMessage() {}

func (x *PermissionUserListResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_permission_user_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionUserListResponseData.ProtoReflect.Descriptor instead.
func (*PermissionUserListResponseData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_permission_user_list_proto_rawDescGZIP(), []int{3}
}

func (x *PermissionUserListResponseData) GetPermissionUsers() []*PermissionUserGetResponseData {
	if x != nil {
		return x.PermissionUsers
	}
	return nil
}

var File_proto_account_aggregates_permission_grpc_permission_user_list_proto protoreflect.FileDescriptor

var file_proto_account_aggregates_permission_grpc_permission_user_list_proto_rawDesc = []byte{
	0x0a, 0x43, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x42, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x19,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc2,
	0x01, 0x0a, 0x1d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x04, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0xa3, 0x01, 0x0a, 0x1a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x56, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x42, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x8e, 0x01, 0x0a, 0x1e, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x6c, 0x0a, 0x10,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x53, 0x5a, 0x51, 0x76, 0x6c,
	0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x67, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_aggregates_permission_grpc_permission_user_list_proto_rawDescOnce sync.Once
	file_proto_account_aggregates_permission_grpc_permission_user_list_proto_rawDescData = file_proto_account_aggregates_permission_grpc_permission_user_list_proto_rawDesc
)

func file_proto_account_aggregates_permission_grpc_permission_user_list_proto_rawDescGZIP() []byte {
	file_proto_account_aggregates_permission_grpc_permission_user_list_proto_rawDescOnce.Do(func() {
		file_proto_account_aggregates_permission_grpc_permission_user_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_aggregates_permission_grpc_permission_user_list_proto_rawDescData)
	})
	return file_proto_account_aggregates_permission_grpc_permission_user_list_proto_rawDescData
}

var file_proto_account_aggregates_permission_grpc_permission_user_list_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_account_aggregates_permission_grpc_permission_user_list_proto_goTypes = []interface{}{
	(*PermissionUserListRequest)(nil),      // 0: account.aggregates.permission.grpc.PermissionUserListRequest
	(*PermissionUserListRequestData)(nil),  // 1: account.aggregates.permission.grpc.PermissionUserListRequestData
	(*PermissionUserListResponse)(nil),     // 2: account.aggregates.permission.grpc.PermissionUserListResponse
	(*PermissionUserListResponseData)(nil), // 3: account.aggregates.permission.grpc.PermissionUserListResponseData
	(*grpc.GRPCRequest)(nil),               // 4: common.grpc.GRPCRequest
	(*grpc.GRPCResponse)(nil),              // 5: common.grpc.GRPCResponse
	(*PermissionUserGetResponseData)(nil),  // 6: account.aggregates.permission.grpc.PermissionUserGetResponseData
}
var file_proto_account_aggregates_permission_grpc_permission_user_list_proto_depIdxs = []int32{
	4, // 0: account.aggregates.permission.grpc.PermissionUserListRequest.base:type_name -> common.grpc.GRPCRequest
	1, // 1: account.aggregates.permission.grpc.PermissionUserListRequest.data:type_name -> account.aggregates.permission.grpc.PermissionUserListRequestData
	5, // 2: account.aggregates.permission.grpc.PermissionUserListResponse.base:type_name -> common.grpc.GRPCResponse
	3, // 3: account.aggregates.permission.grpc.PermissionUserListResponse.data:type_name -> account.aggregates.permission.grpc.PermissionUserListResponseData
	6, // 4: account.aggregates.permission.grpc.PermissionUserListResponseData.permission_users:type_name -> account.aggregates.permission.grpc.PermissionUserGetResponseData
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_account_aggregates_permission_grpc_permission_user_list_proto_init() }
func file_proto_account_aggregates_permission_grpc_permission_user_list_proto_init() {
	if File_proto_account_aggregates_permission_grpc_permission_user_list_proto != nil {
		return
	}
	file_proto_account_aggregates_permission_grpc_permission_user_get_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_account_aggregates_permission_grpc_permission_user_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionUserListRequest); i {
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
		file_proto_account_aggregates_permission_grpc_permission_user_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionUserListRequestData); i {
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
		file_proto_account_aggregates_permission_grpc_permission_user_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionUserListResponse); i {
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
		file_proto_account_aggregates_permission_grpc_permission_user_list_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionUserListResponseData); i {
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
			RawDescriptor: file_proto_account_aggregates_permission_grpc_permission_user_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_aggregates_permission_grpc_permission_user_list_proto_goTypes,
		DependencyIndexes: file_proto_account_aggregates_permission_grpc_permission_user_list_proto_depIdxs,
		MessageInfos:      file_proto_account_aggregates_permission_grpc_permission_user_list_proto_msgTypes,
	}.Build()
	File_proto_account_aggregates_permission_grpc_permission_user_list_proto = out.File
	file_proto_account_aggregates_permission_grpc_permission_user_list_proto_rawDesc = nil
	file_proto_account_aggregates_permission_grpc_permission_user_list_proto_goTypes = nil
	file_proto_account_aggregates_permission_grpc_permission_user_list_proto_depIdxs = nil
}
