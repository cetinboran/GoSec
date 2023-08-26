package delete

import (
	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/gosec/global"
)

func Start(args []cla.Input) {
	userId := global.Auth(args)

	newDelete := DeleteInit(userId)
	newDelete.TakeInputs(args)
	newDelete.HandleInputs()

}
