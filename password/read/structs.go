package read

import (
	"fmt"
	"os"
	"strconv"

	cla "github.com/cetinboran/goarg/CLA"
)

type Read struct {
	userId         int
	PasswordId     int
	Title          string
	Secret         string
	List           bool
	Open           bool
	Copy           bool
	SecretRequired bool
}

func ReadInit(userId int) *Read {
	return &Read{SecretRequired: GetSecretRequired(userId), userId: userId}
}

func (r *Read) TakeInputs(args []cla.Input) {
	for _, i2 := range args {
		if i2.Argument == "i" || i2.Argument == "id" {
			passwordId, err := strconv.Atoi(i2.Value)
			if err != nil {
				fmt.Println(GetErrors(1))
				os.Exit(1)
			}

			// Burada gelen id 0 dan küçük veya 0 ise hata döndür.
			r.PasswordId = passwordId
		}

		if i2.Argument == "t" || i2.Argument == "title" {
			r.Title = i2.Value
		}

		if i2.Argument == "list" {
			r.List = true
		}

		if i2.Argument == "open" {
			r.Open = true
		}

		if i2.Argument == "copy" {
			r.Copy = true
		}

		if i2.Argument == "s" || i2.Argument == "secret" {
			r.Secret = i2.Value
		}
	}
}

func (r *Read) HandleInputs() {
	if r.List {
		if r.PasswordId != 0 || r.Title != "" || r.Open || r.Copy || r.Secret != "" {
			fmt.Println(GetErrors(3))
			os.Exit(3)
		} else {
			List(r.userId)
		}
	}

	if r.Open {
		if r.PasswordId == 0 && r.Title == "" {
			fmt.Println(GetErrors(2))
			os.Exit(2)
		}
	}

	if r.Copy {
		if r.PasswordId == 0 && r.Title == "" {
			fmt.Println(GetErrors(4))
			os.Exit(4)
		} else {
			Copy(r.userId, r.PasswordId, r.Title, r.Secret, r.SecretRequired)
		}
	}
}
