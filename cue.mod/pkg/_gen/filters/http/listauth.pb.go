// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.21.5
// source: source/filters/http/proto/listauth.proto

package proto

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

type ListAuthConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blacklist string `protobuf:"bytes,1,opt,name=blacklist,proto3" json:"blacklist,omitempty"`
	Whitelist string `protobuf:"bytes,2,opt,name=whitelist,proto3" json:"whitelist,omitempty"`
	Denylist  string `protobuf:"bytes,3,opt,name=denylist,proto3" json:"denylist,omitempty"`
	Allowlist string `protobuf:"bytes,4,opt,name=allowlist,proto3" json:"allowlist,omitempty"`
}

func (x *ListAuthConfig) Reset() {
	*x = ListAuthConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_filters_http_proto_listauth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuthConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthConfig) ProtoMessage() {}

func (x *ListAuthConfig) ProtoReflect() protoreflect.Message {
	mi := &file_source_filters_http_proto_listauth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthConfig.ProtoReflect.Descriptor instead.
func (*ListAuthConfig) Descriptor() ([]byte, []int) {
	return file_source_filters_http_proto_listauth_proto_rawDescGZIP(), []int{0}
}

func (x *ListAuthConfig) GetBlacklist() string {
	if x != nil {
		return x.Blacklist
	}
	return ""
}

func (x *ListAuthConfig) GetWhitelist() string {
	if x != nil {
		return x.Whitelist
	}
	return ""
}

func (x *ListAuthConfig) GetDenylist() string {
	if x != nil {
		return x.Denylist
	}
	return ""
}

func (x *ListAuthConfig) GetAllowlist() string {
	if x != nil {
		return x.Allowlist
	}
	return ""
}

var File_source_filters_http_proto_listauth_proto protoreflect.FileDescriptor

var file_source_filters_http_proto_listauth_proto_rawDesc = []byte{
	0x0a, 0x28, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x67, 0x72, 0x65, 0x79,
	0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6f, 0x2e, 0x67, 0x6d, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x22, 0x86, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x61,
	0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c,
	0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x68, 0x69, 0x74, 0x65,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x68, 0x69, 0x74,
	0x65, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x6e, 0x79, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x6e, 0x79, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x6c, 0x69, 0x73, 0x74, 0x42,
	0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72,
	0x65, 0x79, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6d, 0x2d, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_source_filters_http_proto_listauth_proto_rawDescOnce sync.Once
	file_source_filters_http_proto_listauth_proto_rawDescData = file_source_filters_http_proto_listauth_proto_rawDesc
)

func file_source_filters_http_proto_listauth_proto_rawDescGZIP() []byte {
	file_source_filters_http_proto_listauth_proto_rawDescOnce.Do(func() {
		file_source_filters_http_proto_listauth_proto_rawDescData = protoimpl.X.CompressGZIP(file_source_filters_http_proto_listauth_proto_rawDescData)
	})
	return file_source_filters_http_proto_listauth_proto_rawDescData
}

var file_source_filters_http_proto_listauth_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_source_filters_http_proto_listauth_proto_goTypes = []interface{}{
	(*ListAuthConfig)(nil), // 0: greymatter_io.gm_proxy.source.filters.http.ListAuthConfig
}
var file_source_filters_http_proto_listauth_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_source_filters_http_proto_listauth_proto_init() }
func file_source_filters_http_proto_listauth_proto_init() {
	if File_source_filters_http_proto_listauth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_source_filters_http_proto_listauth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuthConfig); i {
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
			RawDescriptor: file_source_filters_http_proto_listauth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_source_filters_http_proto_listauth_proto_goTypes,
		DependencyIndexes: file_source_filters_http_proto_listauth_proto_depIdxs,
		MessageInfos:      file_source_filters_http_proto_listauth_proto_msgTypes,
	}.Build()
	File_source_filters_http_proto_listauth_proto = out.File
	file_source_filters_http_proto_listauth_proto_rawDesc = nil
	file_source_filters_http_proto_listauth_proto_goTypes = nil
	file_source_filters_http_proto_listauth_proto_depIdxs = nil
}
