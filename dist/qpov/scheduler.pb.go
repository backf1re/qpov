// Code generated by protoc-gen-go.
// source: scheduler.proto
// DO NOT EDIT!

package qpov

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GetRequest struct {
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}

type GetReply struct {
	LeaseId         string `protobuf:"bytes,1,opt,name=lease_id" json:"lease_id,omitempty"`
	OrderDefinition string `protobuf:"bytes,2,opt,name=order_definition" json:"order_definition,omitempty"`
}

func (m *GetReply) Reset()         { *m = GetReply{} }
func (m *GetReply) String() string { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()    {}

type RenewRequest struct {
	LeaseId   string `protobuf:"bytes,1,opt,name=lease_id" json:"lease_id,omitempty"`
	ExtendSec int32  `protobuf:"varint,2,opt,name=extend_sec" json:"extend_sec,omitempty"`
}

func (m *RenewRequest) Reset()         { *m = RenewRequest{} }
func (m *RenewRequest) String() string { return proto.CompactTextString(m) }
func (*RenewRequest) ProtoMessage()    {}

type RenewReply struct {
	NewTimeoutSec int64 `protobuf:"varint,1,opt,name=new_timeout_sec" json:"new_timeout_sec,omitempty"`
}

func (m *RenewReply) Reset()         { *m = RenewReply{} }
func (m *RenewReply) String() string { return proto.CompactTextString(m) }
func (*RenewReply) ProtoMessage()    {}

type DoneRequest struct {
	LeaseId      string `protobuf:"bytes,1,opt,name=lease_id" json:"lease_id,omitempty"`
	Image        []byte `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Stdout       []byte `protobuf:"bytes,3,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr       []byte `protobuf:"bytes,4,opt,name=stderr,proto3" json:"stderr,omitempty"`
	JsonMetadata string `protobuf:"bytes,5,opt,name=json_metadata" json:"json_metadata,omitempty"`
}

func (m *DoneRequest) Reset()         { *m = DoneRequest{} }
func (m *DoneRequest) String() string { return proto.CompactTextString(m) }
func (*DoneRequest) ProtoMessage()    {}

type DoneReply struct {
}

func (m *DoneReply) Reset()         { *m = DoneReply{} }
func (m *DoneReply) String() string { return proto.CompactTextString(m) }
func (*DoneReply) ProtoMessage()    {}

type FailedRequest struct {
	LeaseId string `protobuf:"bytes,1,opt,name=lease_id" json:"lease_id,omitempty"`
}

func (m *FailedRequest) Reset()         { *m = FailedRequest{} }
func (m *FailedRequest) String() string { return proto.CompactTextString(m) }
func (*FailedRequest) ProtoMessage()    {}

type FailedReply struct {
}

func (m *FailedReply) Reset()         { *m = FailedReply{} }
func (m *FailedReply) String() string { return proto.CompactTextString(m) }
func (*FailedReply) ProtoMessage()    {}

type AddRequest struct {
	OrderDefinition string `protobuf:"bytes,1,opt,name=order_definition" json:"order_definition,omitempty"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}

type AddReply struct {
	OrderId string `protobuf:"bytes,1,opt,name=order_id" json:"order_id,omitempty"`
}

func (m *AddReply) Reset()         { *m = AddReply{} }
func (m *AddReply) String() string { return proto.CompactTextString(m) }
func (*AddReply) ProtoMessage()    {}

type Lease struct {
	OrderId   string             `protobuf:"bytes,1,opt,name=order_id" json:"order_id,omitempty"`
	LeaseId   string             `protobuf:"bytes,2,opt,name=lease_id" json:"lease_id,omitempty"`
	Done      bool               `protobuf:"varint,3,opt,name=done" json:"done,omitempty"`
	UserId    int64              `protobuf:"varint,4,opt,name=user_id" json:"user_id,omitempty"`
	CreatedMs int64              `protobuf:"varint,5,opt,name=created_ms" json:"created_ms,omitempty"`
	UpdatedMs int64              `protobuf:"varint,6,opt,name=updated_ms" json:"updated_ms,omitempty"`
	ExpiresMs int64              `protobuf:"varint,7,opt,name=expires_ms" json:"expires_ms,omitempty"`
	Order     *Order             `protobuf:"bytes,8,opt,name=order" json:"order,omitempty"`
	Metadata  *RenderingMetadata `protobuf:"bytes,9,opt,name=metadata" json:"metadata,omitempty"`
	Failed    bool               `protobuf:"varint,10,opt,name=failed" json:"failed,omitempty"`
}

func (m *Lease) Reset()         { *m = Lease{} }
func (m *Lease) String() string { return proto.CompactTextString(m) }
func (*Lease) ProtoMessage()    {}

func (m *Lease) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *Lease) GetMetadata() *RenderingMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type LeaseRequest struct {
	LeaseId string `protobuf:"bytes,1,opt,name=lease_id" json:"lease_id,omitempty"`
}

func (m *LeaseRequest) Reset()         { *m = LeaseRequest{} }
func (m *LeaseRequest) String() string { return proto.CompactTextString(m) }
func (*LeaseRequest) ProtoMessage()    {}

type LeaseReply struct {
	Lease *Lease `protobuf:"bytes,1,opt,name=lease" json:"lease,omitempty"`
}

func (m *LeaseReply) Reset()         { *m = LeaseReply{} }
func (m *LeaseReply) String() string { return proto.CompactTextString(m) }
func (*LeaseReply) ProtoMessage()    {}

func (m *LeaseReply) GetLease() *Lease {
	if m != nil {
		return m.Lease
	}
	return nil
}

type LeasesRequest struct {
	Done     bool `protobuf:"varint,1,opt,name=done" json:"done,omitempty"`
	Order    bool `protobuf:"varint,2,opt,name=order" json:"order,omitempty"`
	Metadata bool `protobuf:"varint,3,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *LeasesRequest) Reset()         { *m = LeasesRequest{} }
func (m *LeasesRequest) String() string { return proto.CompactTextString(m) }
func (*LeasesRequest) ProtoMessage()    {}

type LeasesReply struct {
	Lease *Lease `protobuf:"bytes,1,opt,name=lease" json:"lease,omitempty"`
}

func (m *LeasesReply) Reset()         { *m = LeasesReply{} }
func (m *LeasesReply) String() string { return proto.CompactTextString(m) }
func (*LeasesReply) ProtoMessage()    {}

func (m *LeasesReply) GetLease() *Lease {
	if m != nil {
		return m.Lease
	}
	return nil
}

type OrdersRequest struct {
	Done      bool `protobuf:"varint,1,opt,name=done" json:"done,omitempty"`
	Active    bool `protobuf:"varint,2,opt,name=active" json:"active,omitempty"`
	Unstarted bool `protobuf:"varint,3,opt,name=unstarted" json:"unstarted,omitempty"`
}

func (m *OrdersRequest) Reset()         { *m = OrdersRequest{} }
func (m *OrdersRequest) String() string { return proto.CompactTextString(m) }
func (*OrdersRequest) ProtoMessage()    {}

type OrderStat struct {
	OrderId string `protobuf:"bytes,1,opt,name=order_id" json:"order_id,omitempty"`
	Done    bool   `protobuf:"varint,2,opt,name=done" json:"done,omitempty"`
	Active  bool   `protobuf:"varint,3,opt,name=active" json:"active,omitempty"`
}

func (m *OrderStat) Reset()         { *m = OrderStat{} }
func (m *OrderStat) String() string { return proto.CompactTextString(m) }
func (*OrderStat) ProtoMessage()    {}

type OrdersReply struct {
	Order *OrderStat `protobuf:"bytes,1,opt,name=order" json:"order,omitempty"`
}

func (m *OrdersReply) Reset()         { *m = OrdersReply{} }
func (m *OrdersReply) String() string { return proto.CompactTextString(m) }
func (*OrdersReply) ProtoMessage()    {}

func (m *OrdersReply) GetOrder() *OrderStat {
	if m != nil {
		return m.Order
	}
	return nil
}

type SchedulingStats struct {
	Orders       int64 `protobuf:"varint,1,opt,name=orders" json:"orders,omitempty"`
	ActiveOrders int64 `protobuf:"varint,2,opt,name=active_orders" json:"active_orders,omitempty"`
	DoneOrders   int64 `protobuf:"varint,3,opt,name=done_orders" json:"done_orders,omitempty"`
	Leases       int64 `protobuf:"varint,4,opt,name=leases" json:"leases,omitempty"`
	ActiveLeases int64 `protobuf:"varint,5,opt,name=active_leases" json:"active_leases,omitempty"`
	DoneLeases   int64 `protobuf:"varint,6,opt,name=done_leases" json:"done_leases,omitempty"`
}

func (m *SchedulingStats) Reset()         { *m = SchedulingStats{} }
func (m *SchedulingStats) String() string { return proto.CompactTextString(m) }
func (*SchedulingStats) ProtoMessage()    {}

type StatsRequest struct {
	SchedulingStats bool `protobuf:"varint,1,opt,name=scheduling_stats" json:"scheduling_stats,omitempty"`
}

func (m *StatsRequest) Reset()         { *m = StatsRequest{} }
func (m *StatsRequest) String() string { return proto.CompactTextString(m) }
func (*StatsRequest) ProtoMessage()    {}

// Global stats.
type StatsReply struct {
	SchedulingStats *SchedulingStats `protobuf:"bytes,1,opt,name=scheduling_stats" json:"scheduling_stats,omitempty"`
}

func (m *StatsReply) Reset()         { *m = StatsReply{} }
func (m *StatsReply) String() string { return proto.CompactTextString(m) }
func (*StatsReply) ProtoMessage()    {}

func (m *StatsReply) GetSchedulingStats() *SchedulingStats {
	if m != nil {
		return m.SchedulingStats
	}
	return nil
}

type ResultRequest struct {
	LeaseId string `protobuf:"bytes,1,opt,name=lease_id" json:"lease_id,omitempty"`
	Data    bool   `protobuf:"varint,2,opt,name=data" json:"data,omitempty"`
}

func (m *ResultRequest) Reset()         { *m = ResultRequest{} }
func (m *ResultRequest) String() string { return proto.CompactTextString(m) }
func (*ResultRequest) ProtoMessage()    {}

type ResultReply struct {
	ContentType string `protobuf:"bytes,1,opt,name=content_type" json:"content_type,omitempty"`
	Data        []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *ResultReply) Reset()         { *m = ResultReply{} }
func (m *ResultReply) String() string { return proto.CompactTextString(m) }
func (*ResultReply) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Scheduler service

type SchedulerClient interface {
	// Render client API.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	Renew(ctx context.Context, in *RenewRequest, opts ...grpc.CallOption) (*RenewReply, error)
	Done(ctx context.Context, in *DoneRequest, opts ...grpc.CallOption) (*DoneReply, error)
	Failed(ctx context.Context, in *FailedRequest, opts ...grpc.CallOption) (*FailedReply, error)
	// Order handling API. Restricted.
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
	// Stats API. Restricted.
	Lease(ctx context.Context, in *LeaseRequest, opts ...grpc.CallOption) (*LeaseReply, error)
	Leases(ctx context.Context, in *LeasesRequest, opts ...grpc.CallOption) (Scheduler_LeasesClient, error)
	Orders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (Scheduler_OrdersClient, error)
	Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsReply, error)
	// WebUI magic.
	// rpc UserStats (UserStatsRequest) returns (UserStatsReply) {}
	Result(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (Scheduler_ResultClient, error)
}

type schedulerClient struct {
	cc *grpc.ClientConn
}

func NewSchedulerClient(cc *grpc.ClientConn) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := grpc.Invoke(ctx, "/qpov.Scheduler/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) Renew(ctx context.Context, in *RenewRequest, opts ...grpc.CallOption) (*RenewReply, error) {
	out := new(RenewReply)
	err := grpc.Invoke(ctx, "/qpov.Scheduler/Renew", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) Done(ctx context.Context, in *DoneRequest, opts ...grpc.CallOption) (*DoneReply, error) {
	out := new(DoneReply)
	err := grpc.Invoke(ctx, "/qpov.Scheduler/Done", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) Failed(ctx context.Context, in *FailedRequest, opts ...grpc.CallOption) (*FailedReply, error) {
	out := new(FailedReply)
	err := grpc.Invoke(ctx, "/qpov.Scheduler/Failed", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error) {
	out := new(AddReply)
	err := grpc.Invoke(ctx, "/qpov.Scheduler/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) Lease(ctx context.Context, in *LeaseRequest, opts ...grpc.CallOption) (*LeaseReply, error) {
	out := new(LeaseReply)
	err := grpc.Invoke(ctx, "/qpov.Scheduler/Lease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) Leases(ctx context.Context, in *LeasesRequest, opts ...grpc.CallOption) (Scheduler_LeasesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Scheduler_serviceDesc.Streams[0], c.cc, "/qpov.Scheduler/Leases", opts...)
	if err != nil {
		return nil, err
	}
	x := &schedulerLeasesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Scheduler_LeasesClient interface {
	Recv() (*LeasesReply, error)
	grpc.ClientStream
}

type schedulerLeasesClient struct {
	grpc.ClientStream
}

func (x *schedulerLeasesClient) Recv() (*LeasesReply, error) {
	m := new(LeasesReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *schedulerClient) Orders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (Scheduler_OrdersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Scheduler_serviceDesc.Streams[1], c.cc, "/qpov.Scheduler/Orders", opts...)
	if err != nil {
		return nil, err
	}
	x := &schedulerOrdersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Scheduler_OrdersClient interface {
	Recv() (*OrdersReply, error)
	grpc.ClientStream
}

type schedulerOrdersClient struct {
	grpc.ClientStream
}

func (x *schedulerOrdersClient) Recv() (*OrdersReply, error) {
	m := new(OrdersReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *schedulerClient) Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsReply, error) {
	out := new(StatsReply)
	err := grpc.Invoke(ctx, "/qpov.Scheduler/Stats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) Result(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (Scheduler_ResultClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Scheduler_serviceDesc.Streams[2], c.cc, "/qpov.Scheduler/Result", opts...)
	if err != nil {
		return nil, err
	}
	x := &schedulerResultClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Scheduler_ResultClient interface {
	Recv() (*ResultReply, error)
	grpc.ClientStream
}

type schedulerResultClient struct {
	grpc.ClientStream
}

func (x *schedulerResultClient) Recv() (*ResultReply, error) {
	m := new(ResultReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Scheduler service

type SchedulerServer interface {
	// Render client API.
	Get(context.Context, *GetRequest) (*GetReply, error)
	Renew(context.Context, *RenewRequest) (*RenewReply, error)
	Done(context.Context, *DoneRequest) (*DoneReply, error)
	Failed(context.Context, *FailedRequest) (*FailedReply, error)
	// Order handling API. Restricted.
	Add(context.Context, *AddRequest) (*AddReply, error)
	// Stats API. Restricted.
	Lease(context.Context, *LeaseRequest) (*LeaseReply, error)
	Leases(*LeasesRequest, Scheduler_LeasesServer) error
	Orders(*OrdersRequest, Scheduler_OrdersServer) error
	Stats(context.Context, *StatsRequest) (*StatsReply, error)
	// WebUI magic.
	// rpc UserStats (UserStatsRequest) returns (UserStatsReply) {}
	Result(*ResultRequest, Scheduler_ResultServer) error
}

func RegisterSchedulerServer(s *grpc.Server, srv SchedulerServer) {
	s.RegisterService(&_Scheduler_serviceDesc, srv)
}

func _Scheduler_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SchedulerServer).Get(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Scheduler_Renew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RenewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SchedulerServer).Renew(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Scheduler_Done_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SchedulerServer).Done(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Scheduler_Failed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(FailedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SchedulerServer).Failed(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Scheduler_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SchedulerServer).Add(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Scheduler_Lease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(LeaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SchedulerServer).Lease(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Scheduler_Leases_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LeasesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SchedulerServer).Leases(m, &schedulerLeasesServer{stream})
}

type Scheduler_LeasesServer interface {
	Send(*LeasesReply) error
	grpc.ServerStream
}

type schedulerLeasesServer struct {
	grpc.ServerStream
}

func (x *schedulerLeasesServer) Send(m *LeasesReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Scheduler_Orders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OrdersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SchedulerServer).Orders(m, &schedulerOrdersServer{stream})
}

type Scheduler_OrdersServer interface {
	Send(*OrdersReply) error
	grpc.ServerStream
}

type schedulerOrdersServer struct {
	grpc.ServerStream
}

func (x *schedulerOrdersServer) Send(m *OrdersReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Scheduler_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(StatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SchedulerServer).Stats(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Scheduler_Result_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ResultRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SchedulerServer).Result(m, &schedulerResultServer{stream})
}

type Scheduler_ResultServer interface {
	Send(*ResultReply) error
	grpc.ServerStream
}

type schedulerResultServer struct {
	grpc.ServerStream
}

func (x *schedulerResultServer) Send(m *ResultReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Scheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "qpov.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Scheduler_Get_Handler,
		},
		{
			MethodName: "Renew",
			Handler:    _Scheduler_Renew_Handler,
		},
		{
			MethodName: "Done",
			Handler:    _Scheduler_Done_Handler,
		},
		{
			MethodName: "Failed",
			Handler:    _Scheduler_Failed_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Scheduler_Add_Handler,
		},
		{
			MethodName: "Lease",
			Handler:    _Scheduler_Lease_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _Scheduler_Stats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Leases",
			Handler:       _Scheduler_Leases_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Orders",
			Handler:       _Scheduler_Orders_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Result",
			Handler:       _Scheduler_Result_Handler,
			ServerStreams: true,
		},
	},
}
