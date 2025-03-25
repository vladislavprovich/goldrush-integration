package client

import (
	"context"
	"errors"
	"fmt"
	"net/url"
)

func (c *Client) GetERC20TokenTransfersForAddress(
	ctx context.Context,
	req *ReqGetERC20TokenTransfersForAddress,
) (*ResGetERC20TokenTransfersForAddress, error) {
	const op = "client.GetERC20TokenTransfersForAddress"

	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.Balances.BaseURL,
		Path: fmt.Sprintf("%s/%s/%s/%s/%s/",
			c.cfg.Balances.Version,
			req.ChainName,
			c.cfg.Balances.Address,
			req.WalletAddress,
			c.cfg.Balances.TransfersV2,
		),
	}

	var result ResGetERC20TokenTransfersForAddress
	err := c.executeCustomURLRequest(ctx, op, urlForGetStatus, &result)
	if err != nil {
		return nil, errors.New("service error")
	}

	return &result, nil
}
