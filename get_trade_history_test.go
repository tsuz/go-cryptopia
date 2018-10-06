package cryptopia

import (
	"context"
	"fmt"
	"testing"
)

func TestGetTradeHistory(t *testing.T) {
	params := map[string]interface{}{
		"Market": "XBY_BTC",
	}
	cryptopia := newAuthClient()
	ctx := context.Background()
	ret, err := cryptopia.GetTradeHistory(ctx, params)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
