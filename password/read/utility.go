package read

import (
	"fmt"
	"strings"

	"github.com/cetinboran/gosec/database"
)

func GetSecretRequired(userId int) bool {
	configT := database.GosecDb.Tables["config"]
	userSecretRequired := configT.Find("userId", userId)[0]["secretrequired"]

	return userSecretRequired.(bool)
}

func List(userId int) {
	passwords := getPasswords(userId)

	fmt.Println("\nYour passwords are listed below. Enter the initial ids for -i / --id. Other commands work with -i / --id")
	ListWriter(passwords)
}

// Password Writer
func ListWriter(Passwords []map[string]interface{}) {
	var builder strings.Builder

	for _, v := range Passwords {
		passwordID, passwordIDOK := v["passwordId"].(float64)
		title, titleOK := v["title"].(string)
		password, passwordOK := v["password"].(string)

		if passwordIDOK {
			builder.WriteString(fmt.Sprintf("%.0f", passwordID) + " ")
		}

		if titleOK {
			builder.WriteString(title + " ")
		}

		if passwordOK {
			builder.WriteString(password + " ")
		}

		builder.WriteString("\n")
	}

	fmt.Println(builder.String())
}

func getPasswords(userId int) []map[string]interface{} {
	passwordsT := database.GosecDb.Tables["password"]
	passwords := passwordsT.Find("userId", userId)

	return passwords
}
