package models

import (
	"time"

	"github.com/kattana-io/models/pkg-bin/storage"

	"google.golang.org/protobuf/proto"
)

type ApiTradesBlock struct {
	Trades []*ApiTrade `json:"trades"`
}

func (tb *ApiTradesBlock) pack() *storage.ApiTradesBlock {
	trades := make([]*storage.ApiTrade, 0)
	for _, b := range tb.Trades {
		trades = append(trades, b.Pack())
	}

	return &storage.ApiTradesBlock{
		Trades: trades,
	}
}

func (tb *ApiTradesBlock) unpack(data *storage.ApiTradesBlock) *ApiTradesBlock {
	trades := make([]*ApiTrade, 0)
	for _, b := range data.Trades {
		trades = append(trades, new(ApiTrade).Unpack(b))
	}
	tb.Trades = trades

	return tb
}

func (tb *ApiTradesBlock) Marshal() ([]byte, error) {
	return proto.Marshal(tb.pack())
}

func (tb *ApiTradesBlock) UnMarshal(buf []byte) error {
	data := &storage.ApiTradesBlock{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	tb.unpack(data)
	return nil
}

/**
 * Actually this should be a PairSwap
 */
// Trade example
type ApiTrade struct {
	Tx          string    `ch:"tx" json:"tx"`
	Date        time.Time `ch:"date" json:"date"`
	Timestamp   uint32    `ch:"timestamp" json:"timestamp"`
	Chain       string    `ch:"chain" json:"chain"`
	Blocknumber uint64    `ch:"blocknumber" json:"blocknumber"`
	Pair        string    `ch:"pair" json:"pair"`
	Amount0     string    `ch:"amount0" json:"amount0"`
	Amount1     string    `ch:"amount1" json:"amount1"`
	Buy         uint8     `ch:"buy" json:"buy"`
	PriceA      float32   `ch:"pricea" json:"price_a"`
	PriceAUSD   float32   `ch:"pricea_usd" json:"price_a_usd"`
	PriceB      float32   `ch:"priceb" json:"price_b"`
	PriceBUSD   float32   `ch:"priceb_usd" json:"price_b_usd"`
	Bot         uint8     `ch:"bot" json:"bot"`
	Wallet      string    `ch:"wallet" json:"wallet"`
	Order       uint16    `ch:"order" json:"order"`
	ValueUsd    float32   `ch:"value_usd" json:"value_usd"`
	Ticker      any       `json:"ticker,omitempty"`
}

func (t *ApiTrade) Pack() *storage.ApiTrade {
	var buy, bot bool
	if buy = false; t.Buy == 1 {
		buy = true
	}
	if bot = false; t.Bot == 1 {
		bot = true
	}

	return &storage.ApiTrade{
		Tx:          t.Tx,
		Date:        t.Date.Unix(),
		Timestamp:   int64(t.Timestamp),
		Chain:       t.Chain,
		Blocknumber: t.Blocknumber,
		Pair:        t.Pair,
		Amount0:     t.Amount0,
		Amount1:     t.Amount1,
		Buy:         buy,
		Bot:         bot,
		PriceA:      t.PriceA,
		PriceB:      t.PriceB,
		PriceAUSD:   t.PriceAUSD,
		PriceBUSD:   t.PriceBUSD,
		Wallet:      t.Wallet,
		Order:       int32(t.Order),
		ValueUsd:    t.ValueUsd,
	}
}

func (t *ApiTrade) Unpack(data *storage.ApiTrade) *ApiTrade {
	var buy, bot uint8
	if buy = 0; data.Buy {
		buy = 1
	}
	if bot = 0; data.Bot {
		bot = 1
	}

	t.Tx = data.Tx
	t.Date = time.Unix(data.Date, 0)
	t.Timestamp = uint32(data.Timestamp)
	t.Chain = data.Chain
	t.Blocknumber = data.Blocknumber
	t.Pair = data.Pair
	t.Amount0 = data.Amount0
	t.Amount1 = data.Amount1
	t.Buy = buy
	t.PriceA = data.PriceA
	t.PriceB = data.PriceB
	t.PriceAUSD = data.PriceAUSD
	t.PriceBUSD = data.PriceBUSD
	t.Bot = bot
	t.Wallet = data.Wallet
	t.Order = uint16(data.Order)
	t.ValueUsd = data.ValueUsd

	return t
}

func (t *ApiTrade) Marshal() ([]byte, error) {
	return proto.Marshal(t.Pack())
}

func (t *ApiTrade) UnMarshal(buf []byte) error {
	data := &storage.ApiTrade{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	t.Unpack(data)
	return nil
}
