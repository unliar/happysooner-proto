// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account/account.proto

package account

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// 错误码定义 累计7位 账户服务使用 100 [第一位1表示账户模块，第二位表示子模块(表)] ，后四位自由发挥qaq
// 100
type AccountErrors int32

const (
	AccountErrors_None                           AccountErrors = 0
	AccountErrors_UserNotFound                   AccountErrors = 1000001
	AccountErrors_UserCreateFailed               AccountErrors = 1000002
	AccountErrors_UserUpdateFailed               AccountErrors = 1000003
	AccountErrors_UserNicknameExist              AccountErrors = 1000004
	AccountErrors_UserLoginNameExist             AccountErrors = 1000005
	AccountErrors_UserPhoneExist                 AccountErrors = 1000006
	AccountErrors_UserEmailExist                 AccountErrors = 1000007
	AccountErrors_UserPasswordCheckFailed        AccountErrors = 1000008
	AccountErrors_UserBindUserContactEmailFailed AccountErrors = 1000009
	AccountErrors_UserBindUserContactPhoneFailed AccountErrors = 1000010
)

var AccountErrors_name = map[int32]string{
	0:       "None",
	1000001: "UserNotFound",
	1000002: "UserCreateFailed",
	1000003: "UserUpdateFailed",
	1000004: "UserNicknameExist",
	1000005: "UserLoginNameExist",
	1000006: "UserPhoneExist",
	1000007: "UserEmailExist",
	1000008: "UserPasswordCheckFailed",
	1000009: "UserBindUserContactEmailFailed",
	1000010: "UserBindUserContactPhoneFailed",
}

var AccountErrors_value = map[string]int32{
	"None":                           0,
	"UserNotFound":                   1000001,
	"UserCreateFailed":               1000002,
	"UserUpdateFailed":               1000003,
	"UserNicknameExist":              1000004,
	"UserLoginNameExist":             1000005,
	"UserPhoneExist":                 1000006,
	"UserEmailExist":                 1000007,
	"UserPasswordCheckFailed":        1000008,
	"UserBindUserContactEmailFailed": 1000009,
	"UserBindUserContactPhoneFailed": 1000010,
}

func (x AccountErrors) String() string {
	return proto.EnumName(AccountErrors_name, int32(x))
}

func (AccountErrors) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{0}
}

