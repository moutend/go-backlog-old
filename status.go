package backlog

import (
	"context"
	"encoding/json"
)

type Status struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

const (
	getStatusesPath = "/api/v2/statuses"
)

func (c *Client) GetStatuses() ([]Status, error) {
	return c.GetStatusesContext(context.Background())
}

func (c *Client) GetStatusesContext(ctx context.Context) ([]Status, error) {
	var statuses []Status

	path, err := c.root.Parse(getStatusesPath)
	if err != nil {
		return nil, err
	}

	response, err := c.getContext(ctx, path, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(response, &statuses); err != nil {
		return nil, err
	}

	return statuses, nil
}
