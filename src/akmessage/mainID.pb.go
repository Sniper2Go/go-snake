// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: mainID.proto

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

type MSG int32

const (
	MSG_NO                         MSG = 0 // 不可用ID
	MSG_CS_LOGIN                   MSG = 1 //登录
	MSG_SC_LOGIN                   MSG = 2
	MSG_CS_LOGOUT                  MSG = 3 //登出
	MSG_SC_LOGOUT                  MSG = 4 //
	MSG_CS_HEARTBEAT               MSG = 5 //心跳
	MSG_SC_HEARTBEAT               MSG = 6 //
	MSG_CS_SYNC_CLIENT_DEVICE_INFO MSG = 7 // 同步客户端设备信息
	MSG_SC_SYNC_CLIENT_DEVICE_INFO MSG = 8
	MSG_SS_ROUTE                   MSG = 9 //服务器间路由
)

// Enum value maps for MSG.
var (
	MSG_name = map[int32]string{
		0: "NO",
		1: "CS_LOGIN",
		2: "SC_LOGIN",
		3: "CS_LOGOUT",
		4: "SC_LOGOUT",
		5: "CS_HEARTBEAT",
		6: "SC_HEARTBEAT",
		7: "CS_SYNC_CLIENT_DEVICE_INFO",
		8: "SC_SYNC_CLIENT_DEVICE_INFO",
		9: "SS_ROUTE",
	}
	MSG_value = map[string]int32{
		"NO":                         0,
		"CS_LOGIN":                   1,
		"SC_LOGIN":                   2,
		"CS_LOGOUT":                  3,
		"SC_LOGOUT":                  4,
		"CS_HEARTBEAT":               5,
		"SC_HEARTBEAT":               6,
		"CS_SYNC_CLIENT_DEVICE_INFO": 7,
		"SC_SYNC_CLIENT_DEVICE_INFO": 8,
		"SS_ROUTE":                   9,
	}
)

func (x MSG) Enum() *MSG {
	p := new(MSG)
	*p = x
	return p
}

func (x MSG) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MSG) Descriptor() protoreflect.EnumDescriptor {
	return file_mainID_proto_enumTypes[0].Descriptor()
}

func (MSG) Type() protoreflect.EnumType {
	return &file_mainID_proto_enumTypes[0]
}

func (x MSG) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MSG.Descriptor instead.
func (MSG) EnumDescriptor() ([]byte, []int) {
	return file_mainID_proto_rawDescGZIP(), []int{0}
}

var File_mainID_proto protoreflect.FileDescriptor

var file_mainID_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xb9,
	0x01, 0x0a, 0x03, 0x4d, 0x53, 0x47, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x4f, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x43, 0x53, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x53, 0x43, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x53,
	0x5f, 0x4c, 0x4f, 0x47, 0x4f, 0x55, 0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x43, 0x5f,
	0x4c, 0x4f, 0x47, 0x4f, 0x55, 0x54, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x53, 0x5f, 0x48,
	0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41, 0x54, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x43,
	0x5f, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41, 0x54, 0x10, 0x06, 0x12, 0x1e, 0x0a, 0x1a,
	0x43, 0x53, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x44,
	0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x07, 0x12, 0x1e, 0x0a, 0x1a,
	0x53, 0x43, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x44,
	0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08,
	0x53, 0x53, 0x5f, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x10, 0x09, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2e,
	0x2f, 0x61, 0x6b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_mainID_proto_rawDescOnce sync.Once
	file_mainID_proto_rawDescData = file_mainID_proto_rawDesc
)

func file_mainID_proto_rawDescGZIP() []byte {
	file_mainID_proto_rawDescOnce.Do(func() {
		file_mainID_proto_rawDescData = protoimpl.X.CompressGZIP(file_mainID_proto_rawDescData)
	})
	return file_mainID_proto_rawDescData
}

var file_mainID_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mainID_proto_goTypes = []interface{}{
	(MSG)(0), // 0: MSG
}
var file_mainID_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mainID_proto_init() }
func file_mainID_proto_init() {
	if File_mainID_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mainID_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mainID_proto_goTypes,
		DependencyIndexes: file_mainID_proto_depIdxs,
		EnumInfos:         file_mainID_proto_enumTypes,
	}.Build()
	File_mainID_proto = out.File
	file_mainID_proto_rawDesc = nil
	file_mainID_proto_goTypes = nil
	file_mainID_proto_depIdxs = nil
}
