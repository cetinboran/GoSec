package read

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cetinboran/gosec/database"
	"github.com/cetinboran/gosec/myencode"
	"github.com/cetinboran/gosec/settings"
	"github.com/cetinboran/gosec/utilityies"
)

func List(userId int) {
	passwords := GetPasswords(userId)

	fmt.Println("\nYour passwords are listed below. Enter the initial ids for -i / --id. Other commands work with -i / --id or -t / --title")
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

// Copies plain text password after security check
func Copy(r *Read) {
	// Password ve Config Table'ına eriştim.
	ConfigT := database.GosecDb.Tables["config"]

	if r.SecretRequired && r.Secret == "" {
		fmt.Println(GetErrors(5))
		os.Exit(5)
	}

	// Aradığım Passwordu table dan çektim passwordId veya title kullanarak
	passwordMap := getPasswordsIdOrTitle(r.userId, r.PasswordId, r.Title)

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
	// eğer otomatik giriş yapmasını istiyorsan "github.com/chromedp/chromedp" ile yapmak daha kolay.
	// Girilen url ye filter koymalısın garanti olsun.

	passwordMap := getPasswordsIdOrTitle(r.userId, r.PasswordId, r.Title)

	passwordUrl := passwordMap[0]["url"].(string)
	passwordUrl = strings.TrimSpace(passwordUrl)

	if !strings.HasPrefix(passwordUrl, "https://") {
		fmt.Println(GetErrors(9))
		os.Exit(9)
	}

	blacklistSuffix := []string{";", "//", "|", "||", "&&", "&", "/"}

	for _, v := range blacklistSuffix {
		if strings.HasSuffix(passwordUrl, v) {
			fmt.Println(GetErrors(10))
			os.Exit(10)
		}
	}

	cmd := exec.Command("cmd", "/c", "start", "chrome", passwordUrl) // Tarayıcı ve URL'yi burada belirtin
	err := cmd.Start()
	if err != nil {
		fmt.Println(GetErrors(8))
		os.Exit(8)
	}
}

// HELPER FUNCTIONS

func GetSecretRequired(userId int) bool {
	configT := database.GosecDb.Tables["config"]
	userSecretRequired := configT.Find("userId", userId)[0]["secretrequired"]

	return userSecretRequired.(bool)
}

func GetPasswords(userId int) []map[string]interface{} {
	passwordsT := database.GosecDb.Tables["password"]
	passwords := passwordsT.Find("userId", userId)

	return passwords
}

func getValidPasswordId(userId int) []int {
	var passwordIds []int
	passwords := GetPasswords(userId)

	for _, v := range passwords {
		passwordIds = append(passwordIds, int(v["passwordId"].(float64)))
	}
	return passwordIds
}

func checkValidPasswordId(userId int, PasswordId int) {
	validIds := getValidPasswordId(userId)

	for _, v := range validIds {
		if v != PasswordId {
			fmt.Println(GetErrors(11))
			os.Exit(11)
		}
	}
}

func getPasswordsIdOrTitle(userId int, passwordId int, title string) []map[string]interface{} {
	checkValidPasswordId(userId, passwordId)
	PasswordsT := database.GosecDb.Tables["password"]

	passwordMap := make([]map[string]interface{}, 3)
	if passwordId != 0 {
		passwordMap = PasswordsT.Find("passwordId", passwordId)
	}
	if title != "" {
		passwordMap = PasswordsT.Find("title", title)
	}

	return passwordMap
}
