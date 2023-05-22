package models

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
)

type BlockState struct {
	DirectSwaps []*DirectSwap     `json:"direct_swaps"`
	PairSwaps   []*PairSwap       `json:"pair_swaps"`
	Liquidities []*LiquidityEvent `json:"liquidity_events"`
	Transfers   []*TransferEvent  `json:"transfer_events"`
	Pairs       []*NewPair        `json:"new_pairs"`
	Holders     []*Holder         `json:"holders"`
	DexeEvents  []*CustomEvent    `json:"dexe_events"`
	Block       *Block            `json:"block"`
}

func (b *BlockState) pack() *storage.BlockState {
	directSwaps := make([]*storage.DirectSwap, 0)
	for _, b := range b.DirectSwaps {
		directSwaps = append(directSwaps, b.Pack())
	}
	pairSwaps := make([]*storage.PairSwap, 0)
	for _, b := range b.PairSwaps {
		pairSwaps = append(pairSwaps, b.Pack())
	}
	liquidities := make([]*storage.LiquidityEvent, 0)
	for _, l := range b.Liquidities {
		liquidities = append(liquidities, l.pack())
	}
	transfers := make([]*storage.TransferEvent, 0)
	for _, b := range b.Transfers {
		transfers = append(transfers, b.pack())
	}
	pairs := make([]*storage.NewPair, 0)
	for _, b := range b.Pairs {
		pairs = append(pairs, b.pack())
	}
	holders := make([]*storage.Holder, 0)
	for _, h := range b.Holders {
		holders = append(holders, h.pack())
	}
	dexeEvents := make([]*storage.CustomEvent, 0)
	for _, e := range b.DexeEvents {
		dexeEvents = append(dexeEvents, e.pack())
	}

	return &storage.BlockState{
		DirectSwaps: directSwaps,
		PairSwaps:   pairSwaps,
		Liquidities: liquidities,
		Transfers:   transfers,
		Pairs:       pairs,
		Holders:     holders,
		DexeEvents:  dexeEvents,
		Block:       b.Block.pack(),
	}
}

func (b *BlockState) unpack(data *storage.BlockState) {
	directSwaps := make([]*DirectSwap, 0)
	for _, b := range data.DirectSwaps {
		directSwaps = append(directSwaps, new(DirectSwap).Unpack(b))
	}
	b.DirectSwaps = directSwaps

	pairSwaps := make([]*PairSwap, 0)
	for _, b := range data.PairSwaps {
		pairSwaps = append(pairSwaps, new(PairSwap).Unpack(b))
	}
	b.PairSwaps = pairSwaps

	liquidities := make([]*LiquidityEvent, 0)
	for _, b := range data.Liquidities {
		liquidities = append(liquidities, new(LiquidityEvent).unpack(b))
	}
	b.Liquidities = liquidities

	transfers := make([]*TransferEvent, 0)
	for _, b := range data.Transfers {
		transfers = append(transfers, new(TransferEvent).unpack(b))
	}
	b.Transfers = transfers

	pairs := make([]*NewPair, 0)
	for _, b := range data.Pairs {
		pairs = append(pairs, new(NewPair).unpack(b))
	}
	b.Pairs = pairs
	b.Block = new(Block).unpack(data.Block)
}

func (b *BlockState) Marshal() ([]byte, error) {
	return proto.Marshal(b.pack())
}

func (b *BlockState) UnMarshal(buf []byte) error {
	data := &storage.BlockState{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	b.unpack(data)
	return nil
}
