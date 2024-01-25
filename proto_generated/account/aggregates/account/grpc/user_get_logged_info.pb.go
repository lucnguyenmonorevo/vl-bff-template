// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/account/aggregates/account/grpc/user_get_logged_info.proto

package protoaccountgrpc

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

// Request
type UserGetLoggedInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCRequest             `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *UserGetLoggedInfoRequestData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UserGetLoggedInfoRequest) Reset() {
	*x = UserGetLoggedInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGetLoggedInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGetLoggedInfoRequest) ProtoMessage() {}

func (x *UserGetLoggedInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGetLoggedInfoRequest.ProtoReflect.Descriptor instead.
func (*UserGetLoggedInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_rawDescGZIP(), []int{0}
}

func (x *UserGetLoggedInfoRequest) GetBase() *grpc.GRPCRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *UserGetLoggedInfoRequest) GetData() *UserGetLoggedInfoRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type UserGetLoggedInfoRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *UserGetLoggedInfoRequestData) Reset() {
	*x = UserGetLoggedInfoRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGetLoggedInfoRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGetLoggedInfoRequestData) ProtoMessage() {}

func (x *UserGetLoggedInfoRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGetLoggedInfoRequestData.ProtoReflect.Descriptor instead.
func (*UserGetLoggedInfoRequestData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_rawDescGZIP(), []int{1}
}

func (x *UserGetLoggedInfoRequestData) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

// Response
type UserGetLoggedInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCResponse             `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *UserGetLoggedInfoResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UserGetLoggedInfoResponse) Reset() {
	*x = UserGetLoggedInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGetLoggedInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGetLoggedInfoResponse) ProtoMessage() {}

