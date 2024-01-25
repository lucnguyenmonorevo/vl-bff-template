// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/physical_obj/aggregates/transaction/transaction_begin.proto

package prototrx

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

type TransactionBeginUseCaseRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EpId    string `protobuf:"bytes,1,opt,name=ep_id,json=epId,proto3" json:"ep_id,omitempty"`
	EpTrxId string `protobuf:"bytes,2,opt,name=ep_trx_id,json=epTrxId,proto3" json:"ep_trx_id,omitempty"`
}

func (x *TransactionBeginUseCaseRequestData) Reset() {
	*x = TransactionBeginUseCaseRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionBeginUseCaseRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionBeginUseCaseRequestData) ProtoMessage() {}

func (x *TransactionBeginUseCaseRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionBeginUseCaseRequestData.ProtoReflect.Descriptor instead.
func (*TransactionBeginUseCaseRequestData) Descriptor() ([]byte, []int) {
	return file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionBeginUseCaseRequestData) GetEpId() string {
	if x != nil {
		return x.EpId
	}
	return ""
}

func (x *TransactionBeginUseCaseRequestData) GetEpTrxId() string {
	if x != nil {
		return x.EpTrxId
	}
	return ""
}

type TransactionBeginUseCaseResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionLog *TransactionLog `protobuf:"bytes,2,opt,name=transaction_log,json=transactionLog,proto3" json:"transaction_log,omitempty"`
}

func (x *TransactionBeginUseCaseResponseData) Reset() {
	*x = TransactionBeginUseCaseResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionBeginUseCaseResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionBeginUseCaseResponseData) ProtoMessage() {}

func (x *TransactionBeginUseCaseResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionBeginUseCaseResponseData.ProtoReflect.Descriptor instead.
func (*TransactionBeginUseCaseResponseData) Descriptor() ([]byte, []int) {
	return file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionBeginUseCaseResponseData) GetTransactionLog() *TransactionLog {
	if x != nil {
		return x.TransactionLog
	}
	return nil
}

type TransactionBeginUseCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction              `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Client      *messaging.Client                   `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Data        *TransactionBeginUseCaseRequestData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TransactionBeginUseCaseRequest) Reset() {
	*x = TransactionBeginUseCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionBeginUseCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionBeginUseCaseRequest) ProtoMessage() {}

