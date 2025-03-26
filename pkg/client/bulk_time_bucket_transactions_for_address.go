package client

import (
	"context"
	"errors"
	"fmt"
	"net/url"
)

func (c *Client) GetBulkTimeBucketForAddress(
	ctx context.Context,
	req *ReqGetBulkTimeBucketForAddress,
) (*ResGetBulkTimeBucketForAddress, error) {
	const op = "client.GetBulkTimeBucketForAddress"

	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.Transactions.BaseURL,
		Path: fmt.Sprintf("%s/%s/%s/%s/%s/%d/",
			c.cfg.Transactions.Version,
			req.ChainName,
			c.cfg.Transactions.Bulk,
			c.cfg.Transactions.Transactions,
			req.WalletAddress,
			req.TimeBucket,
		),
	}

	var result ResGetBulkTimeBucketForAddress
	err := c.executeCustomURLRequest(ctx, op, urlForGetStatus, &result)
	if err != nil {
		return nil, errors.New("service error")
	}

	return &result, nil
}
