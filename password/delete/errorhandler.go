package delete

func GetErrors(errorId int) string {
	switch errorId {
	case 1:
		return "You have to enter an integer in -i/--id"
	case 2:
		return "-i / --id and --all cannot be used at the same time"
	}

	return ""
}
