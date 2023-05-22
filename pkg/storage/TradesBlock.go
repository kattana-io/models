package models

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
)

type TradesBlock struct {
	Network     string        `json:"network"`
	Date        uint64        `json:"date"`
	Trades      []*Trade      `json:"trades"`
	PairSwaps   []*PairSwap   `json:"pair_swaps"`
	DirectSwaps []*DirectSwap `json:"direct_swaps"`
}

func (t *TradesBlock) pack() *storage.TradesBlock {
	directSwaps := make([]*storage.DirectSwap, 0)
	for _, b := range t.DirectSwaps {
		directSwaps = append(directSwaps, b.Pack())
	}
	pairSwaps := make([]*storage.PairSwap, 0)
	for _, b := range t.PairSwaps {
		pairSwaps = append(pairSwaps, b.Pack())
	}
	trades := make([]*storage.Trade, 0)
	for _, t := range t.Trades {
		trades = append(trades, t.Pack())
	}

	return &storage.TradesBlock{
		Network: t.Network,
		Date:    t.Date,
		Trades:  trades,
	}
}

func (t *TradesBlock) unpack(data *storage.TradesBlock) *TradesBlock {
	directSwaps := make([]*DirectSwap, 0)
	for _, b := range data.DirectSwaps {
		directSwaps = append(directSwaps, new(DirectSwap).Unpack(b))
	}
	pairSwaps := make([]*PairSwap, 0)
	for _, b := range data.PairSwaps {
		pairSwaps = append(pairSwaps, new(PairSwap).Unpack(b))
	}
	trades := make([]*Trade, 0)
	for _, t := range data.Trades {
		trades = append(trades, new(Trade).Unpack(t))
	}

	t.Network = data.Network
	t.Date = data.Date
	t.PairSwaps = pairSwaps
	t.DirectSwaps = directSwaps
	t.Trades = trades

	return t
}

func (t *TradesBlock) Marshal() ([]byte, error) {
	return proto.Marshal(t.pack())
}

func (t *TradesBlock) UnMarshal(buf []byte) error {
	data := &storage.TradesBlock{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	t.unpack(data)
	return nil
}
