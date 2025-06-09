package utils

import (
	"reflect"
)

func IsFieldsAllZero(obj any, excludeFields ...string) bool {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	excludeMap := make(map[string]bool)
	for _, field := range excludeFields {
		excludeMap[field] = true
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)

		if excludeMap[field.Name] {
			continue
		}

		value := v.Field(i)
		if value.Interface() != reflect.Zero(value.Type()).Interface() {
			return false
		}
	}

	return true
}

func FilterZeroFields(obj any, excludeFields ...string) []string {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	excludeMap := make(map[string]bool)
	for _, field := range excludeFields {
		excludeMap[field] = true
	}

	t := v.Type()
	var zeroFields []string

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)

		if excludeMap[field.Name] {
			continue
		}

		value := v.Field(i)
		if value.Interface() == reflect.Zero(value.Type()).Interface() {
			zeroFields = append(zeroFields, field.Name)
		}
	}

	return zeroFields
}

func ReflectAllFieldValuePointers(obj any) []any {
	val := reflect.ValueOf(obj).Elem()
	var ptrs []any
	for i := 0; i < val.NumField(); i++ {
		fieldPtr := val.Field(i).Addr().Interface()
		ptrs = append(ptrs, fieldPtr)
	}

	return ptrs
}
