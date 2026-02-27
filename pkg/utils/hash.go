package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// HashString creates a hash of the input string
// This is used for generating unique identifiers
func HashString(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}
