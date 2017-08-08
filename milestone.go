package backlog

type Milestone struct {
	Id             int    `json:"id"`
	ProjectId      int    `json:"projectId"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	StartDate      Date   `json:"startDate"`
	ReleaseDueDate Date   `json:"releaseDueDate"`
	Archived       bool   `json:"archived"`
	DisplayOrder   int    `json:"displayOrder"`
}
