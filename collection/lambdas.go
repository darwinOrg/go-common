package dgcoll

import (
	"github.com/darwinOrg/go-common/utils"
)

type Predicate[T any] func(t T) bool
type Function[T any, V any] func(t T) V
type Less[T any] func(t1 T, t2 T) bool

func Identity[T any](t T) T { return t }

func FilterList[T any](slice []T, predicate Predicate[T]) []T {
	if len(slice) == 0 {
		return []T{}
	}

	list := make([]T, 0, len(slice))
	for _, t := range slice {
		if predicate(t) {
			list = append(list, t)
		}
	}

	return list
}

func MapToList[T any, V any](slice []T, mapFunc Function[T, V]) []V {
	if len(slice) == 0 {
		return []V{}
	}

	list := make([]V, 0, len(slice))
	for _, t := range slice {
		v := mapFunc(t)
		list = append(list, v)
	}

	return list
}

func FlatMapToList[T any, V any](slice []T, mapFunc Function[T, []V]) []V {
	if len(slice) == 0 {
		return []V{}
	}

	list := make([]V, 0, len(slice))
	for _, t := range slice {
		vs := mapFunc(t)
		for _, v := range vs {
			list = append(list, v)
		}
	}

	return list
}

func FilterAndMapToList[T any, V any](slice []T, predicate Predicate[T], mapFunc Function[T, V]) []V {
	return MapToList(FilterList(slice, predicate), mapFunc)
}

func MapToSet[T any, V comparable](slice []T, mapFunc Function[T, V]) []V {
	list := MapToList(slice, mapFunc)
	if len(list) == 0 {
		return list
	}

	mp := map[V]bool{}
	set := make([]V, 0)
	for _, v := range list {
		if !mp[v] {
			set = append(set, v)
			mp[v] = true
		}
	}

	return set
}

func FlatMapToSet[T any, V comparable](slice []T, mapFunc Function[T, []V]) []V {
	list := FlatMapToList(slice, mapFunc)
	if len(list) == 0 {
		return list
	}

	mp := map[V]bool{}
	set := make([]V, 0)
	for _, v := range list {
		if !mp[v] {
			set = append(set, v)
			mp[v] = true
		}
	}

	return set
}

func DeDupToSet[T comparable](slice []T) []T {
	return MapToSet(slice, Identity[T])
}

func FilterAndMapToSet[T any, V comparable](slice []T, predicate Predicate[T], mapFunc Function[T, V]) []V {
	return MapToSet(FilterList(slice, predicate), mapFunc)
}

func Trans2Map[T any, K comparable](slice []T, keyFunc Function[T, K]) map[K]T {
	return Extract2Map(slice, keyFunc, Identity[T])
}

func Extract2Map[T any, K comparable, V any](slice []T, keyFunc Function[T, K], valueFunc Function[T, V]) map[K]V {
	if len(slice) == 0 {
		return map[K]V{}
	}

	mp := map[K]V{}
	for _, t := range slice {
		k := keyFunc(t)
		v := valueFunc(t)
		mp[k] = v
	}

	return mp
}

func GroupBy[T any, K comparable](slice []T, keyFunc Function[T, K]) map[K][]T {
	if len(slice) == 0 {
		return map[K][]T{}
	}

	mp := map[K][]T{}
	for _, t := range slice {
		k := keyFunc(t)
		if mp[k] == nil {
			mp[k] = []T{t}
		} else {
			mp[k] = append(mp[k], t)
		}
	}

	return mp
}

func AnyMatch[T any](slice []T, predicate Predicate[T]) bool {
	if len(slice) == 0 {
		return false
	}

	for _, t := range slice {
		if predicate(t) {
			return true
		}
	}

	return false
}

func AllMatch[T any](slice []T, predicate Predicate[T]) bool {
	if len(slice) == 0 {
		return false
	}

	for _, t := range slice {
		if !predicate(t) {
			return false
		}
	}

	return true
}

func NoneMatch[T any](slice []T, predicate Predicate[T]) bool {
	if len(slice) == 0 {
		return true
	}

	for _, t := range slice {
		if predicate(t) {
			return false
		}
	}

	return true
}

func FindFirst[T any](slice []T, predicate Predicate[T], defaultValue T) T {
	if len(slice) == 0 {
		return defaultValue
	}

	for _, t := range slice {
		if predicate(t) {
			return t
		}
	}

	return defaultValue
}

func Sort[T any](slice []T, less Less[*T]) {
	if len(slice) < 2 {
		return
	}

	utils.Cmp[T](func(t1, t2 *T) bool {
		return less(t1, t2)
	}).Sort(slice)
}
