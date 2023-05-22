package models

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/proto"
)

type PriceHolder struct {
	Pair      string          `json:"pair"`
	Chain     string          `json:"chain"`
	PriceA    decimal.Decimal `json:"priceA"`
	PriceAUSD decimal.Decimal `json:"PriceAUSD"`
	PriceB    decimal.Decimal `json:"priceB"`
	PriceBUSD decimal.Decimal `json:"priceBUSD"`
}

func (h *PriceHolder) Pack() *storage.PriceHolder {
	return &storage.PriceHolder{
		Pair:      h.Pair,
		Chain:     h.Chain,
		PriceA:    h.PriceA.String(),
		PriceAUSD: h.PriceAUSD.String(),
		PriceB:    h.PriceB.String(),
		PriceBUSD: h.PriceBUSD.String(),
	}
}

func (h *PriceHolder) Unpack(data *storage.PriceHolder) *PriceHolder {
	h.Pair = data.Pair
	h.Chain = data.Chain
	h.PriceA, _ = decimal.NewFromString(data.PriceA)
	h.PriceAUSD, _ = decimal.NewFromString(data.PriceAUSD)
	h.PriceB, _ = decimal.NewFromString(data.PriceB)
	h.PriceBUSD, _ = decimal.NewFromString(data.PriceBUSD)

	return h
}

type PriceUpdates struct {
	Items []*PriceHolder `json:"price_updates"`
}

func (u *PriceUpdates) pack() *storage.PriceUpdates {
	items := make([]*storage.PriceHolder, 0)
	for _, i := range u.Items {
		items = append(items, i.Pack())
	}

	return &storage.PriceUpdates{
		Items: items,
	}
}

func (u *PriceUpdates) unpack(data *storage.PriceUpdates) *PriceUpdates {
	items := make([]*PriceHolder, 0)
	for _, b := range data.Items {
		items = append(items, new(PriceHolder).Unpack(b))
	}
	u.Items = items

	return u
}

func (u *PriceUpdates) Marshal() ([]byte, error) {
	return proto.Marshal(u.pack())
}

func (u *PriceUpdates) UnMarshal(buf []byte) error {
	data := &storage.PriceUpdates{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	u.unpack(data)
	return nil
}
