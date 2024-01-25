// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/production/aggregates/transaction/transaction_revert.proto

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

type TransactionRevertUseCaseRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EpId             string `protobuf:"bytes,1,opt,name=ep_id,json=epId,proto3" json:"ep_id,omitempty"`
	EpTrxId          string `protobuf:"bytes,2,opt,name=ep_trx_id,json=epTrxId,proto3" json:"ep_trx_id,omitempty"`
	TransactionLogId uint64 `protobuf:"varint,3,opt,name=transaction_log_id,json=transactionLogId,proto3" json:"transaction_log_id,omitempty"`
}

func (x *TransactionRevertUseCaseRequestData) Reset() {
	*x = TransactionRevertUseCaseRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_production_aggregates_transaction_transaction_revert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRevertUseCaseRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRevertUseCaseRequestData) ProtoMessage() {}

func (x *TransactionRevertUseCaseRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_production_aggregates_transaction_transaction_revert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRevertUseCaseRequestData.ProtoReflect.Descriptor instead.
func (*TransactionRevertUseCaseRequestData) Descriptor() ([]byte, []int) {
	return file_proto_production_aggregates_transaction_transaction_revert_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionRevertUseCaseRequestData) GetEpId() string {
	if x != nil {
		return x.EpId
	}
	return ""
}

func (x *TransactionRevertUseCaseRequestData) GetEpTrxId() string {
	if x != nil {
		return x.EpTrxId
	}
	return ""
}

func (x *TransactionRevertUseCaseRequestData) GetTransactionLogId() uint64 {
	if x != nil {
		return x.TransactionLogId
	}
	return 0
}

type TransactionRevertUseCaseResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionLog *TransactionLog `protobuf:"bytes,2,opt,name=transaction_log,json=transactionLog,proto3" json:"transaction_log,omitempty"`
}

func (x *TransactionRevertUseCaseResponseData) Reset() {
	*x = TransactionRevertUseCaseResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_production_aggregates_transaction_transaction_revert_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRevertUseCaseResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRevertUseCaseResponseData) ProtoMessage() {}

func (x *TransactionRevertUseCaseResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_production_aggregates_transaction_transaction_revert_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRevertUseCaseResponseData.ProtoReflect.Descriptor instead.
func (*TransactionRevertUseCaseResponseData) Descriptor() ([]byte, []int) {
	return file_proto_production_aggregates_transaction_transaction_revert_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionRevertUseCaseResponseData) GetTransactionLog() *TransactionLog {
	if x != nil {
		return x.TransactionLog
	}
	return nil
}

type TransactionRevertUseCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction               `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Client      *messaging.Client                    `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Data        *TransactionRevertUseCaseRequestData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TransactionRevertUseCaseRequest) Reset() {
	*x = TransactionRevertUseCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_production_aggregates_transaction_transaction_revert_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRevertUseCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRevertUseCaseRequest) ProtoMessage() {}

