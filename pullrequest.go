package backlog

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"
	"path"
)

// PullRequest represents the pull request.
type PullRequest struct {
	Id           uint64  `json:"id"`
	ProjectId    uint64  `json:"projectId"`
	RepositoryId uint64  `json:"repositoryID"`
	Number       int     `json:"number"`
	Summary      string  `json:"summary"`
	Description  string  `json:"description"`
	Base         string  `json:"base"`
	Branch       string  `json:"branch"`
	Status       Status  `json:"status"`
	Assignee     User    `json:"assignee"`
	Issue        Issue   `json:"issue"`
	BaseCommit   string  `json:"baseCommit"`
	BranchCommit string  `json:"branchCommit"`
	CloseAt      *string `json:"closeAt"`
	MergeAt      string  `json:"mergeAt"`
	CreatedUser  User    `json:"createdUser"`
	Created      Date    `json:"created"`
	UpdatedUser  User    `json:"updatedUser"`
	Updated      Date    `json:"update"`
}

func (c *Client) GetPullRequests(projectKeyOrId, repositoryNameOrId string, query url.Values) ([]PullRequest, error) {
	return c.GetPullRequestsContext(context.Background(), projectKeyOrId, repositoryNameOrId, query)
}

func (c *Client) GetPullRequestsContext(ctx context.Context, projectKeyOrId, repositoryNameOrId string, query url.Values) ([]PullRequest, error) {
	var prs []PullRequest

	path, err := c.root.Parse(path.Join(
		getProjectsPath, projectKeyOrId,
		"git", "repositories", repositoryNameOrId,
		"pullRequests"))
	if err != nil {
		return nil, err
	}

	response, err := c.getContext(ctx, path, query)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(response, &prs); err != nil {
		return nil, err
	}

	return prs, nil
}

func (c *Client) GetPullRequestsCount(projectKeyOrId, repositoryNameOrId string, query url.Values) (int, error) {
	return c.GetPullRequestsCountContext(context.Background(), projectKeyOrId, repositoryNameOrId, query)
}

func (c *Client) GetPullRequestsCountContext(ctx context.Context, projectKeyOrId, repositoryNameOrId string, query url.Values) (int, error) {
	path, err := c.root.Parse(path.Join(
		getProjectsPath, projectKeyOrId,
		"git", "repositories", repositoryNameOrId,
		"pullRequests", "count",
	))
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

func (c *Client) GetPullRequest(projectKeyOrId, repositoryNameOrId, number string, query url.Values) (PullRequest, error) {
	return c.GetPullRequestContext(context.Background(), projectKeyOrId, repositoryNameOrId, number, query)
}

func (c *Client) GetPullRequestContext(ctx context.Context, projectKeyOrId, repositoryNameOrId, number string, query url.Values) (pr PullRequest, err error) {
	path, err := c.root.Parse(path.Join(
		getProjectsPath, projectKeyOrId,
		"git", "repositories", repositoryNameOrId,
		"pullRequests", number,
	))
	if err != nil {
		return pr, err
	}

	response, err := c.getContext(ctx, path, query)
	if err != nil {
		return pr, err
	}
	if err := json.Unmarshal(response, &pr); err != nil {
		return pr, err
	}

	return pr, nil
}

func (c *Client) CreatePullRequest(projectKeyOrId, repositoryNameOrId string, query url.Values) (PullRequest, error) {
	return c.CreatePullRequestContext(context.Background(), projectKeyOrId, repositoryNameOrId, query)
}

func (c *Client) CreatePullRequestContext(ctx context.Context, projectKeyOrId, repositoryNameOrId string, query url.Values) (PullRequest, error) {
	var pullRequest PullRequest

	path, err := c.root.Parse(path.Join(
		getProjectsPath, projectKeyOrId,
		"git", "repositories", repositoryNameOrId,
		"pullRequests",
	))
	if err != nil {
		return pullRequest, err
	}

	payload := bytes.NewBufferString(query.Encode())
	response, err := c.postContext(ctx, path, nil, payload)
	if err != nil {
		return pullRequest, err
	}
	if err := json.Unmarshal(response, &pullRequest); err != nil {
		return pullRequest, err
	}

	return pullRequest, nil
}

func (c *Client) UpdatePullRequest(projectKeyOrId, repositoryNameOrId, number string, query url.Values) (PullRequest, error) {
	return c.UpdatePullRequestContext(context.Background(), projectKeyOrId, repositoryNameOrId, number, query)
}

func (c *Client) UpdatePullRequestContext(ctx context.Context, projectKeyOrId, repositoryNameOrId, number string, query url.Values) (PullRequest, error) {
	var pullRequest PullRequest

	path, err := c.root.Parse(path.Join(
		getProjectsPath, projectKeyOrId,
		"git", "repositories", repositoryNameOrId,
		"pullRequests",
	))
	if err != nil {
		return pullRequest, err
	}

	payload := bytes.NewBufferString(query.Encode())
	response, err := c.patchContext(ctx, path, nil, payload)
	if err != nil {
		return pullRequest, err
	}
	if err := json.Unmarshal(response, &pullRequest); err != nil {
		return pullRequest, err
	}

	return pullRequest, nil
}
