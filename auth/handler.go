package auth

import cla "github.com/cetinboran/goarg/CLA"

func Start(args []cla.Input, errors map[string]*cla.OptionError) {
	newAuth := AuthInit(errors)
	newAuth.TakeInputs(args)
	newAuth.HandleInputs()
}
