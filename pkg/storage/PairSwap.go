package models

import (
	"math/big"
	"time"

	"github.com/kattana-io/models/pkg-bin/storage"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/proto"
)

type PairSwap struct {
	Tx               string          `json:"tx"`
	Date             time.Time       `json:"date"`
	Chain            string          `json:"chain"`
	BlockNumber      uint64          `json:"blocknumber"`
	Pair             string          `json:"pair"`
	Amount0          *big.Int        `json:"amount0"`
	Amount1          *big.Int        `json:"amount1"`
	Buy              bool            `json:"buy"`
	PriceA           decimal.Decimal `json:"pricea"`
	PriceAUSD        decimal.Decimal `json:"pricea_usd"`
	PriceB           decimal.Decimal `json:"priceb"`
	PriceBUSD        decimal.Decimal `json:"priceb_usd"`
	Bot              bool            `json:"bot"`
	Wallet           string          `json:"wallet"`
	Order            uint16          `json:"order"`
	ValueUSD         decimal.Decimal `json:"value_usd"`
	TotalTradesCount uint32          `json:"total_trades_count"`
	PairTradesCount  uint32          `json:"pair_trades_count"`
	Klass            string          `json:"klass"`
	Reserve0         string          `json:"reverse0"`
	Reserve1         string          `json:"reverse1"`
	ReserveUSD       string          `json:"reverse_usd"`
	ValidTill        int64           `json:"valid_till"`
	SqrtPriceX96     *big.Int        `json:"sqrt_price_x96"`
	Liquidity        *big.Int        `json:"liquidity"`
	Tick             *big.Int        `json:"tick"`
}

func (p *PairSwap) Pack() *storage.PairSwap {
	return &storage.PairSwap{
		Tx:               p.Tx,
		Date:             p.Date.Unix(),
		Chain:            p.Chain,
		BlockNumber:      p.BlockNumber,
		Pair:             p.Pair,
		Amount0:          p.Amount0.Bytes(),
		Amount1:          p.Amount1.Bytes(),
		Buy:              p.Buy,
		PriceA:           p.PriceA.String(),
		PriceAUSD:        p.PriceAUSD.String(),
		PriceB:           p.PriceB.String(),
		PriceBUSD:        p.PriceBUSD.String(),
		Bot:              p.Bot,
		Wallet:           p.Wallet,
		Order:            uint32(p.Order),
		ValueUSD:         p.ValueUSD.String(),
		TotalTradesCount: p.TotalTradesCount,
		PairTradesCount:  p.PairTradesCount,
		Klass:            p.Klass,
		Reserve0:         p.Reserve0,
		Reserve1:         p.Reserve1,
		ReserveUSD:       p.ReserveUSD,
		ValidTill:        p.ValidTill,
		SqrtPriceX96:     p.SqrtPriceX96.Bytes(),
		Liquidity:        p.Liquidity.Bytes(),
		Tick:             p.Tick.Bytes(),
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
	p.ValueUSD = decimal.RequireFromString(data.ValueUSD)
	p.TotalTradesCount = data.TotalTradesCount
	p.PairTradesCount = data.PairTradesCount
	p.Klass = data.Klass
	p.Reserve0 = data.Reserve0
	p.Reserve1 = data.Reserve1
	p.ReserveUSD = data.ReserveUSD
	p.ValidTill = data.ValidTill
	p.SqrtPriceX96 = new(big.Int).SetBytes(data.SqrtPriceX96)
	p.Liquidity = new(big.Int).SetBytes(data.Liquidity)
	p.Tick = new(big.Int).SetBytes(data.Tick)

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
