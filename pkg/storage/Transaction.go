package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
	"math/big"
)

type Transaction struct {
	From  common.Address `json:"from"`
	Hash  common.Hash    `json:"hash"`
	Input []byte         `json:"input"`
	To    common.Address `json:"to"`
	Value *big.Int       `json:"value"`
}

func (b *Transaction) pack() *storage.Transaction {
	return &storage.Transaction{
		From:  b.From.String(),
		Hash:  b.Hash.String(),
		Input: b.Input,
		To:    b.To.String(),
		Value: b.Value.Bytes(),
	}
}

func (b *Transaction) unpack(data *storage.Transaction) *Transaction {
	b.From = common.HexToAddress(data.From)
	b.Hash = common.HexToHash(data.Hash)
	b.Input = data.Input
	b.To = common.HexToAddress(data.To)
	b.Value = new(big.Int).SetBytes(data.Value)

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
