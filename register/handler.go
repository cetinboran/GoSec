package register

import (
	"fmt"

	"github.com/cetinboran/gosec/structs"
)

func Start(args []structs.Input) {
	newRegister := RegisterInit()
	newRegister.TakeInput(args)
	newRegister.CheckInputs()

	fmt.Println(newRegister)
}
