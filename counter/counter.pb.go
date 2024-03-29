// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: counter/counter.proto

package counter

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CounterErrors int32

const (
	CounterErrors_None         CounterErrors = 0       // 默认值
	CounterErrors_UpdateFailed CounterErrors = 5000001 // 更新失败
)

// Enum value maps for CounterErrors.
var (
	CounterErrors_name = map[int32]string{
		0:       "None",
		5000001: "UpdateFailed",
	}
	CounterErrors_value = map[string]int32{
		"None":         0,
		"UpdateFailed": 5000001,
	}
)

func (x CounterErrors) Enum() *CounterErrors {
	p := new(CounterErrors)
	*p = x
	return p
}

func (x CounterErrors) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CounterErrors) Descriptor() protoreflect.EnumDescriptor {
	return file_counter_counter_proto_enumTypes[0].Descriptor()
}

func (CounterErrors) Type() protoreflect.EnumType {
	return &file_counter_counter_proto_enumTypes[0]
}

func (x CounterErrors) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CounterErrors.Descriptor instead.
func (CounterErrors) EnumDescriptor() ([]byte, []int) {
	return file_counter_counter_proto_rawDescGZIP(), []int{0}
}

type IntRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *IntRequest) Reset() {
	*x = IntRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_counter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntRequest) ProtoMessage() {}

func (x *IntRequest) ProtoReflect() protoreflect.Message {
	mi := &file_counter_counter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntRequest.ProtoReflect.Descriptor instead.
func (*IntRequest) Descriptor() ([]byte, []int) {
	return file_counter_counter_proto_rawDescGZIP(), []int{0}
}

func (x *IntRequest) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorInfo CounterErrors `protobuf:"varint,1,opt,name=ErrorInfo,proto3,enum=counter.CounterErrors" json:"ErrorInfo,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_counter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_counter_counter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_counter_counter_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorResponse) GetErrorInfo() CounterErrors {
	if x != nil {
		return x.ErrorInfo
	}
	return CounterErrors_None
}

type ArticleCountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment     uint64  `protobuf:"varint,1,opt,name=Comment,proto3" json:"Comment,omitempty"`          // 评论次数
	Like        uint64  `protobuf:"varint,2,opt,name=Like,proto3" json:"Like,omitempty"`                // 喜欢次数
	Collect     uint64  `protobuf:"varint,3,opt,name=Collect,proto3" json:"Collect,omitempty"`          // 收藏次数
	DisLike     uint64  `protobuf:"varint,4,opt,name=DisLike,proto3" json:"DisLike,omitempty"`          // 不喜欢次数
	PageView    uint64  `protobuf:"varint,5,opt,name=PageView,proto3" json:"PageView,omitempty"`        // 浏览量
	RewardTimes uint64  `protobuf:"varint,6,opt,name=RewardTimes,proto3" json:"RewardTimes,omitempty"`  // 打赏次数
	RewardTotal float64 `protobuf:"fixed64,7,opt,name=RewardTotal,proto3" json:"RewardTotal,omitempty"` // 打赏总金额
}

func (x *ArticleCountInfo) Reset() {
	*x = ArticleCountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_counter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleCountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleCountInfo) ProtoMessage() {}

func (x *ArticleCountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_counter_counter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleCountInfo.ProtoReflect.Descriptor instead.
func (*ArticleCountInfo) Descriptor() ([]byte, []int) {
	return file_counter_counter_proto_rawDescGZIP(), []int{2}
}

func (x *ArticleCountInfo) GetComment() uint64 {
	if x != nil {
		return x.Comment
	}
	return 0
}

func (x *ArticleCountInfo) GetLike() uint64 {
	if x != nil {
		return x.Like
	}
	return 0
}

func (x *ArticleCountInfo) GetCollect() uint64 {
	if x != nil {
		return x.Collect
	}
	return 0
}

func (x *ArticleCountInfo) GetDisLike() uint64 {
	if x != nil {
		return x.DisLike
	}
	return 0
}

func (x *ArticleCountInfo) GetPageView() uint64 {
	if x != nil {
		return x.PageView
	}
	return 0
}

func (x *ArticleCountInfo) GetRewardTimes() uint64 {
	if x != nil {
		return x.RewardTimes
	}
	return 0
}

func (x *ArticleCountInfo) GetRewardTotal() float64 {
	if x != nil {
		return x.RewardTotal
	}
	return 0
}

type UserCountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Articles    uint64  `protobuf:"varint,1,opt,name=Articles,proto3" json:"Articles,omitempty"`        // 文章个数
	Comments    uint64  `protobuf:"varint,2,opt,name=Comments,proto3" json:"Comments,omitempty"`        // 评论个数
	Commented   uint64  `protobuf:"varint,3,opt,name=Commented,proto3" json:"Commented,omitempty"`      // 被评论个数
	RewardTimes uint64  `protobuf:"varint,4,opt,name=RewardTimes,proto3" json:"RewardTimes,omitempty"`  // 打赏次数
	RewardTotal float64 `protobuf:"fixed64,5,opt,name=RewardTotal,proto3" json:"RewardTotal,omitempty"` // 打赏总金额
}

func (x *UserCountInfo) Reset() {
	*x = UserCountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_counter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCountInfo) ProtoMessage() {}

func (x *UserCountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_counter_counter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCountInfo.ProtoReflect.Descriptor instead.
func (*UserCountInfo) Descriptor() ([]byte, []int) {
	return file_counter_counter_proto_rawDescGZIP(), []int{3}
}

func (x *UserCountInfo) GetArticles() uint64 {
	if x != nil {
		return x.Articles
	}
	return 0
}

func (x *UserCountInfo) GetComments() uint64 {
	if x != nil {
		return x.Comments
	}
	return 0
}

func (x *UserCountInfo) GetCommented() uint64 {
	if x != nil {
		return x.Commented
	}
	return 0
}

func (x *UserCountInfo) GetRewardTimes() uint64 {
	if x != nil {
		return x.RewardTimes
	}
	return 0
}

func (x *UserCountInfo) GetRewardTotal() float64 {
	if x != nil {
		return x.RewardTotal
	}
	return 0
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=Field,proto3" json:"Field,omitempty"` // 字段
	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"` // 字段加减值
	ID    uint64 `protobuf:"varint,3,opt,name=ID,proto3" json:"ID,omitempty"`      // 请求资源ID
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_counter_counter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_counter_counter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_counter_counter_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRequest) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *UpdateRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *UpdateRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

