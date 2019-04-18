package backlog

import (
	"context"
	"encoding/json"
)

type Space struct {
	SpaceKey          string `json:"spaceKey"`
	Name              string `json:"name"`
	OwnerId           uint64 `json:"ownerId"`
	Lang              string `json:"lang"`
	ReportSendTime    string `json:"reportSendTime"`
	TextFormatingRule string `json:"textFormattingRule"`
	Timezone          string `json:"timezone"`
	Created           Date   `json:"created"`
	Updated           Date   `json:"updated"`
}

const (
	getSpacePath             = "/api/v2/space"
	getSpaceActivitiesPath   = "/api/v2/space/activities"
	getSpaceImagePath        = "/api/v2/space/image"
	getSpaceNotificationPath = "/api/v2/space/notification"
	getSpaceDiskUsagePath    = "/api/v2/space/diskUsage"
	getSpaceAttachmentPath   = "/api/v2/space/attachment"
)

func (c *Client) GetSpace() (Space, error) {
	return c.GetSpaceContext(context.Background())
}

func (c *Client) GetSpaceContext(ctx context.Context) (Space, error) {
	var space Space

	path, err := c.root.Parse(getSpacePath)
	if err != nil {
		return space, err
	}

	response, err := c.getContext(ctx, path, nil)
	if err != nil {
		return space, err
	}
	if err := json.Unmarshal(response, &space); err != nil {
		return space, err
	}

	return space, nil
}
