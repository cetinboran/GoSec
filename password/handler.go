package password

import (
	"fmt"

	cla "github.com/cetinboran/goarg/CLA"
)

func Start(args []cla.Input) {
	//global.Auth(args)
	newPassword := PasswordInit()

	fmt.Println(newPassword)
}
