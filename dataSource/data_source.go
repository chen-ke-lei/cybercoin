package dataSource

import (
	"context"
	"cybercoin/dal"
	"time"
)

type DataSource interface {
	QueryHistory(ctx context.Context, query CoinHistoryQuery) []*dal.CoinHisPrice
	QueryNow(ctx context.Context, base, coin string) dal.CoinHisPrice
}

type CoinHistoryQuery struct {
	Coin     string
	Base     string
	Begin    time.Time
	End      time.Time
	Interval Period
	Write    bool
}
