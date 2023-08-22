package register

import (
	"github.com/cetinboran/gosec/structs"
)

func Start(args []structs.Input) {
	// Üye olduğumuzda otomatik bir şekilde bu userId ye uygun bir config dosyası oluşmalı.

	// Kayıt olan kişinin edineceği userId alttaki olucak. Normalde otomatik geliyor ancak bir config dosyası özel oluşaçağı için bize lazım.

	newRegister := RegisterInit()
	newRegister.TakeInput(args)
	newRegister.CheckInputs()
	newRegister.Save()
	newRegister.CreateConfig()
}
