package create

func GetErrors(errorId int) string {
	switch errorId {
	case 1:
		return "Password must be at least 5 characters"
	case 2:
		return "You must either enter a password with -p/--pass or turn on automatic password generation using --create option."
	case 3:
		return "Please do not use -p and --create at the same time."
	}

	return ""
}
