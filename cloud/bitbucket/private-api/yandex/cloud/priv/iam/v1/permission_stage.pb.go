// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.3
// source: yandex/cloud/priv/iam/v1/permission_stage.proto

package iam

import (
	_ "github.com/doublecloud/tross/cloud/bitbucket/private-api/yandex/cloud/priv"
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

type PermissionStage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *PermissionStage) Reset() {
	*x = PermissionStage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionStage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionStage) ProtoMessage() {}

func (x *PermissionStage) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionStage.ProtoReflect.Descriptor instead.
func (*PermissionStage) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_v1_permission_stage_proto_rawDescGZIP(), []int{0}
}

func (x *PermissionStage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PermissionStage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type SetAllPermissionStagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *SetAllPermissionStagesRequest) Reset() {
	*x = SetAllPermissionStagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAllPermissionStagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAllPermissionStagesRequest) ProtoMessage() {}

func (x *SetAllPermissionStagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAllPermissionStagesRequest.ProtoReflect.Descriptor instead.
func (*SetAllPermissionStagesRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_v1_permission_stage_proto_rawDescGZIP(), []int{1}
}

func (x *SetAllPermissionStagesRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type SetPermissionStagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId         string   `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	PermissionStageIds []string `protobuf:"bytes,2,rep,name=permission_stage_ids,json=permissionStageIds,proto3" json:"permission_stage_ids,omitempty"`
}

func (x *SetPermissionStagesRequest) Reset() {
	*x = SetPermissionStagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPermissionStagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPermissionStagesRequest) ProtoMessage() {}

func (x *SetPermissionStagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPermissionStagesRequest.ProtoReflect.Descriptor instead.
func (*SetPermissionStagesRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_v1_permission_stage_proto_rawDescGZIP(), []int{2}
}

func (x *SetPermissionStagesRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *SetPermissionStagesRequest) GetPermissionStageIds() []string {
	if x != nil {
		return x.PermissionStageIds
	}
	return nil
}

type SetPermissionStagesMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *SetPermissionStagesMetadata) Reset() {
	*x = SetPermissionStagesMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPermissionStagesMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPermissionStagesMetadata) ProtoMessage() {}

func (x *SetPermissionStagesMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPermissionStagesMetadata.ProtoReflect.Descriptor instead.
func (*SetPermissionStagesMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_v1_permission_stage_proto_rawDescGZIP(), []int{3}
}

func (x *SetPermissionStagesMetadata) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type UpdatePermissionStagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId               string   `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	AddPermissionStageIds    []string `protobuf:"bytes,2,rep,name=add_permission_stage_ids,json=addPermissionStageIds,proto3" json:"add_permission_stage_ids,omitempty"`
	RemovePermissionStageIds []string `protobuf:"bytes,3,rep,name=remove_permission_stage_ids,json=removePermissionStageIds,proto3" json:"remove_permission_stage_ids,omitempty"`
}

func (x *UpdatePermissionStagesRequest) Reset() {
	*x = UpdatePermissionStagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePermissionStagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePermissionStagesRequest) ProtoMessage() {}

func (x *UpdatePermissionStagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePermissionStagesRequest.ProtoReflect.Descriptor instead.
func (*UpdatePermissionStagesRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_v1_permission_stage_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePermissionStagesRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *UpdatePermissionStagesRequest) GetAddPermissionStageIds() []string {
	if x != nil {
		return x.AddPermissionStageIds
	}
	return nil
}

func (x *UpdatePermissionStagesRequest) GetRemovePermissionStageIds() []string {
	if x != nil {
		return x.RemovePermissionStageIds
	}
	return nil
}

type UpdatePermissionStagesMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *UpdatePermissionStagesMetadata) Reset() {
	*x = UpdatePermissionStagesMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePermissionStagesMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePermissionStagesMetadata) ProtoMessage() {}

func (x *UpdatePermissionStagesMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePermissionStagesMetadata.ProtoReflect.Descriptor instead.
func (*UpdatePermissionStagesMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_priv_iam_v1_permission_stage_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePermissionStagesMetadata) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

var File_yandex_cloud_priv_iam_v1_permission_stage_proto protoreflect.FileDescriptor

var file_yandex_cloud_priv_iam_v1_permission_stage_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x72, 0x69, 0x76, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x18, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x72, 0x69, 0x76, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x43, 0x0a, 0x0f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4e, 0x0a, 0x1d, 0x53, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xa8, 0x89, 0x31, 0x01,
	0xca, 0x89, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x22, 0x7d, 0x0a, 0x1a, 0x53, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xa8, 0x89, 0x31, 0x01, 0xca, 0x89, 0x31,
	0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x12, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x73, 0x22, 0x3e, 0x0a, 0x1b, 0x53, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x22, 0xc6, 0x01, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xa8, 0x89, 0x31, 0x01,
	0xca, 0x89, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x18, 0x61, 0x64, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x15, 0x61, 0x64, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x73, 0x12, 0x3d, 0x0a,
	0x1b, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x18, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x73, 0x22, 0x41, 0x0a, 0x1e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x42,
	0x50, 0x42, 0x03, 0x50, 0x50, 0x53, 0x5a, 0x49, 0x61, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x72, 0x75, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x61,
	0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_priv_iam_v1_permission_stage_proto_rawDescOnce sync.Once
	file_yandex_cloud_priv_iam_v1_permission_stage_proto_rawDescData = file_yandex_cloud_priv_iam_v1_permission_stage_proto_rawDesc
)

func file_yandex_cloud_priv_iam_v1_permission_stage_proto_rawDescGZIP() []byte {
	file_yandex_cloud_priv_iam_v1_permission_stage_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_priv_iam_v1_permission_stage_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_priv_iam_v1_permission_stage_proto_rawDescData)
	})
	return file_yandex_cloud_priv_iam_v1_permission_stage_proto_rawDescData
}

