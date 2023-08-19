package inputhandler

// mainden gelen inputları buraya yolla. Error işlemleri parse işlemleri burada gerçekleşicek.

import (
	cla "github.com/cetinboran/goarg/CLA"

	"github.com/cetinboran/gosec/config"
	"github.com/cetinboran/gosec/key"
	"github.com/cetinboran/gosec/login"
	"github.com/cetinboran/gosec/register"
	"github.com/cetinboran/gosec/structs"
)

func SendInput(args []cla.Input) {
	convertedArgs := convertToMyInput(args)

	// Arg'ın nerden geldiğini anlamamız için.
	ModeName := convertedArgs[0].ModeName

	// Mode Name'e göre start çalışıyor.
	// Artık gerisi kendi klasöründe işlevi yapıcak.
	switch ModeName {
	case "login":
		login.Start(convertedArgs)
		break
	case "register":
		register.Start(convertedArgs)
		break
	case "config":
		config.Start(convertedArgs)
		break
	case "key":
		key.Start(convertedArgs)
		break
	}
}

func convertToMyInput(args []cla.Input) []structs.Input {
	var theArgsArr []structs.Input

	for _, arg := range args {
		theArgsArr = append(theArgsArr, structs.Input{Argument: arg.Argument, Value: arg.Value, Errors: arg.Error, ModeName: arg.ModeName})
	}

	return theArgsArr
}
