// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: common/response/response.proto

package response

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

type ResMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ResMsg) Reset() {
	*x = ResMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_response_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResMsg) ProtoMessage() {}

func (x *ResMsg) ProtoReflect() protoreflect.Message {
	mi := &file_common_response_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResMsg.ProtoReflect.Descriptor instead.
func (*ResMsg) Descriptor() ([]byte, []int) {
	return file_common_response_response_proto_rawDescGZIP(), []int{0}
}

func (x *ResMsg) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Total     int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	TotalPage int32 `protobuf:"varint,4,opt,name=total_page,json=totalPage,proto3" json:"total_page,omitempty"`
}

func (x *PageInfo) Reset() {
	*x = PageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_response_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfo) ProtoMessage() {}

func (x *PageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_response_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfo.ProtoReflect.Descriptor instead.
func (*PageInfo) Descriptor() ([]byte, []int) {
	return file_common_response_response_proto_rawDescGZIP(), []int{1}
}

func (x *PageInfo) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageInfo) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageInfo) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PageInfo) GetTotalPage() int32 {
	if x != nil {
		return x.TotalPage
	}
	return 0
}

var File_common_response_response_proto protoreflect.FileDescriptor

var file_common_response_response_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x0a, 0x06, 0x52, 0x65,
	0x73, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x70, 0x0a, 0x08, 0x50, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x42, 0x2b, 0x5a, 0x29,
	0x61, 0x64, 0x2d, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x3b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_common_response_response_proto_rawDescOnce sync.Once
	file_common_response_response_proto_rawDescData = file_common_response_response_proto_rawDesc
)

func file_common_response_response_proto_rawDescGZIP() []byte {
	file_common_response_response_proto_rawDescOnce.Do(func() {
		file_common_response_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_response_response_proto_rawDescData)
	})
	return file_common_response_response_proto_rawDescData
}

var file_common_response_response_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_response_response_proto_goTypes = []any{
	(*ResMsg)(nil),   // 0: response.ResMsg
	(*PageInfo)(nil), // 1: response.PageInfo
}
var file_common_response_response_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_response_response_proto_init() }
func file_common_response_response_proto_init() {
	if File_common_response_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_response_response_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ResMsg); i {
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
		file_common_response_response_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PageInfo); i {
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
			RawDescriptor: file_common_response_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_response_response_proto_goTypes,
		DependencyIndexes: file_common_response_response_proto_depIdxs,
		MessageInfos:      file_common_response_response_proto_msgTypes,
	}.Build()
	File_common_response_response_proto = out.File
	file_common_response_response_proto_rawDesc = nil
	file_common_response_response_proto_goTypes = nil
	file_common_response_response_proto_depIdxs = nil
}
