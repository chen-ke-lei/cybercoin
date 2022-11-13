package okk_source

import (
	"context"
	"cybercoin/dal/data_object"
	"cybercoin/internal/public_data/history_data"
	"cybercoin/net"
	"cybercoin/util"
	"errors"
	"github.com/spf13/cast"
	"time"
)

const (
	QUERY_URL = "https://www.okx.com/api/v5/market/history-mark-price-candles"
	SOURCE    = "OKK"
)

type OkkService struct{}

func (o OkkService) QueryHistory(ctx context.Context, query *history_data.CoinHistoryQuery) ([]*data_object.CoinHisPrice, error) {
	param, err := buildQueryParam(ctx, query)
	if err != nil {
		return nil, err
	}
	resp, err := net.GetWithParam(ctx, QUERY_URL, param)
	if err != nil {
		return nil, err
	}
	candle := HistoryMarkPriceCandle{}
	err = util.TransferMap2Struct(resp, &candle)
	if err != nil {
		return nil, err
	}
	return buildResByResp(&candle, query)
}

func buildQueryParam(ctx context.Context, query *history_data.CoinHistoryQuery) (map[string]string, error) {
	if query.Coin == "" || query.Anchor == "" || query.Type == "" {
		return nil, errors.New("param is error")
	}
	param := make(map[string]string)
	param["instId"] = query.Coin + "-" + query.Anchor + "-" + query.Type
	if query.End != nil {
		param["after"] = cast.ToString(query.End.UnixMilli())
	}
	if query.Begin != nil {
		param["before"] = cast.ToString(query.End.UnixMilli())
	}
	if query.Period != "" {
		param["bar"] = query.Period
	}
	return param, nil
}
func buildResByResp(candle *HistoryMarkPriceCandle, query *history_data.CoinHistoryQuery) ([]*data_object.CoinHisPrice, error) {
	res := make([]*data_object.CoinHisPrice, 0)
	if candle.Code != 0 {
		return nil, errors.New(candle.Msg)
	}
	for _, item := range candle.Data {

		hisPrice, err := buildHisPrice(query, item)
		if err != nil {
			continue
		}
		res = append(res, hisPrice)

	}
	return res, nil
}
func buildHisPrice(query *history_data.CoinHistoryQuery, item []string) (*data_object.CoinHisPrice, error) {
	if len(item) != 5 {
		return nil, errors.New("return has false format")
	}
	price := data_object.CoinHisPrice{
		Anchor: query.Anchor,
		Coin:   query.Coin,
		Type:   query.Type,
		Period: query.Period,
		Source: SOURCE,
		Time:   time.UnixMilli(cast.ToInt64(item[0])),
		Open:   cast.ToFloat64(item[1]),
		High:   cast.ToFloat64(item[2]),
		Low:    cast.ToFloat64(item[3]),
		Close:  cast.ToFloat64(item[4]),
	}

	return &price, nil
}
