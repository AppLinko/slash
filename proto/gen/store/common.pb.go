// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: store/common.proto

package store

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

type RowStatus int32

const (
	RowStatus_ROW_STATUS_UNSPECIFIED RowStatus = 0
	RowStatus_NORMAL                 RowStatus = 1
	RowStatus_ARCHIVED               RowStatus = 2
)

// Enum value maps for RowStatus.
var (
	RowStatus_name = map[int32]string{
		0: "ROW_STATUS_UNSPECIFIED",
		1: "NORMAL",
		2: "ARCHIVED",
	}
	RowStatus_value = map[string]int32{
		"ROW_STATUS_UNSPECIFIED": 0,
		"NORMAL":                 1,
		"ARCHIVED":               2,
	}
)

func (x RowStatus) Enum() *RowStatus {
	p := new(RowStatus)
	*p = x
	return p
}

func (x RowStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RowStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_store_common_proto_enumTypes[0].Descriptor()
}

func (RowStatus) Type() protoreflect.EnumType {
	return &file_store_common_proto_enumTypes[0]
}

func (x RowStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RowStatus.Descriptor instead.
func (RowStatus) EnumDescriptor() ([]byte, []int) {
	return file_store_common_proto_rawDescGZIP(), []int{0}
}

type Visibility int32

const (
	Visibility_VISIBILITY_UNSPECIFIED Visibility = 0
	Visibility_PRIVATE                Visibility = 1
	Visibility_WORKSPACE              Visibility = 2
	Visibility_PUBLIC                 Visibility = 3
)

// Enum value maps for Visibility.
var (
	Visibility_name = map[int32]string{
		0: "VISIBILITY_UNSPECIFIED",
		1: "PRIVATE",
		2: "WORKSPACE",
		3: "PUBLIC",
	}
	Visibility_value = map[string]int32{
		"VISIBILITY_UNSPECIFIED": 0,
		"PRIVATE":                1,
		"WORKSPACE":              2,
		"PUBLIC":                 3,
	}
)

func (x Visibility) Enum() *Visibility {
	p := new(Visibility)
	*p = x
	return p
}

func (x Visibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Visibility) Descriptor() protoreflect.EnumDescriptor {
	return file_store_common_proto_enumTypes[1].Descriptor()
}

func (Visibility) Type() protoreflect.EnumType {
	return &file_store_common_proto_enumTypes[1]
}

func (x Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Visibility.Descriptor instead.
func (Visibility) EnumDescriptor() ([]byte, []int) {
	return file_store_common_proto_rawDescGZIP(), []int{1}
}

var File_store_common_proto protoreflect.FileDescriptor

var file_store_common_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2a, 0x41, 0x0a, 0x09, 0x52, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a,
	0x0a, 0x16, 0x52, 0x4f, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f,
	0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x52, 0x43, 0x48, 0x49, 0x56,
	0x45, 0x44, 0x10, 0x02, 0x2a, 0x50, 0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x1a, 0x0a, 0x16, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x57,
	0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55,
	0x42, 0x4c, 0x49, 0x43, 0x10, 0x03, 0x42, 0x9c, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x73,
	0x6c, 0x61, 0x73, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x6c, 0x66, 0x68, 0x6f,
	0x73, 0x74, 0x65, 0x64, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0xa2, 0x02, 0x03, 0x53, 0x53, 0x58,
	0xaa, 0x02, 0x0b, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0xca, 0x02,
	0x0b, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0xe2, 0x02, 0x17, 0x53,
	0x6c, 0x61, 0x73, 0x68, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x3a, 0x3a,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_common_proto_rawDescOnce sync.Once
	file_store_common_proto_rawDescData = file_store_common_proto_rawDesc
)

func file_store_common_proto_rawDescGZIP() []byte {
	file_store_common_proto_rawDescOnce.Do(func() {
		file_store_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_common_proto_rawDescData)
	})
	return file_store_common_proto_rawDescData
}

var file_store_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_store_common_proto_goTypes = []any{
	(RowStatus)(0),  // 0: slash.store.RowStatus
	(Visibility)(0), // 1: slash.store.Visibility
}
var file_store_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_store_common_proto_init() }
func file_store_common_proto_init() {
	if File_store_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_common_proto_goTypes,
		DependencyIndexes: file_store_common_proto_depIdxs,
		EnumInfos:         file_store_common_proto_enumTypes,
	}.Build()
	File_store_common_proto = out.File
	file_store_common_proto_rawDesc = nil
	file_store_common_proto_goTypes = nil
	file_store_common_proto_depIdxs = nil
}
