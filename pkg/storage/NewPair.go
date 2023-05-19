package models

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
)

type NewPair struct {
	Factory     string `json:"factory"`
	Pair        string `json:"pair"`
	Klass       string `json:"klass"`
	Network     string `json:"network"`
	Node        string `json:"node"`
	PoolCreated int64  `json:"pool_created"`
}

func (p *NewPair) pack() *storage.NewPair {
	return &storage.NewPair{
		Factory:     p.Factory,
		Pair:        p.Pair,
		Klass:       p.Klass,
		Network:     p.Network,
		Node:        p.Node,
		PoolCreated: p.PoolCreated,
	}
}

func (p *NewPair) unpack(data *storage.NewPair) *NewPair {
	p.Factory = data.Factory
	p.Pair = data.Pair
	p.Klass = data.Klass
	p.Network = data.Network
	p.Node = data.Node
	p.PoolCreated = data.PoolCreated

	return p
}

func (p *NewPair) Marshal() ([]byte, error) {
	return proto.Marshal(p.pack())
}

func (p *NewPair) UnMarshal(buf []byte) error {
	data := &storage.NewPair{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	p.unpack(data)
	return nil
}
