package utilityies

import (
	"crypto/md5"
	"encoding/hex"
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
