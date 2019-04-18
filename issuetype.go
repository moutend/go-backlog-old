package backlog

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
)

type IssueType struct {
	Id           uint64 `json:"id"`
	ProjectId    uint64 `json:"projectId"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	DisplayOrder int    `json:"displayOrder"`
}

func (c *Client) GetIssueTypes(projectId uint64) ([]IssueType, error) {
	return c.GetIssueTypesContext(context.Background(), projectId)
}

func (c *Client) GetIssueTypesContext(ctx context.Context, projectId uint64) ([]IssueType, error) {
	var issueTypes []IssueType

	path, err := c.root.Parse(path.Join(getProjectsPath, fmt.Sprint(projectId), "issueTypes"))
	if err != nil {
		return nil, err
	}

	response, err := c.getContext(ctx, path, nil)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(response, &issueTypes); err != nil {
		return nil, err
	}

	return issueTypes, nil
}
