// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: search/search.proto

package search

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 数据桶类型
type BucketTypes int32

const (
	BucketTypes_NoneType BucketTypes = 0
	BucketTypes_Article  BucketTypes = 1 // 文章类型
	BucketTypes_Author   BucketTypes = 2 // 作者
)

// Enum value maps for BucketTypes.
var (
	BucketTypes_name = map[int32]string{
		0: "NoneType",
		1: "Article",
		2: "Author",
	}
	BucketTypes_value = map[string]int32{
		"NoneType": 0,
		"Article":  1,
		"Author":   2,
	}
)

func (x BucketTypes) Enum() *BucketTypes {
	p := new(BucketTypes)
	*p = x
	return p
}

func (x BucketTypes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BucketTypes) Descriptor() protoreflect.EnumDescriptor {
	return file_search_search_proto_enumTypes[0].Descriptor()
}

func (BucketTypes) Type() protoreflect.EnumType {
	return &file_search_search_proto_enumTypes[0]
}

func (x BucketTypes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BucketTypes.Descriptor instead.
func (BucketTypes) EnumDescriptor() ([]byte, []int) {
	return file_search_search_proto_rawDescGZIP(), []int{0}
}

// 数据仓库类型
type CollectionTypes int32

const (
	CollectionTypes_Test  CollectionTypes = 0
	CollectionTypes_Happy CollectionTypes = 1
)

// Enum value maps for CollectionTypes.
var (
	CollectionTypes_name = map[int32]string{
		0: "Test",
		1: "Happy",
	}
	CollectionTypes_value = map[string]int32{
		"Test":  0,
		"Happy": 1,
	}
)

func (x CollectionTypes) Enum() *CollectionTypes {
	p := new(CollectionTypes)
	*p = x
	return p
}

func (x CollectionTypes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CollectionTypes) Descriptor() protoreflect.EnumDescriptor {
	return file_search_search_proto_enumTypes[1].Descriptor()
}

func (CollectionTypes) Type() protoreflect.EnumType {
	return &file_search_search_proto_enumTypes[1]
}

func (x CollectionTypes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CollectionTypes.Descriptor instead.
func (CollectionTypes) EnumDescriptor() ([]byte, []int) {
	return file_search_search_proto_rawDescGZIP(), []int{1}
}

// 错误类型
type SearchErrors int32

const (
	SearchErrors_NoneError         SearchErrors = 0       // 默认值
	SearchErrors_SearchCommonError SearchErrors = 4000001 // 通用error
)

// Enum value maps for SearchErrors.
var (
	SearchErrors_name = map[int32]string{
		0:       "NoneError",
		4000001: "SearchCommonError",
	}
	SearchErrors_value = map[string]int32{
		"NoneError":         0,
		"SearchCommonError": 4000001,
	}
)

func (x SearchErrors) Enum() *SearchErrors {
	p := new(SearchErrors)
	*p = x
	return p
}

func (x SearchErrors) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchErrors) Descriptor() protoreflect.EnumDescriptor {
	return file_search_search_proto_enumTypes[2].Descriptor()
}

func (SearchErrors) Type() protoreflect.EnumType {
	return &file_search_search_proto_enumTypes[2]
}

func (x SearchErrors) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchErrors.Descriptor instead.
func (SearchErrors) EnumDescriptor() ([]byte, []int) {
	return file_search_search_proto_rawDescGZIP(), []int{2}
}

// 搜索结果
type SearchResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []string `protobuf:"bytes,1,rep,name=Results,proto3" json:"Results,omitempty"`
}

func (x *SearchResult) Reset() {
	*x = SearchResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResult) ProtoMessage() {}

func (x *SearchResult) ProtoReflect() protoreflect.Message {
	mi := &file_search_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResult.ProtoReflect.Descriptor instead.
func (*SearchResult) Descriptor() ([]byte, []int) {
	return file_search_search_proto_rawDescGZIP(), []int{0}
}

func (x *SearchResult) GetResults() []string {
	if x != nil {
		return x.Results
	}
	return nil
}

// 删除搜索
type FlushSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collection CollectionTypes `protobuf:"varint,1,opt,name=Collection,proto3,enum=search.CollectionTypes" json:"Collection,omitempty"` // 类似于database的概念
	Bucket     BucketTypes     `protobuf:"varint,2,opt,name=Bucket,proto3,enum=search.BucketTypes" json:"Bucket,omitempty"`             // 类似于表的概念
	Object     string          `protobuf:"bytes,3,opt,name=Object,proto3" json:"Object,omitempty"`                                      // 这个类似于key的概念 搜索会根据value搜索key🔍
}

