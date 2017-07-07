// Code generated by protoc-gen-go. DO NOT EDIT.
// source: monitorproto.proto

/*
Package monitorproto is a generated protocol buffer package.

It is generated from these files:
	monitorproto.proto

It has these top-level messages:
	Empty
	Heartbeat
	HeartbeatList
	MessageLog
	ValueLog
	LogWriteResponse
	MessageLogReadResponse
	FunctionCall
	Stats
*/
package monitorproto

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

func (m *Heartbeat) GetBeatTime() int64 {
	if m != nil {
		return m.BeatTime
	}
	return 0
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

type MessageLog struct {
	// The entry writing the log
	Entry *discovery.RegistryEntry `protobuf:"bytes,1,opt,name=entry" json:"entry,omitempty"`
	// The message to be written to the logs
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	// The timestamp of the written log
	Timestamps int64 `protobuf:"varint,3,opt,name=timestamps" json:"timestamps,omitempty"`
}

func (m *MessageLog) Reset()                    { *m = MessageLog{} }
func (m *MessageLog) String() string            { return proto.CompactTextString(m) }
func (*MessageLog) ProtoMessage()               {}
func (*MessageLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MessageLog) GetEntry() *discovery.RegistryEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

func (m *MessageLog) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *MessageLog) GetTimestamps() int64 {
	if m != nil {
		return m.Timestamps
	}
	return 0
}

type ValueLog struct {
	Entry *discovery.RegistryEntry `protobuf:"bytes,1,opt,name=entry" json:"entry,omitempty"`
	Value float32                  `protobuf:"fixed32,2,opt,name=value" json:"value,omitempty"`
}

func (m *ValueLog) Reset()                    { *m = ValueLog{} }
func (m *ValueLog) String() string            { return proto.CompactTextString(m) }
func (*ValueLog) ProtoMessage()               {}
func (*ValueLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ValueLog) GetEntry() *discovery.RegistryEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

func (m *ValueLog) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type LogWriteResponse struct {
	// Whether the log was written
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	// The timestamp at which the log was written
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *LogWriteResponse) Reset()                    { *m = LogWriteResponse{} }
func (m *LogWriteResponse) String() string            { return proto.CompactTextString(m) }
func (*LogWriteResponse) ProtoMessage()               {}
func (*LogWriteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LogWriteResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *LogWriteResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type MessageLogReadResponse struct {
	// The response from the logs
	Logs []*MessageLog `protobuf:"bytes,1,rep,name=logs" json:"logs,omitempty"`
}

func (m *MessageLogReadResponse) Reset()                    { *m = MessageLogReadResponse{} }
func (m *MessageLogReadResponse) String() string            { return proto.CompactTextString(m) }
func (*MessageLogReadResponse) ProtoMessage()               {}
func (*MessageLogReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MessageLogReadResponse) GetLogs() []*MessageLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

type FunctionCall struct {
	Binary string `protobuf:"bytes,1,opt,name=binary" json:"binary,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Time   int32  `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
}

func (m *FunctionCall) Reset()                    { *m = FunctionCall{} }
func (m *FunctionCall) String() string            { return proto.CompactTextString(m) }
func (*FunctionCall) ProtoMessage()               {}
func (*FunctionCall) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FunctionCall) GetBinary() string {
	if m != nil {
		return m.Binary
	}
	return ""
}

func (m *FunctionCall) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FunctionCall) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

