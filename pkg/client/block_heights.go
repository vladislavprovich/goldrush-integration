package client

import (
	"context"
	"fmt"
	"net/url"
)

func (c *Client) GetBlockHeights(ctx context.Context, req *ReqGetBlockHeights) (*ResGetBlockHeights, error) {
	const op = "client.GetBlockHeights"

	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.Utility.BaseURL,
		Path: fmt.Sprintf("%s/%s/%s/%s/%s/",
			c.cfg.Utility.Version,
			req.ChainName,
			c.cfg.Utility.BlockV2,
			req.StartDate,
			req.EndDate,
		),
	}

	var result ResGetBlockHeights
	err := c.executeCustomURLRequest(ctx, op, urlForGetStatus, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
