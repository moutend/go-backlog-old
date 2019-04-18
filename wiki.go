package backlog

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
