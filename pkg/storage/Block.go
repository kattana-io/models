package models

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
	"math/big"
)

type Block struct {
	Type         string         `json:"type"`
	Network      string         `json:"network"`
	Number       *big.Int       `json:"number"`
	Node         string         `json:"node"`
	Notify       bool           `json:"notify"`
	Timestamp    uint64         `json:"timestamp"`
	Empty        bool           `json:"empty"`
	Transactions []*Transaction `json:"transactions,omitempty"`
	Logs         []types.Log    `json:"logs,omitempty"`
	Part         uint8          `json:"-"`

	// block trace
	DiscoveredAt       int64 `json:"discovered_at"`
	ReceivedInParserAt int64 `json:"received_in_parser_at"`
	ParsedAt           int64 `json:"parsed_at"`
	ConsumedAt         int64 `json:"consumed_at"`
	LiveSendAt         int64 `json:"live_send_at"`
}

func (b *Block) pack() *storage.Block {
	transactions := make([]*storage.Transaction, 0)
	for _, t := range b.Transactions {
		transactions = append(transactions, t.pack())
	}
	logs := make([][]byte, 0)
	for _, l := range b.Logs {
		log, err := l.MarshalJSON()
		if err != nil {
			continue
		}
		logs = append(logs, log)
	}

	return &storage.Block{
		Type:               b.Type,
		Network:            b.Network,
		Number:             b.Number.Bytes(),
		Node:               b.Node,
		Notify:             b.Notify,
		Timestamp:          b.Timestamp,
		Empty:              b.Empty,
		DiscoveredAt:       b.DiscoveredAt,
		ReceivedInParserAt: b.ReceivedInParserAt,
		ParsedAt:           b.ParsedAt,
		ConsumedAt:         b.ConsumedAt,
		LiveSendAt:         b.LiveSendAt,
		Transactions:       transactions,
		Logs:               logs,
	}
}

func (b *Block) unpack(data *storage.Block) *Block {
	b.Type = data.Type
	b.Network = data.Network
	b.Number = new(big.Int).SetBytes(data.Number)
	b.Node = data.Node
	b.Notify = data.Notify
	b.Timestamp = data.Timestamp
	b.Empty = data.Empty
	b.DiscoveredAt = data.DiscoveredAt
	b.ReceivedInParserAt = data.ReceivedInParserAt
	b.ParsedAt = data.ParsedAt
	b.ConsumedAt = data.ConsumedAt
	b.LiveSendAt = data.LiveSendAt

	transactions := make([]*Transaction, 0)
	for _, t := range data.Transactions {
		transactions = append(transactions, new(Transaction).unpack(t))
	}
	b.Transactions = transactions

	logs := make([]types.Log, 0)
	for _, l := range data.Logs {
		log := new(types.Log)
		err := log.UnmarshalJSON(l)
		if err != nil {
			continue
		}
		logs = append(logs, *log)
	}
	b.Logs = logs

	return b
}

func (b *Block) Marshal() ([]byte, error) {
	return proto.Marshal(b.pack())
}

func (b *Block) UnMarshal(buf []byte) error {
	block := &storage.Block{}
	err := proto.Unmarshal(buf, block)
	if err != nil {
		return err
	}
	// TODO add block validate

	b.unpack(block)
	return nil
}
