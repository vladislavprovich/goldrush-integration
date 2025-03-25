package client

import (
	"context"
	"fmt"
)

func (c *Client) GetAllTransactionsBlockByPage(
	ctx context.Context,
	req *ReqGetAllTransactionsBlockByPage,
) (*ResGetAllTransactionsBlockByPage, error) {
	const op = "client.GetAllTransactionsBlockByPage"
	ctx, span := c.tracer.Start(ctx, op)
	defer span.End()

	path := fmt.Sprintf("%s/%s/%s/%b/%s/%s/%b/",
		c.cfg.Transactions.Version,
		req.ChainName,
		c.cfg.Transactions.Block,
		req.BlockHeight,
		c.cfg.Transactions.TransactionsV3,
		c.cfg.Transactions.Page,
		req.Page,
	)

	urlObj := buildURL(c.cfg.Transactions.BaseURL, path)

	var result ResGetAllTransactionsBlockByPage
	err := c.executeRequest(ctx, urlObj, &result, span)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
