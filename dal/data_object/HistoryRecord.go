package data_object

import "time"

type CoinHisPrice struct {
	id         int `gorm:"-;primary_key;auto_increment"`
	Time       time.Time
	Open       float64
	High       float64
	Low        float64
	Close      float64
	Volumefrom float64
	Base       string
	Coin       string
	Period     string
}

func (CoinHisPrice) TableName() string {
	return "coin_his_price"
}
