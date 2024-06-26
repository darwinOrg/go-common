package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"math/big"
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

func Sha1Hex(key string, content string) string {
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(content))
	bytes := hash.Sum(nil)

	return hex.EncodeToString(bytes)
}

func Sha256Hex(key string, content string) string {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(content))
	bytes := hash.Sum(nil)

	return hex.EncodeToString(bytes)
}

func Md5Hex(content string) string {
	h := md5.New()
	h.Write([]byte(content))

	return hex.EncodeToString(h.Sum(nil))
}

// GenerateCryptoRandomString generates a random string for cryptographic usage.
func GenerateCryptoRandomString(n int, runes string) (string, error) {
	letters := []rune(runes)
	b := make([]rune, n)
	for i := range b {
		v, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		b[i] = letters[v.Int64()]
	}
	return string(b), nil
}

// CryptoUint64 returns cryptographic random uint64.
func CryptoUint64() (uint64, error) {
	var v uint64
	if err := binary.Read(rand.Reader, binary.LittleEndian, &v); err != nil {
		return 0, err
	}
	return v, nil
}
