package load

import (
	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/gosec/global"
)

func Start(args []cla.Input, errors map[string]*cla.OptionError) {
	userId := global.Auth(args)
	newLoad := LoadInit(userId)
	newLoad.Errors = errors // setting errors.

	newLoad.TakeInputs(args)
	newLoad.HandleInputs()

}
