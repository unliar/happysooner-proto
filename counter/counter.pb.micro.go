// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: counter/counter.proto

package counter

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for CounterSV service

func NewCounterSVEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CounterSV service

type CounterSVService interface {
	// 获取文章计数
	GetArticleCount(ctx context.Context, in *IntRequest, opts ...client.CallOption) (*ArticleCountInfo, error)
	// 获取用户计数
	GetUserCount(ctx context.Context, in *IntRequest, opts ...client.CallOption) (*UserCountInfo, error)
	// 增加用户计数
	UpdateUserCount(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*ErrorResponse, error)
	// 增加文章计数
	UpdateArticleCount(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*ErrorResponse, error)
}

type counterSVService struct {
	c    client.Client
	name string
}

func NewCounterSVService(name string, c client.Client) CounterSVService {
	return &counterSVService{
		c:    c,
		name: name,
	}
}

func (c *counterSVService) GetArticleCount(ctx context.Context, in *IntRequest, opts ...client.CallOption) (*ArticleCountInfo, error) {
	req := c.c.NewRequest(c.name, "CounterSV.GetArticleCount", in)
	out := new(ArticleCountInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterSVService) GetUserCount(ctx context.Context, in *IntRequest, opts ...client.CallOption) (*UserCountInfo, error) {
	req := c.c.NewRequest(c.name, "CounterSV.GetUserCount", in)
	out := new(UserCountInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterSVService) UpdateUserCount(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "CounterSV.UpdateUserCount", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterSVService) UpdateArticleCount(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "CounterSV.UpdateArticleCount", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CounterSV service

type CounterSVHandler interface {
	// 获取文章计数
	GetArticleCount(context.Context, *IntRequest, *ArticleCountInfo) error
	// 获取用户计数
	GetUserCount(context.Context, *IntRequest, *UserCountInfo) error
	// 增加用户计数
	UpdateUserCount(context.Context, *UpdateRequest, *ErrorResponse) error
	// 增加文章计数
	UpdateArticleCount(context.Context, *UpdateRequest, *ErrorResponse) error
}

func RegisterCounterSVHandler(s server.Server, hdlr CounterSVHandler, opts ...server.HandlerOption) error {
	type counterSV interface {
		GetArticleCount(ctx context.Context, in *IntRequest, out *ArticleCountInfo) error
		GetUserCount(ctx context.Context, in *IntRequest, out *UserCountInfo) error
		UpdateUserCount(ctx context.Context, in *UpdateRequest, out *ErrorResponse) error
		UpdateArticleCount(ctx context.Context, in *UpdateRequest, out *ErrorResponse) error
	}
	type CounterSV struct {
		counterSV
	}
	h := &counterSVHandler{hdlr}
	return s.Handle(s.NewHandler(&CounterSV{h}, opts...))
}

type counterSVHandler struct {
	CounterSVHandler
}

func (h *counterSVHandler) GetArticleCount(ctx context.Context, in *IntRequest, out *ArticleCountInfo) error {
	return h.CounterSVHandler.GetArticleCount(ctx, in, out)
}

func (h *counterSVHandler) GetUserCount(ctx context.Context, in *IntRequest, out *UserCountInfo) error {
	return h.CounterSVHandler.GetUserCount(ctx, in, out)
}

func (h *counterSVHandler) UpdateUserCount(ctx context.Context, in *UpdateRequest, out *ErrorResponse) error {
	return h.CounterSVHandler.UpdateUserCount(ctx, in, out)
}

func (h *counterSVHandler) UpdateArticleCount(ctx context.Context, in *UpdateRequest, out *ErrorResponse) error {
	return h.CounterSVHandler.UpdateArticleCount(ctx, in, out)
}
