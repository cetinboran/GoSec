package key

import (
	cla "github.com/cetinboran/goarg/CLA"
)

// Start'ın her şeyin başladığı yer olsun.
func Start(args []cla.Input) {
	//global.Auth(args)
	newKey := KeyInit()
	newKey.TakeInput(args)
	newKey.CheckInputs()
	newKey.GenerateKey()
}
