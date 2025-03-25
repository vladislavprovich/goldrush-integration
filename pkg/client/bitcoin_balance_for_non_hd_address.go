package client

import (
	"context"
	"errors"
	"fmt"
	"net/url"
)

func (c *Client) GetBitcoinBalancesForNonHDAddress(
	ctx context.Context,
	req *ReqGetBitcoinBalancesForNonHDAddress,
) (*ResGetBitcoinBalancesForNonHDAddress, error) {
	const op = "client.GetBitcoinBalancesForNonHDAddress"

	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.Bitcoin.BaseURL,
		Path: fmt.Sprintf("%s/%s/%s/%s/%s/",
			c.cfg.Bitcoin.Version,
			c.cfg.Bitcoin.BTCMainnet,
			c.cfg.Bitcoin.Address,
			req.WalletAddress,
			c.cfg.Bitcoin.BalancesV2,
		),
	}

	var result ResGetBitcoinBalancesForNonHDAddress
	err := c.executeCustomURLRequest(ctx, op, urlForGetStatus, &result)
	if err != nil {
		return nil, errors.New("service error")
	}

	return &result, nil
}
