package models

import (
	"github.com/shopspring/decimal"
	"math/big"
)

type NftEvent struct {
	Provider  string          `json:"provider,omitempty"`
	EventType string          `json:"event_type,omitempty"`
	Price     decimal.Decimal `json:"price,omitempty"`
	PriceUsd  decimal.Decimal `json:"price_usd,omitempty"`
	Token     string          `json:"token,omitempty"`
	TokenID   *big.Int        `json:"token_id,omitempty"`
	Wallet    string          `json:"wallet,omitempty"`
	Date      int64           `json:"date,omitempty"`
	Chain     string          `json:"chain,omitempty"`
}
