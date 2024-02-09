// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: expenses.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Microservice_SayHello_FullMethodName = "/proto.Microservice/SayHello"
)

// MicroserviceClient is the client API for Microservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MicroserviceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type microserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewMicroserviceClient(cc grpc.ClientConnInterface) MicroserviceClient {
	return &microserviceClient{cc}
}

func (c *microserviceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, Microservice_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MicroserviceServer is the server API for Microservice service.
// All implementations must embed UnimplementedMicroserviceServer
// for forward compatibility
type MicroserviceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedMicroserviceServer()
}

// UnimplementedMicroserviceServer must be embedded to have forward compatible implementations.
type UnimplementedMicroserviceServer struct {
}

func (UnimplementedMicroserviceServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedMicroserviceServer) mustEmbedUnimplementedMicroserviceServer() {}

// UnsafeMicroserviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MicroserviceServer will
// result in compilation errors.
type UnsafeMicroserviceServer interface {
	mustEmbedUnimplementedMicroserviceServer()
}

func RegisterMicroserviceServer(s grpc.ServiceRegistrar, srv MicroserviceServer) {
	s.RegisterService(&Microservice_ServiceDesc, srv)
}

func _Microservice_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Microservice_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Microservice_ServiceDesc is the grpc.ServiceDesc for Microservice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Microservice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Microservice",
	HandlerType: (*MicroserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Microservice_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "expenses.proto",
}
