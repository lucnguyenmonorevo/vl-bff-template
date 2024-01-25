// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/general_info/aggregates/organizational_settings/use_case/organizational_setting_create.proto

package protoorganizationalsettingsusecase

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

type OrganizationalSettingCreateUseCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction                         `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Client      *messaging.Client                              `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Data        *OrganizationalSettingCreateUseCaseRequestData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OrganizationalSettingCreateUseCaseRequest) Reset() {
	*x = OrganizationalSettingCreateUseCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationalSettingCreateUseCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationalSettingCreateUseCaseRequest) ProtoMessage() {}

func (x *OrganizationalSettingCreateUseCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationalSettingCreateUseCaseRequest.ProtoReflect.Descriptor instead.
func (*OrganizationalSettingCreateUseCaseRequest) Descriptor() ([]byte, []int) {
	return file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_rawDescGZIP(), []int{0}
}

func (x *OrganizationalSettingCreateUseCaseRequest) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *OrganizationalSettingCreateUseCaseRequest) GetClient() *messaging.Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *OrganizationalSettingCreateUseCaseRequest) GetData() *OrganizationalSettingCreateUseCaseRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type OrganizationalSettingCreateUseCaseRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrganizationGeneralSettingId uint64                                                `protobuf:"varint,1,opt,name=organization_general_setting_id,json=organizationGeneralSettingId,proto3" json:"organization_general_setting_id,omitempty"`
	OrganizationId               uint64                                                `protobuf:"varint,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	OrganizationalGeneralSetting *OrganizationalGeneralSettingCreateUseCaseRequestData `protobuf:"bytes,3,opt,name=organizational_general_setting,json=organizationalGeneralSetting,proto3" json:"organizational_general_setting,omitempty"`
}

func (x *OrganizationalSettingCreateUseCaseRequestData) Reset() {
	*x = OrganizationalSettingCreateUseCaseRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationalSettingCreateUseCaseRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationalSettingCreateUseCaseRequestData) ProtoMessage() {}

func (x *OrganizationalSettingCreateUseCaseRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationalSettingCreateUseCaseRequestData.ProtoReflect.Descriptor instead.
func (*OrganizationalSettingCreateUseCaseRequestData) Descriptor() ([]byte, []int) {
	return file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_rawDescGZIP(), []int{1}
}

func (x *OrganizationalSettingCreateUseCaseRequestData) GetOrganizationGeneralSettingId() uint64 {
	if x != nil {
		return x.OrganizationGeneralSettingId
	}
	return 0
}

func (x *OrganizationalSettingCreateUseCaseRequestData) GetOrganizationId() uint64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *OrganizationalSettingCreateUseCaseRequestData) GetOrganizationalGeneralSetting() *OrganizationalGeneralSettingCreateUseCaseRequestData {
	if x != nil {
		return x.OrganizationalGeneralSetting
	}
	return nil
}

type OrganizationalSettingCreateUseCaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Error       *messaging.Error       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data        *OrganizationalSetting `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OrganizationalSettingCreateUseCaseResponse) Reset() {
	*x = OrganizationalSettingCreateUseCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationalSettingCreateUseCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationalSettingCreateUseCaseResponse) ProtoMessage() {}

func (x *OrganizationalSettingCreateUseCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationalSettingCreateUseCaseResponse.ProtoReflect.Descriptor instead.
func (*OrganizationalSettingCreateUseCaseResponse) Descriptor() ([]byte, []int) {
	return file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_rawDescGZIP(), []int{2}
}

func (x *OrganizationalSettingCreateUseCaseResponse) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *OrganizationalSettingCreateUseCaseResponse) GetError() *messaging.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *OrganizationalSettingCreateUseCaseResponse) GetData() *OrganizationalSetting {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto protoreflect.FileDescriptor

var file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_rawDesc = []byte{
	0x0a, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65,
	0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x38, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x1a, 0x5b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x6a, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9b, 0x02, 0x0a, 0x29, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30,
	0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x7b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x67,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xd6, 0x02,
	0x0a, 0x2d, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x45, 0x0a, 0x1f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x1c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0xb4, 0x01, 0x0a, 0x1e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x6e, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63,
	0x61, 0x73, 0x65, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x1c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x81, 0x02, 0x0a, 0x2a, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x63, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x78, 0x5a, 0x76, 0x76, 0x6c,
	0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73,
	0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73,
	0x65, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x75, 0x73, 0x65,
	0x63, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_rawDescOnce sync.Once
	file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_rawDescData = file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_rawDesc
)

func file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_rawDescGZIP() []byte {
	file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_rawDescOnce.Do(func() {
		file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_rawDescData)
	})
	return file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_rawDescData
}

