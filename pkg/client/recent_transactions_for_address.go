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

func (c *Client) GetRecentTransactionForAddress(ctx context.Context, req *ReqGetRecentTransactionForAddress) (*ResGetRecentTransactionForAddress, error) {
	const op = "client.GetRecentTransactionForAddress"
	ctx, span := c.tracer.Start(ctx, op)
	defer span.End()

	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.Transactions.BaseURL,
		Path: fmt.Sprintf("%s/%s/%s/%s/%s/",
			c.cfg.Transactions.Version,
			req.ChainName,
			c.cfg.Transactions.Address,
			req.WalletAddress,
			c.cfg.Transactions.TransactionsV3,
		),
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, urlForGetStatus.String(), nil)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		c.log.ErrorContext(ctx, "error creating http request for get recent transactions for address", "err", err)
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

	var wrapper TransactionResponseRecentTransactionForAddress
	if err = json.Unmarshal(body, &wrapper); err == nil {
		c.log.InfoContext(ctx, "Successfully unmarshaled transaction response with wrapper")
		// Ensure Items is initialized even if it's null in the response
		if wrapper.Data.Items == nil {
			wrapper.Data.Items = []TransactionItemRecent{}
		}
		return &wrapper.Data, nil
	}

	return &wrapper.Data, nil
}
