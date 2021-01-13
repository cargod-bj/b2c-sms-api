// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.9.0
// source: sms/sms.proto

package sms

import (
	common "github.com/cargod-bj/b2c-proto-common/common"
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

type SmsDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone      string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Content    string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	From       string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	TempleteId uint64 `protobuf:"varint,4,opt,name=templeteId,proto3" json:"templeteId,omitempty"`
	//使用模板时要替换的参数
	Param         []string `protobuf:"bytes,5,rep,name=param,proto3" json:"param,omitempty"`
	CountryPrefix string   `protobuf:"bytes,6,opt,name=countryPrefix,proto3" json:"countryPrefix,omitempty"`
	Url           string   `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	Code          string   `protobuf:"bytes,8,opt,name=code,proto3" json:"code,omitempty"` //编码
}

func (x *SmsDto) Reset() {
	*x = SmsDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_sms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmsDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsDto) ProtoMessage() {}

func (x *SmsDto) ProtoReflect() protoreflect.Message {
	mi := &file_sms_sms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsDto.ProtoReflect.Descriptor instead.
func (*SmsDto) Descriptor() ([]byte, []int) {
	return file_sms_sms_proto_rawDescGZIP(), []int{0}
}

func (x *SmsDto) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SmsDto) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SmsDto) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SmsDto) GetTempleteId() uint64 {
	if x != nil {
		return x.TempleteId
	}
	return 0
}

func (x *SmsDto) GetParam() []string {
	if x != nil {
		return x.Param
	}
	return nil
}

func (x *SmsDto) GetCountryPrefix() string {
	if x != nil {
		return x.CountryPrefix
	}
	return ""
}

func (x *SmsDto) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SmsDto) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_sms_sms_proto protoreflect.FileDescriptor

var file_sms_sms_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x6d, 0x73, 0x2f, 0x73, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32,
	0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x06, 0x53, 0x6d, 0x73, 0x44, 0x74, 0x6f, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x71, 0x0a, 0x03, 0x53, 0x6d, 0x73, 0x12, 0x33, 0x0a, 0x0d, 0x53, 0x65,
	0x6e, 0x64, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x53, 0x6d, 0x73, 0x12, 0x0e, 0x2e, 0x64, 0x65,
	0x61, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x6d, 0x73, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x35, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x6d, 0x73, 0x12, 0x0e, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x6d, 0x73, 0x44,
	0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62,
	0x32, 0x63, 0x2d, 0x73, 0x6d, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x6d, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sms_sms_proto_rawDescOnce sync.Once
	file_sms_sms_proto_rawDescData = file_sms_sms_proto_rawDesc
)

func file_sms_sms_proto_rawDescGZIP() []byte {
	file_sms_sms_proto_rawDescOnce.Do(func() {
		file_sms_sms_proto_rawDescData = protoimpl.X.CompressGZIP(file_sms_sms_proto_rawDescData)
	})
	return file_sms_sms_proto_rawDescData
}

var file_sms_sms_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sms_sms_proto_goTypes = []interface{}{
	(*SmsDto)(nil),          // 0: dealer.SmsDto
	(*common.Response)(nil), // 1: common.Response
}
var file_sms_sms_proto_depIdxs = []int32{
	0, // 0: dealer.Sms.SendNormalSms:input_type -> dealer.SmsDto
	0, // 1: dealer.Sms.SendTempleteSms:input_type -> dealer.SmsDto
	1, // 2: dealer.Sms.SendNormalSms:output_type -> common.Response
	1, // 3: dealer.Sms.SendTempleteSms:output_type -> common.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sms_sms_proto_init() }
func file_sms_sms_proto_init() {
	if File_sms_sms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sms_sms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmsDto); i {
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
			RawDescriptor: file_sms_sms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sms_sms_proto_goTypes,
		DependencyIndexes: file_sms_sms_proto_depIdxs,
		MessageInfos:      file_sms_sms_proto_msgTypes,
	}.Build()
	File_sms_sms_proto = out.File
	file_sms_sms_proto_rawDesc = nil
	file_sms_sms_proto_goTypes = nil
	file_sms_sms_proto_depIdxs = nil
}
