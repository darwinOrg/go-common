package dgsys

import (
	"github.com/darwinOrg/go-common/constants"
	"os"
)

var profile = ""

func GetProfile() string {
	if profile != "" {
		return profile
	}

	env, ok := os.LookupEnv(constants.Profile)
	if !ok || len(env) == 0 {
		return ""
	}
	profile = env

	return profile
}

func IsQa() bool {
	return GetProfile() == "qa"
}

func IsPre() bool {
	return GetProfile() == "pre"
}

func IsProd() bool {
	return GetProfile() == "prod"
}

func IsFormalProfile() bool {
	return IsQa() || IsPre() || IsProd()
}
