package strategy

import (
	"context"
	"cybercoin/dal"
	"cybercoin/trade_item"
)

type Strategy interface {
	BuySign(ctx context.Context, base trade_item.Account)
	SellSign(ctx context.Context, base trade_item.Account)
	NewStrategy(ctx context.Context, conf map[string]interface{})
	UpdateStatus(ctx context.Context, price dal.CoinHisPrice)
}
