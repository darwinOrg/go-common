package utils

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	mobileRegexStr   = `(\+?\d{1,3})?[-.\s_]?$?\d{1,4}$?[-.\s_]*\d{1,4}([-.\s_]*\d{1,4}){2,5}`
	mobileRegex      = regexp.MustCompile(mobileRegexStr)
	exactMobileRegex = regexp.MustCompile(fmt.Sprintf(`^%s$`, mobileRegexStr))

	emailRegex = regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
)

func IsMobile(mobile string) bool {
	return exactMobileRegex.MatchString(mobile)
}

func ContainsMobile(text string) bool {
	return mobileRegex.MatchString(text)
}

func ExtractMobiles(text string) []string {
	return mobileRegex.FindAllString(text, -1)
}

func IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ContainsEmail(text string) bool {
	return emailRegex.MatchString(text)
}

func ExtractEmails(text string) []string {
	return emailRegex.FindAllString(text, -1)
}
