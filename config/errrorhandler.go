package config

func GetErrors(errorId int) string {
	switch errorId {
	case 1:
		return "Key Length Must Be 16,24 or 32!"
	case 2:
		return "-k/--key and -req/--required option is used alone"
	case 3:
		return "-req/--required option only takes true or false value"
	}

	return ""
}
