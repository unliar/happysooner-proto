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
	WritingErrors_None                  WritingErrors = 0
	WritingErrors_ArticleNotFound       WritingErrors = 3000001
	WritingErrors_PostArticleFailed     WritingErrors = 3000002
	WritingErrors_PostCategoryFailed    WritingErrors = 3000003
	WritingErrors_PutArticleFailed      WritingErrors = 3000004
	WritingErrors_OperationNotPermitted WritingErrors = 3000006
)

var WritingErrors_name = map[int32]string{
	0:       "None",
	3000001: "ArticleNotFound",
	3000002: "PostArticleFailed",
	3000003: "PostCategoryFailed",
	3000004: "PutArticleFailed",
	3000006: "OperationNotPermitted",
}

var WritingErrors_value = map[string]int32{
	"None":                  0,
	"ArticleNotFound":       3000001,
	"PostArticleFailed":     3000002,
	"PostCategoryFailed":    3000003,
	"PutArticleFailed":      3000004,
	"OperationNotPermitted": 3000006,
}

func (x WritingErrors) String() string {
	return proto.EnumName(WritingErrors_name, int32(x))
}

func (WritingErrors) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_db9b9589ec5f9fbc, []int{0}
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_db9b9589ec5f9fbc, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

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
	return fileDescriptor_db9b9589ec5f9fbc, []int{1}
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
	return fileDescriptor_db9b9589ec5f9fbc, []int{2}
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
	return fileDescriptor_db9b9589ec5f9fbc, []int{6}
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
	return fileDescriptor_db9b9589ec5f9fbc, []int{7}
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

