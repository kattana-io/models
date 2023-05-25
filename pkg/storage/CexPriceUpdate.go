package models

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
)

type CexPriceUpdate struct {
	ID          uint    `json:"-"`
	Price       string  `json:"price"`
	BaseVolume  float64 `json:"base_volume"`
	QuoteVolume float64 `json:"quote_Volume"`
}

func (c *CexPriceUpdate) Pack() *storage.CexPriceUpdate {
	return &storage.CexPriceUpdate{
		Price:       c.Price,
		BaseVolume:  c.BaseVolume,
		QuoteVolume: c.QuoteVolume,
	}
}

func (c *CexPriceUpdate) Unpack(data *storage.CexPriceUpdate) *CexPriceUpdate {
	c.Price = data.Price
	c.BaseVolume = data.BaseVolume
	c.QuoteVolume = data.QuoteVolume

	return c
}

func (c *CexPriceUpdate) Marshal() ([]byte, error) {
	return proto.Marshal(c.Pack())
}

func (c *CexPriceUpdate) UnMarshal(buf []byte) error {
	data := &storage.CexPriceUpdate{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	c.Unpack(data)
	return nil
}
