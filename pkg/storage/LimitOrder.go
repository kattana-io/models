package models

import (
	"github.com/kattana-io/models/pkg-bin/storage"
	"google.golang.org/protobuf/proto"
)

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
	StartPrice          float64 `gorm:"column:start_price;type:float8;" json:"start_price"`
	StartPriceUSD       float64 `gorm:"column:start_price_usd;type:float8;" json:"start_price_usd"`
	EndPrice            float64 `gorm:"column:end_price;type:float8;" json:"end_price"`
	EndPriceUSD         float64 `gorm:"column:end_price_usd;type:float8;" json:"end_price_usd"`
	ExecutionType       string  `gorm:"column:execution_type;type:VARCHAR;" json:"execution_type"`
	Gt                  bool    `gorm:"column:gt;type:BOOL;" json:"gt"`
	Name                string  `gorm:"column:name;type:VARCHAR;" json:"name"`
	Exchange            string  `gorm:"column:exchange;type:VARCHAR;" json:"exchange"`
	Proxy               string  `gorm:"column:proxy;type:VARCHAR;" json:"proxy"`
	AmountIn            string  `gorm:"column:amount_in;type:VARCHAR;" json:"amount_in"`
	AmountOut           string  `gorm:"column:amount_out;type:VARCHAR;" json:"amount_out"`
	SrcAddress          string  `gorm:"column:src_address;type:VARCHAR;" json:"src_address"`
	SrcDecimals         int32   `gorm:"column:src_decimals;type:INT4;default:2;" json:"src_decimals"`
	DestAddress         string  `gorm:"column:dest_address;type:VARCHAR;" json:"dest_address"`
	DestDecimals        int32   `gorm:"column:dest_decimals;type:INT4;default:2;" json:"dest_decimals"`
	GasSent             int64   `gorm:"column:gas_sent;type:int64;" json:"gas_sent"`
	TotalGas            int64   `gorm:"column:total_gas;type:int64;" json:"total_gas"`
	ExpiresAt           int64   `gorm:"column:expires_at;type:int64;" json:"expires_at"`
	TriggerType         string  `gorm:"column:trigger_type;type:VARCHAR;" json:"trigger_type"`
	Slippage            uint    `gorm:"column:slippage;type:int64;" json:"slippage"`
	ExecutedTradesCount int64   `gorm:"column:executed_trades_count;type:int64;" json:"executed_trades_count"`
	TradesCount         int64   `gorm:"column:trades_count;type:int64;" json:"trades_count"`
	CreatedAt           int64   `gorm:"column:created_at;type:int64;" json:"created_at"`
	TradesTimeDistance  int64   `gorm:"column:trades_time_distance;type:int64;" json:"trades_time_distance"`

	LimitTrades []LimitTrade `gorm:"foreignKey:OrderId" json:"limit_trades"`
}

// TableName overrides the table name used by User to `profiles`
func (l *LimitOrder) TableName() string {
	return "limit_orders"
}

func (l *LimitOrder) pack() *storage.LimitOrder {
	return &storage.LimitOrder{
		ID:                  uint32(l.ID),
		Status:              l.Status,
		Uuid:                l.Uuid,
		Wallet:              l.Wallet,
		Chain:               l.Chain,
		Pair:                l.Pair,
		Flipped:             l.Flipped,
		Price:               l.Price,
		PriceUSD:            l.PriceUSD,
		StartPrice:          l.StartPrice,
		StartPriceUSD:       l.StartPriceUSD,
		EndPrice:            l.EndPrice,
		EndPriceUSD:         l.EndPriceUSD,
		ExecutionType:       l.ExecutionType,
		Gt:                  l.Gt,
		Name:                l.Name,
		Exchange:            l.Exchange,
		Proxy:               l.Proxy,
		AmountIn:            l.AmountIn,
		AmountOut:           l.AmountOut,
		SrcAddress:          l.SrcAddress,
		SrcDecimals:         l.SrcDecimals,
		DestAddress:         l.DestAddress,
		DestDecimals:        l.DestDecimals,
		GasSent:             l.GasSent,
		TotalGas:            l.TotalGas,
		ExpiresAt:           l.ExpiresAt,
		TriggerType:         l.TriggerType,
		Slippage:            uint32(l.Slippage),
		ExecutedTradesCount: l.ExecutedTradesCount,
		TradesCount:         l.TradesCount,
		CreatedAt:           l.CreatedAt,
		TradesTimeDistance:  l.TradesTimeDistance,
	}
}

func (l *LimitOrder) unpack(data *storage.LimitOrder) *LimitOrder {
	l.ID = uint(data.ID)
	l.Status = data.Status
	l.Uuid = data.Uuid
	l.Wallet = data.Wallet
	l.Chain = data.Chain
	l.Pair = data.Pair
	l.Flipped = data.Flipped
	l.Price = data.Price
	l.PriceUSD = data.PriceUSD
	l.StartPrice = data.StartPrice
	l.StartPriceUSD = data.StartPriceUSD
	l.EndPrice = data.EndPrice
	l.EndPriceUSD = data.EndPriceUSD
	l.ExecutionType = data.ExecutionType
	l.Gt = data.Gt
	l.Name = data.Name
	l.Exchange = data.Exchange
	l.Proxy = data.Proxy
	l.AmountIn = data.AmountIn
	l.AmountOut = data.AmountOut
	l.SrcAddress = data.SrcAddress
	l.SrcDecimals = data.SrcDecimals
	l.DestAddress = data.DestAddress
	l.DestDecimals = data.DestDecimals
	l.GasSent = data.GasSent
	l.TotalGas = data.TotalGas
	l.ExpiresAt = data.ExpiresAt
	l.TriggerType = data.TriggerType
	l.Slippage = uint(data.Slippage)
	l.ExecutedTradesCount = data.ExecutedTradesCount
	l.TradesCount = data.TradesCount
	l.CreatedAt = data.CreatedAt
	l.TradesTimeDistance = data.TradesTimeDistance
	return l
}

func (l *LimitOrder) Marshal() ([]byte, error) {
	return proto.Marshal(l.pack())
}

func (l *LimitOrder) UnMarshal(buf []byte) error {
	data := &storage.LimitOrder{}
	err := proto.Unmarshal(buf, data)
	if err != nil {
		return err
	}
	// TODO add block validate
	l.unpack(data)
	return nil
}
