package dump

import (
	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/gosec/global"
)

func Start(args []cla.Input) {
	userId := global.Auth(args)
	newDump := DumpInit(userId)
	newDump.TakeInput(args)
	newDump.HandleInput()

	// fmt.Println(newDump)
}
