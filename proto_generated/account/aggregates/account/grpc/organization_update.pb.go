// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/account/aggregates/account/grpc/organization_update.proto

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
type OrganizationUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCRequest              `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *OrganizationUpdateRequestData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OrganizationUpdateRequest) Reset() {
	*x = OrganizationUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_organization_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationUpdateRequest) ProtoMessage() {}

func (x *OrganizationUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_organization_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationUpdateRequest.ProtoReflect.Descriptor instead.
func (*OrganizationUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_organization_update_proto_rawDescGZIP(), []int{0}
}

func (x *OrganizationUpdateRequest) GetBase() *grpc.GRPCRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *OrganizationUpdateRequest) GetData() *OrganizationUpdateRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type OrganizationUpdateRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint64                                       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string                                       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	GeneralSetting *OrganizationGeneralSettingCreateRequestData `protobuf:"bytes,3,opt,name=general_setting,json=generalSetting,proto3" json:"general_setting,omitempty"`
}

func (x *OrganizationUpdateRequestData) Reset() {
	*x = OrganizationUpdateRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_organization_update_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationUpdateRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationUpdateRequestData) ProtoMessage() {}

func (x *OrganizationUpdateRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_organization_update_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationUpdateRequestData.ProtoReflect.Descriptor instead.
func (*OrganizationUpdateRequestData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_organization_update_proto_rawDescGZIP(), []int{1}
}

func (x *OrganizationUpdateRequestData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrganizationUpdateRequestData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrganizationUpdateRequestData) GetGeneralSetting() *OrganizationGeneralSettingCreateRequestData {
	if x != nil {
		return x.GeneralSetting
	}
	return nil
}

// Response
type OrganizationUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCResponse              `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *OrganizationUpdateResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OrganizationUpdateResponse) Reset() {
	*x = OrganizationUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_organization_update_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationUpdateResponse) ProtoMessage() {}

func (x *OrganizationUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_organization_update_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationUpdateResponse.ProtoReflect.Descriptor instead.
func (*OrganizationUpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_organization_update_proto_rawDescGZIP(), []int{2}
}

func (x *OrganizationUpdateResponse) GetBase() *grpc.GRPCResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *OrganizationUpdateResponse) GetData() *OrganizationUpdateResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type OrganizationUpdateResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrganizationUpdateResponseData) Reset() {
	*x = OrganizationUpdateResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_organization_update_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationUpdateResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationUpdateResponseData) ProtoMessage() {}

