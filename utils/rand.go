package utils

import (
	cr "crypto/rand"
	"errors"
	"math/rand"
)

const (
	symbol = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~"
	letter = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

func RandomIntInRange(min, max int) int {
	return min + rand.Intn(max-min)
}

func MustRandomString(s string, length int) string {
	rs, _ := RandomString(s, length)
	return rs
}

func RandomString(s string, length int) (string, error) {
	var chars = []byte(s)
	clen := len(chars)
	if clen < 2 || clen > 256 {
		return "", errors.New("wrong charset length for chars length")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	// storage for random bytes.
	r := make([]byte, length+(length/4))
	i := 0

	for {
		if _, err := cr.Read(r); err != nil {
			return "", err
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b), nil
			}
		}
	}
}

func MustRandomSymbol(length int) string {
	rs, _ := RandomSymbol(length)
	return rs
}

func RandomSymbol(length int) (string, error) {
	return RandomString(symbol, length)
}

func MustRandomLetter(length int) string {
	rs, _ := RandomLetter(length)
	return rs
}

func RandomLetter(length int) (string, error) {
	return RandomString(letter, length)
}

func MustRandomBytes(length int) []byte {
	rb, _ := RandomBytes(length)
	return rb
}

func RandomBytes(length int) ([]byte, error) {
	b := make([]byte, length)
	_, err := cr.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
