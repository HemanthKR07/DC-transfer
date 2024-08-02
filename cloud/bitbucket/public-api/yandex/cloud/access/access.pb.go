// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.3
// source: yandex/cloud/access/access.proto

package access

import (
	_ "github.com/doublecloud/tross/cloud/bitbucket/public-api/yandex/cloud"
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

type AccessBindingAction int32

const (
	AccessBindingAction_ACCESS_BINDING_ACTION_UNSPECIFIED AccessBindingAction = 0
	// Addition of an access binding.
	AccessBindingAction_ADD AccessBindingAction = 1
	// Removal of an access binding.
	AccessBindingAction_REMOVE AccessBindingAction = 2
)

// Enum value maps for AccessBindingAction.
var (
	AccessBindingAction_name = map[int32]string{
		0: "ACCESS_BINDING_ACTION_UNSPECIFIED",
		1: "ADD",
		2: "REMOVE",
	}
	AccessBindingAction_value = map[string]int32{
		"ACCESS_BINDING_ACTION_UNSPECIFIED": 0,
		"ADD":                               1,
		"REMOVE":                            2,
	}
)

func (x AccessBindingAction) Enum() *AccessBindingAction {
	p := new(AccessBindingAction)
	*p = x
	return p
}

func (x AccessBindingAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccessBindingAction) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_access_access_proto_enumTypes[0].Descriptor()
}

func (AccessBindingAction) Type() protoreflect.EnumType {
	return &file_yandex_cloud_access_access_proto_enumTypes[0]
}

func (x AccessBindingAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccessBindingAction.Descriptor instead.
func (AccessBindingAction) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_access_access_proto_rawDescGZIP(), []int{0}
}

type Subject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the subject.
	//
	// It can contain one of the following values:
	//   - `allAuthenticatedUsers`: A special system identifier that represents anyone
	//     who is authenticated. It can be used only if the [type] is `system`.
	//   - `allUsers`: A special system identifier that represents anyone. No authentication is required.
	//     For example, you don't need to specify the IAM token in an API query.
	//   - `<cloud generated id>`: An identifier that represents a user account.
	//     It can be used only if the [type] is `userAccount`, `federatedUser` or `serviceAccount`.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Type of the subject.
	//
	// It can contain one of the following values:
	// * `userAccount`: An account on Yandex or Yandex Connect, added to Yandex Cloud.
	// * `serviceAccount`: A service account. This type represents the [yandex.cloud.iam.v1.ServiceAccount] resource.
	// * `federatedUser`: A federated account. This type represents a user from an identity federation, like Active Directory.
	// * `system`: System group. This type represents several accounts with a common system identifier.
	//
	// For more information, see [Subject to which the role is assigned](/docs/iam/concepts/access-control/#subject).
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Subject) Reset() {
	*x = Subject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_access_access_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subject) ProtoMessage() {}

func (x *Subject) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_access_access_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subject.ProtoReflect.Descriptor instead.
func (*Subject) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_access_access_proto_rawDescGZIP(), []int{0}
}

