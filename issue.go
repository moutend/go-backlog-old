package backlog

type Issue struct {
	Id             int       `json:"id"`
	ProjectID      int       `json:"projectId"`
	IssueKey       string    `json:"issueKey"`
	KeyId          int       `json:"keyId"`
	IssueType      IssueType `json:"issueType"`
	Summary        string    `json:"summary"`
	Description    string    `json:"description"`
	Resolution     string    `json:"resolution"`
	Priority       Priority  `json:"priority"`
	Status         Status    `json:"status"`
	Assignee       User      `json:"assignee"`
	Category       []string  `json:"category"`
	Versions       []string  `json:"versions"`
	Milestone      []string  `json:"milestone"`
	StartDate      Date      `json:"startDate"`
	DueDate        Date      `json:"dueDate"`
	EstimatedHours float64   `json:"estimatedHours"`
	ActualHours    float64   `json:"actualHours"`
	ParentIssueId  int       `json:"parentIssueId"`
	CreatedUser    User      `json:"createdUser"`
	Created        Date      `json:"created"`
	UpdateUser     User      `json:"updatedUser"`
	Updated        Date      `json:"updated"`
	CustomFields   []string  `json:"customFields"`
	Attachments    []string  `json:"attachments"`
	SharedFiles    []string  `json:"sharedFiles"`
	Stars          []Star    `json:"stars"`
}
