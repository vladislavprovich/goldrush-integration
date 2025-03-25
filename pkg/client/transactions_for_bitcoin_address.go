package client

import (
	"context"
	"encoding/json"
	"errors"

	"go.opentelemetry.io/otel/codes"
)

func (c *Client) GetTransactionsForBitcoinAddress(
	ctx context.Context,
) (*ResGetTransactionsForBitcoinAddress, error) {
	const op = "client.GetTransactionsForBitcoinAddress"
	ctx, span := c.tracer.Start(ctx, op)
	defer span.End()

	urlObj := c.buildAPIURL(
		c.cfg.Bitcoin.BaseURL,
		c.cfg.Bitcoin.Version,
		c.cfg.Bitcoin.Cq,
		c.cfg.Bitcoin.Covalent,
		c.cfg.Bitcoin.App,
		c.cfg.Bitcoin.Bitcoin,
		c.cfg.Bitcoin.Transactions,
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

	var res ResGetTransactionsForBitcoinAddress
	if err = json.Unmarshal(respBody, &res); err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		c.log.ErrorContext(ctx, "service error", "err", err)
		return nil, errors.New("service error")
	}

	return &res, nil
}
