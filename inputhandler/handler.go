package inputhandler

// mainden gelen inputları buraya yolla. Error işlemleri parse işlemleri burada gerçekleşicek.

import (
	"fmt"

	cla "github.com/cetinboran/goarg/CLA"
)

func SendInput(args []cla.Input) {
	fmt.Println(args)
}
