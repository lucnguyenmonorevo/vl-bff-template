// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/general_info/aggregates/pages/use_case/page_group_create.proto

package protopagesusecase

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	messaging "vl-account/proto_generated/common/messaging"
	general "vl-account/proto_generated/general_info/aggregates/pages/general"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PageGroupCreateUseCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction             `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Client      *messaging.Client                  `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Data        *PageGroupCreateUseCaseRequestData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PageGroupCreateUseCaseRequest) Reset() {
	*x = PageGroupCreateUseCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageGroupCreateUseCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageGroupCreateUseCaseRequest) ProtoMessage() {}

func (x *PageGroupCreateUseCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageGroupCreateUseCaseRequest.ProtoReflect.Descriptor instead.
func (*PageGroupCreateUseCaseRequest) Descriptor() ([]byte, []int) {
	return file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_rawDescGZIP(), []int{0}
}

func (x *PageGroupCreateUseCaseRequest) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *PageGroupCreateUseCaseRequest) GetClient() *messaging.Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *PageGroupCreateUseCaseRequest) GetData() *PageGroupCreateUseCaseRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type PageGroupCreateUseCaseRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string           `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Slug        string           `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Type        general.PageType `protobuf:"varint,4,opt,name=type,proto3,enum=general_info.aggregates.pages.general.PageType" json:"type,omitempty"`
	Sequence    uint32           `protobuf:"varint,5,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *PageGroupCreateUseCaseRequestData) Reset() {
	*x = PageGroupCreateUseCaseRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageGroupCreateUseCaseRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageGroupCreateUseCaseRequestData) ProtoMessage() {}

func (x *PageGroupCreateUseCaseRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageGroupCreateUseCaseRequestData.ProtoReflect.Descriptor instead.
func (*PageGroupCreateUseCaseRequestData) Descriptor() ([]byte, []int) {
	return file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_rawDescGZIP(), []int{1}
}

func (x *PageGroupCreateUseCaseRequestData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PageGroupCreateUseCaseRequestData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PageGroupCreateUseCaseRequestData) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *PageGroupCreateUseCaseRequestData) GetType() general.PageType {
	if x != nil {
		return x.Type
	}
	return general.PageType(0)
}

func (x *PageGroupCreateUseCaseRequestData) GetSequence() uint32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

type PageGroupCreateUseCaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Error       *messaging.Error       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data        *PageGroup             `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PageGroupCreateUseCaseResponse) Reset() {
	*x = PageGroupCreateUseCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageGroupCreateUseCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageGroupCreateUseCaseResponse) ProtoMessage() {}

func (x *PageGroupCreateUseCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageGroupCreateUseCaseResponse.ProtoReflect.Descriptor instead.
func (*PageGroupCreateUseCaseResponse) Descriptor() ([]byte, []int) {
	return file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_rawDescGZIP(), []int{2}
}

func (x *PageGroupCreateUseCaseResponse) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *PageGroupCreateUseCaseResponse) GetError() *messaging.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *PageGroupCreateUseCaseResponse) GetData() *PageGroup {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_general_info_aggregates_pages_use_case_page_group_create_proto protoreflect.FileDescriptor

var file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_rawDesc = []byte{
	0x0a, 0x44, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f,
	0x70, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e,
	0x70, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x1a, 0x3d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x61,
	0x67, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x01,
	0x0a, 0x1d, 0x50, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x30, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x5d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x49, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0xce, 0x01, 0x0a, 0x21, 0x50, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75,
	0x67, 0x12, 0x43, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2f, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x22, 0xd7, 0x01, 0x0a, 0x1e, 0x50, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x45, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x55, 0x5a, 0x53,
	0x76, 0x6c, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x73, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73,
	0x65, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x75, 0x73, 0x65, 0x63,
	0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_rawDescOnce sync.Once
	file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_rawDescData = file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_rawDesc
)

func file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_rawDescGZIP() []byte {
	file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_rawDescOnce.Do(func() {
		file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_rawDescData)
	})
	return file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_rawDescData
}

var file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_goTypes = []interface{}{
	(*PageGroupCreateUseCaseRequest)(nil),     // 0: general_info.aggregates.pages.use_case.PageGroupCreateUseCaseRequest
	(*PageGroupCreateUseCaseRequestData)(nil), // 1: general_info.aggregates.pages.use_case.PageGroupCreateUseCaseRequestData
	(*PageGroupCreateUseCaseResponse)(nil),    // 2: general_info.aggregates.pages.use_case.PageGroupCreateUseCaseResponse
	(*messaging.Transaction)(nil),             // 3: common.messaging.Transaction
	(*messaging.Client)(nil),                  // 4: common.messaging.Client
	(general.PageType)(0),                     // 5: general_info.aggregates.pages.general.PageType
	(*messaging.Error)(nil),                   // 6: common.messaging.Error
	(*PageGroup)(nil),                         // 7: general_info.aggregates.pages.use_case.PageGroup
}
var file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_depIdxs = []int32{
	3, // 0: general_info.aggregates.pages.use_case.PageGroupCreateUseCaseRequest.transaction:type_name -> common.messaging.Transaction
	4, // 1: general_info.aggregates.pages.use_case.PageGroupCreateUseCaseRequest.client:type_name -> common.messaging.Client
	1, // 2: general_info.aggregates.pages.use_case.PageGroupCreateUseCaseRequest.data:type_name -> general_info.aggregates.pages.use_case.PageGroupCreateUseCaseRequestData
	5, // 3: general_info.aggregates.pages.use_case.PageGroupCreateUseCaseRequestData.type:type_name -> general_info.aggregates.pages.general.PageType
	3, // 4: general_info.aggregates.pages.use_case.PageGroupCreateUseCaseResponse.transaction:type_name -> common.messaging.Transaction
	6, // 5: general_info.aggregates.pages.use_case.PageGroupCreateUseCaseResponse.error:type_name -> common.messaging.Error
	7, // 6: general_info.aggregates.pages.use_case.PageGroupCreateUseCaseResponse.data:type_name -> general_info.aggregates.pages.use_case.PageGroup
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_init() }
func file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_init() {
	if File_proto_general_info_aggregates_pages_use_case_page_group_create_proto != nil {
		return
	}
	file_proto_general_info_aggregates_pages_use_case_page_group_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageGroupCreateUseCaseRequest); i {
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
		file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageGroupCreateUseCaseRequestData); i {
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
		file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageGroupCreateUseCaseResponse); i {
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
			RawDescriptor: file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_goTypes,
		DependencyIndexes: file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_depIdxs,
		MessageInfos:      file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_msgTypes,
	}.Build()
	File_proto_general_info_aggregates_pages_use_case_page_group_create_proto = out.File
	file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_rawDesc = nil
	file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_goTypes = nil
	file_proto_general_info_aggregates_pages_use_case_page_group_create_proto_depIdxs = nil
}
