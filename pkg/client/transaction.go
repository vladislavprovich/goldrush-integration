package client

import (
	"context"
	"encoding/json"
	"fmt"
)

func (c *Client) GetTransaction(ctx context.Context, req *ReqGetTransaction) (*ResGetTransaction, error) {
	const op = "client.GetTransaction"
	ctx, span := c.tracer.Start(ctx, op)
	defer span.End()

	path := fmt.Sprintf("%s/%s/%s/%s/",
		c.cfg.Transactions.Version,
		req.ChainName,
		c.cfg.Transactions.TransactionV2,
		req.TxHash,
	)

	urlObj := buildURL(c.cfg.Transactions.BaseURL, path)

	// For this specific endpoint, we need to handle the response manually.
	// Because it has a special wrapper structure.
	resp, err := c.doRequest(ctx, urlObj.String(), span)
	if err != nil {
		return nil, err
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()

	respBody, err := c.processResponse(ctx, resp, span)
	if err != nil {
		return nil, err
	}

	// Always try to unmarshal with the data wrapper first
	var wrapper TransactionResponseWrapper
	if err = json.Unmarshal(respBody, &wrapper); err == nil {
		return &wrapper.Data, nil
	}

	c.log.InfoContext(ctx, "Successfully unmarshaled transaction response directly")
	return &wrapper.Data, nil
}
