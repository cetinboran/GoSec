package errorhandler

import "fmt"

func GetErrorLogin(errorId int, value string) string {
	switch errorId {
	case 1:
		return fmt.Sprintf("Wrong password or username: %v", value)
	}

	return ""
}
