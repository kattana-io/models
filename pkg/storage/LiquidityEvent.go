package models

import (
	"time"

	"github.com/kattana-io/models/pkg-bin/storage"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/proto"
)

type LiquidityEvent struct {
	// generic info
	BlockNumber uint64    `json:"blocknumber"`
	Date        time.Time `json:"date"`
	Tx          string    `json:"tx"`
	Pair        string    `json:"pair"`
	Chain       string    `json:"chain"`
	Klass       string    `json:"klass"` // mint, burn, sync
	Wallet      string    `json:"wallet"`
	Order       uint      `json:"order"`
	// Mint, Burn
	Amount0 string `json:"amount0"`
	Amount1 string `json:"amount1"`
	// Sync
	Reserve0 string `json:"reserve0"`
	Reserve1 string `json:"reserve1"`
	// Statistics part
	PriceA    decimal.Decimal `json:"pricea"`
	PriceAUSD decimal.Decimal `json:"pricea_usd"`
	PriceB    decimal.Decimal `json:"priceb"`
	PriceBUSD decimal.Decimal `json:"priceb_usd"`
	// USD value of operation
	ValueUSD decimal.Decimal `json:"value_usd"`
	// Reserves value in USD
	ReserveUSD decimal.Decimal `json:"reserve_usd"`
	ValidTill  time.Time       `json:"valid_till"`

	Source int32 `json:"source"`
}

func (l *LiquidityEvent) pack() *storage.LiquidityEvent {
	return &storage.LiquidityEvent{
		BlockNumber: l.BlockNumber,
		Date:        l.Date.Unix(),
		Tx:          l.Tx,
		Pair:        l.Pair,
		Chain:       l.Chain,
		Klass:       l.Klass,
		Wallet:      l.Wallet,
		Order:       uint32(l.Order),
		Amount0:     l.Amount0,
		Amount1:     l.Amount1,
		Reserve0:    l.Reserve0,
		Reserve1:    l.Reserve1,
		PriceA:      l.PriceA.String(),
		PriceAUSD:   l.PriceAUSD.String(),
		PriceB:      l.PriceB.String(),
		PriceBUSD:   l.PriceBUSD.String(),
		ValueUSD:    l.ValueUSD.String(),
		ReserveUSD:  l.ReserveUSD.String(),
		ValidTill:   l.ValidTill.Unix(),
		Source:      storage.LiquidityType(l.Source),
	}
}

func (l *LiquidityEvent) unpack(data *storage.LiquidityEvent) *LiquidityEvent {
	l.BlockNumber = data.BlockNumber
	l.Date = time.Unix(data.Date, 0)
	l.Tx = data.Tx
	l.Pair = data.Pair
	l.Chain = data.Chain
	l.Klass = data.Klass
	l.Wallet = data.Wallet
	l.Order = uint(data.Order)

	l.Amount0 = data.Amount0
	l.Amount1 = data.Amount1
	l.Reserve0 = data.Reserve0
	l.Reserve1 = data.Reserve1

	l.PriceA = decimal.RequireFromString(data.PriceA)
	l.PriceAUSD = decimal.RequireFromString(data.PriceAUSD)
	l.PriceB = decimal.RequireFromString(data.PriceB)
	l.PriceBUSD = decimal.RequireFromString(data.PriceBUSD)
	l.ValueUSD = decimal.RequireFromString(data.ValueUSD)
	l.ReserveUSD = decimal.RequireFromString(data.ReserveUSD)

	l.ValidTill = time.Unix(data.ValidTill, 0)
	l.Source = int32(data.Source)

	return l
}

func (l *LiquidityEvent) Marshal() ([]byte, error) {
	return proto.Marshal(l.pack())
}

func (l *LiquidityEvent) UnMarshal(buf []byte) error {
	data := &storage.LiquidityEvent{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	l.unpack(data)
	return nil
}
