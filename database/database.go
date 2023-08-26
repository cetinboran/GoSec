package database

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/cetinboran/gojson/gojson"
)

// Başka yerlerde DB ye erişmek için global oluşturdum.
var GosecDb gojson.Database

func DatabaseInit() {

	// Sets Path Of Database
	//targetDir := GetPath()

	// Init DB
	GosecDb = gojson.CreateDatabase("gosecDB", "./")

	// Users Table
	UsersT := gojson.CreateTable("users")
	UsersT.AddProperty("userId", "int", "PK")
	UsersT.AddProperty("username", "string", "")
	UsersT.AddProperty("password", "string", "")

	// Config Table
	ConfigT := gojson.CreateTable("config")
	ConfigT.AddProperty("configId", "int", "PK")
	ConfigT.AddProperty("userId", "int", "")
	ConfigT.AddProperty("secret", "string", "")
	ConfigT.AddProperty("secretrequired", "bool", "")

	// Passwords Table
	PasswordsT := gojson.CreateTable("password")
	PasswordsT.AddProperty("passwordId", "int", "PK")
	PasswordsT.AddProperty("userId", "int", "")
	PasswordsT.AddProperty("title", "string", "")
	PasswordsT.AddProperty("url", "string", "")
	PasswordsT.AddProperty("password", "string", "")

	// Adds table to the Database
	GosecDb.AddTable(&UsersT)
	GosecDb.AddTable(&ConfigT)
	GosecDb.AddTable(&PasswordsT)

	// Creates Database Files.
	GosecDb.CreateFiles()

	// UsersT.Save(gojson.DataInit([]string{"username", "password"}, []interface{}{"BORANBORAN", "1"}, &UsersT))

}

func GetPath() string {
	// Gets the file path.
	baseDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Hata:", err)
		return "./"
	}

	var targetDir string

	check := true
	switch runtime.GOOS {
	case "windows":
		fnames := []string{"Documents", "Belgeler"}

		// Bütün olasılıkları deniyorum documents, belgeler gibi
		for _, v := range fnames {
			targetDir = filepath.Join(baseDir, v)

			_, err = os.Stat(targetDir)
			if err != nil {
				// Böyle klasör yok
				check = false
				continue
			} else {
				check = true
				targetDir += "\\"
				break
			}
		}
		// Eğer hiçbir olasılık çalışmadıysa users'ın altına veritabanını yerleştiriyorum
		if !check {
			targetDir = "./"
		}
		break
	case "linux", "darwin":
		fnames := []string{"Documents", "Belgeler"}

		// Bütün olasılıkları deniyorum documents, belgeler gibi
		for _, v := range fnames {
			targetDir = filepath.Join(baseDir, v)

			_, err = os.Stat(targetDir)
			if err != nil {
				// Böyle klasör yok
				check = false
				continue
			} else {
				check = true
				targetDir += "/"
				break
			}
		}
		// Eğer hiçbir olasılık çalışmadıysa users'ın altına veritabanını yerleştiriyorum
		if !check {
			targetDir = "./"
		}
		break
	default:
		targetDir = "./"
		break
	}

	return targetDir
}
