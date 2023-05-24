package models

import (
	"math/big"

	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
)

type Trade struct {
	Pair      string   `json:"pair"`
	PriceA    float64  `json:"price_a"`
	PriceB    float64  `json:"price_b"`
	PriceAUSD float64  `json:"price_a_usd"`
	PriceBUSD float64  `json:"price_b_usd"`
	VolumeA   *big.Int `json:"volume_a"`
	VolumeB   *big.Int `json:"volume_b"`
	VolumeUSD float64  `json:"volume_usd"`
	IsBuy     bool     `json:"is_buy"`
}

func (t *Trade) Pack() *storage.Trade {
	return &storage.Trade{
		Pair:      t.Pair,
		PriceA:    t.PriceA,
		PriceB:    t.PriceB,
		PriceAUSD: t.PriceAUSD,
		PriceBUSD: t.PriceBUSD,
		VolumeA:   t.VolumeA.Bytes(),
		VolumeB:   t.VolumeB.Bytes(),
		VolumeUSD: t.VolumeUSD,
		IsBuy:     t.IsBuy,
	}
}

func (t *Trade) Unpack(data *storage.Trade) *Trade {
	t.Pair = data.Pair
	t.PriceA = data.PriceA
	t.PriceB = data.PriceB
	t.PriceAUSD = data.PriceAUSD
	t.PriceBUSD = data.PriceBUSD
	t.VolumeA = big.NewInt(0).SetBytes(data.VolumeA)
	t.VolumeB = big.NewInt(0).SetBytes(data.VolumeB)
	t.VolumeUSD = data.VolumeUSD
	t.IsBuy = data.IsBuy

	return t
}

func (t *Trade) Marshal() ([]byte, error) {
	return proto.Marshal(t.Pack())
}

func (t *Trade) UnMarshal(buf []byte) error {
	data := &storage.Trade{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	t.Unpack(data)
	return nil
}
