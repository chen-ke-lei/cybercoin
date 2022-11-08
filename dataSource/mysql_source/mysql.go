package mysql_source

import (
	"context"
	"cybercoin/dal"
	"cybercoin/dataSource"
	"fmt"
)

type MysqlService struct {
}

func (m MysqlService) QueryHistory(ctx context.Context, query dataSource.CoinHistoryQuery) []*dal.CoinHisPrice {
	his := []*dal.CoinHisPrice{}
	dal.MySql.Where("time >= ? and time<=? and period = ? and coin= ? and base =?",
		query.Begin,
		query.End,
		dataSource.BuildPeriod(query.Interval),
		query.Coin,
		query.Base).Find(&his)
	fmt.Println(his)
	return his
}

func (m MysqlService) QueryNow(ctx context.Context, base, coin string) dal.CoinHisPrice {
	//TODO implement me
	panic("implement me")
}
