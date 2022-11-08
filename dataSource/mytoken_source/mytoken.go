package mytoken_source

import (
	"context"
	"crypto/md5"
	dal2 "cybercoin/dal"
	"cybercoin/dataSource"
	"encoding/hex"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	Token_Salt     = "9527"
	BaseUrl        = "https://api.mytokenapi.com/currency/kline"
	Market_id      = "1324"
	Platform       = "web_pc"
	Version        = "1.0.0"
	Language       = "en_US"
	Legal_Currency = "USD"
)

type MytokenServie struct {
}

func (s *MytokenServie) QueryNow(ctx context.Context, base, coin string) dal2.CoinHisPrice {
	//TODO implement me
	panic("implement me")
}

func (*MytokenServie) QueryHistory(ctx context.Context, query dataSource.CoinHistoryQuery) []*dal2.CoinHisPrice {
	url := BuildUrl(query)
	history := QueryBatch(url)
	res := []*dal2.CoinHisPrice{}
	for history.Message == "Success" && len(history.Data.KlineList) > 0 {
		early := history.Data.KlineList[len(history.Data.KlineList)-1].Time
		if cast.ToInt(query.Begin.Unix()) >= early {
			break
		}
		query.End = time.Unix(cast.ToInt64(early), 0)
		fmt.Println(query.End)
		fmt.Println(len(history.Data.KlineList))
		list := history.Data.KlineList
		batchRes := BuildResult(list, query)
		if query.Write {
			dal2.MySql.CreateInBatches(batchRes, 240)
		}

		res = append(res, batchRes...)
		history = QueryBatch(BuildUrl(query))

	}

	return res
}

func QueryBatch(url string) *KlineHistory {
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	history := KlineHistory{}

	err = jsoniter.Unmarshal(body, &history)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	sort.Slice(history.Data.KlineList, func(i, j int) bool {
		return history.Data.KlineList[i].Time > history.Data.KlineList[j].Time
	})
	return &history
}
func BuildResult(list []*Param, query dataSource.CoinHistoryQuery) []*dal2.CoinHisPrice {
	histories := []*dal2.CoinHisPrice{}
	for _, param := range list {
		history := dal2.CoinHisPrice{}
		history.Coin = query.Coin
		history.Base = query.Base
		history.Time = time.Unix(cast.ToInt64(param.Time), 0)
		history.High = param.High
		history.Low = param.Low
		history.Open = param.Open
		history.Close = param.Close
		history.Volumefrom = param.Volumefrom
		switch query.Interval {
		case dataSource.Minutely:
			history.Period = "1m"
		case dataSource.Quarterly:
			history.Period = "15m"

		}
		histories = append(histories, &history)

	}
	return histories
}

func BuildUrl(query dataSource.CoinHistoryQuery) string {
	stamp := time.Now().UnixMilli()
	stampStr := cast.ToString(stamp)
	token := BuildKLineToken(stampStr)
	url := BaseUrl
	url += "?com_id=" + strings.ToLower(query.Coin+"_"+query.Base)
	url += "&symbol=" + query.Coin
	url += "&anchor=" + query.Base

	url += "&time=" + cast.ToString(query.End.Unix())
	url += "&market_id=" + Market_id
	switch query.Interval {
	case dataSource.Quarterly:
		url += "&period=15m"
	case dataSource.Minutely:
		url += "&period=1m"
	default:
		url += "&period=15m"
	}
	url += "&timestamp=" + stampStr
	url += "&code=" + token
	url += "&platform=" + Platform
	url += "&v=" + Version
	url += "&language=" + Language
	url += "&legal_currency=" + Legal_Currency
	fmt.Println(url)
	return cast.ToString(url)

}

func BuildKLineToken(timeStamp string) string {
	h := md5.New()
	h.Write([]byte(timeStamp + Token_Salt + timeStamp[0:6]))
	return hex.EncodeToString(h.Sum(nil))
}
