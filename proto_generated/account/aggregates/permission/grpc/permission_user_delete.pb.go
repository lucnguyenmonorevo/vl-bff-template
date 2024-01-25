// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/account/aggregates/permission/grpc/permission_user_delete.proto

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

type PermissionUserDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCRequest                `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *PermissionUserDeleteRequestData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PermissionUserDeleteRequest) Reset() {
	*x = PermissionUserDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionUserDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionUserDeleteRequest) ProtoMessage() {}

func (x *PermissionUserDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionUserDeleteRequest.ProtoReflect.Descriptor instead.
func (*PermissionUserDeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_rawDescGZIP(), []int{0}
}

func (x *PermissionUserDeleteRequest) GetBase() *grpc.GRPCRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *PermissionUserDeleteRequest) GetData() *PermissionUserDeleteRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type PermissionUserDeleteRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PermissionUserDeleteRequestData) Reset() {
	*x = PermissionUserDeleteRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionUserDeleteRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionUserDeleteRequestData) ProtoMessage() {}

func (x *PermissionUserDeleteRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionUserDeleteRequestData.ProtoReflect.Descriptor instead.
func (*PermissionUserDeleteRequestData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_rawDescGZIP(), []int{1}
}

func (x *PermissionUserDeleteRequestData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PermissionUserDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCResponse                `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *PermissionUserDeleteResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PermissionUserDeleteResponse) Reset() {
	*x = PermissionUserDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionUserDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionUserDeleteResponse) ProtoMessage() {}

func (x *PermissionUserDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionUserDeleteResponse.ProtoReflect.Descriptor instead.
func (*PermissionUserDeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_rawDescGZIP(), []int{2}
}

func (x *PermissionUserDeleteResponse) GetBase() *grpc.GRPCResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *PermissionUserDeleteResponse) GetData() *PermissionUserDeleteResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type PermissionUserDeleteResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PermissionUserDeleteResponseData) Reset() {
	*x = PermissionUserDeleteResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionUserDeleteResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionUserDeleteResponseData) ProtoMessage() {}

func (x *PermissionUserDeleteResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionUserDeleteResponseData.ProtoReflect.Descriptor instead.
func (*PermissionUserDeleteResponseData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_rawDescGZIP(), []int{3}
}

func (x *PermissionUserDeleteResponseData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_proto_account_aggregates_permission_grpc_permission_user_delete_proto protoreflect.FileDescriptor

var file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_rawDesc = []byte{
	0x0a, 0x45, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1c, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x01, 0x0a, 0x1b, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x31, 0x0a, 0x1f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xa7, 0x01, 0x0a, 0x1c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x44, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x32, 0x0a,
	0x20, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x42, 0x53, 0x5a, 0x51, 0x76, 0x6c, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_rawDescOnce sync.Once
	file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_rawDescData = file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_rawDesc
)

func file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_rawDescGZIP() []byte {
	file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_rawDescOnce.Do(func() {
		file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_rawDescData)
	})
	return file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_rawDescData
}

var file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_goTypes = []interface{}{
	(*PermissionUserDeleteRequest)(nil),      // 0: account.aggregates.permission.grpc.PermissionUserDeleteRequest
	(*PermissionUserDeleteRequestData)(nil),  // 1: account.aggregates.permission.grpc.PermissionUserDeleteRequestData
	(*PermissionUserDeleteResponse)(nil),     // 2: account.aggregates.permission.grpc.PermissionUserDeleteResponse
	(*PermissionUserDeleteResponseData)(nil), // 3: account.aggregates.permission.grpc.PermissionUserDeleteResponseData
	(*grpc.GRPCRequest)(nil),                 // 4: common.grpc.GRPCRequest
	(*grpc.GRPCResponse)(nil),                // 5: common.grpc.GRPCResponse
}
var file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_depIdxs = []int32{
	4, // 0: account.aggregates.permission.grpc.PermissionUserDeleteRequest.base:type_name -> common.grpc.GRPCRequest
	1, // 1: account.aggregates.permission.grpc.PermissionUserDeleteRequest.data:type_name -> account.aggregates.permission.grpc.PermissionUserDeleteRequestData
	5, // 2: account.aggregates.permission.grpc.PermissionUserDeleteResponse.base:type_name -> common.grpc.GRPCResponse
	3, // 3: account.aggregates.permission.grpc.PermissionUserDeleteResponse.data:type_name -> account.aggregates.permission.grpc.PermissionUserDeleteResponseData
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_init() }
func file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_init() {
	if File_proto_account_aggregates_permission_grpc_permission_user_delete_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionUserDeleteRequest); i {
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
		file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionUserDeleteRequestData); i {
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
		file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionUserDeleteResponse); i {
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
		file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionUserDeleteResponseData); i {
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
			RawDescriptor: file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_goTypes,
		DependencyIndexes: file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_depIdxs,
		MessageInfos:      file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_msgTypes,
	}.Build()
	File_proto_account_aggregates_permission_grpc_permission_user_delete_proto = out.File
	file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_rawDesc = nil
	file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_goTypes = nil
	file_proto_account_aggregates_permission_grpc_permission_user_delete_proto_depIdxs = nil
}