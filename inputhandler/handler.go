package inputhandler

// mainden gelen inputları buraya yolla. Error işlemleri parse işlemleri burada gerçekleşicek.

import (
	cla "github.com/cetinboran/goarg/CLA"

	"github.com/cetinboran/gosec/config"
	"github.com/cetinboran/gosec/deleteuser"
	"github.com/cetinboran/gosec/key"
	"github.com/cetinboran/gosec/password/create"
	"github.com/cetinboran/gosec/password/delete"
	"github.com/cetinboran/gosec/password/dump"
	"github.com/cetinboran/gosec/password/load"
	"github.com/cetinboran/gosec/password/read"
	"github.com/cetinboran/gosec/register"
)

func SendInput(args []cla.Input, errors map[string]*cla.OptionError) {

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
	case "create":
		create.Start(args)
		break
	case "read":
		read.Start(args)
		break
	case "dump":
		dump.Start(args)
		break
	case "delete":
		delete.Start(args)
		break
	case "deleteuser":
		deleteuser.Start(args)
		break
	case "load":
		load.Start(args, errors)
	}
}
