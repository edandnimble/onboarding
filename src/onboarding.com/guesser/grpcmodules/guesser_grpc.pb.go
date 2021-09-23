// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpcmodules

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

// GuesserRpcClient is the client API for GuesserRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GuesserRpcClient interface {
	// api
	Add(ctx context.Context, in *Guesser, opts ...grpc.CallOption) (*AddGuesserResponse, error)
	Remove(ctx context.Context, in *GuesserId, opts ...grpc.CallOption) (*ResponseStatus, error)
	Query(ctx context.Context, in *GuesserId, opts ...grpc.CallOption) (*QueryResponse, error)
}

type guesserRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewGuesserRpcClient(cc grpc.ClientConnInterface) GuesserRpcClient {
	return &guesserRpcClient{cc}
}

func (c *guesserRpcClient) Add(ctx context.Context, in *Guesser, opts ...grpc.CallOption) (*AddGuesserResponse, error) {
	out := new(AddGuesserResponse)
	err := c.cc.Invoke(ctx, "/guesserModel.GuesserRpc/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guesserRpcClient) Remove(ctx context.Context, in *GuesserId, opts ...grpc.CallOption) (*ResponseStatus, error) {
	out := new(ResponseStatus)
	err := c.cc.Invoke(ctx, "/guesserModel.GuesserRpc/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guesserRpcClient) Query(ctx context.Context, in *GuesserId, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/guesserModel.GuesserRpc/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuesserRpcServer is the server API for GuesserRpc service.
// All implementations must embed UnimplementedGuesserRpcServer
// for forward compatibility
type GuesserRpcServer interface {
	// api
	Add(context.Context, *Guesser) (*AddGuesserResponse, error)
	Remove(context.Context, *GuesserId) (*ResponseStatus, error)
	Query(context.Context, *GuesserId) (*QueryResponse, error)
	mustEmbedUnimplementedGuesserRpcServer()
}

// UnimplementedGuesserRpcServer must be embedded to have forward compatible implementations.
type UnimplementedGuesserRpcServer struct {
}

func (UnimplementedGuesserRpcServer) Add(context.Context, *Guesser) (*AddGuesserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedGuesserRpcServer) Remove(context.Context, *GuesserId) (*ResponseStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedGuesserRpcServer) Query(context.Context, *GuesserId) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedGuesserRpcServer) mustEmbedUnimplementedGuesserRpcServer() {}

// UnsafeGuesserRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GuesserRpcServer will
// result in compilation errors.
type UnsafeGuesserRpcServer interface {
	mustEmbedUnimplementedGuesserRpcServer()
}

func RegisterGuesserRpcServer(s grpc.ServiceRegistrar, srv GuesserRpcServer) {
	s.RegisterService(&GuesserRpc_ServiceDesc, srv)
}

func _GuesserRpc_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Guesser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuesserRpcServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/guesserModel.GuesserRpc/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuesserRpcServer).Add(ctx, req.(*Guesser))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuesserRpc_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuesserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuesserRpcServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/guesserModel.GuesserRpc/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuesserRpcServer).Remove(ctx, req.(*GuesserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuesserRpc_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuesserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuesserRpcServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/guesserModel.GuesserRpc/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuesserRpcServer).Query(ctx, req.(*GuesserId))
	}
	return interceptor(ctx, in, info, handler)
}

// GuesserRpc_ServiceDesc is the grpc.ServiceDesc for GuesserRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GuesserRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "guesserModel.GuesserRpc",
	HandlerType: (*GuesserRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _GuesserRpc_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _GuesserRpc_Remove_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _GuesserRpc_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "onboarding.com/guesser/grpcmodules/guesser.proto",
}
