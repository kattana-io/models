package models

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
	"math/big"
)

type DexeEvents struct {
	Number       *big.Int       `json:"number,omitempty"`
	Timestamp    int64          `json:"timestamp,omitempty"`
	Network      string         `json:"network,omitempty"`
	CustomEvents []*CustomEvent `json:"customEvents,omitempty"`
}

func (d *DexeEvents) pack() *storage.DexeEvents {
	customEvents := make([]*storage.CustomEvent, 0)
	for _, c := range d.CustomEvents {
		customEvents = append(customEvents, c.pack())
	}

	return &storage.DexeEvents{
		Number:       d.Number.Bytes(),
		Timestamp:    d.Timestamp,
		Network:      d.Network,
		CustomEvents: customEvents,
	}
}

func (d *DexeEvents) unpack(data *storage.DexeEvents) {
	d.Number = new(big.Int).SetBytes(data.Number)
	d.Timestamp = data.Timestamp
	d.Network = data.Network
	customEvents := make([]*CustomEvent, 0)
	for _, c := range data.CustomEvents {
		customEvents = append(customEvents, new(CustomEvent).unpack(c))
	}
	d.CustomEvents = customEvents
}

func (d *DexeEvents) Marshal() ([]byte, error) {
	return proto.Marshal(d.pack())
}

func (d *DexeEvents) UnMarshal(buf []byte) error {
	data := &storage.DexeEvents{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	d.unpack(data)
	return nil
}
