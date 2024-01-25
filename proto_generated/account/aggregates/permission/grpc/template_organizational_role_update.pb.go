// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/account/aggregates/permission/grpc/template_organizational_role_update.proto

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

type TemplateOrganizationalRoleUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCRequest                            `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *TemplateOrganizationalRoleUpdateRequestData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TemplateOrganizationalRoleUpdateRequest) Reset() {
	*x = TemplateOrganizationalRoleUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateOrganizationalRoleUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateOrganizationalRoleUpdateRequest) ProtoMessage() {}

func (x *TemplateOrganizationalRoleUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateOrganizationalRoleUpdateRequest.ProtoReflect.Descriptor instead.
func (*TemplateOrganizationalRoleUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_rawDescGZIP(), []int{0}
}

func (x *TemplateOrganizationalRoleUpdateRequest) GetBase() *grpc.GRPCRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *TemplateOrganizationalRoleUpdateRequest) GetData() *TemplateOrganizationalRoleUpdateRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type TemplateOrganizationalRoleUpdateRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TemplateOrganizationalRoleUpdateRequestData) Reset() {
	*x = TemplateOrganizationalRoleUpdateRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateOrganizationalRoleUpdateRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateOrganizationalRoleUpdateRequestData) ProtoMessage() {}

func (x *TemplateOrganizationalRoleUpdateRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateOrganizationalRoleUpdateRequestData.ProtoReflect.Descriptor instead.
func (*TemplateOrganizationalRoleUpdateRequestData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_rawDescGZIP(), []int{1}
}

func (x *TemplateOrganizationalRoleUpdateRequestData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TemplateOrganizationalRoleUpdateRequestData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TemplateOrganizationalRoleUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCResponse                            `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *TemplateOrganizationalRoleUpdateResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TemplateOrganizationalRoleUpdateResponse) Reset() {
	*x = TemplateOrganizationalRoleUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateOrganizationalRoleUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateOrganizationalRoleUpdateResponse) ProtoMessage() {}

func (x *TemplateOrganizationalRoleUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateOrganizationalRoleUpdateResponse.ProtoReflect.Descriptor instead.
func (*TemplateOrganizationalRoleUpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_rawDescGZIP(), []int{2}
}

func (x *TemplateOrganizationalRoleUpdateResponse) GetBase() *grpc.GRPCResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *TemplateOrganizationalRoleUpdateResponse) GetData() *TemplateOrganizationalRoleUpdateResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type TemplateOrganizationalRoleUpdateResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TemplateOrganizationalRoleUpdateResponseData) Reset() {
	*x = TemplateOrganizationalRoleUpdateResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateOrganizationalRoleUpdateResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateOrganizationalRoleUpdateResponseData) ProtoMessage() {}

func (x *TemplateOrganizationalRoleUpdateResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateOrganizationalRoleUpdateResponseData.ProtoReflect.Descriptor instead.
func (*TemplateOrganizationalRoleUpdateResponseData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_rawDescGZIP(), []int{3}
}

func (x *TemplateOrganizationalRoleUpdateResponseData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto protoreflect.FileDescriptor

var file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_rawDesc = []byte{
	0x0a, 0x52, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x5f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x27, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x63, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4f,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x51, 0x0a, 0x2b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52,
	0x6f, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xbf, 0x01, 0x0a, 0x28, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04,
	0x62, 0x61, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x50, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f,
	0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3e, 0x0a, 0x2c, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x42, 0x53, 0x5a, 0x51, 0x76, 0x6c,
	0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x67, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_rawDescOnce sync.Once
	file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_rawDescData = file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_rawDesc
)

func file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_rawDescGZIP() []byte {
	file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_rawDescOnce.Do(func() {
		file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_rawDescData)
	})
	return file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_rawDescData
}

var file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_goTypes = []interface{}{
	(*TemplateOrganizationalRoleUpdateRequest)(nil),      // 0: account.aggregates.permission.grpc.TemplateOrganizationalRoleUpdateRequest
	(*TemplateOrganizationalRoleUpdateRequestData)(nil),  // 1: account.aggregates.permission.grpc.TemplateOrganizationalRoleUpdateRequestData
	(*TemplateOrganizationalRoleUpdateResponse)(nil),     // 2: account.aggregates.permission.grpc.TemplateOrganizationalRoleUpdateResponse
	(*TemplateOrganizationalRoleUpdateResponseData)(nil), // 3: account.aggregates.permission.grpc.TemplateOrganizationalRoleUpdateResponseData
	(*grpc.GRPCRequest)(nil),                             // 4: common.grpc.GRPCRequest
	(*grpc.GRPCResponse)(nil),                            // 5: common.grpc.GRPCResponse
}
var file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_depIdxs = []int32{
	4, // 0: account.aggregates.permission.grpc.TemplateOrganizationalRoleUpdateRequest.base:type_name -> common.grpc.GRPCRequest
	1, // 1: account.aggregates.permission.grpc.TemplateOrganizationalRoleUpdateRequest.data:type_name -> account.aggregates.permission.grpc.TemplateOrganizationalRoleUpdateRequestData
	5, // 2: account.aggregates.permission.grpc.TemplateOrganizationalRoleUpdateResponse.base:type_name -> common.grpc.GRPCResponse
	3, // 3: account.aggregates.permission.grpc.TemplateOrganizationalRoleUpdateResponse.data:type_name -> account.aggregates.permission.grpc.TemplateOrganizationalRoleUpdateResponseData
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_init()
}
func file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_init() {
	if File_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateOrganizationalRoleUpdateRequest); i {
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
		file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateOrganizationalRoleUpdateRequestData); i {
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
		file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateOrganizationalRoleUpdateResponse); i {
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
		file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateOrganizationalRoleUpdateResponseData); i {
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
			RawDescriptor: file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_goTypes,
		DependencyIndexes: file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_depIdxs,
		MessageInfos:      file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_msgTypes,
	}.Build()
	File_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto = out.File
	file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_rawDesc = nil
	file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_goTypes = nil
	file_proto_account_aggregates_permission_grpc_template_organizational_role_update_proto_depIdxs = nil
}
