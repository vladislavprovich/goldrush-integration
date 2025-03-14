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

func (c *Client) GetLogEventsByTopicHash(ctx context.Context, req *ReqGetLogEventsByTopicHash) (*ResGetLogEventsByTopicHash, error) {
	urlForGetStatus := &url.URL{
		Scheme: "https",
		Host:   c.cfg.BaseURL,
		Path:   fmt.Sprintf("v1/%s/events/topics/%s/", req.ChainName, req.TopicHash),
	}

	const op = "client.GetLogEventsByTopicHash"
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

	var res ResGetLogEventsByTopicHash
	if err = json.Unmarshal(body, &res); err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		c.log.ErrorContext(ctx, "service error", "err", err)
		return nil, errors.New("service error")
	}

	return &res, nil
}
