package utils

import (
	"reflect"
)

func IsFieldsAllZero(obj any, excludeFields ...string) bool {
	if obj == nil {
		return true
	}

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
	if obj == nil {
		return []string{}
	}

	v := reflect.ValueOf(obj)
	for v.Kind() == reflect.Ptr {
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
	if obj == nil {
		return []any{}
	}

	tpe := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)

	for tpe.Kind() == reflect.Ptr {
		tpe = tpe.Elem()
		val = val.Elem()
	}

	ptrs := make([]any, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		fieldPtr := val.Field(i).Addr().Interface()
		ptrs = append(ptrs, fieldPtr)
	}

	return ptrs
}

func ReflectAllFieldValues(obj any) []any {
	if obj == nil {
		return []any{}
	}

	tpe := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)

	for tpe.Kind() == reflect.Ptr {
		tpe = tpe.Elem()
		val = val.Elem()
	}

	fieldValues := make([]any, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i).Interface()
		fieldValues = append(fieldValues, fieldValue)
	}

	return fieldValues
}
