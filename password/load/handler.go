package load

import (
	"fmt"

	cla "github.com/cetinboran/goarg/CLA"
)

func Start(args []cla.Input, errors map[string]*cla.OptionError) {
	newLoad := LoadInit()
	newLoad.TakeInputs(args)

	fmt.Println(newLoad)
}
