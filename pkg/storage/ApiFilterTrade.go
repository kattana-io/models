package models

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
)

type Filter struct {
	PriceMin      float32
	PriceMax      float32
	PriceUSDMin   float32
	PriceUSDMax   float32
	VolumeMin     float32
	VolumeMax     float32
	VolumeSellMin float32
	VolumeSellMfx float32
	BaseMin       string
	BaseMax       string
	QuoteMin      string
	QuoteMax      string
	Wallet        []string
	Buy           bool
	Sell          bool
	Add           bool
	Del           bool
	Size          storage.TradeSize
}

func (t *Filter) Pack() *storage.TradeFilter {
	return &storage.TradeFilter{
		PriceMin:      t.PriceMin,
		PriceMax:      t.PriceMax,
		PriceUSDMin:   t.PriceUSDMin,
		PriceUSDMax:   t.PriceUSDMax,
		VolumeMin:     t.VolumeMin,
		VolumeMax:     t.VolumeMax,
		VolumeSellMin: t.VolumeSellMin,
		VolumeSellMfx: t.VolumeSellMfx,
		BaseMin:       t.BaseMin,
		BaseMax:       t.BaseMax,
		QuoteMin:      t.QuoteMin,
		QuoteMax:      t.QuoteMax,
		Wallet:        t.Wallet,
		Buy:           t.Buy,
		Sell:          t.Sell,
		Add:           t.Add,
		Del:           t.Del,
		Size:          t.Size,
	}
}

func (t *Filter) Unpack(data *storage.TradeFilter) *Filter {

	t.PriceMin = data.PriceMin
	t.PriceMax = data.PriceMax
	t.PriceUSDMin = data.PriceUSDMin
	t.PriceUSDMax = data.PriceUSDMax
	t.VolumeMin = data.VolumeMin
	t.VolumeMax = data.VolumeMax
	t.VolumeSellMin = data.VolumeSellMin
	t.VolumeSellMfx = data.VolumeSellMfx
	t.BaseMin = data.BaseMin
	t.BaseMax = data.BaseMax
	t.QuoteMin = data.QuoteMin
	t.QuoteMax = data.QuoteMax
	t.Wallet = data.Wallet
	t.Buy = data.Buy
	t.Sell = data.Sell
	t.Add = data.Add
	t.Del = data.Del
	t.Size = data.Size
	return t
}

func (t *Filter) Marshal() ([]byte, error) {
	return proto.Marshal(t.Pack())
}

func (t *Filter) UnMarshal(buf []byte) error {
	data := &storage.TradeFilter{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	t.Unpack(data)
	return nil
}