var file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_goTypes = []interface{}{
	(*OrganizationalSettingCreateUseCaseRequest)(nil),            // 0: general_info.aggregates.organizational_settings.use_case.OrganizationalSettingCreateUseCaseRequest
	(*OrganizationalSettingCreateUseCaseRequestData)(nil),        // 1: general_info.aggregates.organizational_settings.use_case.OrganizationalSettingCreateUseCaseRequestData
	(*OrganizationalSettingCreateUseCaseResponse)(nil),           // 2: general_info.aggregates.organizational_settings.use_case.OrganizationalSettingCreateUseCaseResponse
	(*messaging.Transaction)(nil),                                // 3: common.messaging.Transaction
	(*messaging.Client)(nil),                                     // 4: common.messaging.Client
	(*OrganizationalGeneralSettingCreateUseCaseRequestData)(nil), // 5: general_info.aggregates.organizational_settings.use_case.OrganizationalGeneralSettingCreateUseCaseRequestData
	(*messaging.Error)(nil),                                      // 6: common.messaging.Error
	(*OrganizationalSetting)(nil),                                // 7: general_info.aggregates.organizational_settings.use_case.OrganizationalSetting
}
var file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_depIdxs = []int32{
	3, // 0: general_info.aggregates.organizational_settings.use_case.OrganizationalSettingCreateUseCaseRequest.transaction:type_name -> common.messaging.Transaction
	4, // 1: general_info.aggregates.organizational_settings.use_case.OrganizationalSettingCreateUseCaseRequest.client:type_name -> common.messaging.Client
	1, // 2: general_info.aggregates.organizational_settings.use_case.OrganizationalSettingCreateUseCaseRequest.data:type_name -> general_info.aggregates.organizational_settings.use_case.OrganizationalSettingCreateUseCaseRequestData
	5, // 3: general_info.aggregates.organizational_settings.use_case.OrganizationalSettingCreateUseCaseRequestData.organizational_general_setting:type_name -> general_info.aggregates.organizational_settings.use_case.OrganizationalGeneralSettingCreateUseCaseRequestData
	3, // 4: general_info.aggregates.organizational_settings.use_case.OrganizationalSettingCreateUseCaseResponse.transaction:type_name -> common.messaging.Transaction
	6, // 5: general_info.aggregates.organizational_settings.use_case.OrganizationalSettingCreateUseCaseResponse.error:type_name -> common.messaging.Error
	7, // 6: general_info.aggregates.organizational_settings.use_case.OrganizationalSettingCreateUseCaseResponse.data:type_name -> general_info.aggregates.organizational_settings.use_case.OrganizationalSetting
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() {
	file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_init()
}
func file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_init() {
	if File_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto != nil {
		return
	}
	file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_proto_init()
	file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_create_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationalSettingCreateUseCaseRequest); i {
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
		file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationalSettingCreateUseCaseRequestData); i {
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
		file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationalSettingCreateUseCaseResponse); i {
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
			RawDescriptor: file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_goTypes,
		DependencyIndexes: file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_depIdxs,
		MessageInfos:      file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_msgTypes,
	}.Build()
	File_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto = out.File
	file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_rawDesc = nil
	file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_goTypes = nil
	file_proto_general_info_aggregates_organizational_settings_use_case_organizational_setting_create_proto_depIdxs = nil
}