package key

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

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
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	specialChars := "!@#$%^&*()_-+=[]{}|;:,.<>?/~"

	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	random := rand.New(source)

	secretKey := ""
	for i := 0; i < 16; i++ {
		if i < 4 {
			randIndex := random.Intn(len(lower))
			secretKey += string(lower[randIndex])
		} else if i < 8 {
			randIndex := random.Intn(len(upper))
			secretKey += string(upper[randIndex])
		} else if i < 12 {
			randIndex := random.Intn(len(numbers))
			secretKey += string(numbers[randIndex])
		} else {
			randIndex := random.Intn(len(specialChars))
			secretKey += string(specialChars[randIndex])
		}
	}

	// Secret'ı karıştırıyor.
	strArray := []rune(secretKey)
	for i := len(strArray) - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		strArray[i], strArray[j] = strArray[j], strArray[i]
	}

	shuffledKey := string(strArray)

	fmt.Println("The Generated Secret Key: ", shuffledKey)

	err := utilityies.CopyToClipboard(shuffledKey)
	if err != nil {
		fmt.Println(err)
	}

}
