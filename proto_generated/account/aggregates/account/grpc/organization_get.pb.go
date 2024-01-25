// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/account/aggregates/account/grpc/organization_get.proto

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
type OrganizationGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCRequest           `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *OrganizationGetRequestData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OrganizationGetRequest) Reset() {
	*x = OrganizationGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationGetRequest) ProtoMessage() {}

func (x *OrganizationGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationGetRequest.ProtoReflect.Descriptor instead.
func (*OrganizationGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_organization_get_proto_rawDescGZIP(), []int{0}
}

func (x *OrganizationGetRequest) GetBase() *grpc.GRPCRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *OrganizationGetRequest) GetData() *OrganizationGetRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type OrganizationGetRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrganizationGetRequestData) Reset() {
	*x = OrganizationGetRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationGetRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationGetRequestData) ProtoMessage() {}

func (x *OrganizationGetRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationGetRequestData.ProtoReflect.Descriptor instead.
func (*OrganizationGetRequestData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_organization_get_proto_rawDescGZIP(), []int{1}
}

func (x *OrganizationGetRequestData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OrganizationGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *grpc.GRPCResponse           `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Data *OrganizationGetResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OrganizationGetResponse) Reset() {
	*x = OrganizationGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationGetResponse) ProtoMessage() {}

func (x *OrganizationGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationGetResponse.ProtoReflect.Descriptor instead.
func (*OrganizationGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_organization_get_proto_rawDescGZIP(), []int{2}
}

func (x *OrganizationGetResponse) GetBase() *grpc.GRPCResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *OrganizationGetResponse) GetData() *OrganizationGetResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

// Response
type OrganizationGetResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint64                      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string                      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Setting        *OrganizationSetting        `protobuf:"bytes,3,opt,name=setting,proto3" json:"setting,omitempty"`
	GeneralSetting *OrganizationGeneralSetting `protobuf:"bytes,4,opt,name=general_setting,json=generalSetting,proto3" json:"general_setting,omitempty"`
}

func (x *OrganizationGetResponseData) Reset() {
	*x = OrganizationGetResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationGetResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationGetResponseData) ProtoMessage() {}

func (x *OrganizationGetResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationGetResponseData.ProtoReflect.Descriptor instead.
func (*OrganizationGetResponseData) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_organization_get_proto_rawDescGZIP(), []int{3}
}

func (x *OrganizationGetResponseData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrganizationGetResponseData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrganizationGetResponseData) GetSetting() *OrganizationSetting {
	if x != nil {
		return x.Setting
	}
	return nil
}

func (x *OrganizationGetResponseData) GetGeneralSetting() *OrganizationGeneralSetting {
	if x != nil {
		return x.GeneralSetting
	}
	return nil
}

type OrganizationSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                           uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrganizationId               uint64 `protobuf:"varint,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	OrganizationGeneralSettingId uint64 `protobuf:"varint,3,opt,name=organization_general_setting_id,json=organizationGeneralSettingId,proto3" json:"organization_general_setting_id,omitempty"`
}

func (x *OrganizationSetting) Reset() {
	*x = OrganizationSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationSetting) ProtoMessage() {}

func (x *OrganizationSetting) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationSetting.ProtoReflect.Descriptor instead.
func (*OrganizationSetting) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_organization_get_proto_rawDescGZIP(), []int{4}
}

func (x *OrganizationSetting) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrganizationSetting) GetOrganizationId() uint64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *OrganizationSetting) GetOrganizationGeneralSettingId() uint64 {
	if x != nil {
		return x.OrganizationGeneralSettingId
	}
	return 0
}

type OrganizationGeneralSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Furigana           string `protobuf:"bytes,2,opt,name=furigana,proto3" json:"furigana,omitempty"`
	Zipcode            string `protobuf:"bytes,3,opt,name=zipcode,proto3" json:"zipcode,omitempty"`
	Country            string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Prefecture         string `protobuf:"bytes,5,opt,name=prefecture,proto3" json:"prefecture,omitempty"`
	Address            string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	StreetName         string `protobuf:"bytes,7,opt,name=street_name,json=streetName,proto3" json:"street_name,omitempty"`
	Building           string `protobuf:"bytes,8,opt,name=building,proto3" json:"building,omitempty"`
	Tel                string `protobuf:"bytes,9,opt,name=tel,proto3" json:"tel,omitempty"`
	Fax                string `protobuf:"bytes,10,opt,name=fax,proto3" json:"fax,omitempty"`
	EmailAddress       string `protobuf:"bytes,11,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	LogoImage          string `protobuf:"bytes,12,opt,name=logo_image,json=logoImage,proto3" json:"logo_image,omitempty"`
	SerialImprintImage string `protobuf:"bytes,13,opt,name=serial_imprint_image,json=serialImprintImage,proto3" json:"serial_imprint_image,omitempty"`
}

func (x *OrganizationGeneralSetting) Reset() {
	*x = OrganizationGeneralSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationGeneralSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationGeneralSetting) ProtoMessage() {}

func (x *OrganizationGeneralSetting) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationGeneralSetting.ProtoReflect.Descriptor instead.
func (*OrganizationGeneralSetting) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_account_grpc_organization_get_proto_rawDescGZIP(), []int{5}
}

func (x *OrganizationGeneralSetting) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrganizationGeneralSetting) GetFurigana() string {
	if x != nil {
		return x.Furigana
	}
	return ""
}

func (x *OrganizationGeneralSetting) GetZipcode() string {
	if x != nil {
		return x.Zipcode
	}
	return ""
}

func (x *OrganizationGeneralSetting) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *OrganizationGeneralSetting) GetPrefecture() string {
	if x != nil {
		return x.Prefecture
	}
	return ""
}

