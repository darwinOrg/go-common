package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
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

func Sha1HmacHexEncode(key string, content string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(content))
	return hex.EncodeToString(mac.Sum(nil))
}
