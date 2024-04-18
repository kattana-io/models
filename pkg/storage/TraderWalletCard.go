package models

import "time"

type TraderWalletCard struct {
	FirstTradeAt time.Time         `json:"first_trade_at"`
	LastTradeAt  time.Time         `json:"last_trade_at"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	TotalTrades  uint64            `json:"total_trades"`
	Pairs        map[string]uint64 `json:"pairs"`
}
