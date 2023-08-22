package register

import (
	"fmt"

	"github.com/cetinboran/gosec/structs"
)

func Start(args []structs.Input) {
	// Üye olduğumuzda otomatik bir şekilde bu userId ye uygun bir config dosyası oluşmalı.

	// Kayıt olan kişinin edineceği userId alttaki olucak. Normalde otomatik geliyor ancak bir config dosyası özel oluşaçağı için bize lazım.
	// userId := len(database.GosecDb.Tables["users"].Get()) + 1

	newRegister := RegisterInit()
	newRegister.TakeInput(args)
	newRegister.CheckInputs()
	newRegister.Save()

	fmt.Println(args)
}
