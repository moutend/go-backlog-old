package backlog

type Wiki struct {
	Id          uint64 `json:"id"`
	ProjectId   uint64 `json:"projectId"`
	Name        string `json:"name"`
	Tags        []Tag  `json:"tags"`
	CreatedUser User   `json:"createdUser"`
	Created     Date   `json:"created"`
	UpdateUser  User   `json:"updatedUser"`
	Updated     Date   `json:"updated"`
}
