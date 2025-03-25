package client

import (
	"context"
	"fmt"
)

func (c *Client) GetMultichainMultiaddressTransactions(
	ctx context.Context,
) (*RespMultichainMultiaddressTransactions, error) {
	const op = "client.GetMultichainMultiaddressTransactions"
	ctx, span := c.tracer.Start(ctx, op)
	defer span.End()

	path := fmt.Sprintf("%s/%s/%s/",
		c.cfg.CrossChain.Version,
		c.cfg.CrossChain.AllChains,
		c.cfg.CrossChain.Transactions,
	)

	urlObj := buildURL(c.cfg.CrossChain.BaseURL, path)

	var result RespMultichainMultiaddressTransactions
	err := c.executeRequest(ctx, urlObj, &result, span)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
