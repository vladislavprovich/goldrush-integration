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

func (c *Client) GetLogs(ctx context.Context, req *ReqGetLogs) (*ResGetLogs, error) {
	const op = "client.GetLogs"
	ctx, span := c.tracer.Start(ctx, op)
	defer span.End()

	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.Utility.BaseURL,
		Path: fmt.Sprintf("%s/%s/%s/",
			c.cfg.Utility.Version,
			req.ChainName,
			c.cfg.Utility.Events,
		),
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, urlForGetStatus.String(), nil)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		c.log.ErrorContext(ctx, "error creating http request for get activity across all chains", "err", err)
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

	var res ResGetLogs
	if err = json.Unmarshal(body, &res); err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		c.log.ErrorContext(ctx, "service error", "err", err)
		return nil, errors.New("service error")
	}

	return &res, nil
}
