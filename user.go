package backlog

type User struct {
	Id           int    `json:"id"`
	UserId       string `json:"userId"`
	Name         string `json:"name"`
	RoleType     int    `json:"roleType"`
	Lang         string `json:"lang"`
	MailAddress  string `json:"mailAddress"`
	NulabAccount string `json:"nulabAccount"`
}
