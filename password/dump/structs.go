package dump

import (
	"fmt"
	"os"

	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/gosec/password/read"
)

type Dump struct {
	UserId         int
	Out            string
	Secret         string
	Path           string
	SecretRequired bool
}

func DumpInit(userId int) *Dump {
	// default path
	return &Dump{UserId: userId, SecretRequired: read.GetSecretRequired(userId), Path: "./"}
}

func (d *Dump) TakeInput(args []cla.Input) {
	for _, i2 := range args {
		if i2.Argument == "out" {
			d.Out = i2.Value
		}

		if i2.Argument == "s" || i2.Argument == "secret" {
			d.Secret = i2.Value
		}

		if i2.Argument == "p" || i2.Argument == "path" {
			d.Path = i2.Value
		}
	}
}

func (d *Dump) HandleInput() {
	// Secret Gerekli
	if d.SecretRequired {
		if d.Out == "1" {
			if d.Secret != "" {
				Out(d)
			} else {
				fmt.Println(GetErrors(1))
				os.Exit(1)
			}
		}
	} else { // Secret gerekli değil sormuyoruz.
		if d.Out == "1" {
			Out(d)
		}
	}
}
