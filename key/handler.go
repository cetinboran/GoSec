package key

import (
	"fmt"

	cla "github.com/cetinboran/goarg/CLA"
)

// Start'ın her şeyin başladığı yer olsun.
func Start(args []cla.Input) {
	//global.Auth(args)

	fmt.Println("Welcome to The Key Mode")

	newKey := KeyInit()
	newKey.TakeInput(args)
	newKey.CheckInputs()
	newKey.GenerateKey()

}
