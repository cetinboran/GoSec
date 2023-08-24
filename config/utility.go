package config

import (
	"fmt"

	"github.com/cetinboran/gojson/gojson"
	"github.com/cetinboran/gosec/database"
	"github.com/cetinboran/gosec/encoding"
	"github.com/cetinboran/gosec/settings"
)

func setKey(userId float64, secret string) {
	// Gelen secret'ı şifreliyoruz
	encrypedSecret, _ := encoding.Encrypt(settings.GetSecretForSecrets(), secret)

	// Database'den config Table'a erişiyoruz.
	configT := database.GosecDb.Tables["config"]

	// Yeni bilgileri yazıyoruz.
	newData := gojson.DataInit([]string{"secret"}, []interface{}{encrypedSecret}, configT)

	// Table'a update atıyoruz.
	configT.Update("userId", userId, newData)

	fmt.Println("The secret key has been successfully changed.")
}
