package backlog

import (
	"context"
	"encoding/json"
	"net/url"
	"path"
)

type User struct {
	Id           int    `json:"id"`
	UserId       string `json:"userId"`
	Name         string `json:"name"`
	RoleType     int    `json:"roleType"`
	Lang         string `json:"lang"`
	MailAddress  string `json:"mailAddress"`
	NulabAccount string `json:"nulabAccount"`
}

const (
	getUsersPath = "/api/v2/users"
)

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
