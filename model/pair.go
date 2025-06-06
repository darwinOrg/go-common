package model

import (
	"encoding/json"
)

type IntCodeNamePair struct {
	Code int    `json:"code" remark:"编码"`
	Name string `json:"name" remark:"名称"`
}

type IdNamePair struct {
	Id   int64  `json:"id,omitempty" remark:"id"`
	Name string `json:"name,omitempty" remark:"名称"`
}

type StringCodeNamePair struct {
	Code string `json:"code" remark:"编码"`
	Name string `json:"name" remark:"名称"`
}

type KeyValuePair[K any, V any] struct {
	Key   K `json:"key" remark:"键"`
	Value V `json:"value" remark:"值"`
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
