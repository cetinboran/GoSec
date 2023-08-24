package key

import (
	"fmt"
	"os"
	"strconv"

	cla "github.com/cetinboran/goarg/CLA"
	"github.com/cetinboran/gosec/utilityies"
)

type Key struct {
	Length int
}

func KeyInit() *Key {
	return &Key{}
}

func (k *Key) TakeInput(args []cla.Input) {
	for _, i2 := range args {
		if i2.Argument == "gen" || i2.Argument == "generate" {
			value, err := strconv.Atoi(i2.Value)

			// Int girmedilerse hata veriyoruz. Bu inputta
			if err != nil {
				fmt.Println(GetErrors(1))
				os.Exit(1)
			}

			k.Length = int(value)
		}
	}
}

func (k *Key) CheckInputs() {
	if k.Length != 16 && k.Length != 24 && k.Length != 32 {
		fmt.Println(GetErrors(2))
		os.Exit(2)
	}
}

func (k *Key) GenerateKey() {
	theSecret := utilityies.GenerateKey(k.Length)

	fmt.Println("The Generated Secret Key: ", theSecret)

	err := utilityies.CopyToClipboard(theSecret)
	if err != nil {
		fmt.Println(err)
	}

}
