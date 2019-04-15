package backlog

type Comment struct {
	Id            uint64         `json:"id"`
	Content       string         `json:"content"`
	ChangeLog     []ChangeLog    `json:"changeLog"`
	CreatedUser   User           `json:"createdUser"`
	Created       Date           `json:"created"`
	Updated       Date           `json:"updated"`
	Stars         []Star         `json:"stars"`
	Notifications []Notification `json:"notifications"`
}
