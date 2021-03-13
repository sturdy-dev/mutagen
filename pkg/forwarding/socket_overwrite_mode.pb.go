// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: forwarding/socket_overwrite_mode.proto

package forwarding

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

// SocketOverwriteMode specifies the behavior for overwriting (removing)
// existing Unix domain sockets.
type SocketOverwriteMode int32

const (
	// SocketOverwriteMode_SocketOverwriteModeDefault represents an unspecified
	// socket overwrite mode. It should be converted to one of the following
	// values based on the desired default behavior.
	SocketOverwriteMode_SocketOverwriteModeDefault SocketOverwriteMode = 0
	// SocketOverwriteMode_SocketOverwriteModeLeave specifies that existing
	// sockets should not be overwritten when creating a Unix domain socket
	// listener.
	SocketOverwriteMode_SocketOverwriteModeLeave SocketOverwriteMode = 1
	// SocketOverwriteMode_SocketOverwriteModeOverwrite specifies that existing
	// sockets should be overwritten when creating a Unix domain socket
	// listener.
	SocketOverwriteMode_SocketOverwriteModeOverwrite SocketOverwriteMode = 2
)

// Enum value maps for SocketOverwriteMode.
var (
	SocketOverwriteMode_name = map[int32]string{
		0: "SocketOverwriteModeDefault",
		1: "SocketOverwriteModeLeave",
		2: "SocketOverwriteModeOverwrite",
	}
	SocketOverwriteMode_value = map[string]int32{
		"SocketOverwriteModeDefault":   0,
		"SocketOverwriteModeLeave":     1,
		"SocketOverwriteModeOverwrite": 2,
	}
)

func (x SocketOverwriteMode) Enum() *SocketOverwriteMode {
	p := new(SocketOverwriteMode)
	*p = x
	return p
}

func (x SocketOverwriteMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SocketOverwriteMode) Descriptor() protoreflect.EnumDescriptor {
	return file_forwarding_socket_overwrite_mode_proto_enumTypes[0].Descriptor()
}

func (SocketOverwriteMode) Type() protoreflect.EnumType {
	return &file_forwarding_socket_overwrite_mode_proto_enumTypes[0]
}

func (x SocketOverwriteMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SocketOverwriteMode.Descriptor instead.
func (SocketOverwriteMode) EnumDescriptor() ([]byte, []int) {
	return file_forwarding_socket_overwrite_mode_proto_rawDescGZIP(), []int{0}
}

var File_forwarding_socket_overwrite_mode_proto protoreflect.FileDescriptor

var file_forwarding_socket_overwrite_mode_proto_rawDesc = []byte{
	0x0a, 0x26, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x2a, 0x75, 0x0a, 0x13, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x4f, 0x76,
	0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x53,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x53,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x10, 0x02, 0x42, 0x2e, 0x5a, 0x2c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65,
	0x6e, 0x2d, 0x69, 0x6f, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_forwarding_socket_overwrite_mode_proto_rawDescOnce sync.Once
	file_forwarding_socket_overwrite_mode_proto_rawDescData = file_forwarding_socket_overwrite_mode_proto_rawDesc
)

func file_forwarding_socket_overwrite_mode_proto_rawDescGZIP() []byte {
	file_forwarding_socket_overwrite_mode_proto_rawDescOnce.Do(func() {
		file_forwarding_socket_overwrite_mode_proto_rawDescData = protoimpl.X.CompressGZIP(file_forwarding_socket_overwrite_mode_proto_rawDescData)
	})
	return file_forwarding_socket_overwrite_mode_proto_rawDescData
}

var file_forwarding_socket_overwrite_mode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_forwarding_socket_overwrite_mode_proto_goTypes = []interface{}{
	(SocketOverwriteMode)(0), // 0: forwarding.SocketOverwriteMode
}
var file_forwarding_socket_overwrite_mode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_forwarding_socket_overwrite_mode_proto_init() }
func file_forwarding_socket_overwrite_mode_proto_init() {
	if File_forwarding_socket_overwrite_mode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_forwarding_socket_overwrite_mode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_forwarding_socket_overwrite_mode_proto_goTypes,
		DependencyIndexes: file_forwarding_socket_overwrite_mode_proto_depIdxs,
		EnumInfos:         file_forwarding_socket_overwrite_mode_proto_enumTypes,
	}.Build()
	File_forwarding_socket_overwrite_mode_proto = out.File
	file_forwarding_socket_overwrite_mode_proto_rawDesc = nil
	file_forwarding_socket_overwrite_mode_proto_goTypes = nil
	file_forwarding_socket_overwrite_mode_proto_depIdxs = nil
}
