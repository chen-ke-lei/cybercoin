package mytoken_source

type KlineHistory struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      *Data  `json:"data"`
	Timestamp int    `json:"timestamp"`
}

type Data struct {
	KlineList []*Param `json:"kline"`
	Detail    *Detail  `json:"detail"`
}

type Detail struct {
	PercentChangeDisplay string  `json:"percent_change_display"`
	PercentChangeRange   string  `json:"percent_change_range"`
	Volume24hFrom        float64 `json:"volume_24h_from"`
	Low24h               string  `json:"low_24h"`
	High24h              string  `json:"high_24h"`
	PriceDisplay         string  `json:"price_display"`
	HrPriceDisplay       string  `json:"hr_price_display"`
	LegalCurrencyPrice   float64 `json:"legal_currency_price"`
	Anchor               string  `json:"anchor"`
}
type Param struct {
	Time       int     `json:"time"`
	Open       float64 `json:"open"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	Close      float64 `json:"close"`
	Volumefrom float64 `json:"volumefrom"`
}
