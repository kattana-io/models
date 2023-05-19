package models

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
)

type Holder struct {
	Token  string `json:"token"`
	From   string `json:"from"`
	To     string `json:"to"`
	Tx     string `json:"tx"`
	Amount int64  `json:"amount,omitempty"`
}

func (h *Holder) pack() *storage.Holder {
	return &storage.Holder{
		Token:  h.Token,
		From:   h.From,
		To:     h.To,
		Tx:     h.Tx,
		Amount: h.Amount,
	}
}

func (h *Holder) unpack(data *storage.Holder) *Holder {
	h.Token = data.Token
	h.From = data.From
	h.To = data.To
	h.Tx = data.Tx
	h.Amount = data.Amount

	return h
}

func (h *Holder) Marshal() ([]byte, error) {
	return proto.Marshal(h.pack())
}

func (h *Holder) UnMarshal(buf []byte) error {
	data := &storage.Holder{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	h.unpack(data)
	return nil
}
