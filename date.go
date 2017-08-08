package backlog

import (
	"time"
)

// Date represents the date string. The form of date is "2017-08-08T00:56:08Z".
type Date string

func (d Date) Time() time.Time {
	t, _ := time.Parse(time.RFC3339, string(d))

	return t
}
