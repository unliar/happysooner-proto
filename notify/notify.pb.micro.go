// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: notify/notify.proto

package nofify

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

// Api Endpoints for NofifySV service

func NewNofifySVEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for NofifySV service

type NofifySVService interface {
	SendChatMessage(ctx context.Context, in *SendChatMessageRequest, opts ...client.CallOption) (*ErrorResponse, error)
}

type nofifySVService struct {
	c    client.Client
	name string
}

func NewNofifySVService(name string, c client.Client) NofifySVService {
	return &nofifySVService{
		c:    c,
		name: name,
	}
}

func (c *nofifySVService) SendChatMessage(ctx context.Context, in *SendChatMessageRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "NofifySV.SendChatMessage", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NofifySV service

type NofifySVHandler interface {
	SendChatMessage(context.Context, *SendChatMessageRequest, *ErrorResponse) error
}

func RegisterNofifySVHandler(s server.Server, hdlr NofifySVHandler, opts ...server.HandlerOption) error {
	type nofifySV interface {
		SendChatMessage(ctx context.Context, in *SendChatMessageRequest, out *ErrorResponse) error
	}
	type NofifySV struct {
		nofifySV
	}
	h := &nofifySVHandler{hdlr}
	return s.Handle(s.NewHandler(&NofifySV{h}, opts...))
}

type nofifySVHandler struct {
	NofifySVHandler
}

func (h *nofifySVHandler) SendChatMessage(ctx context.Context, in *SendChatMessageRequest, out *ErrorResponse) error {
	return h.NofifySVHandler.SendChatMessage(ctx, in, out)
}
