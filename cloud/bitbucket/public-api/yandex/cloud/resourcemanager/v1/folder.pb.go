// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.3
// source: yandex/cloud/resourcemanager/v1/folder.proto

package resourcemanager

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

type Folder_Status int32

const (
	Folder_STATUS_UNSPECIFIED Folder_Status = 0
	// The folder is active.
	Folder_ACTIVE Folder_Status = 1
	// The folder is being deleted.
	Folder_DELETING Folder_Status = 2
	// Stopping folder resources and waiting for the deletion start timestamp.
	Folder_PENDING_DELETION Folder_Status = 3
)

// Enum value maps for Folder_Status.
var (
	Folder_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "ACTIVE",
		2: "DELETING",
		3: "PENDING_DELETION",
	}
	Folder_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"ACTIVE":             1,
		"DELETING":           2,
		"PENDING_DELETION":   3,
	}
)

func (x Folder_Status) Enum() *Folder_Status {
	p := new(Folder_Status)
	*p = x
	return p
}

func (x Folder_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Folder_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_resourcemanager_v1_folder_proto_enumTypes[0].Descriptor()
}

func (Folder_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_resourcemanager_v1_folder_proto_enumTypes[0]
}

func (x Folder_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Folder_Status.Descriptor instead.
func (Folder_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_resourcemanager_v1_folder_proto_rawDescGZIP(), []int{0, 0}
}

// A Folder resource. For more information, see [Folder](/docs/resource-manager/concepts/resources-hierarchy#folder).
type Folder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the folder.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the cloud that the folder belongs to.
	CloudId string `protobuf:"bytes,2,opt,name=cloud_id,json=cloudId,proto3" json:"cloud_id,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the folder.
	// The name is unique within the cloud. 3-63 characters long.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the folder. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as “ key:value “ pairs. Maximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Status of the folder.
	Status Folder_Status `protobuf:"varint,7,opt,name=status,proto3,enum=yandex.cloud.resourcemanager.v1.Folder_Status" json:"status,omitempty"`
}

func (x *Folder) Reset() {
	*x = Folder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_resourcemanager_v1_folder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Folder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Folder) ProtoMessage() {}

func (x *Folder) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_resourcemanager_v1_folder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Folder.ProtoReflect.Descriptor instead.
func (*Folder) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_resourcemanager_v1_folder_proto_rawDescGZIP(), []int{0}
}

func (x *Folder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Folder) GetCloudId() string {
	if x != nil {
		return x.CloudId
	}
	return ""
}

func (x *Folder) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Folder) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Folder) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Folder) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Folder) GetStatus() Folder_Status {
	if x != nil {
		return x.Status
	}
	return Folder_STATUS_UNSPECIFIED
}

var File_yandex_cloud_resourcemanager_v1_folder_proto protoreflect.FileDescriptor

var file_yandex_cloud_resourcemanager_v1_folder_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc6, 0x03, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x46, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x39, 0x0a, 0x0b,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x50, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x42, 0x82, 0x01, 0x0a, 0x23, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x5a, 0x5b, 0x61, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x74, 0x65, 0x61, 0x6d,
	0x2e, 0x72, 0x75, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69, 0x74, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_resourcemanager_v1_folder_proto_rawDescOnce sync.Once
	file_yandex_cloud_resourcemanager_v1_folder_proto_rawDescData = file_yandex_cloud_resourcemanager_v1_folder_proto_rawDesc
)

func file_yandex_cloud_resourcemanager_v1_folder_proto_rawDescGZIP() []byte {
	file_yandex_cloud_resourcemanager_v1_folder_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_resourcemanager_v1_folder_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_resourcemanager_v1_folder_proto_rawDescData)
	})
	return file_yandex_cloud_resourcemanager_v1_folder_proto_rawDescData
}

var file_yandex_cloud_resourcemanager_v1_folder_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_resourcemanager_v1_folder_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_resourcemanager_v1_folder_proto_goTypes = []interface{}{
	(Folder_Status)(0),            // 0: yandex.cloud.resourcemanager.v1.Folder.Status
	(*Folder)(nil),                // 1: yandex.cloud.resourcemanager.v1.Folder
	nil,                           // 2: yandex.cloud.resourcemanager.v1.Folder.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_yandex_cloud_resourcemanager_v1_folder_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.resourcemanager.v1.Folder.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: yandex.cloud.resourcemanager.v1.Folder.labels:type_name -> yandex.cloud.resourcemanager.v1.Folder.LabelsEntry
	0, // 2: yandex.cloud.resourcemanager.v1.Folder.status:type_name -> yandex.cloud.resourcemanager.v1.Folder.Status
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_resourcemanager_v1_folder_proto_init() }
func file_yandex_cloud_resourcemanager_v1_folder_proto_init() {
	if File_yandex_cloud_resourcemanager_v1_folder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_resourcemanager_v1_folder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Folder); i {
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
			RawDescriptor: file_yandex_cloud_resourcemanager_v1_folder_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_resourcemanager_v1_folder_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_resourcemanager_v1_folder_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_resourcemanager_v1_folder_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_resourcemanager_v1_folder_proto_msgTypes,
	}.Build()
	File_yandex_cloud_resourcemanager_v1_folder_proto = out.File
	file_yandex_cloud_resourcemanager_v1_folder_proto_rawDesc = nil
	file_yandex_cloud_resourcemanager_v1_folder_proto_goTypes = nil
	file_yandex_cloud_resourcemanager_v1_folder_proto_depIdxs = nil
}
