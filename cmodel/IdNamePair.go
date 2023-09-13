package cmodel

import "encoding/json"

type IdNamePair struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (pair *IdNamePair) String() string {
	js, err := json.Marshal(pair)
	if err != nil {
		return err.Error()
	}
	return string(js)
}
