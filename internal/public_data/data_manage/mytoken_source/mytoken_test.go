package mytoken_source_test

import (
	"context"
	data_manage2 "cybercoin/internal/public_data/data_manage"
	"cybercoin/internal/public_data/data_manage/mytoken_source"
	"fmt"
	"testing"
	"time"
)

func TestBuildKLineToken(t *testing.T) {
	fmt.Println(mytoken_source.BuildKLineToken("1667729633894"))
}

func TestMytokenServie_QueryHistory(t *testing.T) {
	servie := mytoken_source.MytokenServie{}
	query := data_manage2.CoinHistoryQuery{}
	query.Base = "USDT"
	query.Coin = "BTC"
	query.Write = true
	query.Interval = data_manage2.Quarterly

	query.End = time.Now()
	query.End = time.Date(2018, 01, 01, 11, 56, 10, 0, time.Local)
	query.Begin = time.Date(2010, 01, 01, 11, 56, 10, 0, time.Local)
	servie.QueryHistory(context.Background(), query)

}
