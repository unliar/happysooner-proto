// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: file/file.proto

package file

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

// 传图   111.jpg
type PostFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data          string `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`                   // 数据
	FileName      string `protobuf:"bytes,2,opt,name=FileName,proto3" json:"FileName,omitempty"`           // 文件名  111
	FileExtension string `protobuf:"bytes,3,opt,name=FileExtension,proto3" json:"FileExtension,omitempty"` // 文件类型 jpg
}

func (x *PostFileRequest) Reset() {
	*x = PostFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostFileRequest) ProtoMessage() {}

func (x *PostFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostFileRequest.ProtoReflect.Descriptor instead.
func (*PostFileRequest) Descriptor() ([]byte, []int) {
	return file_file_file_proto_rawDescGZIP(), []int{0}
}

func (x *PostFileRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *PostFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *PostFileRequest) GetFileExtension() string {
	if x != nil {
		return x.FileExtension
	}
	return ""
}

type PostFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileURL string `protobuf:"bytes,1,opt,name=FileURL,proto3" json:"FileURL,omitempty"` // 资源地址
}

func (x *PostFileResponse) Reset() {
	*x = PostFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostFileResponse) ProtoMessage() {}

func (x *PostFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostFileResponse.ProtoReflect.Descriptor instead.
func (*PostFileResponse) Descriptor() ([]byte, []int) {
	return file_file_file_proto_rawDescGZIP(), []int{1}
}

func (x *PostFileResponse) GetFileURL() string {
	if x != nil {
		return x.FileURL
	}
	return ""
}

var File_file_file_proto protoreflect.FileDescriptor

var file_file_file_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x67, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a,
	0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x46, 0x69,
	0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x2c, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x52, 0x4c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x52, 0x4c, 0x32, 0x46,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x56, 0x12, 0x3c, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x15, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x6e, 0x6c, 0x69, 0x61, 0x72, 0x2f, 0x68, 0x61, 0x70, 0x70,
	0x79, 0x73, 0x6f, 0x6f, 0x6e, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x32,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_file_proto_rawDescOnce sync.Once
	file_file_file_proto_rawDescData = file_file_file_proto_rawDesc
)

func file_file_file_proto_rawDescGZIP() []byte {
	file_file_file_proto_rawDescOnce.Do(func() {
		file_file_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_file_proto_rawDescData)
	})
	return file_file_file_proto_rawDescData
}

var file_file_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_file_file_proto_goTypes = []interface{}{
	(*PostFileRequest)(nil),  // 0: file.PostFileRequest
	(*PostFileResponse)(nil), // 1: file.PostFileResponse
}
var file_file_file_proto_depIdxs = []int32{
	0, // 0: file.fileSV.PostImage:input_type -> file.PostFileRequest
	1, // 1: file.fileSV.PostImage:output_type -> file.PostFileResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_file_file_proto_init() }
func file_file_file_proto_init() {
	if File_file_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostFileRequest); i {
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
		file_file_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostFileResponse); i {
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
			RawDescriptor: file_file_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_file_proto_goTypes,
		DependencyIndexes: file_file_file_proto_depIdxs,
		MessageInfos:      file_file_file_proto_msgTypes,
	}.Build()
	File_file_file_proto = out.File
	file_file_file_proto_rawDesc = nil
	file_file_file_proto_goTypes = nil
	file_file_file_proto_depIdxs = nil
}
