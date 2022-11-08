package mysql_source_test

import (
	"context"
	"cybercoin/dataSource"
	"cybercoin/dataSource/mysql_source"
	"fmt"
	"testing"
	"time"
)

func TestMysqlService_QueryHistory(t *testing.T) {
	service := mysql_source.MysqlService{}
	query := dataSource.CoinHistoryQuery{}
	query.Base = "USDT"
	query.Coin = "BTC"
	//query.Write = true
	query.Interval = dataSource.Quarterly

	query.End = time.Now()
	//query.End = time.Date(2018, 01, 01, 11, 56, 10, 0, time.Local)
	query.Begin = time.Date(2022, 01, 01, 11, 56, 10, 0, time.Local)
	history := service.QueryHistory(context.Background(), query)
	fmt.Println(len(history))
}
