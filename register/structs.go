package register

import (
	"fmt"
	"os"

	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/gojson/gojson"
	"github.com/cetinboran/gosec/config"
	"github.com/cetinboran/gosec/database"
	"github.com/cetinboran/gosec/myencode"
	"github.com/cetinboran/gosec/settings"
	"github.com/cetinboran/gosec/utilityies"
)

type Register struct {
	Username   string
	Password   string
	Repassword string
	Secret     string
}

func RegisterInit() *Register {
	return &Register{}
}

// This will takes input form args.
func (r *Register) TakeInput(args []cla.Input) {
	for _, i2 := range args {
		if i2.Argument == "u" || i2.Argument == "user" {
			r.Username = i2.Value
		}

		if i2.Argument == "p" || i2.Argument == "pass" {
			r.Password = i2.Value
		}

		if i2.Argument == "cp" || i2.Argument == "cpass" {
			r.Repassword = i2.Value
		}

		if i2.Argument == "s" || i2.Argument == "secret" {
			r.Secret = i2.Value
		}
	}
}

func (r *Register) CheckInputs() {
	if len(r.Username) == 0 || len(r.Password) == 0 || len(r.Repassword) == 0 || len(r.Secret) == 0 {
		fmt.Println(GetErrors(5))
		os.Exit(5)
	}

	if len(r.Username) < 3 {
		fmt.Println(GetErrors(1))
		os.Exit(1)
	}

	if len(r.Password) < 3 {
		fmt.Println(GetErrors(2))
		os.Exit(2)
	}

	if r.Password != r.Repassword {
		fmt.Println(GetErrors(3))
		os.Exit(3)
	}

	validLength := []int{16, 24, 32}

	for _, v := range validLength {
		if len(r.Secret) != v {
			fmt.Println(GetErrors(4))
			os.Exit(4)
		} else {
			break
		}
	}
}

// Eğer böyle bir şifre var ise onu yapamazsın uyarısı versin.
func (r *Register) Save() {
	myDb := database.GosecDb

	UsersT := myDb.Tables["users"]
	md5_password := utilityies.ConvertToMd5(r.Password)

	// Şifre ile programı kullanacakları için 0 değil ise böyle bir şifre vardır dolayısıyla başka şifre gir uyarısı veriyoruz.
	if len(UsersT.Find("password", md5_password)) != 0 {
		fmt.Println(GetErrors(6))
		os.Exit(6)
	}

	if len(UsersT.Find("username", r.Username)) != 0 {
		fmt.Println(GetErrors(7))
		os.Exit(7)
	}

	data := gojson.DataInit([]string{"username", "password"}, []interface{}{r.Username, md5_password}, UsersT)
	UsersT.Save(data)
}

// you have to use this after you use Save Fnction.
func (r *Register) CreateConfig() {
	// Config içinde user ın secreti olucak şifreli bir şiekilde

	userId := len(database.GosecDb.Tables["users"].Get())

	// Zaten secret'ı kontrol edicem 16 24 veya 32 olsun diye o yüzden burda bakmıyorum
	cryptedSecret, _ := myencode.Encrypt(settings.GetSecretForSecrets(), r.Secret)
	config.CreateConfig(userId, cryptedSecret)
}
