package read

func GetErrors(errorId int) string {
	switch errorId {
	case 1:
		return "You have to enter an integer in -i/--id"
	case 2:
		return "To use --open and --copy you must enter --id and --secret. Use config mode to turn off the --secret request"
	case 3:
		return "--list command can be used alone"
	}

	return ""
}
