package models

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/proto"
	"math/big"
	"time"
)

type PairSwap struct {
	Tx          string          `json:"tx"`
	Date        time.Time       `json:"date"`
	Chain       string          `json:"chain"`
	BlockNumber uint64          `json:"blocknumber"`
	Pair        string          `json:"pair"`
	Amount0     *big.Int        `json:"amount0"`
	Amount1     *big.Int        `json:"amount1"`
	Buy         bool            `json:"buy"`
	PriceA      decimal.Decimal `json:"pricea"`
	PriceAUSD   decimal.Decimal `json:"pricea_usd"`
	PriceB      decimal.Decimal `json:"priceb"`
	PriceBUSD   decimal.Decimal `json:"priceb_usd"`
	Bot         bool            `json:"bot"`
	Wallet      string          `json:"wallet"`
	Order       uint16          `json:"order"`
	ValueUSD    decimal.Decimal `json:"value_usd"`
}

func (p *PairSwap) Pack() *storage.PairSwap {
	return &storage.PairSwap{
		Tx:          p.Tx,
		Date:        p.Date.Unix(),
		Chain:       p.Chain,
		BlockNumber: p.BlockNumber,
		Pair:        p.Pair,
		Amount0:     p.Amount0.Bytes(),
		Amount1:     p.Amount1.Bytes(),
		Buy:         p.Buy,
		PriceA:      p.PriceA.String(),
		PriceAUSD:   p.PriceAUSD.String(),
		PriceB:      p.PriceB.String(),
		PriceBUSD:   p.PriceBUSD.String(),
		Bot:         p.Bot,
		Wallet:      p.Wallet,
		Order:       uint32(p.Order),
		ValueUSD:    p.ValueUSD,
	}
}

func (p *PairSwap) Unpack(data *storage.PairSwap) *PairSwap {
	p.Tx = data.Tx
	p.Date = time.Unix(data.Date, 0)
	p.Chain = data.Chain
	p.BlockNumber = data.BlockNumber
	p.Pair = data.Pair
	p.Amount0 = new(big.Int).SetBytes(data.Amount0)
	p.Amount1 = new(big.Int).SetBytes(data.Amount1)
	p.Buy = data.Buy
	p.PriceA = decimal.RequireFromString(data.PriceA)
	p.PriceAUSD = decimal.RequireFromString(data.PriceAUSD)
	p.PriceB = decimal.RequireFromString(data.PriceB)
	p.PriceBUSD = decimal.RequireFromString(data.PriceBUSD)
	p.Bot = data.Bot
	p.Wallet = data.Wallet
	p.Order = uint16(data.Order)
	p.ValueUSD = data.ValueUSD

	return p
}

func (p *PairSwap) Marshal() ([]byte, error) {
	return proto.Marshal(p.Pack())
}

func (p *PairSwap) UnMarshal(buf []byte) error {
	data := &storage.PairSwap{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	p.Unpack(data)
	return nil
}
