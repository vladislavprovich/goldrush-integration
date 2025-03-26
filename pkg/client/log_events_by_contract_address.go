package client

import (
	"context"
)

func (c *Client) GetLogEventsByContractAddress(
	ctx context.Context,
	req *ReqGetLogEventsByContractAddress,
) (*ResGetLogEventsByContractAddress, error) {
	const op = "client.GetLogEventsByContractAddress"

	var result ResGetLogEventsByContractAddress
	err := c.executeAPIRequest(ctx, op, c.cfg.Utility.BaseURL, &result,
		c.cfg.Utility.Version,
		req.ChainName,
		c.cfg.Utility.Events,
		c.cfg.Utility.Address,
		req.ContractAddress,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
