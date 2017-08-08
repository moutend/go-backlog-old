package backlog

type Star struct {
	Id        int    `json:"id"`
	Comment   string `json:"comment"`
	URL       string `json:"url"`
	Title     string `json:"title"`
	Presenter User   `json:"presenter"`
	Created   Date   `json:"created"`
}
