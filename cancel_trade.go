package cryptopia

import (
	"context"

	"github.com/pkg/errors"
)

type CancelTrade struct {
	OrderID      int
	FilledOrders []int
}

// CancelTrade submits an order
func (c *Client) CancelTrade(ctx context.Context, market map[string]interface{}) (CancelTrade, error) {
	req, err := c.newAuthenticatedRequest(ctx, "CancelTrade", market)
	if err != nil {
		return CancelTrade{}, errors.Wrap(err, "Faild to new authenticated request")
	}

	var ret = &CancelTrade{}
	_, err = c.do(req, ret)
	if err != nil {
		return *ret, errors.Wrap(err, "Failed to do request")
	}
	return *ret, nil
}
