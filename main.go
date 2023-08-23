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
	Config.AddOption("-cl, --codeLimit", false, "Set The Code Limit", []string{""})

	// Key Mode
	Key := cla.ModInit()
	Key.SetTitle("Key Mode")
	Key.SetExamples([]string{"Example 1", "Example 2"})
	Key.AddOption("-gen, --generate", false, "Creates The Secret Key!", []string{"This must be 16,24 or 32!"})

	// Mode Init
	Setup.AddMode("register", &Register)
	Setup.AddMode("config", &Config)
	Setup.AddMode("key", &Key)

	// Sets the global options.
	// Global optionlara dikkat et diğer optionlar ile çakışıyor
	Setup.AddGlobalOption("-P", false, "Enter your password for using the program.", []string{""})

	// Automatic Usage
	Register.AutomaticUsage()
	Config.AutomaticUsage()
	Key.AutomaticUsage()

	// Setup.AutomaticUsage()

	args := Setup.Start()

	inputhandler.SendInput(args)
}
