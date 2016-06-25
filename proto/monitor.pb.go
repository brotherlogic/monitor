// Code generated by protoc-gen-go.
// source: monitor.proto
// DO NOT EDIT!

/*
Package monitor is a generated protocol buffer package.

It is generated from these files:
	monitor.proto

It has these top-level messages:
	Empty
	Heartbeat
	HeartbeatList
*/
package monitor

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import discovery "github.com/brotherlogic/discovery/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Heartbeat struct {
	// The entry
	Entry *discovery.RegistryEntry `protobuf:"bytes,1,opt,name=entry" json:"entry,omitempty"`
	// The time of the beat
	BeatTime int64 `protobuf:"varint,2,opt,name=beat_time,json=beatTime" json:"beat_time,omitempty"`
}

func (m *Heartbeat) Reset()                    { *m = Heartbeat{} }
func (m *Heartbeat) String() string            { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()               {}
func (*Heartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Heartbeat) GetEntry() *discovery.RegistryEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type HeartbeatList struct {
	Beats []*Heartbeat `protobuf:"bytes,1,rep,name=beats" json:"beats,omitempty"`
}

func (m *HeartbeatList) Reset()                    { *m = HeartbeatList{} }
func (m *HeartbeatList) String() string            { return proto.CompactTextString(m) }
func (*HeartbeatList) ProtoMessage()               {}
func (*HeartbeatList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HeartbeatList) GetBeats() []*Heartbeat {
	if m != nil {
		return m.Beats
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "monitor.Empty")
	proto.RegisterType((*Heartbeat)(nil), "monitor.Heartbeat")
	proto.RegisterType((*HeartbeatList)(nil), "monitor.HeartbeatList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for MonitorService service

type MonitorServiceClient interface {
	ReceiveHeartbeat(ctx context.Context, in *discovery.RegistryEntry, opts ...grpc.CallOption) (*Heartbeat, error)
	GetHeartbeats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HeartbeatList, error)
}

type monitorServiceClient struct {
	cc *grpc.ClientConn
}

func NewMonitorServiceClient(cc *grpc.ClientConn) MonitorServiceClient {
	return &monitorServiceClient{cc}
}

func (c *monitorServiceClient) ReceiveHeartbeat(ctx context.Context, in *discovery.RegistryEntry, opts ...grpc.CallOption) (*Heartbeat, error) {
	out := new(Heartbeat)
	err := grpc.Invoke(ctx, "/monitor.MonitorService/ReceiveHeartbeat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorServiceClient) GetHeartbeats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HeartbeatList, error) {
	out := new(HeartbeatList)
	err := grpc.Invoke(ctx, "/monitor.MonitorService/GetHeartbeats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MonitorService service

type MonitorServiceServer interface {
	ReceiveHeartbeat(context.Context, *discovery.RegistryEntry) (*Heartbeat, error)
	GetHeartbeats(context.Context, *Empty) (*HeartbeatList, error)
}

func RegisterMonitorServiceServer(s *grpc.Server, srv MonitorServiceServer) {
	s.RegisterService(&_MonitorService_serviceDesc, srv)
}

func _MonitorService_ReceiveHeartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(discovery.RegistryEntry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).ReceiveHeartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.MonitorService/ReceiveHeartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).ReceiveHeartbeat(ctx, req.(*discovery.RegistryEntry))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorService_GetHeartbeats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).GetHeartbeats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.MonitorService/GetHeartbeats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).GetHeartbeats(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _MonitorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "monitor.MonitorService",
	HandlerType: (*MonitorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveHeartbeat",
			Handler:    _MonitorService_ReceiveHeartbeat_Handler,
		},
		{
			MethodName: "GetHeartbeats",
			Handler:    _MonitorService_GetHeartbeats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("monitor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x1b, 0xaa, 0x50, 0x7a, 0xa8, 0x15, 0xf2, 0x80, 0xa2, 0xb0, 0x20, 0x4f, 0x9d, 0x1c,
	0xa9, 0x0c, 0x88, 0x15, 0xa9, 0x82, 0x01, 0x16, 0xc3, 0xc0, 0x86, 0x1a, 0x73, 0x4a, 0x2d, 0x91,
	0xba, 0x72, 0x8e, 0x48, 0xf9, 0x15, 0xfc, 0x65, 0xce, 0x09, 0x38, 0x03, 0xb0, 0xf9, 0xee, 0xbd,
	0xef, 0xe9, 0xf9, 0x60, 0x51, 0xbb, 0xbd, 0x25, 0xe7, 0xd5, 0xc1, 0x3b, 0x72, 0x62, 0xf6, 0x3d,
	0xe6, 0xd7, 0x95, 0xa5, 0xdd, 0x47, 0xa9, 0x8c, 0xab, 0x8b, 0x92, 0xa5, 0x1d, 0xfa, 0x77, 0x57,
	0x59, 0x53, 0xbc, 0xd9, 0xc6, 0xb8, 0x16, 0x7d, 0x57, 0xf4, 0xc4, 0x38, 0x0f, 0x09, 0x72, 0x06,
	0xe9, 0xa6, 0x3e, 0x50, 0x27, 0x5f, 0x60, 0x7e, 0x8f, 0x5b, 0x4f, 0x25, 0x6e, 0x49, 0x28, 0x48,
	0x71, 0x4f, 0xbe, 0xcb, 0x92, 0xcb, 0x64, 0x75, 0xba, 0xce, 0xd4, 0x88, 0x69, 0xac, 0x6c, 0xc3,
	0xd2, 0x26, 0xe8, 0x7a, 0xb0, 0x89, 0x0b, 0x98, 0x07, 0xee, 0x95, 0x6c, 0x8d, 0xd9, 0x11, 0x33,
	0x53, 0x7d, 0x12, 0x16, 0xcf, 0x3c, 0xcb, 0x1b, 0x58, 0xc4, 0xe4, 0x07, 0x66, 0xc5, 0x0a, 0xd2,
	0xf0, 0x6e, 0x38, 0x7d, 0xca, 0xe9, 0x42, 0xfd, 0x7c, 0x2a, 0xda, 0xf4, 0x60, 0x58, 0x7f, 0x26,
	0xb0, 0x7c, 0x1c, 0xc4, 0x27, 0xf4, 0xad, 0x35, 0x28, 0x6e, 0xe1, 0x4c, 0xa3, 0x41, 0xdb, 0xe2,
	0x58, 0xf7, 0xdf, 0x7e, 0xf9, 0x1f, 0xd9, 0x72, 0x22, 0xb8, 0xd1, 0x1d, 0x52, 0xdc, 0x34, 0x62,
	0x19, 0x6d, 0xfd, 0x31, 0xf2, 0xf3, 0xdf, 0x58, 0x68, 0x2e, 0x27, 0xe5, 0x71, 0x7f, 0xb6, 0xab,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x79, 0xf1, 0x21, 0xc6, 0x89, 0x01, 0x00, 0x00,
}
