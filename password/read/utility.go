package read

import (
	"fmt"
	"strings"

	"github.com/cetinboran/gosec/database"
	"github.com/cetinboran/gosec/myencode"
	"github.com/cetinboran/gosec/settings"
	"github.com/cetinboran/gosec/utilityies"
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

func Copy(userId int, passwordId int, title string) {
	PasswordsT := database.GosecDb.Tables["password"]
	ConfigT := database.GosecDb.Tables["config"]

	passwordMap := make([]map[string]interface{}, 3)
	if passwordId != 0 {
		passwordMap = PasswordsT.Find("passwordId", passwordId)
	}

	if title != "" {
		passwordMap = PasswordsT.Find("title", title)
	}

	user := ConfigT.Find("userId", userId)
	userSecret := user[0]["secret"].(string) // içindeki string olduğu için böyle yaparak string yaptım sonra byte a çevirdim diğer türlü hata alıyorum.

	// Sonra şifrelenmiş olan user secret'ı önce decode atıyoruz.
	decryptedUserSecret, _ := myencode.Decrypt(settings.GetSecretForSecrets(), userSecret)

	cryptedPassword := passwordMap[0]["password"]

	decryptedPassword, _ := myencode.Decrypt([]byte(decryptedUserSecret), cryptedPassword.(string))

	fmt.Println("The Password: ", decryptedPassword)
	utilityies.CopyToClipboard(decryptedPassword)
}
