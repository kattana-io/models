// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: limitOrder.proto

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

type LimitOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                  uint32        `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Status              string        `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	Uuid                string        `protobuf:"bytes,3,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
	Wallet              string        `protobuf:"bytes,4,opt,name=Wallet,proto3" json:"Wallet,omitempty"`
	Chain               string        `protobuf:"bytes,5,opt,name=Chain,proto3" json:"Chain,omitempty"`
	Pair                string        `protobuf:"bytes,6,opt,name=Pair,proto3" json:"Pair,omitempty"`
	Flipped             bool          `protobuf:"varint,7,opt,name=Flipped,proto3" json:"Flipped,omitempty"`
	Price               float64       `protobuf:"fixed64,8,opt,name=Price,proto3" json:"Price,omitempty"`
	PriceUSD            float64       `protobuf:"fixed64,9,opt,name=PriceUSD,proto3" json:"PriceUSD,omitempty"`
	StartPrice          float64       `protobuf:"fixed64,10,opt,name=StartPrice,proto3" json:"StartPrice,omitempty"`
	StartPriceUSD       float64       `protobuf:"fixed64,11,opt,name=StartPriceUSD,proto3" json:"StartPriceUSD,omitempty"`
	EndPrice            float64       `protobuf:"fixed64,12,opt,name=EndPrice,proto3" json:"EndPrice,omitempty"`
	EndPriceUSD         float64       `protobuf:"fixed64,13,opt,name=EndPriceUSD,proto3" json:"EndPriceUSD,omitempty"`
	ExecutionType       string        `protobuf:"bytes,14,opt,name=ExecutionType,proto3" json:"ExecutionType,omitempty"`
	Gt                  bool          `protobuf:"varint,15,opt,name=Gt,proto3" json:"Gt,omitempty"`
	Name                string        `protobuf:"bytes,16,opt,name=Name,proto3" json:"Name,omitempty"`
	Exchange            string        `protobuf:"bytes,17,opt,name=Exchange,proto3" json:"Exchange,omitempty"`
	Proxy               string        `protobuf:"bytes,18,opt,name=Proxy,proto3" json:"Proxy,omitempty"`
	AmountIn            string        `protobuf:"bytes,19,opt,name=AmountIn,proto3" json:"AmountIn,omitempty"`
	AmountOut           string        `protobuf:"bytes,20,opt,name=AmountOut,proto3" json:"AmountOut,omitempty"`
	SrcAddress          string        `protobuf:"bytes,21,opt,name=SrcAddress,proto3" json:"SrcAddress,omitempty"`
	SrcDecimals         int32         `protobuf:"varint,22,opt,name=SrcDecimals,proto3" json:"SrcDecimals,omitempty"`
	DestAddress         string        `protobuf:"bytes,23,opt,name=DestAddress,proto3" json:"DestAddress,omitempty"`
	DestDecimals        int32         `protobuf:"varint,24,opt,name=DestDecimals,proto3" json:"DestDecimals,omitempty"`
	GasSent             int64         `protobuf:"varint,25,opt,name=GasSent,proto3" json:"GasSent,omitempty"`
	TotalGas            int64         `protobuf:"varint,26,opt,name=TotalGas,proto3" json:"TotalGas,omitempty"`
	ExpiresAt           int64         `protobuf:"varint,27,opt,name=ExpiresAt,proto3" json:"ExpiresAt,omitempty"`
	TriggerType         string        `protobuf:"bytes,28,opt,name=TriggerType,proto3" json:"TriggerType,omitempty"`
	Slippage            uint32        `protobuf:"varint,29,opt,name=Slippage,proto3" json:"Slippage,omitempty"`
	ExecutedTradesCount int64         `protobuf:"varint,30,opt,name=ExecutedTradesCount,proto3" json:"ExecutedTradesCount,omitempty"`
	TradesCount         int64         `protobuf:"varint,31,opt,name=TradesCount,proto3" json:"TradesCount,omitempty"`
	LimitTrades         []*LimitTrade `protobuf:"bytes,32,rep,name=LimitTrades,proto3" json:"LimitTrades,omitempty"`
}

func (x *LimitOrder) Reset() {
	*x = LimitOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limitOrder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitOrder) ProtoMessage() {}

func (x *LimitOrder) ProtoReflect() protoreflect.Message {
	mi := &file_limitOrder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitOrder.ProtoReflect.Descriptor instead.
func (*LimitOrder) Descriptor() ([]byte, []int) {
	return file_limitOrder_proto_rawDescGZIP(), []int{0}
}

func (x *LimitOrder) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *LimitOrder) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *LimitOrder) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *LimitOrder) GetWallet() string {
	if x != nil {
		return x.Wallet
	}
	return ""
}

func (x *LimitOrder) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *LimitOrder) GetPair() string {
	if x != nil {
		return x.Pair
	}
	return ""
}

func (x *LimitOrder) GetFlipped() bool {
	if x != nil {
		return x.Flipped
	}
	return false
}

func (x *LimitOrder) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *LimitOrder) GetPriceUSD() float64 {
	if x != nil {
		return x.PriceUSD
	}
	return 0
}

func (x *LimitOrder) GetStartPrice() float64 {
	if x != nil {
		return x.StartPrice
	}
	return 0
}

func (x *LimitOrder) GetStartPriceUSD() float64 {
	if x != nil {
		return x.StartPriceUSD
	}
	return 0
}

func (x *LimitOrder) GetEndPrice() float64 {
	if x != nil {
		return x.EndPrice
	}
	return 0
}

func (x *LimitOrder) GetEndPriceUSD() float64 {
	if x != nil {
		return x.EndPriceUSD
	}
	return 0
}

func (x *LimitOrder) GetExecutionType() string {
	if x != nil {
		return x.ExecutionType
	}
	return ""
}

func (x *LimitOrder) GetGt() bool {
	if x != nil {
		return x.Gt
	}
	return false
}

func (x *LimitOrder) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LimitOrder) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *LimitOrder) GetProxy() string {
	if x != nil {
		return x.Proxy
	}
	return ""
}

func (x *LimitOrder) GetAmountIn() string {
	if x != nil {
		return x.AmountIn
	}
	return ""
}

func (x *LimitOrder) GetAmountOut() string {
	if x != nil {
		return x.AmountOut
	}
	return ""
}

func (x *LimitOrder) GetSrcAddress() string {
	if x != nil {
		return x.SrcAddress
	}
	return ""
}

func (x *LimitOrder) GetSrcDecimals() int32 {
	if x != nil {
		return x.SrcDecimals
	}
	return 0
}

func (x *LimitOrder) GetDestAddress() string {
	if x != nil {
		return x.DestAddress
	}
	return ""
}

func (x *LimitOrder) GetDestDecimals() int32 {
	if x != nil {
		return x.DestDecimals
	}
	return 0
}

func (x *LimitOrder) GetGasSent() int64 {
	if x != nil {
		return x.GasSent
	}
	return 0
}

func (x *LimitOrder) GetTotalGas() int64 {
	if x != nil {
		return x.TotalGas
	}
	return 0
}

func (x *LimitOrder) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *LimitOrder) GetTriggerType() string {
	if x != nil {
		return x.TriggerType
	}
	return ""
}

func (x *LimitOrder) GetSlippage() uint32 {
	if x != nil {
		return x.Slippage
	}
	return 0
}

func (x *LimitOrder) GetExecutedTradesCount() int64 {
	if x != nil {
		return x.ExecutedTradesCount
	}
	return 0
}

func (x *LimitOrder) GetTradesCount() int64 {
	if x != nil {
		return x.TradesCount
	}
	return 0
}

func (x *LimitOrder) GetLimitTrades() []*LimitTrade {
	if x != nil {
		return x.LimitTrades
	}
	return nil
}

var File_limitOrder_proto protoreflect.FileDescriptor

var file_limitOrder_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x1a, 0x10, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x07, 0x0a,
	0x0a, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x55, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x6c, 0x69,
	0x70, 0x70, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x46, 0x6c, 0x69, 0x70,
	0x70, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x55, 0x53, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x55, 0x53, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x55, 0x53, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x53, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x45,
	0x6e, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x45,
	0x6e, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x6e, 0x64, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x55, 0x53, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x45, 0x6e,
	0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x53, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x47, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x47, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x53, 0x72, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x72, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x53, 0x72, 0x63, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x53, 0x72, 0x63, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x74, 0x44, 0x65, 0x63, 0x69, 0x6d,
	0x61, 0x6c, 0x73, 0x18, 0x18, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x44, 0x65, 0x73, 0x74, 0x44,
	0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x61, 0x73, 0x53, 0x65,
	0x6e, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x47, 0x61, 0x73, 0x53, 0x65, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x47, 0x61, 0x73, 0x18, 0x1a, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x47, 0x61, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x53, 0x6c, 0x69, 0x70, 0x70, 0x61, 0x67, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x53, 0x6c, 0x69, 0x70, 0x70, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x64, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a,
	0x0b, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x73, 0x18, 0x20, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x0b, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x73, 0x42, 0x12, 0x5a, 0x10, 0x2f, 0x70, 0x6b, 0x67, 0x2d, 0x62, 0x69, 0x6e, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_limitOrder_proto_rawDescOnce sync.Once
	file_limitOrder_proto_rawDescData = file_limitOrder_proto_rawDesc
)

func file_limitOrder_proto_rawDescGZIP() []byte {
	file_limitOrder_proto_rawDescOnce.Do(func() {
		file_limitOrder_proto_rawDescData = protoimpl.X.CompressGZIP(file_limitOrder_proto_rawDescData)
	})
	return file_limitOrder_proto_rawDescData
}

var file_limitOrder_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_limitOrder_proto_goTypes = []interface{}{
	(*LimitOrder)(nil), // 0: binary.LimitOrder
	(*LimitTrade)(nil), // 1: binary.LimitTrade
}
var file_limitOrder_proto_depIdxs = []int32{
	1, // 0: binary.LimitOrder.LimitTrades:type_name -> binary.LimitTrade
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_limitOrder_proto_init() }
func file_limitOrder_proto_init() {
	if File_limitOrder_proto != nil {
		return
	}
	file_limitTrade_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_limitOrder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitOrder); i {
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
			RawDescriptor: file_limitOrder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_limitOrder_proto_goTypes,
		DependencyIndexes: file_limitOrder_proto_depIdxs,
		MessageInfos:      file_limitOrder_proto_msgTypes,
	}.Build()
	File_limitOrder_proto = out.File
	file_limitOrder_proto_rawDesc = nil
	file_limitOrder_proto_goTypes = nil
	file_limitOrder_proto_depIdxs = nil
}
