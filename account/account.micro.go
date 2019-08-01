// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: account/account.proto

package account

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

// Client API for AccountSV service

type AccountSVService interface {
	// 查询用户信息
	GetUserInfo(ctx context.Context, in *UIDInput, opts ...client.CallOption) (*UserInfo, error)
	// 通过UserToken获取用户信息
	GetUserInfoByToken(ctx context.Context, in *UserToken, opts ...client.CallOption) (*UserInfo, error)
	// 创建用户
	CreateUserInfoByLoginName(ctx context.Context, in *CreateUserInput, opts ...client.CallOption) (*ErrorResponse, error)
	// 用户登录
	GetUserToken(ctx context.Context, in *UserLoginRequest, opts ...client.CallOption) (*UserInfoWithToken, error)
	// 检查昵称是否被使用
	CheckNickname(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*ErrorResponse, error)
	// 检查手机是否被绑定
	CheckPhone(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*ErrorResponse, error)
	// 检查邮箱是否被绑定
	CheckEmail(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*ErrorResponse, error)
	// 检查用户名是否被使用
	CheckLoginName(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*ErrorResponse, error)
	// 更新用户密码
	PutUserPassword(ctx context.Context, in *PutPassowrdRequest, opts ...client.CallOption) (*ErrorResponse, error)
	// 通过第三方code获取用户token
	GetUserTokenByOauthCode(ctx context.Context, in *OauthLoginRequest, opts ...client.CallOption) (*UserInfoWithToken, error)
	// 通过第三方code注册用户
	CreateUserInfoByOauthCode(ctx context.Context, in *OauthCreateUserRequest, opts ...client.CallOption) (*ErrorResponse, error)
	// 获取用户密码状态
	GetUserPasswordRandomStatus(ctx context.Context, in *UIDInput, opts ...client.CallOption) (*PasswordRandomStatus, error)
}

type accountSVService struct {
	c    client.Client
	name string
}

func NewAccountSVService(name string, c client.Client) AccountSVService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "account"
	}
	return &accountSVService{
		c:    c,
		name: name,
	}
}

