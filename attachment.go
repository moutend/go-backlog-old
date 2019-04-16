package backlog

type Attachment struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Size        int    `json:"size"`
	CreatedUser User   `json:"createdUser"`
	Created     Date   `json:"created"`
	UpdatedUser User   `json:"updatedUser"`
	Updated     Date   `json:"updated"`
}
