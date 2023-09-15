package model

import "encoding/json"

type IntCodeNamePair struct {
	Code int    `json:"code"`
	Name string `json:"name"`
}

type IdNamePair struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type StringCodeNamePair struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (pair *IdNamePair) String() string {
	js, err := json.Marshal(pair)
	if err != nil {
		return err.Error()
	}
	return string(js)
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
