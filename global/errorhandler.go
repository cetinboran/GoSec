package global

func GetErrors(errorId int) string {
	switch errorId {
	case 1:
		return "This part is forbidden. You must enter password with x to use"
	case 2:
		return "Wrong password. Try again"
	}

	return ""
}
