package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func ToString(value *string) string {
	if value != nil {
		return *value
	}
	return ""
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
