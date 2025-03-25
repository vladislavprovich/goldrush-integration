package client

import (
	"context"
	"errors"
)

func (c *Client) GetAllChains(ctx context.Context) (*ResGetAllChains, error) {
	const op = "client.GetAllChains"
	var res ResGetAllChains

	err := c.executeAPIRequest(
		ctx,
		op,
		c.cfg.Utility.BaseURL,
		&res,
		c.cfg.Utility.Version,
		c.cfg.Utility.Chains,
	)

	if err != nil {
		return nil, errors.New("service error")
	}

	return &res, nil
}
