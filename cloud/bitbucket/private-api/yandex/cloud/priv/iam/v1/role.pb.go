// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.3
// source: yandex/cloud/priv/iam/v1/role.proto

package iam

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

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description   string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	PermissionIds []string `protobuf:"bytes,3,rep,name=permission_ids,json=permissionIds,proto3" json:"permission_ids,omitempty"`
	IsSystem      bool     `protobuf:"varint,4,opt,name=is_system,json=isSystem,proto3" json:"is_system,omitempty"`
	CategoryIds   []string `protobuf:"bytes,5,rep,name=category_ids,json=categoryIds,proto3" json:"category_ids,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_v1_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_v1_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_v1_role_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Role) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Role) GetPermissionIds() []string {
	if x != nil {
		return x.PermissionIds
	}
	return nil
}

func (x *Role) GetIsSystem() bool {
	if x != nil {
		return x.IsSystem
	}
	return false
}

func (x *Role) GetCategoryIds() []string {
	if x != nil {
		return x.CategoryIds
	}
	return nil
}

type RoleCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *RoleCategory) Reset() {
	*x = RoleCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_v1_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleCategory) ProtoMessage() {}

func (x *RoleCategory) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_v1_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleCategory.ProtoReflect.Descriptor instead.
func (*RoleCategory) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_v1_role_proto_rawDescGZIP(), []int{1}
}

func (x *RoleCategory) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoleCategory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleCategory) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_yandex_cloud_priv_iam_v1_role_proto protoreflect.FileDescriptor

var file_yandex_cloud_priv_iam_v1_role_proto_rawDesc = []byte{
	0x0a, 0x23, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x72, 0x69, 0x76, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x22,
	0x9f, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0d, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64,
	0x73, 0x22, 0x54, 0x0a, 0x0c, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x4f, 0x42, 0x02, 0x50, 0x52, 0x5a, 0x49, 0x61,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x72, 0x75, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_priv_iam_v1_role_proto_rawDescOnce sync.Once
	file_yandex_cloud_priv_iam_v1_role_proto_rawDescData = file_yandex_cloud_priv_iam_v1_role_proto_rawDesc
)

func file_yandex_cloud_priv_iam_v1_role_proto_rawDescGZIP() []byte {
	file_yandex_cloud_priv_iam_v1_role_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_priv_iam_v1_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_priv_iam_v1_role_proto_rawDescData)
	})
	return file_yandex_cloud_priv_iam_v1_role_proto_rawDescData
}

var file_yandex_cloud_priv_iam_v1_role_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_priv_iam_v1_role_proto_goTypes = []interface{}{
	(*Role)(nil),         // 0: yandex.cloud.priv.iam.v1.Role
	(*RoleCategory)(nil), // 1: yandex.cloud.priv.iam.v1.RoleCategory
}
var file_yandex_cloud_priv_iam_v1_role_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_priv_iam_v1_role_proto_init() }
func file_yandex_cloud_priv_iam_v1_role_proto_init() {
	if File_yandex_cloud_priv_iam_v1_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_priv_iam_v1_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_yandex_cloud_priv_iam_v1_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleCategory); i {
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
			RawDescriptor: file_yandex_cloud_priv_iam_v1_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_priv_iam_v1_role_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_priv_iam_v1_role_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_priv_iam_v1_role_proto_msgTypes,
	}.Build()
	File_yandex_cloud_priv_iam_v1_role_proto = out.File
	file_yandex_cloud_priv_iam_v1_role_proto_rawDesc = nil
	file_yandex_cloud_priv_iam_v1_role_proto_goTypes = nil
	file_yandex_cloud_priv_iam_v1_role_proto_depIdxs = nil
}
