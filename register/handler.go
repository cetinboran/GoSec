package register

import (
	"github.com/cetinboran/gosec/structs"
)

func Start(args []structs.Input) {
	newRegister := RegisterInit()
	newRegister.TakeInput(args)
	newRegister.CheckInputs()
	newRegister.Save()

	// Save attıktan sonra user için config dosyası oluşturduk.
	newRegister.CreateConfig()
}
