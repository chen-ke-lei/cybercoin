package mysql_source

import (
	"context"
	"cybercoin/dal/data_object"
	data_manage2 "cybercoin/internal/public_data/data_manage"
	"fmt"
)

type MysqlService struct {
}

func (m MysqlService) QueryHistory(ctx context.Context, query data_manage2.CoinHistoryQuery) []*data_object.CoinHisPrice {
	his := []*data_object.CoinHisPrice{}
	data_object.MySql.Where("time >= ? and time<=? and period = ? and coin= ? and base =?",
		query.Begin,
		query.End,
		data_manage2.BuildPeriod(query.Interval),
		query.Coin,
		query.Base).Find(&his)
	fmt.Println(his)
	return his
}

func (m MysqlService) QueryNow(ctx context.Context, base, coin string) data_object.CoinHisPrice {
	//TODO implement me
	panic("implement me")
}
