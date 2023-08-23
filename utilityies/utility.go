package utilityies

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// Converts String To The Md5 String.
func ConvertToMd5(data string) string {
	// Create an MD5 hash object
	hash := md5.New()

	// Write the input string to the hash object
	hash.Write([]byte(data))

	// Get the MD5 hash as a byte slice
	hashBytes := hash.Sum(nil)

	// Convert the byte slice to a hexadecimal string
	hashString := hex.EncodeToString(hashBytes)

	return hashString
}

// CopyToClipBoard
func CopyToClipboard(text string) error {
	// Cmd açıyor ve işletim sistemine göre komut satırı ile kopyalama yapıyor.
	var cmd *exec.Cmd
	if runtime.GOOS == "darwin" {
		cmd = exec.Command("pbcopy")
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("xclip", "-selection", "clipboard")
	} else if runtime.GOOS == "windows" {
		cmd = exec.Command("clip")
	} else {
		return fmt.Errorf("Bu işletim sistemi desteklenmiyor")
	}

	cmd.Stdin = strings.NewReader(text)
	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("Successfully copied.")

	return nil
}
