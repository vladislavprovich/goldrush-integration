package client

import (
	"context"
	"errors"
)

func (c *Client) GetAllChainStatuses(ctx context.Context) (*ResGetAllChainStatuses, error) {
	const op = "client.GetAllChainStatuses"
	var res ResGetAllChainStatuses

	err := c.executeAPIRequest(
		ctx,
		op,
		c.cfg.Utility.BaseURL,
		&res,
		c.cfg.Utility.Version,
		c.cfg.Utility.Chains,
		c.cfg.Utility.Status,
	)

	if err != nil {
		return nil, errors.New("service error")
	}

	return &res, nil
}
