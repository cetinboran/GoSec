package read

func GetErrors(errorId int) string {
	switch errorId {
	case 1:
		return "You have to enter an integer in -i/--id"
	case 2:
		return "To use --open you must enter --id or --title"
	case 3:
		return "--list command can be used alone"
	case 4:
		return "To use --copy you must enter --id or --title and --secret"
	case 5:
		return "To use --copy you must enter --secret. This is for extra security\nYou can disable it at config."
	case 6:
		return "Invalid Secret!"
	case 7:
		return "To use --open you must enter --id or --title"
	case 8:
		return "Error while running chrome. Check if Chrome is installed"
	case 9:
		return "Url must start with https:// "
	case 10:
		return "Command injection detected. Log entry recorded."
	}

	return ""
}
