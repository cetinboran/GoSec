package register

import (
	"fmt"
	"os"

	"github.com/cetinboran/gosec/structs"
)

type Register struct {
	Username   string
	Password   string
	Repassword string
	Secret     string
}

func RegisterInit() *Register {
	return &Register{}
}

// This will takes input form args.
func (r *Register) TakeInput(args []structs.Input) {
	for _, i2 := range args {
		if i2.Argument == "u" || i2.Argument == "user" {
			r.Username = i2.Value
		}

		if i2.Argument == "p" || i2.Argument == "pass" {
			r.Password = i2.Value
		}

		if i2.Argument == "cp" || i2.Argument == "cpass" {
			r.Repassword = i2.Value
		}

		if i2.Argument == "s" || i2.Argument == "secret" {
			r.Secret = i2.Value
		}
	}
}

func (r *Register) CheckInputs() {
	if len(r.Username) == 0 || len(r.Password) == 0 || len(r.Repassword) == 0 || len(r.Secret) == 0 {
		fmt.Println(GetErrors(5))
		os.Exit(5)
	}

	if len(r.Username) < 3 {
		fmt.Println(GetErrors(1))
		os.Exit(1)
	}

	if len(r.Password) < 3 {
		fmt.Println(GetErrors(2))
		os.Exit(2)
	}

	if r.Password != r.Repassword {
		fmt.Println(GetErrors(3))
		os.Exit(3)
	}

	validLength := []int{16, 24, 32}

	for _, v := range validLength {
		if len(r.Secret) != v {
			fmt.Println(GetErrors(4))
			os.Exit(4)
		} else {
			break
		}
	}
}

func (r *Register) Save() {
	// myDb := database.GosecDb

	// Tableları isimleriyle map olarak kaydet işin kolaşlaşsın
	// yani users: *table gibi.
}
