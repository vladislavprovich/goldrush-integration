package client

import (
	"context"
	"errors"
)

func (c *Client) GetHistoricalTokenPrice(
	ctx context.Context,
	req *ReqGetHistoricalTokenPrice,
) (*ResGetHistoricalTokenPrice, error) {
	const op = "client.GetHistoricalTokenPrice"
	var result ResGetHistoricalTokenPrice

	err := c.executeAPIRequest(
		ctx,
		op,
		c.cfg.Prices.BaseURL,
		&result,
		c.cfg.Prices.Version,
		c.cfg.Prices.Pricing,
		c.cfg.Prices.HistoricalByAddressesV2,
		req.ChainName,
		req.QuoteCurrency,
		req.ContractAddress,
	)

	if err != nil {
		return nil, errors.New("service error")
	}

	return &result, nil
}
