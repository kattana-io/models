package models

import (
	"time"
)

type DirectSwap struct {
	Tx          string    `ch:"tx" json:"tx"`
	Date        time.Time `ch:"date" json:"date"`
	Chain       string    `ch:"chain" json:"chain"`
	BlockNumber uint64    `ch:"blocknumber" json:"blocknumber"`
	Protocol    string    `ch:"protocol" json:"protocol"`
	SrcToken    string    `ch:"src_token" json:"src_token"`
	DstToken    string    `ch:"dst_token" json:"dst_token"`
	Amount0     string    `ch:"amount0" json:"amount0"`
	Amount1     string    `ch:"amount1" json:"amount1"`
	PriceA      float32   `ch:"pricea" json:"price_a"`
	PriceAUSD   float32   `ch:"pricea_usd" json:"price_a_usd"`
	PriceB      float32   `ch:"priceb" json:"price_b"`
	PriceBUSD   float32   `ch:"priceb_usd" json:"price_b_usd"`
	Wallet      string    `ch:"wallet" json:"wallet"`
	Order       uint16    `ch:"order" json:"order"`
	ValueUSD    float32   `ch:"value_usd" json:"value_usd"`
}
