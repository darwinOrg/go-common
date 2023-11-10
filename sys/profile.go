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

func IsProd() bool {
	return GetProfile() == "prod"
}
