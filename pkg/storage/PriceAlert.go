package models

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
}

// TableName overrides the table name used by User to `profiles`
func (PriceAlert) TableName() string {
	return "price_alerts"
}
