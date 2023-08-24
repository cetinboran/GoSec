package config

func GetErrors(errorId int) string {
	switch errorId {
	case 1:
		return "Key Length Must Be 16,24 or 32!"
	}

	return ""
}
