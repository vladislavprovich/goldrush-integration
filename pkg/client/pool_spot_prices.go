package client

import (
	"context"
	"errors"
	"fmt"
	"net/url"
)

func (c *Client) GetPoolSpotPrices(
	ctx context.Context,
	req *ReqGetPoolSpotPrices,
) (*ResGetPoolSpotPrices, error) {
	const op = "client.GetPoolSpotPrices"

	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.Prices.BaseURL,
		Path: fmt.Sprintf("%s/%s/%s/%s/%s/%s/",
			c.cfg.Prices.Version,
			c.cfg.Prices.Pricing,
			c.cfg.Prices.SpotPrices,
			req.ChainName,
			c.cfg.Prices.Pools,
			req.ContractAddress,
		),
	}

	var result ResGetPoolSpotPrices
	err := c.executeCustomURLRequest(ctx, op, urlForGetStatus, &result)
	if err != nil {
		return nil, errors.New("service error")
	}

	return &result, nil
}
