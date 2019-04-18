package backlog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	root   *url.URL
	token  string
	logger *log.Logger
}

func New(spaceName, token string) (*Client, error) {
	if spaceName == "" {
		return nil, fmt.Errorf("space name is empty")
	}
	if token == "" {
		return nil, fmt.Errorf("token is empty")
	}

	root, err := url.Parse("https://" + spaceName + ".backlog.jp")
	if err != nil {
		return nil, err
	}

	client := &Client{
		root:   root,
		token:  token,
		logger: log.New(ioutil.Discard, "", log.LstdFlags),
	}

	return client, nil
}

func (c *Client) doContext(ctx context.Context, method string, endpoint *url.URL, query url.Values, payload io.Reader) (response []byte, err error) {
	c.logger.Println(method, endpoint)

	// The value of 'apiKey' is always required.
	q := url.Values{}
	q.Add("apiKey", c.token)

	for k, vs := range query {
		for _, v := range vs {
			q.Add(k, v)
		}
	}

	endpoint.RawQuery = q.Encode()

	c.logger.Println("Query:", endpoint.RawQuery)
	c.logger.Println("Payload:", payload)

	req, err := http.NewRequest(method, endpoint.String(), payload)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	req = req.WithContext(ctx)

	if method == "POST" || method == "PATCH" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if response, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}

	c.logger.Println("Response:", string(response[:]))

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return response, nil
	}

	var errors Errors

	if err = json.Unmarshal(response, &errors); err != nil {
		return nil, err
	}
	if len(errors.Errors) == 0 {
		return nil, fmt.Errorf("error response is broken")
	}

	return nil, errors.Errors[0]
}

func (c *Client) getContext(ctx context.Context, endpoint *url.URL, query url.Values) (response []byte, err error) {
	return c.doContext(ctx, "GET", endpoint, query, nil)
}

func (c *Client) patchContext(ctx context.Context, endpoint *url.URL, query url.Values, payload io.Reader) (response []byte, err error) {
	return c.doContext(ctx, "PATCH", endpoint, query, payload)
}

func (c *Client) postContext(ctx context.Context, endpoint *url.URL, query url.Values, payload io.Reader) (response []byte, err error) {
	return c.doContext(ctx, "POST", endpoint, query, payload)
}

func (c *Client) deleteContext(ctx context.Context, endpoint *url.URL, query url.Values) (response []byte, err error) {
	return c.doContext(ctx, "DELETE", endpoint, query, nil)
}

func (c *Client) SetLogger(logger *log.Logger) {
	c.logger = logger
	c.logger.Println("set logger")

	return
}

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

func (c *Client) DeleteIssue(issueId int) (*Issue, error) {
	return c.DeleteIssueContext(context.Background(), issueId)
}

