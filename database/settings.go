package database

import (
	"fmt"
	"os"
	"runtime"

	"github.com/cetinboran/gojson/gojson"
)

func SetSettings(SettingsT *gojson.Table) {
	if len(SettingsT.Get()) == 0 {
		for {
			var masterKey string
			fmt.Println("please enter the master secret key. This secret key encrypts your secret keys")
			fmt.Print(">: ")
			fmt.Scanln(&masterKey)

			if len(masterKey) != 16 && len(masterKey) != 24 && len(masterKey) != 32 {
				fmt.Println("The Master key length must be 16, 24 or 32")
			} else {
				SettingsT.Save(gojson.DataInit([]string{"masterkey"}, []interface{}{masterKey}, SettingsT))
				break
			}
		}
	}
}

func SetPath() string {
	baseDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Hata:", err)
	}

	switch runtime.GOOS {
	case "windows":
		baseDir += "\\"
	case "linux", "darwin":
		baseDir += "/"
		break
	}

	return baseDir
}
