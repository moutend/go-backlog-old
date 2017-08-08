package backlog

type SharedFile struct {
	d           int    `json:"id"`
	Type        string `json:"type"`
	Dir         string `json:"dir"`
	Name        string `json:"name"`
	Size        int    `json:"size"`
	CreatedUser User   `json:"createdUser"`
	Created     Date   `json:"created"`
	UpdatedUser User   `json:"updatedUser"`
	Updated     Date   `json:"updated"`
}
