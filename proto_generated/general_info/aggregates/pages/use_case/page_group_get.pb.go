// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/general_info/aggregates/pages/use_case/page_group_get.proto

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

type PageGroupGetUseCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction          `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Client      *messaging.Client               `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Data        *PageGroupGetUseCaseRequestData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PageGroupGetUseCaseRequest) Reset() {
	*x = PageGroupGetUseCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageGroupGetUseCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageGroupGetUseCaseRequest) ProtoMessage() {}

func (x *PageGroupGetUseCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageGroupGetUseCaseRequest.ProtoReflect.Descriptor instead.
func (*PageGroupGetUseCaseRequest) Descriptor() ([]byte, []int) {
	return file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_rawDescGZIP(), []int{0}
}

func (x *PageGroupGetUseCaseRequest) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *PageGroupGetUseCaseRequest) GetClient() *messaging.Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *PageGroupGetUseCaseRequest) GetData() *PageGroupGetUseCaseRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type PageGroupGetUseCaseRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PageGroupGetUseCaseRequestData) Reset() {
	*x = PageGroupGetUseCaseRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageGroupGetUseCaseRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageGroupGetUseCaseRequestData) ProtoMessage() {}

func (x *PageGroupGetUseCaseRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageGroupGetUseCaseRequestData.ProtoReflect.Descriptor instead.
func (*PageGroupGetUseCaseRequestData) Descriptor() ([]byte, []int) {
	return file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_rawDescGZIP(), []int{1}
}

func (x *PageGroupGetUseCaseRequestData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PageGroupGetUseCaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction           `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Error       *messaging.Error                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data        *PageGroupGetUseCaseResponseData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PageGroupGetUseCaseResponse) Reset() {
	*x = PageGroupGetUseCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageGroupGetUseCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageGroupGetUseCaseResponse) ProtoMessage() {}

func (x *PageGroupGetUseCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageGroupGetUseCaseResponse.ProtoReflect.Descriptor instead.
func (*PageGroupGetUseCaseResponse) Descriptor() ([]byte, []int) {
	return file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_rawDescGZIP(), []int{2}
}

func (x *PageGroupGetUseCaseResponse) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *PageGroupGetUseCaseResponse) GetError() *messaging.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *PageGroupGetUseCaseResponse) GetData() *PageGroupGetUseCaseResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type PageGroupGetUseCaseResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageGroup *PageGroup                    `protobuf:"bytes,1,opt,name=page_group,json=pageGroup,proto3" json:"page_group,omitempty"`
	Pages     []*PageGetUseCaseResponseData `protobuf:"bytes,2,rep,name=pages,proto3" json:"pages,omitempty"`
}

func (x *PageGroupGetUseCaseResponseData) Reset() {
	*x = PageGroupGetUseCaseResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageGroupGetUseCaseResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageGroupGetUseCaseResponseData) ProtoMessage() {}

func (x *PageGroupGetUseCaseResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageGroupGetUseCaseResponseData.ProtoReflect.Descriptor instead.
func (*PageGroupGetUseCaseResponseData) Descriptor() ([]byte, []int) {
	return file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_rawDescGZIP(), []int{3}
}

func (x *PageGroupGetUseCaseResponseData) GetPageGroup() *PageGroup {
	if x != nil {
		return x.PageGroup
	}
	return nil
}

func (x *PageGroupGetUseCaseResponseData) GetPages() []*PageGetUseCaseResponseData {
	if x != nil {
		return x.Pages
	}
	return nil
}

var File_proto_general_info_aggregates_pages_use_case_page_group_get_proto protoreflect.FileDescriptor

var file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_rawDesc = []byte{
	0x0a, 0x41, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f,
	0x70, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x26, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x1a, 0x3d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xeb, 0x01, 0x0a, 0x1a, 0x50, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x65, 0x74,
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
	0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x30, 0x0a,
	0x1e, 0x50, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22,
	0xea, 0x01, 0x0a, 0x1b, 0x50, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x5b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x75, 0x73,
	0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xcd, 0x01, 0x0a,
	0x1f, 0x50, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x50, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x58, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x42, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x42, 0x55, 0x5a, 0x53,
	0x76, 0x6c, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x73, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73,
	0x65, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x75, 0x73, 0x65, 0x63,
	0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_rawDescOnce sync.Once
	file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_rawDescData = file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_rawDesc
)

func file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_rawDescGZIP() []byte {
	file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_rawDescOnce.Do(func() {
		file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_rawDescData)
	})
	return file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_rawDescData
}

var file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_goTypes = []interface{}{
	(*PageGroupGetUseCaseRequest)(nil),      // 0: general_info.aggregates.pages.use_case.PageGroupGetUseCaseRequest
	(*PageGroupGetUseCaseRequestData)(nil),  // 1: general_info.aggregates.pages.use_case.PageGroupGetUseCaseRequestData
	(*PageGroupGetUseCaseResponse)(nil),     // 2: general_info.aggregates.pages.use_case.PageGroupGetUseCaseResponse
	(*PageGroupGetUseCaseResponseData)(nil), // 3: general_info.aggregates.pages.use_case.PageGroupGetUseCaseResponseData
	(*messaging.Transaction)(nil),           // 4: common.messaging.Transaction
	(*messaging.Client)(nil),                // 5: common.messaging.Client
	(*messaging.Error)(nil),                 // 6: common.messaging.Error
	(*PageGroup)(nil),                       // 7: general_info.aggregates.pages.use_case.PageGroup
	(*PageGetUseCaseResponseData)(nil),      // 8: general_info.aggregates.pages.use_case.PageGetUseCaseResponseData
}
var file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_depIdxs = []int32{
	4, // 0: general_info.aggregates.pages.use_case.PageGroupGetUseCaseRequest.transaction:type_name -> common.messaging.Transaction
	5, // 1: general_info.aggregates.pages.use_case.PageGroupGetUseCaseRequest.client:type_name -> common.messaging.Client
	1, // 2: general_info.aggregates.pages.use_case.PageGroupGetUseCaseRequest.data:type_name -> general_info.aggregates.pages.use_case.PageGroupGetUseCaseRequestData
	4, // 3: general_info.aggregates.pages.use_case.PageGroupGetUseCaseResponse.transaction:type_name -> common.messaging.Transaction
	6, // 4: general_info.aggregates.pages.use_case.PageGroupGetUseCaseResponse.error:type_name -> common.messaging.Error
	3, // 5: general_info.aggregates.pages.use_case.PageGroupGetUseCaseResponse.data:type_name -> general_info.aggregates.pages.use_case.PageGroupGetUseCaseResponseData
	7, // 6: general_info.aggregates.pages.use_case.PageGroupGetUseCaseResponseData.page_group:type_name -> general_info.aggregates.pages.use_case.PageGroup
	8, // 7: general_info.aggregates.pages.use_case.PageGroupGetUseCaseResponseData.pages:type_name -> general_info.aggregates.pages.use_case.PageGetUseCaseResponseData
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_init() }
func file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_init() {
	if File_proto_general_info_aggregates_pages_use_case_page_group_get_proto != nil {
		return
	}
	file_proto_general_info_aggregates_pages_use_case_page_group_proto_init()
	file_proto_general_info_aggregates_pages_use_case_page_get_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageGroupGetUseCaseRequest); i {
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
		file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageGroupGetUseCaseRequestData); i {
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
		file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageGroupGetUseCaseResponse); i {
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
		file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageGroupGetUseCaseResponseData); i {
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
			RawDescriptor: file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_goTypes,
		DependencyIndexes: file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_depIdxs,
		MessageInfos:      file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_msgTypes,
	}.Build()
	File_proto_general_info_aggregates_pages_use_case_page_group_get_proto = out.File
	file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_rawDesc = nil
	file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_goTypes = nil
	file_proto_general_info_aggregates_pages_use_case_page_group_get_proto_depIdxs = nil
}