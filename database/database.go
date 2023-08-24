package database

import "github.com/cetinboran/gojson/gojson"

// Başka yerlerde DB ye erişmek için global oluşturdum.
var GosecDb gojson.Database

func DatabaseInit() {
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