type Stats struct {
	Binary        string `protobuf:"bytes,1,opt,name=binary" json:"binary,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	NumberOfCalls int32  `protobuf:"varint,3,opt,name=number_of_calls,json=numberOfCalls" json:"number_of_calls,omitempty"`
	MeanRunTime   int32  `protobuf:"varint,4,opt,name=mean_run_time,json=meanRunTime" json:"mean_run_time,omitempty"`
}

func (m *Stats) Reset()                    { *m = Stats{} }
func (m *Stats) String() string            { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()               {}
func (*Stats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Stats) GetBinary() string {
	if m != nil {
		return m.Binary
	}
	return ""
}

func (m *Stats) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Stats) GetNumberOfCalls() int32 {
	if m != nil {
		return m.NumberOfCalls
	}
	return 0
}

func (m *Stats) GetMeanRunTime() int32 {
	if m != nil {
		return m.MeanRunTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "monitorproto.Empty")
	proto.RegisterType((*Heartbeat)(nil), "monitorproto.Heartbeat")
	proto.RegisterType((*HeartbeatList)(nil), "monitorproto.HeartbeatList")
	proto.RegisterType((*MessageLog)(nil), "monitorproto.MessageLog")
	proto.RegisterType((*ValueLog)(nil), "monitorproto.ValueLog")
	proto.RegisterType((*LogWriteResponse)(nil), "monitorproto.LogWriteResponse")
	proto.RegisterType((*MessageLogReadResponse)(nil), "monitorproto.MessageLogReadResponse")
	proto.RegisterType((*FunctionCall)(nil), "monitorproto.FunctionCall")
	proto.RegisterType((*Stats)(nil), "monitorproto.Stats")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MonitorService service

type MonitorServiceClient interface {
	ReceiveHeartbeat(ctx context.Context, in *discovery.RegistryEntry, opts ...grpc.CallOption) (*Heartbeat, error)
	GetHeartbeats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HeartbeatList, error)
	WriteMessageLog(ctx context.Context, in *MessageLog, opts ...grpc.CallOption) (*LogWriteResponse, error)
	WriteValueLog(ctx context.Context, in *ValueLog, opts ...grpc.CallOption) (*LogWriteResponse, error)
	ReadMessageLogs(ctx context.Context, in *discovery.RegistryEntry, opts ...grpc.CallOption) (*MessageLogReadResponse, error)
	WriteFunctionCall(ctx context.Context, in *FunctionCall, opts ...grpc.CallOption) (*Empty, error)
	GetStats(ctx context.Context, in *FunctionCall, opts ...grpc.CallOption) (*Stats, error)
}

type monitorServiceClient struct {
	cc *grpc.ClientConn
}

func NewMonitorServiceClient(cc *grpc.ClientConn) MonitorServiceClient {
	return &monitorServiceClient{cc}
}

func (c *monitorServiceClient) ReceiveHeartbeat(ctx context.Context, in *discovery.RegistryEntry, opts ...grpc.CallOption) (*Heartbeat, error) {
	out := new(Heartbeat)
	err := grpc.Invoke(ctx, "/monitorproto.MonitorService/ReceiveHeartbeat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorServiceClient) GetHeartbeats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HeartbeatList, error) {
	out := new(HeartbeatList)
	err := grpc.Invoke(ctx, "/monitorproto.MonitorService/GetHeartbeats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorServiceClient) WriteMessageLog(ctx context.Context, in *MessageLog, opts ...grpc.CallOption) (*LogWriteResponse, error) {
	out := new(LogWriteResponse)
	err := grpc.Invoke(ctx, "/monitorproto.MonitorService/WriteMessageLog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorServiceClient) WriteValueLog(ctx context.Context, in *ValueLog, opts ...grpc.CallOption) (*LogWriteResponse, error) {
	out := new(LogWriteResponse)
	err := grpc.Invoke(ctx, "/monitorproto.MonitorService/WriteValueLog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorServiceClient) ReadMessageLogs(ctx context.Context, in *discovery.RegistryEntry, opts ...grpc.CallOption) (*MessageLogReadResponse, error) {
	out := new(MessageLogReadResponse)
	err := grpc.Invoke(ctx, "/monitorproto.MonitorService/ReadMessageLogs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorServiceClient) WriteFunctionCall(ctx context.Context, in *FunctionCall, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/monitorproto.MonitorService/WriteFunctionCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorServiceClient) GetStats(ctx context.Context, in *FunctionCall, opts ...grpc.CallOption) (*Stats, error) {
	out := new(Stats)
	err := grpc.Invoke(ctx, "/monitorproto.MonitorService/GetStats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MonitorService service

type MonitorServiceServer interface {
	ReceiveHeartbeat(context.Context, *discovery.RegistryEntry) (*Heartbeat, error)
	GetHeartbeats(context.Context, *Empty) (*HeartbeatList, error)
	WriteMessageLog(context.Context, *MessageLog) (*LogWriteResponse, error)
	WriteValueLog(context.Context, *ValueLog) (*LogWriteResponse, error)
	ReadMessageLogs(context.Context, *discovery.RegistryEntry) (*MessageLogReadResponse, error)
	WriteFunctionCall(context.Context, *FunctionCall) (*Empty, error)
	GetStats(context.Context, *FunctionCall) (*Stats, error)
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
		FullMethod: "/monitorproto.MonitorService/ReceiveHeartbeat",
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
		FullMethod: "/monitorproto.MonitorService/GetHeartbeats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).GetHeartbeats(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorService_WriteMessageLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).WriteMessageLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitorproto.MonitorService/WriteMessageLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).WriteMessageLog(ctx, req.(*MessageLog))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorService_WriteValueLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValueLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).WriteValueLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitorproto.MonitorService/WriteValueLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).WriteValueLog(ctx, req.(*ValueLog))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorService_ReadMessageLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(discovery.RegistryEntry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).ReadMessageLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitorproto.MonitorService/ReadMessageLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).ReadMessageLogs(ctx, req.(*discovery.RegistryEntry))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorService_WriteFunctionCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FunctionCall)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).WriteFunctionCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitorproto.MonitorService/WriteFunctionCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).WriteFunctionCall(ctx, req.(*FunctionCall))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorService_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FunctionCall)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitorproto.MonitorService/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).GetStats(ctx, req.(*FunctionCall))
	}
	return interceptor(ctx, in, info, handler)
}

var _MonitorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "monitorproto.MonitorService",
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
		{
			MethodName: "WriteMessageLog",
			Handler:    _MonitorService_WriteMessageLog_Handler,
		},
		{
			MethodName: "WriteValueLog",
			Handler:    _MonitorService_WriteValueLog_Handler,
		},
		{
			MethodName: "ReadMessageLogs",
			Handler:    _MonitorService_ReadMessageLogs_Handler,
		},
		{
			MethodName: "WriteFunctionCall",
			Handler:    _MonitorService_WriteFunctionCall_Handler,
		},
		{
			MethodName: "GetStats",
			Handler:    _MonitorService_GetStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitorproto.proto",
}

func init() { proto.RegisterFile("monitorproto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x4d, 0xb7, 0x66, 0x6b, 0xef, 0x56, 0x36, 0x0c, 0x1a, 0x51, 0x86, 0xa6, 0xc9, 0x42, 0x68,
	0x0f, 0xd0, 0x49, 0xe3, 0x81, 0x27, 0x78, 0x99, 0xb6, 0x01, 0x6a, 0x01, 0xb9, 0x08, 0x78, 0xab,
	0x9c, 0xec, 0x2e, 0xb3, 0x94, 0xc4, 0x95, 0xed, 0x54, 0xea, 0x13, 0x7f, 0x82, 0x1f, 0x8c, 0xec,
	0xb4, 0xf9, 0x40, 0x14, 0x51, 0x5e, 0x22, 0xdf, 0xe3, 0x7b, 0x8e, 0xaf, 0x8f, 0x4f, 0x80, 0x64,
	0x32, 0x17, 0x46, 0xaa, 0x99, 0x92, 0x46, 0x0e, 0xdd, 0x97, 0xec, 0x37, 0xb1, 0xf0, 0x75, 0x22,
	0xcc, 0x7d, 0x11, 0x0d, 0x63, 0x99, 0x9d, 0x47, 0x4a, 0x9a, 0x7b, 0x54, 0xa9, 0x4c, 0x44, 0x7c,
	0x7e, 0x2b, 0x74, 0x2c, 0xe7, 0xa8, 0x16, 0xe7, 0xae, 0xb1, 0xae, 0x4b, 0x19, 0xba, 0x0b, 0xfe,
	0x55, 0x36, 0x33, 0x0b, 0xfa, 0x1d, 0xfa, 0xef, 0x90, 0x2b, 0x13, 0x21, 0x37, 0x64, 0x08, 0x3e,
	0xe6, 0x46, 0x2d, 0x82, 0xce, 0x69, 0xe7, 0x6c, 0xef, 0x22, 0x18, 0xd6, 0x34, 0x86, 0x89, 0xd0,
	0x46, 0x2d, 0xae, 0xec, 0x3e, 0x2b, 0xdb, 0xc8, 0x31, 0xf4, 0x2d, 0x6f, 0x6a, 0x44, 0x86, 0xc1,
	0xd6, 0x69, 0xe7, 0x6c, 0x9b, 0xf5, 0x2c, 0xf0, 0x45, 0x64, 0x48, 0xdf, 0xc2, 0xa0, 0x52, 0x1e,
	0x09, 0x6d, 0xc8, 0x4b, 0xf0, 0xed, 0x5a, 0x07, 0x9d, 0xd3, 0xed, 0xb3, 0xbd, 0x8b, 0x27, 0xc3,
	0xd6, 0xf5, 0xaa, 0x5e, 0x56, 0x76, 0xd1, 0x39, 0xc0, 0x18, 0xb5, 0xe6, 0x09, 0x8e, 0x64, 0xb2,
	0xf1, 0x68, 0x01, 0xec, 0x66, 0x25, 0xdb, 0x0d, 0xd6, 0x67, 0xab, 0x92, 0x9c, 0x00, 0xd8, 0x79,
	0xb5, 0xe1, 0xd9, 0x4c, 0x07, 0xdb, 0x6e, 0xea, 0x06, 0x42, 0x3f, 0x43, 0xef, 0x2b, 0x4f, 0x8b,
	0xff, 0x3a, 0xf5, 0x31, 0xf8, 0x73, 0xcb, 0x75, 0x67, 0x6e, 0xb1, 0xb2, 0xa0, 0x1f, 0xe0, 0x70,
	0x24, 0x93, 0x6f, 0x4a, 0x18, 0x64, 0xa8, 0x67, 0x32, 0xd7, 0x68, 0xe7, 0xd3, 0x45, 0x1c, 0xa3,
	0xd6, 0x4e, 0xbb, 0xc7, 0x56, 0x25, 0x79, 0x0a, 0xfd, 0x6a, 0x9a, 0xa5, 0xa9, 0x35, 0x40, 0xaf,
	0xe1, 0xa8, 0x76, 0x85, 0x21, 0xbf, 0xad, 0x14, 0x5f, 0x40, 0x37, 0x95, 0xc9, 0xca, 0xdd, 0xa0,
	0xed, 0x6e, 0x83, 0xe3, 0xba, 0xe8, 0x47, 0xd8, 0xbf, 0x2e, 0xf2, 0xd8, 0x08, 0x99, 0x5f, 0xf2,
	0x34, 0x25, 0x47, 0xb0, 0x13, 0x89, 0x9c, 0x2f, 0xaf, 0xda, 0x67, 0xcb, 0x8a, 0x10, 0xe8, 0xe6,
	0x3c, 0x5b, 0x99, 0xe8, 0xd6, 0x16, 0x73, 0x2f, 0x6e, 0xbd, 0xf3, 0x99, 0x5b, 0xd3, 0x1f, 0xe0,
	0x4f, 0x0c, 0x37, 0x7a, 0x23, 0xa1, 0xe7, 0x70, 0x90, 0x17, 0x59, 0x84, 0x6a, 0x2a, 0xef, 0xa6,
	0x31, 0x4f, 0x53, 0xbd, 0xd4, 0x1c, 0x94, 0xf0, 0xa7, 0x3b, 0x3b, 0x9b, 0x26, 0x14, 0x06, 0x19,
	0xf2, 0x7c, 0xaa, 0x8a, 0xbc, 0xcc, 0x5a, 0xd7, 0x75, 0xed, 0x59, 0x90, 0x15, 0xb9, 0x8d, 0xdb,
	0xc5, 0xcf, 0x2e, 0x3c, 0x18, 0x97, 0x57, 0x9e, 0xa0, 0x9a, 0x8b, 0x18, 0xc9, 0x0d, 0x1c, 0x32,
	0x8c, 0x51, 0xcc, 0xb1, 0x8e, 0xf8, 0xda, 0x27, 0x0c, 0xd7, 0xe5, 0x91, 0x7a, 0xe4, 0x12, 0x06,
	0x37, 0x68, 0x2a, 0x44, 0x93, 0x47, 0xed, 0x5e, 0xf7, 0x2b, 0x85, 0xc7, 0x6b, 0x04, 0x6c, 0xf8,
	0xa9, 0x47, 0xc6, 0x70, 0xe0, 0x22, 0xd0, 0x08, 0xf5, 0xda, 0x47, 0x0a, 0x4f, 0xda, 0x3b, 0xbf,
	0xc7, 0x87, 0x7a, 0xe4, 0x3d, 0x0c, 0x1c, 0x54, 0x65, 0xf5, 0xa8, 0x4d, 0x59, 0xe1, 0xff, 0x20,
	0x35, 0x81, 0x03, 0x9b, 0xa4, 0xfa, 0x78, 0xfd, 0x17, 0x9b, 0x9e, 0xad, 0x0d, 0x56, 0x23, 0x8c,
	0xd4, 0x23, 0xd7, 0xf0, 0xd0, 0x9d, 0xd3, 0x4a, 0x59, 0xd8, 0x26, 0x37, 0xf7, 0xc2, 0x3f, 0x79,
	0x4a, 0x3d, 0xf2, 0x06, 0x7a, 0x37, 0x68, 0xca, 0x6c, 0x6d, 0x40, 0x77, 0x04, 0xea, 0x45, 0x3b,
	0xae, 0x7c, 0xf5, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xde, 0xe5, 0x72, 0x78, 0x4c, 0x05, 0x00, 0x00,
}
