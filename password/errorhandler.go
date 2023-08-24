package password

func GetErrors(errorId int) string {
	switch errorId {
	case 1:
		return "Password must be at least 3 characters"
	}

	return ""
}