func (c *accountSVService) GetUserInfo(ctx context.Context, in *UIDInput, opts ...client.CallOption) (*UserInfo, error) {
	req := c.c.NewRequest(c.name, "AccountSV.GetUserInfo", in)
	out := new(UserInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountSVService) GetUserInfoByToken(ctx context.Context, in *UserToken, opts ...client.CallOption) (*UserInfo, error) {
	req := c.c.NewRequest(c.name, "AccountSV.GetUserInfoByToken", in)
	out := new(UserInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountSVService) CreateUserInfoByLoginName(ctx context.Context, in *CreateUserInput, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "AccountSV.CreateUserInfoByLoginName", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountSVService) GetUserToken(ctx context.Context, in *UserLoginRequest, opts ...client.CallOption) (*UserInfoWithToken, error) {
	req := c.c.NewRequest(c.name, "AccountSV.GetUserToken", in)
	out := new(UserInfoWithToken)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountSVService) CheckNickname(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "AccountSV.CheckNickname", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountSVService) CheckPhone(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "AccountSV.CheckPhone", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountSVService) CheckEmail(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "AccountSV.CheckEmail", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountSVService) CheckLoginName(ctx context.Context, in *UserInfo, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "AccountSV.CheckLoginName", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountSVService) PutUserPassword(ctx context.Context, in *PutPassowrdRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "AccountSV.PutUserPassword", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountSVService) GetUserTokenByOauthCode(ctx context.Context, in *OauthLoginRequest, opts ...client.CallOption) (*UserInfoWithToken, error) {
	req := c.c.NewRequest(c.name, "AccountSV.GetUserTokenByOauthCode", in)
	out := new(UserInfoWithToken)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountSVService) CreateUserInfoByOauthCode(ctx context.Context, in *OauthCreateUserRequest, opts ...client.CallOption) (*ErrorResponse, error) {
	req := c.c.NewRequest(c.name, "AccountSV.CreateUserInfoByOauthCode", in)
	out := new(ErrorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountSVService) GetUserPasswordRandomStatus(ctx context.Context, in *UIDInput, opts ...client.CallOption) (*PasswordRandomStatus, error) {
	req := c.c.NewRequest(c.name, "AccountSV.GetUserPasswordRandomStatus", in)
	out := new(PasswordRandomStatus)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AccountSV service

type AccountSVHandler interface {
	// 查询用户信息
	GetUserInfo(context.Context, *UIDInput, *UserInfo) error
	// 通过UserToken获取用户信息
	GetUserInfoByToken(context.Context, *UserToken, *UserInfo) error
	// 创建用户
	CreateUserInfoByLoginName(context.Context, *CreateUserInput, *ErrorResponse) error
	// 用户登录
	GetUserToken(context.Context, *UserLoginRequest, *UserInfoWithToken) error
	// 检查昵称是否被使用
	CheckNickname(context.Context, *UserInfo, *ErrorResponse) error
	// 检查手机是否被绑定
	CheckPhone(context.Context, *UserInfo, *ErrorResponse) error
	// 检查邮箱是否被绑定
	CheckEmail(context.Context, *UserInfo, *ErrorResponse) error
	// 检查用户名是否被使用
	CheckLoginName(context.Context, *UserInfo, *ErrorResponse) error
	// 更新用户密码
	PutUserPassword(context.Context, *PutPassowrdRequest, *ErrorResponse) error
	// 通过第三方code获取用户token
	GetUserTokenByOauthCode(context.Context, *OauthLoginRequest, *UserInfoWithToken) error
	// 通过第三方code注册用户
	CreateUserInfoByOauthCode(context.Context, *OauthCreateUserRequest, *ErrorResponse) error
	// 获取用户密码状态
	GetUserPasswordRandomStatus(context.Context, *UIDInput, *PasswordRandomStatus) error
}

func RegisterAccountSVHandler(s server.Server, hdlr AccountSVHandler, opts ...server.HandlerOption) error {
	type accountSV interface {
		GetUserInfo(ctx context.Context, in *UIDInput, out *UserInfo) error
		GetUserInfoByToken(ctx context.Context, in *UserToken, out *UserInfo) error
		CreateUserInfoByLoginName(ctx context.Context, in *CreateUserInput, out *ErrorResponse) error
		GetUserToken(ctx context.Context, in *UserLoginRequest, out *UserInfoWithToken) error
		CheckNickname(ctx context.Context, in *UserInfo, out *ErrorResponse) error
		CheckPhone(ctx context.Context, in *UserInfo, out *ErrorResponse) error
		CheckEmail(ctx context.Context, in *UserInfo, out *ErrorResponse) error
		CheckLoginName(ctx context.Context, in *UserInfo, out *ErrorResponse) error
		PutUserPassword(ctx context.Context, in *PutPassowrdRequest, out *ErrorResponse) error
		GetUserTokenByOauthCode(ctx context.Context, in *OauthLoginRequest, out *UserInfoWithToken) error
		CreateUserInfoByOauthCode(ctx context.Context, in *OauthCreateUserRequest, out *ErrorResponse) error
		GetUserPasswordRandomStatus(ctx context.Context, in *UIDInput, out *PasswordRandomStatus) error
	}
	type AccountSV struct {
		accountSV
	}
	h := &accountSVHandler{hdlr}
	return s.Handle(s.NewHandler(&AccountSV{h}, opts...))
}

type accountSVHandler struct {
	AccountSVHandler
}

func (h *accountSVHandler) GetUserInfo(ctx context.Context, in *UIDInput, out *UserInfo) error {
	return h.AccountSVHandler.GetUserInfo(ctx, in, out)
}

func (h *accountSVHandler) GetUserInfoByToken(ctx context.Context, in *UserToken, out *UserInfo) error {
	return h.AccountSVHandler.GetUserInfoByToken(ctx, in, out)
}

func (h *accountSVHandler) CreateUserInfoByLoginName(ctx context.Context, in *CreateUserInput, out *ErrorResponse) error {
	return h.AccountSVHandler.CreateUserInfoByLoginName(ctx, in, out)
}

func (h *accountSVHandler) GetUserToken(ctx context.Context, in *UserLoginRequest, out *UserInfoWithToken) error {
	return h.AccountSVHandler.GetUserToken(ctx, in, out)
}

func (h *accountSVHandler) CheckNickname(ctx context.Context, in *UserInfo, out *ErrorResponse) error {
	return h.AccountSVHandler.CheckNickname(ctx, in, out)
}

func (h *accountSVHandler) CheckPhone(ctx context.Context, in *UserInfo, out *ErrorResponse) error {
	return h.AccountSVHandler.CheckPhone(ctx, in, out)
}

func (h *accountSVHandler) CheckEmail(ctx context.Context, in *UserInfo, out *ErrorResponse) error {
	return h.AccountSVHandler.CheckEmail(ctx, in, out)
}

func (h *accountSVHandler) CheckLoginName(ctx context.Context, in *UserInfo, out *ErrorResponse) error {
	return h.AccountSVHandler.CheckLoginName(ctx, in, out)
}

func (h *accountSVHandler) PutUserPassword(ctx context.Context, in *PutPassowrdRequest, out *ErrorResponse) error {
	return h.AccountSVHandler.PutUserPassword(ctx, in, out)
}

func (h *accountSVHandler) GetUserTokenByOauthCode(ctx context.Context, in *OauthLoginRequest, out *UserInfoWithToken) error {
	return h.AccountSVHandler.GetUserTokenByOauthCode(ctx, in, out)
}

func (h *accountSVHandler) CreateUserInfoByOauthCode(ctx context.Context, in *OauthCreateUserRequest, out *ErrorResponse) error {
	return h.AccountSVHandler.CreateUserInfoByOauthCode(ctx, in, out)
}

func (h *accountSVHandler) GetUserPasswordRandomStatus(ctx context.Context, in *UIDInput, out *PasswordRandomStatus) error {
	return h.AccountSVHandler.GetUserPasswordRandomStatus(ctx, in, out)
}
