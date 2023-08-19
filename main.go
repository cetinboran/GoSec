package main

import (
	"fmt"

	cla "github.com/cetinboran/goarg/CLA"
)

func main() {
	// Setup ana main onda option şuanlık olmicak belki ilerde olur.
	// Ama ana düşüncem o sadece basıldığında modların helpini göstersin veya modları göstersin

	Setup := cla.Init()
	Setup.SetTitle("GoSec")

	// Login olunca true olsun jsonda ve onu kaydet. true ise diğer komutarı çalıştırmasına izin ver. logout atınca false yap çalıştıramasın başkası.
	// Login Mode
	Login := cla.ModInit()
	Login.SetTitle("Login Mode")
	Login.SetExamples([]string{"Example 1", "Example 2"})
	Login.AddOption("-u,--user", false, "Enter Your Username", []string{"Wrong Username"})
	Login.AddOption("-p, --pass", false, "Enter Your Password", []string{"Wrong Password"})
	Login.AutomaticUsage()

	// Register Mode
	Register := cla.ModInit()
	Register.SetTitle("Register Mode")
	Register.SetExamples([]string{"Example 1", "Example 2"})
	Register.AddOption("-u, --user", false, "Enter Your Username", []string{"Wrong Username"})
	Register.AddOption("-p, --pass", false, "Enter Your Password", []string{"Wrong Password"})
	Register.AddOption("-cp, --cpass", false, "Confirm Your Password", []string{"Passwords do not match!"})
	Register.AddOption("-s, --secret", false, "Enter Your Secret", []string{"Key Length Must Be 16,24 or 32!"})
	Register.AutomaticUsage()

	// Config Mode
	Config := cla.ModInit()
	Config.SetTitle("Config Mode")
	Config.SetExamples([]string{"Example 1", "Example 2"})
	Config.AddOption("-k, --key", false, "Enter The Secret Key's Key", []string{"Key Length Must Be 16,24 or 32!"})

	// Key Mode
	Key := cla.ModInit()
	Key.SetTitle("Key Mode")
	Key.SetExamples([]string{"Example 1", "Example 2"})
	Key.AddOption("-c,--create", true, "Creates Secret Key!", []string{"Key Length Must Be 16,24 or 32!"})
	Key.AutomaticUsage()

	// Mode Init
	Setup.AddMode("login", &Login)
	Setup.AddMode("register", &Register)
	Setup.AddMode("key", &Key)

	// Setup.AutomaticUsage()

	args := Setup.Start()

	fmt.Printf("args: %v\n", args)
}
