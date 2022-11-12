package mytoken_source

import (
	"context"
	"crypto/md5"
	"cybercoin/dal/data_object"
	"cybercoin/dal/mysql"
	data_manage2 "cybercoin/internal/public_data/data_manage"
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

func (s *MytokenServie) QueryNow(ctx context.Context, base, coin string) data_object.CoinHisPrice {
	//TODO implement me
	panic("implement me")
}

func (*MytokenServie) QueryHistory(ctx context.Context, query data_manage2.CoinHistoryQuery) []*data_object.CoinHisPrice {
	url := BuildUrl(query)
	history := QueryBatch(url)
	res := []*data_object.CoinHisPrice{}
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
			mysql.MySql.CreateInBatches(batchRes, 240)
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
func BuildResult(list []*Param, query data_manage2.CoinHistoryQuery) []*data_object.CoinHisPrice {
	histories := []*data_object.CoinHisPrice{}
	for _, param := range list {
		history := data_object.CoinHisPrice{}
		history.Coin = query.Coin
		history.Base = query.Base
		history.Time = time.Unix(cast.ToInt64(param.Time), 0)
		history.High = param.High
		history.Low = param.Low
		history.Open = param.Open
		history.Close = param.Close
		history.Volumefrom = param.Volumefrom
		switch query.Interval {
		case data_manage2.Minutely:
			history.Period = "1m"
		case data_manage2.Quarterly:
			history.Period = "15m"

		}
		histories = append(histories, &history)

	}
	return histories
}

func BuildUrl(query data_manage2.CoinHistoryQuery) string {
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
	case data_manage2.Quarterly:
		url += "&period=15m"
	case data_manage2.Minutely:
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
