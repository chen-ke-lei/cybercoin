package okk_source

import (
	"context"
	"cybercoin/dal"
	"cybercoin/dataSource"
)

type OkkService struct{}

func (o OkkService) QueryHistory(ctx context.Context, query dataSource.CoinHistoryQuery) []*dal.CoinHisPrice {
	//TODO implement me
	panic("implement me")
}

func (o OkkService) QueryNow(ctx context.Context, base, coin string) dal.CoinHisPrice {
	//TODO implement me
	panic("implement me")
}
