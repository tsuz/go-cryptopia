package cryptopia

import (
	"context"

	"github.com/pkg/errors"
)

type GetTradeHistory []TradeHistory

type TradeHistory struct {
	TradeId     int
	TradePairId int
	Market      string
	Type        string
	Rate        float64
	Amount      float64
	Total       float64
	Fee         float64
	TimeStamp   string
}

// GetOpenOrders returns all open orders
func (c *Client) GetTradeHistory(ctx context.Context, market map[string]interface{}) (GetTradeHistory, error) {
	req, err := c.newAuthenticatedRequest(ctx, "GetTradeHistory", market)
	if err != nil {
		return GetTradeHistory{}, errors.Wrap(err, "Faild to new authenticated request")
	}

	var ret = &GetTradeHistory{}
	_, err = c.do(req, ret)
	if err != nil {
		return *ret, errors.Wrap(err, "Failed to do request")
	}
	return *ret, nil
}
