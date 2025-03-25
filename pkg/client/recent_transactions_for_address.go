package client

import (
	"context"
	"encoding/json"
	"errors"
)

func (c *Client) GetRecentTransactionForAddress(
	ctx context.Context,
	req *ReqGetRecentTransactionForAddress,
) (*ResGetRecentTransactionForAddress, error) {
	const op = "client.GetRecentTransactionForAddress"
	ctx, span := c.tracer.Start(ctx, op)
	defer span.End()

	urlObj := c.buildAPIURL(
		c.cfg.Transactions.BaseURL,
		c.cfg.Transactions.Version,
		req.ChainName,
		c.cfg.Transactions.Address,
		req.WalletAddress,
		c.cfg.Transactions.TransactionsV3,
	)

	resp, err := c.doRequest(ctx, urlObj.String(), span)
	if err != nil {
		return nil, errors.New("service error")
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()

	respBody, err := c.processResponse(ctx, resp, span)
	if err != nil {
		return nil, errors.New("service error")
	}

	var wrapper TransactionResponseRecentTransactionForAddress
	if err = json.Unmarshal(respBody, &wrapper); err == nil {
		c.log.InfoContext(ctx, "Successfully unmarshaled transaction response with wrapper")
		// Ensure Items is initialized even if it's null in the response
		if wrapper.Data.Items == nil {
			wrapper.Data.Items = []TransactionItemRecent{}
		}
		return &wrapper.Data, nil
	}

	return &wrapper.Data, nil
}
