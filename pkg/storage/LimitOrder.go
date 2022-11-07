package models

type LimitOrder struct {
	ID                  uint    `gorm:"primaryKey" json:"id"`
	Status              string  `gorm:"column:status;type:VARCHAR;" json:"status"`
	Uuid                string  `gorm:"column:uuid;type:uuid;" json:"uuid"`
	Wallet              string  `gorm:"column:wallet;type:VARCHAR;" json:"wallet"`
	Chain               string  `gorm:"column:chain;type:VARCHAR;" json:"chain"`
	Pair                string  `gorm:"column:pair;type:VARCHAR;" json:"pair"`
	Flipped             bool    `gorm:"column:flipped;type:BOOL;" json:"flipped"`
	Price               float64 `gorm:"column:price;type:float8;" json:"price"`
	PriceUSD            float64 `gorm:"column:price_usd;type:float8;" json:"price_usd"`
	Gt                  bool    `gorm:"column:gt;type:BOOL;" json:"gt"`
	Name                string  `gorm:"column:name;type:VARCHAR;" json:"name"`
	Exchange            string  `gorm:"column:exchange;type:VARCHAR;" json:"exchange"`
	Proxy               string  `gorm:"column:proxy;type:VARCHAR;" json:"proxy"`
	AmountIn            string  `gorm:"column:amount_in;type:VARCHAR;" json:"amount_in"`
	AmountOut           string  `gorm:"column:amount_out;type:VARCHAR;" json:"amount_out"`
	SrcAddress          string  `gorm:"column:src_address;type:VARCHAR;" json:"src_address"`
	DestAddress         string  `gorm:"column:dest_address;type:VARCHAR;" json:"dest_address"`
	GasSent             int64   `gorm:"column:gas_sent;type:int64;" json:"gas_sent"`
	TotalGas            int64   `gorm:"column:total_gas;type:int64;" json:"total_gas"`
	ExpiresAt           int64   `gorm:"column:expires_at;type:int64;" json:"expires_at"`
	TriggerType         string  `gorm:"column:trigger_type;type:VARCHAR;" json:"trigger_type"`
	Slippage            int64   `gorm:"column:slippage;type:int64;" json:"slippage"`
	ExecutedTradesCount int64   `gorm:"column:executed_trades_count;type:int64;" json:"executed_trades_count"`
	TradesCount         int64   `gorm:"column:trades_count;type:int64;" json:"trades_count"`
}

// TableName overrides the table name used by User to `profiles`
func (LimitOrder) TableName() string {
	return "limit_orders"
}