func (x *FlushSearchRequest) Reset() {
	*x = FlushSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlushSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlushSearchRequest) ProtoMessage() {}

func (x *FlushSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlushSearchRequest.ProtoReflect.Descriptor instead.
func (*FlushSearchRequest) Descriptor() ([]byte, []int) {
	return file_search_search_proto_rawDescGZIP(), []int{1}
}

func (x *FlushSearchRequest) GetCollection() CollectionTypes {
	if x != nil {
		return x.Collection
	}
	return CollectionTypes_Test
}

func (x *FlushSearchRequest) GetBucket() BucketTypes {
	if x != nil {
		return x.Bucket
	}
	return BucketTypes_NoneType
}

func (x *FlushSearchRequest) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorInfo SearchErrors `protobuf:"varint,1,opt,name=ErrorInfo,proto3,enum=search.SearchErrors" json:"ErrorInfo,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_search_proto_msgTypes[2]
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
	return file_search_search_proto_rawDescGZIP(), []int{2}
}

func (x *ErrorResponse) GetErrorInfo() SearchErrors {
	if x != nil {
		return x.ErrorInfo
	}
	return SearchErrors_NoneError
}

// 插入搜索数据请求
type PostSearchSourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collection CollectionTypes `protobuf:"varint,1,opt,name=Collection,proto3,enum=search.CollectionTypes" json:"Collection,omitempty"` // 类似于database的概念
	Bucket     BucketTypes     `protobuf:"varint,2,opt,name=Bucket,proto3,enum=search.BucketTypes" json:"Bucket,omitempty"`             // 类似于表的概念
	Object     string          `protobuf:"bytes,3,opt,name=Object,proto3" json:"Object,omitempty"`                                      // 这个类似于key的概念 搜索会根据value搜索key🔍
	Text       string          `protobuf:"bytes,4,opt,name=Text,proto3" json:"Text,omitempty"`                                          // 数据分词内容
}

func (x *PostSearchSourceRequest) Reset() {
	*x = PostSearchSourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostSearchSourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostSearchSourceRequest) ProtoMessage() {}

func (x *PostSearchSourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostSearchSourceRequest.ProtoReflect.Descriptor instead.
func (*PostSearchSourceRequest) Descriptor() ([]byte, []int) {
	return file_search_search_proto_rawDescGZIP(), []int{3}
}

func (x *PostSearchSourceRequest) GetCollection() CollectionTypes {
	if x != nil {
		return x.Collection
	}
	return CollectionTypes_Test
}

func (x *PostSearchSourceRequest) GetBucket() BucketTypes {
	if x != nil {
		return x.Bucket
	}
	return BucketTypes_NoneType
}

func (x *PostSearchSourceRequest) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *PostSearchSourceRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// 获取搜索结果
type GetSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collection CollectionTypes `protobuf:"varint,1,opt,name=Collection,proto3,enum=search.CollectionTypes" json:"Collection,omitempty"` // 类似于database的概念
	Bucket     BucketTypes     `protobuf:"varint,2,opt,name=Bucket,proto3,enum=search.BucketTypes" json:"Bucket,omitempty"`             // 类似于表的概念
	Text       string          `protobuf:"bytes,3,opt,name=Text,proto3" json:"Text,omitempty"`                                          // 搜索数据
}

func (x *GetSearchRequest) Reset() {
	*x = GetSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSearchRequest) ProtoMessage() {}

func (x *GetSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_search_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSearchRequest.ProtoReflect.Descriptor instead.
func (*GetSearchRequest) Descriptor() ([]byte, []int) {
	return file_search_search_proto_rawDescGZIP(), []int{4}
}

func (x *GetSearchRequest) GetCollection() CollectionTypes {
	if x != nil {
		return x.Collection
	}
	return CollectionTypes_Test
}

func (x *GetSearchRequest) GetBucket() BucketTypes {
	if x != nil {
		return x.Bucket
	}
	return BucketTypes_NoneType
}

func (x *GetSearchRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_search_search_proto protoreflect.FileDescriptor

var file_search_search_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x28, 0x0a,
	0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x12, 0x46, 0x6c, 0x75, 0x73,
	0x68, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37,
	0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x0a, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x06, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x43, 0x0a, 0x0d,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x52, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0xab, 0x01, 0x0a, 0x17, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a,
	0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x0a, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x06, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x22,
	0x8c, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x52, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a,
	0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x52, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65,
	0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x2a, 0x34,
	0x0a, 0x0b, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x0c, 0x0a,
	0x08, 0x4e, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x10, 0x02, 0x2a, 0x26, 0x0a, 0x0f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x61, 0x70, 0x70, 0x79, 0x10, 0x01, 0x2a, 0x37, 0x0a, 0x0c,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x0d, 0x0a, 0x09,
	0x4e, 0x6f, 0x6e, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x11, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x10, 0x81, 0x92, 0xf4, 0x01, 0x32, 0xfb, 0x02, 0x0a, 0x08, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x53, 0x56, 0x12, 0x41, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x48, 0x0a, 0x13, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x48, 0x0a, 0x13, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x46, 0x6c, 0x75, 0x73, 0x68, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x17, 0x46, 0x6c, 0x75,
	0x73, 0x68, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x46, 0x6c,
	0x75, 0x73, 0x68, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_search_search_proto_rawDescOnce sync.Once
	file_search_search_proto_rawDescData = file_search_search_proto_rawDesc
)

func file_search_search_proto_rawDescGZIP() []byte {
	file_search_search_proto_rawDescOnce.Do(func() {
		file_search_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_search_search_proto_rawDescData)
	})
	return file_search_search_proto_rawDescData
}

var file_search_search_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_search_search_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_search_search_proto_goTypes = []interface{}{
	(BucketTypes)(0),                // 0: search.BucketTypes
	(CollectionTypes)(0),            // 1: search.CollectionTypes
	(SearchErrors)(0),               // 2: search.SearchErrors
	(*SearchResult)(nil),            // 3: search.SearchResult
	(*FlushSearchRequest)(nil),      // 4: search.FlushSearchRequest
	(*ErrorResponse)(nil),           // 5: search.ErrorResponse
	(*PostSearchSourceRequest)(nil), // 6: search.PostSearchSourceRequest
	(*GetSearchRequest)(nil),        // 7: search.GetSearchRequest
}
var file_search_search_proto_depIdxs = []int32{
	1,  // 0: search.FlushSearchRequest.Collection:type_name -> search.CollectionTypes
	0,  // 1: search.FlushSearchRequest.Bucket:type_name -> search.BucketTypes
	2,  // 2: search.ErrorResponse.ErrorInfo:type_name -> search.SearchErrors
	1,  // 3: search.PostSearchSourceRequest.Collection:type_name -> search.CollectionTypes
	0,  // 4: search.PostSearchSourceRequest.Bucket:type_name -> search.BucketTypes
	1,  // 5: search.GetSearchRequest.Collection:type_name -> search.CollectionTypes
	0,  // 6: search.GetSearchRequest.Bucket:type_name -> search.BucketTypes
	7,  // 7: search.searchSV.GetSearchByType:input_type -> search.GetSearchRequest
	4,  // 8: search.searchSV.FlushSearchByObject:input_type -> search.FlushSearchRequest
	4,  // 9: search.searchSV.FlushSearchByBucket:input_type -> search.FlushSearchRequest
	4,  // 10: search.searchSV.FlushSearchByCollection:input_type -> search.FlushSearchRequest
	6,  // 11: search.searchSV.PostSearchSource:input_type -> search.PostSearchSourceRequest
	3,  // 12: search.searchSV.GetSearchByType:output_type -> search.SearchResult
	5,  // 13: search.searchSV.FlushSearchByObject:output_type -> search.ErrorResponse
	5,  // 14: search.searchSV.FlushSearchByBucket:output_type -> search.ErrorResponse
	5,  // 15: search.searchSV.FlushSearchByCollection:output_type -> search.ErrorResponse
	5,  // 16: search.searchSV.PostSearchSource:output_type -> search.ErrorResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_search_search_proto_init() }
func file_search_search_proto_init() {
	if File_search_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_search_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResult); i {
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
		file_search_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlushSearchRequest); i {
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
		file_search_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_search_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostSearchSourceRequest); i {
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
		file_search_search_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSearchRequest); i {
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
			RawDescriptor: file_search_search_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_search_search_proto_goTypes,
		DependencyIndexes: file_search_search_proto_depIdxs,
		EnumInfos:         file_search_search_proto_enumTypes,
		MessageInfos:      file_search_search_proto_msgTypes,
	}.Build()
	File_search_search_proto = out.File
	file_search_search_proto_rawDesc = nil
	file_search_search_proto_goTypes = nil
	file_search_search_proto_depIdxs = nil
}
