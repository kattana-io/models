package models

import (
	"strings"
)

type Pair struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Name          string `gorm:"column:name;type:VARCHAR;" json:"name"`
	FullName      string `gorm:"column:full_name;type:VARCHAR;" json:"full_name"`
	Chain         string `gorm:"column:chain;type:VARCHAR;" json:"chain"`
	Exchange      string `gorm:"column:exchange;type:VARCHAR;" json:"exchange"`
	Address       string `gorm:"column:address;type:VARCHAR;" json:"address"`
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
	PriceA       float64 `gorm:"-" json:"price_a"`
	PriceAUSD    float64 `gorm:"-" json:"price_a_usd"`
	PriceB       float64 `gorm:"-" json:"price_b"`
	PriceBUSD    float64 `gorm:"-" json:"price_b_usd"`
	Liquidity    float64 `gorm:"-" json:"liquidity"`
	Volume       float64 `gorm:"-" json:"volume"`
	BaseReserve  int64   `gorm:"-" json:"base_reserve"`
	QuoteReserve int64   `gorm:"-" json:"quote_reserve"`
	UpdatedAt    int64   `gorm:"-" json:"updated_at"`
	APY          float64 `gorm:"-" json:"annual_percentage_yield"`
	BaseSlug     string  `gorm:"-" json:"base_slug"`
	QuoteSlug    string  `gorm:"-" json:"quote_slug"`
}

// TableName overrides the table name used by User to `profiles`
func (Pair) TableName() string {
	return "pairs"
}

// GetQuoteAsset WBTC/WETH will return WETH
func (p *Pair) GetQuoteAsset() string {
	split := strings.Split(p.Name, "/")
	if len(split) == 2 {
		return split[1]
	}
	return ""
}

// GetBaseAsset WBTC/WETH will return WBTC
func (p *Pair) GetBaseAsset() string {
	split := strings.Split(p.Name, "/")
	if len(split) == 2 {
		return split[0]
	}
	return ""
}

func (p *Pair) TradesToUSD() bool {
	return strings.Contains(p.GetQuoteAsset(), "USD")
}

func (p *Pair) GetExchange() string {
	return p.Exchange
}

func (p *Pair) GetBaseQuoteAsset() (string, string) {
	split := strings.Split(p.Name, "/")
	return split[0], split[1]
}

func (p *Pair) GetKey() string {
	return p.Exchange + "|" + p.Name
}
