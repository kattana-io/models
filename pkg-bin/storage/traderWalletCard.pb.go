// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: traderWalletCard.proto

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

type TraderWallet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wallet string `protobuf:"bytes,1,opt,name=Wallet,proto3" json:"Wallet,omitempty"`
	Pair   string `protobuf:"bytes,2,opt,name=Pair,proto3" json:"Pair,omitempty"`
	Count  uint64 `protobuf:"varint,3,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *TraderWallet) Reset() {
	*x = TraderWallet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traderWalletCard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraderWallet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraderWallet) ProtoMessage() {}

func (x *TraderWallet) ProtoReflect() protoreflect.Message {
	mi := &file_traderWalletCard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraderWallet.ProtoReflect.Descriptor instead.
func (*TraderWallet) Descriptor() ([]byte, []int) {
	return file_traderWalletCard_proto_rawDescGZIP(), []int{0}
}

func (x *TraderWallet) GetWallet() string {
	if x != nil {
		return x.Wallet
	}
	return ""
}

func (x *TraderWallet) GetPair() string {
	if x != nil {
		return x.Pair
	}
	return ""
}

func (x *TraderWallet) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type TraderWalletsBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain     string          `protobuf:"bytes,1,opt,name=Chain,proto3" json:"Chain,omitempty"`
	Timestamp uint64          `protobuf:"varint,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Live      bool            `protobuf:"varint,3,opt,name=Live,proto3" json:"Live,omitempty"`
	Trades    []*TraderWallet `protobuf:"bytes,4,rep,name=Trades,proto3" json:"Trades,omitempty"`
}

func (x *TraderWalletsBlock) Reset() {
	*x = TraderWalletsBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traderWalletCard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraderWalletsBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraderWalletsBlock) ProtoMessage() {}

func (x *TraderWalletsBlock) ProtoReflect() protoreflect.Message {
	mi := &file_traderWalletCard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraderWalletsBlock.ProtoReflect.Descriptor instead.
func (*TraderWalletsBlock) Descriptor() ([]byte, []int) {
	return file_traderWalletCard_proto_rawDescGZIP(), []int{1}
}

func (x *TraderWalletsBlock) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *TraderWalletsBlock) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *TraderWalletsBlock) GetLive() bool {
	if x != nil {
		return x.Live
	}
	return false
}

func (x *TraderWalletsBlock) GetTrades() []*TraderWallet {
	if x != nil {
		return x.Trades
	}
	return nil
}

type TraderWalletPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuyTradesCount  uint64  `protobuf:"varint,1,opt,name=BuyTradesCount,proto3" json:"BuyTradesCount,omitempty"`
	SellTradesCount uint64  `protobuf:"varint,2,opt,name=SellTradesCount,proto3" json:"SellTradesCount,omitempty"`
	BuyValue        float32 `protobuf:"fixed32,3,opt,name=BuyValue,proto3" json:"BuyValue,omitempty"`
	BuyValueUsd     float32 `protobuf:"fixed32,4,opt,name=BuyValueUsd,proto3" json:"BuyValueUsd,omitempty"`
	SellValue       float32 `protobuf:"fixed32,5,opt,name=SellValue,proto3" json:"SellValue,omitempty"`
	SellValueUsd    float32 `protobuf:"fixed32,6,opt,name=SellValueUsd,proto3" json:"SellValueUsd,omitempty"`
}

func (x *TraderWalletPair) Reset() {
	*x = TraderWalletPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traderWalletCard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraderWalletPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraderWalletPair) ProtoMessage() {}

