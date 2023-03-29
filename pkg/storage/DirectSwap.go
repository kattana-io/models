package models

import (
	"math/big"
	"time"

	"github.com/shopspring/decimal"
)

type DirectSwap struct {
	Tx          string          `ch:"tx" json:"tx"`
	Date        time.Time       `ch:"date" json:"date"`
	Chain       string          `ch:"chain" json:"chain"`
	BlockNumber uint64          `ch:"blocknumber" json:"blocknumber"`
	Protocol    string          `ch:"protocol" json:"protocol"`
	SrcToken    string          `ch:"src_token" json:"src_token"`
	DstToken    string          `ch:"dst_token" json:"dst_token"`
	Amount0     *big.Int        `ch:"amount0" json:"amount0"`
	Amount1     *big.Int        `ch:"amount1" json:"amount1"`
	PriceA      decimal.Decimal `ch:"pricea" json:"pricea"`
	PriceAUSD   decimal.Decimal `ch:"pricea_usd" json:"pricea_usd"`
	PriceB      decimal.Decimal `ch:"priceb" json:"priceb"`
	PriceBUSD   decimal.Decimal `ch:"priceb_usd" json:"priceb_usd"`
	Wallet      string          `ch:"wallet" json:"wallet"`
	Order       uint16          `ch:"order" json:"order"`
	ValueUSD    decimal.Decimal `ch:"value_usd" json:"value_usd"`
}
