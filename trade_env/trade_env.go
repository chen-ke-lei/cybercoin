package trade_env

import (
	"context"
	"cybercoin/dataSource"
	"cybercoin/strategy"
	"cybercoin/trade_item"
	"time"
)

type TradeContext struct {
	Account      trade_item.Account
	Strategy     strategy.Strategy
	DataSource   dataSource.DataSource
	Base         string
	CoinList     []string
	TimeInterval time.Duration
}

func StartEnv(ctx context.Context, tradeContext TradeContext) {
	for true {
		for _, coin := range tradeContext.CoinList {
			newPrice := tradeContext.DataSource.QueryNow(ctx, tradeContext.Base, coin)
			tradeContext.Strategy.UpdateStatus(ctx, newPrice)
		}
		tradeContext.Strategy.BuySign(ctx, tradeContext.Account)
		tradeContext.Strategy.SellSign(ctx, tradeContext.Account)
		time.Sleep(tradeContext.TimeInterval)
	}
}
