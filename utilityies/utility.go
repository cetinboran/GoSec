package utilityies

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os/exec"
	"runtime"
	"strings"
	"time"
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

// Generates
func GenerateKey(length int) string {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	specialChars := "!@#$%^*()_-+=[]{}|;:,.<>?~"

	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	random := rand.New(source)

	var lengthForSecret int
	switch length {
	case 16:
		lengthForSecret = 4
		break
	case 24:
		lengthForSecret = 6
		break
	case 32:
		lengthForSecret = 8
		break
	}

	secretKey := ""
	for i := 0; i < length; i++ {
		if i < lengthForSecret {
			randIndex := random.Intn(len(lower))
			secretKey += string(lower[randIndex])
		} else if i < lengthForSecret+4 {
			randIndex := random.Intn(len(upper))
			secretKey += string(upper[randIndex])
		} else if i < lengthForSecret+8 {
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

	return shuffledKey
}
