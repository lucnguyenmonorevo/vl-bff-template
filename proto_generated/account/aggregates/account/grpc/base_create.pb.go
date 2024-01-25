// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/account/aggregates/account/grpc/base_create.proto

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
type BaseCreateRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OrganizationId uint64 `protobuf:"varint,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	Tel            string `protobuf:"bytes,3,opt,name=tel,proto3" json:"tel,omitempty"`
	Fax            string `protobuf:"bytes,4,opt,name=fax,proto3" json:"fax,omitempty"`
}

func (x *BaseCreateRequestData) Reset() {
	*x = BaseCreateRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_base_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseCreateRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseCreateRequestData) ProtoMessage() {}

func (x *BaseCreateRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_base_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseCreateRequestData.ProtoReflect.Descriptor instead.
func (*BaseCreateRequestData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_base_create_proto_rawDescGZIP(), []int{0}
}

func (x *BaseCreateRequestData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BaseCreateRequestData) GetOrganizationId() uint64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *BaseCreateRequestData) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *BaseCreateRequestData) GetFax() string {
	if x != nil {
		return x.Fax
	}
	return ""
}

type BaseCreateResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BaseCreateResponseData) Reset() {
	*x = BaseCreateResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_base_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseCreateResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseCreateResponseData) ProtoMessage() {}

func (x *BaseCreateResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_base_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseCreateResponseData.ProtoReflect.Descriptor instead.
func (*BaseCreateResponseData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_base_create_proto_rawDescGZIP(), []int{1}
}

func (x *BaseCreateResponseData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Response
type BaseCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCRequest      `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *BaseCreateRequestData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BaseCreateRequest) Reset() {
	*x = BaseCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_base_create_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseCreateRequest) ProtoMessage() {}

func (x *BaseCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_base_create_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseCreateRequest.ProtoReflect.Descriptor instead.
func (*BaseCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_base_create_proto_rawDescGZIP(), []int{2}
}

func (x *BaseCreateRequest) GetBase() *grpc.GRPCRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *BaseCreateRequest) GetData() *BaseCreateRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type BaseCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCResponse      `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *BaseCreateResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BaseCreateResponse) Reset() {
	*x = BaseCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_base_create_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseCreateResponse) ProtoMessage() {}

func (x *BaseCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_base_create_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseCreateResponse.ProtoReflect.Descriptor instead.
func (*BaseCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_base_create_proto_rawDescGZIP(), []int{3}
}

func (x *BaseCreateResponse) GetBase() *grpc.GRPCResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *BaseCreateResponse) GetData() *BaseCreateResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_account_aggregates_account_grpc_base_create_proto protoreflect.FileDescriptor

var file_proto_account_aggregates_account_grpc_base_create_proto_rawDesc = []byte{
	0x0a, 0x37, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x15, 0x42, 0x61, 0x73, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x65, 0x6c,
	0x12, 0x10, 0x0a, 0x03, 0x66, 0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66,
	0x61, 0x78, 0x22, 0x28, 0x0a, 0x16, 0x42, 0x61, 0x73, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8d, 0x01, 0x0a,
	0x11, 0x42, 0x61, 0x73, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x4a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x90, 0x01, 0x0a,
	0x12, 0x42, 0x61, 0x73, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x37, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42,
	0x4d, 0x5a, 0x4b, 0x76, 0x6c, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_aggregates_account_grpc_base_create_proto_rawDescOnce sync.Once
	file_proto_account_aggregates_account_grpc_base_create_proto_rawDescData = file_proto_account_aggregates_account_grpc_base_create_proto_rawDesc
)

func file_proto_account_aggregates_account_grpc_base_create_proto_rawDescGZIP() []byte {
	file_proto_account_aggregates_account_grpc_base_create_proto_rawDescOnce.Do(func() {
		file_proto_account_aggregates_account_grpc_base_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_aggregates_account_grpc_base_create_proto_rawDescData)
	})
	return file_proto_account_aggregates_account_grpc_base_create_proto_rawDescData
}

var file_proto_account_aggregates_account_grpc_base_create_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_account_aggregates_account_grpc_base_create_proto_goTypes = []interface{}{
	(*BaseCreateRequestData)(nil),  // 0: account.aggregates.account.grpc.BaseCreateRequestData
	(*BaseCreateResponseData)(nil), // 1: account.aggregates.account.grpc.BaseCreateResponseData
	(*BaseCreateRequest)(nil),      // 2: account.aggregates.account.grpc.BaseCreateRequest
	(*BaseCreateResponse)(nil),     // 3: account.aggregates.account.grpc.BaseCreateResponse
	(*grpc.GRPCRequest)(nil),       // 4: common.grpc.GRPCRequest
	(*grpc.GRPCResponse)(nil),      // 5: common.grpc.GRPCResponse
}
var file_proto_account_aggregates_account_grpc_base_create_proto_depIdxs = []int32{
	4, // 0: account.aggregates.account.grpc.BaseCreateRequest.base:type_name -> common.grpc.GRPCRequest
	0, // 1: account.aggregates.account.grpc.BaseCreateRequest.data:type_name -> account.aggregates.account.grpc.BaseCreateRequestData
	5, // 2: account.aggregates.account.grpc.BaseCreateResponse.base:type_name -> common.grpc.GRPCResponse
	1, // 3: account.aggregates.account.grpc.BaseCreateResponse.data:type_name -> account.aggregates.account.grpc.BaseCreateResponseData
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_account_aggregates_account_grpc_base_create_proto_init() }
func file_proto_account_aggregates_account_grpc_base_create_proto_init() {
	if File_proto_account_aggregates_account_grpc_base_create_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_account_aggregates_account_grpc_base_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseCreateRequestData); i {
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
		file_proto_account_aggregates_account_grpc_base_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseCreateResponseData); i {
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
		file_proto_account_aggregates_account_grpc_base_create_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseCreateRequest); i {
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
		file_proto_account_aggregates_account_grpc_base_create_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseCreateResponse); i {
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
			RawDescriptor: file_proto_account_aggregates_account_grpc_base_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_aggregates_account_grpc_base_create_proto_goTypes,
		DependencyIndexes: file_proto_account_aggregates_account_grpc_base_create_proto_depIdxs,
		MessageInfos:      file_proto_account_aggregates_account_grpc_base_create_proto_msgTypes,
	}.Build()
	File_proto_account_aggregates_account_grpc_base_create_proto = out.File
	file_proto_account_aggregates_account_grpc_base_create_proto_rawDesc = nil
	file_proto_account_aggregates_account_grpc_base_create_proto_goTypes = nil
	file_proto_account_aggregates_account_grpc_base_create_proto_depIdxs = nil
}
