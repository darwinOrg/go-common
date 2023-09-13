package pay

import (
	"strings"
)

const (
	Supplier   string = "SUPPLIER"
	Contractor string = "CONTRACTOR"
)

var holder = initHolder()

func initHolder() map[string]string {
	m := make(map[string]string)
	m[strings.ToLower(Supplier)] = Supplier
	m[strings.ToLower(Contractor)] = Contractor
	return m
}

func Parse(target string) string {
	return holder[strings.ToLower(target)]
}
