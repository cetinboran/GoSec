package key

func GetErrors(errorId int) string {
	switch errorId {
	case 1:
		return "You have to enter int. Not a string."
	case 2:
		return "Key Length Must Be 16,24 or 32!"
	}

	return ""
}
