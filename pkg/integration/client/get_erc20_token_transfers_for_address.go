package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel/codes"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) GetERC20TokenTransfersForAddress(ctx context.Context, req *ReqGetERC20TokenTransfersForAddress) (*ResGetERC20TokenTransfersForAddress, error) {
	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.BaseURL,
		Path:   fmt.Sprintf("v1/%s/address/%s/transfers_v2/", req.ChainName, req.WalletAddress),
	}

	const op = "client.GetERC20TokenTransfersForAddress"
	ctx, span := c.tracer.Start(ctx, op)
	defer span.End()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, urlForGetStatus.String(), nil)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		c.log.ErrorContext(ctx, "error creating http request for get activity across all chains", "err", err)
		return nil, errors.New("service error")
	}

	httpReq.Header.Set("Content-Type", "application/json")

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

	var res ResGetERC20TokenTransfersForAddress
	if err = json.Unmarshal(body, &res); err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		c.log.ErrorContext(ctx, "service error", "err", err)
		return nil, errors.New("service error")
	}

	return &res, nil
}
