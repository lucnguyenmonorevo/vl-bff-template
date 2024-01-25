// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/account/aggregates/permission/use_case/permission_user.proto

package protopermissionusecase

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

type PermissionUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SemanticLock   uint32                 `protobuf:"varint,2,opt,name=semantic_lock,json=semanticLock,proto3" json:"semantic_lock,omitempty"`
	InstanceId     string                 `protobuf:"bytes,3,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	EpTrxId        string                 `protobuf:"bytes,4,opt,name=ep_trx_id,json=epTrxId,proto3" json:"ep_trx_id,omitempty"`
	CqrsExecId     string                 `protobuf:"bytes,5,opt,name=cqrs_exec_id,json=cqrsExecId,proto3" json:"cqrs_exec_id,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy      string                 `protobuf:"bytes,7,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy      string                 `protobuf:"bytes,9,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	DeletedAt      *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	DeletedBy      string                 `protobuf:"bytes,11,opt,name=deleted_by,json=deletedBy,proto3" json:"deleted_by,omitempty"`
	UserId         uint64                 `protobuf:"varint,12,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrganizationId uint64                 `protobuf:"varint,13,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
}

func (x *PermissionUser) Reset() {
	*x = PermissionUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_aggregates_permission_use_case_permission_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionUser) ProtoMessage() {}

func (x *PermissionUser) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_aggregates_permission_use_case_permission_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionUser.ProtoReflect.Descriptor instead.
func (*PermissionUser) Descriptor() ([]byte, []int) {
	return file_proto_account_aggregates_permission_use_case_permission_user_proto_rawDescGZIP(), []int{0}
}

func (x *PermissionUser) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PermissionUser) GetSemanticLock() uint32 {
	if x != nil {
		return x.SemanticLock
	}
	return 0
}

func (x *PermissionUser) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *PermissionUser) GetEpTrxId() string {
	if x != nil {
		return x.EpTrxId
	}
	return ""
}

func (x *PermissionUser) GetCqrsExecId() string {
	if x != nil {
		return x.CqrsExecId
	}
	return ""
}

func (x *PermissionUser) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PermissionUser) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *PermissionUser) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *PermissionUser) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *PermissionUser) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *PermissionUser) GetDeletedBy() string {
	if x != nil {
		return x.DeletedBy
	}
	return ""
}

func (x *PermissionUser) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PermissionUser) GetOrganizationId() uint64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

var File_proto_account_aggregates_permission_use_case_permission_user_proto protoreflect.FileDescriptor

var file_proto_account_aggregates_permission_use_case_permission_user_proto_rawDesc = []byte{
	0x0a, 0x42, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x03,
	0x0a, 0x0e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x5f, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69,
	0x63, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x09, 0x65, 0x70, 0x5f, 0x74, 0x72, 0x78,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x70, 0x54, 0x72, 0x78,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x71, 0x72, 0x73, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x71, 0x72, 0x73, 0x45, 0x78,
	0x65, 0x63, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x42, 0x5a, 0x5a, 0x58, 0x76, 0x6c, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x75, 0x73, 0x65, 0x63, 0x61, 0x73, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_aggregates_permission_use_case_permission_user_proto_rawDescOnce sync.Once
	file_proto_account_aggregates_permission_use_case_permission_user_proto_rawDescData = file_proto_account_aggregates_permission_use_case_permission_user_proto_rawDesc
)

func file_proto_account_aggregates_permission_use_case_permission_user_proto_rawDescGZIP() []byte {
	file_proto_account_aggregates_permission_use_case_permission_user_proto_rawDescOnce.Do(func() {
		file_proto_account_aggregates_permission_use_case_permission_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_aggregates_permission_use_case_permission_user_proto_rawDescData)
	})
	return file_proto_account_aggregates_permission_use_case_permission_user_proto_rawDescData
}

var file_proto_account_aggregates_permission_use_case_permission_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_account_aggregates_permission_use_case_permission_user_proto_goTypes = []interface{}{
	(*PermissionUser)(nil),        // 0: account.aggregates.permission.use_case.PermissionUser
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_proto_account_aggregates_permission_use_case_permission_user_proto_depIdxs = []int32{
	1, // 0: account.aggregates.permission.use_case.PermissionUser.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: account.aggregates.permission.use_case.PermissionUser.updated_at:type_name -> google.protobuf.Timestamp
	1, // 2: account.aggregates.permission.use_case.PermissionUser.deleted_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_account_aggregates_permission_use_case_permission_user_proto_init() }
func file_proto_account_aggregates_permission_use_case_permission_user_proto_init() {
	if File_proto_account_aggregates_permission_use_case_permission_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_account_aggregates_permission_use_case_permission_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionUser); i {
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
			RawDescriptor: file_proto_account_aggregates_permission_use_case_permission_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_aggregates_permission_use_case_permission_user_proto_goTypes,
		DependencyIndexes: file_proto_account_aggregates_permission_use_case_permission_user_proto_depIdxs,
		MessageInfos:      file_proto_account_aggregates_permission_use_case_permission_user_proto_msgTypes,
	}.Build()
	File_proto_account_aggregates_permission_use_case_permission_user_proto = out.File
	file_proto_account_aggregates_permission_use_case_permission_user_proto_rawDesc = nil
	file_proto_account_aggregates_permission_use_case_permission_user_proto_goTypes = nil
	file_proto_account_aggregates_permission_use_case_permission_user_proto_depIdxs = nil
}
