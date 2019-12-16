// Code generated by protoc-gen-go. DO NOT EDIT.
// source: push/push.proto

package push

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

type SendType int32

const (
	SendType_NoneSendType SendType = 0
	SendType_Mail         SendType = 1
	SendType_Phone        SendType = 2
)

var SendType_name = map[int32]string{
	0: "NoneSendType",
	1: "Mail",
	2: "Phone",
}

var SendType_value = map[string]int32{
	"NoneSendType": 0,
	"Mail":         1,
	"Phone":        2,
}

func (x SendType) String() string {
	return proto.EnumName(SendType_name, int32(x))
}

func (SendType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ae0042da44e9a7a7, []int{0}
}

type PushErrors int32

const (
	PushErrors_None            PushErrors = 0
	PushErrors_PushCommonError PushErrors = 2000001
)

var PushErrors_name = map[int32]string{
	0:       "None",
	2000001: "PushCommonError",
}

var PushErrors_value = map[string]int32{
	"None":            0,
	"PushCommonError": 2000001,
}

func (x PushErrors) String() string {
	return proto.EnumName(PushErrors_name, int32(x))
}

func (PushErrors) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ae0042da44e9a7a7, []int{1}
}

type PushCaptchaRequest struct {
	Code                 string   `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Sender               string   `protobuf:"bytes,2,opt,name=Sender,proto3" json:"Sender,omitempty"`
	Type                 SendType `protobuf:"varint,3,opt,name=Type,proto3,enum=push.SendType" json:"Type,omitempty"`
	UID                  int64    `protobuf:"varint,4,opt,name=UID,proto3" json:"UID,omitempty"`
	Account              string   `protobuf:"bytes,5,opt,name=Account,proto3" json:"Account,omitempty"`
	NationCode           string   `protobuf:"bytes,6,opt,name=NationCode,proto3" json:"NationCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushCaptchaRequest) Reset()         { *m = PushCaptchaRequest{} }
func (m *PushCaptchaRequest) String() string { return proto.CompactTextString(m) }
func (*PushCaptchaRequest) ProtoMessage()    {}
func (*PushCaptchaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae0042da44e9a7a7, []int{0}
}

func (m *PushCaptchaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushCaptchaRequest.Unmarshal(m, b)
}
func (m *PushCaptchaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushCaptchaRequest.Marshal(b, m, deterministic)
}
func (m *PushCaptchaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushCaptchaRequest.Merge(m, src)
}
func (m *PushCaptchaRequest) XXX_Size() int {
	return xxx_messageInfo_PushCaptchaRequest.Size(m)
}
func (m *PushCaptchaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushCaptchaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushCaptchaRequest proto.InternalMessageInfo

func (m *PushCaptchaRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *PushCaptchaRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *PushCaptchaRequest) GetType() SendType {
	if m != nil {
		return m.Type
	}
	return SendType_NoneSendType
}

func (m *PushCaptchaRequest) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *PushCaptchaRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *PushCaptchaRequest) GetNationCode() string {
	if m != nil {
		return m.NationCode
	}
	return ""
}

type ErrorResponse struct {
	ErrorInfo            PushErrors `protobuf:"varint,1,opt,name=ErrorInfo,proto3,enum=push.PushErrors" json:"ErrorInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ErrorResponse) Reset()         { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()    {}
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae0042da44e9a7a7, []int{1}
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

func (m *ErrorResponse) GetErrorInfo() PushErrors {
	if m != nil {
		return m.ErrorInfo
	}
	return PushErrors_None
}

type PushEmailNotificationRequest struct {
	To                   string   `protobuf:"bytes,1,opt,name=To,proto3" json:"To,omitempty"`
	Html                 string   `protobuf:"bytes,2,opt,name=Html,proto3" json:"Html,omitempty"`
	Subject              string   `protobuf:"bytes,3,opt,name=Subject,proto3" json:"Subject,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushEmailNotificationRequest) Reset()         { *m = PushEmailNotificationRequest{} }
func (m *PushEmailNotificationRequest) String() string { return proto.CompactTextString(m) }
func (*PushEmailNotificationRequest) ProtoMessage()    {}
func (*PushEmailNotificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae0042da44e9a7a7, []int{2}
}

func (m *PushEmailNotificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushEmailNotificationRequest.Unmarshal(m, b)
}
func (m *PushEmailNotificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushEmailNotificationRequest.Marshal(b, m, deterministic)
}
func (m *PushEmailNotificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushEmailNotificationRequest.Merge(m, src)
}
func (m *PushEmailNotificationRequest) XXX_Size() int {
	return xxx_messageInfo_PushEmailNotificationRequest.Size(m)
}
func (m *PushEmailNotificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushEmailNotificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushEmailNotificationRequest proto.InternalMessageInfo

func (m *PushEmailNotificationRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *PushEmailNotificationRequest) GetHtml() string {
	if m != nil {
		return m.Html
	}
	return ""
}

func (m *PushEmailNotificationRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

type GetCaptchaVerifyRequest struct {
	Code                 string   `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=Account,proto3" json:"Account,omitempty"`
	UID                  int64    `protobuf:"varint,3,opt,name=UID,proto3" json:"UID,omitempty"`
	Type                 SendType `protobuf:"varint,4,opt,name=Type,proto3,enum=push.SendType" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCaptchaVerifyRequest) Reset()         { *m = GetCaptchaVerifyRequest{} }
func (m *GetCaptchaVerifyRequest) String() string { return proto.CompactTextString(m) }
func (*GetCaptchaVerifyRequest) ProtoMessage()    {}
func (*GetCaptchaVerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae0042da44e9a7a7, []int{3}
}

func (m *GetCaptchaVerifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCaptchaVerifyRequest.Unmarshal(m, b)
}
func (m *GetCaptchaVerifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCaptchaVerifyRequest.Marshal(b, m, deterministic)
}
func (m *GetCaptchaVerifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCaptchaVerifyRequest.Merge(m, src)
}
func (m *GetCaptchaVerifyRequest) XXX_Size() int {
	return xxx_messageInfo_GetCaptchaVerifyRequest.Size(m)
}
func (m *GetCaptchaVerifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCaptchaVerifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCaptchaVerifyRequest proto.InternalMessageInfo

func (m *GetCaptchaVerifyRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *GetCaptchaVerifyRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *GetCaptchaVerifyRequest) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *GetCaptchaVerifyRequest) GetType() SendType {
	if m != nil {
		return m.Type
	}
	return SendType_NoneSendType
}

func init() {
	proto.RegisterEnum("push.SendType", SendType_name, SendType_value)
	proto.RegisterEnum("push.PushErrors", PushErrors_name, PushErrors_value)
	proto.RegisterType((*PushCaptchaRequest)(nil), "push.PushCaptchaRequest")
	proto.RegisterType((*ErrorResponse)(nil), "push.ErrorResponse")
	proto.RegisterType((*PushEmailNotificationRequest)(nil), "push.PushEmailNotificationRequest")
	proto.RegisterType((*GetCaptchaVerifyRequest)(nil), "push.GetCaptchaVerifyRequest")
}

func init() { proto.RegisterFile("push/push.proto", fileDescriptor_ae0042da44e9a7a7) }

var fileDescriptor_ae0042da44e9a7a7 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x4d, 0xcf, 0xd2, 0x40,
	0x10, 0xc7, 0x9f, 0xbe, 0x3c, 0xf5, 0xe9, 0xa8, 0x65, 0x33, 0x06, 0x6d, 0x8c, 0x1a, 0xd2, 0x13,
	0x21, 0x11, 0x23, 0x5e, 0x4d, 0x8c, 0xc1, 0x37, 0x0e, 0x12, 0x52, 0x90, 0x93, 0x97, 0x52, 0x96,
	0xb4, 0x86, 0xee, 0xd6, 0x76, 0x7b, 0xc0, 0xc4, 0x83, 0x47, 0xbf, 0x8d, 0x1f, 0xc9, 0x8f, 0x62,
	0x76, 0x68, 0xd3, 0x62, 0xe0, 0xb9, 0x90, 0x99, 0xf9, 0x0f, 0xb3, 0x33, 0xbf, 0x99, 0x42, 0x2f,
	0xaf, 0xca, 0xe4, 0x85, 0xfe, 0x19, 0xe7, 0x85, 0x54, 0x12, 0x6d, 0x6d, 0x07, 0x7f, 0x0c, 0xc0,
	0x45, 0x55, 0x26, 0xd3, 0x28, 0x57, 0x71, 0x12, 0x85, 0xfc, 0x7b, 0xc5, 0x4b, 0x85, 0x08, 0xf6,
	0x54, 0x6e, 0xb9, 0x6f, 0x0c, 0x8c, 0xa1, 0x1b, 0x92, 0x8d, 0x0f, 0xc1, 0x59, 0x72, 0xb1, 0xe5,
	0x85, 0x6f, 0x52, 0xb4, 0xf6, 0x30, 0x00, 0x7b, 0x75, 0xc8, 0xb9, 0x6f, 0x0d, 0x8c, 0xa1, 0x37,
	0xf1, 0xc6, 0xf4, 0x86, 0xd6, 0x74, 0x34, 0x24, 0x0d, 0x19, 0x58, 0x5f, 0x66, 0xef, 0x7c, 0x7b,
	0x60, 0x0c, 0xad, 0x50, 0x9b, 0xe8, 0xc3, 0x9d, 0xb7, 0x71, 0x2c, 0x2b, 0xa1, 0xfc, 0x6b, 0x2a,
	0xd7, 0xb8, 0xf8, 0x0c, 0x60, 0x1e, 0xa9, 0x54, 0x0a, 0xea, 0xc0, 0x21, 0xb1, 0x13, 0x09, 0xde,
	0xc0, 0xfd, 0xf7, 0x45, 0x21, 0x8b, 0x90, 0x97, 0xb9, 0x14, 0x25, 0xc7, 0x31, 0xb8, 0x14, 0x98,
	0x89, 0x9d, 0xa4, 0x8e, 0xbd, 0x09, 0x3b, 0x76, 0xa1, 0x27, 0x23, 0xa9, 0x0c, 0xdb, 0x94, 0xe0,
	0x2b, 0x3c, 0x21, 0x21, 0x8b, 0xd2, 0xfd, 0x5c, 0xaa, 0x74, 0x97, 0xc6, 0x54, 0xbd, 0x19, 0xde,
	0x03, 0x73, 0x25, 0xeb, 0xd1, 0xcd, 0x95, 0xd4, 0x30, 0x3e, 0xa9, 0x6c, 0x5f, 0x8f, 0x4d, 0xb6,
	0x6e, 0x7f, 0x59, 0x6d, 0xbe, 0xf1, 0x58, 0xd1, 0xdc, 0x6e, 0xd8, 0xb8, 0xc1, 0x4f, 0x78, 0xf4,
	0x91, 0xab, 0x9a, 0xe7, 0x9a, 0x17, 0xe9, 0xee, 0x70, 0x1b, 0xd5, 0x0e, 0x07, 0xf3, 0x94, 0x43,
	0xcd, 0xcc, 0x6a, 0x99, 0x35, 0xa4, 0xed, 0xcb, 0xa4, 0x47, 0x2f, 0xe1, 0xa6, 0x89, 0x20, 0x83,
	0x7b, 0x73, 0x29, 0x78, 0xe3, 0xb3, 0x2b, 0xbc, 0x01, 0xfb, 0x73, 0x94, 0xee, 0x99, 0x81, 0x2e,
	0x5c, 0x2f, 0x12, 0x29, 0x38, 0x33, 0x47, 0xcf, 0x01, 0x5a, 0x50, 0x3a, 0x45, 0xff, 0x89, 0x5d,
	0x61, 0x1f, 0x7a, 0x74, 0x1a, 0x32, 0xcb, 0xa4, 0x20, 0x95, 0xfd, 0xfa, 0xfd, 0x63, 0xf2, 0xd7,
	0x00, 0x47, 0xbf, 0xbc, 0x5c, 0xe3, 0x6b, 0xb8, 0xdb, 0x39, 0x1e, 0xf4, 0x5b, 0xea, 0xa7, 0xf7,
	0xf4, 0xf8, 0xc1, 0x51, 0x39, 0xdd, 0xdb, 0x07, 0x60, 0xff, 0x93, 0xc2, 0xa7, 0xc7, 0xc4, 0x0b,
	0x04, 0xcf, 0xd7, 0x59, 0x40, 0xff, 0xec, 0x3e, 0x31, 0xe8, 0x5c, 0xc1, 0x85, 0x65, 0x9f, 0xad,
	0xb8, 0x71, 0xe8, 0x13, 0x79, 0xf5, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x5d, 0x76, 0xb2, 0x35,
	0x03, 0x00, 0x00,
}
