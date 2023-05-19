// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: candles.proto

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

type Candle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Open  float64 `protobuf:"fixed64,1,opt,name=Open,proto3" json:"Open,omitempty"`
	Low   float64 `protobuf:"fixed64,2,opt,name=Low,proto3" json:"Low,omitempty"`
	High  float64 `protobuf:"fixed64,3,opt,name=High,proto3" json:"High,omitempty"`
	Date  uint64  `protobuf:"varint,4,opt,name=Date,proto3" json:"Date,omitempty"`
	Close float64 `protobuf:"fixed64,5,opt,name=Close,proto3" json:"Close,omitempty"`
}

func (x *Candle) Reset() {
	*x = Candle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_candles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Candle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candle) ProtoMessage() {}

func (x *Candle) ProtoReflect() protoreflect.Message {
	mi := &file_candles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candle.ProtoReflect.Descriptor instead.
func (*Candle) Descriptor() ([]byte, []int) {
	return file_candles_proto_rawDescGZIP(), []int{0}
}

func (x *Candle) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *Candle) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *Candle) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *Candle) GetDate() uint64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *Candle) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

type BlockLite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network     string        `protobuf:"bytes,1,opt,name=Network,proto3" json:"Network,omitempty"`
	Date        uint64        `protobuf:"varint,2,opt,name=Date,proto3" json:"Date,omitempty"`
	Trades      []*Trade      `protobuf:"bytes,3,rep,name=Trades,proto3" json:"Trades,omitempty"`
	PairSwaps   []*PairSwap   `protobuf:"bytes,4,rep,name=PairSwaps,proto3" json:"PairSwaps,omitempty"`
	DirectSwaps []*DirectSwap `protobuf:"bytes,5,rep,name=DirectSwaps,proto3" json:"DirectSwaps,omitempty"`
}

func (x *BlockLite) Reset() {
	*x = BlockLite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_candles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockLite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockLite) ProtoMessage() {}

func (x *BlockLite) ProtoReflect() protoreflect.Message {
	mi := &file_candles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockLite.ProtoReflect.Descriptor instead.
func (*BlockLite) Descriptor() ([]byte, []int) {
	return file_candles_proto_rawDescGZIP(), []int{1}
}

func (x *BlockLite) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *BlockLite) GetDate() uint64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *BlockLite) GetTrades() []*Trade {
	if x != nil {
		return x.Trades
	}
	return nil
}

func (x *BlockLite) GetPairSwaps() []*PairSwap {
	if x != nil {
		return x.PairSwaps
	}
	return nil
}

func (x *BlockLite) GetDirectSwaps() []*DirectSwap {
	if x != nil {
		return x.DirectSwaps
	}
	return nil
}

type Trade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pair      string  `protobuf:"bytes,1,opt,name=Pair,proto3" json:"Pair,omitempty"`
	PriceA    float64 `protobuf:"fixed64,2,opt,name=PriceA,proto3" json:"PriceA,omitempty"`
	PriceB    float64 `protobuf:"fixed64,3,opt,name=PriceB,proto3" json:"PriceB,omitempty"`
	PriceAUSD float64 `protobuf:"fixed64,4,opt,name=PriceAUSD,proto3" json:"PriceAUSD,omitempty"`
	PriceBUSD float64 `protobuf:"fixed64,5,opt,name=PriceBUSD,proto3" json:"PriceBUSD,omitempty"`
	VolumeA   []byte  `protobuf:"bytes,6,opt,name=VolumeA,proto3" json:"VolumeA,omitempty"`
	VolumeB   []byte  `protobuf:"bytes,7,opt,name=VolumeB,proto3" json:"VolumeB,omitempty"`
	VolumeUSD float64 `protobuf:"fixed64,8,opt,name=VolumeUSD,proto3" json:"VolumeUSD,omitempty"`
}

