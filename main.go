package main

import (
	"fmt"

	cla "github.com/cetinboran/goarg/CLA"
)

func main() {
	// goarg modlar kullanırken hangi moddan bu inputun geldiğini söyleyebilir.
	Setup := cla.Init()
	Setup.SetTitle("GoSec")

	// Login Mode
	Login := cla.ModInit()
	Login.SetTitle("Login Mode")
	Login.SetExamples([]string{"Example 1", "Example 2"})
	Login.AddOption("-user,--username", false, "Enter Your Username", []string{"Wrong password or username"})
	Login.AutomaticUsage()

	// Key Mode
	Key := cla.ModInit()
	Key.SetTitle("Key Mode")
	Key.SetExamples([]string{"Example 1", "Example 2"})
	Key.AddOption("-c,--create", true, "Creates Secret Key!", []string{"Key Length Must Be 16,24 or 32!"})
	Key.AutomaticUsage()

	// Mode Init
	Setup.AddMode("key", &Key)
	Setup.AddMode("login", &Login)

	// Setup.AutomaticUsage()

	args := Setup.Start()

	fmt.Printf("args: %v\n", args)
}
