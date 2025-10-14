package utils

import (
	"fmt"
	"net/url"

	"github.com/darwinOrg/go-common/model"
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

func MapToUrlQueryString[V any](params map[string]string) string {
	uv := url.Values{}
	for key, value := range params {
		uv.Set(key, value)
	}

	return uv.Encode()
}
