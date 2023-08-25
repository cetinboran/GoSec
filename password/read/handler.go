package read

import (
	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/gosec/global"
)

func Start(args []cla.Input) {
	userId := global.Auth(args)
	newRead := ReadInit(userId)
	newRead.TakeInputs(args)
	newRead.HandleInputs(userId)

}
