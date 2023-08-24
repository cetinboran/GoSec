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
	UserId   float64 // Veriyi json dan çekiyoruz. Jsonda sayıları float64 yapıyor.
	Secret   string
	Password string
}

func ConfigInit(userId float64) *Config {
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
	}
}

func (c *Config) CheckInputs() {
	if len(c.Secret) != 16 && len(c.Secret) != 24 && len(c.Secret) != 32 {
		fmt.Println(GetErrors(1))
		os.Exit(1)
	}
}

func (c *Config) HandleInputs() {
	// Burada böyle kontrol edicem. Mesela code limit diye bir config varsa configde ve struct içine take ınputs ile geldiyse if içinde eğer codelimit != "" ise uyarı versin dicem.
	// Birden fazla fonk. olabileceği için utiltly.go açıp oraya koydum.
	if c.Secret != "" {
		var choice string

		// Burayı sonradan otomatikleştir şuanlık şifreleri kaydettiğin bir yer yok.
		fmt.Println("If you haven't already obtained the encrypted versions of all passwords from the dump mode, when the 'secret' changes,\nall passwords become unusable.")
		fmt.Print("Are you sure? (Y/N): ")
		fmt.Scan(&choice)

		if choice == "Y" || choice == "y" {
			setKey(c.UserId, c.Secret)
		}
	}
}

// This function just for register.
func CreateConfig(userId int, secret string) {
	myDb := database.GosecDb

	data := gojson.DataInit([]string{"userId", "secret"}, []interface{}{userId, secret}, myDb.Tables["config"])
	myDb.Tables["config"].Save(data)
}
