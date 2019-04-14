package backlog

type Project struct {
	Id                                uint64 `json:"id"`
	ProjectKey                        string `json:"projectKey"`
	Name                              string `json:"name"`
	ChartEnabled                      bool   `json:"chartEnabled"`
	SubtaskingEnabled                 bool   `json:"subtaskingEnabled"`
	ProjectLeaderCanEditProjectLeader bool   `json:"projectLeaderCanEditProjectLeader"`
	TextFormattingRule                string `json:"textFormattingRule"`
	Archived                          bool   `json:"archived"`
}