func (x *TransactionBeginUseCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionBeginUseCaseRequest.ProtoReflect.Descriptor instead.
func (*TransactionBeginUseCaseRequest) Descriptor() ([]byte, []int) {
	return file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionBeginUseCaseRequest) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TransactionBeginUseCaseRequest) GetClient() *messaging.Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *TransactionBeginUseCaseRequest) GetData() *TransactionBeginUseCaseRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type TransactionBeginUseCaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction               `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Error       *messaging.Error                     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data        *TransactionBeginUseCaseResponseData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TransactionBeginUseCaseResponse) Reset() {
	*x = TransactionBeginUseCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionBeginUseCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionBeginUseCaseResponse) ProtoMessage() {}

func (x *TransactionBeginUseCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionBeginUseCaseResponse.ProtoReflect.Descriptor instead.
func (*TransactionBeginUseCaseResponse) Descriptor() ([]byte, []int) {
	return file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionBeginUseCaseResponse) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TransactionBeginUseCaseResponse) GetError() *messaging.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *TransactionBeginUseCaseResponse) GetData() *TransactionBeginUseCaseResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_physical_obj_aggregates_transaction_transaction_begin_proto protoreflect.FileDescriptor

var file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_rawDesc = []byte{
	0x0a, 0x41, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c,
	0x5f, 0x6f, 0x62, 0x6a, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x23, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x6f, 0x62,
	0x6a, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x3f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x6f, 0x62, 0x6a, 0x2f, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x55, 0x0a, 0x22, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x65, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x13, 0x0a, 0x05, 0x65, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x70, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x09,
	0x65, 0x70, 0x5f, 0x74, 0x72, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x70, 0x54, 0x72, 0x78, 0x49, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x23, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65,
	0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x5c, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x70, 0x68, 0x79, 0x73,
	0x69, 0x63, 0x61, 0x6c, 0x5f, 0x6f, 0x62, 0x6a, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x52, 0x0e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x22, 0xf0,
	0x01, 0x0a, 0x1e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x65,
	0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x5b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x47, 0x2e, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x6f, 0x62,
	0x6a, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0xef, 0x01, 0x0a, 0x1f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x5c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x6f,
	0x62, 0x6a, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x4a, 0x5a, 0x48, 0x76, 0x6c, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2f, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x6f, 0x62, 0x6a, 0x2f, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x72, 0x78, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_rawDescOnce sync.Once
	file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_rawDescData = file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_rawDesc
)

func file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_rawDescGZIP() []byte {
	file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_rawDescOnce.Do(func() {
		file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_rawDescData)
	})
	return file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_rawDescData
}

var file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_goTypes = []interface{}{
	(*TransactionBeginUseCaseRequestData)(nil),  // 0: physical_obj.aggregates.transaction.TransactionBeginUseCaseRequestData
	(*TransactionBeginUseCaseResponseData)(nil), // 1: physical_obj.aggregates.transaction.TransactionBeginUseCaseResponseData
	(*TransactionBeginUseCaseRequest)(nil),      // 2: physical_obj.aggregates.transaction.TransactionBeginUseCaseRequest
	(*TransactionBeginUseCaseResponse)(nil),     // 3: physical_obj.aggregates.transaction.TransactionBeginUseCaseResponse
	(*TransactionLog)(nil),                      // 4: physical_obj.aggregates.transaction.TransactionLog
	(*messaging.Transaction)(nil),               // 5: common.messaging.Transaction
	(*messaging.Client)(nil),                    // 6: common.messaging.Client
	(*messaging.Error)(nil),                     // 7: common.messaging.Error
}
var file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_depIdxs = []int32{
	4, // 0: physical_obj.aggregates.transaction.TransactionBeginUseCaseResponseData.transaction_log:type_name -> physical_obj.aggregates.transaction.TransactionLog
	5, // 1: physical_obj.aggregates.transaction.TransactionBeginUseCaseRequest.transaction:type_name -> common.messaging.Transaction
	6, // 2: physical_obj.aggregates.transaction.TransactionBeginUseCaseRequest.client:type_name -> common.messaging.Client
	0, // 3: physical_obj.aggregates.transaction.TransactionBeginUseCaseRequest.data:type_name -> physical_obj.aggregates.transaction.TransactionBeginUseCaseRequestData
	5, // 4: physical_obj.aggregates.transaction.TransactionBeginUseCaseResponse.transaction:type_name -> common.messaging.Transaction
	7, // 5: physical_obj.aggregates.transaction.TransactionBeginUseCaseResponse.error:type_name -> common.messaging.Error
	1, // 6: physical_obj.aggregates.transaction.TransactionBeginUseCaseResponse.data:type_name -> physical_obj.aggregates.transaction.TransactionBeginUseCaseResponseData
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_init() }
func file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_init() {
	if File_proto_physical_obj_aggregates_transaction_transaction_begin_proto != nil {
		return
	}
	file_proto_physical_obj_aggregates_transaction_transaction_log_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionBeginUseCaseRequestData); i {
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
		file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionBeginUseCaseResponseData); i {
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
		file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionBeginUseCaseRequest); i {
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
		file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionBeginUseCaseResponse); i {
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
			RawDescriptor: file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_goTypes,
		DependencyIndexes: file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_depIdxs,
		MessageInfos:      file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_msgTypes,
	}.Build()
	File_proto_physical_obj_aggregates_transaction_transaction_begin_proto = out.File
	file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_rawDesc = nil
	file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_goTypes = nil
	file_proto_physical_obj_aggregates_transaction_transaction_begin_proto_depIdxs = nil
}
