package backlog

import (
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
	root   *url.URL
	token  string
	logger *log.Logger
}
type requestModifyFunc func(*http.Request)

func New(spaceName, token string) (*Client, error) {
	if spaceName == "" {
		return nil, fmt.Errorf("space is empty")
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

func (c *Client) SetLogger(logger *log.Logger) {
	c.logger = logger
	c.logger.Println("set logger")

	return
}

func (c *Client) doContext(ctx context.Context, method string, endpoint *url.URL, query url.Values, payload io.Reader, modifyRequest requestModifyFunc) (response []byte, err error) {
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

	req = req.WithContext(ctx)
	modifyRequest(req)

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if response, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}

	c.logger.Println("Response:", string(response[:]))

	if res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusBadRequest {
		return response, nil
	}

	var errors Errors

	if err := json.Unmarshal(response, &errors); err != nil {
		return nil, err
	}
	if len(errors.Errors) == 0 {
		return nil, fmt.Errorf("error response is broken")
	}

	return nil, errors.Errors[0]
}

func (c *Client) getContext(ctx context.Context, endpoint *url.URL, query url.Values) (response []byte, err error) {
	return c.doContext(ctx, "GET", endpoint, query, nil, func(req *http.Request) {
		return
	})
}

func (c *Client) patchContext(ctx context.Context, endpoint *url.URL, query url.Values, payload io.Reader) (response []byte, err error) {
	return c.doContext(ctx, "PATCH", endpoint, query, payload, func(req *http.Request) {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	})
}

func (c *Client) postContext(ctx context.Context, endpoint *url.URL, query url.Values, payload io.Reader) (response []byte, err error) {
	return c.doContext(ctx, "POST", endpoint, query, payload, func(req *http.Request) {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	})
}

func (c *Client) deleteContext(ctx context.Context, endpoint *url.URL, query url.Values) (response []byte, err error) {
	return c.doContext(ctx, "DELETE", endpoint, query, nil, func(req *http.Request) {
		return
	})
}
