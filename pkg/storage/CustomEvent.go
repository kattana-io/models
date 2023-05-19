package models

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
)

type CustomEvent struct {
	Tx       string `json:"tx,omitempty"`
	Contract string `json:"contract,omitempty"`
	Wallet   string `json:"wallet,omitempty"`
	Entity   string `json:"entity,omitempty"`
	Klass    string `json:"klass,omitempty"`
	Name     string `json:"name,omitempty"`
	Data     string `json:"data,omitempty"`
}

func (c *CustomEvent) pack() *storage.CustomEvent {
	return &storage.CustomEvent{
		Tx:       c.Tx,
		Contract: c.Contract,
		Wallet:   c.Wallet,
		Entity:   c.Entity,
		Klass:    c.Klass,
		Name:     c.Name,
		Data:     c.Data,
	}
}

func (c *CustomEvent) unpack(data *storage.CustomEvent) *CustomEvent {
	c.Tx = data.Tx
	c.Contract = data.Contract
	c.Wallet = data.Wallet
	c.Entity = data.Entity
	c.Klass = data.Klass
	c.Name = data.Name
	c.Data = data.Data
	return c
}
func (c *CustomEvent) Marshal() ([]byte, error) {
	return proto.Marshal(c.pack())
}

func (c *CustomEvent) UnMarshal(buf []byte) error {
	data := &storage.CustomEvent{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	c.unpack(data)
	return nil
}
