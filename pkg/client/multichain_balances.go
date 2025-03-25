package client

import (
	"context"
	"errors"
	"fmt"
	"net/url"
)

func (c *Client) GetMultichainBalances(
	ctx context.Context,
	req *ReqGetMultichainBalances,
) (*ResGetMultichainBalances, error) {
	const op = "client.GetMultichainBalances"

	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.CrossChain.BaseURL,
		Path: fmt.Sprintf("%s/%s/%s/%s/%s/",
			c.cfg.CrossChain.Version,
			c.cfg.CrossChain.AllChains,
			c.cfg.CrossChain.Address,
			req.WalletAddress,
			c.cfg.CrossChain.Balances,
		),
	}

	var result ResGetMultichainBalances
	err := c.executeCustomURLRequest(ctx, op, urlForGetStatus, &result)
	if err != nil {
		return nil, errors.New("service error")
	}

	return &result, nil
}
