package client

import (
	"context"
	"fmt"
)

func (c *Client) GetTransactionSummaryForAddress(
	ctx context.Context,
	req *ReqGetTransactionSummaryForAddress,
) (*ResGetTransactionSummaryForAddress, error) {
	const op = "client.GetTransactionSummaryForAddress"
	ctx, span := c.tracer.Start(ctx, op)
	defer span.End()

	path := fmt.Sprintf("%s/%s/%s/%s/%s/",
		c.cfg.Transactions.Version,
		req.ChainName,
		c.cfg.Transactions.Address,
		req.WalletAddress,
		c.cfg.Transactions.TransactionsSummary,
	)

	urlObj := buildURL(c.cfg.Transactions.BaseURL, path)

	var result ResGetTransactionSummaryForAddress
	err := c.executeRequest(ctx, urlObj, &result, span)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
