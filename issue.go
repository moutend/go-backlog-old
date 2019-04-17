package backlog

type Issue struct {
	Id             uint64       `json:"id"`
	ProjectId      uint64       `json:"projectId"`
	IssueKey       string       `json:"issueKey"`
	KeyId          uint64       `json:"keyId"`
	IssueType      IssueType    `json:"issueType"`
	Summary        string       `json:"summary"`
	Description    string       `json:"description"`
	Resolution     Resolution   `json:"resolution"`
	Priority       Priority     `json:"priority"`
	Status         Status       `json:"status"`
	Assignee       User         `json:"assignee"`
	Category       []Category   `json:"category"`
	Versions       []Version    `json:"versions"`
	Milestone      []Milestone  `json:"milestone"`
	StartDate      Date         `json:"startDate"`
	DueDate        Date         `json:"dueDate"`
	EstimatedHours float64      `json:"estimatedHours"`
	ActualHours    float64      `json:"actualHours"`
	ParentIssueId  uint64       `json:"parentIssueId"`
	CreatedUser    User         `json:"createdUser"`
	Created        Date         `json:"created"`
	UpdatedUser    User         `json:"updatedUser"`
	Updated        Date         `json:"updated"`
	CustomFields   []string     `json:"customFields"`
	Attachments    []Attachment `json:"attachments"`
	SharedFiles    []SharedFile `json:"sharedFiles"`
	Stars          []Star       `json:"stars"`
}
