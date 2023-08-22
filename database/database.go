package database

import "github.com/cetinboran/gojson/gojson"

// Başka yerlerde DB ye erişmek için global oluşturdum.
var GosecDb gojson.Database

func DatabaseInit() {
	// Init DB
	GosecDb = gojson.CreateDatabase("gosecDB", "./")

	// Init Table
	UsersT := gojson.CreateTable("users")

	UsersT.AddProperty("userId", "int", "PK")
	UsersT.AddProperty("username", "string", "")
	UsersT.AddProperty("password", "string", "")
	UsersT.AddProperty("secret", "string", "")

	ConfigT := gojson.CreateTable("config")

	ConfigT.AddProperty("codeperm", "int", "")
	ConfigT.AddProperty("secret", "string", "")

	// Adds table to the Database
	GosecDb.AddTable(&UsersT)
	GosecDb.AddTable(&ConfigT)

	// Creates Database Files.
	GosecDb.CreateFiles()

	// UsersT.Save(gojson.DataInit([]string{"username", "password"}, []interface{}{"BORANBORAN", "1"}, &UsersT))

}
