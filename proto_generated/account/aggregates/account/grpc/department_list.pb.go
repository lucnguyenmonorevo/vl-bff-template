// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/account/aggregates/account/grpc/department_list.proto

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
type DepartmentListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCRequest          `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *DepartmentListRequestData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DepartmentListRequest) Reset() {
	*x = DepartmentListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_department_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepartmentListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepartmentListRequest) ProtoMessage() {}

func (x *DepartmentListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_department_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepartmentListRequest.ProtoReflect.Descriptor instead.
func (*DepartmentListRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_department_list_proto_rawDescGZIP(), []int{0}
}

func (x *DepartmentListRequest) GetBase() *grpc.GRPCRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *DepartmentListRequest) GetData() *DepartmentListRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type DepartmentListRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationId uint64   `protobuf:"varint,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	SearchText     string   `protobuf:"bytes,2,opt,name=searchText,proto3" json:"searchText,omitempty"`
	Offset         uint32   `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit          uint32   `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	BaseId         uint64   `protobuf:"varint,5,opt,name=base_id,json=baseId,proto3" json:"base_id,omitempty"`
	Ids            []uint64 `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DepartmentListRequestData) Reset() {
	*x = DepartmentListRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_department_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepartmentListRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepartmentListRequestData) ProtoMessage() {}

func (x *DepartmentListRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_department_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepartmentListRequestData.ProtoReflect.Descriptor instead.
func (*DepartmentListRequestData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_department_list_proto_rawDescGZIP(), []int{1}
}

func (x *DepartmentListRequestData) GetOrganizationId() uint64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *DepartmentListRequestData) GetSearchText() string {
	if x != nil {
		return x.SearchText
	}
	return ""
}

func (x *DepartmentListRequestData) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *DepartmentListRequestData) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *DepartmentListRequestData) GetBaseId() uint64 {
	if x != nil {
		return x.BaseId
	}
	return 0
}

func (x *DepartmentListRequestData) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

// Response
type DepartmentListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCResponse          `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *DepartmentListResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DepartmentListResponse) Reset() {
	*x = DepartmentListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_department_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepartmentListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepartmentListResponse) ProtoMessage() {}

func (x *DepartmentListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_department_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepartmentListResponse.ProtoReflect.Descriptor instead.
func (*DepartmentListResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_department_list_proto_rawDescGZIP(), []int{2}
}

func (x *DepartmentListResponse) GetBase() *grpc.GRPCResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *DepartmentListResponse) GetData() *DepartmentListResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type DepartmentListResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Departments []*DepartmentGetResponseData `protobuf:"bytes,1,rep,name=departments,proto3" json:"departments,omitempty"`
}

func (x *DepartmentListResponseData) Reset() {
	*x = DepartmentListResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_department_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepartmentListResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepartmentListResponseData) ProtoMessage() {}

func (x *DepartmentListResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_department_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepartmentListResponseData.ProtoReflect.Descriptor instead.
func (*DepartmentListResponseData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_department_list_proto_rawDescGZIP(), []int{3}
}

func (x *DepartmentListResponseData) GetDepartments() []*DepartmentGetResponseData {
	if x != nil {
		return x.Departments
	}
	return nil
}

var File_proto_account_aggregates_account_grpc_department_list_proto protoreflect.FileDescriptor

var file_proto_account_aggregates_account_grpc_department_list_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x3a,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x4e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0xbd, 0x01, 0x0a, 0x19, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27,
	0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x54, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x62, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x04, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x22, 0x98, 0x01, 0x0a, 0x16, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7a, 0x0a, 0x1a, 0x44,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x5c, 0x0a, 0x0b, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x4d, 0x5a, 0x4b, 0x76, 0x6c, 0x2d, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_aggregates_account_grpc_department_list_proto_rawDescOnce sync.Once
	file_proto_account_aggregates_account_grpc_department_list_proto_rawDescData = file_proto_account_aggregates_account_grpc_department_list_proto_rawDesc
)

func file_proto_account_aggregates_account_grpc_department_list_proto_rawDescGZIP() []byte {
	file_proto_account_aggregates_account_grpc_department_list_proto_rawDescOnce.Do(func() {
		file_proto_account_aggregates_account_grpc_department_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_aggregates_account_grpc_department_list_proto_rawDescData)
	})
	return file_proto_account_aggregates_account_grpc_department_list_proto_rawDescData
}

var file_proto_account_aggregates_account_grpc_department_list_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_account_aggregates_account_grpc_department_list_proto_goTypes = []interface{}{
	(*DepartmentListRequest)(nil),      // 0: account.aggregates.account.grpc.DepartmentListRequest
	(*DepartmentListRequestData)(nil),  // 1: account.aggregates.account.grpc.DepartmentListRequestData
	(*DepartmentListResponse)(nil),     // 2: account.aggregates.account.grpc.DepartmentListResponse
	(*DepartmentListResponseData)(nil), // 3: account.aggregates.account.grpc.DepartmentListResponseData
	(*grpc.GRPCRequest)(nil),           // 4: common.grpc.GRPCRequest
	(*grpc.GRPCResponse)(nil),          // 5: common.grpc.GRPCResponse
	(*DepartmentGetResponseData)(nil),  // 6: account.aggregates.account.grpc.DepartmentGetResponseData
}
var file_proto_account_aggregates_account_grpc_department_list_proto_depIdxs = []int32{
	4, // 0: account.aggregates.account.grpc.DepartmentListRequest.base:type_name -> common.grpc.GRPCRequest
	1, // 1: account.aggregates.account.grpc.DepartmentListRequest.data:type_name -> account.aggregates.account.grpc.DepartmentListRequestData
	5, // 2: account.aggregates.account.grpc.DepartmentListResponse.base:type_name -> common.grpc.GRPCResponse
	3, // 3: account.aggregates.account.grpc.DepartmentListResponse.data:type_name -> account.aggregates.account.grpc.DepartmentListResponseData
	6, // 4: account.aggregates.account.grpc.DepartmentListResponseData.departments:type_name -> account.aggregates.account.grpc.DepartmentGetResponseData
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_account_aggregates_account_grpc_department_list_proto_init() }
func file_proto_account_aggregates_account_grpc_department_list_proto_init() {
	if File_proto_account_aggregates_account_grpc_department_list_proto != nil {
		return
	}
	file_proto_account_aggregates_account_grpc_department_get_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_account_aggregates_account_grpc_department_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepartmentListRequest); i {
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
		file_proto_account_aggregates_account_grpc_department_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepartmentListRequestData); i {
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
		file_proto_account_aggregates_account_grpc_department_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepartmentListResponse); i {
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
		file_proto_account_aggregates_account_grpc_department_list_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepartmentListResponseData); i {
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
			RawDescriptor: file_proto_account_aggregates_account_grpc_department_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_aggregates_account_grpc_department_list_proto_goTypes,
		DependencyIndexes: file_proto_account_aggregates_account_grpc_department_list_proto_depIdxs,
		MessageInfos:      file_proto_account_aggregates_account_grpc_department_list_proto_msgTypes,
	}.Build()
	File_proto_account_aggregates_account_grpc_department_list_proto = out.File
	file_proto_account_aggregates_account_grpc_department_list_proto_rawDesc = nil
	file_proto_account_aggregates_account_grpc_department_list_proto_goTypes = nil
	file_proto_account_aggregates_account_grpc_department_list_proto_depIdxs = nil
}