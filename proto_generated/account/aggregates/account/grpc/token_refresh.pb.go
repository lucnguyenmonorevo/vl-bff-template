// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/account/aggregates/account/grpc/token_refresh.proto

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
type TokenRefreshRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCRequest        `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *TokenRefreshRequestData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TokenRefreshRequest) Reset() {
	*x = TokenRefreshRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_token_refresh_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenRefreshRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRefreshRequest) ProtoMessage() {}

func (x *TokenRefreshRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_token_refresh_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRefreshRequest.ProtoReflect.Descriptor instead.
func (*TokenRefreshRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_token_refresh_proto_rawDescGZIP(), []int{0}
}

func (x *TokenRefreshRequest) GetBase() *grpc.GRPCRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *TokenRefreshRequest) GetData() *TokenRefreshRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type TokenRefreshRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationId uint64 `protobuf:"varint,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	Username       string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	RefreshToken   string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *TokenRefreshRequestData) Reset() {
	*x = TokenRefreshRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_token_refresh_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenRefreshRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRefreshRequestData) ProtoMessage() {}

func (x *TokenRefreshRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_token_refresh_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRefreshRequestData.ProtoReflect.Descriptor instead.
func (*TokenRefreshRequestData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_token_refresh_proto_rawDescGZIP(), []int{1}
}

func (x *TokenRefreshRequestData) GetOrganizationId() uint64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *TokenRefreshRequestData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *TokenRefreshRequestData) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

// Response
type TokenRefreshResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCResponse `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *Token             `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TokenRefreshResponse) Reset() {
	*x = TokenRefreshResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_token_refresh_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenRefreshResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRefreshResponse) ProtoMessage() {}

func (x *TokenRefreshResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_token_refresh_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRefreshResponse.ProtoReflect.Descriptor instead.
func (*TokenRefreshResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_token_refresh_proto_rawDescGZIP(), []int{2}
}

func (x *TokenRefreshResponse) GetBase() *grpc.GRPCResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *TokenRefreshResponse) GetData() *Token {
	if x != nil {
		return x.Data
	}
	return nil
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdToken        string `protobuf:"bytes,1,opt,name=id_token,json=idToken,proto3" json:"id_token,omitempty"`
	AccessToken    string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken   string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiredIn      int64  `protobuf:"varint,4,opt,name=expired_in,json=expiredIn,proto3" json:"expired_in,omitempty"`
	OrganizationId uint64 `protobuf:"varint,5,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	Username       string `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_token_refresh_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_token_refresh_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_token_refresh_proto_rawDescGZIP(), []int{3}
}

func (x *Token) GetIdToken() string {
	if x != nil {
		return x.IdToken
	}
	return ""
}

func (x *Token) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Token) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *Token) GetExpiredIn() int64 {
	if x != nil {
		return x.ExpiredIn
	}
	return 0
}

func (x *Token) GetOrganizationId() uint64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *Token) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_proto_account_aggregates_account_grpc_token_refresh_proto protoreflect.FileDescriptor

var file_proto_account_aggregates_account_grpc_token_refresh_proto_rawDesc = []byte{
	0x0a, 0x39, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1c, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x13, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x4c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x83,
	0x01, 0x0a, 0x17, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x81, 0x01, 0x0a, 0x14, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xce, 0x01, 0x0a, 0x05, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64,
	0x5f, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x49, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x4d, 0x5a, 0x4b, 0x76, 0x6c, 0x2d,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_aggregates_account_grpc_token_refresh_proto_rawDescOnce sync.Once
	file_proto_account_aggregates_account_grpc_token_refresh_proto_rawDescData = file_proto_account_aggregates_account_grpc_token_refresh_proto_rawDesc
)

func file_proto_account_aggregates_account_grpc_token_refresh_proto_rawDescGZIP() []byte {
	file_proto_account_aggregates_account_grpc_token_refresh_proto_rawDescOnce.Do(func() {
		file_proto_account_aggregates_account_grpc_token_refresh_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_aggregates_account_grpc_token_refresh_proto_rawDescData)
	})
	return file_proto_account_aggregates_account_grpc_token_refresh_proto_rawDescData
}

var file_proto_account_aggregates_account_grpc_token_refresh_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_account_aggregates_account_grpc_token_refresh_proto_goTypes = []interface{}{
	(*TokenRefreshRequest)(nil),     // 0: account.aggregates.account.grpc.TokenRefreshRequest
	(*TokenRefreshRequestData)(nil), // 1: account.aggregates.account.grpc.TokenRefreshRequestData
	(*TokenRefreshResponse)(nil),    // 2: account.aggregates.account.grpc.TokenRefreshResponse
	(*Token)(nil),                   // 3: account.aggregates.account.grpc.Token
	(*grpc.GRPCRequest)(nil),        // 4: common.grpc.GRPCRequest
	(*grpc.GRPCResponse)(nil),       // 5: common.grpc.GRPCResponse
}
var file_proto_account_aggregates_account_grpc_token_refresh_proto_depIdxs = []int32{
	4, // 0: account.aggregates.account.grpc.TokenRefreshRequest.base:type_name -> common.grpc.GRPCRequest
	1, // 1: account.aggregates.account.grpc.TokenRefreshRequest.data:type_name -> account.aggregates.account.grpc.TokenRefreshRequestData
	5, // 2: account.aggregates.account.grpc.TokenRefreshResponse.base:type_name -> common.grpc.GRPCResponse
	3, // 3: account.aggregates.account.grpc.TokenRefreshResponse.data:type_name -> account.aggregates.account.grpc.Token
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_account_aggregates_account_grpc_token_refresh_proto_init() }
func file_proto_account_aggregates_account_grpc_token_refresh_proto_init() {
	if File_proto_account_aggregates_account_grpc_token_refresh_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_account_aggregates_account_grpc_token_refresh_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenRefreshRequest); i {
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
		file_proto_account_aggregates_account_grpc_token_refresh_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenRefreshRequestData); i {
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
		file_proto_account_aggregates_account_grpc_token_refresh_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenRefreshResponse); i {
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
		file_proto_account_aggregates_account_grpc_token_refresh_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
			RawDescriptor: file_proto_account_aggregates_account_grpc_token_refresh_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_aggregates_account_grpc_token_refresh_proto_goTypes,
		DependencyIndexes: file_proto_account_aggregates_account_grpc_token_refresh_proto_depIdxs,
		MessageInfos:      file_proto_account_aggregates_account_grpc_token_refresh_proto_msgTypes,
	}.Build()
	File_proto_account_aggregates_account_grpc_token_refresh_proto = out.File
	file_proto_account_aggregates_account_grpc_token_refresh_proto_rawDesc = nil
	file_proto_account_aggregates_account_grpc_token_refresh_proto_goTypes = nil
	file_proto_account_aggregates_account_grpc_token_refresh_proto_depIdxs = nil
}
