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
	Register.AddOption("-gen,--generate", false, "Sets Secret Key Automaticly")

	DeleteUser := cla.ModInit()
	DeleteUser.SetUsage("GoSec Delete Mode", "You can delete users in this field.", []string{"Example"})
	DeleteUser.AddOption("-p,--pass", false, "Enter the password of the user to be deleted")

	// Config Mode
	Config := cla.ModInit()
	Config.SetTitle("Config Mode")
	Config.SetExamples([]string{"Example 1", "Example 2"})
	Config.AddOption("-k, --key", false, "Sets The Secret Key")
	Config.AddOption("-req,--required", false, "Sets the Secret Required true / false.")

	// Key Mode
	Key := cla.ModInit()
	Key.SetTitle("Key Mode")
	Key.SetExamples([]string{"Example 1", "Example 2"})
	Key.AddOption("-gen, --generate", false, "Creates The Secret Key!")

	// Password Mode **********************************************************************
	Password := cla.ModInit()
	Password.SetTitle("Password Mode")
	Password.SetDescription("This mod allows you to save your passwords.\nYou can even generate an automatic password")
	Password.SetExamples([]string{"gosec password <mode name>"})

	Create := cla.ModInit()
	Create.SetTitle("Password's Create Mode")
	Create.SetDescription("You can save your passwords in this field.")
	Create.SetExamples([]string{"gosec password create -P <loginpassword> -t <title> -u <url> -p <password>", "gosec password create -P <loginpassword> --create"})
	Create.AddOption("-t,--title", false, "Enter the Title of your Password. Like Instagram etc.")
	Create.AddOption("-u,--url", false, "Enter where this password is being used")
	Create.AddOption("-p,--pass", false, "Enter the password")
	Create.AddOption("--generate", true, "Creates Password For You. 16 chars")

	Read := cla.ModInit()
	Read.SetUsage("Password's Read Mode", "You can read your password in this field.", []string{"Example"})
	Read.AddOption("-i,--id", false, "You can choose from password id")
	Read.AddOption("-t,--title", false, "You can choose from title of your password")
	Read.AddOption("-s, --secret", false, "Enter your secret.")
	Read.AddOption("--list", true, "Shows Your Password")
	Read.AddOption("--open", true, "Opens the url in browser.")
	Read.AddOption("--copy", true, "Copies the password")

	Dump := cla.ModInit()
	Dump.SetUsage("Password's Dump Mode", "You can dump all of your passwords in this field.", []string{"Example"})
	Dump.AddOption("-s,--secret", false, "Enter your secret. For Extra Security.")
	Dump.AddOption("--out", true, "Dumps All Of Your Passwords")
	Dump.AddOption("-p,--path", false, "Add your out path.")

	Delete := cla.ModInit()
	Delete.SetUsage("Password's Delete Mode", "You can delete password in this field.", []string{"Example"})
	Delete.AddOption("-i,--id", false, "You can choose from password id")
	Delete.AddOption("--all", true, "Deletes all passwords")

	Load := cla.ModInit()
	Load.SetUsage("Password's Load Mode", "You can load passwords.", []string{"Example"})
	Load.AddOption("-p,--path", false, "Enter the path of file.")
	Load.AddOption("--format", true, "Create example file.")

	// Load Add Error
	Load.AddError("-p,--path", []string{"Invalid Path", "Error While Reading File"})

	// Password Mode Init
	Password.AddMode("create", &Create)
	Password.AddMode("read", &Read)
	Password.AddMode("dump", &Dump)
	Password.AddMode("delete", &Delete)
	Password.AddMode("load", &Load)

	// Global Option For Mods
	Password.AddGlobalOption("-P", false, "Enter your password for using the program.")

	// Password Mode Automatic Usage
	Create.AutomaticUsage()
	Read.AutomaticUsage()
	Dump.AutomaticUsage()
	Load.AutomaticUsage()
	// Password Mods Usage **********************************************************************

	// Main Mode Init
	Setup.AddMode("register", &Register)
	Setup.AddMode("config", &Config)
	Setup.AddMode("key", &Key)
	Setup.AddMode("password", &Password)
	Setup.AddMode("deleteuser", &DeleteUser)

	// Sets the global options.
	Setup.AddGlobalOption("-P", false, "Enter your password for using the program.")

	// Automatic Usage
	Register.AutomaticUsage()
	Config.AutomaticUsage()
	Key.AutomaticUsage()
	Password.AutomaticUsage()
	DeleteUser.AutomaticUsage()
	Setup.AutomaticUsage()

	args, errors := Setup.Start()

	inputhandler.SendInput(args, errors)
}
