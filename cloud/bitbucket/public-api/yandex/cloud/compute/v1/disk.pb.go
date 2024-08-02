// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.3
// source: yandex/cloud/compute/v1/disk.proto

package compute

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

type Disk_Status int32

const (
	Disk_STATUS_UNSPECIFIED Disk_Status = 0
	// Disk is being created.
	Disk_CREATING Disk_Status = 1
	// Disk is ready to use.
	Disk_READY Disk_Status = 2
	// Disk encountered a problem and cannot operate.
	Disk_ERROR Disk_Status = 3
	// Disk is being deleted.
	Disk_DELETING Disk_Status = 4
)

// Enum value maps for Disk_Status.
var (
	Disk_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "CREATING",
		2: "READY",
		3: "ERROR",
		4: "DELETING",
	}
	Disk_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATING":           1,
		"READY":              2,
		"ERROR":              3,
		"DELETING":           4,
	}
)

func (x Disk_Status) Enum() *Disk_Status {
	p := new(Disk_Status)
	*p = x
	return p
}

func (x Disk_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Disk_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_compute_v1_disk_proto_enumTypes[0].Descriptor()
}

func (Disk_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_compute_v1_disk_proto_enumTypes[0]
}

func (x Disk_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Disk_Status.Descriptor instead.
func (Disk_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_compute_v1_disk_proto_rawDescGZIP(), []int{0, 0}
}

// A Disk resource. For more information, see [Disks](/docs/compute/concepts/disk).
type Disk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the disk.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the disk belongs to.
	FolderId  string                 `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the disk. 1-63 characters long.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the disk. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `key:value` pairs. Maximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ID of the disk type.
	TypeId string `protobuf:"bytes,7,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	// ID of the availability zone where the disk resides.
	ZoneId string `protobuf:"bytes,8,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	// Size of the disk, specified in bytes.
	Size int64 `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	// Block size of the disk, specified in bytes.
	BlockSize int64 `protobuf:"varint,15,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
	// License IDs that indicate which licenses are attached to this resource.
	// License IDs are used to calculate additional charges for the use of the virtual machine.
	//
	// The correct license ID is generated by the platform. IDs are inherited by new resources created from this resource.
	//
	// If you know the license IDs, specify them when you create the image.
	// For example, if you create a disk image using a third-party utility and load it into Object Storage, the license IDs will be lost.
	// You can specify them in the [yandex.cloud.compute.v1.ImageService.Create] request.
	ProductIds []string `protobuf:"bytes,10,rep,name=product_ids,json=productIds,proto3" json:"product_ids,omitempty"`
	// Current status of the disk.
	Status Disk_Status `protobuf:"varint,11,opt,name=status,proto3,enum=yandex.cloud.compute.v1.Disk_Status" json:"status,omitempty"`
	// Types that are assignable to Source:
	//
	//	*Disk_SourceImageId
	//	*Disk_SourceSnapshotId
	Source isDisk_Source `protobuf_oneof:"source"`
	// Array of instances to which the disk is attached.
	InstanceIds []string `protobuf:"bytes,14,rep,name=instance_ids,json=instanceIds,proto3" json:"instance_ids,omitempty"`
	// Placement policy configuration.
	DiskPlacementPolicy *DiskPlacementPolicy `protobuf:"bytes,16,opt,name=disk_placement_policy,json=diskPlacementPolicy,proto3" json:"disk_placement_policy,omitempty"`
}

func (x *Disk) Reset() {
	*x = Disk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_compute_v1_disk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Disk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Disk) ProtoMessage() {}

func (x *Disk) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_compute_v1_disk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Disk.ProtoReflect.Descriptor instead.
func (*Disk) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_compute_v1_disk_proto_rawDescGZIP(), []int{0}
}

func (x *Disk) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Disk) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Disk) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Disk) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Disk) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Disk) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Disk) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

func (x *Disk) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *Disk) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Disk) GetBlockSize() int64 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

func (x *Disk) GetProductIds() []string {
	if x != nil {
		return x.ProductIds
	}
	return nil
}

func (x *Disk) GetStatus() Disk_Status {
	if x != nil {
		return x.Status
	}
	return Disk_STATUS_UNSPECIFIED
}

func (m *Disk) GetSource() isDisk_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *Disk) GetSourceImageId() string {
	if x, ok := x.GetSource().(*Disk_SourceImageId); ok {
		return x.SourceImageId
	}
	return ""
}

func (x *Disk) GetSourceSnapshotId() string {
	if x, ok := x.GetSource().(*Disk_SourceSnapshotId); ok {
		return x.SourceSnapshotId
	}
	return ""
}

func (x *Disk) GetInstanceIds() []string {
	if x != nil {
		return x.InstanceIds
	}
	return nil
}

func (x *Disk) GetDiskPlacementPolicy() *DiskPlacementPolicy {
	if x != nil {
		return x.DiskPlacementPolicy
	}
	return nil
}

type isDisk_Source interface {
	isDisk_Source()
}

type Disk_SourceImageId struct {
	// ID of the image that was used for disk creation.
	SourceImageId string `protobuf:"bytes,12,opt,name=source_image_id,json=sourceImageId,proto3,oneof"`
}

type Disk_SourceSnapshotId struct {
	// ID of the snapshot that was used for disk creation.
	SourceSnapshotId string `protobuf:"bytes,13,opt,name=source_snapshot_id,json=sourceSnapshotId,proto3,oneof"`
}

func (*Disk_SourceImageId) isDisk_Source() {}

func (*Disk_SourceSnapshotId) isDisk_Source() {}

type DiskPlacementPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Placement group ID.
	PlacementGroupId        string `protobuf:"bytes,1,opt,name=placement_group_id,json=placementGroupId,proto3" json:"placement_group_id,omitempty"`
	PlacementGroupPartition int64  `protobuf:"varint,2,opt,name=placement_group_partition,json=placementGroupPartition,proto3" json:"placement_group_partition,omitempty"`
}

func (x *DiskPlacementPolicy) Reset() {
	*x = DiskPlacementPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_compute_v1_disk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskPlacementPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskPlacementPolicy) ProtoMessage() {}

func (x *DiskPlacementPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_compute_v1_disk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskPlacementPolicy.ProtoReflect.Descriptor instead.
func (*DiskPlacementPolicy) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_compute_v1_disk_proto_rawDescGZIP(), []int{1}
}

func (x *DiskPlacementPolicy) GetPlacementGroupId() string {
	if x != nil {
		return x.PlacementGroupId
	}
	return ""
}

func (x *DiskPlacementPolicy) GetPlacementGroupPartition() int64 {
	if x != nil {
		return x.PlacementGroupPartition
	}
	return 0
}

type DiskPlacementPolicyChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Disk ID.
	DiskId string `protobuf:"bytes,1,opt,name=disk_id,json=diskId,proto3" json:"disk_id,omitempty"`
	// Placement policy configuration for given disk.
	DiskPlacementPolicy *DiskPlacementPolicy `protobuf:"bytes,2,opt,name=disk_placement_policy,json=diskPlacementPolicy,proto3" json:"disk_placement_policy,omitempty"`
}

func (x *DiskPlacementPolicyChange) Reset() {
	*x = DiskPlacementPolicyChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_compute_v1_disk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskPlacementPolicyChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskPlacementPolicyChange) ProtoMessage() {}

func (x *DiskPlacementPolicyChange) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_compute_v1_disk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskPlacementPolicyChange.ProtoReflect.Descriptor instead.
func (*DiskPlacementPolicyChange) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_compute_v1_disk_proto_rawDescGZIP(), []int{2}
}

func (x *DiskPlacementPolicyChange) GetDiskId() string {
	if x != nil {
		return x.DiskId
	}
	return ""
}

func (x *DiskPlacementPolicyChange) GetDiskPlacementPolicy() *DiskPlacementPolicy {
	if x != nil {
		return x.DiskPlacementPolicy
	}
	return nil
}

var File_yandex_cloud_compute_v1_disk_proto protoreflect.FileDescriptor

var file_yandex_cloud_compute_v1_disk_proto_rawDesc = []byte{
	0x0a, 0x22, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3,
	0x06, 0x0a, 0x04, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x69, 0x73, 0x6b, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x7a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x73, 0x12, 0x3c,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x0f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12, 0x60, 0x0a, 0x15, 0x64, 0x69, 0x73,
	0x6b, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x13, 0x64, 0x69, 0x73, 0x6b, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x39, 0x0a, 0x0b, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x52, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x22, 0x7f, 0x0a, 0x13, 0x44, 0x69, 0x73, 0x6b, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x19, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x96, 0x01, 0x0a, 0x19, 0x44, 0x69, 0x73, 0x6b, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x60, 0x0a, 0x15,
	0x64, 0x69, 0x73, 0x6b, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x13, 0x64, 0x69, 0x73, 0x6b, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x6a,
	0x0a, 0x1b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x5a, 0x4b, 0x61,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x72, 0x75, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f,
	0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_yandex_cloud_compute_v1_disk_proto_rawDescOnce sync.Once
	file_yandex_cloud_compute_v1_disk_proto_rawDescData = file_yandex_cloud_compute_v1_disk_proto_rawDesc
)

func file_yandex_cloud_compute_v1_disk_proto_rawDescGZIP() []byte {
	file_yandex_cloud_compute_v1_disk_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_compute_v1_disk_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_compute_v1_disk_proto_rawDescData)
	})
	return file_yandex_cloud_compute_v1_disk_proto_rawDescData
}

var file_yandex_cloud_compute_v1_disk_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_compute_v1_disk_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_compute_v1_disk_proto_goTypes = []interface{}{
	(Disk_Status)(0),                  // 0: yandex.cloud.compute.v1.Disk.Status
	(*Disk)(nil),                      // 1: yandex.cloud.compute.v1.Disk
	(*DiskPlacementPolicy)(nil),       // 2: yandex.cloud.compute.v1.DiskPlacementPolicy
	(*DiskPlacementPolicyChange)(nil), // 3: yandex.cloud.compute.v1.DiskPlacementPolicyChange
	nil,                               // 4: yandex.cloud.compute.v1.Disk.LabelsEntry
	(*timestamppb.Timestamp)(nil),     // 5: google.protobuf.Timestamp
}
var file_yandex_cloud_compute_v1_disk_proto_depIdxs = []int32{
	5, // 0: yandex.cloud.compute.v1.Disk.created_at:type_name -> google.protobuf.Timestamp
	4, // 1: yandex.cloud.compute.v1.Disk.labels:type_name -> yandex.cloud.compute.v1.Disk.LabelsEntry
	0, // 2: yandex.cloud.compute.v1.Disk.status:type_name -> yandex.cloud.compute.v1.Disk.Status
	2, // 3: yandex.cloud.compute.v1.Disk.disk_placement_policy:type_name -> yandex.cloud.compute.v1.DiskPlacementPolicy
	2, // 4: yandex.cloud.compute.v1.DiskPlacementPolicyChange.disk_placement_policy:type_name -> yandex.cloud.compute.v1.DiskPlacementPolicy
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_yandex_cloud_compute_v1_disk_proto_init() }
func file_yandex_cloud_compute_v1_disk_proto_init() {
	if File_yandex_cloud_compute_v1_disk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_compute_v1_disk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Disk); i {
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
		file_yandex_cloud_compute_v1_disk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiskPlacementPolicy); i {
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
		file_yandex_cloud_compute_v1_disk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiskPlacementPolicyChange); i {
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
	file_yandex_cloud_compute_v1_disk_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Disk_SourceImageId)(nil),
		(*Disk_SourceSnapshotId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_compute_v1_disk_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_compute_v1_disk_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_compute_v1_disk_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_compute_v1_disk_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_compute_v1_disk_proto_msgTypes,
	}.Build()
	File_yandex_cloud_compute_v1_disk_proto = out.File
	file_yandex_cloud_compute_v1_disk_proto_rawDesc = nil
	file_yandex_cloud_compute_v1_disk_proto_goTypes = nil
	file_yandex_cloud_compute_v1_disk_proto_depIdxs = nil
}
