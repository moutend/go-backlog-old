package backlog

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"
	"path"
)

type Issue struct {
	Id             uint64       `json:"id"`
	ProjectId      uint64       `json:"projectId"`
	IssueKey       string       `json:"issueKey"`
	KeyId          uint64       `json:"keyId"`
	IssueType      IssueType    `json:"issueType"`
	Summary        string       `json:"summary"`
	Description    string       `json:"description"`
	Resolution     Resolution   `json:"resolution"`
	Priority       Priority     `json:"priority"`
	Status         Status       `json:"status"`
	Assignee       User         `json:"assignee"`
	Category       []Category   `json:"category"`
	Versions       []Version    `json:"versions"`
	Milestone      []Milestone  `json:"milestone"`
	StartDate      Date         `json:"startDate"`
	DueDate        Date         `json:"dueDate"`
	EstimatedHours float64      `json:"estimatedHours"`
	ActualHours    float64      `json:"actualHours"`
	ParentIssueId  uint64       `json:"parentIssueId"`
	CreatedUser    User         `json:"createdUser"`
	Created        Date         `json:"created"`
	UpdatedUser    User         `json:"updatedUser"`
	Updated        Date         `json:"updated"`
	CustomFields   []string     `json:"customFields"`
	Attachments    []Attachment `json:"attachments"`
	SharedFiles    []SharedFile `json:"sharedFiles"`
	Stars          []Star       `json:"stars"`
}

const (
	getIssuesPath      = "/api/v2/issues"
	getIssuesCountPath = "/api/v2/issues/count"
)

func (c *Client) GetIssues(query url.Values) ([]Issue, error) {
	return c.GetIssuesContext(context.Background(), query)
}

func (c *Client) GetIssuesContext(ctx context.Context, query url.Values) ([]Issue, error) {
	var issues []Issue

	path, err := c.root.Parse(getIssuesPath)
	if err != nil {
		return nil, err
	}

	response, err := c.getContext(ctx, path, query)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(response, &issues); err != nil {
		return nil, err
	}

	return issues, nil
}

func (c *Client) GetIssuesCount(query url.Values) (int, error) {
	return c.GetIssuesCountContext(context.Background(), query)
}

func (c *Client) GetIssuesCountContext(ctx context.Context, query url.Values) (int, error) {
	path, err := c.root.Parse(getIssuesCountPath)
	if err != nil {
		return -1, err
	}

	response, err := c.getContext(ctx, path, query)
	if err != nil {
		return -1, err
	}

	var count struct {
		Count int `json:"count"`
	}

	if err := json.Unmarshal(response, &count); err != nil {
		return -1, err
	}

	return count.Count, nil
}

func (c *Client) GetIssue(issueKeyOrId string) (Issue, error) {
	return c.GetIssueContext(context.Background(), issueKeyOrId)
}

func (c *Client) GetIssueContext(ctx context.Context, issueKeyOrId string) (Issue, error) {
	var issue Issue

	path, err := c.root.Parse(path.Join(getIssuesPath, issueKeyOrId))
	if err != nil {
		return issue, err
	}

	response, err := c.getContext(ctx, path, nil)
	if err != nil {
		return issue, err
	}
	if err := json.Unmarshal(response, &issue); err != nil {
		return issue, err
	}

	return issue, nil
}

func (c *Client) CreateIssue(query url.Values) (Issue, error) {
	return c.CreateIssueContext(context.Background(), query)
}

func (c *Client) CreateIssueContext(ctx context.Context, query url.Values) (Issue, error) {
	var issue Issue

	path, err := c.root.Parse(getIssuesPath)
	if err != nil {
		return issue, err
	}

	payload := bytes.NewBufferString(query.Encode())
	response, err := c.postContext(ctx, path, nil, payload)
	if err != nil {
		return issue, err
	}
	if err := json.Unmarshal(response, &issue); err != nil {
		return issue, err
	}

	return issue, nil
}

func (c *Client) UpdateIssue(issueKeyOrId string, query url.Values) (Issue, error) {
	return c.UpdateIssueContext(context.Background(), issueKeyOrId, query)
}

func (c *Client) UpdateIssueContext(ctx context.Context, issueKeyOrId string, query url.Values) (Issue, error) {
	var issue Issue

	path, err := c.root.Parse(path.Join(getIssuesPath, issueKeyOrId))
	if err != nil {
		return issue, err
	}

	payload := bytes.NewBufferString(query.Encode())
	response, err := c.patchContext(ctx, path, nil, payload)
	if err != nil {
		return issue, err
	}
	if err := json.Unmarshal(response, &issue); err != nil {
		return issue, err
	}

	return issue, nil
}

func (c *Client) DeleteIssue(issueKeyOrId string) (Issue, error) {
	return c.DeleteIssueContext(context.Background(), issueKeyOrId)
}

func (c *Client) DeleteIssueContext(ctx context.Context, issueKeyOrId string) (Issue, error) {
	var issue Issue

	path, err := c.root.Parse(path.Join(getIssuesPath, issueKeyOrId))
	if err != nil {
		return issue, err
	}

	response, err := c.deleteContext(ctx, path, nil)
	if err != nil {
		return issue, err
	}
	if err := json.Unmarshal(response, &issue); err != nil {
		return issue, err
	}

	return issue, nil
}
