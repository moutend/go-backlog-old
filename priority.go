package backlog

import (
	"context"
	"encoding/json"
)

type Priority struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

const (
	getPrioritiesPath = "/api/v2/priorities"
)

func (c *Client) GetPriorities() ([]Priority, error) {
	return c.GetPrioritiesContext(context.Background())
}

func (c *Client) GetPrioritiesContext(ctx context.Context) ([]Priority, error) {
	var priorities []Priority

	path, err := c.root.Parse(getPrioritiesPath)
	if err != nil {
		return nil, err
	}

	response, err := c.getContext(ctx, path, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(response, &priorities); err != nil {
		return nil, err
	}

	return priorities, nil
}
