package client

import (
	"context"
	"errors"
)

func (c *Client) GetAllTransactionsInBlock(
	ctx context.Context,
	req *ReqGetAllTransactionsInBlock,
) (*ResGetAllTransactionsInBlock, error) {
	const op = "client.GetAllTransactionsInBlock"
	var res ResGetAllTransactionsInBlock

	err := c.executeAPIRequest(
		ctx,
		op,
		c.cfg.Transactions.BaseURL,
		&res,
		c.cfg.Transactions.Version,
		req.ChainName,
		c.cfg.Transactions.BlockHash,
		req.BlockHash,
		c.cfg.Transactions.TransactionsV3,
	)

	if err != nil {
		return nil, errors.New("service error")
	}

	return &res, nil
}
