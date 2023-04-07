// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gateway.proto

package storegatewaypb

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	storepb "github.com/grafana/mimir/pkg/storegateway/storepb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("gateway.proto", fileDescriptor_f1a937782ebbded5) }

var fileDescriptor_f1a937782ebbded5 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbd, 0x4a, 0x04, 0x31,
	0x14, 0x46, 0x93, 0x66, 0xc1, 0xf8, 0x53, 0x04, 0x14, 0x5c, 0xe1, 0x3e, 0xc2, 0x44, 0xb4, 0x12,
	0x3b, 0x15, 0x6d, 0xc4, 0xc2, 0x05, 0x0b, 0xbb, 0x9b, 0xe5, 0x3a, 0x3b, 0xb8, 0x33, 0x89, 0x49,
	0x06, 0xb1, 0xf3, 0x11, 0x7c, 0x0c, 0x1f, 0xc5, 0x72, 0x2a, 0xd9, 0xd2, 0xc9, 0x34, 0x96, 0xfb,
	0x08, 0xe2, 0x66, 0x07, 0x7f, 0xd8, 0xf2, 0x3b, 0xf7, 0x70, 0x8a, 0x2b, 0x36, 0x73, 0x0c, 0xf4,
	0x88, 0x4f, 0x99, 0x75, 0x26, 0x18, 0xb9, 0xb6, 0x9c, 0x56, 0x0f, 0x8f, 0xf3, 0x22, 0x4c, 0x6a,
	0x9d, 0x8d, 0x4d, 0xa9, 0x72, 0x87, 0x77, 0x58, 0xa1, 0x2a, 0x8b, 0xb2, 0x70, 0xca, 0xde, 0xe7,
	0xca, 0x07, 0xe3, 0x68, 0x29, 0xa7, 0x61, 0xb5, 0x72, 0x76, 0x9c, 0x3a, 0x07, 0xef, 0x5c, 0x6c,
	0x8c, 0xbe, 0xe9, 0x45, 0x52, 0xe4, 0x91, 0x18, 0x8c, 0xc8, 0x15, 0xe4, 0xe5, 0x76, 0x16, 0x26,
	0x58, 0x19, 0x9f, 0xa5, 0x7d, 0x4d, 0x0f, 0x35, 0xf9, 0x30, 0xdc, 0xf9, 0x8f, 0xbd, 0x35, 0x95,
	0xa7, 0x7d, 0x2e, 0x4f, 0x85, 0xb8, 0x44, 0x4d, 0xd3, 0x2b, 0x2c, 0xc9, 0xcb, 0xdd, 0xde, 0xfb,
	0x61, 0x7d, 0x62, 0xb8, 0xea, 0x94, 0x32, 0xf2, 0x5c, 0xac, 0x2f, 0xe8, 0x0d, 0x4e, 0x6b, 0xf2,
	0xf2, 0xaf, 0x9a, 0x60, 0x9f, 0xd9, 0x5b, 0x79, 0x4b, 0x9d, 0x93, 0xb3, 0xa6, 0x05, 0x36, 0x6b,
	0x81, 0xcd, 0x5b, 0xe0, 0xcf, 0x11, 0xf8, 0x6b, 0x04, 0xfe, 0x16, 0x81, 0x37, 0x11, 0xf8, 0x47,
	0x04, 0xfe, 0x19, 0x81, 0xcd, 0x23, 0xf0, 0x97, 0x0e, 0x58, 0xd3, 0x01, 0x9b, 0x75, 0xc0, 0x6e,
	0xb7, 0x7e, 0xbf, 0xcb, 0x6a, 0x3d, 0x58, 0x7c, 0xe9, 0xf0, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xb9,
	0xbe, 0x15, 0x37, 0x7e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StoreGatewayClient is the client API for StoreGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreGatewayClient interface {
	// Series streams each Series for given label matchers and time range.
	//
	// Series should strictly stream full series after series, optionally split by time. This means that a single frame can contain
	// partition of the single series, but once a new series is started to be streamed it means that no more data will
	// be sent for previous one.
	//
	// Series are sorted.
	Series(ctx context.Context, in *storepb.SeriesRequest, opts ...grpc.CallOption) (StoreGateway_SeriesClient, error)
	// LabelNames returns all label names that is available.
	LabelNames(ctx context.Context, in *storepb.LabelNamesRequest, opts ...grpc.CallOption) (*storepb.LabelNamesResponse, error)
	// LabelValues returns all label values for given label name.
	LabelValues(ctx context.Context, in *storepb.LabelValuesRequest, opts ...grpc.CallOption) (*storepb.LabelValuesResponse, error)
}

