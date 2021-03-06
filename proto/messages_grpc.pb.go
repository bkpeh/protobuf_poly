// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package messages

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// GetSystemEventsClient is the client API for GetSystemEvents service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetSystemEventsClient interface {
	GetEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Pid, error)
}

type getSystemEventsClient struct {
	cc grpc.ClientConnInterface
}

func NewGetSystemEventsClient(cc grpc.ClientConnInterface) GetSystemEventsClient {
	return &getSystemEventsClient{cc}
}

func (c *getSystemEventsClient) GetEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Pid, error) {
	out := new(Pid)
	err := c.cc.Invoke(ctx, "/messages.GetSystemEvents/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetSystemEventsServer is the server API for GetSystemEvents service.
// All implementations must embed UnimplementedGetSystemEventsServer
// for forward compatibility
type GetSystemEventsServer interface {
	GetEvent(context.Context, *Event) (*Pid, error)
	mustEmbedUnimplementedGetSystemEventsServer()
}

// UnimplementedGetSystemEventsServer must be embedded to have forward compatible implementations.
type UnimplementedGetSystemEventsServer struct {
}

func (UnimplementedGetSystemEventsServer) GetEvent(context.Context, *Event) (*Pid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedGetSystemEventsServer) mustEmbedUnimplementedGetSystemEventsServer() {}

// UnsafeGetSystemEventsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetSystemEventsServer will
// result in compilation errors.
type UnsafeGetSystemEventsServer interface {
	mustEmbedUnimplementedGetSystemEventsServer()
}

func RegisterGetSystemEventsServer(s grpc.ServiceRegistrar, srv GetSystemEventsServer) {
	s.RegisterService(&_GetSystemEvents_serviceDesc, srv)
}

func _GetSystemEvents_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetSystemEventsServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.GetSystemEvents/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetSystemEventsServer).GetEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetSystemEvents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messages.GetSystemEvents",
	HandlerType: (*GetSystemEventsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvent",
			Handler:    _GetSystemEvents_GetEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messages.proto",
}
