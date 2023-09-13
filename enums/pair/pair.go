package pair

type StringCodeNamePair struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func FindStringPairName(pairs []*StringCodeNamePair, code string, defaultVal string) string {
	if code == "" {
		return defaultVal
	}

	for _, pair := range pairs {
		if pair.Code == code {
			return pair.Name
		}
	}

	return defaultVal
}

type IntCodeNamePair struct {
	Code int    `json:"code"`
	Name string `json:"name"`
}

func FindIntPairName(pairs []*IntCodeNamePair, code int, defaultVal string) string {
	if code == 0 {
		return defaultVal
	}

	for _, pair := range pairs {
		if pair.Code == code {
			return pair.Name
		}
	}

	return defaultVal
}
