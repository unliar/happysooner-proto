// Code generated by protoc-gen-go. DO NOT EDIT.
// source: writing/writing.proto

// 内容创作模块

package writing

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

type ArticleInfo struct {
	Id                   int64              `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title                string             `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Summary              string             `protobuf:"bytes,3,opt,name=Summary,proto3" json:"Summary,omitempty"`
	Tags                 []*ArticleTag      `protobuf:"bytes,4,rep,name=Tags,proto3" json:"Tags,omitempty"`
	CreatedAt            int64              `protobuf:"varint,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Status               int64              `protobuf:"varint,6,opt,name=Status,proto3" json:"Status,omitempty"`
	Content              string             `protobuf:"bytes,7,opt,name=Content,proto3" json:"Content,omitempty"`
	Category             *Category          `protobuf:"bytes,8,opt,name=Category,proto3" json:"Category,omitempty"`
	ArticleCountInfo     *ArticleCountInfo  `protobuf:"bytes,9,opt,name=ArticleCountInfo,proto3" json:"ArticleCountInfo,omitempty"`
	ArticleAuthorInfo    *ArticleAuthorInfo `protobuf:"bytes,10,opt,name=ArticleAuthorInfo,proto3" json:"ArticleAuthorInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ArticleInfo) Reset()         { *m = ArticleInfo{} }
func (m *ArticleInfo) String() string { return proto.CompactTextString(m) }
func (*ArticleInfo) ProtoMessage()    {}
func (*ArticleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_db9b9589ec5f9fbc, []int{0}
}

func (m *ArticleInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleInfo.Unmarshal(m, b)
}
func (m *ArticleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleInfo.Marshal(b, m, deterministic)
}
func (m *ArticleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleInfo.Merge(m, src)
}
func (m *ArticleInfo) XXX_Size() int {
	return xxx_messageInfo_ArticleInfo.Size(m)
}
func (m *ArticleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleInfo proto.InternalMessageInfo

func (m *ArticleInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArticleInfo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ArticleInfo) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *ArticleInfo) GetTags() []*ArticleTag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ArticleInfo) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *ArticleInfo) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ArticleInfo) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ArticleInfo) GetCategory() *Category {
	if m != nil {
		return m.Category
	}
	return nil
}

func (m *ArticleInfo) GetArticleCountInfo() *ArticleCountInfo {
	if m != nil {
		return m.ArticleCountInfo
	}
	return nil
}

func (m *ArticleInfo) GetArticleAuthorInfo() *ArticleAuthorInfo {
	if m != nil {
		return m.ArticleAuthorInfo
	}
	return nil
}

// Tag解构体
type ArticleTag struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Tag                  string   `protobuf:"bytes,2,opt,name=Tag,proto3" json:"Tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleTag) Reset()         { *m = ArticleTag{} }
func (m *ArticleTag) String() string { return proto.CompactTextString(m) }
func (*ArticleTag) ProtoMessage()    {}
func (*ArticleTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_db9b9589ec5f9fbc, []int{1}
}

func (m *ArticleTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleTag.Unmarshal(m, b)
}
func (m *ArticleTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleTag.Marshal(b, m, deterministic)
}
func (m *ArticleTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleTag.Merge(m, src)
}
func (m *ArticleTag) XXX_Size() int {
	return xxx_messageInfo_ArticleTag.Size(m)
}
func (m *ArticleTag) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleTag.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleTag proto.InternalMessageInfo

func (m *ArticleTag) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArticleTag) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type IntRequest struct {
	Value                int64    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntRequest) Reset()         { *m = IntRequest{} }
func (m *IntRequest) String() string { return proto.CompactTextString(m) }
func (*IntRequest) ProtoMessage()    {}
func (*IntRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db9b9589ec5f9fbc, []int{2}
}

func (m *IntRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntRequest.Unmarshal(m, b)
}
func (m *IntRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntRequest.Marshal(b, m, deterministic)
}
func (m *IntRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntRequest.Merge(m, src)
}
func (m *IntRequest) XXX_Size() int {
	return xxx_messageInfo_IntRequest.Size(m)
}
func (m *IntRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IntRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IntRequest proto.InternalMessageInfo

func (m *IntRequest) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// 数据统计信息
type ArticleCountInfo struct {
	Comment              int64    `protobuf:"varint,1,opt,name=Comment,proto3" json:"Comment,omitempty"`
	Like                 int64    `protobuf:"varint,2,opt,name=Like,proto3" json:"Like,omitempty"`
	Collect              int64    `protobuf:"varint,3,opt,name=Collect,proto3" json:"Collect,omitempty"`
	DisLike              int64    `protobuf:"varint,4,opt,name=DisLike,proto3" json:"DisLike,omitempty"`
	PageView             int64    `protobuf:"varint,5,opt,name=PageView,proto3" json:"PageView,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleCountInfo) Reset()         { *m = ArticleCountInfo{} }
