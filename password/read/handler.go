package read

import (
	"fmt"

	cla "github.com/cetinboran/goarg/CLA"
)

func Start(args []cla.Input) {
	// global.Auth(args)
	newRead := ReadInit()
	newRead.TakeInputs(args)
	newRead.CheckInputs()

	fmt.Println(args)
}
