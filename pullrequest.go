package backlog

type PullRequest struct {
	Id           uint64  `json:"id"`
	ProjectId    uint64  `json:"projectId"`
	RepositoryId uint64  `json:"repositoryID"`
	Number       int     `json:"number"`
	Summary      string  `json:"summary"`
	Description  string  `json:"description"`
	Base         string  `json:"base"`
	Branch       string  `json:"branch"`
	Status       Status  `json:"status"`
	Assignee     User    `json:"assignee"`
	Issue        Issue   `json:"issue"`
	BaseCommit   string  `json:"baseCommit"`
	BranchCommit string  `json:"branchCommit"`
	CloseAt      *string `json:"closeAt"`
	MergeAt      string  `json:"mergeAt"`
	CreateUser   User    `json:"createUser"`
	Created      string  `json:"created"`
	UpdateUser   User    `json:"updateUser"`
	Update       string  `json:"update"`
}