func (x *UserGetLoggedInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGetLoggedInfoResponse.ProtoReflect.Descriptor instead.
func (*UserGetLoggedInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_rawDescGZIP(), []int{2}
}

func (x *UserGetLoggedInfoResponse) GetBase() *grpc.GRPCResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *UserGetLoggedInfoResponse) GetData() *UserGetLoggedInfoResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type UserGetLoggedInfoResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo     *UserInfo     `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	Organization *Organization `protobuf:"bytes,2,opt,name=organization,proto3" json:"organization,omitempty"`
	Departments  []*Department `protobuf:"bytes,3,rep,name=departments,proto3" json:"departments,omitempty"`
	Bases        []*Base       `protobuf:"bytes,4,rep,name=bases,proto3" json:"bases,omitempty"`
	Bookmarks    []*Bookmark   `protobuf:"bytes,5,rep,name=bookmarks,proto3" json:"bookmarks,omitempty"`
}

func (x *UserGetLoggedInfoResponseData) Reset() {
	*x = UserGetLoggedInfoResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGetLoggedInfoResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGetLoggedInfoResponseData) ProtoMessage() {}

func (x *UserGetLoggedInfoResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGetLoggedInfoResponseData.ProtoReflect.Descriptor instead.
func (*UserGetLoggedInfoResponseData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_rawDescGZIP(), []int{3}
}

func (x *UserGetLoggedInfoResponseData) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *UserGetLoggedInfoResponseData) GetOrganization() *Organization {
	if x != nil {
		return x.Organization
	}
	return nil
}

func (x *UserGetLoggedInfoResponseData) GetDepartments() []*Department {
	if x != nil {
		return x.Departments
	}
	return nil
}

func (x *UserGetLoggedInfoResponseData) GetBases() []*Base {
	if x != nil {
		return x.Bases
	}
	return nil
}

func (x *UserGetLoggedInfoResponseData) GetBookmarks() []*Bookmark {
	if x != nil {
		return x.Bookmarks
	}
	return nil
}

type Bookmark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Bookmark) Reset() {
	*x = Bookmark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bookmark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bookmark) ProtoMessage() {}

func (x *Bookmark) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bookmark.ProtoReflect.Descriptor instead.
func (*Bookmark) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_rawDescGZIP(), []int{4}
}

func (x *Bookmark) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Bookmark) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_proto_account_aggregates_account_grpc_user_get_logged_info_proto protoreflect.FileDescriptor

var file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_rawDesc = []byte{
	0x0a, 0x40, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x65, 0x74,
	0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x30, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x38, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x01, 0x0a,
	0x18, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x4c,
	0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x41, 0x0a, 0x1c, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9e, 0x01,
	0x0a, 0x19, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47,
	0x65, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x8f,
	0x03, 0x0a, 0x1d, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x46, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x51, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x0b, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x05, 0x62, 0x61,
	0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x05, 0x62, 0x61, 0x73, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x6d,
	0x61, 0x72, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73,
	0x22, 0x33, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x4d, 0x5a, 0x4b, 0x76, 0x6c, 0x2d, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_rawDescOnce sync.Once
	file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_rawDescData = file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_rawDesc
)

func file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_rawDescGZIP() []byte {
	file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_rawDescOnce.Do(func() {
		file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_rawDescData)
	})
	return file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_rawDescData
}

var file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_goTypes = []interface{}{
	(*UserGetLoggedInfoRequest)(nil),      // 0: account.aggregates.account.grpc.UserGetLoggedInfoRequest
	(*UserGetLoggedInfoRequestData)(nil),  // 1: account.aggregates.account.grpc.UserGetLoggedInfoRequestData
	(*UserGetLoggedInfoResponse)(nil),     // 2: account.aggregates.account.grpc.UserGetLoggedInfoResponse
	(*UserGetLoggedInfoResponseData)(nil), // 3: account.aggregates.account.grpc.UserGetLoggedInfoResponseData
	(*Bookmark)(nil),                      // 4: account.aggregates.account.grpc.Bookmark
	(*grpc.GRPCRequest)(nil),              // 5: common.grpc.GRPCRequest
	(*grpc.GRPCResponse)(nil),             // 6: common.grpc.GRPCResponse
	(*UserInfo)(nil),                      // 7: account.aggregates.account.grpc.UserInfo
	(*Organization)(nil),                  // 8: account.aggregates.account.grpc.Organization
	(*Department)(nil),                    // 9: account.aggregates.account.grpc.Department
	(*Base)(nil),                          // 10: account.aggregates.account.grpc.Base
}
var file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_depIdxs = []int32{
	5,  // 0: account.aggregates.account.grpc.UserGetLoggedInfoRequest.base:type_name -> common.grpc.GRPCRequest
	1,  // 1: account.aggregates.account.grpc.UserGetLoggedInfoRequest.data:type_name -> account.aggregates.account.grpc.UserGetLoggedInfoRequestData
	6,  // 2: account.aggregates.account.grpc.UserGetLoggedInfoResponse.base:type_name -> common.grpc.GRPCResponse
	3,  // 3: account.aggregates.account.grpc.UserGetLoggedInfoResponse.data:type_name -> account.aggregates.account.grpc.UserGetLoggedInfoResponseData
	7,  // 4: account.aggregates.account.grpc.UserGetLoggedInfoResponseData.user_info:type_name -> account.aggregates.account.grpc.UserInfo
	8,  // 5: account.aggregates.account.grpc.UserGetLoggedInfoResponseData.organization:type_name -> account.aggregates.account.grpc.Organization
	9,  // 6: account.aggregates.account.grpc.UserGetLoggedInfoResponseData.departments:type_name -> account.aggregates.account.grpc.Department
	10, // 7: account.aggregates.account.grpc.UserGetLoggedInfoResponseData.bases:type_name -> account.aggregates.account.grpc.Base
	4,  // 8: account.aggregates.account.grpc.UserGetLoggedInfoResponseData.bookmarks:type_name -> account.aggregates.account.grpc.Bookmark
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_init() }
func file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_init() {
	if File_proto_account_aggregates_account_grpc_user_get_logged_info_proto != nil {
		return
	}
	file_proto_account_aggregates_account_grpc_base_proto_init()
	file_proto_account_aggregates_account_grpc_department_proto_init()
	file_proto_account_aggregates_account_grpc_user_info_proto_init()
	file_proto_account_aggregates_account_grpc_organization_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGetLoggedInfoRequest); i {
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
		file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGetLoggedInfoRequestData); i {
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
		file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGetLoggedInfoResponse); i {
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
		file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGetLoggedInfoResponseData); i {
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
		file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bookmark); i {
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
			RawDescriptor: file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_goTypes,
		DependencyIndexes: file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_depIdxs,
		MessageInfos:      file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_msgTypes,
	}.Build()
	File_proto_account_aggregates_account_grpc_user_get_logged_info_proto = out.File
	file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_rawDesc = nil
	file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_goTypes = nil
	file_proto_account_aggregates_account_grpc_user_get_logged_info_proto_depIdxs = nil
}
