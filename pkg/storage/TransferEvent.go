package models

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/proto"
	"time"
)

type TransferEvent struct {
	Chain       string          `json:"chain"`
	Contract    string          `json:"contract"`
	BlockNumber uint64          `json:"blocknumber"`
	Date        time.Time       `json:"date"`
	Order       uint            `json:"order"` // log order
	Tx          string          `json:"tx"`
	From        string          `json:"from"`   // usually contract
	To          string          `json:"to"`     // usually wallet
	Amount      string          `json:"amount"` // bigint amount
	ValueUSD    decimal.Decimal `json:"value_usd"`
}

func (t *TransferEvent) pack() *storage.TransferEvent {
	return &storage.TransferEvent{
		Chain:       t.Chain,
		Contract:    t.Contract,
		BlockNumber: t.BlockNumber,
		Date:        t.Date.Unix(),
		Order:       uint32(t.Order),
		Tx:          t.Tx,
		From:        t.From,
		To:          t.To,
		Amount:      t.Amount,
		ValueUSD:    t.ValueUSD.String(),
	}
}

func (t *TransferEvent) unpack(data *storage.TransferEvent) *TransferEvent {
	t.Chain = data.Chain
	t.Contract = data.Contract
	t.BlockNumber = data.BlockNumber
	t.Date = time.Unix(data.Date, 0)
	t.Order = uint(data.Order)
	t.Tx = data.Tx
	t.From = data.From
	t.To = data.To
	t.Amount = data.Amount
	t.ValueUSD = decimal.RequireFromString(data.ValueUSD)
	return t
}

func (t *TransferEvent) Marshal() ([]byte, error) {
	return proto.Marshal(t.pack())
}

func (t *TransferEvent) UnMarshal(buf []byte) error {
	data := &storage.TransferEvent{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	t.unpack(data)
	return nil
}
