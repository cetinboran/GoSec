package auth

import (
	"fmt"

	"github.com/cetinboran/gosec/database"
	"github.com/cetinboran/gosec/utilityies"
)

func Check(a *Auth) {
	usersT := database.GosecDb.Tables["users"]

	if len(usersT.Find("password", utilityies.ConvertToMd5(a.Password))) != 0 {
		fmt.Println("This is valid password.")
	} else {
		a.Errors["-p"].GetErrors(1)
	}
}
