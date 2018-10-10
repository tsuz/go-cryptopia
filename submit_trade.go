package cryptopia

import (
	"context"

	"github.com/pkg/errors"
)

type SubmitTrade struct {
	OrderID      int
	FilledOrders []int
}

// SubmitTrade submits an order
func (c *Client) SubmitTrade(ctx context.Context, market map[string]interface{}) (SubmitTrade, error) {
	req, err := c.newAuthenticatedRequest(ctx, "SubmitTrade", market)
	if err != nil {
		return SubmitTrade{}, errors.Wrap(err, "Faild to new authenticated request")
	}

	var ret = &SubmitTrade{}
	_, err = c.do(req, ret)
	if err != nil {
		return *ret, errors.Wrap(err, "Failed to do request")
	}
	return *ret, nil
}
