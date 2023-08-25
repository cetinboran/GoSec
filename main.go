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
	Register.AddOption("-u, --user", false, "Enter Your Username")
	Register.AddOption("-p, --pass", false, "Enter Your Password")
	Register.AddOption("-cp, --cpass", false, "Confirm Your Password")
	Register.AddOption("-s, --secret", false, "Enter Your Secret")

	// Config Mode
	Config := cla.ModInit()
	Config.SetTitle("Config Mode")
	Config.SetExamples([]string{"Example 1", "Example 2"})
	Config.AddOption("-k, --key", false, "Sets The Secret Key")

	// Key Mode
	Key := cla.ModInit()
	Key.SetTitle("Key Mode")
	Key.SetExamples([]string{"Example 1", "Example 2"})
	Key.AddOption("-gen, --generate", false, "Creates The Secret Key!")

	// Password Mode **********************************************************************
	Password := cla.ModInit()
	Password.SetTitle("Password Mode")
	Password.SetDescription("This mod allows you to save your passwords.\nYou can even generate an automatic password")
	Password.SetExamples([]string{"gosec password create -P <password> --create"})

	Create := cla.ModInit()
	Create.SetTitle("Password's Create Mode")
	Create.SetDescription("You can save your passwords in this field.")
	Create.SetExamples([]string{"gosec password create -P <loginpassword> -t <title> -u <url> -p <password>", "gosec password create -P <loginpassword> --create"})
	Create.AddOption("-t,--title", false, "Enter the Title of your Password. Like Instagram etc.")
	Create.AddOption("-u,--url", false, "Enter where this password is being used")
	Create.AddOption("-p,--pass", false, "Enter the password")
	Create.AddOption("--generate", true, "Creates Password For You.")

	Read := cla.ModInit()
	Read.SetUsage("Password's Read Mode", "You can read your password in this field.", []string{"Example"})
	Read.AddOption("-i,--id", false, "You need to enter password id")
	Read.AddOption("-s, --secret", false, "Enter your secret.")
	Read.AddOption("--list", true, "Shows Your Password")
	Read.AddOption("--open", true, "Opens the url in browser.")
	Read.AddOption("--copy", true, "Copies the password")

	// Password Mode Init
	Password.AddMode("create", &Create)
	Password.AddMode("read", &Read)

	// Global Option For Mods
	Password.AddGlobalOption("-P", false, "Enter your password for using the program.")

	// Password Mode Automatic Usage
	Create.AutomaticUsage()
	Read.AutomaticUsage()
	// Password Mods Usage **********************************************************************

	// Main Mode Init
	Setup.AddMode("register", &Register)
	Setup.AddMode("config", &Config)
	Setup.AddMode("key", &Key)
	Setup.AddMode("password", &Password)

	// Sets the global options.
	Setup.AddGlobalOption("-P", false, "Enter your password for using the program.")

	// Automatic Usage
	Register.AutomaticUsage()
	Config.AutomaticUsage()
	Key.AutomaticUsage()
	Password.AutomaticUsage()
	Setup.AutomaticUsage()

	args := Setup.Start()

	inputhandler.SendInput(args)
}
