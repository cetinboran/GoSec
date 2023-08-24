package global

import (
	"fmt"
	"math"
	"os"

	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/gosec/database"
	"github.com/cetinboran/gosec/utilityies"
)

// Register mode hariç eğer -ps yani şifre girilmez ise ve doğru değil ise hata vermeliyiz. Karşm şifre gir diye o yüzden her yerde olacak şeyi buraya yazıcam

// Auth Control
func Auth(args []cla.Input) int {
	// Bu kısma gelmesi için kullıcının en az bir option girmesi lazım bu nedenle help in çalışmasına engel değil.

	theIndex := -1
	for i, v := range args {
		// P aradığım şifre option'u Maindeki globalden onu değiştirirsen bunu da değiştirmelisin.
		if v.Argument == "P" {
			theIndex = i
			break
		}
	}
	if theIndex == -1 {
		fmt.Println(GetErrors(1))
		os.Exit(1)
	}

	// User bilgilerinin olduğu table
	usersT := database.GosecDb.Tables["users"]

	// Kullanıcıdan aldığımız şifreyi md5 çevirip db de kontrol ediyoruz.
	md5_password := utilityies.ConvertToMd5(args[theIndex].Value)

	// Len 0 ise yoktur dolayısıyla error atıcaz.

	user := usersT.Find("password", md5_password)
	if len(user) == 0 {
		fmt.Println(GetErrors(2))
		os.Exit(2)
	}

	// Float64 ü int çevirir. Burada .(float64) gibi yazılımlar yani .(type) interfacelerde çalışıyor sadece
	intUserId := int(math.Floor(user[0]["userId"].(float64)))
	// Zaten programda kimsenin aynı şifresi olamıyor o yüzden böyle girebilirim

	return intUserId
}
