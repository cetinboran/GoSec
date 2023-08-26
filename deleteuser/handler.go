package deleteuser

import (
	cla "github.com/cetinboran/goarg/CLA"
)

func Start(args []cla.Input) {
	newDeleteUser := DeleteUserInit()
	newDeleteUser.TakeInputs(args)
	newDeleteUser.HandleInputs()
}