func (x *Subject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Subject) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type AccessBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the [yandex.cloud.iam.v1.Role] that is assigned to the [subject].
	RoleId string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	// Identity for which access binding is being created.
	// It can represent an account with a unique ID or several accounts with a system identifier.
	Subject *Subject `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *AccessBinding) Reset() {
	*x = AccessBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_access_access_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessBinding) ProtoMessage() {}

func (x *AccessBinding) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_access_access_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessBinding.ProtoReflect.Descriptor instead.
func (*AccessBinding) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_access_access_proto_rawDescGZIP(), []int{1}
}

func (x *AccessBinding) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *AccessBinding) GetSubject() *Subject {
	if x != nil {
		return x.Subject
	}
	return nil
}

type ListAccessBindingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the resource to list access bindings for.
	//
	// To get the resource ID, use a corresponding List request.
	// For example, use the [yandex.cloud.resourcemanager.v1.CloudService.List] request to get the Cloud resource ID.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// The maximum number of results per page that should be returned. If the number of available
	// results is larger than [page_size],
	// the service returns a [ListAccessBindingsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. Set [page_token]
	// to the [ListAccessBindingsResponse.next_page_token]
	// returned by a previous list request to get the next page of results.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListAccessBindingsRequest) Reset() {
	*x = ListAccessBindingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_access_access_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccessBindingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccessBindingsRequest) ProtoMessage() {}

func (x *ListAccessBindingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_access_access_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccessBindingsRequest.ProtoReflect.Descriptor instead.
func (*ListAccessBindingsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_access_access_proto_rawDescGZIP(), []int{2}
}

func (x *ListAccessBindingsRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *ListAccessBindingsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListAccessBindingsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListAccessBindingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of access bindings for the specified resource.
	AccessBindings []*AccessBinding `protobuf:"bytes,1,rep,name=access_bindings,json=accessBindings,proto3" json:"access_bindings,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListAccessBindingsRequest.page_size], use
	// the [next_page_token] as the value
	// for the [ListAccessBindingsRequest.page_token] query parameter
	// in the next list request. Each subsequent list request will have its own
	// [next_page_token] to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListAccessBindingsResponse) Reset() {
	*x = ListAccessBindingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_access_access_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccessBindingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccessBindingsResponse) ProtoMessage() {}

func (x *ListAccessBindingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_access_access_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccessBindingsResponse.ProtoReflect.Descriptor instead.
func (*ListAccessBindingsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_access_access_proto_rawDescGZIP(), []int{3}
}

func (x *ListAccessBindingsResponse) GetAccessBindings() []*AccessBinding {
	if x != nil {
		return x.AccessBindings
	}
	return nil
}

func (x *ListAccessBindingsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type SetAccessBindingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the resource for which access bindings are being set.
	//
	// To get the resource ID, use a corresponding List request.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// Access bindings to be set. For more information, see [Access Bindings](/docs/iam/concepts/access-control/#access-bindings).
	AccessBindings []*AccessBinding `protobuf:"bytes,2,rep,name=access_bindings,json=accessBindings,proto3" json:"access_bindings,omitempty"`
}

func (x *SetAccessBindingsRequest) Reset() {
	*x = SetAccessBindingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_access_access_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAccessBindingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAccessBindingsRequest) ProtoMessage() {}

func (x *SetAccessBindingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_access_access_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAccessBindingsRequest.ProtoReflect.Descriptor instead.
func (*SetAccessBindingsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_access_access_proto_rawDescGZIP(), []int{4}
}

func (x *SetAccessBindingsRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *SetAccessBindingsRequest) GetAccessBindings() []*AccessBinding {
	if x != nil {
		return x.AccessBindings
	}
	return nil
}

type SetAccessBindingsMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the resource for which access bindings are being set.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *SetAccessBindingsMetadata) Reset() {
	*x = SetAccessBindingsMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_access_access_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAccessBindingsMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAccessBindingsMetadata) ProtoMessage() {}

func (x *SetAccessBindingsMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_access_access_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAccessBindingsMetadata.ProtoReflect.Descriptor instead.
func (*SetAccessBindingsMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_access_access_proto_rawDescGZIP(), []int{5}
}

func (x *SetAccessBindingsMetadata) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type UpdateAccessBindingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the resource for which access bindings are being updated.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// Updates to access bindings.
	AccessBindingDeltas []*AccessBindingDelta `protobuf:"bytes,2,rep,name=access_binding_deltas,json=accessBindingDeltas,proto3" json:"access_binding_deltas,omitempty"`
}

func (x *UpdateAccessBindingsRequest) Reset() {
	*x = UpdateAccessBindingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_access_access_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccessBindingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccessBindingsRequest) ProtoMessage() {}

func (x *UpdateAccessBindingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_access_access_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccessBindingsRequest.ProtoReflect.Descriptor instead.
func (*UpdateAccessBindingsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_access_access_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateAccessBindingsRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *UpdateAccessBindingsRequest) GetAccessBindingDeltas() []*AccessBindingDelta {
	if x != nil {
		return x.AccessBindingDeltas
	}
	return nil
}

type UpdateAccessBindingsMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the resource for which access bindings are being updated.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *UpdateAccessBindingsMetadata) Reset() {
	*x = UpdateAccessBindingsMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_access_access_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccessBindingsMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccessBindingsMetadata) ProtoMessage() {}

func (x *UpdateAccessBindingsMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_access_access_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccessBindingsMetadata.ProtoReflect.Descriptor instead.
func (*UpdateAccessBindingsMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_access_access_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateAccessBindingsMetadata) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type AccessBindingDelta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The action that is being performed on an access binding.
	Action AccessBindingAction `protobuf:"varint,1,opt,name=action,proto3,enum=yandex.cloud.access.AccessBindingAction" json:"action,omitempty"`
	// Access binding. For more information, see [Access Bindings](/docs/iam/concepts/access-control/#access-bindings).
	AccessBinding *AccessBinding `protobuf:"bytes,2,opt,name=access_binding,json=accessBinding,proto3" json:"access_binding,omitempty"`
}

func (x *AccessBindingDelta) Reset() {
	*x = AccessBindingDelta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_access_access_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessBindingDelta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessBindingDelta) ProtoMessage() {}

func (x *AccessBindingDelta) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_access_access_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessBindingDelta.ProtoReflect.Descriptor instead.
func (*AccessBindingDelta) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_access_access_proto_rawDescGZIP(), []int{8}
}

