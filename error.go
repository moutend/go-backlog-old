package backlog

type Error struct {
	Message  string `json:"message"`
	Code     int    `json:"code"`
	MoreInfo string `json:"moreInfo"`
}

func (e Error) Error() string {
	switch e.Code {
	case 1:
		return "InternalError"
	case 2:
		return "LicenceError"
	case 3:
		return "LicenceExpiredError"
	case 4:
		return "AccessDeniedError"
	case 5:
		return "UnauthorizedOperationError"
	case 6:
		return "NoResourceError"
	case 7:
		return "InvalidRequestError"
	case 8:
		return "SpaceOverCapacityError"
	case 9:
		return "ResourceOverflowError"
	case 10:
		return "TooLargeFileError"
	case 11:
		return "AuthenticationError"
	default:
		return "UnexpectedError"
	}
}