type PutArticleRequest struct {
	PostID               int64    `protobuf:"varint,1,opt,name=PostID,proto3" json:"PostID,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
	AuthorUID            int64    `protobuf:"varint,4,opt,name=AuthorUID,proto3" json:"AuthorUID,omitempty"`
	Summary              string   `protobuf:"bytes,5,opt,name=Summary,proto3" json:"Summary,omitempty"`
	CategoryID           int64    `protobuf:"varint,6,opt,name=CategoryID,proto3" json:"CategoryID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutArticleRequest) Reset()         { *m = PutArticleRequest{} }
func (m *PutArticleRequest) String() string { return proto.CompactTextString(m) }
func (*PutArticleRequest) ProtoMessage()    {}
func (*PutArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db9b9589ec5f9fbc, []int{8}
}

func (m *PutArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutArticleRequest.Unmarshal(m, b)
}
func (m *PutArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutArticleRequest.Marshal(b, m, deterministic)
}
func (m *PutArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutArticleRequest.Merge(m, src)
}
func (m *PutArticleRequest) XXX_Size() int {
	return xxx_messageInfo_PutArticleRequest.Size(m)
}
func (m *PutArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutArticleRequest proto.InternalMessageInfo

func (m *PutArticleRequest) GetPostID() int64 {
	if m != nil {
		return m.PostID
	}
	return 0
}

func (m *PutArticleRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PutArticleRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *PutArticleRequest) GetAuthorUID() int64 {
	if m != nil {
		return m.AuthorUID
	}
	return 0
}

func (m *PutArticleRequest) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *PutArticleRequest) GetCategoryID() int64 {
	if m != nil {
		return m.CategoryID
	}
	return 0
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
	return fileDescriptor_db9b9589ec5f9fbc, []int{9}
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
	return fileDescriptor_db9b9589ec5f9fbc, []int{10}
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
	return fileDescriptor_db9b9589ec5f9fbc, []int{11}
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

type CategoriesResponse struct {
	List                 []*Category `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CategoriesResponse) Reset()         { *m = CategoriesResponse{} }
func (m *CategoriesResponse) String() string { return proto.CompactTextString(m) }
func (*CategoriesResponse) ProtoMessage()    {}
func (*CategoriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_db9b9589ec5f9fbc, []int{12}
}

func (m *CategoriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoriesResponse.Unmarshal(m, b)
}
func (m *CategoriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoriesResponse.Marshal(b, m, deterministic)
}
func (m *CategoriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoriesResponse.Merge(m, src)
}
func (m *CategoriesResponse) XXX_Size() int {
	return xxx_messageInfo_CategoriesResponse.Size(m)
}
func (m *CategoriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CategoriesResponse proto.InternalMessageInfo

func (m *CategoriesResponse) GetList() []*Category {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterEnum("writing.WritingErrors", WritingErrors_name, WritingErrors_value)
	proto.RegisterType((*Empty)(nil), "writing.Empty")
	proto.RegisterType((*ArticleInfo)(nil), "writing.ArticleInfo")
	proto.RegisterType((*ArticleTag)(nil), "writing.ArticleTag")
	proto.RegisterType((*ArticleCountInfo)(nil), "writing.ArticleCountInfo")
	proto.RegisterType((*ArticleAuthorInfo)(nil), "writing.ArticleAuthorInfo")
	proto.RegisterType((*Category)(nil), "writing.Category")
	proto.RegisterType((*IntRequest)(nil), "writing.IntRequest")
	proto.RegisterType((*PostArticleRequest)(nil), "writing.PostArticleRequest")
	proto.RegisterType((*PutArticleRequest)(nil), "writing.PutArticleRequest")
	proto.RegisterType((*GetArticleListRequest)(nil), "writing.GetArticleListRequest")
	proto.RegisterType((*ErrorResponse)(nil), "writing.ErrorResponse")
	proto.RegisterType((*ArticleListResponse)(nil), "writing.ArticleListResponse")
	proto.RegisterType((*CategoriesResponse)(nil), "writing.CategoriesResponse")
}

func init() { proto.RegisterFile("writing/writing.proto", fileDescriptor_db9b9589ec5f9fbc) }

var fileDescriptor_db9b9589ec5f9fbc = []byte{
	// 863 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcd, 0x6e, 0xeb, 0x44,
	0x14, 0xae, 0xe3, 0xfc, 0x9e, 0xdc, 0xe6, 0x26, 0x73, 0x6f, 0x83, 0x6f, 0x1a, 0x55, 0xd1, 0x48,
	0x40, 0x84, 0x44, 0x41, 0x81, 0x05, 0xbf, 0x12, 0x69, 0x9a, 0x96, 0x48, 0x95, 0x1b, 0x39, 0xa1,
	0xac, 0x4d, 0x3d, 0x84, 0x51, 0x63, 0x3b, 0xd8, 0x63, 0xaa, 0x20, 0xb1, 0x65, 0x57, 0x89, 0x07,
	0x80, 0x35, 0x0f, 0xc0, 0xa2, 0x6b, 0x7e, 0xc4, 0x1b, 0xf0, 0x3e, 0x68, 0xc6, 0x63, 0x8f, 0x63,
	0xa7, 0x1b, 0x16, 0xac, 0x32, 0xdf, 0x39, 0x67, 0xce, 0xef, 0x77, 0x26, 0x86, 0xa3, 0xfb, 0x80,
	0x32, 0xea, 0xad, 0xde, 0x91, 0xbf, 0xa7, 0x9b, 0xc0, 0x67, 0x3e, 0xaa, 0x49, 0x88, 0x6b, 0x50,
	0x99, 0xba, 0x1b, 0xb6, 0xc5, 0x3f, 0xea, 0xd0, 0x1c, 0x07, 0x8c, 0xde, 0xae, 0xc9, 0xcc, 0xfb,
	0xda, 0x47, 0x2d, 0x28, 0xcd, 0x1c, 0x43, 0x1b, 0x68, 0x43, 0xdd, 0x2a, 0xcd, 0x1c, 0xf4, 0x12,
	0x2a, 0x4b, 0xca, 0xd6, 0xc4, 0x28, 0x0d, 0xb4, 0x61, 0xc3, 0x8a, 0x01, 0x32, 0xa0, 0xb6, 0x88,
	0x5c, 0xd7, 0x0e, 0xb6, 0x86, 0x2e, 0xe4, 0x09, 0x44, 0x6f, 0x42, 0x79, 0x69, 0xaf, 0x42, 0xa3,
	0x3c, 0xd0, 0x87, 0xcd, 0xd1, 0x8b, 0xd3, 0x24, 0xbe, 0x8c, 0xb1, 0xb4, 0x57, 0x96, 0x30, 0x40,
	0x7d, 0x68, 0x4c, 0x02, 0x62, 0x33, 0xe2, 0x8c, 0x99, 0x51, 0x11, 0xf1, 0x94, 0x00, 0x75, 0xa1,
	0xba, 0x60, 0x36, 0x8b, 0x42, 0xa3, 0x2a, 0x54, 0x12, 0xf1, 0xc0, 0x13, 0xdf, 0x63, 0xc4, 0x63,
	0x46, 0x2d, 0x0e, 0x2c, 0x21, 0x7a, 0x1b, 0xea, 0x13, 0x9b, 0x91, 0x95, 0x1f, 0x6c, 0x8d, 0xfa,
	0x40, 0x1b, 0x36, 0x47, 0x9d, 0x34, 0x78, 0xa2, 0xb0, 0x52, 0x13, 0x34, 0x85, 0xb6, 0x4c, 0x69,
	0xe2, 0x47, 0x1e, 0xe3, 0xb5, 0x1b, 0x0d, 0x71, 0xed, 0x55, 0x3e, 0xe7, 0xd4, 0xc0, 0x2a, 0x5c,
	0x41, 0x9f, 0x43, 0x47, 0xca, 0xc6, 0x11, 0xfb, 0xc6, 0x0f, 0x84, 0x1f, 0x10, 0x7e, 0x7a, 0x79,
	0x3f, 0xca, 0xc2, 0x2a, 0x5e, 0xc2, 0xa7, 0x00, 0xaa, 0x47, 0x85, 0x31, 0xb4, 0x41, 0x5f, 0xda,
	0x2b, 0x39, 0x04, 0x7e, 0xc4, 0x3f, 0x69, 0xc5, 0x0a, 0xe2, 0xf6, 0xb8, 0x2e, 0x6f, 0x4f, 0x7c,
	0x37, 0x81, 0x08, 0x41, 0xf9, 0x8a, 0xde, 0xc5, 0x63, 0xd4, 0x2d, 0x71, 0x8e, 0xad, 0xd7, 0x6b,
	0x72, 0xcb, 0xc4, 0x14, 0x85, 0xb5, 0x80, 0x5c, 0x73, 0x4e, 0x43, 0x71, 0xa1, 0x1c, 0x6b, 0x24,
	0x44, 0x3d, 0xa8, 0xcf, 0xed, 0x15, 0xb9, 0xa1, 0xe4, 0x5e, 0x4e, 0x2d, 0xc5, 0xf8, 0x87, 0x3d,
	0xcd, 0xe0, 0x73, 0x8e, 0xd1, 0x17, 0xb3, 0x73, 0x99, 0x94, 0x12, 0xa0, 0x37, 0xa0, 0x15, 0x03,
	0x93, 0xde, 0xde, 0x79, 0xb6, 0x9b, 0xf0, 0x2c, 0x27, 0x45, 0x18, 0x9e, 0xc5, 0x92, 0xf1, 0x77,
	0x36, 0xb3, 0x03, 0xc9, 0xba, 0x1d, 0x19, 0xfe, 0x48, 0x31, 0xa0, 0xd0, 0xbf, 0x16, 0x94, 0x26,
	0xa6, 0xf4, 0x5d, 0x9a, 0x98, 0x1c, 0x4f, 0x4d, 0xe9, 0xa5, 0x34, 0x35, 0x31, 0x06, 0x98, 0x79,
	0xcc, 0x22, 0xdf, 0x46, 0x24, 0x64, 0x9c, 0xf4, 0x37, 0xf6, 0x3a, 0x22, 0xd2, 0x41, 0x0c, 0xf0,
	0xcf, 0x1a, 0xa0, 0xb9, 0x1f, 0x32, 0x59, 0x63, 0x62, 0x7c, 0x02, 0x90, 0x84, 0x4d, 0x2b, 0xcc,
	0x48, 0x9e, 0xde, 0xa0, 0x84, 0xc8, 0xfa, 0x2e, 0x91, 0x77, 0x1a, 0x56, 0xce, 0x37, 0x2c, 0xb3,
	0x79, 0x95, 0x9d, 0xcd, 0xc3, 0xbf, 0x69, 0xd0, 0x99, 0x47, 0xf9, 0xec, 0xba, 0x50, 0xe5, 0x39,
	0xa7, 0x99, 0x49, 0xf4, 0x7f, 0x65, 0x95, 0xeb, 0x4e, 0x35, 0xdf, 0x1d, 0xfc, 0xab, 0x06, 0x47,
	0x97, 0x24, 0xc9, 0xfa, 0x8a, 0x86, 0xe9, 0x10, 0x86, 0xf0, 0xfc, 0x8a, 0xba, 0x94, 0x8d, 0xf9,
	0x16, 0xf1, 0xe7, 0x29, 0x29, 0x21, 0x2f, 0xe6, 0xdc, 0xe6, 0x1c, 0x4c, 0xb8, 0xcd, 0xcf, 0x5c,
	0xb6, 0xa0, 0xdf, 0x13, 0x49, 0x6c, 0x71, 0xce, 0xe5, 0x52, 0x2e, 0x4c, 0xaa, 0x0f, 0x8d, 0xeb,
	0xc0, 0x21, 0xc1, 0x72, 0xbb, 0x21, 0xb2, 0x0e, 0x25, 0xc0, 0x53, 0x38, 0x9c, 0x06, 0x81, 0x1f,
	0x58, 0x24, 0xdc, 0xf8, 0x5e, 0x48, 0xd0, 0xfb, 0xd0, 0x10, 0x02, 0xb1, 0xf3, 0x3c, 0xb5, 0xd6,
	0xa8, 0x9b, 0xee, 0xfc, 0x97, 0xf1, 0xaf, 0x30, 0x08, 0x2d, 0x65, 0x88, 0x2f, 0xe1, 0xc5, 0x4e,
	0xb1, 0xd2, 0xd9, 0xbb, 0x50, 0x97, 0xe2, 0xd0, 0xd0, 0xc4, 0xdb, 0xf9, 0x32, 0xff, 0x7e, 0x88,
	0x97, 0x23, 0xb5, 0xc2, 0x1f, 0x03, 0x92, 0xb9, 0x53, 0x12, 0xa6, 0x7e, 0x5e, 0xe7, 0x7b, 0x1e,
	0x32, 0xe9, 0x63, 0xcf, 0x13, 0x28, 0xd4, 0x6f, 0xfd, 0xa2, 0xc1, 0xe1, 0x4e, 0x8a, 0xa8, 0x0e,
	0x65, 0xd3, 0xf7, 0x48, 0xfb, 0x00, 0x75, 0xe1, 0xb9, 0x0c, 0x62, 0xfa, 0xec, 0xc2, 0x8f, 0x3c,
	0xa7, 0xfd, 0xfb, 0xc3, 0xa3, 0x86, 0x0c, 0xe8, 0x64, 0xe8, 0x7f, 0x61, 0xd3, 0x35, 0x71, 0xda,
	0x7f, 0x70, 0xcd, 0xab, 0x78, 0x31, 0x92, 0x18, 0x52, 0xf5, 0x27, 0x57, 0xbd, 0x06, 0x6d, 0x45,
	0x4a, 0xa9, 0xf8, 0x8b, 0x2b, 0xfa, 0x70, 0x74, 0xbd, 0x21, 0x81, 0xcd, 0xa8, 0xef, 0x99, 0x3e,
	0x9b, 0x93, 0xc0, 0xa5, 0x8c, 0x11, 0xa7, 0xfd, 0xf7, 0xc3, 0xa3, 0x36, 0xfa, 0x47, 0x87, 0x86,
	0x4c, 0x7d, 0x71, 0x83, 0x3e, 0x85, 0x96, 0xe2, 0xc8, 0xd9, 0x76, 0xe6, 0x20, 0xf5, 0xc7, 0xa2,
	0xd6, 0xb6, 0xb7, 0xb7, 0x63, 0xf8, 0x00, 0xcd, 0xb3, 0xd7, 0x79, 0xf9, 0xe8, 0x24, 0xb5, 0xdc,
	0xcb, 0xbd, 0x5e, 0x3f, 0xef, 0x29, 0x3b, 0x2b, 0x7c, 0x80, 0xce, 0xa0, 0x99, 0x69, 0x05, 0x3a,
	0x4e, 0xcd, 0x8b, 0xef, 0x43, 0x4f, 0x71, 0x62, 0x97, 0x3e, 0x9f, 0x01, 0xa8, 0xce, 0x20, 0xf5,
	0x6f, 0x51, 0xd8, 0xe1, 0x27, 0x3d, 0x7c, 0x08, 0xcf, 0xb2, 0x6d, 0x47, 0xc5, 0x69, 0x3f, 0x79,
	0xf5, 0x13, 0x38, 0xbc, 0x24, 0x4c, 0xf1, 0x07, 0xb5, 0x94, 0x21, 0xff, 0x2e, 0xe8, 0x1d, 0xe7,
	0x7d, 0x65, 0x49, 0xf6, 0x01, 0x34, 0xe7, 0xd1, 0x7f, 0x89, 0xfb, 0x55, 0x55, 0x7c, 0x87, 0xbc,
	0xf7, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x3d, 0x52, 0x66, 0xa0, 0x08, 0x00, 0x00,
}
