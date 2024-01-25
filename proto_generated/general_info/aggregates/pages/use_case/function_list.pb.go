// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/general_info/aggregates/pages/use_case/function_list.proto

package protopagesusecase

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	messaging "vl-account/proto_generated/common/messaging"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FunctionListUseCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction          `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Client      *messaging.Client               `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Data        *FunctionListUseCaseRequestData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FunctionListUseCaseRequest) Reset() {
	*x = FunctionListUseCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_general_info_aggregates_pages_use_case_function_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionListUseCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionListUseCaseRequest) ProtoMessage() {}

func (x *FunctionListUseCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_general_info_aggregates_pages_use_case_function_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionListUseCaseRequest.ProtoReflect.Descriptor instead.
func (*FunctionListUseCaseRequest) Descriptor() ([]byte, []int) {
	return file_proto_general_info_aggregates_pages_use_case_function_list_proto_rawDescGZIP(), []int{0}
}

func (x *FunctionListUseCaseRequest) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *FunctionListUseCaseRequest) GetClient() *messaging.Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *FunctionListUseCaseRequest) GetData() *FunctionListUseCaseRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type FunctionListUseCaseRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit      uint64   `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset     uint64   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	SearchText string   `protobuf:"bytes,3,opt,name=search_text,json=searchText,proto3" json:"search_text,omitempty"`
	PageId     uint64   `protobuf:"varint,4,opt,name=page_id,json=pageId,proto3" json:"page_id,omitempty"`
	Ids        []uint64 `protobuf:"varint,5,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *FunctionListUseCaseRequestData) Reset() {
	*x = FunctionListUseCaseRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_general_info_aggregates_pages_use_case_function_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionListUseCaseRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionListUseCaseRequestData) ProtoMessage() {}

func (x *FunctionListUseCaseRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_general_info_aggregates_pages_use_case_function_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionListUseCaseRequestData.ProtoReflect.Descriptor instead.
func (*FunctionListUseCaseRequestData) Descriptor() ([]byte, []int) {
	return file_proto_general_info_aggregates_pages_use_case_function_list_proto_rawDescGZIP(), []int{1}
}

func (x *FunctionListUseCaseRequestData) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FunctionListUseCaseRequestData) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *FunctionListUseCaseRequestData) GetSearchText() string {
	if x != nil {
		return x.SearchText
	}
	return ""
}

func (x *FunctionListUseCaseRequestData) GetPageId() uint64 {
	if x != nil {
		return x.PageId
	}
	return 0
}

