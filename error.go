package backlog

import "fmt"

type Error struct {
	Message  string `json:"message"`
	Code     int    `json:"code"`
	MoreInfo string `json:"moreInfo"`
}

func (e Error) Error() string {
	var errorType string

	switch e.Code {
	case 1:
		errorType += "InternalError"
	case 2:
		errorType += "LicenceError"
	case 3:
		errorType += "LicenceExpiredError"
	case 4:
		errorType += "AccessDeniedError"
	case 5:
		errorType += "UnauthorizedOperationError"
	case 6:
		errorType += "NoResourceError"
	case 7:
		errorType += "InvalidRequestError"
	case 8:
		errorType += "SpaceOverCapacityError"
	case 9:
		errorType += "ResourceOverflowError"
	case 10:
		errorType += "TooLargeFileError"
	case 11:
		errorType += "AuthenticationError"
	default:
		errorType += "UnexpectedError"
	}
	return fmt.Sprintf("%v (%v)", e.Message, errorType)
}
