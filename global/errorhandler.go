package global

func GetErrors(errorId int) string {
	switch errorId {
	case 1:
		return "This part is forbidden. You must enter password. So we can understand who u are."
	case 2:
		return "Wrong password. Try again"
	}

	return ""
}