var file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_yandex_cloud_priv_iam_v1_permission_stage_proto_goTypes = []interface{}{
	(*PermissionStage)(nil),                // 0: yandex.cloud.priv.iam.v1.PermissionStage
	(*SetAllPermissionStagesRequest)(nil),  // 1: yandex.cloud.priv.iam.v1.SetAllPermissionStagesRequest
	(*SetPermissionStagesRequest)(nil),     // 2: yandex.cloud.priv.iam.v1.SetPermissionStagesRequest
	(*SetPermissionStagesMetadata)(nil),    // 3: yandex.cloud.priv.iam.v1.SetPermissionStagesMetadata
	(*UpdatePermissionStagesRequest)(nil),  // 4: yandex.cloud.priv.iam.v1.UpdatePermissionStagesRequest
	(*UpdatePermissionStagesMetadata)(nil), // 5: yandex.cloud.priv.iam.v1.UpdatePermissionStagesMetadata
}
var file_yandex_cloud_priv_iam_v1_permission_stage_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_priv_iam_v1_permission_stage_proto_init() }
func file_yandex_cloud_priv_iam_v1_permission_stage_proto_init() {
	if File_yandex_cloud_priv_iam_v1_permission_stage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionStage); i {
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
		file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAllPermissionStagesRequest); i {
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
		file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPermissionStagesRequest); i {
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
		file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPermissionStagesMetadata); i {
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
		file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePermissionStagesRequest); i {
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
		file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePermissionStagesMetadata); i {
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
			RawDescriptor: file_yandex_cloud_priv_iam_v1_permission_stage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_priv_iam_v1_permission_stage_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_priv_iam_v1_permission_stage_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_priv_iam_v1_permission_stage_proto_msgTypes,
	}.Build()
	File_yandex_cloud_priv_iam_v1_permission_stage_proto = out.File
	file_yandex_cloud_priv_iam_v1_permission_stage_proto_rawDesc = nil
	file_yandex_cloud_priv_iam_v1_permission_stage_proto_goTypes = nil
	file_yandex_cloud_priv_iam_v1_permission_stage_proto_depIdxs = nil
}
