package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	defaultLog "log"
	"net/http"
	"net/url"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// doRequest performs an HTTP request and handles common error scenarios.
func (c *Client) doRequest(
	ctx context.Context,
	urlStr string,
	span trace.Span,
) (*http.Response, error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		c.log.ErrorContext(ctx, "error creating http request", "err", err)
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

	return resp, nil
}

// processResponse handles the HTTP response, checks status code, and reads the body.
func (c *Client) processResponse(ctx context.Context, resp *http.Response, span trace.Span) ([]byte, error) {
	defer func() {
		if err := resp.Body.Close(); err != nil {
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

	return body, nil
}

// unmarshalResponse unmarshals the response body into the provided result.
func (c *Client) unmarshalResponse(ctx context.Context, data []byte, result interface{}, span trace.Span) error {
	if err := json.Unmarshal(data, result); err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		c.log.ErrorContext(ctx, "service error", "err", err)
		return errors.New("service error")
	}

	return nil
}

// buildURL creates a URL with the given parameters.
func buildURL(host, path string) *url.URL {
	return &url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path,
	}
}

// buildAPIURL creates a URL for API endpoints with formatted path.
func (c *Client) buildAPIURL(host string, pathSegments ...string) *url.URL {
	path := ""
	for i, segment := range pathSegments {
		if i > 0 {
			path += "/"
		}
		path += segment
	}
	if len(path) > 0 {
		path += "/"
	}

	return &url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path,
	}
}

// executeRequest performs the full request cycle: sends request, processes response, and unmarshals data.
func (c *Client) executeRequest(
	ctx context.Context,
	urlObj *url.URL,
	result interface{},
	span trace.Span,
) error {
	resp, err := c.doRequest(ctx, urlObj.String(), span)
	if err != nil {
		return err
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			span.SetStatus(codes.Error, err.Error())
			span.RecordError(err)
			defaultLog.Println("error closing response body", err)
		}
	}()

	respBody, err := c.processResponse(ctx, resp, span)
	if err != nil {
		return err
	}

	return c.unmarshalResponse(ctx, respBody, result, span)
}

// executeAPIRequest is a high-level function that handles the complete API request flow.
func (c *Client) executeAPIRequest(
	ctx context.Context,
	operation string,
	host string,
	result interface{},
	pathSegments ...string,
) error {
	ctx, span := c.tracer.Start(ctx, operation)
	defer span.End()

	urlObj := c.buildAPIURL(host, pathSegments...)
	return c.executeRequest(ctx, urlObj, result, span)
}

// executeCustomURLRequest handles requests with custom URL construction.
func (c *Client) executeCustomURLRequest(
	ctx context.Context,
	operation string,
	urlObj *url.URL,
	result interface{},
) error {
	ctx, span := c.tracer.Start(ctx, operation)
	defer span.End()

	return c.executeRequest(ctx, urlObj, result, span)
}
