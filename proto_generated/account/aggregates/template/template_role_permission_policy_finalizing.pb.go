// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/account/aggregates/template/template_role_permission_policy_finalizing.proto

package prototemplate

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

type TemplateRolePermissionPolicyFinalizingCQRSRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TemplateRolePermissionPolicyFinalizingCQRSRequestData) Reset() {
	*x = TemplateRolePermissionPolicyFinalizingCQRSRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRolePermissionPolicyFinalizingCQRSRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRolePermissionPolicyFinalizingCQRSRequestData) ProtoMessage() {}

func (x *TemplateRolePermissionPolicyFinalizingCQRSRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRolePermissionPolicyFinalizingCQRSRequestData.ProtoReflect.Descriptor instead.
func (*TemplateRolePermissionPolicyFinalizingCQRSRequestData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_rawDescGZIP(), []int{0}
}

func (x *TemplateRolePermissionPolicyFinalizingCQRSRequestData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TemplateRolePermissionPolicyFinalizingCQRSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction                                 `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Client      *messaging.Client                                      `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Data        *TemplateRolePermissionPolicyFinalizingCQRSRequestData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TemplateRolePermissionPolicyFinalizingCQRSRequest) Reset() {
	*x = TemplateRolePermissionPolicyFinalizingCQRSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRolePermissionPolicyFinalizingCQRSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRolePermissionPolicyFinalizingCQRSRequest) ProtoMessage() {}

func (x *TemplateRolePermissionPolicyFinalizingCQRSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRolePermissionPolicyFinalizingCQRSRequest.ProtoReflect.Descriptor instead.
func (*TemplateRolePermissionPolicyFinalizingCQRSRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_rawDescGZIP(), []int{1}
}

func (x *TemplateRolePermissionPolicyFinalizingCQRSRequest) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TemplateRolePermissionPolicyFinalizingCQRSRequest) GetClient() *messaging.Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *TemplateRolePermissionPolicyFinalizingCQRSRequest) GetData() *TemplateRolePermissionPolicyFinalizingCQRSRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type TemplateRolePermissionPolicyFinalizingCQRSResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction        `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Error       *messaging.Error              `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data        *TemplateRolePermissionPolicy `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TemplateRolePermissionPolicyFinalizingCQRSResponse) Reset() {
	*x = TemplateRolePermissionPolicyFinalizingCQRSResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRolePermissionPolicyFinalizingCQRSResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRolePermissionPolicyFinalizingCQRSResponse) ProtoMessage() {}

func (x *TemplateRolePermissionPolicyFinalizingCQRSResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRolePermissionPolicyFinalizingCQRSResponse.ProtoReflect.Descriptor instead.
func (*TemplateRolePermissionPolicyFinalizingCQRSResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_rawDescGZIP(), []int{2}
}

func (x *TemplateRolePermissionPolicyFinalizingCQRSResponse) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TemplateRolePermissionPolicyFinalizingCQRSResponse) GetError() *messaging.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *TemplateRolePermissionPolicyFinalizingCQRSResponse) GetData() *TemplateRolePermissionPolicy {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto protoreflect.FileDescriptor

var file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_rawDesc = []byte{
	0x0a, 0x52, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6c,
	0x65, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x1a, 0x47, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f,
	0x6c, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x47, 0x0a, 0x35, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x43, 0x51, 0x52, 0x53,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8e, 0x02, 0x0a, 0x31,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x46, 0x69, 0x6e, 0x61,
	0x6c, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x43, 0x51, 0x52, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x66, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x52, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x46, 0x69, 0x6e,
	0x61, 0x6c, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x43, 0x51, 0x52, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xf3, 0x01, 0x0a,
	0x32, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x46, 0x69, 0x6e,
	0x61, 0x6c, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x43, 0x51, 0x52, 0x53, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x4d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x39, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x47, 0x5a, 0x45, 0x76, 0x6c, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_rawDescOnce sync.Once
	file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_rawDescData = file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_rawDesc
)

func file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_rawDescGZIP() []byte {
	file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_rawDescOnce.Do(func() {
		file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_rawDescData)
	})
	return file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_rawDescData
}

var file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_goTypes = []interface{}{
	(*TemplateRolePermissionPolicyFinalizingCQRSRequestData)(nil), // 0: account.aggregates.template.TemplateRolePermissionPolicyFinalizingCQRSRequestData
	(*TemplateRolePermissionPolicyFinalizingCQRSRequest)(nil),     // 1: account.aggregates.template.TemplateRolePermissionPolicyFinalizingCQRSRequest
	(*TemplateRolePermissionPolicyFinalizingCQRSResponse)(nil),    // 2: account.aggregates.template.TemplateRolePermissionPolicyFinalizingCQRSResponse
	(*messaging.Transaction)(nil),                                 // 3: common.messaging.Transaction
	(*messaging.Client)(nil),                                      // 4: common.messaging.Client
	(*messaging.Error)(nil),                                       // 5: common.messaging.Error
	(*TemplateRolePermissionPolicy)(nil),                          // 6: account.aggregates.template.TemplateRolePermissionPolicy
}
var file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_depIdxs = []int32{
	3, // 0: account.aggregates.template.TemplateRolePermissionPolicyFinalizingCQRSRequest.transaction:type_name -> common.messaging.Transaction
	4, // 1: account.aggregates.template.TemplateRolePermissionPolicyFinalizingCQRSRequest.client:type_name -> common.messaging.Client
	0, // 2: account.aggregates.template.TemplateRolePermissionPolicyFinalizingCQRSRequest.data:type_name -> account.aggregates.template.TemplateRolePermissionPolicyFinalizingCQRSRequestData
	3, // 3: account.aggregates.template.TemplateRolePermissionPolicyFinalizingCQRSResponse.transaction:type_name -> common.messaging.Transaction
	5, // 4: account.aggregates.template.TemplateRolePermissionPolicyFinalizingCQRSResponse.error:type_name -> common.messaging.Error
	6, // 5: account.aggregates.template.TemplateRolePermissionPolicyFinalizingCQRSResponse.data:type_name -> account.aggregates.template.TemplateRolePermissionPolicy
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_init()
}
func file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_init() {
	if File_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto != nil {
		return
	}
	file_proto_account_aggregates_template_template_role_permission_policy_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRolePermissionPolicyFinalizingCQRSRequestData); i {
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
		file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRolePermissionPolicyFinalizingCQRSRequest); i {
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
		file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRolePermissionPolicyFinalizingCQRSResponse); i {
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
			RawDescriptor: file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_goTypes,
		DependencyIndexes: file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_depIdxs,
		MessageInfos:      file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_msgTypes,
	}.Build()
	File_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto = out.File
	file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_rawDesc = nil
	file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_goTypes = nil
	file_proto_account_aggregates_template_template_role_permission_policy_finalizing_proto_depIdxs = nil
}