package client

import (
	"context"
	"errors"
	"fmt"
	"net/url"
)

func (c *Client) GetTokenApprovalsForAddress(
	ctx context.Context,
	req *ReqGetTokenApprovalsForAddress,
) (*ResGetTokenApprovalsForAddress, error) {
	const op = "client.GetTokenApprovalsForAddress"

	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.Security.BaseURL,
		Path: fmt.Sprintf("%s/%s/%s/%s/",
			c.cfg.Security.Version,
			req.ChainName,
			c.cfg.Security.Approvals,
			req.WalletAddress,
		),
	}

	var result ResGetTokenApprovalsForAddress
	err := c.executeCustomURLRequest(ctx, op, urlForGetStatus, &result)
	if err != nil {
		return nil, errors.New("service error")
	}

	return &result, nil
}
