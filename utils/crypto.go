package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
)

func Sha1Base64Encode(key string, content string) string {
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(content))
	bytes := hash.Sum(nil)

	return base64.StdEncoding.EncodeToString(bytes)
}

func Sha256Base64Encode(key string, content string) string {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(content))
	bytes := hash.Sum(nil)

	return base64.StdEncoding.EncodeToString(bytes)
}
