package config

import (
	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/gosec/global"
)

func Start(args []cla.Input) {
	// Auth Check
	userId := global.Auth(args)
	newConfig := ConfigInit(userId)
	newConfig.TakeInputs(args)
	newConfig.HandleInputs()
}
