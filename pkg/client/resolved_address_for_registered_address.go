package client

import (
	"context"
	"errors"
	"fmt"
	"net/url"
)

func (c *Client) GetResolverAddressForRegisteredAddress(
	ctx context.Context,
	req *ReqGetResolverAddressForRegisteredAddress,
) (*ResGetResolverAddressForRegisteredAddress, error) {
	const op = "client.GetResolverAddressForRegisteredAddress"

	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.Utility.BaseURL,
		Path: fmt.Sprintf("%s/%s/%s/%s/%s/",
			c.cfg.Utility.Version,
			req.ChainName,
			c.cfg.Utility.Address,
			req.WalletAddress,
			c.cfg.Utility.ResolveAddress,
		),
	}

	var result ResGetResolverAddressForRegisteredAddress
	err := c.executeCustomURLRequest(ctx, op, urlForGetStatus, &result)
	if err != nil {
		return nil, errors.New("service error")
	}

	return &result, nil
}
