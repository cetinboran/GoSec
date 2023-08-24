package main

import (
	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/gosec/database"
	"github.com/cetinboran/gosec/inputhandler"
)

func main() {
	// Init Database JSON
	database.DatabaseInit()

	// İlk başta config başla ordan secret almamız lazım.
	// Her jsona kayıt ettiğinde json var mı diye kontrol et yoksa oluştur.

	// Setup ana main onda option şuanlık olmicak belki ilerde olur.
	// Ama ana düşüncem o sadece basıldığında modların helpini göstersin veya modları göstersin

	Setup := cla.Init()
	Setup.SetTitle("GoSec")
	Setup.SetDescription("A terminal-based password manager application that securely stores and manages passwords through simple commands.\nFeel free to create an account through the 'register' command. The password you set while signing up can also grant you access to other moderators.")

	// Login olunca true olsun jsonda ve onu kaydet. true ise diğer komutarı çalıştırmasına izin ver. logout atınca false yap çalıştıramasın başkası.

	// Register Mode
	Register := cla.ModInit()
	Register.SetTitle("Register Mode")
	Register.SetExamples([]string{"Example 1", "Example 2"})
	Register.AddOption("-u, --user", false, "Enter Your Username", []string{"Username must be at least 3 characters"})
	Register.AddOption("-p, --pass", false, "Enter Your Password", []string{"Password must be at least 3 characters"})
	Register.AddOption("-cp, --cpass", false, "Confirm Your Password", []string{"Password is not the same as confirm password"})
	Register.AddOption("-s, --secret", false, "Enter Your Secret", []string{"Key Length Must Be 16,24 or 32!"})

	// Config Mode
	Config := cla.ModInit()
	Config.SetTitle("Config Mode")
	Config.SetExamples([]string{"Example 1", "Example 2"})
	Config.AddOption("-k, --key", false, "Sets The Secret Key", []string{"Key Length Must Be 16,24 or 32!"})

	// Key Mode
	Key := cla.ModInit()
	Key.SetTitle("Key Mode")
	Key.SetExamples([]string{"Example 1", "Example 2"})
	Key.AddOption("-gen, --generate", false, "Creates The Secret Key!", []string{"This must be 16,24 or 32!"})

	// Password Mode
	Password := cla.ModInit()
	Password.SetTitle("Password Mode")
	Password.SetDescription("This mod allows you to save your passwords.\nYou can even generate an automatic password")
	Password.SetExamples([]string{"Example 1", "Example 2"})
	Password.AddOption("-t,--title", false, "Enter the Title of your Password. Like Instagram etc.", []string{})
	Password.AddOption("-u,--url", false, "Enter where this password is being used", []string{})
	Password.AddOption("-p,--pass", false, "Enter the password", []string{})
	Password.AddOption("--create", true, "Creates Password For You.", []string{})

	// Main Mode Init
	Setup.AddMode("register", &Register)
	Setup.AddMode("config", &Config)
	Setup.AddMode("key", &Key)
	Setup.AddMode("password", &Password)

	// Sets the global options.
	Setup.AddGlobalOption("-P", false, "Enter your password for using the program.", []string{""})

	// Automatic Usage
	Register.AutomaticUsage()
	Config.AutomaticUsage()
	Key.AutomaticUsage()
	Password.AutomaticUsage()
	Setup.AutomaticUsage()

	args := Setup.Start()

	inputhandler.SendInput(args)
}
