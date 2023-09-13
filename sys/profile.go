package dgsys

import "os"

var profile = ""

func GetProfile() string {
	if profile != "" {
		return profile
	}

	env, ok := os.LookupEnv("profile")
	if !ok || len(env) == 0 {
		return ""
	}
	profile = env

	return profile
}

func IsProd() bool {
	return GetProfile() == "prod"
}
