package models

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
)

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
func (l *LimitTrade) TableName() string {
	return "limit_trades"
}

func (l *LimitTrade) pack() *storage.LimitTrade {
	return &storage.LimitTrade{
		ID:            uint32(l.ID),
		OrderId:       uint32(l.OrderId),
		Status:        l.Status,
		AmountIn:      l.AmountIn,
		AmountOut:     l.AmountOut,
		GasPrice:      l.GasPrice,
		GasValue:      l.GasValue,
		TriggerTx:     l.TriggerTx,
		ExecutionTx:   l.ExecutionTx,
		FailureReason: l.FailureReason,
		Price:         l.Price,
		PriceUSD:      l.PriceUSD,
	}
}

func (l *LimitTrade) unpack(data *storage.LimitTrade) *LimitTrade {
	l.ID = uint(data.ID)
	l.OrderId = uint(data.OrderId)
	l.Status = data.Status
	l.AmountIn = data.AmountIn
	l.AmountOut = data.AmountOut
	l.GasPrice = data.GasPrice
	l.GasValue = data.GasValue
	l.TriggerTx = data.TriggerTx
	l.ExecutionTx = data.ExecutionTx
	l.FailureReason = data.FailureReason
	l.Price = data.Price
	l.PriceUSD = data.PriceUSD

	return l
}

func (l *LimitTrade) Marshal() ([]byte, error) {
	return proto.Marshal(l.pack())
}

func (l *LimitTrade) UnMarshal(buf []byte) error {
	data := &storage.LimitTrade{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	l.unpack(data)
	return nil
}
