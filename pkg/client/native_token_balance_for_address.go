package client

import (
	"context"
	"errors"
	"fmt"
	"net/url"
)

func (c *Client) GetNativeTokenBalancesForAddress(
	ctx context.Context,
	req *ReqGetNativeTokenBalancesForAddress,
) (*ResGetNativeTokenBalancesForAddress, error) {
	const op = "client.GetNativeTokenBalancesForAddress"

	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.Balances.BaseURL,
		Path: fmt.Sprintf("%s/%s/%s/%s/%s/",
			c.cfg.Balances.Version,
			req.ChainName,
			c.cfg.Balances.Address,
			req.WalletAddress,
			c.cfg.Balances.BalancesNative,
		),
	}

	var result ResGetNativeTokenBalancesForAddress
	err := c.executeCustomURLRequest(ctx, op, urlForGetStatus, &result)
	if err != nil {
		return nil, errors.New("service error")
	}

	return &result, nil
}
