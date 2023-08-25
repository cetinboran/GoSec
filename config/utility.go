package config

import (
	"fmt"

	"github.com/cetinboran/gojson/gojson"
	"github.com/cetinboran/gosec/database"
	"github.com/cetinboran/gosec/myencode"
	"github.com/cetinboran/gosec/settings"
)

func changeTheSecretOfPasswords(userId int, newSecret string) {
	passwordT := database.GosecDb.Tables["password"]
	Passwords := passwordT.Find("userId", userId)

	// Bu userId ye bağllı bütün passwordları map içine attım. Şifresini kırıcam şimdi
	// BÜTÜN ŞİFRELERİ BURDA TUTUYORM
	mapPasswords := make(map[float64]string)

	for _, v := range Passwords {
		passwordId := v["passwordId"]
		password := v["password"]

		mapPasswords[passwordId.(float64)] = password.(string)
	}

	// Config table çekiyoruz
	ConfigT := database.GosecDb.Tables["config"]

	// Kullandığımız user'ı çekiyoruz
	user := ConfigT.Find("userId", userId)

	// User'ın şifrelenmiş secret'ını çekiyoruz
	userSecret := user[0]["secret"].(string)

	// Userın şifrelenmiş secret'ını kırıyoruz
	decryptedUserSecret, _ := myencode.Decrypt(settings.GetSecretForSecrets(), userSecret)

	// Userın secretı ile şifreleri kırıyoruz ve kaydediyoruz.
	for k, v := range mapPasswords {
		decryptedPassword, _ := myencode.Decrypt([]byte(decryptedUserSecret), v)
		mapPasswords[k] = decryptedPassword
	}

	// Şimdi bu şifreleri yeni secret ile şifreleyeceğiz.

	// burada yeni secret ile şifreliyoruz şimdi kaydedicez
	for k, v := range mapPasswords {
		ecryptedPassword, _ := myencode.Encrypt([]byte(newSecret), v)
		mapPasswords[k] = ecryptedPassword
	}

	for k, v := range mapPasswords {
		newData := gojson.DataInit([]string{"password"}, []interface{}{v}, passwordT)

		passwordT.Update("passwordId", k, newData)
	}
}

func setKey(userId int, secret string) {
	changeTheSecretOfPasswords(userId, secret)
	// Gelen secret'ı şifreliyoruz
	encrypedSecret, _ := myencode.Encrypt(settings.GetSecretForSecrets(), secret)

	// Database'den config Table'a erişiyoruz.
	configT := database.GosecDb.Tables["config"]

	// Yeni bilgileri yazıyoruz.
	newData := gojson.DataInit([]string{"secret", "secretrequired"}, []interface{}{encrypedSecret, true}, configT)

	// Table'a update atıyoruz.
	configT.Update("userId", userId, newData)

	fmt.Println("The secret key has been successfully changed.")
}

func setSecretReq(userId int, secretReqValue string) {
	configT := database.GosecDb.Tables["config"]

	boolValue := false

	if secretReqValue == "true" || secretReqValue == "True" {
		boolValue = true
	} else if secretReqValue == "false" || secretReqValue == "False" {
		boolValue = false
	}

	newData := gojson.DataInit([]string{"secretrequired"}, []interface{}{boolValue}, configT)

	// hata olmadığı halde değişmiyorsa o değişme tipi gojsonda yoktur büyük ihtimalle eklersin.
	configT.Update("userId", userId, newData)
}
