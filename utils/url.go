package dgutils

import (
	"fmt"
	"github.com/darwinOrg/go-common/model"
	"net/url"
	"strings"
)

func FormUrlEncodedParams[V any](paramPairs []*model.KeyValuePair[string, V]) string {
	var paramArr []string
	for _, p := range paramPairs {
		if p.Key != "" {
			param := url.QueryEscape(p.Key) + "=" + url.QueryEscape(fmt.Sprintf("%v", p.Value))
			paramArr = append(paramArr, param)
		}
	}

	return strings.Join(paramArr, "&")
}
