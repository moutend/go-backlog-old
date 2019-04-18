package backlog

import (
	"context"
	"encoding/json"
	"net/url"
)

type Notification struct {
	Id                  int         `json:"id"`
	AlreadyRead         bool        `json:"alreadyRead"`
	Reason              int         `json:"reason"`
	ResourceAlreadyRead bool        `json:"resourceAlreadyRead"`
	Project             Project     `json:"project"`
	Issue               Issue       `json:"issue"`
	Comment             Comment     `json:"comment"`
	PullRequest         PullRequest `json:"pullRequest"`
	PullRequestComment  Comment     `json:"pullRequestComment"`
	Sender              User        `json:"sender"`
	Created             Date        `json:"created"`
}

const (
	getNotificationsPath = "/api/v2/notifications"
)

func (c *Client) GetNotifications(query url.Values) ([]Notification, error) {
	return c.GetNotificationsContext(context.Background(), query)
}

func (c *Client) GetNotificationsContext(ctx context.Context, query url.Values) ([]Notification, error) {
	var wikis []Notification

	path, err := c.root.Parse(getNotificationsPath)
	if err != nil {
		return nil, err
	}

	response, err := c.getContext(ctx, path, query)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(response, &wikis); err != nil {
		return nil, err
	}

	return wikis, nil
}
