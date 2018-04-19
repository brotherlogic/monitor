// Code generated by protoc-gen-go. DO NOT EDIT.
// source: monitorproto.proto

/*
Package monitorproto is a generated protocol buffer package.

It is generated from these files:
	monitorproto.proto

It has these top-level messages:
	Empty
	MessageLog
	LogWriteResponse
	MessageLogReadResponse
	Milestone
	FunctionCall
	Stats
	StatsList
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
func (*MessageLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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

type LogWriteResponse struct {
	// Whether the log was written
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	// The timestamp at which the log was written
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *LogWriteResponse) Reset()                    { *m = LogWriteResponse{} }
func (m *LogWriteResponse) String() string            { return proto.CompactTextString(m) }
func (*LogWriteResponse) ProtoMessage()               {}
func (*LogWriteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
func (*MessageLogReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MessageLogReadResponse) GetLogs() []*MessageLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

type Milestone struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Time int32  `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
}

func (m *Milestone) Reset()                    { *m = Milestone{} }
func (m *Milestone) String() string            { return proto.CompactTextString(m) }
func (*Milestone) ProtoMessage()               {}
func (*Milestone) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Milestone) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Milestone) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

type FunctionCall struct {
	Binary     string       `protobuf:"bytes,1,opt,name=binary" json:"binary,omitempty"`
	Name       string       `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Time       int32        `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	Milestones []*Milestone `protobuf:"bytes,4,rep,name=milestones" json:"milestones,omitempty"`
}

func (m *FunctionCall) Reset()                    { *m = FunctionCall{} }
func (m *FunctionCall) String() string            { return proto.CompactTextString(m) }
func (*FunctionCall) ProtoMessage()               {}
func (*FunctionCall) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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

func (m *FunctionCall) GetMilestones() []*Milestone {
	if m != nil {
		return m.Milestones
	}
	return nil
}

type Stats struct {
	Binary        string        `protobuf:"bytes,1,opt,name=binary" json:"binary,omitempty"`
	Name          string        `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	NumberOfCalls int32         `protobuf:"varint,3,opt,name=number_of_calls,json=numberOfCalls" json:"number_of_calls,omitempty"`
	MeanRunTime   int32         `protobuf:"varint,4,opt,name=mean_run_time,json=meanRunTime" json:"mean_run_time,omitempty"`
	RunTimes      []int32       `protobuf:"varint,5,rep,packed,name=run_times,json=runTimes" json:"run_times,omitempty"`
	Slowest       *FunctionCall `protobuf:"bytes,6,opt,name=slowest" json:"slowest,omitempty"`
}

func (m *Stats) Reset()                    { *m = Stats{} }
func (m *Stats) String() string            { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()               {}
func (*Stats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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

func (m *Stats) GetRunTimes() []int32 {
	if m != nil {
		return m.RunTimes
	}
	return nil
}

func (m *Stats) GetSlowest() *FunctionCall {
	if m != nil {
		return m.Slowest
	}
	return nil
}

type StatsList struct {
	Stats []*Stats `protobuf:"bytes,1,rep,name=stats" json:"stats,omitempty"`
}

func (m *StatsList) Reset()                    { *m = StatsList{} }
func (m *StatsList) String() string            { return proto.CompactTextString(m) }
func (*StatsList) ProtoMessage()               {}
func (*StatsList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *StatsList) GetStats() []*Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "monitorproto.Empty")
	proto.RegisterType((*MessageLog)(nil), "monitorproto.MessageLog")
	proto.RegisterType((*LogWriteResponse)(nil), "monitorproto.LogWriteResponse")
	proto.RegisterType((*MessageLogReadResponse)(nil), "monitorproto.MessageLogReadResponse")
	proto.RegisterType((*Milestone)(nil), "monitorproto.Milestone")
	proto.RegisterType((*FunctionCall)(nil), "monitorproto.FunctionCall")
	proto.RegisterType((*Stats)(nil), "monitorproto.Stats")
	proto.RegisterType((*StatsList)(nil), "monitorproto.StatsList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MonitorService service

type MonitorServiceClient interface {
	WriteMessageLog(ctx context.Context, in *MessageLog, opts ...grpc.CallOption) (*LogWriteResponse, error)
	ReadMessageLogs(ctx context.Context, in *discovery.RegistryEntry, opts ...grpc.CallOption) (*MessageLogReadResponse, error)
	WriteFunctionCall(ctx context.Context, in *FunctionCall, opts ...grpc.CallOption) (*Empty, error)
	GetStats(ctx context.Context, in *FunctionCall, opts ...grpc.CallOption) (*StatsList, error)
	ClearStats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type monitorServiceClient struct {
	cc *grpc.ClientConn
}

func NewMonitorServiceClient(cc *grpc.ClientConn) MonitorServiceClient {
	return &monitorServiceClient{cc}
}

func (c *monitorServiceClient) WriteMessageLog(ctx context.Context, in *MessageLog, opts ...grpc.CallOption) (*LogWriteResponse, error) {
	out := new(LogWriteResponse)
	err := grpc.Invoke(ctx, "/monitorproto.MonitorService/WriteMessageLog", in, out, c.cc, opts...)
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

func (c *monitorServiceClient) GetStats(ctx context.Context, in *FunctionCall, opts ...grpc.CallOption) (*StatsList, error) {
	out := new(StatsList)
	err := grpc.Invoke(ctx, "/monitorproto.MonitorService/GetStats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorServiceClient) ClearStats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/monitorproto.MonitorService/ClearStats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MonitorService service

type MonitorServiceServer interface {
	WriteMessageLog(context.Context, *MessageLog) (*LogWriteResponse, error)
	ReadMessageLogs(context.Context, *discovery.RegistryEntry) (*MessageLogReadResponse, error)
	WriteFunctionCall(context.Context, *FunctionCall) (*Empty, error)
	GetStats(context.Context, *FunctionCall) (*StatsList, error)
	ClearStats(context.Context, *Empty) (*Empty, error)
}

func RegisterMonitorServiceServer(s *grpc.Server, srv MonitorServiceServer) {
	s.RegisterService(&_MonitorService_serviceDesc, srv)
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

func _MonitorService_ClearStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).ClearStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitorproto.MonitorService/ClearStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).ClearStats(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _MonitorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "monitorproto.MonitorService",
	HandlerType: (*MonitorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteMessageLog",
			Handler:    _MonitorService_WriteMessageLog_Handler,
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
		{
			MethodName: "ClearStats",
			Handler:    _MonitorService_ClearStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitorproto.proto",
}

func init() { proto.RegisterFile("monitorproto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xd1, 0x6a, 0xdb, 0x4a,
	0x10, 0xb5, 0x22, 0xcb, 0xb6, 0x26, 0xc9, 0xf5, 0xed, 0x16, 0x12, 0xe1, 0x96, 0x60, 0x96, 0x52,
	0x5c, 0x28, 0x36, 0x24, 0xa5, 0xe9, 0x6b, 0x09, 0x71, 0xa1, 0xd8, 0x14, 0xd6, 0x85, 0x3e, 0x1a,
	0x59, 0xd9, 0x28, 0x0b, 0xd2, 0xae, 0xd9, 0x59, 0xb9, 0xf8, 0x0b, 0xfa, 0x7f, 0xfd, 0x9b, 0xbe,
	0x15, 0xad, 0x2c, 0x4b, 0x36, 0x76, 0xa0, 0x2f, 0x46, 0x33, 0x73, 0xe6, 0xcc, 0x99, 0xf1, 0x59,
	0x20, 0xa9, 0x92, 0xc2, 0x28, 0xbd, 0xd4, 0xca, 0xa8, 0xa1, 0xfd, 0x25, 0x67, 0xf5, 0x5c, 0xef,
	0x36, 0x16, 0xe6, 0x29, 0x5b, 0x0c, 0x23, 0x95, 0x8e, 0x16, 0x5a, 0x99, 0x27, 0xae, 0x13, 0x15,
	0x8b, 0x68, 0xf4, 0x20, 0x30, 0x52, 0x2b, 0xae, 0xd7, 0x23, 0x0b, 0xac, 0xe2, 0x82, 0x86, 0xb6,
	0xc1, 0xbb, 0x4f, 0x97, 0x66, 0x4d, 0x57, 0x00, 0x53, 0x8e, 0x18, 0xc6, 0x7c, 0xa2, 0x62, 0x32,
	0x04, 0x8f, 0x4b, 0xa3, 0xd7, 0x81, 0xd3, 0x77, 0x06, 0xa7, 0xd7, 0xc1, 0xb0, 0xea, 0x63, 0x3c,
	0x16, 0x68, 0xf4, 0xfa, 0x3e, 0xaf, 0xb3, 0x02, 0x46, 0x02, 0x68, 0xa7, 0x45, 0x77, 0x70, 0xd2,
	0x77, 0x06, 0x3e, 0x2b, 0x43, 0x72, 0x05, 0x60, 0x44, 0xca, 0xd1, 0x84, 0xe9, 0x12, 0x03, 0xb7,
	0xef, 0x0c, 0x5c, 0x56, 0xcb, 0xd0, 0xaf, 0xf0, 0xff, 0x44, 0xc5, 0x3f, 0xb4, 0x30, 0x9c, 0x71,
	0x5c, 0x2a, 0x89, 0x3c, 0x67, 0xc3, 0x2c, 0x8a, 0x38, 0xa2, 0x9d, 0xdf, 0x61, 0x65, 0x48, 0x5e,
	0x83, 0xbf, 0xed, 0xb5, 0x93, 0x5c, 0x56, 0x25, 0xe8, 0x18, 0x2e, 0xaa, 0x1d, 0x18, 0x0f, 0x1f,
	0xb6, 0x8c, 0xef, 0xa1, 0x99, 0xa8, 0x38, 0xa7, 0x73, 0xed, 0x3a, 0x3b, 0x07, 0xad, 0xf5, 0x58,
	0x14, 0xbd, 0x01, 0x7f, 0x2a, 0x12, 0x8e, 0x46, 0x49, 0x4e, 0x08, 0x34, 0x65, 0x98, 0x72, 0xab,
	0xc4, 0x67, 0xf6, 0x3b, 0xcf, 0xe5, 0x53, 0xad, 0x02, 0x8f, 0xd9, 0x6f, 0xfa, 0xcb, 0x81, 0xb3,
	0x71, 0x26, 0x23, 0x23, 0x94, 0xbc, 0x0b, 0x93, 0x84, 0x5c, 0x40, 0x6b, 0x21, 0x64, 0xb8, 0x39,
	0xa2, 0xcf, 0x36, 0xd1, 0x96, 0xf0, 0xe4, 0x00, 0xa1, 0x5b, 0x11, 0x92, 0x5b, 0x80, 0xb4, 0x54,
	0x81, 0x41, 0xd3, 0x2a, 0xbf, 0xdc, 0x53, 0x5e, 0xd6, 0x59, 0x0d, 0x4a, 0x7f, 0x3b, 0xe0, 0xcd,
	0x4c, 0x68, 0xf0, 0x9f, 0x24, 0xbc, 0x85, 0xae, 0xcc, 0xd2, 0x05, 0xd7, 0x73, 0xf5, 0x38, 0x8f,
	0xc2, 0x24, 0xc1, 0x8d, 0x9a, 0xf3, 0x22, 0xfd, 0xed, 0x31, 0xdf, 0x0a, 0x09, 0x85, 0xf3, 0x94,
	0x87, 0x72, 0xae, 0x33, 0x39, 0xb7, 0x9a, 0x9b, 0x16, 0x75, 0x9a, 0x27, 0x59, 0x26, 0xbf, 0xe7,
	0xd2, 0x5f, 0x81, 0x5f, 0x96, 0x31, 0xf0, 0xfa, 0xee, 0xc0, 0x63, 0x1d, 0x5d, 0xd4, 0x90, 0x7c,
	0x80, 0x36, 0x26, 0xea, 0x27, 0x47, 0x13, 0xb4, 0xac, 0xbb, 0x7a, 0xbb, 0x4b, 0xd5, 0x8f, 0xc8,
	0x4a, 0x28, 0xfd, 0x08, 0xbe, 0xdd, 0x69, 0x22, 0xd0, 0x90, 0x77, 0xe0, 0x61, 0x1e, 0x6c, 0xfe,
	0xcf, 0x97, 0xbb, 0x04, 0x16, 0xc7, 0x0a, 0xc4, 0xf5, 0x9f, 0x13, 0xf8, 0x6f, 0x5a, 0x54, 0x67,
	0x5c, 0xaf, 0x44, 0xc4, 0xc9, 0x14, 0xba, 0xd6, 0x6f, 0x35, 0xbf, 0x1f, 0x75, 0x44, 0xef, 0x6a,
	0xb7, 0xb2, 0xef, 0x55, 0xda, 0x20, 0x33, 0xe8, 0xe6, 0x5e, 0xab, 0x7a, 0x90, 0x1c, 0x7d, 0x2f,
	0xbd, 0x37, 0x47, 0xad, 0x57, 0xb3, 0x2b, 0x6d, 0x90, 0x31, 0xbc, 0xb0, 0x73, 0x76, 0x1c, 0xf5,
	0xcc, 0xa1, 0x7a, 0x7b, 0x37, 0x28, 0x1e, 0x75, 0x83, 0x7c, 0x86, 0xce, 0x17, 0x6e, 0x0a, 0x37,
	0x3c, 0xd7, 0x7e, 0x79, 0xe0, 0x84, 0xf9, 0xa9, 0x69, 0x83, 0x7c, 0x02, 0xb8, 0x4b, 0x78, 0xa8,
	0x0b, 0x92, 0x43, 0x73, 0x8e, 0x0c, 0x5f, 0xb4, 0x6c, 0x78, 0xf3, 0x37, 0x00, 0x00, 0xff, 0xff,
	0xa1, 0x64, 0x33, 0x75, 0xc0, 0x04, 0x00, 0x00,
}
