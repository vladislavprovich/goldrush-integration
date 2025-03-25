package client

import (
	"context"
	"fmt"
	"net/url"
)

func (c *Client) GetBitcoinBalancesForHDAddress(
	ctx context.Context,
	req *ReqGetBitcoinBalancesForHDAddress,
) (*ResGetBitcoinBalancesForHDAddress, error) {
	const op = "client.GetBitcoinBalancesForHDAddress"

	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.Bitcoin.BaseURL,
		Path: fmt.Sprintf("%s/%s/%s/%s/%s/",
			c.cfg.Bitcoin.Version,
			c.cfg.Bitcoin.BTCMainnet,
			c.cfg.Bitcoin.Address,
			req.WalletAddress,
			c.cfg.Bitcoin.HdWallets,
		),
	}

	var result ResGetBitcoinBalancesForHDAddress
	err := c.executeCustomURLRequest(ctx, op, urlForGetStatus, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
