package client

import (
	"context"
	"errors"
	"fmt"
	"net/url"
)

func (c *Client) GetPaginatedTransactionForAddress(
	ctx context.Context,
	req *ReqGetPaginatedTransactionForAddress,
) (*ResGetPaginatedTransactionForAddress, error) {
	const op = "client.GetPaginatedTransactionForAddress"

	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.Transactions.BaseURL,
		Path: fmt.Sprintf("%s/%s/%s/%s/%s/%s/%d/",
			c.cfg.Transactions.Version,
			req.ChainName,
			c.cfg.Transactions.Address,
			req.WalletAddress,
			c.cfg.Transactions.TransactionsV3,
			c.cfg.Transactions.Page,
			req.Page,
		),
	}

	var result ResGetPaginatedTransactionForAddress
	err := c.executeCustomURLRequest(ctx, op, urlForGetStatus, &result)
	if err != nil {
		return nil, errors.New("service error")
	}

	return &result, nil
}
