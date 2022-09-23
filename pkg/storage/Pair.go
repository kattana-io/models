package models

import (
	"strings"
)

type Pair struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	PairName      string `gorm:"column:name;type:VARCHAR;" json:"name"`
	Chain         string `gorm:"column:chain;type:VARCHAR;" json:"chain"`
	Exchange      string `gorm:"column:exchange;type:VARCHAR;" json:"exchange"`
	PairAddress   string `gorm:"column:address;type:VARCHAR;" json:"address"`
	BaseAddress   string `gorm:"column:base_address;type:VARCHAR;" json:"base_address"`
	BaseDecimals  int32  `gorm:"column:base_decimals;type:INT4;default:2;" json:"base_decimals"`
	QuoteAddress  string `gorm:"column:quote_address;type:VARCHAR;" json:"quote_address"`
	QuoteDecimals int32  `gorm:"column:quote_decimals;type:INT4;default:2;" json:"quote_decimals"`
	Hidden        bool   `gorm:"column:hidden;type:BOOL;" json:"hidden"`
	Hot           bool   `gorm:"column:hot;type:BOOL;" json:"hot"`
	Flipped       bool   `gorm:"column:flipped;type:BOOL;" json:"flipped"`
	Rank          byte   `gorm:"column:rank;type:int32;" json:"rank"`
	Kind          byte   `gorm:"column:kind;type:int32;" json:"kind"`
	PoolCreated   int64  `gorm:"column:pool_created;type:INT4;default:0;" json:"pool_created"`
	BaseId        uint   `gorm:"column:base_id;type:INT4;" json:"base_id"`
	QuoteId       uint   `gorm:"column:quote_id;type:INT4;" json:"quote_id"`
	// there are no columns in DB (OpenSearch only)
	Volume       float64 `json:"volume"`
	Liquidity    float64 `json:"liquidity"`
	Price        float64 `json:"price"`
	BaseVolume   float64 `json:"base_volume"`
	BaseReserve  float64 `json:"base_reserve"`
	QuoteReserve float64 `json:"quote_reserve"`
	PriceUSD     float64 `json:"price_usd"`
	LastTrade    int64   `json:"last_trade"`
}

// TableName overrides the table name used by User to `profiles`
func (Pair) TableName() string {
	return "pairs"
}

// GetQuoteAsset WBTC/WETH will return WETH
func (t *Pair) GetQuoteAsset() string {
	split := strings.Split(t.PairName, "/")
	if len(split) == 2 {
		return split[1]
	}
	return ""
}

// GetBaseAsset WBTC/WETH will return WBTC
func (t *Pair) GetBaseAsset() string {
	split := strings.Split(t.PairName, "/")
	if len(split) == 2 {
		return split[0]
	}
	return ""
}

func (t *Pair) TradesToUSD() bool {
	return strings.Contains(t.GetQuoteAsset(), "USD")
}

func (t *Pair) GetExchange() string {
	return t.Exchange
}

func (t *Pair) GetBaseQuoteAsset() (string, string) {
	split := strings.Split(t.PairName, "/")
	return split[0], split[1]
}

func (t *Pair) GetKey() string {
	return t.Exchange + "|" + t.PairName
}
