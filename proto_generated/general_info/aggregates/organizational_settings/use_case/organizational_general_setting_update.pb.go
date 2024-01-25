// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/general_info/aggregates/organizational_settings/use_case/organizational_general_setting_update.proto

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

type OrganizationalGeneralSettingUpdateUseCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction                                `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Client      *messaging.Client                                     `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Data        *OrganizationalGeneralSettingUpdateUseCaseRequestData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequest) Reset() {
	*x = OrganizationalGeneralSettingUpdateUseCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationalGeneralSettingUpdateUseCaseRequest) ProtoMessage() {}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationalGeneralSettingUpdateUseCaseRequest.ProtoReflect.Descriptor instead.
func (*OrganizationalGeneralSettingUpdateUseCaseRequest) Descriptor() ([]byte, []int) {
	return file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_rawDescGZIP(), []int{0}
}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequest) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequest) GetClient() *messaging.Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequest) GetData() *OrganizationalGeneralSettingUpdateUseCaseRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type OrganizationalGeneralSettingUpdateUseCaseRequestData struct {
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

func (x *OrganizationalGeneralSettingUpdateUseCaseRequestData) Reset() {
	*x = OrganizationalGeneralSettingUpdateUseCaseRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationalGeneralSettingUpdateUseCaseRequestData) ProtoMessage() {}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationalGeneralSettingUpdateUseCaseRequestData.ProtoReflect.Descriptor instead.
func (*OrganizationalGeneralSettingUpdateUseCaseRequestData) Descriptor() ([]byte, []int) {
	return file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_rawDescGZIP(), []int{1}
}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequestData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequestData) GetFurigana() string {
	if x != nil {
		return x.Furigana
	}
	return ""
}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequestData) GetZipcode() string {
	if x != nil {
		return x.Zipcode
	}
	return ""
}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequestData) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequestData) GetPrefecture() string {
	if x != nil {
		return x.Prefecture
	}
	return ""
}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequestData) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequestData) GetStreetName() string {
	if x != nil {
		return x.StreetName
	}
	return ""
}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequestData) GetBuilding() string {
	if x != nil {
		return x.Building
	}
	return ""
}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequestData) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequestData) GetFax() string {
	if x != nil {
		return x.Fax
	}
	return ""
}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequestData) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequestData) GetLogoImage() string {
	if x != nil {
		return x.LogoImage
	}
	return ""
}

func (x *OrganizationalGeneralSettingUpdateUseCaseRequestData) GetSerialImprintImage() string {
	if x != nil {
		return x.SerialImprintImage
	}
	return ""
}

type OrganizationalGeneralSettingUpdateUseCaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *messaging.Transaction        `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Error       *messaging.Error              `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data        *OrganizationalGeneralSetting `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OrganizationalGeneralSettingUpdateUseCaseResponse) Reset() {
	*x = OrganizationalGeneralSettingUpdateUseCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationalGeneralSettingUpdateUseCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationalGeneralSettingUpdateUseCaseResponse) ProtoMessage() {}

func (x *OrganizationalGeneralSettingUpdateUseCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationalGeneralSettingUpdateUseCaseResponse.ProtoReflect.Descriptor instead.
func (*OrganizationalGeneralSettingUpdateUseCaseResponse) Descriptor() ([]byte, []int) {
	return file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_rawDescGZIP(), []int{2}
}

func (x *OrganizationalGeneralSettingUpdateUseCaseResponse) GetTransaction() *messaging.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *OrganizationalGeneralSettingUpdateUseCaseResponse) GetError() *messaging.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *OrganizationalGeneralSettingUpdateUseCaseResponse) GetData() *OrganizationalGeneralSetting {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto protoreflect.FileDescriptor

var file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_rawDesc = []byte{
	0x0a, 0x6a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65,
	0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x38, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x75, 0x73,
	0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x1a, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x75, 0x73,
	0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x02, 0x0a, 0x30, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x06, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x82, 0x01, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x6e, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x75, 0x73, 0x65,
	0x5f, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0xa7, 0x03, 0x0a, 0x34, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x72,
	0x69, 0x67, 0x61, 0x6e, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x72,
	0x69, 0x67, 0x61, 0x6e, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x65,
	0x66, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74,
	0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x61, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x66, 0x61, 0x78, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67,
	0x6f, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c,
	0x6f, 0x67, 0x6f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x5f, 0x69, 0x6d, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x6d,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x8f, 0x02, 0x0a, 0x31, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x6a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x56,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x78, 0x5a, 0x76,
	0x76, 0x6c, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63,
	0x61, 0x73, 0x65, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x75,
	0x73, 0x65, 0x63, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_rawDescOnce sync.Once
	file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_rawDescData = file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_rawDesc
)

func file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_rawDescGZIP() []byte {
	file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_rawDescOnce.Do(func() {
		file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_rawDescData)
	})
	return file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_rawDescData
}

var file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_goTypes = []interface{}{
	(*OrganizationalGeneralSettingUpdateUseCaseRequest)(nil),     // 0: general_info.aggregates.organizational_settings.use_case.OrganizationalGeneralSettingUpdateUseCaseRequest
	(*OrganizationalGeneralSettingUpdateUseCaseRequestData)(nil), // 1: general_info.aggregates.organizational_settings.use_case.OrganizationalGeneralSettingUpdateUseCaseRequestData
	(*OrganizationalGeneralSettingUpdateUseCaseResponse)(nil),    // 2: general_info.aggregates.organizational_settings.use_case.OrganizationalGeneralSettingUpdateUseCaseResponse
	(*messaging.Transaction)(nil),                                // 3: common.messaging.Transaction
	(*messaging.Client)(nil),                                     // 4: common.messaging.Client
	(*messaging.Error)(nil),                                      // 5: common.messaging.Error
	(*OrganizationalGeneralSetting)(nil),                         // 6: general_info.aggregates.organizational_settings.use_case.OrganizationalGeneralSetting
}
var file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_depIdxs = []int32{
	3, // 0: general_info.aggregates.organizational_settings.use_case.OrganizationalGeneralSettingUpdateUseCaseRequest.transaction:type_name -> common.messaging.Transaction
	4, // 1: general_info.aggregates.organizational_settings.use_case.OrganizationalGeneralSettingUpdateUseCaseRequest.client:type_name -> common.messaging.Client
	1, // 2: general_info.aggregates.organizational_settings.use_case.OrganizationalGeneralSettingUpdateUseCaseRequest.data:type_name -> general_info.aggregates.organizational_settings.use_case.OrganizationalGeneralSettingUpdateUseCaseRequestData
	3, // 3: general_info.aggregates.organizational_settings.use_case.OrganizationalGeneralSettingUpdateUseCaseResponse.transaction:type_name -> common.messaging.Transaction
	5, // 4: general_info.aggregates.organizational_settings.use_case.OrganizationalGeneralSettingUpdateUseCaseResponse.error:type_name -> common.messaging.Error
	6, // 5: general_info.aggregates.organizational_settings.use_case.OrganizationalGeneralSettingUpdateUseCaseResponse.data:type_name -> general_info.aggregates.organizational_settings.use_case.OrganizationalGeneralSetting
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_init()
}
func file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_init() {
	if File_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto != nil {
		return
	}
	file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationalGeneralSettingUpdateUseCaseRequest); i {
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
		file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationalGeneralSettingUpdateUseCaseRequestData); i {
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
		file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationalGeneralSettingUpdateUseCaseResponse); i {
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
			RawDescriptor: file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_goTypes,
		DependencyIndexes: file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_depIdxs,
		MessageInfos:      file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_msgTypes,
	}.Build()
	File_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto = out.File
	file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_rawDesc = nil
	file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_goTypes = nil
	file_proto_general_info_aggregates_organizational_settings_use_case_organizational_general_setting_update_proto_depIdxs = nil
}