func (x *OrganizationGeneralSetting) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *OrganizationGeneralSetting) GetStreetName() string {
	if x != nil {
		return x.StreetName
	}
	return ""
}

func (x *OrganizationGeneralSetting) GetBuilding() string {
	if x != nil {
		return x.Building
	}
	return ""
}

func (x *OrganizationGeneralSetting) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *OrganizationGeneralSetting) GetFax() string {
	if x != nil {
		return x.Fax
	}
	return ""
}

func (x *OrganizationGeneralSetting) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *OrganizationGeneralSetting) GetLogoImage() string {
	if x != nil {
		return x.LogoImage
	}
	return ""
}

func (x *OrganizationGeneralSetting) GetSerialImprintImage() string {
	if x != nil {
		return x.SerialImprintImage
	}
	return ""
}

var File_proto_account_aggregates_account_grpc_organization_get_proto protoreflect.FileDescriptor

var file_proto_account_aggregates_account_grpc_organization_get_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a,
	0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01,
	0x0a, 0x16, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2c, 0x0a, 0x1a, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9a, 0x01, 0x0a, 0x17, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x52,
	0x50, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x50, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0xf7, 0x01, 0x0a, 0x1b, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x07, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x64, 0x0a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3b, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x95, 0x01, 0x0a,
	0x13, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x45, 0x0a,
	0x1f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x1c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x49, 0x64, 0x22, 0x8d, 0x03, 0x0a, 0x1a, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x72, 0x69, 0x67, 0x61, 0x6e, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x72, 0x69, 0x67, 0x61, 0x6e, 0x61, 0x12,
	0x18, 0x0a, 0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x66, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x66, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x65,
	0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03,
	0x66, 0x61, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x61, 0x78, 0x12, 0x23,
	0x0a, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x6f, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x6d, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6d, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x42, 0x4d, 0x5a, 0x4b, 0x76, 0x6c, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_aggregates_account_grpc_organization_get_proto_rawDescOnce sync.Once
	file_proto_account_aggregates_account_grpc_organization_get_proto_rawDescData = file_proto_account_aggregates_account_grpc_organization_get_proto_rawDesc
)

func file_proto_account_aggregates_account_grpc_organization_get_proto_rawDescGZIP() []byte {
	file_proto_account_aggregates_account_grpc_organization_get_proto_rawDescOnce.Do(func() {
		file_proto_account_aggregates_account_grpc_organization_get_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_aggregates_account_grpc_organization_get_proto_rawDescData)
	})
	return file_proto_account_aggregates_account_grpc_organization_get_proto_rawDescData
}

var file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_account_aggregates_account_grpc_organization_get_proto_goTypes = []interface{}{
	(*OrganizationGetRequest)(nil),      // 0: account.aggregates.account.grpc.OrganizationGetRequest
	(*OrganizationGetRequestData)(nil),  // 1: account.aggregates.account.grpc.OrganizationGetRequestData
	(*OrganizationGetResponse)(nil),     // 2: account.aggregates.account.grpc.OrganizationGetResponse
	(*OrganizationGetResponseData)(nil), // 3: account.aggregates.account.grpc.OrganizationGetResponseData
	(*OrganizationSetting)(nil),         // 4: account.aggregates.account.grpc.OrganizationSetting
	(*OrganizationGeneralSetting)(nil),  // 5: account.aggregates.account.grpc.OrganizationGeneralSetting
	(*grpc.GRPCRequest)(nil),            // 6: common.grpc.GRPCRequest
	(*grpc.GRPCResponse)(nil),           // 7: common.grpc.GRPCResponse
}
var file_proto_account_aggregates_account_grpc_organization_get_proto_depIdxs = []int32{
	6, // 0: account.aggregates.account.grpc.OrganizationGetRequest.base:type_name -> common.grpc.GRPCRequest
	1, // 1: account.aggregates.account.grpc.OrganizationGetRequest.data:type_name -> account.aggregates.account.grpc.OrganizationGetRequestData
	7, // 2: account.aggregates.account.grpc.OrganizationGetResponse.base:type_name -> common.grpc.GRPCResponse
	3, // 3: account.aggregates.account.grpc.OrganizationGetResponse.data:type_name -> account.aggregates.account.grpc.OrganizationGetResponseData
	4, // 4: account.aggregates.account.grpc.OrganizationGetResponseData.setting:type_name -> account.aggregates.account.grpc.OrganizationSetting
	5, // 5: account.aggregates.account.grpc.OrganizationGetResponseData.general_setting:type_name -> account.aggregates.account.grpc.OrganizationGeneralSetting
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_account_aggregates_account_grpc_organization_get_proto_init() }
func file_proto_account_aggregates_account_grpc_organization_get_proto_init() {
	if File_proto_account_aggregates_account_grpc_organization_get_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationGetRequest); i {
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
		file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationGetRequestData); i {
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
		file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationGetResponse); i {
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
		file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationGetResponseData); i {
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
		file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationSetting); i {
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
		file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationGeneralSetting); i {
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
			RawDescriptor: file_proto_account_aggregates_account_grpc_organization_get_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_aggregates_account_grpc_organization_get_proto_goTypes,
		DependencyIndexes: file_proto_account_aggregates_account_grpc_organization_get_proto_depIdxs,
		MessageInfos:      file_proto_account_aggregates_account_grpc_organization_get_proto_msgTypes,
	}.Build()
	File_proto_account_aggregates_account_grpc_organization_get_proto = out.File
	file_proto_account_aggregates_account_grpc_organization_get_proto_rawDesc = nil
	file_proto_account_aggregates_account_grpc_organization_get_proto_goTypes = nil
	file_proto_account_aggregates_account_grpc_organization_get_proto_depIdxs = nil
}
