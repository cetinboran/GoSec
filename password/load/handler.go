package load

import (
	cla "github.com/cetinboran/goarg/CLA"
)

func Start(args []cla.Input, errors map[string]*cla.OptionError) {
	newLoad := LoadInit()
	newLoad.Errors = errors // setting errors.

	newLoad.TakeInputs(args)
	newLoad.HandleInputs()

}
