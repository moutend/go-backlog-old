package backlog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
)

type Wiki struct {
	Id          uint64       `json:"id"`
	ProjectId   uint64       `json:"projectId"`
	Name        string       `json:"name"`
	Content     string       `json:"content"`
	Tags        []Tag        `json:"tags"`
	Attachments []Attachment `json:"attachment"`
	SharedFiles []SharedFile `json:"sharedFiles"`
	Stars       []Star       `json:"stars"`
	CreatedUser User         `json:"createdUser"`
	Created     Date         `json:"created"`
	UpdatedUser User         `json:"updatedUser"`
	Updated     Date         `json:"updated"`
}

const (
	getWikisPath = "/api/v2/wikis"
)

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
