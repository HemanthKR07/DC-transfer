// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.3
// source: yandex/cloud/priv/iam/v1/api_key.proto

package iam

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

type ApiKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ServiceAccountId string                 `protobuf:"bytes,2,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	CreatedAt        *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Description      string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	LastUsedAt       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_used_at,json=lastUsedAt,proto3" json:"last_used_at,omitempty"`
	Scope            string                 `protobuf:"bytes,6,opt,name=scope,proto3" json:"scope,omitempty"`
	ExpiresAt        *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *ApiKey) Reset() {
	*x = ApiKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_v1_api_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKey) ProtoMessage() {}

func (x *ApiKey) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_v1_api_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKey.ProtoReflect.Descriptor instead.
func (*ApiKey) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_v1_api_key_proto_rawDescGZIP(), []int{0}
}

func (x *ApiKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ApiKey) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *ApiKey) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ApiKey) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ApiKey) GetLastUsedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUsedAt
	}
	return nil
}

func (x *ApiKey) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *ApiKey) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

var File_yandex_cloud_priv_iam_v1_api_key_proto protoreflect.FileDescriptor

var file_yandex_cloud_priv_iam_v1_api_key_proto_rawDesc = []byte{
	0x0a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x72, 0x69, 0x76, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6b,
	0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x02, 0x0a, 0x06, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c,
	0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x0c, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x42, 0x52, 0x42, 0x05, 0x50, 0x41, 0x50, 0x49,
	0x4b, 0x5a, 0x49, 0x61, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x74, 0x65, 0x61, 0x6d,
	0x2e, 0x72, 0x75, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69, 0x74, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x69,
	0x76, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_priv_iam_v1_api_key_proto_rawDescOnce sync.Once
	file_yandex_cloud_priv_iam_v1_api_key_proto_rawDescData = file_yandex_cloud_priv_iam_v1_api_key_proto_rawDesc
)

func file_yandex_cloud_priv_iam_v1_api_key_proto_rawDescGZIP() []byte {
	file_yandex_cloud_priv_iam_v1_api_key_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_priv_iam_v1_api_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_priv_iam_v1_api_key_proto_rawDescData)
	})
	return file_yandex_cloud_priv_iam_v1_api_key_proto_rawDescData
}

var file_yandex_cloud_priv_iam_v1_api_key_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_priv_iam_v1_api_key_proto_goTypes = []interface{}{
	(*ApiKey)(nil),                // 0: yandex.cloud.priv.iam.v1.ApiKey
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_yandex_cloud_priv_iam_v1_api_key_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.priv.iam.v1.ApiKey.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: yandex.cloud.priv.iam.v1.ApiKey.last_used_at:type_name -> google.protobuf.Timestamp
	1, // 2: yandex.cloud.priv.iam.v1.ApiKey.expires_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_priv_iam_v1_api_key_proto_init() }
func file_yandex_cloud_priv_iam_v1_api_key_proto_init() {
	if File_yandex_cloud_priv_iam_v1_api_key_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_priv_iam_v1_api_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiKey); i {
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
			RawDescriptor: file_yandex_cloud_priv_iam_v1_api_key_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_priv_iam_v1_api_key_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_priv_iam_v1_api_key_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_priv_iam_v1_api_key_proto_msgTypes,
	}.Build()
	File_yandex_cloud_priv_iam_v1_api_key_proto = out.File
	file_yandex_cloud_priv_iam_v1_api_key_proto_rawDesc = nil
	file_yandex_cloud_priv_iam_v1_api_key_proto_goTypes = nil
	file_yandex_cloud_priv_iam_v1_api_key_proto_depIdxs = nil
}
