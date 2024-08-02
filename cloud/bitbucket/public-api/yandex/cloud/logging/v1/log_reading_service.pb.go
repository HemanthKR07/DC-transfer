// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.3
// source: yandex/cloud/logging/v1/log_reading_service.proto

package logging

import (
	_ "github.com/doublecloud/tross/cloud/bitbucket/public-api/yandex/cloud"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Read selector.
	//
	// Types that are assignable to Selector:
	//
	//	*ReadRequest_PageToken
	//	*ReadRequest_Criteria
	Selector isReadRequest_Selector `protobuf_oneof:"selector"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_logging_v1_log_reading_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_logging_v1_log_reading_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_logging_v1_log_reading_service_proto_rawDescGZIP(), []int{0}
}

func (m *ReadRequest) GetSelector() isReadRequest_Selector {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (x *ReadRequest) GetPageToken() string {
	if x, ok := x.GetSelector().(*ReadRequest_PageToken); ok {
		return x.PageToken
	}
	return ""
}

func (x *ReadRequest) GetCriteria() *Criteria {
	if x, ok := x.GetSelector().(*ReadRequest_Criteria); ok {
		return x.Criteria
	}
	return nil
}

type isReadRequest_Selector interface {
	isReadRequest_Selector()
}

type ReadRequest_PageToken struct {
	// Page token. To get the next page of results, set `page_token` to the
	// [ReadResponse.next_page_token] or [ReadResponse.previous_page_token] returned by a previous read request.
	PageToken string `protobuf:"bytes,1,opt,name=page_token,json=pageToken,proto3,oneof"`
}

type ReadRequest_Criteria struct {
	// Read criteria.
	//
	// See [Criteria] for details.
	Criteria *Criteria `protobuf:"bytes,2,opt,name=criteria,proto3,oneof"`
}

func (*ReadRequest_PageToken) isReadRequest_Selector() {}

func (*ReadRequest_Criteria) isReadRequest_Selector() {}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Log group ID the read was performed from.
	LogGroupId string `protobuf:"bytes,1,opt,name=log_group_id,json=logGroupId,proto3" json:"log_group_id,omitempty"`
	// List of matching log entries.
	Entries []*LogEntry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
	// Token for getting the next page of the log entries.
	//
	// After getting log entries initially with [Criteria], you can use `next_page_token` as the value
	// for the [ReadRequest.page_token] parameter in the next read request.
	//
	// Each subsequent page will have its own `next_page_token` to continue paging through the results.
	NextPageToken string `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Token for getting the previous page of the log entries.
	//
	// After getting log entries initially with [Criteria], you can use `previous_page_token` as the value
	// for the [ReadRequest.page_token] parameter in the next read request.
	//
	// Each subsequent page will have its own `next_page_token` to continue paging through the results.
	PreviousPageToken string `protobuf:"bytes,4,opt,name=previous_page_token,json=previousPageToken,proto3" json:"previous_page_token,omitempty"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_logging_v1_log_reading_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_logging_v1_log_reading_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_logging_v1_log_reading_service_proto_rawDescGZIP(), []int{1}
}

func (x *ReadResponse) GetLogGroupId() string {
	if x != nil {
		return x.LogGroupId
	}
	return ""
}

func (x *ReadResponse) GetEntries() []*LogEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *ReadResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ReadResponse) GetPreviousPageToken() string {
	if x != nil {
		return x.PreviousPageToken
	}
	return ""
}

// Read criteria. Should be used in initial [ReadRequest].
type Criteria struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the log group to return.
	//
	// To get a log group ID make a [LogGroupService.List] request.
	LogGroupId string `protobuf:"bytes,1,opt,name=log_group_id,json=logGroupId,proto3" json:"log_group_id,omitempty"`
	// List of resource types to limit log entries to.
	//
	// Empty list disables filter.
	ResourceTypes []string `protobuf:"bytes,2,rep,name=resource_types,json=resourceTypes,proto3" json:"resource_types,omitempty"`
	// List of resource IDs to limit log entries to.
	//
	// Empty list disables filter.
	ResourceIds []string `protobuf:"bytes,3,rep,name=resource_ids,json=resourceIds,proto3" json:"resource_ids,omitempty"`
	// Lower bound of log entries timestamps.
	Since *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=since,proto3" json:"since,omitempty"`
	// Upper bound of log entries timestamps.
	Until *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=until,proto3" json:"until,omitempty"`
	// List of log levels to limit log entries to.
	//
	// Empty list disables filter.
	Levels []LogLevel_Level `protobuf:"varint,6,rep,packed,name=levels,proto3,enum=yandex.cloud.logging.v1.LogLevel_Level" json:"levels,omitempty"`
	// Filter expression. For details about filtering, see [documentation](/docs/logging/concepts/filter).
	Filter string `protobuf:"bytes,7,opt,name=filter,proto3" json:"filter,omitempty"`
	// List of stream names to limit log entries to.
	//
	// Empty list disables filter.
	StreamNames []string `protobuf:"bytes,10,rep,name=stream_names,json=streamNames,proto3" json:"stream_names,omitempty"`
	// The maximum number of results per page to return.
	PageSize int64 `protobuf:"varint,8,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Limits response to maximum size in bytes. Prevents gRPC resource exhaustion.
	//
	// Default value for max response size is 3.5 MiB
	MaxResponseSize int64 `protobuf:"varint,9,opt,name=max_response_size,json=maxResponseSize,proto3" json:"max_response_size,omitempty"`
}

func (x *Criteria) Reset() {
	*x = Criteria{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_logging_v1_log_reading_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Criteria) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Criteria) ProtoMessage() {}

func (x *Criteria) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_logging_v1_log_reading_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Criteria.ProtoReflect.Descriptor instead.
func (*Criteria) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_logging_v1_log_reading_service_proto_rawDescGZIP(), []int{2}
}

func (x *Criteria) GetLogGroupId() string {
	if x != nil {
		return x.LogGroupId
	}
	return ""
}

func (x *Criteria) GetResourceTypes() []string {
	if x != nil {
		return x.ResourceTypes
	}
	return nil
}

func (x *Criteria) GetResourceIds() []string {
	if x != nil {
		return x.ResourceIds
	}
	return nil
}

func (x *Criteria) GetSince() *timestamppb.Timestamp {
	if x != nil {
		return x.Since
	}
	return nil
}

func (x *Criteria) GetUntil() *timestamppb.Timestamp {
	if x != nil {
		return x.Until
	}
	return nil
}

func (x *Criteria) GetLevels() []LogLevel_Level {
	if x != nil {
		return x.Levels
	}
	return nil
}

func (x *Criteria) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *Criteria) GetStreamNames() []string {
	if x != nil {
		return x.StreamNames
	}
	return nil
}

func (x *Criteria) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *Criteria) GetMaxResponseSize() int64 {
	if x != nil {
		return x.MaxResponseSize
	}
	return 0
}

var File_yandex_cloud_logging_v1_log_reading_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_logging_v1_log_reading_service_proto_rawDesc = []byte{
	0x0a, 0x31, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x65,
	0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3f, 0x0a, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x48, 0x00, 0x52, 0x08, 0x63, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x42, 0x0a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x22, 0xc5, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75,
	0x73, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x98, 0x04, 0x0a, 0x08, 0x43,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x2e, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8,
	0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x36, 0x34, 0x52, 0x0a, 0x6c, 0x6f, 0x67,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x11, 0x82, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d,
	0x36, 0x33, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x12, 0x34, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x11, 0x82, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x31,
	0x30, 0x30, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x36, 0x33, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x75, 0x6e, 0x74,
	0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x12, 0x49, 0x0a, 0x06, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x2e, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x42, 0x08, 0x82, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x31, 0x30, 0x52, 0x06,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x12, 0x22, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x8a, 0xc8, 0x31, 0x06, 0x3c, 0x3d, 0x31, 0x30,
	0x30, 0x30, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x0c, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x11, 0x82, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x8a, 0xc8, 0x31, 0x04, 0x3c,
	0x3d, 0x36, 0x33, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x12, 0x27, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x0a, 0xfa, 0xc7, 0x31, 0x06, 0x30, 0x2d, 0x31, 0x30, 0x30, 0x30, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x3a, 0x0a, 0x11, 0x6d, 0x61, 0x78,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x0e, 0xfa, 0xc7, 0x31, 0x0a, 0x30, 0x2d, 0x31, 0x30, 0x34, 0x38,
	0x35, 0x37, 0x36, 0x30, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x32, 0x68, 0x0a, 0x11, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x04, 0x52, 0x65,
	0x61, 0x64, 0x12, 0x24, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x6a, 0x0a, 0x1b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x5a, 0x4b,
	0x61, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x72, 0x75,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_logging_v1_log_reading_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_logging_v1_log_reading_service_proto_rawDescData = file_yandex_cloud_logging_v1_log_reading_service_proto_rawDesc
)

func file_yandex_cloud_logging_v1_log_reading_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_logging_v1_log_reading_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_logging_v1_log_reading_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_logging_v1_log_reading_service_proto_rawDescData)
	})
	return file_yandex_cloud_logging_v1_log_reading_service_proto_rawDescData
}

var file_yandex_cloud_logging_v1_log_reading_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_logging_v1_log_reading_service_proto_goTypes = []interface{}{
	(*ReadRequest)(nil),           // 0: yandex.cloud.logging.v1.ReadRequest
	(*ReadResponse)(nil),          // 1: yandex.cloud.logging.v1.ReadResponse
	(*Criteria)(nil),              // 2: yandex.cloud.logging.v1.Criteria
	(*LogEntry)(nil),              // 3: yandex.cloud.logging.v1.LogEntry
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(LogLevel_Level)(0),           // 5: yandex.cloud.logging.v1.LogLevel.Level
}
var file_yandex_cloud_logging_v1_log_reading_service_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.logging.v1.ReadRequest.criteria:type_name -> yandex.cloud.logging.v1.Criteria
	3, // 1: yandex.cloud.logging.v1.ReadResponse.entries:type_name -> yandex.cloud.logging.v1.LogEntry
	4, // 2: yandex.cloud.logging.v1.Criteria.since:type_name -> google.protobuf.Timestamp
	4, // 3: yandex.cloud.logging.v1.Criteria.until:type_name -> google.protobuf.Timestamp
	5, // 4: yandex.cloud.logging.v1.Criteria.levels:type_name -> yandex.cloud.logging.v1.LogLevel.Level
	0, // 5: yandex.cloud.logging.v1.LogReadingService.Read:input_type -> yandex.cloud.logging.v1.ReadRequest
	1, // 6: yandex.cloud.logging.v1.LogReadingService.Read:output_type -> yandex.cloud.logging.v1.ReadResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_yandex_cloud_logging_v1_log_reading_service_proto_init() }
func file_yandex_cloud_logging_v1_log_reading_service_proto_init() {
	if File_yandex_cloud_logging_v1_log_reading_service_proto != nil {
		return
	}
	file_yandex_cloud_logging_v1_log_entry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_logging_v1_log_reading_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_yandex_cloud_logging_v1_log_reading_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadResponse); i {
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
		file_yandex_cloud_logging_v1_log_reading_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Criteria); i {
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
	file_yandex_cloud_logging_v1_log_reading_service_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ReadRequest_PageToken)(nil),
		(*ReadRequest_Criteria)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_logging_v1_log_reading_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_logging_v1_log_reading_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_logging_v1_log_reading_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_logging_v1_log_reading_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_logging_v1_log_reading_service_proto = out.File
	file_yandex_cloud_logging_v1_log_reading_service_proto_rawDesc = nil
	file_yandex_cloud_logging_v1_log_reading_service_proto_goTypes = nil
	file_yandex_cloud_logging_v1_log_reading_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogReadingServiceClient is the client API for LogReadingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogReadingServiceClient interface {
	// Read log entries from the specified log group.
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
}

type logReadingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogReadingServiceClient(cc grpc.ClientConnInterface) LogReadingServiceClient {
	return &logReadingServiceClient{cc}
}

func (c *logReadingServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.logging.v1.LogReadingService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogReadingServiceServer is the server API for LogReadingService service.
type LogReadingServiceServer interface {
	// Read log entries from the specified log group.
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
}

// UnimplementedLogReadingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLogReadingServiceServer struct {
}

func (*UnimplementedLogReadingServiceServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}

func RegisterLogReadingServiceServer(s *grpc.Server, srv LogReadingServiceServer) {
	s.RegisterService(&_LogReadingService_serviceDesc, srv)
}

func _LogReadingService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogReadingServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.logging.v1.LogReadingService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogReadingServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogReadingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.logging.v1.LogReadingService",
	HandlerType: (*LogReadingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _LogReadingService_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/logging/v1/log_reading_service.proto",
}
