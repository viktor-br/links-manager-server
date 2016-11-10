package security

import (
	"crypto/sha1"
	"fmt"
)

// Hash function hashing string with specified secret.
func Hash(str string, secret string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	hash.Write([]byte(secret))

	return fmt.Sprintf("%x", hash.Sum(nil))
}
