// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: pay/pay.proto

package pay

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

// Client API for PaySV service

type PaySVService interface {
	// 获取用户支付数据信息
	GetPayInfo(ctx context.Context, in *UIDInput, opts ...client.CallOption) (*UserPayWays, error)
	// 新增一个支付数据类型
	CreateUserPayWay(ctx context.Context, in *CreateUserPayWayRequest, opts ...client.CallOption) (*ErrorResponse, error)
	// 删除一个支付数据类型
	DeleteUserPayWay(ctx context.Context, in *DeleteUserPayWayRequest, opts ...client.CallOption) (*ErrorResponse, error)
	// 更新一个支付数据类型
	PutUserPayWay(ctx context.Context, in *PutUserPayWayRequest, opts ...client.CallOption) (*ErrorResponse, error)
	// 创建alipay账单
	CreateAliPayOrder(ctx context.Context, in *CreateAliPayOrderRequest, opts ...client.CallOption) (*OrderResultResponse, error)
	// 查询订单结果
	GetAliPayOrderResult(ctx context.Context, in *PayOrderResultRequest, opts ...client.CallOption) (*OrderResultResponse, error)
	// 更新订单结果
	UpdateAliPayOrderResult(ctx context.Context, in *PayOrderResultRequest, opts ...client.CallOption) (*ErrorResponse, error)
}

type paySVService struct {
	c    client.Client
	name string
}

func NewPaySVService(name string, c client.Client) PaySVService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "pay"
	}
	return &paySVService{
		c:    c,
		name: name,
	}
}

func (c *paySVService) GetPayInfo(ctx context.Context, in *UIDInput, opts ...client.CallOption) (*UserPayWays, error) {
	req := c.c.NewRequest(c.name, "PaySV.GetPayInfo", in)
	out := new(UserPayWays)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paySVService) CreateUserPayWay(ctx context.Context, in *CreateUserPayWayRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "PaySV.CreateUserPayWay", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paySVService) DeleteUserPayWay(ctx context.Context, in *DeleteUserPayWayRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "PaySV.DeleteUserPayWay", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paySVService) PutUserPayWay(ctx context.Context, in *PutUserPayWayRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "PaySV.PutUserPayWay", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paySVService) CreateAliPayOrder(ctx context.Context, in *CreateAliPayOrderRequest, opts ...client.CallOption) (*OrderResultResponse, error) {
	req := c.c.NewRequest(c.name, "PaySV.CreateAliPayOrder", in)
	out := new(OrderResultResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paySVService) GetAliPayOrderResult(ctx context.Context, in *PayOrderResultRequest, opts ...client.CallOption) (*OrderResultResponse, error) {
	req := c.c.NewRequest(c.name, "PaySV.GetAliPayOrderResult", in)
	out := new(OrderResultResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paySVService) UpdateAliPayOrderResult(ctx context.Context, in *PayOrderResultRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "PaySV.UpdateAliPayOrderResult", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PaySV service

type PaySVHandler interface {
	// 获取用户支付数据信息
	GetPayInfo(context.Context, *UIDInput, *UserPayWays) error
	// 新增一个支付数据类型
	CreateUserPayWay(context.Context, *CreateUserPayWayRequest, *ErrorResponse) error
	// 删除一个支付数据类型
	DeleteUserPayWay(context.Context, *DeleteUserPayWayRequest, *ErrorResponse) error
	// 更新一个支付数据类型
	PutUserPayWay(context.Context, *PutUserPayWayRequest, *ErrorResponse) error
	// 创建alipay账单
	CreateAliPayOrder(context.Context, *CreateAliPayOrderRequest, *OrderResultResponse) error
	// 查询订单结果
	GetAliPayOrderResult(context.Context, *PayOrderResultRequest, *OrderResultResponse) error
	// 更新订单结果
	UpdateAliPayOrderResult(context.Context, *PayOrderResultRequest, *ErrorResponse) error
}

func RegisterPaySVHandler(s server.Server, hdlr PaySVHandler, opts ...server.HandlerOption) error {
	type paySV interface {
		GetPayInfo(ctx context.Context, in *UIDInput, out *UserPayWays) error
		CreateUserPayWay(ctx context.Context, in *CreateUserPayWayRequest, out *ErrorResponse) error
		DeleteUserPayWay(ctx context.Context, in *DeleteUserPayWayRequest, out *ErrorResponse) error
		PutUserPayWay(ctx context.Context, in *PutUserPayWayRequest, out *ErrorResponse) error
		CreateAliPayOrder(ctx context.Context, in *CreateAliPayOrderRequest, out *OrderResultResponse) error
		GetAliPayOrderResult(ctx context.Context, in *PayOrderResultRequest, out *OrderResultResponse) error
		UpdateAliPayOrderResult(ctx context.Context, in *PayOrderResultRequest, out *ErrorResponse) error
	}
	type PaySV struct {
		paySV
	}
	h := &paySVHandler{hdlr}
	return s.Handle(s.NewHandler(&PaySV{h}, opts...))
}

type paySVHandler struct {
	PaySVHandler
}

func (h *paySVHandler) GetPayInfo(ctx context.Context, in *UIDInput, out *UserPayWays) error {
	return h.PaySVHandler.GetPayInfo(ctx, in, out)
}

func (h *paySVHandler) CreateUserPayWay(ctx context.Context, in *CreateUserPayWayRequest, out *ErrorResponse) error {
	return h.PaySVHandler.CreateUserPayWay(ctx, in, out)
}

func (h *paySVHandler) DeleteUserPayWay(ctx context.Context, in *DeleteUserPayWayRequest, out *ErrorResponse) error {
	return h.PaySVHandler.DeleteUserPayWay(ctx, in, out)
}

func (h *paySVHandler) PutUserPayWay(ctx context.Context, in *PutUserPayWayRequest, out *ErrorResponse) error {
	return h.PaySVHandler.PutUserPayWay(ctx, in, out)
}

func (h *paySVHandler) CreateAliPayOrder(ctx context.Context, in *CreateAliPayOrderRequest, out *OrderResultResponse) error {
	return h.PaySVHandler.CreateAliPayOrder(ctx, in, out)
}

func (h *paySVHandler) GetAliPayOrderResult(ctx context.Context, in *PayOrderResultRequest, out *OrderResultResponse) error {
	return h.PaySVHandler.GetAliPayOrderResult(ctx, in, out)
}

func (h *paySVHandler) UpdateAliPayOrderResult(ctx context.Context, in *PayOrderResultRequest, out *ErrorResponse) error {
	return h.PaySVHandler.UpdateAliPayOrderResult(ctx, in, out)
}
