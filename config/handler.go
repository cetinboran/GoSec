package config

import (
	"fmt"

	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/gosec/global"
)

func Start(args []cla.Input) {
	// Auth Check
	userId := global.Auth(args)
	fmt.Println("Welcome to The Config Mode")

	newConfig := ConfigInit(userId)
	newConfig.TakeInputs(args)
	newConfig.CheckInputs()
	newConfig.HandleInputs()
}
