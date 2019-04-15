package backlog

type Repository struct {
	Id           uint64  `json:"id"`
	ProjectId    uint64  `json:"projectId"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	HookURL      *string `json:"hookUrl"`
	HTTPURL      string  `json:"httpUrl"`
	SSHURL       string  `json:"sshUrl"`
	DisplayOrder int     `json:"displayOrder"`
	PushedAt     *string `json:"pushedAt"`
	CreatedUser  User    `json:"createdUser"`
	Created      string  `json:"created"`
	UpdatedUser  User    `json:"createdUser"`
	Updated      string  `json:"created"`
}
