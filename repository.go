package backlog

import (
	"context"
	"encoding/json"
	"net/url"
	"path"
)

// Repository represents the git repository.
type Repository struct {
	Id           uint64  `json:"id"`
	ProjectId    uint64  `json:"projectId"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	HookURL      *string `json:"hookUrl"`
	HTTPURL      string  `json:"httpUrl"`
	SSHURL       string  `json:"sshUrl"`
	DisplayOrder int     `json:"displayOrder"`
	PushedAt     *string `json:"pushedAt"`
	CreatedUser  User    `json:"createdUser"`
	Created      string  `json:"created"`
	UpdatedUser  User    `json:"createdUser"`
	Updated      string  `json:"created"`
}

// GetRepositories gets the git repositories.
func (c *Client) GetRepositories(projectKeyOrId string, query url.Values) ([]Repository, error) {
	return c.GetRepositoriesContext(context.Background(), projectKeyOrId, query)
}

// GetRepositoriesContext gets the git repositories.
func (c *Client) GetRepositoriesContext(ctx context.Context, projectKeyOrId string, query url.Values) ([]Repository, error) {
	var repositories []Repository

	path, err := c.root.Parse(path.Join(getProjectsPath, projectKeyOrId, "git", "repositories"))
	if err != nil {
		return nil, err
	}

	response, err := c.getContext(ctx, path, query)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(response, &repositories); err != nil {
		return nil, err
	}

	return repositories, nil
}

// GetRepository gets the repository.
func (c *Client) GetRepository(projectKeyOrId, repositoryNameOrId string, query url.Values) (Repository, error) {
	return c.GetRepositoryContext(context.Background(), projectKeyOrId, repositoryNameOrId, query)
}

// GetRepositoryContext gets the repository.
func (c *Client) GetRepositoryContext(ctx context.Context, projectKeyOrId, repositoryNameOrId string, query url.Values) (repository Repository, err error) {
	path, err := c.root.Parse(path.Join(
		getProjectsPath, projectKeyOrId,
		"git", "repositories", repositoryNameOrId,
	))
	if err != nil {
		return repository, err
	}

	response, err := c.getContext(ctx, path, query)
	if err != nil {
		return repository, err
	}
	if err := json.Unmarshal(response, &repository); err != nil {
		return repository, err
	}

	return repository, nil
}