func (x *Trade) Reset() {
	*x = Trade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_candles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trade) ProtoMessage() {}

func (x *Trade) ProtoReflect() protoreflect.Message {
	mi := &file_candles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trade.ProtoReflect.Descriptor instead.
func (*Trade) Descriptor() ([]byte, []int) {
	return file_candles_proto_rawDescGZIP(), []int{2}
}

func (x *Trade) GetPair() string {
	if x != nil {
		return x.Pair
	}
	return ""
}

func (x *Trade) GetPriceA() float64 {
	if x != nil {
		return x.PriceA
	}
	return 0
}

func (x *Trade) GetPriceB() float64 {
	if x != nil {
		return x.PriceB
	}
	return 0
}

func (x *Trade) GetPriceAUSD() float64 {
	if x != nil {
		return x.PriceAUSD
	}
	return 0
}

func (x *Trade) GetPriceBUSD() float64 {
	if x != nil {
		return x.PriceBUSD
	}
	return 0
}

func (x *Trade) GetVolumeA() []byte {
	if x != nil {
		return x.VolumeA
	}
	return nil
}

func (x *Trade) GetVolumeB() []byte {
	if x != nil {
		return x.VolumeB
	}
	return nil
}

func (x *Trade) GetVolumeUSD() float64 {
	if x != nil {
		return x.VolumeUSD
	}
	return 0
}

type Volume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date      uint64  `protobuf:"varint,1,opt,name=Date,proto3" json:"Date,omitempty"`
	VolumeA   []byte  `protobuf:"bytes,2,opt,name=VolumeA,proto3" json:"VolumeA,omitempty"`
	VolumeB   []byte  `protobuf:"bytes,3,opt,name=VolumeB,proto3" json:"VolumeB,omitempty"`
	VolumeUSD float64 `protobuf:"fixed64,4,opt,name=VolumeUSD,proto3" json:"VolumeUSD,omitempty"`
}

func (x *Volume) Reset() {
	*x = Volume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_candles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Volume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Volume) ProtoMessage() {}

func (x *Volume) ProtoReflect() protoreflect.Message {
	mi := &file_candles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Volume.ProtoReflect.Descriptor instead.
func (*Volume) Descriptor() ([]byte, []int) {
	return file_candles_proto_rawDescGZIP(), []int{3}
}

func (x *Volume) GetDate() uint64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *Volume) GetVolumeA() []byte {
	if x != nil {
		return x.VolumeA
	}
	return nil
}

func (x *Volume) GetVolumeB() []byte {
	if x != nil {
		return x.VolumeB
	}
	return nil
}

func (x *Volume) GetVolumeUSD() float64 {
	if x != nil {
		return x.VolumeUSD
	}
	return 0
}

type Candles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Candle []*Candle `protobuf:"bytes,1,rep,name=Candle,proto3" json:"Candle,omitempty"`
	Volume *Volume   `protobuf:"bytes,2,opt,name=Volume,proto3" json:"Volume,omitempty"`
}

func (x *Candles) Reset() {
	*x = Candles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_candles_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Candles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candles) ProtoMessage() {}

