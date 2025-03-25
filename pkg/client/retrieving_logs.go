package client

import (
	"context"
	"fmt"
)

func (c *Client) GetLogs(ctx context.Context, req *ReqGetLogs) (*ResGetLogs, error) {
	const op = "client.GetLogs"
	ctx, span := c.tracer.Start(ctx, op)
	defer span.End()

	path := fmt.Sprintf("%s/%s/%s/",
		c.cfg.Utility.Version,
		req.ChainName,
		c.cfg.Utility.Events,
	)

	urlObj := buildURL(c.cfg.Utility.BaseURL, path)

	var result ResGetLogs
	err := c.executeRequest(ctx, urlObj, &result, span)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
