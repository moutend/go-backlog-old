package backlog

type IssueType struct {
	Id           uint64 `json:"id"`
	ProjectId    int    `json:"projectId"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	DisplayOrder int    `json:"displayOrder"`
}
