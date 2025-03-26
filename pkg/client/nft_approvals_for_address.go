package client

import (
	"context"
	"errors"
	"fmt"
	"net/url"
)

func (c *Client) GetNFTApprovalsForAddress(
	ctx context.Context,
	req *ReqGetNFTApprovalsForAddress,
) (*ResGetNFTApprovalsForAddress, error) {
	const op = "client.GetNFTApprovalsForAddress"

	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.Security.BaseURL,
		Path: fmt.Sprintf("%s/%s/%s/%s/%s/",
			c.cfg.Security.Version,
			req.ChainName,
			c.cfg.Security.Nft,
			c.cfg.Security.Approvals,
			req.WalletAddress,
		),
	}

	var result ResGetNFTApprovalsForAddress
	err := c.executeCustomURLRequest(ctx, op, urlForGetStatus, &result)
	if err != nil {
		return nil, errors.New("service error")
	}

	return &result, nil
}
