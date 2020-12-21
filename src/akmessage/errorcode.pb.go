// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: errorcode.proto

package akmessage

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

type ErrorCode int32

const (
	ErrorCode_Invaild        ErrorCode = 0 //非法操作错误
	ErrorCode_Success        ErrorCode = 1 //成功
	ErrorCode_AccountExisted ErrorCode = 2 //账号已存在
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "Invaild",
		1: "Success",
		2: "AccountExisted",
	}
	ErrorCode_value = map[string]int32{
		"Invaild":        0,
		"Success":        1,
		"AccountExisted": 2,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_errorcode_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_errorcode_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_errorcode_proto_rawDescGZIP(), []int{0}
}

var File_errorcode_proto protoreflect.FileDescriptor

var file_errorcode_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0x39, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x49, 0x6e, 0x76, 0x61, 0x69, 0x6c, 0x64, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x64, 0x10, 0x02, 0x42, 0x0e, 0x5a, 0x0c,
	0x2e, 0x2e, 0x2f, 0x61, 0x6b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errorcode_proto_rawDescOnce sync.Once
	file_errorcode_proto_rawDescData = file_errorcode_proto_rawDesc
)

func file_errorcode_proto_rawDescGZIP() []byte {
	file_errorcode_proto_rawDescOnce.Do(func() {
		file_errorcode_proto_rawDescData = protoimpl.X.CompressGZIP(file_errorcode_proto_rawDescData)
	})
	return file_errorcode_proto_rawDescData
}

var file_errorcode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errorcode_proto_goTypes = []interface{}{
	(ErrorCode)(0), // 0: ErrorCode
}
var file_errorcode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errorcode_proto_init() }
func file_errorcode_proto_init() {
	if File_errorcode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_errorcode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errorcode_proto_goTypes,
		DependencyIndexes: file_errorcode_proto_depIdxs,
		EnumInfos:         file_errorcode_proto_enumTypes,
	}.Build()
	File_errorcode_proto = out.File
	file_errorcode_proto_rawDesc = nil
	file_errorcode_proto_goTypes = nil
	file_errorcode_proto_depIdxs = nil
}
