package client

import (
	"context"
	"encoding/json"
	"errors"
)

func (c *Client) GetHistoricalPortfolioValueOverTime(
	ctx context.Context,
	req *ReqGetHistoricalPortfolioValueOverTime,
) (*ResGetHistoricalPortfolioValueOverTime, error) {
	const op = "client.GetHistoricalPortfolioValueOverTime"
	ctx, span := c.tracer.Start(ctx, op)
	defer span.End()

	urlObj := c.buildAPIURL(
		c.cfg.Balances.BaseURL,
		c.cfg.Balances.Version,
		req.ChainName,
		c.cfg.Balances.Address,
		req.WalletAddress,
		c.cfg.Balances.PortfolioV2,
	)

	resp, err := c.doRequest(ctx, urlObj.String(), span)
	if err != nil {
		return nil, errors.New("service error")
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()

	respBody, err := c.processResponse(ctx, resp, span)
	if err != nil {
		return nil, errors.New("service error")
	}

	var wrapper TransactionResponsePortfolioValueOverTime
	if err = json.Unmarshal(respBody, &wrapper); err == nil {
		return &wrapper.Data, nil
	}

	return &wrapper.Data, nil
}
