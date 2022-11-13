package history_data

import (
	"context"
	"cybercoin/dal/data_object"
	"time"
)

type DataSource interface {
	QueryHistory(ctx context.Context, query CoinHistoryQuery) []*data_object.CoinHisPrice
}

type CoinHistoryQuery struct {
	Coin   string
	Anchor string
	Begin  *time.Time
	End    *time.Time
	Period string
	Type   string
}

type InstType int

const (
	SPOT    InstType = 1 //币币
	SWAP    InstType = 2 //永续合约
	FUTURES InstType = 3 //交割合约
	OPTION  InstType = 4 //期权
)
