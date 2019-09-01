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

// --------------------------response  start-----------------------------
type WritingErrors int32

const (
	WritingErrors_None               WritingErrors = 0
	WritingErrors_ArticleNotFound    WritingErrors = 3000001
	WritingErrors_PostArticleFailed  WritingErrors = 3000002
	WritingErrors_PostCategoryFailed WritingErrors = 3000003
)

var WritingErrors_name = map[int32]string{
	0:       "None",
	3000001: "ArticleNotFound",
	3000002: "PostArticleFailed",
	3000003: "PostCategoryFailed",
}

var WritingErrors_value = map[string]int32{
	"None":               0,
	"ArticleNotFound":    3000001,
	"PostArticleFailed":  3000002,
	"PostCategoryFailed": 3000003,
}

func (x WritingErrors) String() string {
	return proto.EnumName(WritingErrors_name, int32(x))
}

func (WritingErrors) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_db9b9589ec5f9fbc, []int{0}
}

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
	return fileDescriptor_db9b9589ec5f9fbc, []int{2}
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
	return fileDescriptor_db9b9589ec5f9fbc, []int{3}
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
	return fileDescriptor_db9b9589ec5f9fbc, []int{4}
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

// --------------------------request  start------------------------------
// int参数请求
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
	return fileDescriptor_db9b9589ec5f9fbc, []int{5}
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

