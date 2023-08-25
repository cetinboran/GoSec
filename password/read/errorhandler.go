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
	}

	return ""
}
