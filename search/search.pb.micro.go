// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: search/search.proto

package search

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for SearchSV service

type SearchSVService interface {
	// 获取搜索结果
	GetSearchByType(ctx context.Context, in *GetSearchRequest, opts ...client.CallOption) (*SearchResult, error)
	// Flush某个搜索Object
	FlushSearchByObject(ctx context.Context, in *FlushSearchRequest, opts ...client.CallOption) (*ErrorResponse, error)
	// Flush某个搜索Bucket
	FlushSearchByBucket(ctx context.Context, in *FlushSearchRequest, opts ...client.CallOption) (*ErrorResponse, error)
	// Flush某个搜索Collection
	FlushSearchByCollection(ctx context.Context, in *FlushSearchRequest, opts ...client.CallOption) (*ErrorResponse, error)
	// 插入搜索资源数据
	PostSearchSource(ctx context.Context, in *PostSearchSourceRequest, opts ...client.CallOption) (*ErrorResponse, error)
}

type searchSVService struct {
	c    client.Client
	name string
}

func NewSearchSVService(name string, c client.Client) SearchSVService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "search"
	}
	return &searchSVService{
		c:    c,
		name: name,
	}
}

func (c *searchSVService) GetSearchByType(ctx context.Context, in *GetSearchRequest, opts ...client.CallOption) (*SearchResult, error) {
	req := c.c.NewRequest(c.name, "SearchSV.GetSearchByType", in)
	out := new(SearchResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchSVService) FlushSearchByObject(ctx context.Context, in *FlushSearchRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "SearchSV.FlushSearchByObject", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchSVService) FlushSearchByBucket(ctx context.Context, in *FlushSearchRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "SearchSV.FlushSearchByBucket", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchSVService) FlushSearchByCollection(ctx context.Context, in *FlushSearchRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "SearchSV.FlushSearchByCollection", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchSVService) PostSearchSource(ctx context.Context, in *PostSearchSourceRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "SearchSV.PostSearchSource", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SearchSV service

type SearchSVHandler interface {
	// 获取搜索结果
	GetSearchByType(context.Context, *GetSearchRequest, *SearchResult) error
	// Flush某个搜索Object
	FlushSearchByObject(context.Context, *FlushSearchRequest, *ErrorResponse) error
	// Flush某个搜索Bucket
	FlushSearchByBucket(context.Context, *FlushSearchRequest, *ErrorResponse) error
	// Flush某个搜索Collection
	FlushSearchByCollection(context.Context, *FlushSearchRequest, *ErrorResponse) error
	// 插入搜索资源数据
	PostSearchSource(context.Context, *PostSearchSourceRequest, *ErrorResponse) error
}

func RegisterSearchSVHandler(s server.Server, hdlr SearchSVHandler, opts ...server.HandlerOption) error {
	type searchSV interface {
		GetSearchByType(ctx context.Context, in *GetSearchRequest, out *SearchResult) error
		FlushSearchByObject(ctx context.Context, in *FlushSearchRequest, out *ErrorResponse) error
		FlushSearchByBucket(ctx context.Context, in *FlushSearchRequest, out *ErrorResponse) error
		FlushSearchByCollection(ctx context.Context, in *FlushSearchRequest, out *ErrorResponse) error
		PostSearchSource(ctx context.Context, in *PostSearchSourceRequest, out *ErrorResponse) error
	}
	type SearchSV struct {
		searchSV
	}
	h := &searchSVHandler{hdlr}
	return s.Handle(s.NewHandler(&SearchSV{h}, opts...))
}

type searchSVHandler struct {
	SearchSVHandler
}

func (h *searchSVHandler) GetSearchByType(ctx context.Context, in *GetSearchRequest, out *SearchResult) error {
	return h.SearchSVHandler.GetSearchByType(ctx, in, out)
}

func (h *searchSVHandler) FlushSearchByObject(ctx context.Context, in *FlushSearchRequest, out *ErrorResponse) error {
	return h.SearchSVHandler.FlushSearchByObject(ctx, in, out)
}

func (h *searchSVHandler) FlushSearchByBucket(ctx context.Context, in *FlushSearchRequest, out *ErrorResponse) error {
	return h.SearchSVHandler.FlushSearchByBucket(ctx, in, out)
}

func (h *searchSVHandler) FlushSearchByCollection(ctx context.Context, in *FlushSearchRequest, out *ErrorResponse) error {
	return h.SearchSVHandler.FlushSearchByCollection(ctx, in, out)
}

func (h *searchSVHandler) PostSearchSource(ctx context.Context, in *PostSearchSourceRequest, out *ErrorResponse) error {
	return h.SearchSVHandler.PostSearchSource(ctx, in, out)
}