type PostArticleRequest struct {
	CategoryID           int64    `protobuf:"varint,1,opt,name=CategoryID,proto3" json:"CategoryID,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
	AuthorUID            int64    `protobuf:"varint,4,opt,name=AuthorUID,proto3" json:"AuthorUID,omitempty"`
	Summary              string   `protobuf:"bytes,5,opt,name=Summary,proto3" json:"Summary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostArticleRequest) Reset()         { *m = PostArticleRequest{} }
func (m *PostArticleRequest) String() string { return proto.CompactTextString(m) }
func (*PostArticleRequest) ProtoMessage()    {}
func (*PostArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db9b9589ec5f9fbc, []int{6}
}

func (m *PostArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostArticleRequest.Unmarshal(m, b)
}
func (m *PostArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostArticleRequest.Marshal(b, m, deterministic)
}
func (m *PostArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostArticleRequest.Merge(m, src)
}
func (m *PostArticleRequest) XXX_Size() int {
	return xxx_messageInfo_PostArticleRequest.Size(m)
}
func (m *PostArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostArticleRequest proto.InternalMessageInfo

func (m *PostArticleRequest) GetCategoryID() int64 {
	if m != nil {
		return m.CategoryID
	}
	return 0
}

func (m *PostArticleRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PostArticleRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *PostArticleRequest) GetAuthorUID() int64 {
	if m != nil {
		return m.AuthorUID
	}
	return 0
}

func (m *PostArticleRequest) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

type GetArticleListRequest struct {
	LimitAriticleID      int64    `protobuf:"varint,1,opt,name=LimitAriticleID,proto3" json:"LimitAriticleID,omitempty"`
	Page                 int64    `protobuf:"varint,2,opt,name=Page,proto3" json:"Page,omitempty"`
	Size                 int64    `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
	CategoryID           int64    `protobuf:"varint,4,opt,name=CategoryID,proto3" json:"CategoryID,omitempty"`
	OrderType            string   `protobuf:"bytes,5,opt,name=OrderType,proto3" json:"OrderType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetArticleListRequest) Reset()         { *m = GetArticleListRequest{} }
func (m *GetArticleListRequest) String() string { return proto.CompactTextString(m) }
func (*GetArticleListRequest) ProtoMessage()    {}
func (*GetArticleListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db9b9589ec5f9fbc, []int{7}
}

func (m *GetArticleListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArticleListRequest.Unmarshal(m, b)
}
func (m *GetArticleListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArticleListRequest.Marshal(b, m, deterministic)
}
func (m *GetArticleListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArticleListRequest.Merge(m, src)
}
func (m *GetArticleListRequest) XXX_Size() int {
	return xxx_messageInfo_GetArticleListRequest.Size(m)
}
func (m *GetArticleListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArticleListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArticleListRequest proto.InternalMessageInfo

func (m *GetArticleListRequest) GetLimitAriticleID() int64 {
	if m != nil {
		return m.LimitAriticleID
	}
	return 0
}

func (m *GetArticleListRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetArticleListRequest) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *GetArticleListRequest) GetCategoryID() int64 {
	if m != nil {
		return m.CategoryID
	}
	return 0
}

func (m *GetArticleListRequest) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

type ErrorResponse struct {
	ErrorInfo            WritingErrors `protobuf:"varint,1,opt,name=ErrorInfo,proto3,enum=writing.WritingErrors" json:"ErrorInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ErrorResponse) Reset()         { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()    {}
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_db9b9589ec5f9fbc, []int{8}
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

func (m *ErrorResponse) GetErrorInfo() WritingErrors {
	if m != nil {
		return m.ErrorInfo
	}
	return WritingErrors_None
}

type ArticleListResponse struct {
	Articles             []*ArticleInfo `protobuf:"bytes,1,rep,name=Articles,proto3" json:"Articles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ArticleListResponse) Reset()         { *m = ArticleListResponse{} }
func (m *ArticleListResponse) String() string { return proto.CompactTextString(m) }
func (*ArticleListResponse) ProtoMessage()    {}
func (*ArticleListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_db9b9589ec5f9fbc, []int{9}
}

func (m *ArticleListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleListResponse.Unmarshal(m, b)
}
func (m *ArticleListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleListResponse.Marshal(b, m, deterministic)
}
func (m *ArticleListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleListResponse.Merge(m, src)
}
func (m *ArticleListResponse) XXX_Size() int {
	return xxx_messageInfo_ArticleListResponse.Size(m)
}
func (m *ArticleListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleListResponse proto.InternalMessageInfo

func (m *ArticleListResponse) GetArticles() []*ArticleInfo {
	if m != nil {
		return m.Articles
	}
	return nil
}

func init() {
	proto.RegisterEnum("writing.WritingErrors", WritingErrors_name, WritingErrors_value)
	proto.RegisterType((*ArticleInfo)(nil), "writing.ArticleInfo")
	proto.RegisterType((*ArticleTag)(nil), "writing.ArticleTag")
	proto.RegisterType((*ArticleCountInfo)(nil), "writing.ArticleCountInfo")
	proto.RegisterType((*ArticleAuthorInfo)(nil), "writing.ArticleAuthorInfo")
	proto.RegisterType((*Category)(nil), "writing.Category")
	proto.RegisterType((*IntRequest)(nil), "writing.IntRequest")
	proto.RegisterType((*PostArticleRequest)(nil), "writing.PostArticleRequest")
	proto.RegisterType((*GetArticleListRequest)(nil), "writing.GetArticleListRequest")
	proto.RegisterType((*ErrorResponse)(nil), "writing.ErrorResponse")
	proto.RegisterType((*ArticleListResponse)(nil), "writing.ArticleListResponse")
}

func init() { proto.RegisterFile("writing/writing.proto", fileDescriptor_db9b9589ec5f9fbc) }

var fileDescriptor_db9b9589ec5f9fbc = []byte{
	// 730 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0x4b, 0x4f, 0xdb, 0x4a,
	0x14, 0xc6, 0x0f, 0x20, 0x39, 0x81, 0x10, 0x0e, 0x0f, 0x99, 0x5c, 0x84, 0xa2, 0x59, 0xdc, 0x1b,
	0x5d, 0xe9, 0x72, 0xab, 0xb4, 0x9b, 0x56, 0xea, 0x22, 0x84, 0x40, 0x23, 0x21, 0x17, 0x39, 0x29,
	0x5d, 0xbb, 0x78, 0x48, 0x2d, 0x12, 0x9b, 0xda, 0xe3, 0xa2, 0x54, 0xea, 0xb6, 0x3b, 0xa4, 0xee,
	0xdb, 0x7d, 0x7f, 0x02, 0xeb, 0xb6, 0x7f, 0xac, 0x9a, 0xf1, 0xf8, 0x0d, 0xab, 0xcc, 0xf9, 0xe6,
	0x9b, 0xf3, 0xfc, 0x8e, 0x03, 0x3b, 0xb7, 0x81, 0xcb, 0x5c, 0x6f, 0xfa, 0xbf, 0xfc, 0x3d, 0xbc,
	0x09, 0x7c, 0xe6, 0xe3, 0xaa, 0x34, 0xc9, 0x17, 0x0d, 0x1a, 0xfd, 0x80, 0xb9, 0x97, 0x33, 0x3a,
	0xf2, 0xae, 0x7c, 0x6c, 0x82, 0x3a, 0x72, 0x0c, 0xa5, 0xa3, 0x74, 0x35, 0x4b, 0x1d, 0x39, 0xb8,
	0x0d, 0xcb, 0x13, 0x97, 0xcd, 0xa8, 0xa1, 0x76, 0x94, 0x6e, 0xdd, 0x8a, 0x0d, 0x34, 0x60, 0x75,
	0x1c, 0xcd, 0xe7, 0x76, 0xb0, 0x30, 0x34, 0x81, 0x27, 0x26, 0xfe, 0x03, 0xfa, 0xc4, 0x9e, 0x86,
	0x86, 0xde, 0xd1, 0xba, 0x8d, 0xde, 0xd6, 0x61, 0x12, 0x56, 0xc6, 0x98, 0xd8, 0x53, 0x4b, 0x10,
	0x70, 0x1f, 0xea, 0x83, 0x80, 0xda, 0x8c, 0x3a, 0x7d, 0x66, 0x2c, 0x8b, 0x78, 0x19, 0x80, 0xbb,
	0xb0, 0x32, 0x66, 0x36, 0x8b, 0x42, 0x63, 0x45, 0x5c, 0x49, 0x8b, 0x07, 0x1e, 0xf8, 0x1e, 0xa3,
	0x1e, 0x33, 0x56, 0xe3, 0xc0, 0xd2, 0xc4, 0xff, 0xa0, 0x36, 0xb0, 0x19, 0x9d, 0xfa, 0xc1, 0xc2,
	0xa8, 0x75, 0x94, 0x6e, 0xa3, 0xb7, 0x99, 0x06, 0x4f, 0x2e, 0xac, 0x94, 0x82, 0x43, 0x68, 0xc9,
	0x94, 0x06, 0x7e, 0xe4, 0x31, 0x5e, 0xbb, 0x51, 0x17, 0xcf, 0xf6, 0xca, 0x39, 0xa7, 0x04, 0xab,
	0xf2, 0x04, 0x5f, 0xc1, 0xa6, 0xc4, 0xfa, 0x11, 0x7b, 0xef, 0x07, 0xc2, 0x0f, 0x08, 0x3f, 0xed,
	0xb2, 0x9f, 0x8c, 0x61, 0x55, 0x1f, 0x91, 0x43, 0x80, 0xac, 0x47, 0x95, 0x31, 0xb4, 0x40, 0x9b,
	0xd8, 0x53, 0x39, 0x04, 0x7e, 0x24, 0x5f, 0x95, 0x6a, 0x05, 0x71, 0x7b, 0xe6, 0x73, 0xde, 0x9e,
	0xf8, 0x6d, 0x62, 0x22, 0x82, 0x7e, 0xe6, 0x5e, 0xc7, 0x63, 0xd4, 0x2c, 0x71, 0x8e, 0xd9, 0xb3,
	0x19, 0xbd, 0x64, 0x62, 0x8a, 0x82, 0x2d, 0x4c, 0x7e, 0x73, 0xec, 0x86, 0xe2, 0x81, 0x1e, 0xdf,
	0x48, 0x13, 0xdb, 0x50, 0x3b, 0xb7, 0xa7, 0xf4, 0xc2, 0xa5, 0xb7, 0x72, 0x6a, 0xa9, 0x4d, 0x3e,
	0x3f, 0xd0, 0x0c, 0x3e, 0xe7, 0xd8, 0x7a, 0x33, 0x3a, 0x96, 0x49, 0x65, 0x00, 0xfe, 0x0d, 0xcd,
	0xd8, 0x30, 0xdd, 0xcb, 0x6b, 0xcf, 0x9e, 0x27, 0x3a, 0x2b, 0xa1, 0x48, 0x60, 0x2d, 0x46, 0xfa,
	0x1f, 0x6d, 0x66, 0x07, 0x52, 0x75, 0x05, 0x8c, 0xbc, 0xc8, 0x14, 0x50, 0xe9, 0x5f, 0x13, 0xd4,
	0x81, 0x29, 0x7d, 0xab, 0x03, 0x93, 0xdb, 0x43, 0x53, 0x7a, 0x51, 0x87, 0x26, 0x21, 0x00, 0x23,
	0x8f, 0x59, 0xf4, 0x43, 0x44, 0x43, 0xc6, 0x45, 0x7f, 0x61, 0xcf, 0x22, 0x2a, 0x1d, 0xc4, 0x06,
	0xf9, 0xae, 0x00, 0x9e, 0xfb, 0x21, 0x93, 0x35, 0x26, 0xe4, 0x03, 0x80, 0x24, 0x6c, 0x5a, 0x61,
	0x0e, 0x79, 0x7c, 0x83, 0x12, 0x21, 0x6b, 0x45, 0x21, 0x17, 0x1a, 0xa6, 0x97, 0x1b, 0x96, 0xdb,
	0xbc, 0xe5, 0xc2, 0xe6, 0x91, 0x1f, 0x0a, 0xec, 0x9c, 0xd2, 0x24, 0xbb, 0x33, 0x37, 0x4c, 0xcb,
	0xe9, 0xc2, 0xc6, 0x99, 0x3b, 0x77, 0x59, 0x9f, 0xeb, 0x91, 0x2f, 0x7a, 0x92, 0x66, 0x19, 0xe6,
	0x2a, 0xe1, 0xd3, 0x4c, 0x54, 0xc2, 0xcf, 0x1c, 0x1b, 0xbb, 0x9f, 0xa8, 0x94, 0x88, 0x38, 0x97,
	0x6a, 0xd6, 0x2b, 0x35, 0xef, 0x43, 0xfd, 0x75, 0xe0, 0xd0, 0x60, 0xb2, 0xb8, 0xa1, 0x32, 0xcf,
	0x0c, 0x20, 0x43, 0x58, 0x1f, 0x06, 0x81, 0x1f, 0x58, 0x34, 0xbc, 0xf1, 0xbd, 0x90, 0xe2, 0x33,
	0xa8, 0x0b, 0x40, 0x6c, 0x0f, 0x4f, 0xad, 0xd9, 0xdb, 0x4d, 0xb7, 0xe7, 0x6d, 0xfc, 0x2b, 0x08,
	0xa1, 0x95, 0x11, 0xc9, 0x29, 0x6c, 0x15, 0x8a, 0x95, 0xce, 0x9e, 0x40, 0x4d, 0xc2, 0xa1, 0xa1,
	0x88, 0xaf, 0xd0, 0x76, 0x79, 0x13, 0xc5, 0x0e, 0xa6, 0xac, 0x7f, 0xaf, 0x60, 0xbd, 0x10, 0x04,
	0x6b, 0xa0, 0x9b, 0xbe, 0x47, 0x5b, 0x4b, 0xb8, 0x0b, 0x1b, 0x92, 0x66, 0xfa, 0xec, 0xc4, 0x8f,
	0x3c, 0xa7, 0xf5, 0xf3, 0xee, 0x5e, 0x41, 0x03, 0x36, 0x73, 0x52, 0x38, 0xb1, 0xdd, 0x19, 0x75,
	0x5a, 0xbf, 0xf8, 0xcd, 0x5e, 0x2c, 0x92, 0xa4, 0x19, 0xf2, 0xea, 0xf7, 0xdd, 0xbd, 0xd2, 0xfb,
	0xa6, 0x42, 0x5d, 0x66, 0x32, 0xbe, 0xc0, 0x97, 0xd0, 0xcc, 0xc6, 0x75, 0xb4, 0x18, 0x39, 0x98,
	0x7d, 0x2d, 0x33, 0x2d, 0xb6, 0x1f, 0x4c, 0x9e, 0x2c, 0xe1, 0x79, 0xfe, 0x39, 0x6f, 0x00, 0x1e,
	0xa4, 0xcc, 0x07, 0x65, 0xd0, 0xde, 0x2f, 0x7b, 0xca, 0xb7, 0x8d, 0x2c, 0xe1, 0x11, 0x34, 0x72,
	0x35, 0xe1, 0x5f, 0x29, 0xbd, 0x2a, 0xfa, 0x76, 0x36, 0x9e, 0xe2, 0x24, 0x9f, 0xc3, 0x5a, 0xbe,
	0x7a, 0xac, 0x7e, 0x83, 0x1f, 0x7b, 0xfa, 0x6e, 0x45, 0xfc, 0x33, 0x3d, 0xfd, 0x13, 0x00, 0x00,
	0xff, 0xff, 0xe1, 0xfa, 0x09, 0xd1, 0xb2, 0x06, 0x00, 0x00,
}