package utils

import (
	"net/mail"
	"regexp"
)

var mobileRegex = regexp.MustCompile(`^(?:\+?\d{1,3}[-.\s_]*)?(?:$?(\d{1,4})$?[-.\s_]*)?(\d{1,4}[-.\s_]*?){2,5}\d{1,9}$`)

func IsMobile(mobile string) bool {
	return mobileRegex.MatchString(mobile)
}

func IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
