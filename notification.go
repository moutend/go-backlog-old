package backlog

type Notification struct {
	Id          int  `json:"id"`
	AlreadyRead bool `json:"alreadyRead"`
	Reason      int  `json:"reason"`
	User        User `json:"user"`
}
