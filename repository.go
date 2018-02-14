package backlog

type Repository struct {
	Id           int     `json:"id"`
	ProjectId    int     `json:"projectId"`
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
