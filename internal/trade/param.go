package trade

import (
	"context"
	"time"
)

type Account interface {
	Trade(ctx context.Context, base, coin string, price string, interval time.Duration)
}
type Asset struct {
	Name string
	Num  float64
}

type AccountBase struct {
	Assets map[string]Asset
}

func (a *AccountBase) Trade(ctx context.Context, base, coin string, sellMount, buyMount float64, interval time.Duration) {
	baseAsset, baseOk := a.Assets[base]
	coinAsset, coinOk := a.Assets[base]
	if !baseOk || !coinOk || baseAsset.Num < sellMount {
		return
	}
	baseAsset.Num -= sellMount
	coinAsset.Num += buyMount

}
