package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func ConvertBeanToJsonString(obj any) (string, error) {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func MustConvertBeanToJsonString(obj any) string {
	jsonBytes, _ := json.Marshal(obj)
	return string(jsonBytes)
}

func MustConvertBeanToJsonStringPretty(obj any) string {
	jsonBytes, _ := json.MarshalIndent(obj, "", "	")
	return string(jsonBytes)
}

func ConvertJsonStringToBean[T any](str string) (*T, error) {
	if len(str) == 0 {
		return nil, nil
	}

	t := new(T)
	err := json.Unmarshal([]byte(str), t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func MustConvertJsonStringToBean[T any](str string) *T {
	if len(str) == 0 {
		return nil
	}

	t := new(T)
	_ = json.Unmarshal([]byte(str), t)
	return t
}

func ConvertJsonStringToBeanList[T any](str string) ([]*T, error) {
	return ConvertJsonStringToList[*T](str)
}

func ConvertJsonStringToList[T any](str string) ([]T, error) {
	if len(str) == 0 {
		return []T{}, nil
	}

	var t []T
	err := json.Unmarshal([]byte(str), &t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func MustConvertJsonStringToList[T any](str string) []T {
	if len(str) == 0 {
		return []T{}
	}

	var t []T
	_ = json.Unmarshal([]byte(str), &t)
	return t
}

func ConvertJsonBytesToBean[T any](bytes []byte) (*T, error) {
	if len(bytes) == 0 {
		return nil, nil
	}

	t := new(T)
	err := json.Unmarshal(bytes, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func MustConvertJsonBytesToBean[T any](bytes []byte) *T {
	if len(bytes) == 0 {
		return nil
	}

	t := new(T)
	_ = json.Unmarshal(bytes, t)
	return t
}

func ConvertJsonStringToMap(str string) (map[string]any, error) {
	return ConvertJsonBytesToMap([]byte(str))
}

func MustConvertJsonStringToMap(str string) map[string]any {
	mp, _ := ConvertJsonStringToMap(str)
	return mp
}

func ConvertJsonBytesToMap(bytes []byte) (map[string]any, error) {
	var mp map[string]any
	err := json.Unmarshal(bytes, &mp)
	if err != nil {
		return nil, err
	}
	return mp, nil
}

func MustConvertJsonBytesToMap(bytes []byte) map[string]any {
	mp, _ := ConvertJsonBytesToMap(bytes)
	return mp
}

func ConvertJsonBytesToBeanList[T any](bytes []byte) ([]*T, error) {
	return ConvertJsonBytesToList[*T](bytes)
}

func ConvertJsonBytesToList[T any](bytes []byte) ([]T, error) {
	if len(bytes) == 0 {
		return []T{}, nil
	}

	var t []T
	err := json.Unmarshal(bytes, &t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func ConvertToNewBeanByJson[T any](obj any) (*T, error) {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	t := new(T)
	err = json.Unmarshal(jsonBytes, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func MustConvertToNewBeanByJson[T any](obj any) *T {
	jsonBytes, _ := json.Marshal(obj)
	t := new(T)
	_ = json.Unmarshal(jsonBytes, t)
	return t
}

func ConvertToNewBeanListByJson[T any](obj any) ([]*T, error) {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var t []*T
	err = json.Unmarshal(jsonBytes, &t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func ConvertBeanToJsonStringWithoutEscaping(obj any) (string, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)

	err := encoder.Encode(obj)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func MustConvertBeanToJsonStringWithoutEscaping(obj any) string {
	str, err := ConvertBeanToJsonStringWithoutEscaping(obj)
	if err != nil {
		return ""
	}

	return str
}

func MergeJsonStrings(jsonStrings ...string) (string, error) {
	var mergedMap map[string]any

	for _, jsonString := range jsonStrings {
		var currentMap map[string]any
		if err := json.Unmarshal([]byte(jsonString), &currentMap); err != nil {
			fmt.Printf("error unmarshalling JSON string: %v", err)
			continue
		}

		if mergedMap == nil {
			mergedMap = currentMap
		} else {
			mergedMap = MergeMaps(mergedMap, currentMap)
		}
	}

	mergedJSON, err := json.Marshal(mergedMap)
	if err != nil {
		return "", fmt.Errorf("error marshalling merged map to JSON: %v", err)
	}

	return string(mergedJSON), nil
}
