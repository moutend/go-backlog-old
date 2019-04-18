package backlog

import (
	"context"
	"encoding/json"
	"net/url"
	"path"
)

type Project struct {
	Id                                uint64 `json:"id"`
	ProjectKey                        string `json:"projectKey"`
	Name                              string `json:"name"`
	ChartEnabled                      bool   `json:"chartEnabled"`
	SubtaskingEnabled                 bool   `json:"subtaskingEnabled"`
	ProjectLeaderCanEditProjectLeader bool   `json:"projectLeaderCanEditProjectLeader"`
	TextFormattingRule                string `json:"textFormattingRule"`
	Archived                          bool   `json:"archived"`
}

const (
	getProjectsPath = "/api/v2/projects"
)

func (c *Client) GetProjects(query url.Values) ([]Project, error) {
	return c.GetProjectsContext(context.Background(), query)
}

func (c *Client) GetProjectsContext(ctx context.Context, query url.Values) ([]Project, error) {
	var projects []Project

	path, err := c.root.Parse(getProjectsPath)
	if err != nil {
		return nil, err
	}

	response, err := c.getContext(ctx, path, query)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(response, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func (c *Client) GetProject(projectKeyOrId string) (Project, error) {
	return c.GetProjectContext(context.Background(), projectKeyOrId)
}

func (c *Client) GetProjectContext(ctx context.Context, projectKeyOrId string) (project Project, err error) {
	path, err := c.root.Parse(path.Join(getProjectsPath, projectKeyOrId))
	if err != nil {
		return project, err
	}

	response, err := c.getContext(ctx, path, nil)
	if err != nil {
		return project, err
	}
	if err := json.Unmarshal(response, &project); err != nil {
		return project, err
	}

	return project, nil
}