func (x *FunctionListUseCaseRequestData) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type FunctionListUseCaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction           `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Error       *messaging.Error                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data        *FunctionListUseCaseResponseData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FunctionListUseCaseResponse) Reset() {
	*x = FunctionListUseCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_general_info_aggregates_pages_use_case_function_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionListUseCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionListUseCaseResponse) ProtoMessage() {}

func (x *FunctionListUseCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_general_info_aggregates_pages_use_case_function_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionListUseCaseResponse.ProtoReflect.Descriptor instead.
func (*FunctionListUseCaseResponse) Descriptor() ([]byte, []int) {
	return file_proto_general_info_aggregates_pages_use_case_function_list_proto_rawDescGZIP(), []int{2}
}

func (x *FunctionListUseCaseResponse) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *FunctionListUseCaseResponse) GetError() *messaging.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *FunctionListUseCaseResponse) GetData() *FunctionListUseCaseResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type FunctionListUseCaseResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Functions []*Function `protobuf:"bytes,1,rep,name=functions,proto3" json:"functions,omitempty"`
}

func (x *FunctionListUseCaseResponseData) Reset() {
	*x = FunctionListUseCaseResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_general_info_aggregates_pages_use_case_function_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionListUseCaseResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionListUseCaseResponseData) ProtoMessage() {}

func (x *FunctionListUseCaseResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_general_info_aggregates_pages_use_case_function_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionListUseCaseResponseData.ProtoReflect.Descriptor instead.
func (*FunctionListUseCaseResponseData) Descriptor() ([]byte, []int) {
	return file_proto_general_info_aggregates_pages_use_case_function_list_proto_rawDescGZIP(), []int{3}
}

func (x *FunctionListUseCaseResponseData) GetFunctions() []*Function {
	if x != nil {
		return x.Functions
	}
	return nil
}

var File_proto_general_info_aggregates_pages_use_case_function_list_proto protoreflect.FileDescriptor

var file_proto_general_info_aggregates_pages_use_case_function_list_proto_rawDesc = []byte{
	0x0a, 0x40, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f,
	0x70, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x26, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x1a, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xeb, 0x01, 0x0a, 0x1a, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f,
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x30, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x5a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x46, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9a, 0x01,
	0x0a, 0x1e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x78, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x04, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xea, 0x01, 0x0a, 0x1b, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x43, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x5b, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x73, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73,
	0x65, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x71, 0x0a, 0x1f, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x4e, 0x0a, 0x09, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x75, 0x73,
	0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x55, 0x5a, 0x53, 0x76, 0x6c,
	0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73,
	0x2f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x75, 0x73, 0x65, 0x63, 0x61, 0x73,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_general_info_aggregates_pages_use_case_function_list_proto_rawDescOnce sync.Once
	file_proto_general_info_aggregates_pages_use_case_function_list_proto_rawDescData = file_proto_general_info_aggregates_pages_use_case_function_list_proto_rawDesc
)

func file_proto_general_info_aggregates_pages_use_case_function_list_proto_rawDescGZIP() []byte {
	file_proto_general_info_aggregates_pages_use_case_function_list_proto_rawDescOnce.Do(func() {
		file_proto_general_info_aggregates_pages_use_case_function_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_general_info_aggregates_pages_use_case_function_list_proto_rawDescData)
	})
	return file_proto_general_info_aggregates_pages_use_case_function_list_proto_rawDescData
}

var file_proto_general_info_aggregates_pages_use_case_function_list_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_general_info_aggregates_pages_use_case_function_list_proto_goTypes = []interface{}{
	(*FunctionListUseCaseRequest)(nil),      // 0: general_info.aggregates.pages.use_case.FunctionListUseCaseRequest
	(*FunctionListUseCaseRequestData)(nil),  // 1: general_info.aggregates.pages.use_case.FunctionListUseCaseRequestData
	(*FunctionListUseCaseResponse)(nil),     // 2: general_info.aggregates.pages.use_case.FunctionListUseCaseResponse
	(*FunctionListUseCaseResponseData)(nil), // 3: general_info.aggregates.pages.use_case.FunctionListUseCaseResponseData
	(*messaging.Transaction)(nil),           // 4: common.messaging.Transaction
	(*messaging.Client)(nil),                // 5: common.messaging.Client
	(*messaging.Error)(nil),                 // 6: common.messaging.Error
	(*Function)(nil),                        // 7: general_info.aggregates.pages.use_case.Function
}
var file_proto_general_info_aggregates_pages_use_case_function_list_proto_depIdxs = []int32{
	4, // 0: general_info.aggregates.pages.use_case.FunctionListUseCaseRequest.transaction:type_name -> common.messaging.Transaction
	5, // 1: general_info.aggregates.pages.use_case.FunctionListUseCaseRequest.client:type_name -> common.messaging.Client
	1, // 2: general_info.aggregates.pages.use_case.FunctionListUseCaseRequest.data:type_name -> general_info.aggregates.pages.use_case.FunctionListUseCaseRequestData
	4, // 3: general_info.aggregates.pages.use_case.FunctionListUseCaseResponse.transaction:type_name -> common.messaging.Transaction
	6, // 4: general_info.aggregates.pages.use_case.FunctionListUseCaseResponse.error:type_name -> common.messaging.Error
	3, // 5: general_info.aggregates.pages.use_case.FunctionListUseCaseResponse.data:type_name -> general_info.aggregates.pages.use_case.FunctionListUseCaseResponseData
	7, // 6: general_info.aggregates.pages.use_case.FunctionListUseCaseResponseData.functions:type_name -> general_info.aggregates.pages.use_case.Function
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_general_info_aggregates_pages_use_case_function_list_proto_init() }
func file_proto_general_info_aggregates_pages_use_case_function_list_proto_init() {
	if File_proto_general_info_aggregates_pages_use_case_function_list_proto != nil {
		return
	}
	file_proto_general_info_aggregates_pages_use_case_function_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_general_info_aggregates_pages_use_case_function_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionListUseCaseRequest); i {
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
		file_proto_general_info_aggregates_pages_use_case_function_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionListUseCaseRequestData); i {
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
		file_proto_general_info_aggregates_pages_use_case_function_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionListUseCaseResponse); i {
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
		file_proto_general_info_aggregates_pages_use_case_function_list_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionListUseCaseResponseData); i {
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
			RawDescriptor: file_proto_general_info_aggregates_pages_use_case_function_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_general_info_aggregates_pages_use_case_function_list_proto_goTypes,
		DependencyIndexes: file_proto_general_info_aggregates_pages_use_case_function_list_proto_depIdxs,
		MessageInfos:      file_proto_general_info_aggregates_pages_use_case_function_list_proto_msgTypes,
	}.Build()
	File_proto_general_info_aggregates_pages_use_case_function_list_proto = out.File
	file_proto_general_info_aggregates_pages_use_case_function_list_proto_rawDesc = nil
	file_proto_general_info_aggregates_pages_use_case_function_list_proto_goTypes = nil
	file_proto_general_info_aggregates_pages_use_case_function_list_proto_depIdxs = nil
}
