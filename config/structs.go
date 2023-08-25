package config

import (
	"fmt"
	"os"

	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/gojson/gojson"
	"github.com/cetinboran/gosec/database"
)

// Config Id yazmadım otomatik eklenicek DB ye zaten bu struct kolaylık için
type Config struct {
	UserId            int // Veriyi json dan çekiyoruz. Jsonda sayıları float64 yapıyor.
	Secret            string
	Password          string
	SecretRequired    bool
	SetSecretRequired string
}

func ConfigInit(userId int) *Config {
	return &Config{UserId: userId}
}

func (c *Config) TakeInputs(args []cla.Input) {
	for _, i2 := range args {
		if i2.Argument == "k" || i2.Argument == "key" {
			c.Secret = i2.Value
		}

		// Bu config yerine şifreyle girileceği için -P yi de set etmesi gerekiyor. Bunun ile hangi user'ın secret'ini değiştireceğimizi öğreneceğiz.
		// Eğer P yi vermezse zaten buraya gelemez Authtan geçemez.
		if i2.Argument == "P" {
			c.Password = i2.Value
		}

		if i2.Argument == "req" || i2.Argument == "required" {
			if i2.Value == "true" || i2.Value == "True" {
				c.SetSecretRequired = i2.Value
			} else if i2.Value == "false" || i2.Value == "False" {
				c.SetSecretRequired = i2.Value
			} else {
				fmt.Println(GetErrors(3))
				os.Exit(3)
			}
		}
	}
}

func (c *Config) HandleInputs() {
	if c.Secret == "" && c.SetSecretRequired != "" {
		setSecretReq(c.UserId, c.SetSecretRequired)
	}

	if c.Secret != "" && c.SetSecretRequired == "" {
		var choice string

		if len(c.Secret) != 16 && len(c.Secret) != 24 && len(c.Secret) != 32 {
			fmt.Println(GetErrors(1))
			os.Exit(1)
		}

		// Burayı sonradan otomatikleştir şuanlık şifreleri kaydettiğin bir yer yok.
		fmt.Println("If you haven't already obtained the encrypted versions of all passwords from the dump mode, when the 'secret' changes,\nall passwords become unusable.")
		fmt.Print("Are you sure? (Y/N): ")
		fmt.Scan(&choice)

		if choice == "Y" || choice == "y" {
			setKey(c.UserId, c.Secret)
		}
	}

	if c.Secret != "" && c.SetSecretRequired != "" {
		fmt.Println(GetErrors(2))
		os.Exit(2)
	}

}

// This function just for register.
func CreateConfig(userId int, secret string) {
	myDb := database.GosecDb

	data := gojson.DataInit([]string{"userId", "secret", "secretrequired"}, []interface{}{userId, secret, true}, myDb.Tables["config"])
	myDb.Tables["config"].Save(data)
}
