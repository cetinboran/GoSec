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
// => --list
func ListWriter(Passwords []map[string]interface{}) {
	var builder strings.Builder

	// Hizalama için kullanılacak genişlik değerleri ilerde bu değerleri otomatik bulsun.
	// mesela max url yi bulsun ve kendi urlsinden çıkarsın o değer kadar boşluk bıraksın.
	idWidth := 3
	titleWidth := 20
	urlWidth := 30
	usernameWidth := 20
	passwordWidth := 30

	// Başlık satırları
	builder.WriteString(fmt.Sprintf("%-*s | %-*s | %-*s | %-*s | %-*s\n", idWidth, "ID", titleWidth, "Title", urlWidth, "URL", usernameWidth, "Username", passwordWidth, "Password"))

	for _, v := range Passwords {
		passwordID, passwordIDOK := v["passwordId"].(float64)
		title, titleOK := v["title"].(string)
		username, usernameOK := v["username"].(string)
		password, passwordOK := v["password"].(string)
		url, urlOK := v["url"].(string)

		if passwordIDOK {
			builder.WriteString(fmt.Sprintf("%-*s | ", idWidth, fmt.Sprintf("%.0f", passwordID)))
		} else {
			builder.WriteString(fmt.Sprintf("%-*s | ", idWidth, ""))
		}

		if titleOK {
			builder.WriteString(fmt.Sprintf("%-*s | ", titleWidth, title))
		} else {
			builder.WriteString(fmt.Sprintf("%-*s | ", titleWidth, ""))
		}

		if urlOK {
			builder.WriteString(fmt.Sprintf("%-*s | ", urlWidth, url))
		} else {
			builder.WriteString(fmt.Sprintf("%-*s | ", urlWidth, ""))
		}

		if usernameOK {
			builder.WriteString(fmt.Sprintf("%-*s | ", usernameWidth, username))
		} else {
			builder.WriteString(fmt.Sprintf("%-*s | ", usernameWidth, ""))
		}

		if passwordOK {
			builder.WriteString(fmt.Sprintf("%-*s", passwordWidth, password))
		} else {
			builder.WriteString(fmt.Sprintf("%-*s", passwordWidth, ""))
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
	passwordMap := getPasswordsIdOrTitle(r.UserId, r.PasswordId, r.Title)

	user := ConfigT.Find("userId", r.UserId)

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

	passwordMap := getPasswordsIdOrTitle(r.UserId, r.PasswordId, r.Title)

	passwordUrl := passwordMap[0]["url"].(string)
	passwordUrl = strings.TrimSpace(passwordUrl)

	if !strings.HasPrefix(passwordUrl, "https://") {
		fmt.Println(GetErrors(9))
		os.Exit(9)
	}

	blacklistSuffix := []string{";", "//", "|", "||", "&&", "&"}

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

	fmt.Printf("successfully opened %v\n", passwordUrl)
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

// Turns Valid Password's Id
func GetValidPasswordId(userId int) []int {
	var passwordIds []int

	// userId veriyoruz bize şifrelerinin mapini döndürüyor.
	passwords := GetPasswords(userId)

	// int yaparak geri döndürüyoruz.
	for _, v := range passwords {
		passwordIds = append(passwordIds, int(v["passwordId"].(float64)))
	}
	return passwordIds
}

func getValidTitles(userId int) []string {
	var passwordIds []string
	passwords := GetPasswords(userId)

	for _, v := range passwords {
		passwordIds = append(passwordIds, v["title"].(string))
	}
	return passwordIds
}

func CheckValidPasswordId(userId int, PasswordId int) {
	validIds := GetValidPasswordId(userId)

	check := false
	for _, v := range validIds {
		if v != PasswordId {
			check = false
		} else {
			check = true
			break
		}
	}

	if !check {
		fmt.Println(GetErrors(11))
		os.Exit(11)
	}
}

// Finds valid titles for database search.
func checkValidTitle(userId int, title string) {
	// Eğer yanlış title atarsa kullanıcı yakalamak istiyoruz yoksa error yeriz.
	ValidTitles := getValidTitles(userId)

	check := false
	for _, v := range ValidTitles {
		if v != title {
			check = false
		} else {
			check = true
			break
		}
	}

	if !check {
		fmt.Println(GetErrors(12))
		os.Exit(12)
	}
}

func getPasswordsIdOrTitle(userId int, passwordId int, title string) []map[string]interface{} {
	// Title boş ise title kontrol ediyor.
	if passwordId != 0 && title == "" {
		CheckValidPasswordId(userId, passwordId)
	}

	if passwordId == 0 && title != "" {
		checkValidTitle(userId, title)
	}

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
