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
)

type Client struct {
	token  string
	root   *url.URL
	logger *log.Logger
}

func New(spaceName, token string) (*Client, error) {
	var err error
	var root *url.URL

	if spaceName == "" {
		return nil, fmt.Errorf("space name is empty")
	}
	if root, err = url.Parse("https://" + spaceName + ".backlog.jp/api/v2/"); err != nil {
		return nil, err
	}

	logger := log.New(ioutil.Discard, "", log.LstdFlags)
	client := &Client{
		token,
		root,
		logger,
	}

	return client, nil
}

func (c *Client) getContext(ctx context.Context, endpoint *url.URL, query url.Values) (response []byte, err error) {
	var req *http.Request
	var res *http.Response

	if req, err = http.NewRequest("GET", endpoint.String(), nil); err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	req = req.WithContext(ctx)
	q := req.URL.Query()

	for key, value := range query {
		q.Add(key, value[0])
	}

	// The value of `apiKey` is always required.
	q.Add("apiKey", c.token)
	rawQuery := q.Encode()
	req.URL.RawQuery = rawQuery

	c.logger.Println("GET", endpoint)
	c.logger.Println("query parameter:", rawQuery)

	if res, err = httpClient.Do(req); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if response, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}
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

func (c *Client) patchContext(ctx context.Context, endpoint *url.URL, values url.Values) (response []byte, err error) {
	var req *http.Request
	var res *http.Response

	payload := bytes.NewBufferString(values.Encode())

	if req, err = http.NewRequest("PATCH", endpoint.String(), payload); err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	q := req.URL.Query()

	// The value of `apiKey` is always required.
	q.Add("apiKey", c.token)
	rawQuery := q.Encode()
	req.URL.RawQuery = rawQuery

	c.logger.Println("PATCH", endpoint)
	c.logger.Println(rawQuery)
	c.logger.Println(payload)

	if res, err = httpClient.Do(req); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if response, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}
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

func (c *Client) postContext(ctx context.Context, endpoint *url.URL, query url.Values, payload io.Reader) (response []byte, err error) {
	var req *http.Request
	var res *http.Response

	c.logger.Println(payload)
	c.logger.Println("POST", endpoint)

	if req, err = http.NewRequest("POST", endpoint.String(), payload); err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	q := req.URL.Query()

	for key, value := range query {
		q.Add(key, value[0])
	}

	// The value of `apiKey` is always required.
	q.Add("apiKey", c.token)
	rawQuery := q.Encode()
	req.URL.RawQuery = rawQuery
	c.logger.Println("query parameter:", rawQuery)

	if res, err = httpClient.Do(req); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if response, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}
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

func (c *Client) SetLogger(logger *log.Logger) {
	c.logger = logger
	c.logger.Println("set logger")

	return
}

func (c *Client) GetProjects(query url.Values) ([]*Project, error) {
	return c.GetProjectsContext(context.Background(), query)
}

func (c *Client) GetProjectsContext(ctx context.Context, query url.Values) ([]*Project, error) {
	var err error
	var response []byte
	var projects []*Project
	var path *url.URL

	if query == nil {
		query = url.Values{}
	}
	if path, err = c.root.Parse("./projects"); err != nil {
		return nil, err
	}
	if response, err = c.getContext(ctx, path, query); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(response, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func (c *Client) GetIssues(query url.Values) ([]*Issue, error) {
	return c.GetIssuesContext(context.Background(), query)
}

func (c *Client) GetIssuesContext(ctx context.Context, query url.Values) ([]*Issue, error) {
	var err error
	var response []byte
	var issues []*Issue
	var path *url.URL

	if query == nil {
		query = url.Values{}
	}
	if path, err = c.root.Parse("./issues"); err != nil {
		return nil, err
	}
	if response, err = c.getContext(ctx, path, query); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(response, &issues); err != nil {
		return nil, err
	}

	return issues, nil
}

func (c *Client) GetIssue(issueId int) (*Issue, error) {
	return c.GetIssueContext(context.Background(), issueId)
}

func (c *Client) GetIssueContext(ctx context.Context, issueId int) (*Issue, error) {
	var err error
	var response []byte
	var issue Issue
	var path *url.URL

	if path, err = c.root.Parse(fmt.Sprintf("./issues/%v", issueId)); err != nil {
		return nil, err
	}
	if response, err = c.getContext(ctx, path, nil); err != nil {
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

func (c *Client) SetIssue(issueId int, values url.Values) (*Issue, error) {
	return c.SetIssueContext(context.Background(), issueId, values)
}

func (c *Client) SetIssueContext(ctx context.Context, issueId int, values url.Values) (*Issue, error) {
	var err error
	var response []byte
	var issue Issue
	var path *url.URL

	if path, err = c.root.Parse(fmt.Sprintf("./issues/%v", issueId)); err != nil {
		return nil, err
	}
	if response, err = c.patchContext(ctx, path, values); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(response, &issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

func (c *Client) GetIssuesCount(ctx context.Context, query url.Values) (int, error) {
	return c.GetIssuesCount(context.Background(), query)
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

func (c *Client) GetStatuses() ([]*Status, error) {
	return c.GetStatusesContext(context.Background())
}

func (c *Client) GetStatusesContext(ctx context.Context) ([]*Status, error) {
	var err error
	var response []byte
	var statuses []*Status
	var path *url.URL

	if path, err = c.root.Parse("./statuses"); err != nil {
		return nil, err
	}
	if response, err = c.getContext(ctx, path, nil); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(response, &statuses); err != nil {
		return nil, err
	}

	return statuses, nil
}

func (c *Client) GetIssueTypes(projectId int) ([]*IssueType, error) {
	return c.GetIssueTypesContext(context.Background(), projectId)
}

func (c *Client) GetIssueTypesContext(ctx context.Context, projectId int) ([]*IssueType, error) {
	var err error
	var response []byte
	var issueTypes []*IssueType
	var path *url.URL

	if path, err = c.root.Parse(fmt.Sprintf("./projects/%v/issueTypes", projectId)); err != nil {
		return nil, err
	}
	if response, err = c.getContext(ctx, path, nil); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(response, &issueTypes); err != nil {
		return nil, err
	}

	return issueTypes, nil
}

func (c *Client) GetPriorities() ([]*Priority, error) {
	return c.GetPrioritiesContext(context.Background())
}

func (c *Client) GetPrioritiesContext(ctx context.Context) ([]*Priority, error) {
	var err error
	var response []byte
	var priorities []*Priority
	var path *url.URL

	if path, err = c.root.Parse("./priorities"); err != nil {
		return nil, err
	}
	if response, err = c.getContext(ctx, path, nil); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(response, &priorities); err != nil {
		return nil, err
	}

	return priorities, nil
}
