package domains

import (
	"strings"
)

type Domain string

func (d Domain) String() string {
	return string(d)
}

const (
	Business Domain = "BUSINESS"
	Supplier Domain = "SUPPLIER"
	Platform Domain = "PLATFORM"
	Employee Domain = "EMPLOYEE"
)

var holder = initHolder()

func initHolder() map[string]Domain {
	m := make(map[string]Domain)
	m[strings.ToLower(string(Business))] = Business
	m[strings.ToLower(string(Supplier))] = Supplier
	m[strings.ToLower(string(Platform))] = Platform
	m[strings.ToLower(string(Employee))] = Employee
	return m
}

func (d Domain) Is(that Domain) bool {
	return Equal(d.String(), that)
}

func (d Domain) IsAny(that ...Domain) bool {
	if len(that) == 0 {
		return false
	}
	for _, e := range that {
		if d.Is(e) {
			return true
		}
	}
	return false
}

func (d Domain) IsNot(that Domain) bool {
	return !d.Is(that)
}

func Equal(domain string, target Domain) bool {
	if strings.EqualFold(domain, string(target)) {
		return true
	}
	return false
}

func NotEqual(domain string, target Domain) bool {
	return !Equal(domain, target)
}

func Parse(dm string) Domain {
	return holder[strings.ToLower(dm)]
}

func IsNotAccessible(this Domain, target Domain) bool {
	return !IsAccessible(this, target)
}

func IsAccessible(this Domain, target Domain) bool {
	switch target {
	case Business:
		if this.Is(Platform) || this.Is(Business) {
			return true
		}
	case Platform:
		if this.Is(Platform) {
			return true
		}
	case Supplier:
		if this.Is(Platform) || this.Is(Supplier) {
			return true
		}
	case Employee:
		if this.Is(Platform) || this.Is(Employee) {
			return true
		}
	}
	return false
}
