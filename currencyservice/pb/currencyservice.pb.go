// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: currencyservice.proto

package pb

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

// 当前货币描述
type Currency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 货币code 例如：人民币、美元、欧元
	CurrencyCode string `protobuf:"bytes,1,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// 货币单位
	Units int32 `protobuf:"varint,2,opt,name=units,proto3" json:"units,omitempty"`
	// 数量的纳米单位
	Nanos int32 `protobuf:"varint,3,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *Currency) Reset() {
	*x = Currency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currencyservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Currency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Currency) ProtoMessage() {}

func (x *Currency) ProtoReflect() protoreflect.Message {
	mi := &file_currencyservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Currency.ProtoReflect.Descriptor instead.
func (*Currency) Descriptor() ([]byte, []int) {
	return file_currencyservice_proto_rawDescGZIP(), []int{0}
}

func (x *Currency) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *Currency) GetUnits() int32 {
	if x != nil {
		return x.Units
	}
	return 0
}

func (x *Currency) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currencyservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_currencyservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_currencyservice_proto_rawDescGZIP(), []int{1}
}

// 货币转换请求消息
type CurrencyConvertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin *Currency `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	ToCode string    `protobuf:"bytes,2,opt,name=to_code,json=toCode,proto3" json:"to_code,omitempty"`
}

func (x *CurrencyConvertRequest) Reset() {
	*x = CurrencyConvertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currencyservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyConvertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyConvertRequest) ProtoMessage() {}

func (x *CurrencyConvertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_currencyservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyConvertRequest.ProtoReflect.Descriptor instead.
func (*CurrencyConvertRequest) Descriptor() ([]byte, []int) {
	return file_currencyservice_proto_rawDescGZIP(), []int{2}
}

func (x *CurrencyConvertRequest) GetOrigin() *Currency {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *CurrencyConvertRequest) GetToCode() string {
	if x != nil {
		return x.ToCode
	}
	return ""
}

// 获得支持的币种的响应信息
type GetSupportedCurrenciesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrencyCode []string `protobuf:"bytes,1,rep,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
}

func (x *GetSupportedCurrenciesResponse) Reset() {
	*x = GetSupportedCurrenciesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currencyservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSupportedCurrenciesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSupportedCurrenciesResponse) ProtoMessage() {}

func (x *GetSupportedCurrenciesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_currencyservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSupportedCurrenciesResponse.ProtoReflect.Descriptor instead.
func (*GetSupportedCurrenciesResponse) Descriptor() ([]byte, []int) {
	return file_currencyservice_proto_rawDescGZIP(), []int{3}
}

func (x *GetSupportedCurrenciesResponse) GetCurrencyCode() []string {
	if x != nil {
		return x.CurrencyCode
	}
	return nil
}

var File_currencyservice_proto protoreflect.FileDescriptor

var file_currencyservice_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b,
	0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x75, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5a, 0x0a, 0x16, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x45, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x53, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x32, 0xa6,
	0x01, 0x0a, 0x0f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x56, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x72, 0x73, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_currencyservice_proto_rawDescOnce sync.Once
	file_currencyservice_proto_rawDescData = file_currencyservice_proto_rawDesc
)

func file_currencyservice_proto_rawDescGZIP() []byte {
	file_currencyservice_proto_rawDescOnce.Do(func() {
		file_currencyservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_currencyservice_proto_rawDescData)
	})
	return file_currencyservice_proto_rawDescData
}

var file_currencyservice_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_currencyservice_proto_goTypes = []interface{}{
	(*Currency)(nil),                       // 0: proto.Currency
	(*EmptyRequest)(nil),                   // 1: proto.EmptyRequest
	(*CurrencyConvertRequest)(nil),         // 2: proto.CurrencyConvertRequest
	(*GetSupportedCurrenciesResponse)(nil), // 3: proto.GetSupportedCurrenciesResponse
}
var file_currencyservice_proto_depIdxs = []int32{
	0, // 0: proto.CurrencyConvertRequest.origin:type_name -> proto.Currency
	1, // 1: proto.CurrencyService.GetSupportedCurrencies:input_type -> proto.EmptyRequest
	2, // 2: proto.CurrencyService.Convert:input_type -> proto.CurrencyConvertRequest
	3, // 3: proto.CurrencyService.GetSupportedCurrencies:output_type -> proto.GetSupportedCurrenciesResponse
	0, // 4: proto.CurrencyService.Convert:output_type -> proto.Currency
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_currencyservice_proto_init() }
func file_currencyservice_proto_init() {
	if File_currencyservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_currencyservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Currency); i {
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
		file_currencyservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_currencyservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyConvertRequest); i {
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
		file_currencyservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSupportedCurrenciesResponse); i {
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
			RawDescriptor: file_currencyservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_currencyservice_proto_goTypes,
		DependencyIndexes: file_currencyservice_proto_depIdxs,
		MessageInfos:      file_currencyservice_proto_msgTypes,
	}.Build()
	File_currencyservice_proto = out.File
	file_currencyservice_proto_rawDesc = nil
	file_currencyservice_proto_goTypes = nil
	file_currencyservice_proto_depIdxs = nil
}