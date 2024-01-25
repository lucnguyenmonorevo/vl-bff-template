// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/notification/aggregates/role_notification/organizational_role_notification_setting.proto

package protorolenotification

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrganizationalRoleNotificationSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductionLotEnable                          uint32 `protobuf:"varint,1,opt,name=production_lot_enable,json=productionLotEnable,proto3" json:"production_lot_enable,omitempty"`
	ProductionLotSystemInternalNotification      uint32 `protobuf:"varint,2,opt,name=production_lot_system_internal_notification,json=productionLotSystemInternalNotification,proto3" json:"production_lot_system_internal_notification,omitempty"`
	ProductionLotEmailNotification               uint32 `protobuf:"varint,3,opt,name=production_lot_email_notification,json=productionLotEmailNotification,proto3" json:"production_lot_email_notification,omitempty"`
	ProcurementLotEnable                         uint32 `protobuf:"varint,4,opt,name=procurement_lot_enable,json=procurementLotEnable,proto3" json:"procurement_lot_enable,omitempty"`
	ProcurementLotSystemInternalNotification     uint32 `protobuf:"varint,5,opt,name=procurement_lot_system_internal_notification,json=procurementLotSystemInternalNotification,proto3" json:"procurement_lot_system_internal_notification,omitempty"`
	ProcurementLotEmailNotification              uint32 `protobuf:"varint,6,opt,name=procurement_lot_email_notification,json=procurementLotEmailNotification,proto3" json:"procurement_lot_email_notification,omitempty"`
	ProcessMemoEnable                            uint32 `protobuf:"varint,7,opt,name=process_memo_enable,json=processMemoEnable,proto3" json:"process_memo_enable,omitempty"`
	ProcessMemoSystemInternalNotification        uint32 `protobuf:"varint,8,opt,name=process_memo_system_internal_notification,json=processMemoSystemInternalNotification,proto3" json:"process_memo_system_internal_notification,omitempty"`
	ProcessMemoEmailNotification                 uint32 `protobuf:"varint,9,opt,name=process_memo_email_notification,json=processMemoEmailNotification,proto3" json:"process_memo_email_notification,omitempty"`
	DefectEnable                                 uint32 `protobuf:"varint,10,opt,name=defect_enable,json=defectEnable,proto3" json:"defect_enable,omitempty"`
	DefectSystemInternalNotification             uint32 `protobuf:"varint,11,opt,name=defect_system_internal_notification,json=defectSystemInternalNotification,proto3" json:"defect_system_internal_notification,omitempty"`
	DefectEmailNotification                      uint32 `protobuf:"varint,12,opt,name=defect_email_notification,json=defectEmailNotification,proto3" json:"defect_email_notification,omitempty"`
	UnexpectedStoppageEnable                     uint32 `protobuf:"varint,13,opt,name=unexpected_stoppage_enable,json=unexpectedStoppageEnable,proto3" json:"unexpected_stoppage_enable,omitempty"`
	UnexpectedStoppageSystemInternalNotification uint32 `protobuf:"varint,14,opt,name=unexpected_stoppage_system_internal_notification,json=unexpectedStoppageSystemInternalNotification,proto3" json:"unexpected_stoppage_system_internal_notification,omitempty"`
	UnexpectedStoppageEmailNotification          uint32 `protobuf:"varint,15,opt,name=unexpected_stoppage_email_notification,json=unexpectedStoppageEmailNotification,proto3" json:"unexpected_stoppage_email_notification,omitempty"`
}

func (x *OrganizationalRoleNotificationSetting) Reset() {
	*x = OrganizationalRoleNotificationSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationalRoleNotificationSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationalRoleNotificationSetting) ProtoMessage() {}

func (x *OrganizationalRoleNotificationSetting) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationalRoleNotificationSetting.ProtoReflect.Descriptor instead.
func (*OrganizationalRoleNotificationSetting) Descriptor() ([]byte, []int) {
	return file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_rawDescGZIP(), []int{0}
}

func (x *OrganizationalRoleNotificationSetting) GetProductionLotEnable() uint32 {
	if x != nil {
		return x.ProductionLotEnable
	}
	return 0
}

func (x *OrganizationalRoleNotificationSetting) GetProductionLotSystemInternalNotification() uint32 {
	if x != nil {
		return x.ProductionLotSystemInternalNotification
	}
	return 0
}

func (x *OrganizationalRoleNotificationSetting) GetProductionLotEmailNotification() uint32 {
	if x != nil {
		return x.ProductionLotEmailNotification
	}
	return 0
}

func (x *OrganizationalRoleNotificationSetting) GetProcurementLotEnable() uint32 {
	if x != nil {
		return x.ProcurementLotEnable
	}
	return 0
}