func (x *TransactionRevertUseCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_production_aggregates_transaction_transaction_revert_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRevertUseCaseRequest.ProtoReflect.Descriptor instead.
func (*TransactionRevertUseCaseRequest) Descriptor() ([]byte, []int) {
	return file_proto_production_aggregates_transaction_transaction_revert_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionRevertUseCaseRequest) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TransactionRevertUseCaseRequest) GetClient() *messaging.Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *TransactionRevertUseCaseRequest) GetData() *TransactionRevertUseCaseRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type TransactionRevertUseCaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction                `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Error       *messaging.Error                      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data        *TransactionRevertUseCaseResponseData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TransactionRevertUseCaseResponse) Reset() {
	*x = TransactionRevertUseCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_production_aggregates_transaction_transaction_revert_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRevertUseCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRevertUseCaseResponse) ProtoMessage() {}

func (x *TransactionRevertUseCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_production_aggregates_transaction_transaction_revert_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRevertUseCaseResponse.ProtoReflect.Descriptor instead.
func (*TransactionRevertUseCaseResponse) Descriptor() ([]byte, []int) {
	return file_proto_production_aggregates_transaction_transaction_revert_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionRevertUseCaseResponse) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TransactionRevertUseCaseResponse) GetError() *messaging.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *TransactionRevertUseCaseResponse) GetData() *TransactionRevertUseCaseResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_production_aggregates_transaction_transaction_revert_proto protoreflect.FileDescriptor

var file_proto_production_aggregates_transaction_transaction_revert_proto_rawDesc = []byte{
	0x0a, 0x40, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x21, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x3d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a,
	0x23, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x65,
	0x72, 0x74, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x13, 0x0a, 0x05, 0x65, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x70, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x09, 0x65, 0x70, 0x5f,
	0x74, 0x72, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x70,
	0x54, 0x72, 0x78, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f,
	0x67, 0x49, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x24, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x5a, 0x0a, 0x0f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x22, 0xf0, 0x01, 0x0a, 0x1f, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x55, 0x73,
	0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a,
	0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x5a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76,
	0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xef, 0x01, 0x0a, 0x20,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x65, 0x72,
	0x74, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x5b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x76, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x48, 0x5a,
	0x46, 0x76, 0x6c, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x74, 0x72, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_production_aggregates_transaction_transaction_revert_proto_rawDescOnce sync.Once
	file_proto_production_aggregates_transaction_transaction_revert_proto_rawDescData = file_proto_production_aggregates_transaction_transaction_revert_proto_rawDesc
)

func file_proto_production_aggregates_transaction_transaction_revert_proto_rawDescGZIP() []byte {
	file_proto_production_aggregates_transaction_transaction_revert_proto_rawDescOnce.Do(func() {
		file_proto_production_aggregates_transaction_transaction_revert_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_production_aggregates_transaction_transaction_revert_proto_rawDescData)
	})
	return file_proto_production_aggregates_transaction_transaction_revert_proto_rawDescData
}

var file_proto_production_aggregates_transaction_transaction_revert_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_production_aggregates_transaction_transaction_revert_proto_goTypes = []interface{}{
	(*TransactionRevertUseCaseRequestData)(nil),  // 0: production.aggregates.transaction.TransactionRevertUseCaseRequestData
	(*TransactionRevertUseCaseResponseData)(nil), // 1: production.aggregates.transaction.TransactionRevertUseCaseResponseData
	(*TransactionRevertUseCaseRequest)(nil),      // 2: production.aggregates.transaction.TransactionRevertUseCaseRequest
	(*TransactionRevertUseCaseResponse)(nil),     // 3: production.aggregates.transaction.TransactionRevertUseCaseResponse
	(*TransactionLog)(nil),                       // 4: production.aggregates.transaction.TransactionLog
	(*messaging.Transaction)(nil),                // 5: common.messaging.Transaction
	(*messaging.Client)(nil),                     // 6: common.messaging.Client
	(*messaging.Error)(nil),                      // 7: common.messaging.Error
}
var file_proto_production_aggregates_transaction_transaction_revert_proto_depIdxs = []int32{
	4, // 0: production.aggregates.transaction.TransactionRevertUseCaseResponseData.transaction_log:type_name -> production.aggregates.transaction.TransactionLog
	5, // 1: production.aggregates.transaction.TransactionRevertUseCaseRequest.transaction:type_name -> common.messaging.Transaction
	6, // 2: production.aggregates.transaction.TransactionRevertUseCaseRequest.client:type_name -> common.messaging.Client
	0, // 3: production.aggregates.transaction.TransactionRevertUseCaseRequest.data:type_name -> production.aggregates.transaction.TransactionRevertUseCaseRequestData
	5, // 4: production.aggregates.transaction.TransactionRevertUseCaseResponse.transaction:type_name -> common.messaging.Transaction
	7, // 5: production.aggregates.transaction.TransactionRevertUseCaseResponse.error:type_name -> common.messaging.Error
	1, // 6: production.aggregates.transaction.TransactionRevertUseCaseResponse.data:type_name -> production.aggregates.transaction.TransactionRevertUseCaseResponseData
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_production_aggregates_transaction_transaction_revert_proto_init() }
func file_proto_production_aggregates_transaction_transaction_revert_proto_init() {
	if File_proto_production_aggregates_transaction_transaction_revert_proto != nil {
		return
	}
	file_proto_production_aggregates_transaction_transaction_log_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_production_aggregates_transaction_transaction_revert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRevertUseCaseRequestData); i {
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
		file_proto_production_aggregates_transaction_transaction_revert_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRevertUseCaseResponseData); i {
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
		file_proto_production_aggregates_transaction_transaction_revert_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRevertUseCaseRequest); i {
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
		file_proto_production_aggregates_transaction_transaction_revert_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRevertUseCaseResponse); i {
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
			RawDescriptor: file_proto_production_aggregates_transaction_transaction_revert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_production_aggregates_transaction_transaction_revert_proto_goTypes,
		DependencyIndexes: file_proto_production_aggregates_transaction_transaction_revert_proto_depIdxs,
		MessageInfos:      file_proto_production_aggregates_transaction_transaction_revert_proto_msgTypes,
	}.Build()
	File_proto_production_aggregates_transaction_transaction_revert_proto = out.File
	file_proto_production_aggregates_transaction_transaction_revert_proto_rawDesc = nil
	file_proto_production_aggregates_transaction_transaction_revert_proto_goTypes = nil
	file_proto_production_aggregates_transaction_transaction_revert_proto_depIdxs = nil
}
