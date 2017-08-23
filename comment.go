package backlog

type Comment struct {
	Id            int            `json:"id"`
	Content       string         `json:"content"`
	Changelog     string         `json:"changeLog"`
	CreatedUser   User           `json:"createdUser"`
	Created       Date           `json:"created"`
	Updated       Date           `json:"updated"`
	Stars         []Star         `json:"stars"`
	Notifications []Notification `json:"notifications"`
}