type storeGatewayClient struct {
	cc *grpc.ClientConn
}

func NewStoreGatewayClient(cc *grpc.ClientConn) StoreGatewayClient {
	return &storeGatewayClient{cc}
}

func (c *storeGatewayClient) Series(ctx context.Context, in *storepb.SeriesRequest, opts ...grpc.CallOption) (StoreGateway_SeriesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StoreGateway_serviceDesc.Streams[0], "/gatewaypb.StoreGateway/Series", opts...)
	if err != nil {
		return nil, err
	}
	x := &storeGatewaySeriesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StoreGateway_SeriesClient interface {
	Recv() (*storepb.SeriesResponse, error)
	grpc.ClientStream
}

type storeGatewaySeriesClient struct {
	grpc.ClientStream
}

func (x *storeGatewaySeriesClient) Recv() (*storepb.SeriesResponse, error) {
	m := new(storepb.SeriesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storeGatewayClient) LabelNames(ctx context.Context, in *storepb.LabelNamesRequest, opts ...grpc.CallOption) (*storepb.LabelNamesResponse, error) {
	out := new(storepb.LabelNamesResponse)
	err := c.cc.Invoke(ctx, "/gatewaypb.StoreGateway/LabelNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeGatewayClient) LabelValues(ctx context.Context, in *storepb.LabelValuesRequest, opts ...grpc.CallOption) (*storepb.LabelValuesResponse, error) {
	out := new(storepb.LabelValuesResponse)
	err := c.cc.Invoke(ctx, "/gatewaypb.StoreGateway/LabelValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreGatewayServer is the server API for StoreGateway service.
type StoreGatewayServer interface {
	// Series streams each Series for given label matchers and time range.
	//
	// Series should strictly stream full series after series, optionally split by time. This means that a single frame can contain
	// partition of the single series, but once a new series is started to be streamed it means that no more data will
	// be sent for previous one.
	//
	// Series are sorted.
	Series(*storepb.SeriesRequest, StoreGateway_SeriesServer) error
	// LabelNames returns all label names that is available.
	LabelNames(context.Context, *storepb.LabelNamesRequest) (*storepb.LabelNamesResponse, error)
	// LabelValues returns all label values for given label name.
	LabelValues(context.Context, *storepb.LabelValuesRequest) (*storepb.LabelValuesResponse, error)
}

// UnimplementedStoreGatewayServer can be embedded to have forward compatible implementations.
type UnimplementedStoreGatewayServer struct {
}

func (*UnimplementedStoreGatewayServer) Series(req *storepb.SeriesRequest, srv StoreGateway_SeriesServer) error {
	return status.Errorf(codes.Unimplemented, "method Series not implemented")
}
func (*UnimplementedStoreGatewayServer) LabelNames(ctx context.Context, req *storepb.LabelNamesRequest) (*storepb.LabelNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LabelNames not implemented")
}
func (*UnimplementedStoreGatewayServer) LabelValues(ctx context.Context, req *storepb.LabelValuesRequest) (*storepb.LabelValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LabelValues not implemented")
}

func RegisterStoreGatewayServer(s *grpc.Server, srv StoreGatewayServer) {
	s.RegisterService(&_StoreGateway_serviceDesc, srv)
}

func _StoreGateway_Series_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(storepb.SeriesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StoreGatewayServer).Series(m, &storeGatewaySeriesServer{stream})
}

type StoreGateway_SeriesServer interface {
	Send(*storepb.SeriesResponse) error
	grpc.ServerStream
}

type storeGatewaySeriesServer struct {
	grpc.ServerStream
}

func (x *storeGatewaySeriesServer) Send(m *storepb.SeriesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StoreGateway_LabelNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(storepb.LabelNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreGatewayServer).LabelNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gatewaypb.StoreGateway/LabelNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreGatewayServer).LabelNames(ctx, req.(*storepb.LabelNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreGateway_LabelValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(storepb.LabelValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreGatewayServer).LabelValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gatewaypb.StoreGateway/LabelValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreGatewayServer).LabelValues(ctx, req.(*storepb.LabelValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoreGateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gatewaypb.StoreGateway",
	HandlerType: (*StoreGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LabelNames",
			Handler:    _StoreGateway_LabelNames_Handler,
		},
		{
			MethodName: "LabelValues",
			Handler:    _StoreGateway_LabelValues_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Series",
			Handler:       _StoreGateway_Series_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gateway.proto",
}
