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
	PayErrors_None                  PayErrors = 0
	PayErrors_PayWayNotFound        PayErrors = 2000001
	PayErrors_PayWayOperationFailed PayErrors = 2000002
	PayErrors_CreateOrderFailed     PayErrors = 2000003
	PayErrors_UpdateOrderFailed     PayErrors = 2000004
)

var PayErrors_name = map[int32]string{
	0:       "None",
	2000001: "PayWayNotFound",
	2000002: "PayWayOperationFailed",
	2000003: "CreateOrderFailed",
	2000004: "UpdateOrderFailed",
}

var PayErrors_value = map[string]int32{
	"None":                  0,
	"PayWayNotFound":        2000001,
	"PayWayOperationFailed": 2000002,
	"CreateOrderFailed":     2000003,
	"UpdateOrderFailed":     2000004,
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

type CreateOrderCommonRequest struct {
	CreateUserID         int64    `protobuf:"varint,1,opt,name=CreateUserID,proto3" json:"CreateUserID,omitempty"`
	OrderType            string   `protobuf:"bytes,2,opt,name=OrderType,proto3" json:"OrderType,omitempty"`
	CreateForUserID      int64    `protobuf:"varint,3,opt,name=CreateForUserID,proto3" json:"CreateForUserID,omitempty"`
	TotalAmount          string   `protobuf:"bytes,4,opt,name=TotalAmount,proto3" json:"TotalAmount,omitempty"`
	Subject              string   `protobuf:"bytes,5,opt,name=Subject,proto3" json:"Subject,omitempty"`
	OrderExtends         string   `protobuf:"bytes,6,opt,name=OrderExtends,proto3" json:"OrderExtends,omitempty"`
	OrderPlatform        string   `protobuf:"bytes,7,opt,name=OrderPlatform,proto3" json:"OrderPlatform,omitempty"`
	OrderPlatformMode    string   `protobuf:"bytes,8,opt,name=OrderPlatformMode,proto3" json:"OrderPlatformMode,omitempty"`
	QuitURL              string   `protobuf:"bytes,9,opt,name=QuitURL,proto3" json:"QuitURL,omitempty"`
	ReturnURL            string   `protobuf:"bytes,10,opt,name=ReturnURL,proto3" json:"ReturnURL,omitempty"`
	OpenID               string   `protobuf:"bytes,11,opt,name=OpenID,proto3" json:"OpenID,omitempty"`
	ClientIP             string   `protobuf:"bytes,12,opt,name=ClientIP,proto3" json:"ClientIP,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderCommonRequest) Reset()         { *m = CreateOrderCommonRequest{} }
func (m *CreateOrderCommonRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOrderCommonRequest) ProtoMessage()    {}
func (*CreateOrderCommonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_587e2c2585dcd479, []int{7}
}

func (m *CreateOrderCommonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderCommonRequest.Unmarshal(m, b)
}
func (m *CreateOrderCommonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderCommonRequest.Marshal(b, m, deterministic)
}
func (m *CreateOrderCommonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderCommonRequest.Merge(m, src)
}
func (m *CreateOrderCommonRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOrderCommonRequest.Size(m)
}
func (m *CreateOrderCommonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderCommonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderCommonRequest proto.InternalMessageInfo

func (m *CreateOrderCommonRequest) GetCreateUserID() int64 {
	if m != nil {
		return m.CreateUserID
	}
	return 0
}

func (m *CreateOrderCommonRequest) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

func (m *CreateOrderCommonRequest) GetCreateForUserID() int64 {
	if m != nil {
		return m.CreateForUserID
	}
	return 0
}

func (m *CreateOrderCommonRequest) GetTotalAmount() string {
	if m != nil {
		return m.TotalAmount
	}
	return ""
}

func (m *CreateOrderCommonRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *CreateOrderCommonRequest) GetOrderExtends() string {
	if m != nil {
		return m.OrderExtends
	}
	return ""
}

func (m *CreateOrderCommonRequest) GetOrderPlatform() string {
	if m != nil {
		return m.OrderPlatform
	}
	return ""
}

func (m *CreateOrderCommonRequest) GetOrderPlatformMode() string {
	if m != nil {
		return m.OrderPlatformMode
	}
	return ""
}

func (m *CreateOrderCommonRequest) GetQuitURL() string {
	if m != nil {
		return m.QuitURL
	}
	return ""
}

func (m *CreateOrderCommonRequest) GetReturnURL() string {
	if m != nil {
		return m.ReturnURL
	}
	return ""
}

func (m *CreateOrderCommonRequest) GetOpenID() string {
	if m != nil {
		return m.OpenID
	}
	return ""
}

func (m *CreateOrderCommonRequest) GetClientIP() string {
	if m != nil {
		return m.ClientIP
	}
	return ""
}

type PayOrderResultRequest struct {
	OutTradeNo           string   `protobuf:"bytes,1,opt,name=OutTradeNo,proto3" json:"OutTradeNo,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	RawData              string   `protobuf:"bytes,3,opt,name=RawData,proto3" json:"RawData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayOrderResultRequest) Reset()         { *m = PayOrderResultRequest{} }
func (m *PayOrderResultRequest) String() string { return proto.CompactTextString(m) }
func (*PayOrderResultRequest) ProtoMessage()    {}
func (*PayOrderResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_587e2c2585dcd479, []int{8}
}

func (m *PayOrderResultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayOrderResultRequest.Unmarshal(m, b)
}
func (m *PayOrderResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayOrderResultRequest.Marshal(b, m, deterministic)
}
func (m *PayOrderResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayOrderResultRequest.Merge(m, src)
}
func (m *PayOrderResultRequest) XXX_Size() int {
	return xxx_messageInfo_PayOrderResultRequest.Size(m)
}
func (m *PayOrderResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PayOrderResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PayOrderResultRequest proto.InternalMessageInfo

func (m *PayOrderResultRequest) GetOutTradeNo() string {
	if m != nil {
		return m.OutTradeNo
	}
	return ""
}

func (m *PayOrderResultRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *PayOrderResultRequest) GetRawData() string {
	if m != nil {
		return m.RawData
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

type GetOrderListRequest struct {
	OrderCreateByUID     string   `protobuf:"bytes,1,opt,name=OrderCreateByUID,proto3" json:"OrderCreateByUID,omitempty"`
	OrderType            string   `protobuf:"bytes,2,opt,name=OrderType,proto3" json:"OrderType,omitempty"`
	OrdercreateForUID    string   `protobuf:"bytes,3,opt,name=OrdercreateForUID,proto3" json:"OrdercreateForUID,omitempty"`
	Page                 int64    `protobuf:"varint,4,opt,name=Page,proto3" json:"Page,omitempty"`
	Size                 int64    `protobuf:"varint,5,opt,name=Size,proto3" json:"Size,omitempty"`
	Status               string   `protobuf:"bytes,6,opt,name=Status,proto3" json:"Status,omitempty"`
	OrderPlatForm        string   `protobuf:"bytes,7,opt,name=OrderPlatForm,proto3" json:"OrderPlatForm,omitempty"`
	OrderExtends         string   `protobuf:"bytes,9,opt,name=OrderExtends,proto3" json:"OrderExtends,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderListRequest) Reset()         { *m = GetOrderListRequest{} }
func (m *GetOrderListRequest) String() string { return proto.CompactTextString(m) }
func (*GetOrderListRequest) ProtoMessage()    {}
func (*GetOrderListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_587e2c2585dcd479, []int{10}
}

func (m *GetOrderListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderListRequest.Unmarshal(m, b)
}
func (m *GetOrderListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderListRequest.Marshal(b, m, deterministic)
}
func (m *GetOrderListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderListRequest.Merge(m, src)
}
func (m *GetOrderListRequest) XXX_Size() int {
	return xxx_messageInfo_GetOrderListRequest.Size(m)
}
func (m *GetOrderListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderListRequest proto.InternalMessageInfo

func (m *GetOrderListRequest) GetOrderCreateByUID() string {
	if m != nil {
		return m.OrderCreateByUID
	}
	return ""
}

func (m *GetOrderListRequest) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

func (m *GetOrderListRequest) GetOrdercreateForUID() string {
	if m != nil {
		return m.OrdercreateForUID
	}
	return ""
}

func (m *GetOrderListRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetOrderListRequest) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *GetOrderListRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *GetOrderListRequest) GetOrderPlatForm() string {
	if m != nil {
		return m.OrderPlatForm
	}
	return ""
}

func (m *GetOrderListRequest) GetOrderExtends() string {
	if m != nil {
		return m.OrderExtends
	}
	return ""
}

type OrderInfo struct {
	OutTradeNo           string   `protobuf:"bytes,1,opt,name=OutTradeNo,proto3" json:"OutTradeNo,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	PlatForm             string   `protobuf:"bytes,3,opt,name=PlatForm,proto3" json:"PlatForm,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	CreateByUID          string   `protobuf:"bytes,5,opt,name=CreateByUID,proto3" json:"CreateByUID,omitempty"`
	CreateForUID         string   `protobuf:"bytes,6,opt,name=CreateForUID,proto3" json:"CreateForUID,omitempty"`
	Extends              string   `protobuf:"bytes,7,opt,name=Extends,proto3" json:"Extends,omitempty"`
	Amount               string   `protobuf:"bytes,8,opt,name=Amount,proto3" json:"Amount,omitempty"`
	CreatedAt            int64    `protobuf:"varint,9,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Expired              bool     `protobuf:"varint,10,opt,name=Expired,proto3" json:"Expired,omitempty"`
	QRCode               string   `protobuf:"bytes,11,opt,name=QRCode,proto3" json:"QRCode,omitempty"`
	PlatFormMode         string   `protobuf:"bytes,12,opt,name=PlatFormMode,proto3" json:"PlatFormMode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderInfo) Reset()         { *m = OrderInfo{} }
func (m *OrderInfo) String() string { return proto.CompactTextString(m) }
func (*OrderInfo) ProtoMessage()    {}
func (*OrderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_587e2c2585dcd479, []int{11}
}

func (m *OrderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderInfo.Unmarshal(m, b)
}
func (m *OrderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderInfo.Marshal(b, m, deterministic)
}
func (m *OrderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderInfo.Merge(m, src)
}
func (m *OrderInfo) XXX_Size() int {
	return xxx_messageInfo_OrderInfo.Size(m)
}
func (m *OrderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OrderInfo proto.InternalMessageInfo

func (m *OrderInfo) GetOutTradeNo() string {
	if m != nil {
		return m.OutTradeNo
	}
	return ""
}

func (m *OrderInfo) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *OrderInfo) GetPlatForm() string {
	if m != nil {
		return m.PlatForm
	}
	return ""
}

func (m *OrderInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *OrderInfo) GetCreateByUID() string {
	if m != nil {
		return m.CreateByUID
	}
	return ""
}

func (m *OrderInfo) GetCreateForUID() string {
	if m != nil {
		return m.CreateForUID
	}
	return ""
}

func (m *OrderInfo) GetExtends() string {
	if m != nil {
		return m.Extends
	}
	return ""
}

func (m *OrderInfo) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *OrderInfo) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *OrderInfo) GetExpired() bool {
	if m != nil {
		return m.Expired
	}
	return false
}

func (m *OrderInfo) GetQRCode() string {
	if m != nil {
		return m.QRCode
	}
	return ""
}

func (m *OrderInfo) GetPlatFormMode() string {
	if m != nil {
		return m.PlatFormMode
	}
	return ""
}

type OrderListResponse struct {
	List                 []*OrderInfo `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *OrderListResponse) Reset()         { *m = OrderListResponse{} }
func (m *OrderListResponse) String() string { return proto.CompactTextString(m) }
func (*OrderListResponse) ProtoMessage()    {}
func (*OrderListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_587e2c2585dcd479, []int{12}
}

func (m *OrderListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderListResponse.Unmarshal(m, b)
}
func (m *OrderListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderListResponse.Marshal(b, m, deterministic)
}
func (m *OrderListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderListResponse.Merge(m, src)
}
func (m *OrderListResponse) XXX_Size() int {
	return xxx_messageInfo_OrderListResponse.Size(m)
}
func (m *OrderListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderListResponse proto.InternalMessageInfo

func (m *OrderListResponse) GetList() []*OrderInfo {
	if m != nil {
		return m.List
	}
	return nil
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
	proto.RegisterType((*CreateOrderCommonRequest)(nil), "pay.CreateOrderCommonRequest")
	proto.RegisterType((*PayOrderResultRequest)(nil), "pay.PayOrderResultRequest")
	proto.RegisterType((*OrderResultResponse)(nil), "pay.OrderResultResponse")
	proto.RegisterType((*GetOrderListRequest)(nil), "pay.GetOrderListRequest")
	proto.RegisterType((*OrderInfo)(nil), "pay.OrderInfo")
	proto.RegisterType((*OrderListResponse)(nil), "pay.OrderListResponse")
}

func init() {
	proto.RegisterFile("pay/pay.proto", fileDescriptor_587e2c2585dcd479)
}

var fileDescriptor_587e2c2585dcd479 = []byte{
	// 1031 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x5f, 0x6f, 0x1b, 0x45,
	0x10, 0xaf, 0x7d, 0xae, 0x6b, 0x8f, 0xff, 0xe0, 0x6e, 0x43, 0x72, 0x98, 0x80, 0xa2, 0x13, 0x48,
	0x51, 0x55, 0xb5, 0x52, 0x40, 0xe2, 0x05, 0x24, 0xda, 0xb8, 0x0e, 0x16, 0x69, 0xe2, 0x5e, 0xe2,
	0xf6, 0x79, 0x9b, 0xdb, 0x46, 0x87, 0x9c, 0xdb, 0xe3, 0x6e, 0x4f, 0x70, 0xe6, 0x89, 0x3f, 0x2f,
	0x79, 0x47, 0xe2, 0x1b, 0xf0, 0xb5, 0x78, 0xe4, 0x6b, 0xa0, 0x9d, 0xdd, 0xbd, 0xbb, 0xb5, 0x7d,
	0x80, 0xe0, 0x6d, 0x67, 0x66, 0x77, 0xfe, 0xfe, 0x66, 0x66, 0x61, 0x10, 0xd3, 0xfc, 0x49, 0x4c,
	0xf3, 0xc7, 0x71, 0xc2, 0x05, 0x27, 0x4e, 0x4c, 0x73, 0x6f, 0x1f, 0x3a, 0x8b, 0xd9, 0x64, 0x16,
	0xc5, 0x99, 0x20, 0x23, 0x70, 0x16, 0xb3, 0x89, 0xdb, 0x38, 0x68, 0x1c, 0x3a, 0xbe, 0x3c, 0x7a,
	0x02, 0xda, 0x73, 0x9a, 0xbf, 0xa6, 0x39, 0x19, 0x43, 0x47, 0x9d, 0x8a, 0x0b, 0x05, 0x4d, 0x5c,
	0xb8, 0x37, 0xa7, 0xf9, 0x65, 0x1e, 0x33, 0xb7, 0x79, 0xd0, 0x38, 0xec, 0xfa, 0x86, 0xd4, 0xaf,
	0x5e, 0xd1, 0x65, 0xc6, 0x5c, 0x07, 0x45, 0x05, 0xad, 0x5f, 0x4d, 0x58, 0x7a, 0xe5, 0xb6, 0x8a,
	0x57, 0x92, 0xf4, 0x7e, 0x80, 0xbd, 0xe3, 0x84, 0x51, 0xc1, 0x16, 0x29, 0x4b, 0x94, 0x15, 0x9f,
	0x7d, 0x9b, 0xb1, 0x54, 0x54, 0x4d, 0x35, 0xea, 0x4d, 0x35, 0xeb, 0x4d, 0x39, 0x96, 0x29, 0x13,
	0x72, 0xab, 0x0c, 0xf9, 0x04, 0xf6, 0x26, 0x6c, 0xc9, 0xb6, 0x19, 0xdf, 0xc8, 0x8f, 0x95, 0x95,
	0xa6, 0x9d, 0x15, 0xef, 0xd7, 0x06, 0xec, 0xcc, 0x33, 0xf1, 0x3f, 0xd5, 0x54, 0x23, 0x76, 0xea,
	0x23, 0x6e, 0xd5, 0x47, 0x7c, 0xd7, 0x4e, 0xee, 0xa7, 0xd0, 0x2b, 0x5d, 0x4a, 0xc9, 0xc7, 0x78,
	0x51, 0x1e, 0xdd, 0xc6, 0x81, 0x73, 0xd8, 0x3b, 0xea, 0x3d, 0x96, 0x08, 0xd1, 0x1e, 0x1b, 0x99,
	0xf7, 0x05, 0x0c, 0x9e, 0x27, 0x09, 0x4f, 0x7c, 0x96, 0xc6, 0x3c, 0x4a, 0x19, 0x79, 0x04, 0x5d,
	0x64, 0xcc, 0xa2, 0xb7, 0x1c, 0x43, 0x19, 0x1e, 0x0d, 0xcd, 0x4b, 0x14, 0xa4, 0x7e, 0x79, 0xc1,
	0xfb, 0xdd, 0x01, 0x57, 0x95, 0xf4, 0x3c, 0x09, 0x58, 0x72, 0xcc, 0x6f, 0x6e, 0x78, 0x64, 0xf2,
	0xe1, 0x41, 0xbf, 0x2c, 0x77, 0x91, 0x18, 0x8b, 0x47, 0xf6, 0xa1, 0x8b, 0x2f, 0x2b, 0x20, 0x2b,
	0x19, 0xe4, 0x10, 0xde, 0x51, 0xb7, 0xa7, 0x3c, 0xd1, 0x4a, 0x1c, 0x54, 0xb2, 0xce, 0x26, 0x07,
	0xd0, 0xbb, 0xe4, 0x82, 0x2e, 0x9f, 0xde, 0xf0, 0x2c, 0x12, 0x3a, 0x6d, 0x55, 0x96, 0xcc, 0xdc,
	0x45, 0xf6, 0xe6, 0x1b, 0x76, 0x25, 0x4c, 0xe6, 0x34, 0x29, 0xfd, 0x44, 0x93, 0xcf, 0xbf, 0x17,
	0x2c, 0x0a, 0x52, 0xb7, 0x8d, 0x62, 0x8b, 0x47, 0x3e, 0x82, 0x01, 0xd2, 0xf3, 0x25, 0x15, 0x6f,
	0x79, 0x72, 0xe3, 0xde, 0xc3, 0x4b, 0x36, 0x93, 0x3c, 0x82, 0xfb, 0x16, 0xe3, 0x05, 0x0f, 0x98,
	0xdb, 0xc1, 0x9b, 0x9b, 0x02, 0xe9, 0xd1, 0xcb, 0x2c, 0x14, 0x0b, 0xff, 0xd4, 0xed, 0x2a, 0x8f,
	0x34, 0x29, 0xb3, 0xe2, 0x33, 0x91, 0x25, 0x91, 0x94, 0x81, 0xca, 0x4a, 0xc1, 0x20, 0xbb, 0xd0,
	0x3e, 0x8f, 0x59, 0x34, 0x9b, 0xb8, 0x3d, 0x14, 0x69, 0x4a, 0xe2, 0xe6, 0x78, 0x19, 0xb2, 0x48,
	0xcc, 0xe6, 0x6e, 0x5f, 0xe1, 0xc6, 0xd0, 0x5e, 0x08, 0xef, 0xce, 0x69, 0x8e, 0x3e, 0xf8, 0x2c,
	0xcd, 0x96, 0xc2, 0x14, 0xe9, 0x43, 0x80, 0xf3, 0x4c, 0x5c, 0x26, 0x34, 0x60, 0x67, 0x5c, 0xf7,
	0x5e, 0x85, 0x23, 0x8d, 0x5d, 0x08, 0x2a, 0xb2, 0x54, 0x57, 0x47, 0x53, 0xd2, 0x79, 0x9f, 0x7e,
	0x37, 0xa1, 0x82, 0x1a, 0xf8, 0x6a, 0xd2, 0xbb, 0x6d, 0xc2, 0x03, 0xcb, 0x90, 0x46, 0x16, 0x81,
	0xd6, 0xb1, 0xcc, 0x87, 0xb2, 0x81, 0x67, 0xd9, 0x32, 0x2f, 0xd2, 0x6b, 0xad, 0x5a, 0x1e, 0xd7,
	0xfc, 0x71, 0xb6, 0xf9, 0xf3, 0xd2, 0x47, 0x3d, 0xaa, 0xc6, 0x9a, 0x92, 0xda, 0x2f, 0xc2, 0xeb,
	0x48, 0xd7, 0x16, 0xcf, 0x08, 0x0a, 0xf9, 0x4c, 0x07, 0xd0, 0xd6, 0xa0, 0x28, 0x59, 0x32, 0x0a,
	0x63, 0x4a, 0x15, 0xd4, 0x90, 0xeb, 0x80, 0xea, 0x6c, 0x02, 0xca, 0x83, 0xfe, 0xb3, 0x2c, 0x67,
	0xc9, 0x29, 0xbf, 0xe6, 0xb2, 0x18, 0xaa, 0x86, 0x16, 0xcf, 0xfb, 0xad, 0x09, 0x0f, 0x4e, 0x98,
	0xc0, 0x74, 0x9c, 0x86, 0x69, 0x91, 0xf5, 0x87, 0x30, 0x52, 0x0d, 0x83, 0x30, 0x7e, 0x96, 0x9b,
	0xb9, 0xd1, 0xf5, 0x37, 0xf8, 0xff, 0xd0, 0x22, 0x06, 0x72, 0x57, 0x45, 0x43, 0xe8, 0x26, 0x31,
	0x90, 0xab, 0x0a, 0x64, 0x96, 0xe6, 0xf4, 0x9a, 0xe9, 0xb9, 0x88, 0x67, 0x95, 0xb9, 0x15, 0xc3,
	0xcc, 0x39, 0x3e, 0x9e, 0x2b, 0x55, 0x6f, 0x5b, 0x55, 0xaf, 0xb6, 0xc1, 0x74, 0x5b, 0x1b, 0x48,
	0xe6, 0x46, 0x43, 0x75, 0x37, 0x1b, 0xca, 0xfb, 0xa3, 0xa9, 0xc3, 0x92, 0x73, 0xe4, 0x3f, 0xa3,
	0x50, 0x8e, 0x4a, 0xe3, 0x8a, 0xd9, 0x43, 0xc6, 0x0b, 0x02, 0x2d, 0x4c, 0x99, 0xc2, 0x09, 0x9e,
	0x65, 0x55, 0xab, 0x29, 0x57, 0x60, 0xa9, 0xb2, 0xca, 0xa1, 0xa5, 0x53, 0xa9, 0x87, 0x41, 0x95,
	0x27, 0x51, 0x63, 0x42, 0xd3, 0xa8, 0x31, 0x63, 0x62, 0x17, 0xda, 0x16, 0x60, 0x34, 0x25, 0x6b,
	0xa8, 0x34, 0x04, 0x4f, 0x05, 0xa6, 0xc3, 0xf1, 0x4b, 0x86, 0xd2, 0x17, 0x87, 0x09, 0x0b, 0xb0,
	0xd9, 0x3b, 0xbe, 0x21, 0x2b, 0x68, 0xef, 0x59, 0x68, 0xf7, 0xa0, 0x6f, 0xe2, 0xc4, 0x19, 0xa3,
	0xda, 0xdd, 0xe2, 0x79, 0x9f, 0x69, 0x64, 0x28, 0xdc, 0xe9, 0x26, 0xf4, 0xa0, 0x25, 0x69, 0xbd,
	0x13, 0xd4, 0x64, 0x2f, 0xca, 0xe0, 0xa3, 0xec, 0xe1, 0x0a, 0xba, 0xc5, 0xb0, 0x27, 0x1d, 0x68,
	0x9d, 0xf1, 0x88, 0x8d, 0xee, 0x90, 0x1d, 0x18, 0xaa, 0xad, 0x71, 0xc6, 0xc5, 0x94, 0x67, 0x51,
	0x30, 0xfa, 0xf1, 0x76, 0x45, 0xde, 0xc7, 0xc1, 0xf2, 0x9a, 0xe6, 0xe7, 0x31, 0x4b, 0xa8, 0x08,
	0x79, 0x34, 0xa5, 0xe1, 0x92, 0x05, 0xa3, 0x9f, 0x6e, 0x57, 0x64, 0x0f, 0xee, 0x57, 0xb6, 0x83,
	0x16, 0xfc, 0xac, 0x04, 0x8b, 0x38, 0x58, 0x13, 0xfc, 0x72, 0xbb, 0x3a, 0xfa, 0xb3, 0x05, 0x77,
	0x63, 0x9a, 0x5f, 0xbc, 0x22, 0x4f, 0x00, 0x4e, 0x98, 0x98, 0xd3, 0x1c, 0x01, 0x32, 0x40, 0x4f,
	0xcd, 0x8f, 0x66, 0x3c, 0x52, 0x64, 0xb9, 0xef, 0xbc, 0x3b, 0x64, 0x0a, 0xa3, 0xf5, 0xdf, 0x05,
	0xd9, 0xc7, 0x7b, 0x35, 0x9f, 0x8e, 0x31, 0x41, 0xa9, 0xb5, 0xff, 0x94, 0x9e, 0xf5, 0x8f, 0x82,
	0xd6, 0x53, 0xf3, 0x7f, 0xa8, 0xd1, 0xf3, 0x25, 0x0c, 0xac, 0x6f, 0x02, 0x79, 0x4f, 0xed, 0xd1,
	0x2d, 0x5f, 0x87, 0x1a, 0x0d, 0x5f, 0xc3, 0x50, 0xb9, 0x6e, 0x46, 0x37, 0xf9, 0xa0, 0x12, 0xcf,
	0xe6, 0xc6, 0x1d, 0xbb, 0x65, 0x3d, 0xed, 0xe1, 0xeb, 0xdd, 0x21, 0x5f, 0xc1, 0x8e, 0x4a, 0xb9,
	0xbd, 0x07, 0xc8, 0xd8, 0x6c, 0xf7, 0xcd, 0xe5, 0x50, 0x1b, 0x58, 0xbf, 0x3a, 0xd3, 0x88, 0xb2,
	0xba, 0x65, 0xcc, 0x8d, 0x77, 0x4b, 0x7f, 0x2c, 0x14, 0x7e, 0x0e, 0x43, 0x73, 0x7d, 0xc2, 0x04,
	0x0d, 0x97, 0x7f, 0xeb, 0xc5, 0x1a, 0x4a, 0xc9, 0x04, 0x76, 0xcd, 0x6b, 0x03, 0x78, 0x33, 0x0e,
	0xfe, 0xbd, 0x96, 0x37, 0x6d, 0xfc, 0x2c, 0x7f, 0xf2, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x28,
	0x92, 0x51, 0x34, 0x3d, 0x0b, 0x00, 0x00,
}
