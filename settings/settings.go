package settings

import "github.com/cetinboran/gosec/database"

func GetSecretForSecrets() []byte {
	SettingsT := database.GosecDb.Tables["settings"]
	masterKey := SettingsT.Get()[0]["masterkey"].(string)

	return []byte(masterKey)
}
