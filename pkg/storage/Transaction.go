package models

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
)

type Transaction struct {
	From  string `json:"from"`
	Hash  string `json:"hash"`
	Input string `json:"input"`
	To    string `json:"to"`
	Value string `json:"value"`
}

func (b *Transaction) pack() *storage.Transaction {
	return &storage.Transaction{
		From:  b.From,
		Hash:  b.Hash,
		Input: b.Input,
		To:    b.To,
		Value: b.Value,
	}
}

func (b *Transaction) unpack(data *storage.Transaction) *Transaction {
	b.From = data.From
	b.Hash = data.Hash
	b.Input = data.Input
	b.To = data.To
	b.Value = data.Value

	return b
}

func (b *Transaction) Marshal() ([]byte, error) {
	return proto.Marshal(b.pack())
}

func (b *Transaction) UnMarshal(buf []byte) error {
	transaction := &storage.Transaction{}
	err := proto.Unmarshal(buf, transaction)
	if err != nil {
		return err
	}

	b.unpack(transaction)
	return nil
}
