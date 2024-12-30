// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.4
// source: cacheserver/v1/cacheserver.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CacheServer_Set_FullMethodName       = "/cacheserver.v1.CacheServer/Set"
	CacheServer_Get_FullMethodName       = "/cacheserver.v1.CacheServer/Get"
	CacheServer_Del_FullMethodName       = "/cacheserver.v1.CacheServer/Del"
	CacheServer_SetSecret_FullMethodName = "/cacheserver.v1.CacheServer/SetSecret"
	CacheServer_GetSecret_FullMethodName = "/cacheserver.v1.CacheServer/GetSecret"
	CacheServer_DelSecret_FullMethodName = "/cacheserver.v1.CacheServer/DelSecret"
)

// CacheServerClient is the client API for CacheServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CacheServerClient interface {
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetSecret(ctx context.Context, in *SetSecretRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error)
	DelSecret(ctx context.Context, in *DelSecretRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type cacheServerClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheServerClient(cc grpc.ClientConnInterface) CacheServerClient {
	return &cacheServerClient{cc}
}

func (c *cacheServerClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CacheServer_Set_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServerClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, CacheServer_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServerClient) Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CacheServer_Del_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServerClient) SetSecret(ctx context.Context, in *SetSecretRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CacheServer_SetSecret_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServerClient) GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSecretResponse)
	err := c.cc.Invoke(ctx, CacheServer_GetSecret_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServerClient) DelSecret(ctx context.Context, in *DelSecretRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CacheServer_DelSecret_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServerServer is the server API for CacheServer service.
// All implementations must embed UnimplementedCacheServerServer
// for forward compatibility.
type CacheServerServer interface {
	Set(context.Context, *SetRequest) (*emptypb.Empty, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Del(context.Context, *DelRequest) (*emptypb.Empty, error)
	SetSecret(context.Context, *SetSecretRequest) (*emptypb.Empty, error)
	GetSecret(context.Context, *GetSecretRequest) (*GetSecretResponse, error)
	DelSecret(context.Context, *DelSecretRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCacheServerServer()
}

// UnimplementedCacheServerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCacheServerServer struct{}

func (UnimplementedCacheServerServer) Set(context.Context, *SetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedCacheServerServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCacheServerServer) Del(context.Context, *DelRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (UnimplementedCacheServerServer) SetSecret(context.Context, *SetSecretRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSecret not implemented")
}
func (UnimplementedCacheServerServer) GetSecret(context.Context, *GetSecretRequest) (*GetSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecret not implemented")
}
func (UnimplementedCacheServerServer) DelSecret(context.Context, *DelSecretRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelSecret not implemented")
}
func (UnimplementedCacheServerServer) mustEmbedUnimplementedCacheServerServer() {}
func (UnimplementedCacheServerServer) testEmbeddedByValue()                     {}

// UnsafeCacheServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacheServerServer will
// result in compilation errors.
type UnsafeCacheServerServer interface {
	mustEmbedUnimplementedCacheServerServer()
}

func RegisterCacheServerServer(s grpc.ServiceRegistrar, srv CacheServerServer) {
	// If the following call pancis, it indicates UnimplementedCacheServerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CacheServer_ServiceDesc, srv)
}

func _CacheServer_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServerServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheServer_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServerServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheServer_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheServer_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServerServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheServer_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServerServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheServer_Del_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServerServer).Del(ctx, req.(*DelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheServer_SetSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServerServer).SetSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheServer_SetSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServerServer).SetSecret(ctx, req.(*SetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheServer_GetSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServerServer).GetSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheServer_GetSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServerServer).GetSecret(ctx, req.(*GetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheServer_DelSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServerServer).DelSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheServer_DelSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServerServer).DelSecret(ctx, req.(*DelSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CacheServer_ServiceDesc is the grpc.ServiceDesc for CacheServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CacheServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cacheserver.v1.CacheServer",
	HandlerType: (*CacheServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _CacheServer_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CacheServer_Get_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _CacheServer_Del_Handler,
		},
		{
			MethodName: "SetSecret",
			Handler:    _CacheServer_SetSecret_Handler,
		},
		{
			MethodName: "GetSecret",
			Handler:    _CacheServer_GetSecret_Handler,
		},
		{
			MethodName: "DelSecret",
			Handler:    _CacheServer_DelSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cacheserver/v1/cacheserver.proto",
}
