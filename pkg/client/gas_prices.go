package client

import (
	"context"
	"fmt"
	"net/url"
)

func (c *Client) GetGasPrices(ctx context.Context, req *ReqGetGasPrices) (*ResGetGasPrices, error) {
	const op = "client.GetGasPrices"

	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.Prices.BaseURL,
		Path: fmt.Sprintf("%s/%s/%s/%s/%s/",
			c.cfg.Prices.Version,
			req.ChainName,
			c.cfg.Prices.Event,
			req.EventType,
			c.cfg.Prices.GasPrices,
		),
	}

	var result ResGetGasPrices
	err := c.executeCustomURLRequest(ctx, op, urlForGetStatus, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
