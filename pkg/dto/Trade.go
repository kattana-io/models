package dto

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
	"math/big"
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
}

func (t *Trade) pack() *storage.Trade {
	return &storage.Trade{
		Pair:      t.Pair,
		PriceA:    t.PriceA,
		PriceB:    t.PriceB,
		PriceAUSD: t.PriceAUSD,
		PriceBUSD: t.PriceBUSD,
		VolumeA:   t.VolumeA.Bytes(),
		VolumeB:   t.VolumeB.Bytes(),
		VolumeUSD: t.VolumeUSD,
	}
}

func (t *Trade) unpack(data *storage.Trade) *Trade {
	t.Pair = data.Pair
	t.PriceA = data.PriceA
	t.PriceB = data.PriceB
	t.PriceAUSD = data.PriceAUSD
	t.PriceBUSD = data.PriceBUSD
	t.VolumeA = new(big.Int).SetBytes(data.VolumeA)
	t.VolumeB = new(big.Int).SetBytes(data.VolumeB)
	t.VolumeUSD = data.VolumeUSD
	return t
}

func (t *Trade) Marshal() ([]byte, error) {
	return proto.Marshal(t.pack())
}

func (t *Trade) UnMarshal(buf []byte) error {
	data := &storage.Trade{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	t.unpack(data)
	return nil
}
