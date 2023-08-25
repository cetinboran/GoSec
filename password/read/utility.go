package read

import (
	"fmt"
	"os"
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

func Copy(r *Read) {
	// Password ve Config Table'ına eriştim.
	PasswordsT := database.GosecDb.Tables["password"]
	ConfigT := database.GosecDb.Tables["config"]

	if r.SecretRequired && r.Secret == "" {
		fmt.Println(GetErrors(5))
		os.Exit(5)
	}

	// Aradığım Passwordu table dan çektim passwordId veya title kullanarak
	passwordMap := make([]map[string]interface{}, 3)
	if r.PasswordId != 0 {
		passwordMap = PasswordsT.Find("passwordId", r.PasswordId)
	}
	if r.Title != "" {
		passwordMap = PasswordsT.Find("title", r.Title)
	}

	user := ConfigT.Find("userId", r.userId)

	// Config dosyasından user'ın secretını çektim
	// Bununla password'un şifresini kırıcaz.
	userSecret := user[0]["secret"].(string) // içindeki string olduğu için böyle yaparak string yaptım sonra byte a çevirdim diğer türlü hata alıyorum.

	// Sonra şifrelenmiş olan user secret'ı önce decode atıyoruz.
	decryptedUserSecret, _ := myencode.Decrypt(settings.GetSecretForSecrets(), userSecret)

	if r.SecretRequired {
		if len(r.Secret) != 16 && len(r.Secret) == 24 && len(r.Secret) == 32 {
			fmt.Println(GetErrors(6))
			os.Exit(6)
		}

		if r.Secret != decryptedUserSecret {
			fmt.Println(GetErrors(6))
			os.Exit(6)
		}

		cryptedPassword := passwordMap[0]["password"]
		decryptedPassword, _ := myencode.Decrypt([]byte(decryptedUserSecret), cryptedPassword.(string))

		fmt.Println("The Password: ", decryptedPassword)

		// Şifreyi koypalıyoruz.
		utilityies.CopyToClipboard(decryptedPassword)
		return
	}

	cryptedPassword := passwordMap[0]["password"]
	decryptedPassword, _ := myencode.Decrypt([]byte(decryptedUserSecret), cryptedPassword.(string))

	fmt.Println("The Password: ", decryptedPassword)

	// Şifreyi koypalıyoruz.
	utilityies.CopyToClipboard(decryptedPassword)
}

func Open(r *Read) {

}