type UserInfo struct {
	Id                   int64       `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	LoginName            string      `protobuf:"bytes,2,opt,name=LoginName,proto3" json:"LoginName,omitempty"`
	Nickname             string      `protobuf:"bytes,3,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	Age                  int64       `protobuf:"varint,4,opt,name=Age,proto3" json:"Age,omitempty"`
	Gender               int64       `protobuf:"varint,5,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Avatar               string      `protobuf:"bytes,6,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Location             string      `protobuf:"bytes,7,opt,name=Location,proto3" json:"Location,omitempty"`
	Profession           string      `protobuf:"bytes,8,opt,name=Profession,proto3" json:"Profession,omitempty"`
	Status               int64       `protobuf:"varint,9,opt,name=Status,proto3" json:"Status,omitempty"`
	Brief                string      `protobuf:"bytes,12,opt,name=Brief,proto3" json:"Brief,omitempty"`
	Email                string      `protobuf:"bytes,13,opt,name=Email,proto3" json:"Email,omitempty"`
	Phone                string      `protobuf:"bytes,14,opt,name=Phone,proto3" json:"Phone,omitempty"`
	PhoneCode            string      `protobuf:"bytes,15,opt,name=PhoneCode,proto3" json:"PhoneCode,omitempty"`
	WeChatId             string      `protobuf:"bytes,16,opt,name=WeChatId,proto3" json:"WeChatId,omitempty"`
	QQId                 string      `protobuf:"bytes,17,opt,name=QQId,proto3" json:"QQId,omitempty"`
	Password             string      `protobuf:"bytes,18,opt,name=Password,proto3" json:"Password,omitempty"`
	Roles                []*RoleInfo `protobuf:"bytes,19,rep,name=Roles,proto3" json:"Roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{0}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfo) GetLoginName() string {
	if m != nil {
		return m.LoginName
	}
	return ""
}

func (m *UserInfo) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UserInfo) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserInfo) GetGender() int64 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *UserInfo) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *UserInfo) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *UserInfo) GetProfession() string {
	if m != nil {
		return m.Profession
	}
	return ""
}

func (m *UserInfo) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UserInfo) GetBrief() string {
	if m != nil {
		return m.Brief
	}
	return ""
}

func (m *UserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *UserInfo) GetPhoneCode() string {
	if m != nil {
		return m.PhoneCode
	}
	return ""
}

func (m *UserInfo) GetWeChatId() string {
	if m != nil {
		return m.WeChatId
	}
	return ""
}

func (m *UserInfo) GetQQId() string {
	if m != nil {
		return m.QQId
	}
	return ""
}

func (m *UserInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserInfo) GetRoles() []*RoleInfo {
	if m != nil {
		return m.Roles
	}
	return nil
}

type PutUserBaseRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	LoginName            string   `protobuf:"bytes,2,opt,name=LoginName,proto3" json:"LoginName,omitempty"`
	Nickname             string   `protobuf:"bytes,3,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	Age                  int64    `protobuf:"varint,4,opt,name=Age,proto3" json:"Age,omitempty"`
	Gender               int64    `protobuf:"varint,5,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Avatar               string   `protobuf:"bytes,6,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Location             string   `protobuf:"bytes,7,opt,name=Location,proto3" json:"Location,omitempty"`
	Profession           string   `protobuf:"bytes,8,opt,name=Profession,proto3" json:"Profession,omitempty"`
	Brief                string   `protobuf:"bytes,9,opt,name=Brief,proto3" json:"Brief,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutUserBaseRequest) Reset()         { *m = PutUserBaseRequest{} }
func (m *PutUserBaseRequest) String() string { return proto.CompactTextString(m) }
func (*PutUserBaseRequest) ProtoMessage()    {}
func (*PutUserBaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{1}
}

func (m *PutUserBaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutUserBaseRequest.Unmarshal(m, b)
}
func (m *PutUserBaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutUserBaseRequest.Marshal(b, m, deterministic)
}
func (m *PutUserBaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutUserBaseRequest.Merge(m, src)
}
func (m *PutUserBaseRequest) XXX_Size() int {
	return xxx_messageInfo_PutUserBaseRequest.Size(m)
}
func (m *PutUserBaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutUserBaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutUserBaseRequest proto.InternalMessageInfo

func (m *PutUserBaseRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PutUserBaseRequest) GetLoginName() string {
	if m != nil {
		return m.LoginName
	}
	return ""
}

func (m *PutUserBaseRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *PutUserBaseRequest) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *PutUserBaseRequest) GetGender() int64 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *PutUserBaseRequest) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *PutUserBaseRequest) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *PutUserBaseRequest) GetProfession() string {
	if m != nil {
		return m.Profession
	}
	return ""
}

func (m *PutUserBaseRequest) GetBrief() string {
	if m != nil {
		return m.Brief
	}
	return ""
}

type UserInfoWithToken struct {
	UserInfo             *UserInfo `protobuf:"bytes,1,opt,name=UserInfo,proto3" json:"UserInfo,omitempty"`
	Token                string    `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UserInfoWithToken) Reset()         { *m = UserInfoWithToken{} }
func (m *UserInfoWithToken) String() string { return proto.CompactTextString(m) }
func (*UserInfoWithToken) ProtoMessage()    {}
func (*UserInfoWithToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{2}
}

func (m *UserInfoWithToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoWithToken.Unmarshal(m, b)
}
func (m *UserInfoWithToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoWithToken.Marshal(b, m, deterministic)
}
func (m *UserInfoWithToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoWithToken.Merge(m, src)
}
func (m *UserInfoWithToken) XXX_Size() int {
	return xxx_messageInfo_UserInfoWithToken.Size(m)
}
func (m *UserInfoWithToken) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoWithToken.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoWithToken proto.InternalMessageInfo

