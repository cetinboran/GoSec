package create

import (
	"fmt"
	"os"

	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/gojson/gojson"
	"github.com/cetinboran/gosec/database"
	"github.com/cetinboran/gosec/myencode"
	"github.com/cetinboran/gosec/settings"
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
		if i2.Argument == "generate" {
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

	ConfigT := database.GosecDb.Tables["config"]

	user := ConfigT.Find("userId", p.UserId)
	userSecret := user[0]["secret"].(string) // içindeki string olduğu için böyle yaparak string yaptım sonra byte a çevirdim diğer türlü hata alıyorum.

	// Sonra şifrelenmiş olan user secret'ı önce decode atıyoruz.
	decryptedUserSecret, _ := myencode.Decrypt(settings.GetSecretForSecrets(), userSecret)

	// interface olduğu için userSecret değeri böyle casting yapılıyor.
	// ardından decode edilmiş user secret ile şifreyi şifreliyoruz.
	cryptedPassword, _ := myencode.Encrypt([]byte(decryptedUserSecret), p.Password)

	// Eğer o title'dan başka var ise bu title kullanılıyor diyorum. Beliki şuanlık gereksiz olabilir ID den çekmek daha mantıklı
	if len(PasswordT.Find("title", p.Title)) != 0 {
		fmt.Println(GetErrors(4))
		os.Exit(4)
	}

	// passwordId yi db de pk yaptığım için otomatik ayarlanacak
	// sonra kayıt işlemi gerçekleştiriliyor.
	newData := gojson.DataInit([]string{"userId", "title", "url", "password"}, []interface{}{p.UserId, p.Title, p.Url, cryptedPassword}, PasswordT)

	PasswordT.Save(newData)
}
