package password

import (
	"fmt"

	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/gosec/global"
)

func Start(args []cla.Input) {
	userId := global.Auth(args)
	newPassword := PasswordInit(userId)
	newPassword.TakeInputs(args)
	newPassword.CheckInputs()
	newPassword.Save()

	fmt.Println(newPassword)
}
