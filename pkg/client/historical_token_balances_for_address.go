package client

import (
	"context"
	"errors"
)

func (c *Client) GetHistoricalTokenBalancesForAddress(
	ctx context.Context,
	req *ReqGetHistoricalTokenBalancesForAddress,
) (*ResGetHistoricalTokenBalancesForAddress, error) {
	const op = "client.GetHistoricalTokenBalancesForAddress"
	var res ResGetHistoricalTokenBalancesForAddress

	err := c.executeAPIRequest(
		ctx,
		op,
		c.cfg.Balances.BaseURL,
		&res,
		c.cfg.Balances.Version,
		req.ChainName,
		c.cfg.Balances.Address,
		req.WalletAddress,
		c.cfg.Balances.HistoricalBalances,
	)

	if err != nil {
		return nil, errors.New("service error")
	}

	return &res, nil
}
