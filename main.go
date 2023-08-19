package main

import (
	"fmt"

	cla "github.com/cetinboran/goarg/CLA"
)

func main() {
	Setup := cla.Init()
	Setup.SetTitle("GoSec")
	Setup.AddOption("-p,--password", false, "Enter your password.", []string{""})
	Setup.AutomaticUsage()
	args := Setup.Start()

	fmt.Printf("args: %v\n", args)
}
