package models

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
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
	Order        int32   `gorm:"-" json:"order"`
}

// TableName overrides the table name used by User to `profiles`
func (p *Pair) TableName() string {
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

func (p *Pair) pack() *storage.Pair {
	return &storage.Pair{
		ID:            uint32(p.ID),
		Name:          p.Name,
		FullName:      p.FullName,
		Chain:         p.Chain,
		Exchange:      p.Exchange,
		Address:       p.Address,
		BaseAddress:   p.BaseAddress,
		BaseDecimals:  p.BaseDecimals,
		QuoteAddress:  p.QuoteAddress,
		QuoteDecimals: p.QuoteDecimals,
		Hidden:        p.Hidden,
		Hot:           p.Hot,
		Flipped:       p.Flipped,
		Rank:          int32(p.Rank),
		Kind:          int32(p.Kind),
		PoolCreated:   p.PoolCreated,
		BaseId:        uint32(p.BaseId),
		QuoteId:       uint32(p.QuoteId),
		PriceA:        p.PriceA,
		PriceAUSD:     p.PriceAUSD,
		PriceB:        p.PriceB,
		PriceBUSD:     p.PriceBUSD,
		Liquidity:     p.Liquidity,
		Volume:        p.Volume,
		BaseReserve:   p.BaseReserve,
		QuoteReserve:  p.QuoteReserve,
		UpdatedAt:     p.UpdatedAt,
		APY:           p.APY,
		BaseSlug:      p.BaseSlug,
		QuoteSlug:     p.QuoteSlug,
		Order:         p.Order,
	}
}

func (p *Pair) unpack(data *storage.Pair) *Pair {
	p.ID = uint(data.ID)
	p.Name = data.Name
	p.FullName = data.FullName
	p.Chain = data.Chain
	p.Exchange = data.Exchange
	p.Address = data.Address
	p.BaseAddress = data.BaseAddress
	p.BaseDecimals = data.BaseDecimals
	p.QuoteAddress = data.QuoteAddress
	p.QuoteDecimals = data.QuoteDecimals
	p.Hidden = data.Hidden
	p.Hot = data.Hot
	p.Flipped = data.Flipped
	p.Rank = byte(data.Rank)
	p.Kind = byte(data.Kind)
	p.PoolCreated = data.PoolCreated
	p.BaseId = uint(data.BaseId)
	p.QuoteId = uint(data.QuoteId)
	p.PriceA = data.PriceA
	p.PriceAUSD = data.PriceAUSD
	p.PriceB = data.PriceB
	p.PriceBUSD = data.PriceBUSD
	p.Liquidity = data.Liquidity
	p.Volume = data.Volume
	p.BaseReserve = data.BaseReserve
	p.QuoteReserve = data.QuoteReserve
	p.UpdatedAt = data.UpdatedAt
	p.APY = data.APY
	p.BaseSlug = data.BaseSlug
	p.QuoteSlug = data.QuoteSlug
	p.Order = data.Order

	return p
}

func (p *Pair) Marshal() ([]byte, error) {
	return proto.Marshal(p.pack())
}

func (p *Pair) UnMarshal(buf []byte) error {
	data := &storage.Pair{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	p.unpack(data)
	return nil
}
