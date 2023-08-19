package main

import (
	"fmt"

	"github.com/cetinboran/gosec/database"
)

func main() {
	err := database.InitDB()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("DB connected")
}
