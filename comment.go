package backlog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
)

type Comment struct {
	Id            uint64         `json:"id"`
	Content       string         `json:"content"`
	ChangeLog     []ChangeLog    `json:"changeLog"`
	CreatedUser   User           `json:"createdUser"`
	Created       Date           `json:"created"`
	Updated       Date           `json:"updated"`
	Stars         []Star         `json:"stars"`
	Notifications []Notification `json:"notifications"`
}

func (c *Client) GetIssueComments(issueId uint64, query url.Values) ([]Comment, error) {
	return c.GetIssueCommentsContext(context.Background(), issueId, query)
}

func (c *Client) GetIssueCommentsContext(ctx context.Context, issueId uint64, query url.Values) ([]Comment, error) {
	var comments []Comment

	path, err := c.root.Parse(path.Join(getIssuesPath, fmt.Sprint(issueId), "comments"))
	if err != nil {
		return nil, err
	}

	response, err := c.getContext(ctx, path, query)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(response, &comments); err != nil {
		return nil, err
	}
	return comments, nil
}

func (c *Client) GetPullRequestComments(projectKeyOrId, repositoryNameOrId string, number string, query url.Values) ([]Comment, error) {
	return c.GetPullRequestCommentsContext(context.Background(), projectKeyOrId, repositoryNameOrId, number, query)
}

func (c *Client) GetPullRequestCommentsContext(ctx context.Context, projectKeyOrId, repositoryNameOrId string, number string, query url.Values) ([]Comment, error) {
	var comments []Comment

	path, err := c.root.Parse(path.Join(
		getProjectsPath, projectKeyOrId,
		"git", "repositories", repositoryNameOrId,
		"pullRequests", number, "comments",
	))
	if err != nil {
		return nil, err
	}

	response, err := c.getContext(ctx, path, query)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(response, &comments); err != nil {
		return nil, err
	}
	return comments, nil
}