func (x *Candles) ProtoReflect() protoreflect.Message {
	mi := &file_candles_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candles.ProtoReflect.Descriptor instead.
func (*Candles) Descriptor() ([]byte, []int) {
	return file_candles_proto_rawDescGZIP(), []int{4}
}

func (x *Candles) GetCandle() []*Candle {
	if x != nil {
		return x.Candle
	}
	return nil
}

func (x *Candles) GetVolume() *Volume {
	if x != nil {
		return x.Volume
	}
	return nil
}

var File_candles_proto protoreflect.FileDescriptor

var file_candles_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x1a, 0x0e, 0x70, 0x61, 0x69, 0x72, 0x53, 0x77, 0x61,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53,
	0x77, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x06, 0x43, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x6f, 0x77, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x4c, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x69, 0x67,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x48, 0x69, 0x67, 0x68, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x22, 0xc6, 0x01, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x4c, 0x69, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x52, 0x06, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x09, 0x50, 0x61,
	0x69, 0x72, 0x53, 0x77, 0x61, 0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x53, 0x77, 0x61, 0x70, 0x52,
	0x09, 0x50, 0x61, 0x69, 0x72, 0x53, 0x77, 0x61, 0x70, 0x73, 0x12, 0x34, 0x0a, 0x0b, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x53, 0x77, 0x61, 0x70, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53,
	0x77, 0x61, 0x70, 0x52, 0x0b, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x77, 0x61, 0x70, 0x73,
	0x22, 0xd9, 0x01, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61,
	0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x50, 0x72, 0x69, 0x63, 0x65, 0x41, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x41, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x12, 0x1c,
	0x0a, 0x09, 0x50, 0x72, 0x69, 0x63, 0x65, 0x41, 0x55, 0x53, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x09, 0x50, 0x72, 0x69, 0x63, 0x65, 0x41, 0x55, 0x53, 0x44, 0x12, 0x1c, 0x0a, 0x09,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x55, 0x53, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x55, 0x53, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x41, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x41, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x12, 0x1c,
	0x0a, 0x09, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x55, 0x53, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x09, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x55, 0x53, 0x44, 0x22, 0x6e, 0x0a, 0x06,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x41, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x41, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x12, 0x1c,
	0x0a, 0x09, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x55, 0x53, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x09, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x55, 0x53, 0x44, 0x22, 0x59, 0x0a, 0x07,
	0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x2e, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x06, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12,
	0x26, 0x0a, 0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52,
	0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x2f, 0x70, 0x6b, 0x67, 0x2d,
	0x62, 0x69, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_candles_proto_rawDescOnce sync.Once
	file_candles_proto_rawDescData = file_candles_proto_rawDesc
)

func file_candles_proto_rawDescGZIP() []byte {
	file_candles_proto_rawDescOnce.Do(func() {
		file_candles_proto_rawDescData = protoimpl.X.CompressGZIP(file_candles_proto_rawDescData)
	})
	return file_candles_proto_rawDescData
}

var file_candles_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_candles_proto_goTypes = []interface{}{
	(*Candle)(nil),     // 0: binary.Candle
	(*BlockLite)(nil),  // 1: binary.BlockLite
	(*Trade)(nil),      // 2: binary.Trade
	(*Volume)(nil),     // 3: binary.Volume
	(*Candles)(nil),    // 4: binary.Candles
	(*PairSwap)(nil),   // 5: binary.PairSwap
	(*DirectSwap)(nil), // 6: binary.DirectSwap
}
var file_candles_proto_depIdxs = []int32{
	2, // 0: binary.BlockLite.Trades:type_name -> binary.Trade
	5, // 1: binary.BlockLite.PairSwaps:type_name -> binary.PairSwap
	6, // 2: binary.BlockLite.DirectSwaps:type_name -> binary.DirectSwap
	0, // 3: binary.Candles.Candle:type_name -> binary.Candle
	3, // 4: binary.Candles.Volume:type_name -> binary.Volume
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_candles_proto_init() }
func file_candles_proto_init() {
	if File_candles_proto != nil {
		return
	}
	file_pairSwap_proto_init()
	file_directSwap_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_candles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Candle); i {
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
		file_candles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockLite); i {
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
		file_candles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trade); i {
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
		file_candles_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Volume); i {
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
		file_candles_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Candles); i {
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
			RawDescriptor: file_candles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_candles_proto_goTypes,
		DependencyIndexes: file_candles_proto_depIdxs,
		MessageInfos:      file_candles_proto_msgTypes,
	}.Build()
	File_candles_proto = out.File
	file_candles_proto_rawDesc = nil
	file_candles_proto_goTypes = nil
	file_candles_proto_depIdxs = nil
}
