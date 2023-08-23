package register

import cla "github.com/cetinboran/goarg/CLA"

func Start(args []cla.Input) {
	newRegister := RegisterInit()
	newRegister.TakeInput(args)
	newRegister.CheckInputs()
	newRegister.Save()

	// Save attıktan sonra user için config dosyası oluşturduk.
	newRegister.CreateConfig()
}
