package models

import "time"

type TraderWalletCard struct {
	FirstTradeAt time.Time         `json:"first_trade_at"`
	LastTradeAt  time.Time         `json:"last_trade_at"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	TotalTrades  uint64            `json:"total_trades"`
	Pairs        map[string]uint64 `json:"pairs"`
	PairInfo     *TraderWalletPair `json:"pair_info"`
}

type TraderWalletPair struct {
	BuyTradesCount  uint64  `json:"buy_trades_count"`
	SellTradesCount uint64  `json:"sell_trades_count"`
	BuyValue        float64 `json:"buy_value"`
	BuyValueUsd     float64 `json:"buy_value_usd"`
	SellValue       float64 `json:"sell_value"`
	SellValueUsd    float64 `json:"sell_value_usd"`
}
