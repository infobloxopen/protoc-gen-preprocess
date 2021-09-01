// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: github.com/infobloxopen/protoc-gen-preprocess/options/preprocess.proto

package preprocess

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PreprocessString_Methods int32

const (
	PreprocessString_none  PreprocessString_Methods = 0
	PreprocessString_trim  PreprocessString_Methods = 1
	PreprocessString_upper PreprocessString_Methods = 2
	PreprocessString_lower PreprocessString_Methods = 3
	PreprocessString_clear PreprocessString_Methods = 4
)

// Enum value maps for PreprocessString_Methods.
var (
	PreprocessString_Methods_name = map[int32]string{
		0: "none",
		1: "trim",
		2: "upper",
		3: "lower",
		4: "clear",
	}
	PreprocessString_Methods_value = map[string]int32{
		"none":  0,
		"trim":  1,
		"upper": 2,
		"lower": 3,
		"clear": 4,
	}
)

func (x PreprocessString_Methods) Enum() *PreprocessString_Methods {
	p := new(PreprocessString_Methods)
	*p = x
	return p
}

func (x PreprocessString_Methods) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PreprocessString_Methods) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_enumTypes[0].Descriptor()
}

func (PreprocessString_Methods) Type() protoreflect.EnumType {
	return &file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_enumTypes[0]
}

func (x PreprocessString_Methods) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PreprocessString_Methods.Descriptor instead.
func (PreprocessString_Methods) EnumDescriptor() ([]byte, []int) {
	return file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_rawDescGZIP(), []int{2, 0}
}

type PreprocessMessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	String_ *PreprocessString `protobuf:"bytes,1,opt,name=string,proto3" json:"string,omitempty"`
}

func (x *PreprocessMessageOptions) Reset() {
	*x = PreprocessMessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreprocessMessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreprocessMessageOptions) ProtoMessage() {}

func (x *PreprocessMessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreprocessMessageOptions.ProtoReflect.Descriptor instead.
func (*PreprocessMessageOptions) Descriptor() ([]byte, []int) {
	return file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_rawDescGZIP(), []int{0}
}

func (x *PreprocessMessageOptions) GetString_() *PreprocessString {
	if x != nil {
		return x.String_
	}
	return nil
}

type PreprocessFieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	String_ *PreprocessString `protobuf:"bytes,1,opt,name=string,proto3" json:"string,omitempty"`
}

func (x *PreprocessFieldOptions) Reset() {
	*x = PreprocessFieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreprocessFieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreprocessFieldOptions) ProtoMessage() {}

func (x *PreprocessFieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreprocessFieldOptions.ProtoReflect.Descriptor instead.
func (*PreprocessFieldOptions) Descriptor() ([]byte, []int) {
	return file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_rawDescGZIP(), []int{1}
}

func (x *PreprocessFieldOptions) GetString_() *PreprocessString {
	if x != nil {
		return x.String_
	}
	return nil
}

type PreprocessString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Methods []PreprocessString_Methods `protobuf:"varint,1,rep,packed,name=methods,proto3,enum=preprocess.PreprocessString_Methods" json:"methods,omitempty"`
}

func (x *PreprocessString) Reset() {
	*x = PreprocessString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreprocessString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreprocessString) ProtoMessage() {}

func (x *PreprocessString) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreprocessString.ProtoReflect.Descriptor instead.
func (*PreprocessString) Descriptor() ([]byte, []int) {
	return file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_rawDescGZIP(), []int{2}
}

func (x *PreprocessString) GetMethods() []PreprocessString_Methods {
	if x != nil {
		return x.Methods
	}
	return nil
}

var file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*PreprocessMessageOptions)(nil),
		Field:         11111,
		Name:          "preprocess.each",
		Tag:           "bytes,11111,opt,name=each",
		Filename:      "github.com/infobloxopen/protoc-gen-preprocess/options/preprocess.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*PreprocessFieldOptions)(nil),
		Field:         11111,
		Name:          "preprocess.field",
		Tag:           "bytes,11111,opt,name=field",
		Filename:      "github.com/infobloxopen/protoc-gen-preprocess/options/preprocess.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional preprocess.PreprocessMessageOptions each = 11111;
	E_Each = &file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional preprocess.PreprocessFieldOptions field = 11111;
	E_Field = &file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_extTypes[1]
)

var File_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto protoreflect.FileDescriptor

var file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_rawDesc = []byte{
	0x0a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66,
	0x6f, 0x62, 0x6c, 0x6f, 0x78, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x70, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x65, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x18, 0x50, 0x72, 0x65, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x50, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x4e, 0x0a, 0x16, 0x50, 0x72, 0x65, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x50, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x92, 0x01, 0x0a, 0x10, 0x50, 0x72, 0x65,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x3e, 0x0a,
	0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x24,
	0x2e, 0x70, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x65, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x22, 0x3e, 0x0a,
	0x07, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x6e, 0x6f, 0x6e, 0x65,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x74, 0x72, 0x69, 0x6d, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x75, 0x70, 0x70, 0x65, 0x72, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x10, 0x04, 0x3a, 0x5a, 0x0a,
	0x04, 0x65, 0x61, 0x63, 0x68, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe7, 0x56, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x70, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x04, 0x65, 0x61, 0x63, 0x68, 0x3a, 0x58, 0x0a, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xe7, 0x56, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x65, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x6c, 0x6f, 0x78, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x70, 0x72, 0x65, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x70, 0x72, 0x65,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_rawDescOnce sync.Once
	file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_rawDescData = file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_rawDesc
)

func file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_rawDescGZIP() []byte {
	file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_rawDescOnce.Do(func() {
		file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_rawDescData)
	})
	return file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_rawDescData
}

var file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_goTypes = []interface{}{
	(PreprocessString_Methods)(0),       // 0: preprocess.PreprocessString.Methods
	(*PreprocessMessageOptions)(nil),    // 1: preprocess.PreprocessMessageOptions
	(*PreprocessFieldOptions)(nil),      // 2: preprocess.PreprocessFieldOptions
	(*PreprocessString)(nil),            // 3: preprocess.PreprocessString
	(*descriptorpb.MessageOptions)(nil), // 4: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 5: google.protobuf.FieldOptions
}
var file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_depIdxs = []int32{
	3, // 0: preprocess.PreprocessMessageOptions.string:type_name -> preprocess.PreprocessString
	3, // 1: preprocess.PreprocessFieldOptions.string:type_name -> preprocess.PreprocessString
	0, // 2: preprocess.PreprocessString.methods:type_name -> preprocess.PreprocessString.Methods
	4, // 3: preprocess.each:extendee -> google.protobuf.MessageOptions
	5, // 4: preprocess.field:extendee -> google.protobuf.FieldOptions
	1, // 5: preprocess.each:type_name -> preprocess.PreprocessMessageOptions
	2, // 6: preprocess.field:type_name -> preprocess.PreprocessFieldOptions
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	5, // [5:7] is the sub-list for extension type_name
	3, // [3:5] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_init() }
func file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_init() {
	if File_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreprocessMessageOptions); i {
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
		file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreprocessFieldOptions); i {
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
		file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreprocessString); i {
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
			RawDescriptor: file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_goTypes,
		DependencyIndexes: file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_depIdxs,
		EnumInfos:         file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_enumTypes,
		MessageInfos:      file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_msgTypes,
		ExtensionInfos:    file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_extTypes,
	}.Build()
	File_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto = out.File
	file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_rawDesc = nil
	file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_goTypes = nil
	file_github_com_infobloxopen_protoc_gen_preprocess_options_preprocess_proto_depIdxs = nil
}
