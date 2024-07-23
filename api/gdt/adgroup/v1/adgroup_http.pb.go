// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.26.1
// source: gdt/adgroup/v1/adgroup.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAdgroupCreateAdgroup = "/api.gdt.adgroup.v1.Adgroup/CreateAdgroup"
const OperationAdgroupGetAdgroup = "/api.gdt.adgroup.v1.Adgroup/GetAdgroup"
const OperationAdgroupListAdgroup = "/api.gdt.adgroup.v1.Adgroup/ListAdgroup"

type AdgroupHTTPServer interface {
	CreateAdgroup(context.Context, *CreateAdgroupRequest) (*CreateAdgroupReply, error)
	GetAdgroup(context.Context, *GetAdgroupRequest) (*GetAdgroupReply, error)
	ListAdgroup(context.Context, *ListAdgroupRequest) (*ListAdgroupReply, error)
}

func RegisterAdgroupHTTPServer(s *http.Server, srv AdgroupHTTPServer) {
	r := s.Route("/")
	r.POST("/gdt/v1/adgroup", _Adgroup_CreateAdgroup0_HTTP_Handler(srv))
	r.GET("/gdt/v1/adgroup", _Adgroup_GetAdgroup0_HTTP_Handler(srv))
	r.GET("/gdt/v1/adgroup/list", _Adgroup_ListAdgroup0_HTTP_Handler(srv))
}

func _Adgroup_CreateAdgroup0_HTTP_Handler(srv AdgroupHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateAdgroupRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdgroupCreateAdgroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAdgroup(ctx, req.(*CreateAdgroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateAdgroupReply)
		return ctx.Result(200, reply)
	}
}

func _Adgroup_GetAdgroup0_HTTP_Handler(srv AdgroupHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAdgroupRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdgroupGetAdgroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAdgroup(ctx, req.(*GetAdgroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAdgroupReply)
		return ctx.Result(200, reply)
	}
}

func _Adgroup_ListAdgroup0_HTTP_Handler(srv AdgroupHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListAdgroupRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdgroupListAdgroup)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListAdgroup(ctx, req.(*ListAdgroupRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListAdgroupReply)
		return ctx.Result(200, reply)
	}
}

type AdgroupHTTPClient interface {
	CreateAdgroup(ctx context.Context, req *CreateAdgroupRequest, opts ...http.CallOption) (rsp *CreateAdgroupReply, err error)
	GetAdgroup(ctx context.Context, req *GetAdgroupRequest, opts ...http.CallOption) (rsp *GetAdgroupReply, err error)
	ListAdgroup(ctx context.Context, req *ListAdgroupRequest, opts ...http.CallOption) (rsp *ListAdgroupReply, err error)
}

type AdgroupHTTPClientImpl struct {
	cc *http.Client
}

func NewAdgroupHTTPClient(client *http.Client) AdgroupHTTPClient {
	return &AdgroupHTTPClientImpl{client}
}

func (c *AdgroupHTTPClientImpl) CreateAdgroup(ctx context.Context, in *CreateAdgroupRequest, opts ...http.CallOption) (*CreateAdgroupReply, error) {
	var out CreateAdgroupReply
	pattern := "/gdt/v1/adgroup"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAdgroupCreateAdgroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AdgroupHTTPClientImpl) GetAdgroup(ctx context.Context, in *GetAdgroupRequest, opts ...http.CallOption) (*GetAdgroupReply, error) {
	var out GetAdgroupReply
	pattern := "/gdt/v1/adgroup"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAdgroupGetAdgroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *AdgroupHTTPClientImpl) ListAdgroup(ctx context.Context, in *ListAdgroupRequest, opts ...http.CallOption) (*ListAdgroupReply, error) {
	var out ListAdgroupReply
	pattern := "/gdt/v1/adgroup/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAdgroupListAdgroup))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}