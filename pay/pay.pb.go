// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pay/pay.proto

package pay

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

type PayErrors int32

const (
	PayErrors_None                    PayErrors = 0
	PayErrors_PayWayNotFound          PayErrors = 2000001
	PayErrors_PayWayOperationFailed   PayErrors = 2000002
	PayErrors_CreateAlipayOrderFailed PayErrors = 2000003
)

var PayErrors_name = map[int32]string{
	0:       "None",
	2000001: "PayWayNotFound",
	2000002: "PayWayOperationFailed",
	2000003: "CreateAlipayOrderFailed",
}

var PayErrors_value = map[string]int32{
	"None":                    0,
	"PayWayNotFound":          2000001,
	"PayWayOperationFailed":   2000002,
	"CreateAlipayOrderFailed": 2000003,
}

func (x PayErrors) String() string {
	return proto.EnumName(PayErrors_name, int32(x))
}

func (PayErrors) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_587e2c2585dcd479, []int{0}
}

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
	return fileDescriptor_587e2c2585dcd479, []int{0}
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

type PayWay struct {
	PayWayID             int64    `protobuf:"varint,1,opt,name=PayWayID,proto3" json:"PayWayID,omitempty"`
	PayType              string   `protobuf:"bytes,2,opt,name=PayType,proto3" json:"PayType,omitempty"`
	PayValue             string   `protobuf:"bytes,3,opt,name=PayValue,proto3" json:"PayValue,omitempty"`
	PayDesc              string   `protobuf:"bytes,4,opt,name=PayDesc,proto3" json:"PayDesc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayWay) Reset()         { *m = PayWay{} }
func (m *PayWay) String() string { return proto.CompactTextString(m) }
func (*PayWay) ProtoMessage()    {}
func (*PayWay) Descriptor() ([]byte, []int) {
	return fileDescriptor_587e2c2585dcd479, []int{1}
}

func (m *PayWay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayWay.Unmarshal(m, b)
}
func (m *PayWay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayWay.Marshal(b, m, deterministic)
}
func (m *PayWay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayWay.Merge(m, src)
}
func (m *PayWay) XXX_Size() int {
	return xxx_messageInfo_PayWay.Size(m)
}
func (m *PayWay) XXX_DiscardUnknown() {
	xxx_messageInfo_PayWay.DiscardUnknown(m)
}

var xxx_messageInfo_PayWay proto.InternalMessageInfo

func (m *PayWay) GetPayWayID() int64 {
	if m != nil {
		return m.PayWayID
	}
	return 0
}

func (m *PayWay) GetPayType() string {
	if m != nil {
		return m.PayType
	}
	return ""
}

func (m *PayWay) GetPayValue() string {
	if m != nil {
		return m.PayValue
	}
	return ""
}

func (m *PayWay) GetPayDesc() string {
	if m != nil {
		return m.PayDesc
	}
	return ""
}

type CreateUserPayWayRequest struct {
	PayType              string   `protobuf:"bytes,1,opt,name=PayType,proto3" json:"PayType,omitempty"`
	PayValue             string   `protobuf:"bytes,2,opt,name=PayValue,proto3" json:"PayValue,omitempty"`
	PayDesc              string   `protobuf:"bytes,3,opt,name=PayDesc,proto3" json:"PayDesc,omitempty"`
	UID                  int64    `protobuf:"varint,4,opt,name=UID,proto3" json:"UID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserPayWayRequest) Reset()         { *m = CreateUserPayWayRequest{} }
func (m *CreateUserPayWayRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserPayWayRequest) ProtoMessage()    {}
func (*CreateUserPayWayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_587e2c2585dcd479, []int{2}
}

func (m *CreateUserPayWayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserPayWayRequest.Unmarshal(m, b)
}
func (m *CreateUserPayWayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserPayWayRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserPayWayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserPayWayRequest.Merge(m, src)
}
func (m *CreateUserPayWayRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserPayWayRequest.Size(m)
}
func (m *CreateUserPayWayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserPayWayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserPayWayRequest proto.InternalMessageInfo

func (m *CreateUserPayWayRequest) GetPayType() string {
	if m != nil {
		return m.PayType
	}
	return ""
}

func (m *CreateUserPayWayRequest) GetPayValue() string {
	if m != nil {
		return m.PayValue
	}
	return ""
}

func (m *CreateUserPayWayRequest) GetPayDesc() string {
	if m != nil {
		return m.PayDesc
	}
	return ""
}

func (m *CreateUserPayWayRequest) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

type DeleteUserPayWayRequest struct {
	UID                  int64    `protobuf:"varint,1,opt,name=UID,proto3" json:"UID,omitempty"`
	PayWayID             int64    `protobuf:"varint,2,opt,name=PayWayID,proto3" json:"PayWayID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserPayWayRequest) Reset()         { *m = DeleteUserPayWayRequest{} }
func (m *DeleteUserPayWayRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserPayWayRequest) ProtoMessage()    {}
func (*DeleteUserPayWayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_587e2c2585dcd479, []int{3}
}

func (m *DeleteUserPayWayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserPayWayRequest.Unmarshal(m, b)
}
func (m *DeleteUserPayWayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserPayWayRequest.Marshal(b, m, deterministic)
}
func (m *DeleteUserPayWayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserPayWayRequest.Merge(m, src)
}
func (m *DeleteUserPayWayRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserPayWayRequest.Size(m)
}
func (m *DeleteUserPayWayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserPayWayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserPayWayRequest proto.InternalMessageInfo

func (m *DeleteUserPayWayRequest) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *DeleteUserPayWayRequest) GetPayWayID() int64 {
	if m != nil {
		return m.PayWayID
	}
	return 0
}

type PutUserPayWayRequest struct {
	UID                  int64    `protobuf:"varint,1,opt,name=UID,proto3" json:"UID,omitempty"`
	PayWayID             int64    `protobuf:"varint,2,opt,name=PayWayID,proto3" json:"PayWayID,omitempty"`
	PayType              string   `protobuf:"bytes,3,opt,name=PayType,proto3" json:"PayType,omitempty"`
	PayValue             string   `protobuf:"bytes,4,opt,name=PayValue,proto3" json:"PayValue,omitempty"`
	PayDesc              string   `protobuf:"bytes,5,opt,name=PayDesc,proto3" json:"PayDesc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutUserPayWayRequest) Reset()         { *m = PutUserPayWayRequest{} }
func (m *PutUserPayWayRequest) String() string { return proto.CompactTextString(m) }
func (*PutUserPayWayRequest) ProtoMessage()    {}
func (*PutUserPayWayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_587e2c2585dcd479, []int{4}
}

func (m *PutUserPayWayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutUserPayWayRequest.Unmarshal(m, b)
}
func (m *PutUserPayWayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutUserPayWayRequest.Marshal(b, m, deterministic)
}
func (m *PutUserPayWayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutUserPayWayRequest.Merge(m, src)
}
func (m *PutUserPayWayRequest) XXX_Size() int {
	return xxx_messageInfo_PutUserPayWayRequest.Size(m)
}
func (m *PutUserPayWayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutUserPayWayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutUserPayWayRequest proto.InternalMessageInfo

func (m *PutUserPayWayRequest) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *PutUserPayWayRequest) GetPayWayID() int64 {
	if m != nil {
		return m.PayWayID
	}
	return 0
}

func (m *PutUserPayWayRequest) GetPayType() string {
	if m != nil {
		return m.PayType
	}
	return ""
}

func (m *PutUserPayWayRequest) GetPayValue() string {
	if m != nil {
		return m.PayValue
	}
	return ""
}

func (m *PutUserPayWayRequest) GetPayDesc() string {
	if m != nil {
		return m.PayDesc
	}
	return ""
}

type UserPayWays struct {
	PayWays              []*PayWay `protobuf:"bytes,1,rep,name=PayWays,proto3" json:"PayWays,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UserPayWays) Reset()         { *m = UserPayWays{} }
func (m *UserPayWays) String() string { return proto.CompactTextString(m) }
func (*UserPayWays) ProtoMessage()    {}
func (*UserPayWays) Descriptor() ([]byte, []int) {
	return fileDescriptor_587e2c2585dcd479, []int{5}
}

func (m *UserPayWays) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserPayWays.Unmarshal(m, b)
}
func (m *UserPayWays) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserPayWays.Marshal(b, m, deterministic)
}
func (m *UserPayWays) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPayWays.Merge(m, src)
}
func (m *UserPayWays) XXX_Size() int {
	return xxx_messageInfo_UserPayWays.Size(m)
}
func (m *UserPayWays) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPayWays.DiscardUnknown(m)
}

var xxx_messageInfo_UserPayWays proto.InternalMessageInfo

func (m *UserPayWays) GetPayWays() []*PayWay {
	if m != nil {
		return m.PayWays
	}
	return nil
}

type ErrorResponse struct {
	ErrorInfo            PayErrors `protobuf:"varint,1,opt,name=ErrorInfo,proto3,enum=pay.PayErrors" json:"ErrorInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ErrorResponse) Reset()         { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()    {}
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_587e2c2585dcd479, []int{6}
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

func (m *ErrorResponse) GetErrorInfo() PayErrors {
	if m != nil {
		return m.ErrorInfo
	}
	return PayErrors_None
}

type CreateAliPayOrderRequest struct {
	CreateUserID         int64    `protobuf:"varint,1,opt,name=CreateUserID,proto3" json:"CreateUserID,omitempty"`
	OrderType            string   `protobuf:"bytes,2,opt,name=OrderType,proto3" json:"OrderType,omitempty"`
	CreateForUserID      int64    `protobuf:"varint,3,opt,name=CreateForUserID,proto3" json:"CreateForUserID,omitempty"`
	TotalAmount          string   `protobuf:"bytes,4,opt,name=TotalAmount,proto3" json:"TotalAmount,omitempty"`
	Subject              string   `protobuf:"bytes,5,opt,name=Subject,proto3" json:"Subject,omitempty"`
	OrderExtends         string   `protobuf:"bytes,6,opt,name=OrderExtends,proto3" json:"OrderExtends,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAliPayOrderRequest) Reset()         { *m = CreateAliPayOrderRequest{} }
func (m *CreateAliPayOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAliPayOrderRequest) ProtoMessage()    {}
func (*CreateAliPayOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_587e2c2585dcd479, []int{7}
}

func (m *CreateAliPayOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAliPayOrderRequest.Unmarshal(m, b)
}
func (m *CreateAliPayOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAliPayOrderRequest.Marshal(b, m, deterministic)
}
func (m *CreateAliPayOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAliPayOrderRequest.Merge(m, src)
}
func (m *CreateAliPayOrderRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAliPayOrderRequest.Size(m)
}
func (m *CreateAliPayOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAliPayOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAliPayOrderRequest proto.InternalMessageInfo

func (m *CreateAliPayOrderRequest) GetCreateUserID() int64 {
	if m != nil {
		return m.CreateUserID
	}
	return 0
}

func (m *CreateAliPayOrderRequest) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

func (m *CreateAliPayOrderRequest) GetCreateForUserID() int64 {
	if m != nil {
		return m.CreateForUserID
	}
	return 0
}

func (m *CreateAliPayOrderRequest) GetTotalAmount() string {
	if m != nil {
		return m.TotalAmount
	}
	return ""
}

func (m *CreateAliPayOrderRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *CreateAliPayOrderRequest) GetOrderExtends() string {
	if m != nil {
		return m.OrderExtends
	}
	return ""
}

type GetPayOrderResultRequest struct {
	OutTradeNo           string   `protobuf:"bytes,1,opt,name=OutTradeNo,proto3" json:"OutTradeNo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPayOrderResultRequest) Reset()         { *m = GetPayOrderResultRequest{} }
func (m *GetPayOrderResultRequest) String() string { return proto.CompactTextString(m) }
func (*GetPayOrderResultRequest) ProtoMessage()    {}
func (*GetPayOrderResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_587e2c2585dcd479, []int{8}
}

func (m *GetPayOrderResultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPayOrderResultRequest.Unmarshal(m, b)
}
func (m *GetPayOrderResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPayOrderResultRequest.Marshal(b, m, deterministic)
}
func (m *GetPayOrderResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPayOrderResultRequest.Merge(m, src)
}
func (m *GetPayOrderResultRequest) XXX_Size() int {
	return xxx_messageInfo_GetPayOrderResultRequest.Size(m)
}
func (m *GetPayOrderResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPayOrderResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPayOrderResultRequest proto.InternalMessageInfo

func (m *GetPayOrderResultRequest) GetOutTradeNo() string {
	if m != nil {
		return m.OutTradeNo
	}
	return ""
}

type OrderResultResponse struct {
	Code                 string   `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	OutTradeNo           string   `protobuf:"bytes,3,opt,name=OutTradeNo,proto3" json:"OutTradeNo,omitempty"`
	QRCode               string   `protobuf:"bytes,4,opt,name=QRCode,proto3" json:"QRCode,omitempty"`
	Sign                 string   `protobuf:"bytes,5,opt,name=Sign,proto3" json:"Sign,omitempty"`
	TradeStatus          string   `protobuf:"bytes,6,opt,name=TradeStatus,proto3" json:"TradeStatus,omitempty"`
	TradeNo              string   `protobuf:"bytes,7,opt,name=TradeNo,proto3" json:"TradeNo,omitempty"`
	TotalAmount          string   `protobuf:"bytes,8,opt,name=TotalAmount,proto3" json:"TotalAmount,omitempty"`
	BuyerLogonID         string   `protobuf:"bytes,9,opt,name=BuyerLogonID,proto3" json:"BuyerLogonID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderResultResponse) Reset()         { *m = OrderResultResponse{} }
func (m *OrderResultResponse) String() string { return proto.CompactTextString(m) }
func (*OrderResultResponse) ProtoMessage()    {}
func (*OrderResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_587e2c2585dcd479, []int{9}
}

func (m *OrderResultResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderResultResponse.Unmarshal(m, b)
}
func (m *OrderResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderResultResponse.Marshal(b, m, deterministic)
}
func (m *OrderResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderResultResponse.Merge(m, src)
}
func (m *OrderResultResponse) XXX_Size() int {
	return xxx_messageInfo_OrderResultResponse.Size(m)
}
func (m *OrderResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderResultResponse proto.InternalMessageInfo

func (m *OrderResultResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *OrderResultResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *OrderResultResponse) GetOutTradeNo() string {
	if m != nil {
		return m.OutTradeNo
	}
	return ""
}

func (m *OrderResultResponse) GetQRCode() string {
	if m != nil {
		return m.QRCode
	}
	return ""
}

func (m *OrderResultResponse) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func (m *OrderResultResponse) GetTradeStatus() string {
	if m != nil {
		return m.TradeStatus
	}
	return ""
}

func (m *OrderResultResponse) GetTradeNo() string {
	if m != nil {
		return m.TradeNo
	}
	return ""
}

func (m *OrderResultResponse) GetTotalAmount() string {
	if m != nil {
		return m.TotalAmount
	}
	return ""
}

func (m *OrderResultResponse) GetBuyerLogonID() string {
	if m != nil {
		return m.BuyerLogonID
	}
	return ""
}

func init() {
	proto.RegisterEnum("pay.PayErrors", PayErrors_name, PayErrors_value)
	proto.RegisterType((*UIDInput)(nil), "pay.UIDInput")
	proto.RegisterType((*PayWay)(nil), "pay.PayWay")
	proto.RegisterType((*CreateUserPayWayRequest)(nil), "pay.CreateUserPayWayRequest")
	proto.RegisterType((*DeleteUserPayWayRequest)(nil), "pay.DeleteUserPayWayRequest")
	proto.RegisterType((*PutUserPayWayRequest)(nil), "pay.PutUserPayWayRequest")
	proto.RegisterType((*UserPayWays)(nil), "pay.UserPayWays")
	proto.RegisterType((*ErrorResponse)(nil), "pay.ErrorResponse")
	proto.RegisterType((*CreateAliPayOrderRequest)(nil), "pay.CreateAliPayOrderRequest")
	proto.RegisterType((*GetPayOrderResultRequest)(nil), "pay.GetPayOrderResultRequest")
	proto.RegisterType((*OrderResultResponse)(nil), "pay.OrderResultResponse")
}

func init() { proto.RegisterFile("pay/pay.proto", fileDescriptor_587e2c2585dcd479) }

var fileDescriptor_587e2c2585dcd479 = []byte{
	// 675 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0xe3, 0x34, 0x4d, 0x26, 0x4d, 0x09, 0x4b, 0xa1, 0x26, 0xb4, 0x28, 0xb2, 0x84, 0x14,
	0x21, 0xd4, 0x4a, 0x85, 0x13, 0x12, 0x12, 0xa5, 0x69, 0xaa, 0x48, 0xa5, 0x09, 0x49, 0x5b, 0xce,
	0xdb, 0x66, 0x89, 0x82, 0x8c, 0xd7, 0xac, 0xd7, 0x12, 0x2e, 0x27, 0xe0, 0x94, 0x3b, 0x2f, 0xc1,
	0x8b, 0xf1, 0x1c, 0x68, 0xff, 0xfc, 0x97, 0xfa, 0xc4, 0x6d, 0x67, 0x66, 0xe7, 0x1b, 0x7f, 0xdf,
	0xce, 0x8c, 0xa1, 0x15, 0xe0, 0xf8, 0x20, 0xc0, 0xf1, 0x7e, 0xc0, 0x28, 0xa7, 0xc8, 0x0e, 0x70,
	0xec, 0xee, 0x42, 0xfd, 0x72, 0xd8, 0x1f, 0xfa, 0x41, 0xc4, 0x51, 0x1b, 0xec, 0xcb, 0x61, 0xdf,
	0xb1, 0xba, 0x56, 0xcf, 0x9e, 0x88, 0xa3, 0xcb, 0xa1, 0x36, 0xc6, 0xf1, 0x47, 0x1c, 0xa3, 0x0e,
	0xd4, 0xd5, 0x29, 0xb9, 0x90, 0xd8, 0xc8, 0x81, 0x8d, 0x31, 0x8e, 0x2f, 0xe2, 0x80, 0x38, 0x95,
	0xae, 0xd5, 0x6b, 0x4c, 0x8c, 0xa9, 0xb3, 0xae, 0xb0, 0x17, 0x11, 0xc7, 0x96, 0xa1, 0xc4, 0xd6,
	0x59, 0x7d, 0x12, 0xde, 0x38, 0xd5, 0x24, 0x4b, 0x98, 0xee, 0x77, 0xd8, 0x39, 0x66, 0x04, 0x73,
	0x72, 0x19, 0x12, 0xa6, 0xaa, 0x4c, 0xc8, 0xd7, 0x88, 0x84, 0x3c, 0x5b, 0xca, 0x2a, 0x2f, 0x55,
	0x29, 0x2f, 0x65, 0xe7, 0x4a, 0x19, 0xca, 0xd5, 0x94, 0xf2, 0x29, 0xec, 0xf4, 0x89, 0x47, 0xee,
	0x2a, 0xbe, 0xa2, 0x4f, 0x4e, 0x95, 0x4a, 0x5e, 0x15, 0xf7, 0xb7, 0x05, 0xdb, 0xe3, 0x88, 0xff,
	0x27, 0x4c, 0x96, 0xb1, 0x5d, 0xce, 0xb8, 0x5a, 0xce, 0x78, 0x3d, 0x2f, 0xee, 0x2b, 0x68, 0xa6,
	0x9f, 0x14, 0xa2, 0x67, 0xf2, 0xa2, 0x38, 0x3a, 0x56, 0xd7, 0xee, 0x35, 0x0f, 0x9b, 0xfb, 0xa2,
	0x43, 0xf4, 0x17, 0x9b, 0x98, 0xfb, 0x06, 0x5a, 0x27, 0x8c, 0x51, 0x36, 0x21, 0x61, 0x40, 0xfd,
	0x90, 0xa0, 0x17, 0xd0, 0x90, 0x8e, 0xa1, 0xff, 0x89, 0x4a, 0x2a, 0x5b, 0x87, 0x5b, 0x26, 0x53,
	0x06, 0xc2, 0x49, 0x7a, 0xc1, 0xfd, 0x6b, 0x81, 0xa3, 0x9e, 0xf4, 0xc8, 0x5b, 0x8c, 0x71, 0x3c,
	0x62, 0x33, 0xc2, 0x8c, 0x1e, 0x2e, 0x6c, 0xa6, 0xcf, 0x9d, 0x08, 0x93, 0xf3, 0xa1, 0x5d, 0x68,
	0xc8, 0x9c, 0x4c, 0x93, 0xa5, 0x0e, 0xd4, 0x83, 0x7b, 0xea, 0xf6, 0x80, 0x32, 0x0d, 0x62, 0x4b,
	0x90, 0xa2, 0x1b, 0x75, 0xa1, 0x79, 0x41, 0x39, 0xf6, 0x8e, 0xbe, 0xd0, 0xc8, 0xe7, 0x5a, 0xb6,
	0xac, 0x4b, 0x28, 0x37, 0x8d, 0xae, 0x3f, 0x93, 0x1b, 0x6e, 0x94, 0xd3, 0xa6, 0xf8, 0x4e, 0x59,
	0xf2, 0xe4, 0x1b, 0x27, 0xfe, 0x2c, 0x74, 0x6a, 0x32, 0x9c, 0xf3, 0xb9, 0xaf, 0xc1, 0x39, 0x25,
	0x3c, 0x65, 0x18, 0x46, 0x1e, 0x37, 0x3c, 0x9f, 0x02, 0x8c, 0x22, 0x7e, 0xc1, 0xf0, 0x8c, 0x9c,
	0x53, 0xdd, 0xbe, 0x19, 0x8f, 0xbb, 0xac, 0xc0, 0x83, 0x5c, 0x9a, 0x96, 0x1a, 0x41, 0xf5, 0x98,
	0xce, 0x4c, 0xc3, 0xcb, 0xb3, 0xe8, 0xa1, 0xf7, 0xe1, 0x5c, 0x2b, 0x21, 0x8e, 0x05, 0x74, 0xbb,
	0x88, 0x8e, 0x1e, 0x41, 0xed, 0xc3, 0x44, 0xe2, 0x28, 0xd2, 0xda, 0x12, 0xe8, 0xd3, 0xc5, 0xdc,
	0xd7, 0x64, 0xe5, 0x59, 0xaa, 0x24, 0xd2, 0xa6, 0x1c, 0xf3, 0xc8, 0x10, 0xcd, 0xba, 0x84, 0x4a,
	0xa6, 0xd4, 0x86, 0x52, 0xc9, 0xd4, 0x29, 0x28, 0x5c, 0x5f, 0x55, 0xd8, 0x85, 0xcd, 0x77, 0x51,
	0x4c, 0xd8, 0x19, 0x9d, 0x53, 0x7f, 0xd8, 0x77, 0x1a, 0x4a, 0xc7, 0xac, 0xef, 0xf9, 0x1c, 0x1a,
	0x49, 0x23, 0xa1, 0x3a, 0x54, 0xcf, 0xa9, 0x4f, 0xda, 0x6b, 0x68, 0x1b, 0xb6, 0x54, 0x47, 0x9e,
	0x53, 0x3e, 0xa0, 0x91, 0x3f, 0x6b, 0xff, 0x58, 0xde, 0xa2, 0x27, 0xf0, 0x50, 0x79, 0x47, 0x01,
	0x61, 0x98, 0x2f, 0xa8, 0x3f, 0xc0, 0x0b, 0x8f, 0xcc, 0xda, 0x3f, 0x97, 0xb7, 0x68, 0xcf, 0x2c,
	0x93, 0x23, 0x6f, 0x11, 0xe8, 0x77, 0xd1, 0xe1, 0x5f, 0xcb, 0xdb, 0xc3, 0x3f, 0x36, 0xac, 0x07,
	0x38, 0x9e, 0x5e, 0xa1, 0x03, 0x00, 0xf5, 0x74, 0xa2, 0x63, 0x51, 0x4b, 0x36, 0xb3, 0x59, 0x8d,
	0x9d, 0xb6, 0x32, 0xd3, 0xc1, 0x71, 0xd7, 0xd0, 0x00, 0xda, 0xc5, 0x35, 0x85, 0x76, 0xe5, 0xbd,
	0x92, 0xed, 0xd5, 0x41, 0x32, 0x9a, 0x1b, 0x24, 0x85, 0x53, 0xdc, 0x38, 0x1a, 0xa7, 0x64, 0x11,
	0x95, 0xe0, 0xbc, 0x85, 0x56, 0x6e, 0xdf, 0xa0, 0xc7, 0x6a, 0x20, 0xef, 0xd8, 0x41, 0x25, 0x08,
	0x67, 0x70, 0x7f, 0x65, 0x4a, 0xd1, 0x5e, 0x86, 0xd2, 0xea, 0xf4, 0x76, 0x1c, 0x19, 0xbe, 0xab,
	0x6f, 0x47, 0xb0, 0x7d, 0x4a, 0x78, 0x2e, 0x45, 0xc4, 0x35, 0x60, 0xd9, 0x98, 0x94, 0x03, 0x5e,
	0xd7, 0xe4, 0x7f, 0xeb, 0xe5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61, 0x8e, 0x5d, 0xfd, 0xc8,
	0x06, 0x00, 0x00,
}