func (x *AccessBindingDelta) GetAction() AccessBindingAction {
	if x != nil {
		return x.Action
	}
	return AccessBindingAction_ACCESS_BINDING_ACTION_UNSPECIFIED
}

func (x *AccessBindingDelta) GetAccessBinding() *AccessBinding {
	if x != nil {
		return x.AccessBinding
	}
	return nil
}

type AccessBindingsOperationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Result access binding deltas.
	EffectiveDeltas []*AccessBindingDelta `protobuf:"bytes,1,rep,name=effective_deltas,json=effectiveDeltas,proto3" json:"effective_deltas,omitempty"`
}

func (x *AccessBindingsOperationResult) Reset() {
	*x = AccessBindingsOperationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_access_access_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessBindingsOperationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessBindingsOperationResult) ProtoMessage() {}

func (x *AccessBindingsOperationResult) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_access_access_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessBindingsOperationResult.ProtoReflect.Descriptor instead.
func (*AccessBindingsOperationResult) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_access_access_proto_rawDescGZIP(), []int{9}
}

func (x *AccessBindingsOperationResult) GetEffectiveDeltas() []*AccessBindingDelta {
	if x != nil {
		return x.EffectiveDeltas
	}
	return nil
}

var File_yandex_cloud_access_access_proto protoreflect.FileDescriptor

var file_yandex_cloud_access_access_proto_rawDesc = []byte{
	0x0a, 0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x1d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xe8,
	0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x21, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d,
	0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x74, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c,
	0x3d, 0x35, 0x30, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01,
	0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x9d, 0x01, 0x0a, 0x19, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7,
	0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0xfa, 0xc7, 0x31, 0x06, 0x3c,
	0x3d, 0x31, 0x30, 0x30, 0x30, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x28, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x91, 0x01, 0x0a, 0x1a, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa2, 0x01,
	0x0a, 0x18, 0x53, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x57, 0x0a, 0x0f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x0a, 0x82, 0xc8, 0x31, 0x06, 0x3c, 0x3d, 0x31, 0x30,
	0x30, 0x30, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x22, 0x3c, 0x0a, 0x19, 0x53, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x22, 0xb5, 0x01, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c,
	0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x67, 0x0a, 0x15, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x42, 0x0a, 0x82, 0xc8, 0x31, 0x06, 0x31, 0x2d, 0x31,
	0x30, 0x30, 0x30, 0x52, 0x13, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x73, 0x22, 0x3f, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0xad, 0x01, 0x0a, 0x12, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x74, 0x61,
	0x12, 0x46, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x73, 0x0a, 0x1d, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x52, 0x0a, 0x10, 0x65, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x52, 0x0f, 0x65,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x73, 0x2a, 0x51,
	0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x21, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f,
	0x42, 0x49, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x41, 0x44, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10,
	0x02, 0x42, 0x61, 0x0a, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5a, 0x46, 0x61, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x72, 0x75, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x3b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_access_access_proto_rawDescOnce sync.Once
	file_yandex_cloud_access_access_proto_rawDescData = file_yandex_cloud_access_access_proto_rawDesc
)

