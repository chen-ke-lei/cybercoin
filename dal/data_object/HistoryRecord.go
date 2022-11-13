package data_object

import "time"

type CoinHisPrice struct {
	id     uint `gorm:"-;primary_key;auto_increment"`
	Time   time.Time
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Anchor string
	Coin   string
	Period string
	Source string
	Type   int
	Market int
}

func (CoinHisPrice) TableName() string {
	return "coin_his_price"
}
