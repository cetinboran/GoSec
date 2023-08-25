package dump

func GetErrors(errorId int) string {
	switch errorId {
	case 1:
		return "--secret input is required for --out option. Extra security.\nYou can disable it at config."
	case 2:
		return "Invalid Secret!"
	}

	return ""
}
