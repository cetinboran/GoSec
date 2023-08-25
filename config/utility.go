package config

import (
	"fmt"

	"github.com/cetinboran/gojson/gojson"
	"github.com/cetinboran/gosec/database"
	"github.com/cetinboran/gosec/myencode"
	"github.com/cetinboran/gosec/settings"
)

func setKey(userId int, secret string) {
	// Gelen secret'ı şifreliyoruz
	encrypedSecret, _ := myencode.Encrypt(settings.GetSecretForSecrets(), secret)

	// Database'den config Table'a erişiyoruz.
	configT := database.GosecDb.Tables["config"]

	// Yeni bilgileri yazıyoruz.
	newData := gojson.DataInit([]string{"secret"}, []interface{}{encrypedSecret}, configT)

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
