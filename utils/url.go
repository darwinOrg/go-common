package utils

import (
	"fmt"
	"github.com/darwinOrg/go-common/model"
	"net/url"
)

func FormUrlEncodedParams[V any](paramPairs []*model.KeyValuePair[string, V]) string {
	uv := url.Values{}
	for _, p := range paramPairs {
		if p.Key != "" {
			uv.Set(p.Key, fmt.Sprintf("%v", p.Value))
		}
	}

	return uv.Encode()
}
