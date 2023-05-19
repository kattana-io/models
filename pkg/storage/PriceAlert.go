package models

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
)

type PriceAlert struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Uuid string `gorm:"column:uuid;type:uuid;" json:"uuid"`

	Status     string  `gorm:"column:status;type:VARCHAR;" json:"status"`
	Exchange   string  `gorm:"column:exchange;type:VARCHAR;" json:"exchange"`
	Name       string  `gorm:"column:name;type:VARCHAR;" json:"name"`
	Wallet     string  `gorm:"column:wallet;type:VARCHAR;" json:"wallet"`
	Token      string  `gorm:"column:token;type:VARCHAR;" json:"token"`
	Chain      string  `gorm:"column:chain;type:VARCHAR;" json:"chain"`
	Pair       string  `gorm:"column:pair;type:VARCHAR;" json:"pair"`
	Flipped    bool    `gorm:"column:flipped;type:BOOL;" json:"flipped"`
	Price      float64 `gorm:"column:price;type:float8;" json:"price"`
	PriceUSD   float64 `gorm:"column:price_usd;type:float8;" json:"price_usd"`
	Gt         bool    `gorm:"column:gt;type:BOOL;" json:"gt"`
	CreatedAt  int64   `gorm:"column:created_at;type:int64;" json:"created_at"`
	ExecutedAt int64   `gorm:"column:executed_at;type:int64;" json:"executed_at"`
	Hidden     bool    `gorm:"column:hidden;type:BOOL;" json:"hidden"`
}

// TableName overrides the table name used by User to `profiles`
func (p *PriceAlert) TableName() string {
	return "price_alerts"
}

func (p *PriceAlert) pack() *storage.PriceAlert {
	return &storage.PriceAlert{
		ID:         uint32(p.ID),
		Uuid:       p.Uuid,
		Status:     p.Status,
		Exchange:   p.Exchange,
		Name:       p.Name,
		Wallet:     p.Wallet,
		Token:      p.Token,
		Chain:      p.Chain,
		Pair:       p.Pair,
		Flipped:    p.Flipped,
		Price:      p.Price,
		PriceUSD:   p.PriceUSD,
		Gt:         p.Gt,
		CreatedAt:  p.CreatedAt,
		ExecutedAt: p.ExecutedAt,
		Hidden:     p.Hidden,
	}
}

func (p *PriceAlert) unpack(data *storage.PriceAlert) *PriceAlert {
	p.ID = uint(data.ID)
	p.Uuid = data.Uuid
	p.Status = data.Status
	p.Exchange = data.Exchange
	p.Name = data.Name
	p.Wallet = data.Wallet
	p.Token = data.Token
	p.Chain = data.Chain
	p.Pair = data.Pair
	p.Flipped = data.Flipped
	p.Price = data.Price
	p.PriceUSD = data.PriceUSD
	p.Gt = data.Gt
	p.CreatedAt = data.CreatedAt
	p.ExecutedAt = data.ExecutedAt
	p.Hidden = data.Hidden
	return p
}

func (p *PriceAlert) Marshal() ([]byte, error) {
	return proto.Marshal(p.pack())
}

func (p *PriceAlert) UnMarshal(buf []byte) error {
	data := &storage.PriceAlert{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	p.unpack(data)
	return nil
}