func (x *OrganizationUpdateResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_organization_update_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationUpdateResponseData.ProtoReflect.Descriptor instead.
func (*OrganizationUpdateResponseData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_organization_update_proto_rawDescGZIP(), []int{3}
}

func (x *OrganizationUpdateResponseData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OrganizationGeneralSettingCreateRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Furigana           string `protobuf:"bytes,1,opt,name=furigana,proto3" json:"furigana,omitempty"`
	Zipcode            string `protobuf:"bytes,2,opt,name=zipcode,proto3" json:"zipcode,omitempty"`
	Country            string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Prefecture         string `protobuf:"bytes,4,opt,name=prefecture,proto3" json:"prefecture,omitempty"`
	Address            string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	StreetName         string `protobuf:"bytes,6,opt,name=street_name,json=streetName,proto3" json:"street_name,omitempty"`
	Building           string `protobuf:"bytes,7,opt,name=building,proto3" json:"building,omitempty"`
	Tel                string `protobuf:"bytes,8,opt,name=tel,proto3" json:"tel,omitempty"`
	Fax                string `protobuf:"bytes,9,opt,name=fax,proto3" json:"fax,omitempty"`
	EmailAddress       string `protobuf:"bytes,10,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	LogoImage          string `protobuf:"bytes,11,opt,name=logo_image,json=logoImage,proto3" json:"logo_image,omitempty"`
	SerialImprintImage string `protobuf:"bytes,12,opt,name=serial_imprint_image,json=serialImprintImage,proto3" json:"serial_imprint_image,omitempty"`
}

func (x *OrganizationGeneralSettingCreateRequestData) Reset() {
	*x = OrganizationGeneralSettingCreateRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_organization_update_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationGeneralSettingCreateRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationGeneralSettingCreateRequestData) ProtoMessage() {}

func (x *OrganizationGeneralSettingCreateRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_organization_update_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationGeneralSettingCreateRequestData.ProtoReflect.Descriptor instead.
func (*OrganizationGeneralSettingCreateRequestData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_organization_update_proto_rawDescGZIP(), []int{4}
}

func (x *OrganizationGeneralSettingCreateRequestData) GetFurigana() string {
	if x != nil {
		return x.Furigana
	}
	return ""
}

func (x *OrganizationGeneralSettingCreateRequestData) GetZipcode() string {
	if x != nil {
		return x.Zipcode
	}
	return ""
}

func (x *OrganizationGeneralSettingCreateRequestData) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *OrganizationGeneralSettingCreateRequestData) GetPrefecture() string {
	if x != nil {
		return x.Prefecture
	}
	return ""
}

func (x *OrganizationGeneralSettingCreateRequestData) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *OrganizationGeneralSettingCreateRequestData) GetStreetName() string {
	if x != nil {
		return x.StreetName
	}
	return ""
}

func (x *OrganizationGeneralSettingCreateRequestData) GetBuilding() string {
	if x != nil {
		return x.Building
	}
	return ""
}

func (x *OrganizationGeneralSettingCreateRequestData) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *OrganizationGeneralSettingCreateRequestData) GetFax() string {
	if x != nil {
		return x.Fax
	}
	return ""
}

func (x *OrganizationGeneralSettingCreateRequestData) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *OrganizationGeneralSettingCreateRequestData) GetLogoImage() string {
	if x != nil {
		return x.LogoImage
	}
	return ""
}

func (x *OrganizationGeneralSettingCreateRequestData) GetSerialImprintImage() string {
	if x != nil {
		return x.SerialImprintImage
	}
	return ""
}

var File_proto_account_aggregates_account_grpc_organization_update_proto protoreflect.FileDescriptor

var file_proto_account_aggregates_account_grpc_organization_update_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9d, 0x01, 0x0a, 0x19, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c,
	0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0xba, 0x01, 0x0a, 0x1d, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x75, 0x0a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x4c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x22, 0xa0, 0x01,
	0x0a, 0x1a, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04,
	0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x30, 0x0a, 0x1e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x8e, 0x03, 0x0a, 0x2b, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x72, 0x69, 0x67, 0x61, 0x6e, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x72, 0x69, 0x67, 0x61, 0x6e, 0x61, 0x12, 0x18,
	0x0a, 0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x66, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x66, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x65, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x66,
	0x61, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x61, 0x78, 0x12, 0x23, 0x0a,
	0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x6f, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x6d, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6d, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x42, 0x4d, 0x5a, 0x4b, 0x76, 0x6c, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x67, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_aggregates_account_grpc_organization_update_proto_rawDescOnce sync.Once
	file_proto_account_aggregates_account_grpc_organization_update_proto_rawDescData = file_proto_account_aggregates_account_grpc_organization_update_proto_rawDesc
)

func file_proto_account_aggregates_account_grpc_organization_update_proto_rawDescGZIP() []byte {
	file_proto_account_aggregates_account_grpc_organization_update_proto_rawDescOnce.Do(func() {
		file_proto_account_aggregates_account_grpc_organization_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_aggregates_account_grpc_organization_update_proto_rawDescData)
	})
	return file_proto_account_aggregates_account_grpc_organization_update_proto_rawDescData
}

var file_proto_account_aggregates_account_grpc_organization_update_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_account_aggregates_account_grpc_organization_update_proto_goTypes = []interface{}{
	(*OrganizationUpdateRequest)(nil),                   // 0: account.aggregates.account.grpc.OrganizationUpdateRequest
	(*OrganizationUpdateRequestData)(nil),               // 1: account.aggregates.account.grpc.OrganizationUpdateRequestData
	(*OrganizationUpdateResponse)(nil),                  // 2: account.aggregates.account.grpc.OrganizationUpdateResponse
	(*OrganizationUpdateResponseData)(nil),              // 3: account.aggregates.account.grpc.OrganizationUpdateResponseData
	(*OrganizationGeneralSettingCreateRequestData)(nil), // 4: account.aggregates.account.grpc.OrganizationGeneralSettingCreateRequestData
	(*grpc.GRPCRequest)(nil),                            // 5: common.grpc.GRPCRequest
	(*grpc.GRPCResponse)(nil),                           // 6: common.grpc.GRPCResponse
}
var file_proto_account_aggregates_account_grpc_organization_update_proto_depIdxs = []int32{
	5, // 0: account.aggregates.account.grpc.OrganizationUpdateRequest.base:type_name -> common.grpc.GRPCRequest
	1, // 1: account.aggregates.account.grpc.OrganizationUpdateRequest.data:type_name -> account.aggregates.account.grpc.OrganizationUpdateRequestData
	4, // 2: account.aggregates.account.grpc.OrganizationUpdateRequestData.general_setting:type_name -> account.aggregates.account.grpc.OrganizationGeneralSettingCreateRequestData
	6, // 3: account.aggregates.account.grpc.OrganizationUpdateResponse.base:type_name -> common.grpc.GRPCResponse
	3, // 4: account.aggregates.account.grpc.OrganizationUpdateResponse.data:type_name -> account.aggregates.account.grpc.OrganizationUpdateResponseData
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_account_aggregates_account_grpc_organization_update_proto_init() }
func file_proto_account_aggregates_account_grpc_organization_update_proto_init() {
	if File_proto_account_aggregates_account_grpc_organization_update_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_account_aggregates_account_grpc_organization_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationUpdateRequest); i {
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
		file_proto_account_aggregates_account_grpc_organization_update_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationUpdateRequestData); i {
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
		file_proto_account_aggregates_account_grpc_organization_update_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationUpdateResponse); i {
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
		file_proto_account_aggregates_account_grpc_organization_update_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationUpdateResponseData); i {
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
		file_proto_account_aggregates_account_grpc_organization_update_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationGeneralSettingCreateRequestData); i {
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
			RawDescriptor: file_proto_account_aggregates_account_grpc_organization_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_aggregates_account_grpc_organization_update_proto_goTypes,
		DependencyIndexes: file_proto_account_aggregates_account_grpc_organization_update_proto_depIdxs,
		MessageInfos:      file_proto_account_aggregates_account_grpc_organization_update_proto_msgTypes,
	}.Build()
	File_proto_account_aggregates_account_grpc_organization_update_proto = out.File
	file_proto_account_aggregates_account_grpc_organization_update_proto_rawDesc = nil
	file_proto_account_aggregates_account_grpc_organization_update_proto_goTypes = nil
	file_proto_account_aggregates_account_grpc_organization_update_proto_depIdxs = nil
}
