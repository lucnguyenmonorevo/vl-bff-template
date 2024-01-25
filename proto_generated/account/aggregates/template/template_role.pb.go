// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/account/aggregates/template/template_role.proto

package prototemplate

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TemplateRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SemanticLock  uint32                 `protobuf:"varint,2,opt,name=semantic_lock,json=semanticLock,proto3" json:"semantic_lock,omitempty"`
	InstanceId    string                 `protobuf:"bytes,3,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	SecondaryEpId string                 `protobuf:"bytes,5,opt,name=secondary_ep_id,json=secondaryEpId,proto3" json:"secondary_ep_id,omitempty"`
	EpTrxId       string                 `protobuf:"bytes,6,opt,name=ep_trx_id,json=epTrxId,proto3" json:"ep_trx_id,omitempty"`
	CqrsExecId    string                 `protobuf:"bytes,7,opt,name=cqrs_exec_id,json=cqrsExecId,proto3" json:"cqrs_exec_id,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy     string                 `protobuf:"bytes,9,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy     string                 `protobuf:"bytes,11,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	DeletedAt     *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	DeletedBy     string                 `protobuf:"bytes,13,opt,name=deleted_by,json=deletedBy,proto3" json:"deleted_by,omitempty"`
	Name          string                 `protobuf:"bytes,14,opt,name=name,proto3" json:"name,omitempty"`
	TemplateId    uint64                 `protobuf:"varint,15,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
}

func (x *TemplateRole) Reset() {
	*x = TemplateRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_template_template_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRole) ProtoMessage() {}

func (x *TemplateRole) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_template_template_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRole.ProtoReflect.Descriptor instead.
func (*TemplateRole) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_template_template_role_proto_rawDescGZIP(), []int{0}
}

func (x *TemplateRole) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TemplateRole) GetSemanticLock() uint32 {
	if x != nil {
		return x.SemanticLock
	}
	return 0
}

func (x *TemplateRole) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *TemplateRole) GetSecondaryEpId() string {
	if x != nil {
		return x.SecondaryEpId
	}
	return ""
}

func (x *TemplateRole) GetEpTrxId() string {
	if x != nil {
		return x.EpTrxId
	}
	return ""
}

func (x *TemplateRole) GetCqrsExecId() string {
	if x != nil {
		return x.CqrsExecId
	}
	return ""
}

func (x *TemplateRole) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *TemplateRole) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *TemplateRole) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *TemplateRole) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *TemplateRole) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *TemplateRole) GetDeletedBy() string {
	if x != nil {
		return x.DeletedBy
	}
	return ""
}

func (x *TemplateRole) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TemplateRole) GetTemplateId() uint64 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

var File_proto_account_aggregates_template_template_role_proto protoreflect.FileDescriptor

var file_proto_account_aggregates_template_template_role_proto_rawDesc = []byte{
	0x0a, 0x35, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x04, 0x0a, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74,
	0x69, 0x63, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73,
	0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x65, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79,
	0x45, 0x70, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x09, 0x65, 0x70, 0x5f, 0x74, 0x72, 0x78, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x70, 0x54, 0x72, 0x78, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0c, 0x63, 0x71, 0x72, 0x73, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x71, 0x72, 0x73, 0x45, 0x78, 0x65, 0x63,
	0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x42, 0x47, 0x5a, 0x45, 0x76, 0x6c, 0x2d, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_aggregates_template_template_role_proto_rawDescOnce sync.Once
	file_proto_account_aggregates_template_template_role_proto_rawDescData = file_proto_account_aggregates_template_template_role_proto_rawDesc
)

func file_proto_account_aggregates_template_template_role_proto_rawDescGZIP() []byte {
	file_proto_account_aggregates_template_template_role_proto_rawDescOnce.Do(func() {
		file_proto_account_aggregates_template_template_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_aggregates_template_template_role_proto_rawDescData)
	})
	return file_proto_account_aggregates_template_template_role_proto_rawDescData
}

var file_proto_account_aggregates_template_template_role_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_account_aggregates_template_template_role_proto_goTypes = []interface{}{
	(*TemplateRole)(nil),          // 0: account.aggregates.template.TemplateRole
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_proto_account_aggregates_template_template_role_proto_depIdxs = []int32{
	1, // 0: account.aggregates.template.TemplateRole.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: account.aggregates.template.TemplateRole.updated_at:type_name -> google.protobuf.Timestamp
	1, // 2: account.aggregates.template.TemplateRole.deleted_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_account_aggregates_template_template_role_proto_init() }
func file_proto_account_aggregates_template_template_role_proto_init() {
	if File_proto_account_aggregates_template_template_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_account_aggregates_template_template_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRole); i {
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
			RawDescriptor: file_proto_account_aggregates_template_template_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_aggregates_template_template_role_proto_goTypes,
		DependencyIndexes: file_proto_account_aggregates_template_template_role_proto_depIdxs,
		MessageInfos:      file_proto_account_aggregates_template_template_role_proto_msgTypes,
	}.Build()
	File_proto_account_aggregates_template_template_role_proto = out.File
	file_proto_account_aggregates_template_template_role_proto_rawDesc = nil
	file_proto_account_aggregates_template_template_role_proto_goTypes = nil
	file_proto_account_aggregates_template_template_role_proto_depIdxs = nil
}
