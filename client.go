package backlog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	token   string
	apiroot *url.URL
}

func New(spaceName, token string) (*Client, error) {
	var err error
	var apiroot *url.URL

	if spaceName == "" {
		return nil, fmt.Errorf("space name is empty")
	}
	if apiroot, err = url.Parse("https://" + spaceName + ".backlog.jp/api/v2/"); err != nil {
		return nil, err
	}

	client := &Client{
		token,
		apiroot,
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
	req.URL.RawQuery = q.Encode()

	if res, err = httpClient.Do(req); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if response, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}
	if res.StatusCode == 200 {
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
	req.URL.RawQuery = q.Encode()

	if res, err = httpClient.Do(req); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if response, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}
	if res.StatusCode == 200 {
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

func (c *Client) Issues(query url.Values) ([]*Issue, error) {
	return c.IssuesContext(context.Background(), query)
}

func (c *Client) IssuesContext(ctx context.Context, query url.Values) ([]*Issue, error) {
	var err error
	var response []byte
	var issues []*Issue
	var path *url.URL

	if query == nil {
		query = url.Values{}
	}
	if path, err = c.apiroot.Parse("./issues"); err != nil {
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

	if path, err = c.apiroot.Parse(fmt.Sprintf("./issues/%v", issueId)); err != nil {
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

func (c *Client) SetIssue(issueId int, values url.Values) (*Issue, error) {
	return c.SetIssueContext(context.Background(), issueId, values)
}

func (c *Client) SetIssueContext(ctx context.Context, issueId int, values url.Values) (*Issue, error) {
	var err error
	var response []byte
	var issue Issue
	var path *url.URL

	if path, err = c.apiroot.Parse(fmt.Sprintf("./issues/%v", issueId)); err != nil {
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
