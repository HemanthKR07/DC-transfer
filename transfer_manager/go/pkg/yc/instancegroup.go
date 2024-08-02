package yc

import (
	"context"

	"github.com/doublecloud/tross/cloud/bitbucket/public-api/yandex/cloud/compute/v1/instancegroup"
	"github.com/doublecloud/tross/cloud/bitbucket/public-api/yandex/cloud/operation"
	"github.com/doublecloud/tross/library/go/core/xerrors"
	"google.golang.org/grpc"
)

// InstanceGroupServiceClient is a instancegroup.InstanceGroupServiceClient with
// lazy GRPC connection initialization.
type InstanceGroupServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) Create(ctx context.Context, in *instancegroup.CreateInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot get connection: %w", err)
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).Create(ctx, in, opts...)
}

// CreateFromYaml implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) CreateFromYaml(ctx context.Context, in *instancegroup.CreateInstanceGroupFromYamlRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot get connection: %w", err)
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).CreateFromYaml(ctx, in, opts...)
}

// Delete implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) Delete(ctx context.Context, in *instancegroup.DeleteInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot get connection: %w", err)
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) Get(ctx context.Context, in *instancegroup.GetInstanceGroupRequest, opts ...grpc.CallOption) (*instancegroup.InstanceGroup, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot get connection: %w", err)
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).Get(ctx, in, opts...)
}

// List implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) List(ctx context.Context, in *instancegroup.ListInstanceGroupsRequest, opts ...grpc.CallOption) (*instancegroup.ListInstanceGroupsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot get connection: %w", err)
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).List(ctx, in, opts...)
}

type InstanceGroupIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *InstanceGroupServiceClient
	request *instancegroup.ListInstanceGroupsRequest

	items []*instancegroup.InstanceGroup
}

func (c *InstanceGroupServiceClient) InstanceGroupIterator(ctx context.Context, folderID string, opts ...grpc.CallOption) *InstanceGroupIterator {
	return &InstanceGroupIterator{
		ctx:     ctx,
		opts:    opts,
		err:     nil,
		started: false,
		client:  c,
		request: &instancegroup.ListInstanceGroupsRequest{
			FolderId: folderID,
			PageSize: 1000,
		},
		items: []*instancegroup.InstanceGroup{},
	}
}

func (it *InstanceGroupIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.InstanceGroups
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *InstanceGroupIterator) Value() *instancegroup.InstanceGroup {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *InstanceGroupIterator) Error() error {
	return it.err
}

// ListInstances implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) ListInstances(ctx context.Context, in *instancegroup.ListInstanceGroupInstancesRequest, opts ...grpc.CallOption) (*instancegroup.ListInstanceGroupInstancesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannt get connection: %w", err)
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).ListInstances(ctx, in, opts...)
}

type InstanceGroupInstancesIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *InstanceGroupServiceClient
	request *instancegroup.ListInstanceGroupInstancesRequest

	items []*instancegroup.ManagedInstance
}

func (c *InstanceGroupServiceClient) InstanceGroupInstancesIterator(ctx context.Context, instanceGroupID string, opts ...grpc.CallOption) *InstanceGroupInstancesIterator {
	return &InstanceGroupInstancesIterator{
		ctx:     ctx,
		opts:    opts,
		client:  c,
		err:     nil,
		started: false,
		request: &instancegroup.ListInstanceGroupInstancesRequest{
			InstanceGroupId: instanceGroupID,
			PageSize:        1000,
		},
		items: []*instancegroup.ManagedInstance{},
	}
}

func (it *InstanceGroupInstancesIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	response, err := it.client.ListInstances(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Instances
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *InstanceGroupInstancesIterator) Value() *instancegroup.ManagedInstance {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *InstanceGroupInstancesIterator) Error() error {
	return it.err
}

// ListLogRecords implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) ListLogRecords(ctx context.Context, in *instancegroup.ListInstanceGroupLogRecordsRequest, opts ...grpc.CallOption) (*instancegroup.ListInstanceGroupLogRecordsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot get connection: %w", err)
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).ListLogRecords(ctx, in, opts...)
}

type InstanceGroupLogRecordsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *InstanceGroupServiceClient
	request *instancegroup.ListInstanceGroupLogRecordsRequest

	items []*instancegroup.LogRecord
}

func (c *InstanceGroupServiceClient) InstanceGroupLogRecordsIterator(ctx context.Context, instanceGroupID string, opts ...grpc.CallOption) *InstanceGroupLogRecordsIterator {
	return &InstanceGroupLogRecordsIterator{
		ctx:     ctx,
		opts:    opts,
		err:     nil,
		started: false,
		client:  c,
		request: &instancegroup.ListInstanceGroupLogRecordsRequest{
			InstanceGroupId: instanceGroupID,
			PageSize:        1000,
		},
		items: []*instancegroup.LogRecord{},
	}
}

func (it *InstanceGroupLogRecordsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	response, err := it.client.ListLogRecords(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.LogRecords
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *InstanceGroupLogRecordsIterator) Value() *instancegroup.LogRecord {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *InstanceGroupLogRecordsIterator) Error() error {
	return it.err
}

// ListOperations implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) ListOperations(ctx context.Context, in *instancegroup.ListInstanceGroupOperationsRequest, opts ...grpc.CallOption) (*instancegroup.ListInstanceGroupOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot get connection: %w", err)
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).ListOperations(ctx, in, opts...)
}

type InstanceGroupOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *InstanceGroupServiceClient
	request *instancegroup.ListInstanceGroupOperationsRequest

	items []*operation.Operation
}

func (c *InstanceGroupServiceClient) InstanceGroupOperationsIterator(ctx context.Context, instanceGroupID string, opts ...grpc.CallOption) *InstanceGroupOperationsIterator {
	return &InstanceGroupOperationsIterator{
		ctx:     ctx,
		opts:    opts,
		err:     nil,
		started: false,
		client:  c,
		request: &instancegroup.ListInstanceGroupOperationsRequest{
			InstanceGroupId: instanceGroupID,
			PageSize:        1000,
		},
		items: nil,
	}
}

func (it *InstanceGroupOperationsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *InstanceGroupOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *InstanceGroupOperationsIterator) Error() error {
	return it.err
}

// Start implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) Start(ctx context.Context, in *instancegroup.StartInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot get connection: %w", err)
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).Start(ctx, in, opts...)
}

// Stop implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) Stop(ctx context.Context, in *instancegroup.StopInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot get connection: %w", err)
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).Stop(ctx, in, opts...)
}

// Update implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) Update(ctx context.Context, in *instancegroup.UpdateInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot get connection: %w", err)
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateFromYaml implements instancegroup.InstanceGroupServiceClient
func (c *InstanceGroupServiceClient) UpdateFromYaml(ctx context.Context, in *instancegroup.UpdateInstanceGroupFromYamlRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot get connection: %w", err)
	}
	return instancegroup.NewInstanceGroupServiceClient(conn).UpdateFromYaml(ctx, in, opts...)
}