func (x *OrganizationalRoleNotificationSetting) GetProcurementLotSystemInternalNotification() uint32 {
	if x != nil {
		return x.ProcurementLotSystemInternalNotification
	}
	return 0
}

func (x *OrganizationalRoleNotificationSetting) GetProcurementLotEmailNotification() uint32 {
	if x != nil {
		return x.ProcurementLotEmailNotification
	}
	return 0
}

func (x *OrganizationalRoleNotificationSetting) GetProcessMemoEnable() uint32 {
	if x != nil {
		return x.ProcessMemoEnable
	}
	return 0
}

func (x *OrganizationalRoleNotificationSetting) GetProcessMemoSystemInternalNotification() uint32 {
	if x != nil {
		return x.ProcessMemoSystemInternalNotification
	}
	return 0
}

func (x *OrganizationalRoleNotificationSetting) GetProcessMemoEmailNotification() uint32 {
	if x != nil {
		return x.ProcessMemoEmailNotification
	}
	return 0
}

func (x *OrganizationalRoleNotificationSetting) GetDefectEnable() uint32 {
	if x != nil {
		return x.DefectEnable
	}
	return 0
}

func (x *OrganizationalRoleNotificationSetting) GetDefectSystemInternalNotification() uint32 {
	if x != nil {
		return x.DefectSystemInternalNotification
	}
	return 0
}

func (x *OrganizationalRoleNotificationSetting) GetDefectEmailNotification() uint32 {
	if x != nil {
		return x.DefectEmailNotification
	}
	return 0
}

func (x *OrganizationalRoleNotificationSetting) GetUnexpectedStoppageEnable() uint32 {
	if x != nil {
		return x.UnexpectedStoppageEnable
	}
	return 0
}

func (x *OrganizationalRoleNotificationSetting) GetUnexpectedStoppageSystemInternalNotification() uint32 {
	if x != nil {
		return x.UnexpectedStoppageSystemInternalNotification
	}
	return 0
}

func (x *OrganizationalRoleNotificationSetting) GetUnexpectedStoppageEmailNotification() uint32 {
	if x != nil {
		return x.UnexpectedStoppageEmailNotification
	}
	return 0
}

var File_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto protoreflect.FileDescriptor

var file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_rawDesc = []byte{
	0x0a, 0x5e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x29, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe3, 0x08, 0x0a, 0x25,
	0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x6f,
	0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x32, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x74, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x6f, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x5c, 0x0a, 0x2b, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x27,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x74, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x21, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x1e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f,
	0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x16, 0x70, 0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x6f, 0x74, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x14, 0x70, 0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c,
	0x6f, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x5e, 0x0a, 0x2c, 0x70, 0x72, 0x6f, 0x63,
	0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x28,
	0x70, 0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x74, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x22, 0x70, 0x72, 0x6f, 0x63,
	0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x1f, 0x70, 0x72, 0x6f, 0x63, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x4c, 0x6f, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x6d, 0x6f, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x58, 0x0a, 0x29, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x25, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x4d, 0x65, 0x6d, 0x6f, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x45, 0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1c, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x4d, 0x65, 0x6d, 0x6f, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x65, 0x63, 0x74,
	0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x64,
	0x65, 0x66, 0x65, 0x63, 0x74, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x4d, 0x0a, 0x23, 0x64,
	0x65, 0x66, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x20, 0x64, 0x65, 0x66, 0x65, 0x63, 0x74,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x19, 0x64, 0x65,
	0x66, 0x65, 0x63, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x17, 0x64,
	0x65, 0x66, 0x65, 0x63, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x1a, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x75, 0x6e, 0x65, 0x78,
	0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x70, 0x70, 0x61, 0x67, 0x65, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x66, 0x0a, 0x30, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x2c,
	0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x70, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x26,
	0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x23, 0x75, 0x6e,
	0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x70, 0x70, 0x61, 0x67, 0x65,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x5c, 0x5a, 0x5a, 0x76, 0x6c, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72,
	0x6f, 0x6c, 0x65, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_rawDescOnce sync.Once
	file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_rawDescData = file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_rawDesc
)

func file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_rawDescGZIP() []byte {
	file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_rawDescOnce.Do(func() {
		file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_rawDescData)
	})
	return file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_rawDescData
}

var file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_goTypes = []interface{}{
	(*OrganizationalRoleNotificationSetting)(nil), // 0: notification.aggregates.role_notification.OrganizationalRoleNotificationSetting
}
var file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_init()
}
func file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_init() {
	if File_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationalRoleNotificationSetting); i {
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
			RawDescriptor: file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_goTypes,
		DependencyIndexes: file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_depIdxs,
		MessageInfos:      file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_msgTypes,
	}.Build()
	File_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto = out.File
	file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_rawDesc = nil
	file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_goTypes = nil
	file_proto_notification_aggregates_role_notification_organizational_role_notification_setting_proto_depIdxs = nil
}
