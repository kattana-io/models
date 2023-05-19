package models

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
)

type HoldersBlock struct {
	Block     uint64    `json:"block,omitempty"`
	Timestamp int64     `json:"timestamp,omitempty"`
	Chain     string    `json:"chain,omitempty"`
	Notify    bool      `json:"notify,omitempty"`
	Holders   []*Holder `json:"holders,omitempty"`
}

func (h *HoldersBlock) pack() *storage.HoldersBlock {
	holders := make([]*storage.Holder, 0)
	for _, d := range h.Holders {
		holders = append(holders, d.pack())
	}

	return &storage.HoldersBlock{
		Block:     h.Block,
		Timestamp: h.Timestamp,
		Chain:     h.Chain,
		Notify:    h.Notify,
		Holders:   holders,
	}
}

func (h *HoldersBlock) unpack(data *storage.HoldersBlock) *HoldersBlock {
	h.Block = data.Block
	h.Timestamp = data.Timestamp
	h.Chain = data.Chain
	h.Notify = data.Notify

	holders := make([]*Holder, 0)
	for _, d := range data.Holders {
		holders = append(holders, new(Holder).unpack(d))
	}
	h.Holders = holders

	return h
}

func (h *HoldersBlock) MarshalBin() ([]byte, error) {
	return proto.Marshal(h.pack())
}

func (h *HoldersBlock) UnMarshal(buf []byte) error {
	data := &storage.HoldersBlock{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	h.unpack(data)
	return nil
}
