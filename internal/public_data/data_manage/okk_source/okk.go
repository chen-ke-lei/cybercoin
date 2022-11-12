package okk_source

import (
	"context"
	"cybercoin/dal/data_object"
	"cybercoin/internal/public_data/data_manage"
)

type OkkService struct{}

func (o OkkService) QueryHistory(ctx context.Context, query data_manage.CoinHistoryQuery) []*data_object.CoinHisPrice {
	//TODO implement me
	panic("implement me")
}

func (o OkkService) QueryNow(ctx context.Context, base, coin string) data_object.CoinHisPrice {
	//TODO implement me
	panic("implement me")
}
