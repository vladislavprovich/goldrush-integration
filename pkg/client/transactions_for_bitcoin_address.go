package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"go.opentelemetry.io/otel/codes"
)

func (c *Client) GetTransactionsForBitcoinAddress(
	ctx context.Context,
) (*ResGetTransactionsForBitcoinAddress, error) {
	const op = "client.GetTransactionsForBitcoinAddress"
	ctx, span := c.tracer.Start(ctx, op)
	defer span.End()

	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.Bitcoin.BaseURL,
		Path: fmt.Sprintf("%s/%s/%s/%s/%s/%s/",
			c.cfg.Bitcoin.Version,
			c.cfg.Bitcoin.Cq,
			c.cfg.Bitcoin.Covalent,
			c.cfg.Bitcoin.App,
			c.cfg.Bitcoin.Bitcoin,
			c.cfg.Bitcoin.Transactions,
		),
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, urlForGetStatus.String(), nil)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		c.log.ErrorContext(ctx, "error creating http request for get activity across all chains",
			"err", err)
		return nil, errors.New("service error")
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.SetBasicAuth(c.cfg.APIKey, "")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		c.log.ErrorContext(ctx, "service error", "err", err)
		return nil, errors.New("service error")
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			span.SetStatus(codes.Error, err.Error())
			span.RecordError(err)

			panic(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		span.SetStatus(codes.Error, fmt.Sprintf("%s status code %d", resp.Status, resp.StatusCode))
		span.RecordError(fmt.Errorf("%s status code %d", resp.Status, resp.StatusCode))

		c.log.ErrorContext(ctx, "service error", "status", resp.StatusCode)
		return nil, errors.New("service error")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		c.log.ErrorContext(ctx, "service error", "err", err)
		return nil, errors.New("service error")
	}

	var res ResGetTransactionsForBitcoinAddress
	if err = json.Unmarshal(body, &res); err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		c.log.ErrorContext(ctx, "service error", "err", err)
		return nil, errors.New("service error")
	}

	return &res, nil
}