func (c *Client) DeleteIssueContext(ctx context.Context, issueId int) (*Issue, error) {
	var err error
	var response []byte
	var issue Issue
	var path *url.URL

	if path, err = c.root.Parse(fmt.Sprintf("./issues/%v", issueId)); err != nil {
		return nil, err
	}
	if response, err = c.deleteContext(ctx, path, nil); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(response, &issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

func (c *Client) CreateIssue(values url.Values) (*Issue, error) {
	return c.CreateIssueContext(context.Background(), values)
}

func (c *Client) CreateIssueContext(ctx context.Context, values url.Values) (*Issue, error) {
	var err error
	var response []byte
	var issue Issue
	var path *url.URL

	errorPrefix := "CreateIssueContext"
	payload := bytes.NewBufferString(values.Encode())

	if path, err = c.root.Parse("./issues"); err != nil {
		return nil, fmt.Errorf("%s: %s", errorPrefix, err)
	}
	if response, err = c.postContext(ctx, path, nil, payload); err != nil {
		return nil, fmt.Errorf("%s: %s", errorPrefix, err)
	}
	if err = json.Unmarshal(response, &issue); err != nil {
		return nil, fmt.Errorf("%s: %s", errorPrefix, err)
	}

	return &issue, nil
}

func (c *Client) SetIssue(issueId string, values url.Values) (*Issue, error) {
	return c.SetIssueContext(context.Background(), issueId, values)
}

func (c *Client) SetIssueContext(ctx context.Context, issueId string, values url.Values) (*Issue, error) {
	var err error
	var response []byte
	var issue Issue
	var path *url.URL

	payload := bytes.NewBufferString(values.Encode())

	if path, err = c.root.Parse(fmt.Sprintf("./issues/%v", issueId)); err != nil {
		return nil, err
	}
	if response, err = c.patchContext(ctx, path, nil, payload); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(response, &issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

func (c *Client) GetIssuesCount(query url.Values) (int, error) {
	return c.GetIssuesCountContext(context.Background(), query)
}

func (c *Client) GetIssuesCountContext(ctx context.Context, query url.Values) (int, error) {
	var err error
	var response []byte
	var count struct {
		Count int `json:"count"`
	}
	var path *url.URL

	if path, err = c.root.Parse("./issues/count"); err != nil {
		return 0, err
	}
	if response, err = c.getContext(ctx, path, query); err != nil {
		return 0, err
	}
	if err = json.Unmarshal(response, &count); err != nil {
		return 0, err
	}

	return count.Count, nil
}

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

func (c *Client) GetMyself() (myself User, err error) {
	return c.GetMyselfContext(context.Background())
}

func (c *Client) GetMyselfContext(ctx context.Context) (myself User, err error) {
	path, err := c.root.Parse(path.Join(getUsersPath, "myself"))
	if err != nil {
		return myself, err
	}

	response, err := c.getContext(ctx, path, nil)
	if err != nil {
		return myself, err
	}

	if err := json.Unmarshal(response, &myself); err != nil {
		return myself, err
	}

	return myself, nil
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

func (c *Client) GetPullRequestsCount(projectID, repositoryID string, query url.Values) (int, error) {
	return c.GetPullRequestsCountContext(context.Background(), projectID, repositoryID, query)
}

func (c *Client) GetPullRequestsCountContext(ctx context.Context, projectID, repositoryID string, query url.Values) (int, error) {
	var err error
	var response []byte
	var path *url.URL

	if query == nil {
		query = url.Values{}
	}
	if path, err = c.root.Parse(fmt.Sprintf("./projects/%v/git/repositories/%v/pullRequests/count", projectID, repositoryID)); err != nil {
		return -1, err
	}
	if response, err = c.getContext(ctx, path, query); err != nil {
		return -1, err
	}
	var count struct {
		Count int `json:"count"`
	}
	if err = json.Unmarshal(response, &count); err != nil {
		return -1, err
	}

	return count.Count, nil
}

func (c *Client) GetRepositories(projectKeyOrId string, query url.Values) ([]Repository, error) {
	return c.GetRepositoriesContext(context.Background(), projectKeyOrId, query)
}

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

func (c *Client) GetRepository(projectKeyOrId, repositoryNameOrId string, query url.Values) (Repository, error) {
	return c.GetRepositoryContext(context.Background(), projectKeyOrId, repositoryNameOrId, query)
}

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

func (c *Client) CreatePullRequest(projectId, repositoryId string, values url.Values) (*PullRequest, error) {
	return c.CreatePullRequestContext(context.Background(), projectId, repositoryId, values)
}

func (c *Client) CreatePullRequestContext(ctx context.Context, projectId, repositoryId string, values url.Values) (*PullRequest, error) {
	var err error
	var response []byte
	var pullRequest PullRequest
	var path *url.URL

	errorPrefix := "CreatePullRequestContext"
	payload := bytes.NewBufferString(values.Encode())

	if path, err = c.root.Parse(fmt.Sprintf("./projects/%v/git/repositories/%v/pullRequests", projectId, repositoryId)); err != nil {
		return nil, fmt.Errorf("%s: %s", errorPrefix, err)
	}
	if response, err = c.postContext(ctx, path, nil, payload); err != nil {
		return nil, fmt.Errorf("%s: %s", errorPrefix, err)
	}
	if err = json.Unmarshal(response, &pullRequest); err != nil {
		return nil, fmt.Errorf("%s: %s", errorPrefix, err)
	}

	return &pullRequest, nil
}

func (c *Client) UpdatePullRequest(projectId, repositoryId string, number int, values url.Values) (*PullRequest, error) {
	return c.UpdatePullRequestContext(context.Background(), projectId, repositoryId, number, values)
}

func (c *Client) UpdatePullRequestContext(ctx context.Context, projectId, repositoryId string, number int, values url.Values) (*PullRequest, error) {
	var err error
	var response []byte
	var pullRequest PullRequest
	var path *url.URL

	errorPrefix := "UpdatePullRequestContext"
	payload := bytes.NewBufferString(values.Encode())

	if path, err = c.root.Parse(fmt.Sprintf("./projects/%v/git/repositories/%v/pullRequests", projectId, repositoryId)); err != nil {
		return nil, fmt.Errorf("%s: %s", errorPrefix, err)
	}
	if response, err = c.patchContext(ctx, path, nil, payload); err != nil {
		return nil, fmt.Errorf("%s: %s", errorPrefix, err)
	}
	if err = json.Unmarshal(response, &pullRequest); err != nil {
		return nil, fmt.Errorf("%s: %s", errorPrefix, err)
	}

	return &pullRequest, nil
}

func (c *Client) GetUsers() ([]*User, error) {
	return c.GetUsersContext(context.Background())
}

func (c *Client) GetUsersContext(ctx context.Context) ([]*User, error) {
	var err error
	var response []byte
	var users []*User
	var path *url.URL

	if path, err = c.root.Parse("./users"); err != nil {
		return nil, err
	}
	if response, err = c.getContext(ctx, path, nil); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(response, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (c *Client) GetWikis(query url.Values) ([]Wiki, error) {
	return c.GetWikisContext(context.Background(), query)
}

func (c *Client) GetWikisContext(ctx context.Context, query url.Values) ([]Wiki, error) {
	var wikis []Wiki

	path, err := c.root.Parse(getWikisPath)
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

func (c *Client) GetWiki(wikiId uint64) (Wiki, error) {
	return c.GetWikiContext(context.Background(), wikiId)
}

func (c *Client) GetWikiContext(ctx context.Context, wikiId uint64) (wiki Wiki, err error) {
	path, err := c.root.Parse(path.Join(getWikisPath, fmt.Sprint(wikiId)))
	if err != nil {
		return wiki, err
	}

	response, err := c.getContext(ctx, path, nil)
	if err != nil {
		return wiki, err
	}
	if err := json.Unmarshal(response, &wiki); err != nil {
		return wiki, err
	}

	return wiki, nil
}
