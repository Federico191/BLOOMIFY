package helper

import (
	"crypto/sha512"
	"encoding/hex"
)

func Hash512(input string) string {
	hash := sha512.New()
	hash.Write([]byte(input))
	pass := hex.EncodeToString(hash.Sum(nil))
	return pass
}
