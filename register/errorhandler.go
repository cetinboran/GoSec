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
	case 6:
		return "This password is being used, please enter another password."
	case 7:
		return "This username is being used, please enter another username."
	case 8:
		return "You have to enter int. Not a string in -gen / --generate"
	case 9:
		return "You must choose one option --generate cannot be used simultaneously with -s / --secret"
	case 10:
		return "-gen / --generate setting only takes 16 24 or 32 as input"
	}

	return ""
}
