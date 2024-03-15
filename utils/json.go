package utils

import "encoding/json"

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

func ConvertJsonStringToBean[T any](str string) (*T, error) {
	t := new(T)
	err := json.Unmarshal([]byte(str), t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func MustConvertJsonStringToBean[T any](str string) *T {
	t := new(T)
	_ = json.Unmarshal([]byte(str), t)
	return t
}

func ConvertJsonStringToBeanList[T any](str string) ([]*T, error) {
	return ConvertJsonStringToList[*T](str)
}

func ConvertJsonStringToList[T any](str string) ([]T, error) {
	var t []T
	err := json.Unmarshal([]byte(str), &t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func MustConvertJsonStringToList[T any](str string) []T {
	var t []T
	_ = json.Unmarshal([]byte(str), &t)
	return t
}

func ConvertJsonBytesToBean[T any](bytes []byte) (*T, error) {
	t := new(T)
	err := json.Unmarshal(bytes, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func MustConvertJsonBytesToBean[T any](bytes []byte) *T {
	t := new(T)
	_ = json.Unmarshal(bytes, t)
	return t
}

func ConvertJsonBytesToBeanList[T any](bytes []byte) ([]*T, error) {
	return ConvertJsonBytesToList[*T](bytes)
}

func ConvertJsonBytesToList[T any](bytes []byte) ([]T, error) {
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
