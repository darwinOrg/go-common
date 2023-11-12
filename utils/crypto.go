package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
)

func Sha1Base64Encode(key string, content string) string {
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(content))
	bytes := hash.Sum(nil)

	return base64.StdEncoding.EncodeToString(bytes)
}