var File_counter_counter_proto protoreflect.FileDescriptor

var file_counter_counter_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x22, 0x22, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x45, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x52, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x10,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69,
	0x6b, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x4c,
	0x69, 0x6b, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x44, 0x69, 0x73, 0x4c, 0x69,
	0x6b, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x56, 0x69, 0x65, 0x77, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x56, 0x69, 0x65, 0x77, 0x12, 0x20,
	0x0a, 0x0b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x22, 0xa9, 0x01, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x4b,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x2a, 0x2e, 0x0a, 0x0d, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0xc1, 0x96, 0xb1, 0x02, 0x32, 0x94, 0x02, 0x0a, 0x09,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x56, 0x12, 0x41, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3b, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x41, 0x0a, 0x0f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x3b,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_counter_counter_proto_rawDescOnce sync.Once
	file_counter_counter_proto_rawDescData = file_counter_counter_proto_rawDesc
)

func file_counter_counter_proto_rawDescGZIP() []byte {
	file_counter_counter_proto_rawDescOnce.Do(func() {
		file_counter_counter_proto_rawDescData = protoimpl.X.CompressGZIP(file_counter_counter_proto_rawDescData)
	})
	return file_counter_counter_proto_rawDescData
}

var file_counter_counter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_counter_counter_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_counter_counter_proto_goTypes = []interface{}{
	(CounterErrors)(0),       // 0: counter.CounterErrors
	(*IntRequest)(nil),       // 1: counter.IntRequest
	(*ErrorResponse)(nil),    // 2: counter.ErrorResponse
	(*ArticleCountInfo)(nil), // 3: counter.ArticleCountInfo
	(*UserCountInfo)(nil),    // 4: counter.UserCountInfo
	(*UpdateRequest)(nil),    // 5: counter.UpdateRequest
}
var file_counter_counter_proto_depIdxs = []int32{
	0, // 0: counter.ErrorResponse.ErrorInfo:type_name -> counter.CounterErrors
	1, // 1: counter.counterSV.GetArticleCount:input_type -> counter.IntRequest
	1, // 2: counter.counterSV.GetUserCount:input_type -> counter.IntRequest
	5, // 3: counter.counterSV.UpdateUserCount:input_type -> counter.UpdateRequest
	5, // 4: counter.counterSV.UpdateArticleCount:input_type -> counter.UpdateRequest
	3, // 5: counter.counterSV.GetArticleCount:output_type -> counter.ArticleCountInfo
	4, // 6: counter.counterSV.GetUserCount:output_type -> counter.UserCountInfo
	2, // 7: counter.counterSV.UpdateUserCount:output_type -> counter.ErrorResponse
	2, // 8: counter.counterSV.UpdateArticleCount:output_type -> counter.ErrorResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_counter_counter_proto_init() }
func file_counter_counter_proto_init() {
	if File_counter_counter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_counter_counter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_counter_counter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_counter_counter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleCountInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_counter_counter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCountInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_counter_counter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_counter_counter_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_counter_counter_proto_goTypes,
		DependencyIndexes: file_counter_counter_proto_depIdxs,
		EnumInfos:         file_counter_counter_proto_enumTypes,
		MessageInfos:      file_counter_counter_proto_msgTypes,
	}.Build()
	File_counter_counter_proto = out.File
	file_counter_counter_proto_rawDesc = nil
	file_counter_counter_proto_goTypes = nil
	file_counter_counter_proto_depIdxs = nil
}
