package models

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/proto"
	"math/big"
	"time"
)

type DirectSwap struct {
	Tx          string          `ch:"tx" json:"tx"`
	Date        time.Time       `ch:"date" json:"date"`
	Chain       string          `ch:"chain" json:"chain"`
	BlockNumber uint64          `ch:"blocknumber" json:"blocknumber"`
	Protocol    string          `ch:"protocol" json:"protocol"`
	SrcToken    string          `ch:"src_token" json:"src_token"`
	DstToken    string          `ch:"dst_token" json:"dst_token"`
	Amount0     *big.Int        `ch:"amount0" json:"amount0"`
	Amount1     *big.Int        `ch:"amount1" json:"amount1"`
	PriceA      decimal.Decimal `ch:"pricea" json:"pricea"`
	PriceAUSD   decimal.Decimal `ch:"pricea_usd" json:"pricea_usd"`
	PriceB      decimal.Decimal `ch:"priceb" json:"priceb"`
	PriceBUSD   decimal.Decimal `ch:"priceb_usd" json:"priceb_usd"`
	Wallet      string          `ch:"wallet" json:"wallet"`
	Order       uint16          `ch:"order" json:"order"`
	ValueUSD    decimal.Decimal `ch:"value_usd" json:"value_usd"`
}

func (ds *DirectSwap) Pack() *storage.DirectSwap {
	return &storage.DirectSwap{
		Tx:          ds.Tx,
		Date:        ds.Date.Unix(),
		Chain:       ds.Chain,
		BlockNumber: ds.BlockNumber,
		Protocol:    ds.Protocol,
		SrcToken:    ds.SrcToken,
		DstToken:    ds.DstToken,
		Amount0:     ds.Amount0.Bytes(),
		Amount1:     ds.Amount1.Bytes(),
		PriceA:      ds.PriceA.String(),
		PriceAUSD:   ds.PriceAUSD.String(),
		PriceB:      ds.PriceB.String(),
		PriceBUSD:   ds.PriceBUSD.String(),
		Wallet:      ds.Wallet,
		Order:       uint32(ds.Order),
		ValueUSD:    ds.ValueUSD.String(),
	}
}

func (ds *DirectSwap) Unpack(data *storage.DirectSwap) *DirectSwap {
	ds.Tx = data.Tx
	ds.Date = time.Unix(data.Date, 0)
	ds.Chain = data.Chain
	ds.BlockNumber = data.BlockNumber
	ds.Protocol = data.Protocol
	ds.SrcToken = data.SrcToken
	ds.DstToken = data.DstToken
	ds.Amount0 = new(big.Int).SetBytes(data.Amount0)
	ds.Amount1 = new(big.Int).SetBytes(data.Amount1)

	//TODO change to decimal.FromString()?
	ds.PriceA = decimal.RequireFromString(data.PriceA)
	ds.PriceAUSD = decimal.RequireFromString(data.PriceAUSD)
	ds.PriceB = decimal.RequireFromString(data.PriceB)
	ds.PriceBUSD = decimal.RequireFromString(data.PriceBUSD)
	ds.Wallet = data.Wallet
	ds.Order = uint16(data.Order)
	ds.ValueUSD = decimal.RequireFromString(data.ValueUSD)

	return ds
}

func (ds *DirectSwap) Marshal() ([]byte, error) {
	return proto.Marshal(ds.Pack())
}
func (ds *DirectSwap) UnMarshal(buf []byte) error {
	data := &storage.DirectSwap{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add validate?

	ds.Unpack(data)
	return nil
}
