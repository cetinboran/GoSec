package delete

import (
	"fmt"
	"os"
	"strconv"

	cla "github.com/cetinboran/goarg/CLA"
)

type Delete struct {
	UserId     int
	PasswordId int
	All        string
}

func DeleteInit(userId int) *Delete {
	return &Delete{UserId: userId}
}

func (d *Delete) TakeInputs(args []cla.Input) {
	for _, i2 := range args {
		if i2.Argument == "i" || i2.Argument == "id" {
			passwordId, err := strconv.Atoi(i2.Value)
			if err != nil {
				fmt.Println(GetErrors(1))
				os.Exit(1)
			}

			// Burada gelen id 0 dan küçük veya 0 ise hata döndür.
			d.PasswordId = passwordId
		}

		if i2.Argument == "all" {
			d.All = "true"
		}
	}
}

func (d *Delete) HandleInputs() {
	if d.PasswordId != 0 && d.All == "true" {
		fmt.Println(GetErrors(2))
		os.Exit(2)
	}

	if d.PasswordId == 0 && d.All == "true" {
		DeleteAll(d)
	}

	if d.PasswordId != 0 && d.All == "" {
		DeleteById(d)
	}

}
