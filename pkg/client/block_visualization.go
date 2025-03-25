package client

import (
	"context"
	"errors"
	"fmt"
	"net/url"
)

func (c *Client) GetABlock(
	ctx context.Context,
	req *ReqGetABlock,
) (*ResGetABlock, error) {
	const op = "client.GetABlock"

	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.Utility.BaseURL,
		Path: fmt.Sprintf("%s/%s/%s/%s/",
			c.cfg.Utility.Version,
			req.ChainName,
			c.cfg.Utility.BlockV2,
			req.BlockHeight,
		),
	}

	var result ResGetABlock
	err := c.executeCustomURLRequest(ctx, op, urlForGetStatus, &result)
	if err != nil {
		return nil, errors.New("service error")
	}

	return &result, nil
}
