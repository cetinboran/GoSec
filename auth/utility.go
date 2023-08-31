package auth

import (
	"fmt"
	"os"

	"github.com/cetinboran/gosec/database"
	"github.com/cetinboran/gosec/utilityies"
)

func Check(a *Auth) {
	usersT := database.GosecDb.Tables["users"]

	if len(usersT.Find("password", utilityies.ConvertToMd5(a.Password))) != 0 {
		fmt.Println("This is valid password.")
		os.Exit(0)
	} else {
		fmt.Println("false")
		os.Exit(0)
		// a.Errors["-p"].GetErrors(1)
	}
}
