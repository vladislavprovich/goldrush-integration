package client

import (
	"context"
	"errors"
	"fmt"
	"net/url"
)

func (c *Client) GetTokenHoldersBlockHeight(
	ctx context.Context,
	req *ReqGetTokenHoldersBlockHeight,
) (*ResGetTokenHoldersBlockHeight, error) {
	const op = "client.GetTokenHoldersBlockHeight"

	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.Balances.BaseURL,
		Path: fmt.Sprintf("%s/%s/%s/%s/%s/",
			c.cfg.Balances.Version,
			req.ChainName,
			c.cfg.Balances.Address,
			req.WalletAddress,
			c.cfg.Balances.TokenHoldersV2,
		),
	}

	var result ResGetTokenHoldersBlockHeight
	err := c.executeCustomURLRequest(ctx, op, urlForGetStatus, &result)
	if err != nil {
		return nil, errors.New("service error")
	}

	return &result, nil
}
