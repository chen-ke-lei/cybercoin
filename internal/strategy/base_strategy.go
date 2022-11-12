package strategy

import (
	"context"
	"cybercoin/dal/data_object"
	"cybercoin/internal/trade"
)

type Strategy interface {
	BuySign(ctx context.Context, base trade.Account)
	SellSign(ctx context.Context, base trade.Account)
	NewStrategy(ctx context.Context, conf map[string]interface{})
	UpdateStatus(ctx context.Context, price data_object.CoinHisPrice)
}
