// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.26.1
// source: common/advertiser/v1/advertiser.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Advertiser_CreateAdvertiser_FullMethodName = "/api.common.advertiser.v1.Advertiser/CreateAdvertiser"
	Advertiser_UpdateAdvertiser_FullMethodName = "/api.common.advertiser.v1.Advertiser/UpdateAdvertiser"
	Advertiser_DeleteAdvertiser_FullMethodName = "/api.common.advertiser.v1.Advertiser/DeleteAdvertiser"
	Advertiser_GetAdvertiser_FullMethodName    = "/api.common.advertiser.v1.Advertiser/GetAdvertiser"
	Advertiser_ListAdvertiser_FullMethodName   = "/api.common.advertiser.v1.Advertiser/ListAdvertiser"
)

// AdvertiserClient is the client API for Advertiser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdvertiserClient interface {
	CreateAdvertiser(ctx context.Context, in *CreateAdvertiserRequest, opts ...grpc.CallOption) (*CreateAdvertiserReply, error)
	UpdateAdvertiser(ctx context.Context, in *UpdateAdvertiserRequest, opts ...grpc.CallOption) (*UpdateAdvertiserReply, error)
	DeleteAdvertiser(ctx context.Context, in *DeleteAdvertiserRequest, opts ...grpc.CallOption) (*DeleteAdvertiserReply, error)
	GetAdvertiser(ctx context.Context, in *GetAdvertiserRequest, opts ...grpc.CallOption) (*GetAdvertiserReply, error)
	ListAdvertiser(ctx context.Context, in *ListAdvertiserRequest, opts ...grpc.CallOption) (*ListAdvertiserReply, error)
}

type advertiserClient struct {
	cc grpc.ClientConnInterface
}

func NewAdvertiserClient(cc grpc.ClientConnInterface) AdvertiserClient {
	return &advertiserClient{cc}
}

func (c *advertiserClient) CreateAdvertiser(ctx context.Context, in *CreateAdvertiserRequest, opts ...grpc.CallOption) (*CreateAdvertiserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAdvertiserReply)
	err := c.cc.Invoke(ctx, Advertiser_CreateAdvertiser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertiserClient) UpdateAdvertiser(ctx context.Context, in *UpdateAdvertiserRequest, opts ...grpc.CallOption) (*UpdateAdvertiserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateAdvertiserReply)
	err := c.cc.Invoke(ctx, Advertiser_UpdateAdvertiser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertiserClient) DeleteAdvertiser(ctx context.Context, in *DeleteAdvertiserRequest, opts ...grpc.CallOption) (*DeleteAdvertiserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteAdvertiserReply)
	err := c.cc.Invoke(ctx, Advertiser_DeleteAdvertiser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertiserClient) GetAdvertiser(ctx context.Context, in *GetAdvertiserRequest, opts ...grpc.CallOption) (*GetAdvertiserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAdvertiserReply)
	err := c.cc.Invoke(ctx, Advertiser_GetAdvertiser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertiserClient) ListAdvertiser(ctx context.Context, in *ListAdvertiserRequest, opts ...grpc.CallOption) (*ListAdvertiserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAdvertiserReply)
	err := c.cc.Invoke(ctx, Advertiser_ListAdvertiser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdvertiserServer is the server API for Advertiser service.
// All implementations must embed UnimplementedAdvertiserServer
// for forward compatibility
type AdvertiserServer interface {
	CreateAdvertiser(context.Context, *CreateAdvertiserRequest) (*CreateAdvertiserReply, error)
	UpdateAdvertiser(context.Context, *UpdateAdvertiserRequest) (*UpdateAdvertiserReply, error)
	DeleteAdvertiser(context.Context, *DeleteAdvertiserRequest) (*DeleteAdvertiserReply, error)
	GetAdvertiser(context.Context, *GetAdvertiserRequest) (*GetAdvertiserReply, error)
	ListAdvertiser(context.Context, *ListAdvertiserRequest) (*ListAdvertiserReply, error)
	mustEmbedUnimplementedAdvertiserServer()
}

// UnimplementedAdvertiserServer must be embedded to have forward compatible implementations.
type UnimplementedAdvertiserServer struct {
}

func (UnimplementedAdvertiserServer) CreateAdvertiser(context.Context, *CreateAdvertiserRequest) (*CreateAdvertiserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdvertiser not implemented")
}
func (UnimplementedAdvertiserServer) UpdateAdvertiser(context.Context, *UpdateAdvertiserRequest) (*UpdateAdvertiserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAdvertiser not implemented")
}
func (UnimplementedAdvertiserServer) DeleteAdvertiser(context.Context, *DeleteAdvertiserRequest) (*DeleteAdvertiserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAdvertiser not implemented")
}
func (UnimplementedAdvertiserServer) GetAdvertiser(context.Context, *GetAdvertiserRequest) (*GetAdvertiserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdvertiser not implemented")
}
func (UnimplementedAdvertiserServer) ListAdvertiser(context.Context, *ListAdvertiserRequest) (*ListAdvertiserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAdvertiser not implemented")
}
func (UnimplementedAdvertiserServer) mustEmbedUnimplementedAdvertiserServer() {}

// UnsafeAdvertiserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdvertiserServer will
// result in compilation errors.
type UnsafeAdvertiserServer interface {
	mustEmbedUnimplementedAdvertiserServer()
}

func RegisterAdvertiserServer(s grpc.ServiceRegistrar, srv AdvertiserServer) {
	s.RegisterService(&Advertiser_ServiceDesc, srv)
}

func _Advertiser_CreateAdvertiser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdvertiserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertiserServer).CreateAdvertiser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Advertiser_CreateAdvertiser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertiserServer).CreateAdvertiser(ctx, req.(*CreateAdvertiserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Advertiser_UpdateAdvertiser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAdvertiserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertiserServer).UpdateAdvertiser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Advertiser_UpdateAdvertiser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertiserServer).UpdateAdvertiser(ctx, req.(*UpdateAdvertiserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Advertiser_DeleteAdvertiser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAdvertiserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertiserServer).DeleteAdvertiser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Advertiser_DeleteAdvertiser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertiserServer).DeleteAdvertiser(ctx, req.(*DeleteAdvertiserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Advertiser_GetAdvertiser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdvertiserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertiserServer).GetAdvertiser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Advertiser_GetAdvertiser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertiserServer).GetAdvertiser(ctx, req.(*GetAdvertiserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Advertiser_ListAdvertiser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAdvertiserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertiserServer).ListAdvertiser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Advertiser_ListAdvertiser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertiserServer).ListAdvertiser(ctx, req.(*ListAdvertiserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Advertiser_ServiceDesc is the grpc.ServiceDesc for Advertiser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Advertiser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.common.advertiser.v1.Advertiser",
	HandlerType: (*AdvertiserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAdvertiser",
			Handler:    _Advertiser_CreateAdvertiser_Handler,
		},
		{
			MethodName: "UpdateAdvertiser",
			Handler:    _Advertiser_UpdateAdvertiser_Handler,
		},
		{
			MethodName: "DeleteAdvertiser",
			Handler:    _Advertiser_DeleteAdvertiser_Handler,
		},
		{
			MethodName: "GetAdvertiser",
			Handler:    _Advertiser_GetAdvertiser_Handler,
		},
		{
			MethodName: "ListAdvertiser",
			Handler:    _Advertiser_ListAdvertiser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/advertiser/v1/advertiser.proto",
}
