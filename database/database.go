package database

import "github.com/cetinboran/gojson/gojson"

func DatabaseInit() {
	// Init DB
	gosecDB := gojson.CreateDatabase("gosecDB", "./")

	// Init Table
	UsersT := gojson.CreateTable("users")

	UsersT.AddProperty("userId", "int", "PK")
	UsersT.AddProperty("username", "string", "")
	UsersT.AddProperty("password", "string", "")

	// Adds table to the Database
	gosecDB.AddTable(&UsersT)

	// Creates Database Files.
	gosecDB.CreateFiles()

}
