package dump

import (
	"fmt"
	"os"
	"strings"

	"github.com/cetinboran/gosec/database"
	"github.com/cetinboran/gosec/myencode"
	"github.com/cetinboran/gosec/password/read"
	"github.com/cetinboran/gosec/settings"
)

func Out(d *Dump) {
	ConfigT := database.GosecDb.Tables["config"]
	passwords := read.GetPasswords(d.UserId)

	user := ConfigT.Find("userId", d.UserId)

	// Config dosyasından user'ın secretını çektim
	// Bununla password'un şifresini kırıcaz.
	userSecret := user[0]["secret"].(string) // içindeki string olduğu için böyle yaparak string yaptım sonra byte a çevirdim diğer türlü hata alıyorum.

	// Sonra şifrelenmiş olan user secret'ı önce decode atıyoruz.
	decryptedUserSecret, _ := myencode.Decrypt(settings.GetSecretForSecrets(), userSecret)

	// Eğer secret gerekli ise burada kontrol ediyoruz doğru değil ise hata yolluyoruz.
	if d.SecretRequired {
		if d.Secret != decryptedUserSecret {
			fmt.Println(GetErrors(2))
			os.Exit(2)
		}
	}

	// Şifreleri decrypt ediyorum ve geri password map'ine kaydediyorum
	for i, v := range passwords {
		decryptedPassword, _ := myencode.Decrypt([]byte(decryptedUserSecret), v["password"].(string))
		passwords[i]["password"] = decryptedPassword
	}

	filePath := d.Path + "out.txt"
	dataStr := mapSliceToString(passwords)

	err := os.WriteFile(filePath, []byte(dataStr), 0644)
	if err != nil {
		fmt.Println("File write error:", err)
		return
	}

	fmt.Println("The out.json file was created at " + filePath)
}

func mapSliceToString(data []map[string]interface{}) string {
	var sb strings.Builder

	for _, item := range data {
		sb.WriteString(mapToString(item))
		sb.WriteString("\n")
	}

	return sb.String()
}

func mapToString(m map[string]interface{}) string {
	var sb strings.Builder

	keys := []string{"passwordId", "userId", "title", "url", "password"}

	for _, key := range keys {
		value := m[key]
		sb.WriteString(fmt.Sprintf("%s : %v\n", key, value))
	}

	return sb.String()
}