func (m *ArticleCountInfo) String() string { return proto.CompactTextString(m) }
func (*ArticleCountInfo) ProtoMessage()    {}
func (*ArticleCountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_db9b9589ec5f9fbc, []int{3}
}

func (m *ArticleCountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleCountInfo.Unmarshal(m, b)
}
func (m *ArticleCountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleCountInfo.Marshal(b, m, deterministic)
}
func (m *ArticleCountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleCountInfo.Merge(m, src)
}
func (m *ArticleCountInfo) XXX_Size() int {
	return xxx_messageInfo_ArticleCountInfo.Size(m)
}
func (m *ArticleCountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleCountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleCountInfo proto.InternalMessageInfo

func (m *ArticleCountInfo) GetComment() int64 {
	if m != nil {
		return m.Comment
	}
	return 0
}

func (m *ArticleCountInfo) GetLike() int64 {
	if m != nil {
		return m.Like
	}
	return 0
}

func (m *ArticleCountInfo) GetCollect() int64 {
	if m != nil {
		return m.Collect
	}
	return 0
}

func (m *ArticleCountInfo) GetDisLike() int64 {
	if m != nil {
		return m.DisLike
	}
	return 0
}

func (m *ArticleCountInfo) GetPageView() int64 {
	if m != nil {
		return m.PageView
	}
	return 0
}

// 作者信息
type ArticleAuthorInfo struct {
	AuthorUID            int64    `protobuf:"varint,1,opt,name=AuthorUID,proto3" json:"AuthorUID,omitempty"`
	AuthorNickname       string   `protobuf:"bytes,2,opt,name=AuthorNickname,proto3" json:"AuthorNickname,omitempty"`
	AuthorAvatar         string   `protobuf:"bytes,3,opt,name=AuthorAvatar,proto3" json:"AuthorAvatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleAuthorInfo) Reset()         { *m = ArticleAuthorInfo{} }
func (m *ArticleAuthorInfo) String() string { return proto.CompactTextString(m) }
func (*ArticleAuthorInfo) ProtoMessage()    {}
func (*ArticleAuthorInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_db9b9589ec5f9fbc, []int{4}
}

func (m *ArticleAuthorInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleAuthorInfo.Unmarshal(m, b)
}
func (m *ArticleAuthorInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleAuthorInfo.Marshal(b, m, deterministic)
}
func (m *ArticleAuthorInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleAuthorInfo.Merge(m, src)
}
func (m *ArticleAuthorInfo) XXX_Size() int {
	return xxx_messageInfo_ArticleAuthorInfo.Size(m)
}
func (m *ArticleAuthorInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleAuthorInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleAuthorInfo proto.InternalMessageInfo

func (m *ArticleAuthorInfo) GetAuthorUID() int64 {
	if m != nil {
		return m.AuthorUID
	}
	return 0
}

func (m *ArticleAuthorInfo) GetAuthorNickname() string {
	if m != nil {
		return m.AuthorNickname
	}
	return ""
}

func (m *ArticleAuthorInfo) GetAuthorAvatar() string {
	if m != nil {
		return m.AuthorAvatar
	}
	return ""
}

// 文章类别
type Category struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	CN                   string   `protobuf:"bytes,2,opt,name=CN,proto3" json:"CN,omitempty"`
	EN                   string   `protobuf:"bytes,3,opt,name=EN,proto3" json:"EN,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_db9b9589ec5f9fbc, []int{5}
}

func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (m *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(m, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Category) GetCN() string {
	if m != nil {
		return m.CN
	}
	return ""
}

func (m *Category) GetEN() string {
	if m != nil {
		return m.EN
	}
	return ""
}

func init() {
	proto.RegisterType((*ArticleInfo)(nil), "writing.ArticleInfo")
	proto.RegisterType((*ArticleTag)(nil), "writing.ArticleTag")
	proto.RegisterType((*IntRequest)(nil), "writing.IntRequest")
	proto.RegisterType((*ArticleCountInfo)(nil), "writing.ArticleCountInfo")
	proto.RegisterType((*ArticleAuthorInfo)(nil), "writing.ArticleAuthorInfo")
	proto.RegisterType((*Category)(nil), "writing.Category")
}

func init() { proto.RegisterFile("writing/writing.proto", fileDescriptor_db9b9589ec5f9fbc) }

var fileDescriptor_db9b9589ec5f9fbc = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x3d, 0x8f, 0xd3, 0x40,
	0x10, 0xc5, 0x76, 0x2e, 0x1f, 0x13, 0x14, 0xdd, 0x0d, 0x07, 0x5a, 0x22, 0x8a, 0xc8, 0x05, 0xa4,
	0x21, 0x48, 0xa1, 0x43, 0xa2, 0x08, 0xbe, 0x08, 0x8c, 0x90, 0x85, 0x9c, 0x90, 0x7e, 0x49, 0x16,
	0x63, 0x9d, 0x3f, 0xc0, 0x19, 0x73, 0x4a, 0x41, 0x4b, 0xcd, 0x4f, 0x46, 0xde, 0x1d, 0xc7, 0xc2,
	0xbe, 0xca, 0xfb, 0xde, 0xbc, 0x99, 0xd9, 0x9d, 0x37, 0x86, 0xc7, 0x77, 0x45, 0x4c, 0x71, 0x16,
	0xbd, 0xe2, 0xef, 0xe2, 0x47, 0x91, 0x53, 0x8e, 0x03, 0x86, 0xee, 0x1f, 0x07, 0xc6, 0xab, 0x82,
	0xe2, 0x7d, 0xa2, 0xfc, 0xec, 0x5b, 0x8e, 0x13, 0xb0, 0xfd, 0x83, 0xb0, 0x66, 0xd6, 0xdc, 0x09,
	0x6d, 0xff, 0x80, 0xd7, 0x70, 0xb1, 0x8d, 0x29, 0x51, 0xc2, 0x9e, 0x59, 0xf3, 0x51, 0x68, 0x00,
	0x0a, 0x18, 0x6c, 0xca, 0x34, 0x95, 0xc5, 0x49, 0x38, 0x9a, 0xaf, 0x21, 0xbe, 0x80, 0xde, 0x56,
	0x46, 0x47, 0xd1, 0x9b, 0x39, 0xf3, 0xf1, 0xf2, 0xd1, 0xa2, 0x6e, 0xcb, 0x3d, 0xb6, 0x32, 0x0a,
	0xb5, 0x00, 0x9f, 0xc1, 0xc8, 0x2b, 0x94, 0x24, 0x75, 0x58, 0x91, 0xb8, 0xd0, 0xfd, 0x1a, 0x02,
	0x9f, 0x40, 0x7f, 0x43, 0x92, 0xca, 0xa3, 0xe8, 0xeb, 0x10, 0xa3, 0xaa, 0xb1, 0x97, 0x67, 0xa4,
	0x32, 0x12, 0x03, 0xd3, 0x98, 0x21, 0xbe, 0x84, 0xa1, 0x27, 0x49, 0x45, 0x79, 0x71, 0x12, 0xc3,
	0x99, 0x35, 0x1f, 0x2f, 0xaf, 0xce, 0xcd, 0xeb, 0x40, 0x78, 0x96, 0xe0, 0x1a, 0x2e, 0xf9, 0x4a,
	0x5e, 0x5e, 0x66, 0x54, 0xbd, 0x5d, 0x8c, 0x74, 0xda, 0xd3, 0xf6, 0x9d, 0xcf, 0x82, 0xb0, 0x93,
	0x82, 0x1f, 0xe0, 0x8a, 0xb9, 0x55, 0x49, 0xdf, 0xf3, 0x42, 0xd7, 0x01, 0x5d, 0x67, 0xda, 0xae,
	0xd3, 0x28, 0xc2, 0x6e, 0x92, 0xbb, 0x00, 0x68, 0x66, 0xd4, 0xb1, 0xe1, 0x12, 0x9c, 0xad, 0x8c,
	0xd8, 0x84, 0xea, 0xe8, 0xba, 0x00, 0x7e, 0x46, 0xa1, 0xfa, 0x59, 0xaa, 0x23, 0x55, 0x36, 0xed,
	0x64, 0x52, 0x2a, 0x4e, 0x31, 0xc0, 0xfd, 0x6b, 0x75, 0x5f, 0x69, 0x46, 0x98, 0xa6, 0xd5, 0x08,
	0x8d, 0xb8, 0x86, 0x88, 0xd0, 0xfb, 0x14, 0xdf, 0x1a, 0xab, 0x9d, 0x50, 0x9f, 0x8d, 0x3a, 0x49,
	0xd4, 0x9e, 0xb4, 0xd3, 0x5a, 0xad, 0x61, 0x15, 0xb9, 0x89, 0x8f, 0x3a, 0xa1, 0x67, 0x22, 0x0c,
	0x71, 0x0a, 0xc3, 0xcf, 0x32, 0x52, 0xbb, 0x58, 0xdd, 0xb1, 0xb3, 0x67, 0xec, 0xfe, 0xbe, 0x67,
	0x60, 0xd5, 0x2e, 0x18, 0xf4, 0xc5, 0xbf, 0xe1, 0x4b, 0x35, 0x04, 0x3e, 0x87, 0x89, 0x01, 0x41,
	0xbc, 0xbf, 0xcd, 0x64, 0x5a, 0xef, 0x62, 0x8b, 0x45, 0x17, 0x1e, 0x1a, 0x66, 0xf5, 0x4b, 0x92,
	0x2c, 0x78, 0x33, 0xff, 0xe3, 0xdc, 0x37, 0xcd, 0x96, 0x74, 0x66, 0x3c, 0x01, 0xdb, 0x0b, 0xb8,
	0xb6, 0xed, 0x05, 0x15, 0x5e, 0x07, 0x5c, 0xc5, 0x5e, 0x07, 0xcb, 0x8f, 0x30, 0x62, 0x47, 0x37,
	0x3b, 0x7c, 0x0b, 0x93, 0xf7, 0x8a, 0xf8, 0x29, 0xef, 0x4e, 0xfe, 0x01, 0x9b, 0x5d, 0x6f, 0x7c,
	0x99, 0x5e, 0xb7, 0x97, 0x40, 0x7b, 0xfd, 0xe0, 0x6b, 0x5f, 0xff, 0x86, 0xaf, 0xff, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x15, 0x4c, 0xfc, 0x42, 0x9f, 0x03, 0x00, 0x00,
}
