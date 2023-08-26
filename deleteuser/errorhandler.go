package deleteuser

func GetErrors(errorId int) string {

	switch errorId {
	case 1:
		return "Invalid Password!"
	}

	return ""
}
