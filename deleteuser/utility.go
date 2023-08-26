package deleteuser

import (
	"fmt"
	"math"
	"os"

	"github.com/cetinboran/gosec/database"
	deletepass "github.com/cetinboran/gosec/password/delete"
	"github.com/cetinboran/gosec/utilityies"
)

func FindUserIdByPassword(password string) int {
	// Bu fonksiyonları lazım olur diye yazdım.

	UserT := database.GosecDb.Tables["users"]

	// UserT içinde şifreler md5 olarak tutulduğu için böyle bakıyorum.
	user := UserT.Find("password", utilityies.ConvertToMd5(password))

	// Herkesin farklı şifre koyması zorunlu olduğudan bu kontrol yeter
	if len(user) == 0 {
		fmt.Println(GetErrors(1))
		os.Exit(1)
	}

	// İnt çeviriyoruz ve yolluyoruz.
	if userIdFloat, ok := user[0]["userId"].(float64); ok {
		userIdInt := int(math.Floor(userIdFloat))
		user[0]["userId"] = userIdInt
	}

	return user[0]["userId"].(int)
}

func CheckValidPassword(password string) {
	// Bu fonksiyonları lazım olur diye yazdım.

	UserT := database.GosecDb.Tables["users"]

	// UserT içinde şifreler md5 olarak tutulduğu için böyle bakıyorum.
	user := UserT.Find("password", utilityies.ConvertToMd5(password))

	// Herkesin farklı şifre koyması zorunlu olduğudan bu kontrol yeter
	if len(user) == 0 {
		fmt.Println(GetErrors(1))
		os.Exit(1)
	}
}

func UserDelete(d *DeleteUser) {
	ConfigT := database.GosecDb.Tables["config"]
	UsersT := database.GosecDb.Tables["users"]

	// Buradaki ise gojsonun içinden. Table içindekileri siliyor.
	ConfigT.Delete("userId", d.UserId)
	UsersT.Delete("userId", d.UserId)

	// Password'un delete modunu kullanarak bütün passwordları siliyorum.
	deletepass.DeleteAll(&deletepass.Delete{UserId: d.UserId})
}
