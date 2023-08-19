package main

import (
	"fmt"

	cla "github.com/cetinboran/goarg/CLA"
)

func main() {
	Setup := cla.Init()
	Setup.SetTitle("GoSec")
	Setup.AddOption("-p,--password", false, "Enter your password.", []string{"TOO LOW"})

	Generate := cla.ModInit()
	Generate.SetTitle("Generate Mode")
	Generate.SetExamples([]string{"Example 1", "Example 2"})
	Generate.AddOption("--generate", true, "Generate Password!", []string{"This is error!"})

	Setup.AddMode("generate", &Generate)

	Setup.AutomaticUsage()
	Generate.AutomaticUsage()

	args := Setup.Start()

	fmt.Printf("args: %v\n", args)
}
