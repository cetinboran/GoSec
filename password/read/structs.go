package read

import (
	"fmt"
	"os"
	"strconv"

	cla "github.com/cetinboran/goarg/CLA"
)

type Read struct {
	PasswordId int
	Secret     string
	List       bool
	Open       bool
	Copy       bool
}

func ReadInit() *Read {
	return &Read{}
}

func (r *Read) TakeInputs(args []cla.Input) {
	for _, i2 := range args {
		if i2.Argument == "i" || i2.Argument == "id" {
			passwordId, err := strconv.Atoi(i2.Value)
			if err != nil {
				fmt.Println(GetErrors(1))
				os.Exit(1)
			}

			r.PasswordId = passwordId
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
	}
}

func (r *Read) CheckInputs() {
	if r.PasswordId == 0 && (r.Open || r.Copy) {
		fmt.Println(GetErrors(2))
		os.Exit(1)
	}
}
