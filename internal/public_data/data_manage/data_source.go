package data_manage

import (
	"context"
	"cybercoin/dal/data_object"
	"time"
)

type DataSource interface {
	QueryHistory(ctx context.Context, query CoinHistoryQuery) []*data_object.CoinHisPrice
	QueryNow(ctx context.Context, base, coin string) data_object.CoinHisPrice
}

type CoinHistoryQuery struct {
	Coin     string
	Base     string
	Begin    time.Time
	End      time.Time
	Interval Period
	Write    bool
}
