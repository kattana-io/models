package models

type LimitTrade struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	OrderId       uint    `gorm:"column:order_id;type:int64" json:"order_id"`
	Status        string  `gorm:"column:status;type:VARCHAR;" json:"status"`
	AmountIn      string  `gorm:"column:amount_in;type:VARCHAR;" json:"amount_in"`
	AmountOut     string  `gorm:"column:amount_out;type:VARCHAR;" json:"amount_out"`
	GasPrice      int64   `gorm:"column:gas_price;type:int64;" json:"gas_price"`
	GasValue      int64   `gorm:"column:gas_value;type:int64;" json:"gas_value"`
	TriggerTx     string  `gorm:"column:trigger_tx;type:VARCHAR;" json:"trigger_tx"`
	ExecutionTx   string  `gorm:"column:execution_tx;type:VARCHAR;" json:"execution_tx"`
	FailureReason string  `gorm:"column:failure_reason;type:VARCHAR;" json:"failure_reason"`
	Price         float64 `gorm:"column:price;type:float8;" json:"price"`
	PriceUSD      float64 `gorm:"column:price_usd;type:float8;" json:"price_usd"`
}

// TableName overrides the table name used by User to `profiles`
func (LimitTrade) TableName() string {
	return "limit_trades"
}
