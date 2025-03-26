package client

import (
	"context"
	"errors"
)

func (c *Client) GetEarliesrTransactionForAddress(
	ctx context.Context,
	req *ReqGetEarliesrTransactionForAddress,
) (*ResGetEarliestTransactionForAddress, error) {
	const op = "client.GetEarliesrTransactionForAddress"
	var res ResGetEarliestTransactionForAddress

	err := c.executeAPIRequest(
		ctx,
		op,
		c.cfg.Transactions.BaseURL,
		&res,
		c.cfg.Transactions.Version,
		req.ChainName,
		c.cfg.Transactions.Bulk,
		c.cfg.Transactions.Transactions,
		req.WalletAddress,
	)

	if err != nil {
		return nil, errors.New("service error")
	}

	return &res, nil
}
