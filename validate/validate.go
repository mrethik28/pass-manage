package validate

import (
	"crypto/md5"
	"encoding/hex"
)

func ValidateMasterKey(key string) bool {
	hash := md5.Sum([]byte(key))
	if hex.EncodeToString(hash[:]) == "5d41402abc4b2a76b9719d911017c592" {
		return true
	} else {
		return false
	}
}
