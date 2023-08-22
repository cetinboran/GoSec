package register

func GetErrors(errorId int) string {
	switch errorId {
	case 1:
		return "Username must be at least 3 characters"
	case 2:
		return "Password must be at least 3 characters"
	case 3:
		return "Password is not the same as confirm password"
	case 4:
		return "Key Length Must Be 16,24 or 32!"
	case 5:
		return "You forgot to enter an argument"
	}

	return ""
}