func (m *UserInfoWithToken) GetUserInfo() *UserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func (m *UserInfoWithToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// id查询用户信息
type UIDInput struct {
	UID                  int64    `protobuf:"varint,1,opt,name=UID,proto3" json:"UID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UIDInput) Reset()         { *m = UIDInput{} }
func (m *UIDInput) String() string { return proto.CompactTextString(m) }
func (*UIDInput) ProtoMessage()    {}
func (*UIDInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{3}
}

func (m *UIDInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UIDInput.Unmarshal(m, b)
}
func (m *UIDInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UIDInput.Marshal(b, m, deterministic)
}
func (m *UIDInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UIDInput.Merge(m, src)
}
func (m *UIDInput) XXX_Size() int {
	return xxx_messageInfo_UIDInput.Size(m)
}
func (m *UIDInput) XXX_DiscardUnknown() {
	xxx_messageInfo_UIDInput.DiscardUnknown(m)
}

var xxx_messageInfo_UIDInput proto.InternalMessageInfo

func (m *UIDInput) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

type RoleInfo struct {
	Role                 string   `protobuf:"bytes,1,opt,name=Role,proto3" json:"Role,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleInfo) Reset()         { *m = RoleInfo{} }
func (m *RoleInfo) String() string { return proto.CompactTextString(m) }
func (*RoleInfo) ProtoMessage()    {}
func (*RoleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{4}
}

func (m *RoleInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleInfo.Unmarshal(m, b)
}
func (m *RoleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleInfo.Marshal(b, m, deterministic)
}
func (m *RoleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleInfo.Merge(m, src)
}
func (m *RoleInfo) XXX_Size() int {
	return xxx_messageInfo_RoleInfo.Size(m)
}
func (m *RoleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RoleInfo proto.InternalMessageInfo

func (m *RoleInfo) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *RoleInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// 第三方登录请求
type OauthLoginRequest struct {
	Type                 string   `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OauthLoginRequest) Reset()         { *m = OauthLoginRequest{} }
func (m *OauthLoginRequest) String() string { return proto.CompactTextString(m) }
func (*OauthLoginRequest) ProtoMessage()    {}
func (*OauthLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{5}
}

func (m *OauthLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OauthLoginRequest.Unmarshal(m, b)
}
func (m *OauthLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OauthLoginRequest.Marshal(b, m, deterministic)
}
func (m *OauthLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OauthLoginRequest.Merge(m, src)
}
func (m *OauthLoginRequest) XXX_Size() int {
	return xxx_messageInfo_OauthLoginRequest.Size(m)
}
func (m *OauthLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OauthLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OauthLoginRequest proto.InternalMessageInfo

func (m *OauthLoginRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *OauthLoginRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type ErrorResponse struct {
	ErrorInfo            AccountErrors `protobuf:"varint,1,opt,name=ErrorInfo,proto3,enum=account.AccountErrors" json:"ErrorInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ErrorResponse) Reset()         { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()    {}
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{6}
}

func (m *ErrorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorResponse.Unmarshal(m, b)
}
func (m *ErrorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorResponse.Marshal(b, m, deterministic)
}
func (m *ErrorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorResponse.Merge(m, src)
}
func (m *ErrorResponse) XXX_Size() int {
	return xxx_messageInfo_ErrorResponse.Size(m)
}
func (m *ErrorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorResponse proto.InternalMessageInfo

func (m *ErrorResponse) GetErrorInfo() AccountErrors {
	if m != nil {
		return m.ErrorInfo
	}
	return AccountErrors_None
}

// 创建用户
type CreateUserInput struct {
	Type                 string   `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Nickname             string   `protobuf:"bytes,7,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	Extension            string   `protobuf:"bytes,4,opt,name=Extension,proto3" json:"Extension,omitempty"`
	IP                   string   `protobuf:"bytes,5,opt,name=IP,proto3" json:"IP,omitempty"`
	RegisterFrom         string   `protobuf:"bytes,6,opt,name=RegisterFrom,proto3" json:"RegisterFrom,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserInput) Reset()         { *m = CreateUserInput{} }
func (m *CreateUserInput) String() string { return proto.CompactTextString(m) }
func (*CreateUserInput) ProtoMessage()    {}
func (*CreateUserInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{7}
}

func (m *CreateUserInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserInput.Unmarshal(m, b)
}
func (m *CreateUserInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserInput.Marshal(b, m, deterministic)
}
func (m *CreateUserInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserInput.Merge(m, src)
}
func (m *CreateUserInput) XXX_Size() int {
	return xxx_messageInfo_CreateUserInput.Size(m)
}
func (m *CreateUserInput) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserInput.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserInput proto.InternalMessageInfo

func (m *CreateUserInput) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateUserInput) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *CreateUserInput) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateUserInput) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *CreateUserInput) GetExtension() string {
	if m != nil {
		return m.Extension
	}
	return ""
}

func (m *CreateUserInput) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *CreateUserInput) GetRegisterFrom() string {
	if m != nil {
		return m.RegisterFrom
	}
	return ""
}

// 用户登录
type UserLoginRequest struct {
	Type                 string   `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Validate             bool     `protobuf:"varint,4,opt,name=Validate,proto3" json:"Validate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLoginRequest) Reset()         { *m = UserLoginRequest{} }
func (m *UserLoginRequest) String() string { return proto.CompactTextString(m) }
func (*UserLoginRequest) ProtoMessage()    {}
func (*UserLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{8}
}

func (m *UserLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginRequest.Unmarshal(m, b)
}
func (m *UserLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginRequest.Marshal(b, m, deterministic)
}
func (m *UserLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginRequest.Merge(m, src)
}
func (m *UserLoginRequest) XXX_Size() int {
	return xxx_messageInfo_UserLoginRequest.Size(m)
}
func (m *UserLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginRequest proto.InternalMessageInfo

func (m *UserLoginRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *UserLoginRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *UserLoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserLoginRequest) GetValidate() bool {
	if m != nil {
		return m.Validate
	}
	return false
}

type UserToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserToken) Reset()         { *m = UserToken{} }
func (m *UserToken) String() string { return proto.CompactTextString(m) }
func (*UserToken) ProtoMessage()    {}
func (*UserToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{9}
}

func (m *UserToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserToken.Unmarshal(m, b)
}
func (m *UserToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserToken.Marshal(b, m, deterministic)
}
func (m *UserToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserToken.Merge(m, src)
}
func (m *UserToken) XXX_Size() int {
	return xxx_messageInfo_UserToken.Size(m)
}
func (m *UserToken) XXX_DiscardUnknown() {
	xxx_messageInfo_UserToken.DiscardUnknown(m)
}

var xxx_messageInfo_UserToken proto.InternalMessageInfo

func (m *UserToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// 更新用户密码
type PutPassowrdRequest struct {
	UID                  int64    `protobuf:"varint,1,opt,name=UID,proto3" json:"UID,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	FreshPassword        string   `protobuf:"bytes,3,opt,name=FreshPassword,proto3" json:"FreshPassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutPassowrdRequest) Reset()         { *m = PutPassowrdRequest{} }
func (m *PutPassowrdRequest) String() string { return proto.CompactTextString(m) }
func (*PutPassowrdRequest) ProtoMessage()    {}
func (*PutPassowrdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{10}
}

func (m *PutPassowrdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutPassowrdRequest.Unmarshal(m, b)
}
func (m *PutPassowrdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutPassowrdRequest.Marshal(b, m, deterministic)
}
func (m *PutPassowrdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutPassowrdRequest.Merge(m, src)
}
func (m *PutPassowrdRequest) XXX_Size() int {
	return xxx_messageInfo_PutPassowrdRequest.Size(m)
}
func (m *PutPassowrdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutPassowrdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutPassowrdRequest proto.InternalMessageInfo

func (m *PutPassowrdRequest) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *PutPassowrdRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *PutPassowrdRequest) GetFreshPassword() string {
	if m != nil {
		return m.FreshPassword
	}
	return ""
}

// 获取用户密码是否是随机生成的
type PasswordRandomStatus struct {
	Random               int64    `protobuf:"varint,1,opt,name=Random,proto3" json:"Random,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PasswordRandomStatus) Reset()         { *m = PasswordRandomStatus{} }
func (m *PasswordRandomStatus) String() string { return proto.CompactTextString(m) }
func (*PasswordRandomStatus) ProtoMessage()    {}
func (*PasswordRandomStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{11}
}

func (m *PasswordRandomStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasswordRandomStatus.Unmarshal(m, b)
}
func (m *PasswordRandomStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasswordRandomStatus.Marshal(b, m, deterministic)
}
func (m *PasswordRandomStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasswordRandomStatus.Merge(m, src)
}
func (m *PasswordRandomStatus) XXX_Size() int {
	return xxx_messageInfo_PasswordRandomStatus.Size(m)
}
func (m *PasswordRandomStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_PasswordRandomStatus.DiscardUnknown(m)
}

var xxx_messageInfo_PasswordRandomStatus proto.InternalMessageInfo

func (m *PasswordRandomStatus) GetRandom() int64 {
	if m != nil {
		return m.Random
	}
	return 0
}

type PutUserContactEmailRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	UID                  int64    `protobuf:"varint,2,opt,name=UID,proto3" json:"UID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutUserContactEmailRequest) Reset()         { *m = PutUserContactEmailRequest{} }
func (m *PutUserContactEmailRequest) String() string { return proto.CompactTextString(m) }
func (*PutUserContactEmailRequest) ProtoMessage()    {}
func (*PutUserContactEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{12}
}

func (m *PutUserContactEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutUserContactEmailRequest.Unmarshal(m, b)
}
func (m *PutUserContactEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutUserContactEmailRequest.Marshal(b, m, deterministic)
}
func (m *PutUserContactEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutUserContactEmailRequest.Merge(m, src)
}
func (m *PutUserContactEmailRequest) XXX_Size() int {
	return xxx_messageInfo_PutUserContactEmailRequest.Size(m)
}
func (m *PutUserContactEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutUserContactEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutUserContactEmailRequest proto.InternalMessageInfo

func (m *PutUserContactEmailRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *PutUserContactEmailRequest) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

type PutUserContactPhoneRequest struct {
	Phone                string   `protobuf:"bytes,1,opt,name=Phone,proto3" json:"Phone,omitempty"`
	PhoneCode            string   `protobuf:"bytes,2,opt,name=PhoneCode,proto3" json:"PhoneCode,omitempty"`
	UID                  int64    `protobuf:"varint,3,opt,name=UID,proto3" json:"UID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutUserContactPhoneRequest) Reset()         { *m = PutUserContactPhoneRequest{} }
func (m *PutUserContactPhoneRequest) String() string { return proto.CompactTextString(m) }
func (*PutUserContactPhoneRequest) ProtoMessage()    {}
func (*PutUserContactPhoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{13}
}

func (m *PutUserContactPhoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutUserContactPhoneRequest.Unmarshal(m, b)
}
func (m *PutUserContactPhoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutUserContactPhoneRequest.Marshal(b, m, deterministic)
}
func (m *PutUserContactPhoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutUserContactPhoneRequest.Merge(m, src)
}
func (m *PutUserContactPhoneRequest) XXX_Size() int {
	return xxx_messageInfo_PutUserContactPhoneRequest.Size(m)
}
func (m *PutUserContactPhoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutUserContactPhoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutUserContactPhoneRequest proto.InternalMessageInfo

func (m *PutUserContactPhoneRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *PutUserContactPhoneRequest) GetPhoneCode() string {
	if m != nil {
		return m.PhoneCode
	}
	return ""
}

func (m *PutUserContactPhoneRequest) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

type UserContact struct {
	Phone                string   `protobuf:"bytes,1,opt,name=Phone,proto3" json:"Phone,omitempty"`
	PhoneCode            string   `protobuf:"bytes,2,opt,name=PhoneCode,proto3" json:"PhoneCode,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	WeChatId             string   `protobuf:"bytes,4,opt,name=WeChatId,proto3" json:"WeChatId,omitempty"`
	QQId                 string   `protobuf:"bytes,5,opt,name=QQId,proto3" json:"QQId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserContact) Reset()         { *m = UserContact{} }
func (m *UserContact) String() string { return proto.CompactTextString(m) }
func (*UserContact) ProtoMessage()    {}
func (*UserContact) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{14}
}

func (m *UserContact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserContact.Unmarshal(m, b)
}
func (m *UserContact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserContact.Marshal(b, m, deterministic)
}
func (m *UserContact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserContact.Merge(m, src)
}
func (m *UserContact) XXX_Size() int {
	return xxx_messageInfo_UserContact.Size(m)
}
func (m *UserContact) XXX_DiscardUnknown() {
	xxx_messageInfo_UserContact.DiscardUnknown(m)
}

var xxx_messageInfo_UserContact proto.InternalMessageInfo

func (m *UserContact) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *UserContact) GetPhoneCode() string {
	if m != nil {
		return m.PhoneCode
	}
	return ""
}

func (m *UserContact) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserContact) GetWeChatId() string {
	if m != nil {
		return m.WeChatId
	}
	return ""
}

func (m *UserContact) GetQQId() string {
	if m != nil {
		return m.QQId
	}
	return ""
}

func init() {
	proto.RegisterEnum("account.AccountErrors", AccountErrors_name, AccountErrors_value)
	proto.RegisterType((*UserInfo)(nil), "account.UserInfo")
	proto.RegisterType((*PutUserBaseRequest)(nil), "account.PutUserBaseRequest")
	proto.RegisterType((*UserInfoWithToken)(nil), "account.UserInfoWithToken")
	proto.RegisterType((*UIDInput)(nil), "account.UIDInput")
	proto.RegisterType((*RoleInfo)(nil), "account.RoleInfo")
	proto.RegisterType((*OauthLoginRequest)(nil), "account.OauthLoginRequest")
	proto.RegisterType((*ErrorResponse)(nil), "account.ErrorResponse")
	proto.RegisterType((*CreateUserInput)(nil), "account.CreateUserInput")
	proto.RegisterType((*UserLoginRequest)(nil), "account.UserLoginRequest")
	proto.RegisterType((*UserToken)(nil), "account.UserToken")
	proto.RegisterType((*PutPassowrdRequest)(nil), "account.PutPassowrdRequest")
	proto.RegisterType((*PasswordRandomStatus)(nil), "account.PasswordRandomStatus")
	proto.RegisterType((*PutUserContactEmailRequest)(nil), "account.PutUserContactEmailRequest")
	proto.RegisterType((*PutUserContactPhoneRequest)(nil), "account.PutUserContactPhoneRequest")
	proto.RegisterType((*UserContact)(nil), "account.UserContact")
}

func init() { proto.RegisterFile("account/account.proto", fileDescriptor_d66906c5773c9d08) }

var fileDescriptor_d66906c5773c9d08 = []byte{
	// 1081 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x57, 0x5b, 0x4f, 0x1b, 0x47,
	0x14, 0xc6, 0x37, 0xf0, 0x1e, 0xb0, 0x31, 0x53, 0x17, 0x36, 0x0e, 0x89, 0xe8, 0x36, 0x52, 0x51,
	0xa5, 0x52, 0x89, 0xb6, 0x0f, 0x55, 0x6b, 0xa9, 0x60, 0x0c, 0x5d, 0x29, 0x32, 0xce, 0x12, 0x48,
	0x5f, 0x37, 0xde, 0x01, 0xaf, 0x30, 0x3b, 0xee, 0xee, 0xb8, 0x09, 0xef, 0x96, 0xfa, 0x53, 0xfa,
	0xdc, 0x5f, 0x50, 0xf5, 0x92, 0xde, 0x7e, 0x4e, 0x7f, 0x41, 0x35, 0x67, 0x66, 0x67, 0x2f, 0xb6,
	0x93, 0x86, 0xc7, 0x3e, 0x31, 0xe7, 0x9b, 0x39, 0x67, 0xce, 0xe5, 0x9b, 0x6f, 0x0d, 0xbc, 0xeb,
	0x0e, 0x06, 0x6c, 0x12, 0xf0, 0x8f, 0xd5, 0xdf, 0xbd, 0x71, 0xc8, 0x38, 0x23, 0x2b, 0xca, 0xb4,
	0x7e, 0x2a, 0x41, 0xf5, 0x3c, 0xa2, 0xa1, 0x1d, 0x5c, 0x32, 0x52, 0x87, 0xa2, 0xed, 0x99, 0x85,
	0x9d, 0xc2, 0x6e, 0xc9, 0x29, 0xda, 0x1e, 0xd9, 0x06, 0xe3, 0x31, 0xbb, 0xf2, 0x83, 0x9e, 0x7b,
	0x43, 0xcd, 0xe2, 0x4e, 0x61, 0xd7, 0x70, 0x12, 0x80, 0xb4, 0xa0, 0xda, 0xf3, 0x07, 0xd7, 0x81,
	0xd8, 0x2c, 0xe1, 0xa6, 0xb6, 0x49, 0x03, 0x4a, 0x07, 0x57, 0xd4, 0x2c, 0x63, 0x28, 0xb1, 0x24,
	0x9b, 0xb0, 0x7c, 0x42, 0x03, 0x8f, 0x86, 0x66, 0x05, 0x41, 0x65, 0x09, 0xfc, 0xe0, 0x3b, 0x97,
	0xbb, 0xa1, 0xb9, 0x8c, 0x31, 0x94, 0x25, 0xa2, 0x3f, 0x66, 0x03, 0x97, 0xfb, 0x2c, 0x30, 0x57,
	0x64, 0xf4, 0xd8, 0x26, 0x0f, 0x01, 0xfa, 0x21, 0xbb, 0xa4, 0x51, 0x24, 0x76, 0xab, 0xb8, 0x9b,
	0x42, 0x44, 0xcc, 0x33, 0xee, 0xf2, 0x49, 0x64, 0x1a, 0xf2, 0x2e, 0x69, 0x91, 0x26, 0x54, 0x0e,
	0x43, 0x9f, 0x5e, 0x9a, 0x6b, 0xe8, 0x22, 0x0d, 0x81, 0x76, 0x6f, 0x5c, 0x7f, 0x64, 0xd6, 0x24,
	0x8a, 0x86, 0x40, 0xfb, 0x43, 0x16, 0x50, 0xb3, 0x2e, 0x51, 0x34, 0x44, 0x47, 0x70, 0xd1, 0x61,
	0x1e, 0x35, 0xd7, 0x65, 0x47, 0x34, 0x20, 0x72, 0x7e, 0x46, 0x3b, 0x43, 0x97, 0xdb, 0x9e, 0xd9,
	0x90, 0x39, 0xc7, 0x36, 0x21, 0x50, 0x7e, 0xf2, 0xc4, 0xf6, 0xcc, 0x0d, 0xc4, 0x71, 0x2d, 0xce,
	0xf7, 0xdd, 0x28, 0x7a, 0xc1, 0x42, 0xcf, 0x24, 0xf2, 0x7c, 0x6c, 0x93, 0x0f, 0xa0, 0xe2, 0xb0,
	0x11, 0x8d, 0xcc, 0x77, 0x76, 0x4a, 0xbb, 0xab, 0xfb, 0x1b, 0x7b, 0xf1, 0x00, 0x05, 0x2a, 0xa6,
	0xe5, 0xc8, 0x7d, 0xeb, 0x9f, 0x02, 0x90, 0xfe, 0x84, 0x8b, 0x21, 0x1e, 0xba, 0x11, 0x75, 0xe8,
	0xb7, 0x13, 0x1a, 0xf1, 0xff, 0xcd, 0x2c, 0xf5, 0xcc, 0x8c, 0xd4, 0xcc, 0xac, 0x6f, 0x60, 0x23,
	0x66, 0xed, 0x33, 0x9f, 0x0f, 0x9f, 0xb2, 0x6b, 0x1a, 0x90, 0x8f, 0x12, 0x2a, 0x63, 0xe1, 0xe9,
	0xae, 0xc5, 0x1b, 0x4e, 0xc2, 0xf6, 0x26, 0x54, 0xd0, 0x4f, 0x75, 0x43, 0x1a, 0xd6, 0x36, 0x54,
	0xcf, 0xed, 0x23, 0x3b, 0x18, 0x4f, 0xb8, 0xa8, 0xfc, 0xdc, 0x3e, 0x52, 0x4d, 0x14, 0x4b, 0xeb,
	0x2b, 0xa8, 0xc6, 0xfd, 0x17, 0x13, 0x15, 0x6b, 0xdc, 0x36, 0x1c, 0x5c, 0x93, 0x1d, 0x58, 0x3d,
	0xa2, 0xd1, 0x20, 0xf4, 0xc7, 0x58, 0xac, 0x8c, 0x9c, 0x86, 0xac, 0x2f, 0x60, 0xe3, 0xd4, 0x9d,
	0xf0, 0x21, 0xf6, 0x3e, 0x1e, 0x16, 0x81, 0xf2, 0xd3, 0xdb, 0xb1, 0x0e, 0x25, 0xd6, 0x02, 0x43,
	0x96, 0xc9, 0x18, 0xb8, 0xb6, 0xba, 0x50, 0xeb, 0x86, 0x21, 0x0b, 0x1d, 0x1a, 0x8d, 0x59, 0x10,
	0x51, 0xf2, 0x29, 0x18, 0x08, 0xe8, 0x9a, 0xeb, 0xfb, 0x9b, 0xba, 0xe6, 0x03, 0xf9, 0x17, 0x0f,
	0x44, 0x4e, 0x72, 0xd0, 0x7a, 0x55, 0x80, 0xf5, 0x4e, 0x48, 0x5d, 0x4e, 0x65, 0x33, 0x44, 0xad,
	0xf3, 0x52, 0x68, 0x42, 0xe5, 0xc2, 0x1d, 0x4d, 0xe2, 0x1c, 0xa4, 0x91, 0x61, 0x6d, 0x29, 0xc7,
	0xda, 0x34, 0x8f, 0x56, 0x72, 0x3c, 0xda, 0x06, 0xa3, 0xfb, 0x92, 0xd3, 0x00, 0x07, 0x5d, 0x96,
	0x0c, 0xd4, 0x00, 0xf2, 0xb5, 0x8f, 0x7c, 0x32, 0x9c, 0xa2, 0xdd, 0x27, 0x16, 0xac, 0x39, 0xf4,
	0xca, 0x8f, 0x38, 0x0d, 0x8f, 0x43, 0x76, 0xa3, 0x18, 0x95, 0xc1, 0x2c, 0x0e, 0x0d, 0x51, 0xc0,
	0x1b, 0x5b, 0x79, 0xa7, 0x3a, 0x2e, 0xdc, 0x91, 0xef, 0xb9, 0x5c, 0x12, 0xbf, 0xea, 0x68, 0xdb,
	0x7a, 0x0f, 0x0c, 0x71, 0xab, 0xe4, 0x9c, 0x26, 0x51, 0x21, 0x4d, 0xa2, 0x21, 0x3e, 0x49, 0x11,
	0x8d, 0xbd, 0x08, 0xbd, 0x38, 0xb5, 0x19, 0x3a, 0x65, 0x52, 0x28, 0xe6, 0x52, 0x78, 0x04, 0xb5,
	0xe3, 0x90, 0x46, 0xc3, 0x5c, 0x8e, 0x59, 0xd0, 0xda, 0x83, 0x66, 0xbc, 0x76, 0xdc, 0xc0, 0x63,
	0x37, 0x4a, 0xea, 0x36, 0x61, 0x59, 0xda, 0xea, 0x3a, 0x65, 0x59, 0x47, 0xd0, 0x52, 0x62, 0xd1,
	0x61, 0x01, 0x77, 0x07, 0x1c, 0xd5, 0x2e, 0xce, 0x50, 0x4b, 0x61, 0x21, 0x2d, 0x85, 0x2a, 0xef,
	0x62, 0xf2, 0x0c, 0x9e, 0xe7, 0xa3, 0xa0, 0x06, 0xa6, 0xa2, 0x48, 0xe9, 0x2c, 0x2c, 0x94, 0xce,
	0x62, 0x5e, 0x3a, 0xd5, 0x1d, 0xa5, 0xe4, 0x8e, 0xef, 0x0b, 0xb0, 0x9a, 0xba, 0xe1, 0x4e, 0x51,
	0x75, 0x3d, 0xa5, 0x74, 0x3d, 0x69, 0x99, 0x2e, 0x2f, 0x90, 0xe9, 0x4a, 0x22, 0xd3, 0x1f, 0xfe,
	0x50, 0x84, 0x5a, 0xe6, 0x2d, 0x91, 0x2a, 0x94, 0x7b, 0x2c, 0xa0, 0x8d, 0x25, 0x42, 0x60, 0x4d,
	0x24, 0xd9, 0x63, 0xfc, 0x98, 0x4d, 0x02, 0xaf, 0xf1, 0xf3, 0xb4, 0x4d, 0x36, 0x25, 0x2d, 0xe5,
	0x0b, 0x3b, 0x76, 0xfd, 0x11, 0xf5, 0x1a, 0xbf, 0x24, 0xf8, 0xf9, 0xd8, 0x4b, 0xf0, 0x5f, 0xa7,
	0x6d, 0xb2, 0x25, 0xc5, 0x2c, 0x7e, 0x28, 0xdd, 0x97, 0x7e, 0xc4, 0x1b, 0xbf, 0x4d, 0xdb, 0xc4,
	0x04, 0xa2, 0xf9, 0xdd, 0xd3, 0x3b, 0xaf, 0xa6, 0x6d, 0xd2, 0x84, 0xba, 0xd8, 0xc1, 0x4a, 0x25,
	0xfa, 0x7b, 0x82, 0x62, 0x95, 0x12, 0xfd, 0x63, 0xda, 0x26, 0x0f, 0x60, 0x0b, 0xcf, 0x2a, 0x9a,
	0x74, 0x86, 0x74, 0x70, 0xad, 0x6e, 0xff, 0x73, 0xda, 0x26, 0x8f, 0xe0, 0x21, 0x7e, 0x3b, 0xfc,
	0xc0, 0xcb, 0xd3, 0x42, 0x9d, 0xfa, 0x6b, 0xe1, 0x29, 0xbc, 0x5f, 0x9d, 0xfa, 0x7b, 0xda, 0xde,
	0xff, 0xb1, 0x0a, 0x86, 0xea, 0xd4, 0xd9, 0x05, 0xf9, 0x0c, 0x56, 0x4f, 0x28, 0xd7, 0x7a, 0x9b,
	0x12, 0x63, 0x25, 0xb0, 0xad, 0x59, 0x7d, 0xb6, 0x96, 0x48, 0x1b, 0x48, 0xca, 0xed, 0xf0, 0x56,
	0x3e, 0x34, 0x92, 0x39, 0x8a, 0xd8, 0x7c, 0xf7, 0x53, 0xb8, 0x97, 0xd6, 0x36, 0x11, 0x21, 0xf9,
	0xce, 0x99, 0xda, 0x23, 0xa7, 0x7f, 0xad, 0x44, 0x36, 0x33, 0x0a, 0x6b, 0x2d, 0x91, 0x13, 0x58,
	0x53, 0xf9, 0xc8, 0x4c, 0xee, 0x65, 0x6e, 0x4d, 0x8b, 0x4f, 0xab, 0x35, 0x93, 0x90, 0xfe, 0x3a,
	0x59, 0x4b, 0xe4, 0x4b, 0xa8, 0x61, 0xf3, 0xb5, 0x22, 0xce, 0xe6, 0xff, 0x9a, 0x34, 0x3e, 0x07,
	0x40, 0x6f, 0xc9, 0xfb, 0x3b, 0xb9, 0x4a, 0xfa, 0xbf, 0x95, 0x6b, 0x1b, 0xea, 0xe8, 0x9a, 0xb4,
	0xf0, 0xad, 0xdc, 0xbf, 0x86, 0x75, 0x25, 0x14, 0x5a, 0xd7, 0xee, 0xeb, 0xc3, 0xb3, 0x12, 0xf9,
	0x9f, 0x22, 0x89, 0x5f, 0x39, 0x48, 0xa8, 0x4c, 0xa4, 0xdc, 0xef, 0x9f, 0xd7, 0x44, 0x3a, 0x83,
	0xad, 0xf4, 0x3c, 0x0f, 0x6f, 0xf1, 0x7b, 0x2c, 0x7f, 0xc0, 0x69, 0xa7, 0x99, 0x6f, 0xf4, 0x1b,
	0x66, 0x7b, 0x0a, 0xf7, 0x55, 0xd0, 0xb9, 0x72, 0x3c, 0x87, 0xfb, 0x0f, 0x92, 0xec, 0xe7, 0x79,
	0x9c, 0x41, 0x73, 0xde, 0x93, 0x24, 0xef, 0xe7, 0x8b, 0x9e, 0xa3, 0xe3, 0x8b, 0x8a, 0x9f, 0x13,
	0x54, 0xb2, 0x69, 0x51, 0xd0, 0xb4, 0xac, 0x2f, 0x0c, 0x9a, 0xbc, 0x57, 0xe5, 0xb5, 0xe8, 0xb5,
	0x37, 0x33, 0x1d, 0x54, 0x87, 0x9f, 0x2f, 0xe3, 0x7f, 0x24, 0x9f, 0xfc, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0x1c, 0x97, 0x93, 0xfc, 0xaa, 0x0c, 0x00, 0x00,
}
