package delete

import (
	"github.com/cetinboran/gosec/database"
	"github.com/cetinboran/gosec/password/read"
)

func DeleteById(d *Delete) {
	read.CheckValidPasswordId(d.UserId, d.PasswordId)

	passwordT := database.GosecDb.Tables["password"]
	passwordT.Delete("passwordId", d.PasswordId)
}

func DeleteAll(d *Delete) {
	passwordT := database.GosecDb.Tables["password"]

	// userId alıyor bize bütün şifrelerinin passwordId lerini döndürüyor.
	ValidPasswordIds := read.GetValidPasswordId(d.UserId)

	// passwordId ye göre siliyoruz.
	for _, v := range ValidPasswordIds {
		passwordT.Delete("passwordId", v)
	}

}
