package config

import (
	"fmt"

	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/gosec/global"
)

func Start(args []cla.Input) {
	// Auth Check
	global.Auth(args)

	fmt.Println("Welcome to the config")

}
