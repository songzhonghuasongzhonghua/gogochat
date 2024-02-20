package tool

import (
	"crypto/sha256"
	"fmt"
)

func HashPasswordSha256(pwd string) string {
	sha := sha256.New()
	sha.Write([]byte(pwd))
	return fmt.Sprintf("%x", sha.Sum(nil))
}

func ValidHashPassword(plainPwd, password string) bool {
	sha := sha256.New()
	sha.Write([]byte(plainPwd))
	return fmt.Sprintf("%x", sha.Sum(nil)) == password
}
