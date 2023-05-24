package dto

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
)

type Candle struct {
	Open  float64 `json:"open"`
	Low   float64 `json:"low"`
	High  float64 `json:"high"`
	Date  int32   `json:"date"`
	Close float64 `json:"close"`
}

type Candles struct {
	Candle []*Candle `json:"c"`
	Volume *Volume   `json:"v"`
}

func (c *Candle) pack() *storage.Candle {
	return &storage.Candle{
		Open:  c.Open,
		Low:   c.Low,
		High:  c.High,
		Date:  c.Date,
		Close: c.Close,
	}
}

func (c *Candle) unpack(data *storage.Candle) *Candle {
	c.Open = data.Open
	c.Low = data.Low
	c.High = data.High
	c.Date = data.Date
	c.Close = data.Close
	return c
}

func (c *Candle) Marshal() ([]byte, error) {
	return proto.Marshal(c.pack())
}

func (c *Candle) UnMarshal(buf []byte) error {
	data := &storage.Candle{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	c.unpack(data)
	return nil
}

func (cs *Candles) pack() *storage.Candles {
	candles := make([]*storage.Candle, 0)
	for _, c := range cs.Candle {
		candles = append(candles, c.pack())
	}
	return &storage.Candles{
		Candle: candles,
		Volume: cs.Volume.pack(),
	}
}

func (cs *Candles) unpack(data *storage.Candles) *Candles {
	candle := make([]*Candle, 0)
	for _, c := range data.Candle {
		candle = append(candle, new(Candle).unpack(c))
	}
	cs.Candle = candle
	cs.Volume = new(Volume).unpack(data.Volume)
	return cs
}

func (cs *Candles) Marshal() ([]byte, error) {
	return proto.Marshal(cs.pack())
}

func (cs *Candles) UnMarshal(buf []byte) error {
	data := &storage.Candles{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	cs.unpack(data)
	return nil
}
