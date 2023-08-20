package main

import (
	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/gosec/inputhandler"
)

func main() {
	// İlk başta config başla ordan secret almamız lazım.
	// Her jsona kayıt ettiğinde json var mı diye kontrol et yoksa oluştur.

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

	// Register Mode
	Register := cla.ModInit()
	Register.SetTitle("Register Mode")
	Register.SetExamples([]string{"Example 1", "Example 2"})
	Register.AddOption("-u, --user", false, "Enter Your Username", []string{"Wrong Username"})
	Register.AddOption("-p, --pass", false, "Enter Your Password", []string{"Wrong Password"})
	Register.AddOption("-cp, --cpass", false, "Confirm Your Password", []string{"Passwords do not match!"})
	Register.AddOption("-s, --secret", false, "Enter Your Secret", []string{"Key Length Must Be 16,24 or 32!"})

	// Config Mode
	Config := cla.ModInit()
	Config.SetTitle("Config Mode")
	Config.SetExamples([]string{"Example 1", "Example 2"})
	Config.AddOption("-k, --key", false, "Sets The Secret Key's Key", []string{"Key Length Must Be 16,24 or 32!"})

	// Key Mode
	Key := cla.ModInit()
	Key.SetTitle("Key Mode")
	Key.SetExamples([]string{"Example 1", "Example 2"})
	Key.AddOption("-setk,--setkey", false, "Sets The Manuel Secret Key!", []string{"This must be 16,24 or 32!"})
	Key.AddOption("-ck, --createkey", false, "Creates The Secret Key!", []string{"This must be 16,24 or 32!"})

	// Mode Init
	Setup.AddMode("login", &Login)
	Setup.AddMode("register", &Register)
	Setup.AddMode("config", &Config)
	Setup.AddMode("key", &Key)

	// Sets the global options.
	Setup.AddGlobalOption("--logout", true, "Logs out so no one can use it.", []string{"You are already logged"})

	// Automatic Usage
	Login.AutomaticUsage()
	Register.AutomaticUsage()
	Config.AutomaticUsage()
	Key.AutomaticUsage()

	// Setup.AutomaticUsage()

	args := Setup.Start()

	inputhandler.SendInput(args)
}
