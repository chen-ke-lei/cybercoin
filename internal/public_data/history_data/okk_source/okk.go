package okk_source

import (
	"context"
	"cybercoin/dal/data_object"
	"cybercoin/internal/public_data/history_data"
)

const (
	QUERY_URL = "https://www.okx.com/api/v5/market/history-mark-price-candles"
)

type OkkService struct{}

func (o OkkService) QueryHistory(ctx context.Context, query history_data.CoinHistoryQuery) []*data_object.CoinHisPrice {
	//TODO implement me
	panic("implement me")
}
