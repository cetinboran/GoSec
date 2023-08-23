package config

import (
	"github.com/cetinboran/gojson/gojson"
	"github.com/cetinboran/gosec/database"
)

type Config struct {
	UserId    int
	CodeLimit int
	Secret    string
}

func ConfigInit() *Config {
	return &Config{}
}

// This function just for register.
func CreateConfig(userId int, secret string) {
	myDb := database.GosecDb

	data := gojson.DataInit([]string{"userId", "secret"}, []interface{}{userId, secret}, myDb.Tables["config"])
	myDb.Tables["config"].Save(data)
}
