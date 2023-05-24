// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: priceAlert.proto

package storage

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

type PriceAlert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         uint32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Uuid       string  `protobuf:"bytes,2,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
	Status     string  `protobuf:"bytes,3,opt,name=Status,proto3" json:"Status,omitempty"`
	Exchange   string  `protobuf:"bytes,4,opt,name=Exchange,proto3" json:"Exchange,omitempty"`
	Name       string  `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
	Wallet     string  `protobuf:"bytes,6,opt,name=Wallet,proto3" json:"Wallet,omitempty"`
	Token      string  `protobuf:"bytes,7,opt,name=Token,proto3" json:"Token,omitempty"`
	Chain      string  `protobuf:"bytes,8,opt,name=Chain,proto3" json:"Chain,omitempty"`
	Pair       string  `protobuf:"bytes,9,opt,name=Pair,proto3" json:"Pair,omitempty"`
	Flipped    bool    `protobuf:"varint,10,opt,name=Flipped,proto3" json:"Flipped,omitempty"`
	Price      float64 `protobuf:"fixed64,11,opt,name=Price,proto3" json:"Price,omitempty"`
	PriceUSD   float64 `protobuf:"fixed64,12,opt,name=PriceUSD,proto3" json:"PriceUSD,omitempty"`
	Gt         bool    `protobuf:"varint,13,opt,name=Gt,proto3" json:"Gt,omitempty"`
	CreatedAt  int64   `protobuf:"varint,14,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	ExecutedAt int64   `protobuf:"varint,15,opt,name=ExecutedAt,proto3" json:"ExecutedAt,omitempty"`
	Hidden     bool    `protobuf:"varint,16,opt,name=Hidden,proto3" json:"Hidden,omitempty"`
}

func (x *PriceAlert) Reset() {
	*x = PriceAlert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_priceAlert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceAlert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceAlert) ProtoMessage() {}

func (x *PriceAlert) ProtoReflect() protoreflect.Message {
	mi := &file_priceAlert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceAlert.ProtoReflect.Descriptor instead.
func (*PriceAlert) Descriptor() ([]byte, []int) {
	return file_priceAlert_proto_rawDescGZIP(), []int{0}
}

func (x *PriceAlert) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PriceAlert) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *PriceAlert) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PriceAlert) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *PriceAlert) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PriceAlert) GetWallet() string {
	if x != nil {
		return x.Wallet
	}
	return ""
}

func (x *PriceAlert) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *PriceAlert) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *PriceAlert) GetPair() string {
	if x != nil {
		return x.Pair
	}
	return ""
}

func (x *PriceAlert) GetFlipped() bool {
	if x != nil {
		return x.Flipped
	}
	return false
}

func (x *PriceAlert) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *PriceAlert) GetPriceUSD() float64 {
	if x != nil {
		return x.PriceUSD
	}
	return 0
}

func (x *PriceAlert) GetGt() bool {
	if x != nil {
		return x.Gt
	}
	return false
}

func (x *PriceAlert) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *PriceAlert) GetExecutedAt() int64 {
	if x != nil {
		return x.ExecutedAt
	}
	return 0
}

func (x *PriceAlert) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

var File_priceAlert_proto protoreflect.FileDescriptor

var file_priceAlert_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x69, 0x63, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x22, 0x82, 0x03, 0x0a, 0x0a, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x75, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x69,
	0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x46, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x46, 0x6c, 0x69, 0x70, 0x70, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x53, 0x44, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x53, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x47, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x47, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x69, 0x64, 0x64, 0x65,
	0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x42,
	0x12, 0x5a, 0x10, 0x2f, 0x70, 0x6b, 0x67, 0x2d, 0x62, 0x69, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_priceAlert_proto_rawDescOnce sync.Once
	file_priceAlert_proto_rawDescData = file_priceAlert_proto_rawDesc
)

func file_priceAlert_proto_rawDescGZIP() []byte {
	file_priceAlert_proto_rawDescOnce.Do(func() {
		file_priceAlert_proto_rawDescData = protoimpl.X.CompressGZIP(file_priceAlert_proto_rawDescData)
	})
	return file_priceAlert_proto_rawDescData
}

var file_priceAlert_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_priceAlert_proto_goTypes = []interface{}{
	(*PriceAlert)(nil), // 0: binary.PriceAlert
}
var file_priceAlert_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_priceAlert_proto_init() }
func file_priceAlert_proto_init() {
	if File_priceAlert_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_priceAlert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceAlert); i {
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
			RawDescriptor: file_priceAlert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_priceAlert_proto_goTypes,
		DependencyIndexes: file_priceAlert_proto_depIdxs,
		MessageInfos:      file_priceAlert_proto_msgTypes,
	}.Build()
	File_priceAlert_proto = out.File
	file_priceAlert_proto_rawDesc = nil
	file_priceAlert_proto_goTypes = nil
	file_priceAlert_proto_depIdxs = nil
}