func (x *TraderWalletPair) ProtoReflect() protoreflect.Message {
	mi := &file_traderWalletCard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraderWalletPair.ProtoReflect.Descriptor instead.
func (*TraderWalletPair) Descriptor() ([]byte, []int) {
	return file_traderWalletCard_proto_rawDescGZIP(), []int{2}
}

func (x *TraderWalletPair) GetBuyTradesCount() uint64 {
	if x != nil {
		return x.BuyTradesCount
	}
	return 0
}

func (x *TraderWalletPair) GetSellTradesCount() uint64 {
	if x != nil {
		return x.SellTradesCount
	}
	return 0
}

func (x *TraderWalletPair) GetBuyValue() float32 {
	if x != nil {
		return x.BuyValue
	}
	return 0
}

func (x *TraderWalletPair) GetBuyValueUsd() float32 {
	if x != nil {
		return x.BuyValueUsd
	}
	return 0
}

func (x *TraderWalletPair) GetSellValue() float32 {
	if x != nil {
		return x.SellValue
	}
	return 0
}

func (x *TraderWalletPair) GetSellValueUsd() float32 {
	if x != nil {
		return x.SellValueUsd
	}
	return 0
}

type TraderWalletCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstTradeAt uint64            `protobuf:"varint,1,opt,name=FirstTradeAt,proto3" json:"FirstTradeAt,omitempty"`
	LastTradeAt  uint64            `protobuf:"varint,2,opt,name=LastTradeAt,proto3" json:"LastTradeAt,omitempty"`
	CreatedAt    uint64            `protobuf:"varint,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt    uint64            `protobuf:"varint,4,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	TotalTrades  uint64            `protobuf:"varint,5,opt,name=TotalTrades,proto3" json:"TotalTrades,omitempty"`
	Pairs        map[string]uint64 `protobuf:"bytes,6,rep,name=Pairs,proto3" json:"Pairs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	PairInfo     *TraderWalletPair `protobuf:"bytes,7,opt,name=pairInfo,proto3" json:"pairInfo,omitempty"`
}

func (x *TraderWalletCard) Reset() {
	*x = TraderWalletCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traderWalletCard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraderWalletCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraderWalletCard) ProtoMessage() {}

func (x *TraderWalletCard) ProtoReflect() protoreflect.Message {
	mi := &file_traderWalletCard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraderWalletCard.ProtoReflect.Descriptor instead.
func (*TraderWalletCard) Descriptor() ([]byte, []int) {
	return file_traderWalletCard_proto_rawDescGZIP(), []int{3}
}

func (x *TraderWalletCard) GetFirstTradeAt() uint64 {
	if x != nil {
		return x.FirstTradeAt
	}
	return 0
}

func (x *TraderWalletCard) GetLastTradeAt() uint64 {
	if x != nil {
		return x.LastTradeAt
	}
	return 0
}

func (x *TraderWalletCard) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *TraderWalletCard) GetUpdatedAt() uint64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *TraderWalletCard) GetTotalTrades() uint64 {
	if x != nil {
		return x.TotalTrades
	}
	return 0
}

func (x *TraderWalletCard) GetPairs() map[string]uint64 {
	if x != nil {
		return x.Pairs
	}
	return nil
}

func (x *TraderWalletCard) GetPairInfo() *TraderWalletPair {
	if x != nil {
		return x.PairInfo
	}
	return nil
}

var File_traderWalletCard_proto protoreflect.FileDescriptor

var file_traderWalletCard_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x72, 0x61, 0x64, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x22, 0x50, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x64, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x8a, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x64, 0x65, 0x72, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x4c, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x4c, 0x69, 0x76,
	0x65, 0x12, 0x2c, 0x0a, 0x06, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x06, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x22,
	0xe4, 0x01, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x64, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x50, 0x61, 0x69, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x42, 0x75, 0x79, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x42, 0x75,
	0x79, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f,
	0x53, 0x65, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x53, 0x65, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x75, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x42, 0x75, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x75, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x55, 0x73,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x42, 0x75, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x55, 0x73, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x53, 0x65, 0x6c, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x65, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x55,
	0x73, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x53, 0x65, 0x6c, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x55, 0x73, 0x64, 0x22, 0xe1, 0x02, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x46,
	0x69, 0x72, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x41, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0c, 0x46, 0x69, 0x72, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x41, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x41, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x41,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x12,
	0x39, 0x0a, 0x05, 0x50, 0x61, 0x69, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x72, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x50, 0x61, 0x69, 0x72, 0x73, 0x12, 0x34, 0x0a, 0x08, 0x70, 0x61,
	0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x50, 0x61, 0x69, 0x72, 0x52, 0x08, 0x70, 0x61, 0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x38, 0x0a, 0x0a, 0x50, 0x61, 0x69, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x12, 0x5a, 0x10, 0x2f, 0x70,
	0x6b, 0x67, 0x2d, 0x62, 0x69, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_traderWalletCard_proto_rawDescOnce sync.Once
	file_traderWalletCard_proto_rawDescData = file_traderWalletCard_proto_rawDesc
)

func file_traderWalletCard_proto_rawDescGZIP() []byte {
	file_traderWalletCard_proto_rawDescOnce.Do(func() {
		file_traderWalletCard_proto_rawDescData = protoimpl.X.CompressGZIP(file_traderWalletCard_proto_rawDescData)
	})
	return file_traderWalletCard_proto_rawDescData
}

var file_traderWalletCard_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_traderWalletCard_proto_goTypes = []interface{}{
	(*TraderWallet)(nil),       // 0: binary.TraderWallet
	(*TraderWalletsBlock)(nil), // 1: binary.TraderWalletsBlock
	(*TraderWalletPair)(nil),   // 2: binary.TraderWalletPair
	(*TraderWalletCard)(nil),   // 3: binary.TraderWalletCard
	nil,                        // 4: binary.TraderWalletCard.PairsEntry
}
var file_traderWalletCard_proto_depIdxs = []int32{
	0, // 0: binary.TraderWalletsBlock.Trades:type_name -> binary.TraderWallet
	4, // 1: binary.TraderWalletCard.Pairs:type_name -> binary.TraderWalletCard.PairsEntry
	2, // 2: binary.TraderWalletCard.pairInfo:type_name -> binary.TraderWalletPair
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_traderWalletCard_proto_init() }
func file_traderWalletCard_proto_init() {
	if File_traderWalletCard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_traderWalletCard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraderWallet); i {
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
		file_traderWalletCard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraderWalletsBlock); i {
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
		file_traderWalletCard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraderWalletPair); i {
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
		file_traderWalletCard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraderWalletCard); i {
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
			RawDescriptor: file_traderWalletCard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_traderWalletCard_proto_goTypes,
		DependencyIndexes: file_traderWalletCard_proto_depIdxs,
		MessageInfos:      file_traderWalletCard_proto_msgTypes,
	}.Build()
	File_traderWalletCard_proto = out.File
	file_traderWalletCard_proto_rawDesc = nil
	file_traderWalletCard_proto_goTypes = nil
	file_traderWalletCard_proto_depIdxs = nil
}