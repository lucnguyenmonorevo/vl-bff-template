// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/account/aggregates/permission/grpc/organization_subscription_delete.proto

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

type OrganizationSubscriptionDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCRequest                          `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *OrganizationSubscriptionDeleteRequestData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OrganizationSubscriptionDeleteRequest) Reset() {
	*x = OrganizationSubscriptionDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationSubscriptionDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationSubscriptionDeleteRequest) ProtoMessage() {}

func (x *OrganizationSubscriptionDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationSubscriptionDeleteRequest.ProtoReflect.Descriptor instead.
func (*OrganizationSubscriptionDeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_rawDescGZIP(), []int{0}
}

func (x *OrganizationSubscriptionDeleteRequest) GetBase() *grpc.GRPCRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *OrganizationSubscriptionDeleteRequest) GetData() *OrganizationSubscriptionDeleteRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type OrganizationSubscriptionDeleteRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrganizationSubscriptionDeleteRequestData) Reset() {
	*x = OrganizationSubscriptionDeleteRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationSubscriptionDeleteRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationSubscriptionDeleteRequestData) ProtoMessage() {}

func (x *OrganizationSubscriptionDeleteRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationSubscriptionDeleteRequestData.ProtoReflect.Descriptor instead.
func (*OrganizationSubscriptionDeleteRequestData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_rawDescGZIP(), []int{1}
}

func (x *OrganizationSubscriptionDeleteRequestData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OrganizationSubscriptionDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCResponse                          `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *OrganizationSubscriptionDeleteResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OrganizationSubscriptionDeleteResponse) Reset() {
	*x = OrganizationSubscriptionDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationSubscriptionDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationSubscriptionDeleteResponse) ProtoMessage() {}

func (x *OrganizationSubscriptionDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationSubscriptionDeleteResponse.ProtoReflect.Descriptor instead.
func (*OrganizationSubscriptionDeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_rawDescGZIP(), []int{2}
}

func (x *OrganizationSubscriptionDeleteResponse) GetBase() *grpc.GRPCResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *OrganizationSubscriptionDeleteResponse) GetData() *OrganizationSubscriptionDeleteResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type OrganizationSubscriptionDeleteResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrganizationSubscriptionDeleteResponseData) Reset() {
	*x = OrganizationSubscriptionDeleteResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationSubscriptionDeleteResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationSubscriptionDeleteResponseData) ProtoMessage() {}

func (x *OrganizationSubscriptionDeleteResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationSubscriptionDeleteResponseData.ProtoReflect.Descriptor instead.
func (*OrganizationSubscriptionDeleteResponseData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_rawDescGZIP(), []int{3}
}

func (x *OrganizationSubscriptionDeleteResponseData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto protoreflect.FileDescriptor

var file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x22, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x25, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3b,
	0x0a, 0x29, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0xbb, 0x01, 0x0a, 0x26,
	0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3c, 0x0a, 0x2a, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x42, 0x53, 0x5a, 0x51, 0x76, 0x6c, 0x2d, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_rawDescOnce sync.Once
	file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_rawDescData = file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_rawDesc
)

func file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_rawDescGZIP() []byte {
	file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_rawDescOnce.Do(func() {
		file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_rawDescData)
	})
	return file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_rawDescData
}

var file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_goTypes = []interface{}{
	(*OrganizationSubscriptionDeleteRequest)(nil),      // 0: account.aggregates.permission.grpc.OrganizationSubscriptionDeleteRequest
	(*OrganizationSubscriptionDeleteRequestData)(nil),  // 1: account.aggregates.permission.grpc.OrganizationSubscriptionDeleteRequestData
	(*OrganizationSubscriptionDeleteResponse)(nil),     // 2: account.aggregates.permission.grpc.OrganizationSubscriptionDeleteResponse
	(*OrganizationSubscriptionDeleteResponseData)(nil), // 3: account.aggregates.permission.grpc.OrganizationSubscriptionDeleteResponseData
	(*grpc.GRPCRequest)(nil),                           // 4: common.grpc.GRPCRequest
	(*grpc.GRPCResponse)(nil),                          // 5: common.grpc.GRPCResponse
}
var file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_depIdxs = []int32{
	4, // 0: account.aggregates.permission.grpc.OrganizationSubscriptionDeleteRequest.base:type_name -> common.grpc.GRPCRequest
	1, // 1: account.aggregates.permission.grpc.OrganizationSubscriptionDeleteRequest.data:type_name -> account.aggregates.permission.grpc.OrganizationSubscriptionDeleteRequestData
	5, // 2: account.aggregates.permission.grpc.OrganizationSubscriptionDeleteResponse.base:type_name -> common.grpc.GRPCResponse
	3, // 3: account.aggregates.permission.grpc.OrganizationSubscriptionDeleteResponse.data:type_name -> account.aggregates.permission.grpc.OrganizationSubscriptionDeleteResponseData
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_init()
}
func file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_init() {
	if File_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationSubscriptionDeleteRequest); i {
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
		file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationSubscriptionDeleteRequestData); i {
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
		file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationSubscriptionDeleteResponse); i {
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
		file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationSubscriptionDeleteResponseData); i {
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
			RawDescriptor: file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_goTypes,
		DependencyIndexes: file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_depIdxs,
		MessageInfos:      file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_msgTypes,
	}.Build()
	File_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto = out.File
	file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_rawDesc = nil
	file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_goTypes = nil
	file_proto_account_aggregates_permission_grpc_organization_subscription_delete_proto_depIdxs = nil
}
