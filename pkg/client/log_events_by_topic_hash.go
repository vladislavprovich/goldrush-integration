package client

import (
	"context"
)

func (c *Client) GetLogEventsByTopicHash(
	ctx context.Context,
	req *ReqGetLogEventsByTopicHash,
) (*ResGetLogEventsByTopicHash, error) {
	const op = "client.GetLogEventsByTopicHash"

	var result ResGetLogEventsByTopicHash
	err := c.executeAPIRequest(ctx, op, c.cfg.Utility.BaseURL, &result,
		c.cfg.Utility.Version,
		req.ChainName,
		c.cfg.Utility.Events,
		c.cfg.Utility.Topics,
		req.TopicHash,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