func file_yandex_cloud_access_access_proto_rawDescGZIP() []byte {
	file_yandex_cloud_access_access_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_access_access_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_access_access_proto_rawDescData)
	})
	return file_yandex_cloud_access_access_proto_rawDescData
}

var file_yandex_cloud_access_access_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_access_access_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_yandex_cloud_access_access_proto_goTypes = []interface{}{
	(AccessBindingAction)(0),              // 0: yandex.cloud.access.AccessBindingAction
	(*Subject)(nil),                       // 1: yandex.cloud.access.Subject
	(*AccessBinding)(nil),                 // 2: yandex.cloud.access.AccessBinding
	(*ListAccessBindingsRequest)(nil),     // 3: yandex.cloud.access.ListAccessBindingsRequest
	(*ListAccessBindingsResponse)(nil),    // 4: yandex.cloud.access.ListAccessBindingsResponse
	(*SetAccessBindingsRequest)(nil),      // 5: yandex.cloud.access.SetAccessBindingsRequest
	(*SetAccessBindingsMetadata)(nil),     // 6: yandex.cloud.access.SetAccessBindingsMetadata
	(*UpdateAccessBindingsRequest)(nil),   // 7: yandex.cloud.access.UpdateAccessBindingsRequest
	(*UpdateAccessBindingsMetadata)(nil),  // 8: yandex.cloud.access.UpdateAccessBindingsMetadata
	(*AccessBindingDelta)(nil),            // 9: yandex.cloud.access.AccessBindingDelta
	(*AccessBindingsOperationResult)(nil), // 10: yandex.cloud.access.AccessBindingsOperationResult
}
var file_yandex_cloud_access_access_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.access.AccessBinding.subject:type_name -> yandex.cloud.access.Subject
	2, // 1: yandex.cloud.access.ListAccessBindingsResponse.access_bindings:type_name -> yandex.cloud.access.AccessBinding
	2, // 2: yandex.cloud.access.SetAccessBindingsRequest.access_bindings:type_name -> yandex.cloud.access.AccessBinding
	9, // 3: yandex.cloud.access.UpdateAccessBindingsRequest.access_binding_deltas:type_name -> yandex.cloud.access.AccessBindingDelta
	0, // 4: yandex.cloud.access.AccessBindingDelta.action:type_name -> yandex.cloud.access.AccessBindingAction
	2, // 5: yandex.cloud.access.AccessBindingDelta.access_binding:type_name -> yandex.cloud.access.AccessBinding
	9, // 6: yandex.cloud.access.AccessBindingsOperationResult.effective_deltas:type_name -> yandex.cloud.access.AccessBindingDelta
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_yandex_cloud_access_access_proto_init() }
func file_yandex_cloud_access_access_proto_init() {
	if File_yandex_cloud_access_access_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_access_access_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subject); i {
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
		file_yandex_cloud_access_access_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessBinding); i {
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
		file_yandex_cloud_access_access_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccessBindingsRequest); i {
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
		file_yandex_cloud_access_access_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccessBindingsResponse); i {
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
		file_yandex_cloud_access_access_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAccessBindingsRequest); i {
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
		file_yandex_cloud_access_access_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAccessBindingsMetadata); i {
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
		file_yandex_cloud_access_access_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccessBindingsRequest); i {
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
		file_yandex_cloud_access_access_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccessBindingsMetadata); i {
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
		file_yandex_cloud_access_access_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessBindingDelta); i {
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
		file_yandex_cloud_access_access_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessBindingsOperationResult); i {
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
			RawDescriptor: file_yandex_cloud_access_access_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_access_access_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_access_access_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_access_access_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_access_access_proto_msgTypes,
	}.Build()
	File_yandex_cloud_access_access_proto = out.File
	file_yandex_cloud_access_access_proto_rawDesc = nil
	file_yandex_cloud_access_access_proto_goTypes = nil
	file_yandex_cloud_access_access_proto_depIdxs = nil
}
