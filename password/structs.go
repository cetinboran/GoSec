package password

import (
	"fmt"
	"os"

	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/gojson/gojson"
	"github.com/cetinboran/gosec/database"
	"github.com/cetinboran/gosec/utilityies"
)

// Password Id yazmadım otomatik eklenicek DB ye zaten bu struct kolaylık için
type Password struct {
	UserId   int
	Title    string
	Url      string
	Password string
	Create   string
}

func PasswordInit(userId int) *Password {
	return &Password{UserId: userId}
}

func (p *Password) TakeInputs(args []cla.Input) {
	for _, i2 := range args {
		if i2.Argument == "t" || i2.Argument == "title" {
			p.Title = i2.Value
		}
		if i2.Argument == "u" || i2.Argument == "url" {
			p.Url = i2.Value
		}
		if i2.Argument == "p" || i2.Argument == "pass" {
			p.Password = i2.Value
		}
		if i2.Argument == "create" {
			p.Create = i2.Value
		}
	}
}

func (p *Password) CheckInputs() {
	// Bunların sırası önemli!

	if p.Password == "" && p.Create != "1" {
		fmt.Println(GetErrors(2))
		os.Exit(2)
	}

	if p.Password != "" && p.Create == "1" {
		fmt.Println(GetErrors(3))
		os.Exit(3)
	}

	if p.Password == "" && p.Create == "1" {
		p.Password = utilityies.GenerateKey(16)
	}

	if len(p.Password) < 5 {
		fmt.Println(GetErrors(1))
		os.Exit(1)
	}

	// Default Değerleri Değiştirdim.
	if p.Title == "" {
		p.Title = "Title"
	}

	if p.Url == "" {
		p.Url = "Url"
	}
}

func (p *Password) Save() {
	// userId yi global ın auth fonksiyonundan alıyoruz.
	PasswordT := database.GosecDb.Tables["password"]

	// passwordId yi db de pk yaptığım için otomatik ayarlanacak
	newData := gojson.DataInit([]string{"userId", "title", "url", "password"}, []interface{}{p.UserId, p.Title, p.Url, p.Password}, PasswordT)

	PasswordT.Save(newData)
}
