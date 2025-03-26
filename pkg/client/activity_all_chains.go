package client

import (
	"context"
	"fmt"
	"net/url"
)

func (c *Client) GetActivityAcrossAllChains(
	ctx context.Context,
	req *ReqGetActivityAcrossAllChains,
) (*RespGetActivityAcrossAllChains, error) {
	const op = "client.GetActivityAcrossAllChains"

	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.CrossChain.BaseURL,
		Path: fmt.Sprintf("%s/%s/%s/%s/",
			c.cfg.CrossChain.Version,
			c.cfg.CrossChain.Address,
			req.WalletAddress,
			c.cfg.CrossChain.Activity,
		),
	}

	var result RespGetActivityAcrossAllChains
	err := c.executeCustomURLRequest(ctx, op, urlForGetStatus, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
