package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
	"os"
)

func AesCBCEncrypt(key, content string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	contentBytes := []byte(content)
	blockSize := block.BlockSize()

	// Apply PKCS7 padding
	paddedContent := PKCS7Padding(contentBytes, blockSize)

	ciphertext := make([]byte, aes.BlockSize+len(paddedContent))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedContent)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func AesCBCDecrypt(key, content string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	if len(ciphertext)%block.BlockSize() != 0 {
		return "", fmt.Errorf("ciphertext length not aligned to block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintextPadded := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintextPadded, ciphertext)

	// Remove PKCS7 padding
	plaintext, err := PKCS7UnPadding(plaintextPadded)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padByte := byte(padding)
	paddingBytes := make([]byte, padding)
	for i := range paddingBytes {
		paddingBytes[i] = padByte
	}
	return append(data, paddingBytes...)
}

func PKCS7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, fmt.Errorf("data is empty")
	}

	padding := int(data[length-1])
	if padding > length || padding > aes.BlockSize {
		return nil, fmt.Errorf("invalid padding")
	}

	return data[:(length - padding)], nil
}

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
	return Md5HexBytes([]byte(content))
}

func Md5HexBytes(bytes []byte) string {
	h := md5.New()
	h.Write(bytes)

	return hex.EncodeToString(h.Sum(nil))
}

func Md5HexFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
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
