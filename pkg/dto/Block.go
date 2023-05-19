package dto

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	models "github.com/kattana-io/models/pkg/storage"
	"google.golang.org/protobuf/proto"
)

type Block struct {
	Network     string               `json:"network"`
	Date        uint64               `json:"date"`
	Trades      []*Trade             `json:"trades"`
	PairSwaps   []*models.PairSwap   `json:"pair_swaps"`
	DirectSwaps []*models.DirectSwap `json:"direct_swaps"`
}

func (b *Block) pack() *storage.BlockLite {
	trades := make([]*storage.Trade, 0)
	for _, b := range b.Trades {
		trades = append(trades, b.pack())
	}

	directSwaps := make([]*storage.DirectSwap, 0)
	for _, b := range b.DirectSwaps {
		directSwaps = append(directSwaps, b.Pack())
	}
	pairSwaps := make([]*storage.PairSwap, 0)
	for _, b := range b.PairSwaps {
		pairSwaps = append(pairSwaps, b.Pack())
	}

	return &storage.BlockLite{
		Network:     b.Network,
		Date:        b.Date,
		Trades:      trades,
		PairSwaps:   pairSwaps,
		DirectSwaps: directSwaps,
	}
}

func (b *Block) unpack(data *storage.BlockLite) *Block {
	b.Network = data.Network
	b.Date = data.Date

	trades := make([]*Trade, 0)
	for _, b := range data.Trades {
		trades = append(trades, new(Trade).unpack(b))
	}
	b.Trades = trades

	pairSwaps := make([]*models.PairSwap, 0)
	for _, b := range data.PairSwaps {
		pairSwaps = append(pairSwaps, new(models.PairSwap).Unpack(b))
	}
	b.PairSwaps = pairSwaps
	directSwaps := make([]*models.DirectSwap, 0)
	for _, b := range data.DirectSwaps {
		directSwaps = append(directSwaps, new(models.DirectSwap).Unpack(b))
	}
	b.DirectSwaps = directSwaps
	return b
}

func (b *Block) Marshal() ([]byte, error) {
	return proto.Marshal(b.pack())
}

func (b *Block) UnMarshal(buf []byte) error {
	data := &storage.BlockLite{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	b.unpack(data)
	return nil
}
