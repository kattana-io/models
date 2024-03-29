// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: newPair.proto

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

type NewPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Factory     string `protobuf:"bytes,1,opt,name=Factory,proto3" json:"Factory,omitempty"`
	Pair        string `protobuf:"bytes,2,opt,name=Pair,proto3" json:"Pair,omitempty"`
	Klass       string `protobuf:"bytes,3,opt,name=Klass,proto3" json:"Klass,omitempty"`
	Network     string `protobuf:"bytes,4,opt,name=Network,proto3" json:"Network,omitempty"`
	Node        string `protobuf:"bytes,5,opt,name=Node,proto3" json:"Node,omitempty"`
	PoolCreated int64  `protobuf:"varint,6,opt,name=PoolCreated,proto3" json:"PoolCreated,omitempty"`
}

func (x *NewPair) Reset() {
	*x = NewPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_newPair_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPair) ProtoMessage() {}

func (x *NewPair) ProtoReflect() protoreflect.Message {
	mi := &file_newPair_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPair.ProtoReflect.Descriptor instead.
func (*NewPair) Descriptor() ([]byte, []int) {
	return file_newPair_proto_rawDescGZIP(), []int{0}
}

func (x *NewPair) GetFactory() string {
	if x != nil {
		return x.Factory
	}
	return ""
}

func (x *NewPair) GetPair() string {
	if x != nil {
		return x.Pair
	}
	return ""
}

func (x *NewPair) GetKlass() string {
	if x != nil {
		return x.Klass
	}
	return ""
}

func (x *NewPair) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *NewPair) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *NewPair) GetPoolCreated() int64 {
	if x != nil {
		return x.PoolCreated
	}
	return 0
}

var File_newPair_proto protoreflect.FileDescriptor

var file_newPair_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x22, 0x9d, 0x01, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x50,
	0x61, 0x69, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x50, 0x61, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x69,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x4b, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x4b, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x6f, 0x6f, 0x6c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x50, 0x6f, 0x6f, 0x6c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x12, 0x5a, 0x10, 0x2f, 0x70, 0x6b, 0x67, 0x2d,
	0x62, 0x69, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_newPair_proto_rawDescOnce sync.Once
	file_newPair_proto_rawDescData = file_newPair_proto_rawDesc
)

func file_newPair_proto_rawDescGZIP() []byte {
	file_newPair_proto_rawDescOnce.Do(func() {
		file_newPair_proto_rawDescData = protoimpl.X.CompressGZIP(file_newPair_proto_rawDescData)
	})
	return file_newPair_proto_rawDescData
}

var file_newPair_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_newPair_proto_goTypes = []interface{}{
	(*NewPair)(nil), // 0: binary.NewPair
}
var file_newPair_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_newPair_proto_init() }
func file_newPair_proto_init() {
	if File_newPair_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_newPair_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPair); i {
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
			RawDescriptor: file_newPair_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_newPair_proto_goTypes,
		DependencyIndexes: file_newPair_proto_depIdxs,
		MessageInfos:      file_newPair_proto_msgTypes,
	}.Build()
	File_newPair_proto = out.File
	file_newPair_proto_rawDesc = nil
	file_newPair_proto_goTypes = nil
	file_newPair_proto_depIdxs = nil
}
