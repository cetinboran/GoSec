package inputhandler

// mainden gelen inputları buraya yolla. Error işlemleri parse işlemleri burada gerçekleşicek.

import (
	cla "github.com/cetinboran/goarg/CLA"

	"github.com/cetinboran/gosec/config"
	"github.com/cetinboran/gosec/key"
	"github.com/cetinboran/gosec/password"
	"github.com/cetinboran/gosec/register"
)

func SendInput(args []cla.Input) {

	// Arg'ın nerden geldiğini anlamamız için.
	ModeName := args[0].ModeName

	// Mode Name'e göre start çalışıyor.
	// Artık gerisi kendi klasöründe işlevi yapıcak.
	switch ModeName {
	case "register":
		register.Start(args)
		break
	case "config":
		config.Start(args)
		break
	case "key":
		key.Start(args)
		break
	case "password":
		password.Start(args)
		break
	}
}
