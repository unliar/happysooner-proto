// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: push/push.proto

package push

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for PushSV service

func NewPushSVEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PushSV service

type PushSVService interface {
	// 发送验证码
	PushCaptcha(ctx context.Context, in *PushCaptchaRequest, opts ...client.CallOption) (*ErrorResponse, error)
	// 校验验证码
	GetCaptchaVerify(ctx context.Context, in *GetCaptchaVerifyRequest, opts ...client.CallOption) (*ErrorResponse, error)
	// 发送邮件通知
	PushEmailNotification(ctx context.Context, in *PushEmailNotificationRequest, opts ...client.CallOption) (*ErrorResponse, error)
	// 存储订单临时邮件信息
	StoreOrderMail(ctx context.Context, in *StoreOrderMailRequest, opts ...client.CallOption) (*ErrorResponse, error)
	// 消费订单邮件信息
	PostOrderMail(ctx context.Context, in *StoreOrderMailRequest, opts ...client.CallOption) (*ErrorResponse, error)
	// 发送微信机器人通知
	PushWechatWorkGroupRobotNotification(ctx context.Context, in *PushWechatWorkGroupRobotNotificationRequest, opts ...client.CallOption) (*ErrorResponse, error)
	// 发送公众号模板消息
	PushWechatTemplateMessage(ctx context.Context, in *PushWechatTemplateMessageRequest, opts ...client.CallOption) (*ErrorResponse, error)
	// 通用发送消息接口
	PushCommonMessage(ctx context.Context, in *PushCommonMessageRequest, opts ...client.CallOption) (*ErrorResponse, error)
}

type pushSVService struct {
	c    client.Client
	name string
}

func NewPushSVService(name string, c client.Client) PushSVService {
	return &pushSVService{
		c:    c,
		name: name,
	}
}

func (c *pushSVService) PushCaptcha(ctx context.Context, in *PushCaptchaRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "PushSV.PushCaptcha", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushSVService) GetCaptchaVerify(ctx context.Context, in *GetCaptchaVerifyRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "PushSV.GetCaptchaVerify", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushSVService) PushEmailNotification(ctx context.Context, in *PushEmailNotificationRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "PushSV.PushEmailNotification", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushSVService) StoreOrderMail(ctx context.Context, in *StoreOrderMailRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "PushSV.StoreOrderMail", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushSVService) PostOrderMail(ctx context.Context, in *StoreOrderMailRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "PushSV.PostOrderMail", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushSVService) PushWechatWorkGroupRobotNotification(ctx context.Context, in *PushWechatWorkGroupRobotNotificationRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "PushSV.PushWechatWorkGroupRobotNotification", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushSVService) PushWechatTemplateMessage(ctx context.Context, in *PushWechatTemplateMessageRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "PushSV.PushWechatTemplateMessage", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushSVService) PushCommonMessage(ctx context.Context, in *PushCommonMessageRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "PushSV.PushCommonMessage", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PushSV service

type PushSVHandler interface {
	// 发送验证码
	PushCaptcha(context.Context, *PushCaptchaRequest, *ErrorResponse) error
	// 校验验证码
	GetCaptchaVerify(context.Context, *GetCaptchaVerifyRequest, *ErrorResponse) error
	// 发送邮件通知
	PushEmailNotification(context.Context, *PushEmailNotificationRequest, *ErrorResponse) error
	// 存储订单临时邮件信息
	StoreOrderMail(context.Context, *StoreOrderMailRequest, *ErrorResponse) error
	// 消费订单邮件信息
	PostOrderMail(context.Context, *StoreOrderMailRequest, *ErrorResponse) error
	// 发送微信机器人通知
	PushWechatWorkGroupRobotNotification(context.Context, *PushWechatWorkGroupRobotNotificationRequest, *ErrorResponse) error
	// 发送公众号模板消息
	PushWechatTemplateMessage(context.Context, *PushWechatTemplateMessageRequest, *ErrorResponse) error
	// 通用发送消息接口
	PushCommonMessage(context.Context, *PushCommonMessageRequest, *ErrorResponse) error
}

func RegisterPushSVHandler(s server.Server, hdlr PushSVHandler, opts ...server.HandlerOption) error {
	type pushSV interface {
		PushCaptcha(ctx context.Context, in *PushCaptchaRequest, out *ErrorResponse) error
		GetCaptchaVerify(ctx context.Context, in *GetCaptchaVerifyRequest, out *ErrorResponse) error
		PushEmailNotification(ctx context.Context, in *PushEmailNotificationRequest, out *ErrorResponse) error
		StoreOrderMail(ctx context.Context, in *StoreOrderMailRequest, out *ErrorResponse) error
		PostOrderMail(ctx context.Context, in *StoreOrderMailRequest, out *ErrorResponse) error
		PushWechatWorkGroupRobotNotification(ctx context.Context, in *PushWechatWorkGroupRobotNotificationRequest, out *ErrorResponse) error
		PushWechatTemplateMessage(ctx context.Context, in *PushWechatTemplateMessageRequest, out *ErrorResponse) error
		PushCommonMessage(ctx context.Context, in *PushCommonMessageRequest, out *ErrorResponse) error
	}
	type PushSV struct {
		pushSV
	}
	h := &pushSVHandler{hdlr}
	return s.Handle(s.NewHandler(&PushSV{h}, opts...))
}

type pushSVHandler struct {
	PushSVHandler
}

func (h *pushSVHandler) PushCaptcha(ctx context.Context, in *PushCaptchaRequest, out *ErrorResponse) error {
	return h.PushSVHandler.PushCaptcha(ctx, in, out)
}

func (h *pushSVHandler) GetCaptchaVerify(ctx context.Context, in *GetCaptchaVerifyRequest, out *ErrorResponse) error {
	return h.PushSVHandler.GetCaptchaVerify(ctx, in, out)
}

func (h *pushSVHandler) PushEmailNotification(ctx context.Context, in *PushEmailNotificationRequest, out *ErrorResponse) error {
	return h.PushSVHandler.PushEmailNotification(ctx, in, out)
}

func (h *pushSVHandler) StoreOrderMail(ctx context.Context, in *StoreOrderMailRequest, out *ErrorResponse) error {
	return h.PushSVHandler.StoreOrderMail(ctx, in, out)
}

func (h *pushSVHandler) PostOrderMail(ctx context.Context, in *StoreOrderMailRequest, out *ErrorResponse) error {
	return h.PushSVHandler.PostOrderMail(ctx, in, out)
}

func (h *pushSVHandler) PushWechatWorkGroupRobotNotification(ctx context.Context, in *PushWechatWorkGroupRobotNotificationRequest, out *ErrorResponse) error {
	return h.PushSVHandler.PushWechatWorkGroupRobotNotification(ctx, in, out)
}

func (h *pushSVHandler) PushWechatTemplateMessage(ctx context.Context, in *PushWechatTemplateMessageRequest, out *ErrorResponse) error {
	return h.PushSVHandler.PushWechatTemplateMessage(ctx, in, out)
}

func (h *pushSVHandler) PushCommonMessage(ctx context.Context, in *PushCommonMessageRequest, out *ErrorResponse) error {
	return h.PushSVHandler.PushCommonMessage(ctx, in, out)
}
