package dgcoll

import (
	dgutils "github.com/darwinOrg/go-common/utils"
	"math/rand"
	"strconv"
	"strings"
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

func Sort[T any](slice []T, less Less[T]) {
	if len(slice) < 2 {
		return
	}

	dgutils.Cmp[T](func(t1, t2 *T) bool {
		return less(*t1, *t2)
	}).Sort(slice)
}

func SortAsc[T any, V int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64](slice []T, mapFunc Function[T, V]) {
	Sort(slice, func(t1, t2 T) bool {
		return mapFunc(t1) < mapFunc(t2)
	})
}

func SortDesc[T any, V int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64](slice []T, mapFunc Function[T, V]) {
	Sort(slice, func(t1, t2 T) bool {
		return mapFunc(t1) > mapFunc(t2)
	})
}

func SortByIds[T any, V comparable](slice []T, ids []V, mapFunc Function[T, V]) []T {
	if len(slice) == 0 {
		return []T{}
	}
	if len(ids) == 0 {
		return slice
	}

	sliceMap := Trans2Map(slice, mapFunc)
	var newSlice []T
	for _, id := range ids {
		value, ok := sliceMap[id]
		if ok {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}

func Contains[T comparable](slice []T, t T) bool {
	if len(slice) == 0 {
		return false
	}

	return AnyMatch(slice, func(s T) bool {
		return s == t
	})
}

func ContainsAny[T comparable](slice1 []T, slice2 []T) bool {
	if len(slice1) == 0 || len(slice2) == 0 {
		return false
	}

	return AnyMatch(slice1, func(t1 T) bool {
		return AnyMatch(slice2, func(t2 T) bool {
			return t1 == t2
		})
	})
}

func MergeToList[T any](slices ...[]T) []T {
	if len(slices) == 0 {
		return []T{}
	}

	var newSlice []T
	for _, slice := range slices {
		if len(slice) > 0 {
			newSlice = append(newSlice, slice...)
		}
	}

	return newSlice
}

func MergeToSet[T comparable](slices ...[]T) []T {
	return DeDupToSet(MergeToList(slices...))
}

func JoinInts[T int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64](slice []T, sep string) string {
	if len(slice) == 0 {
		return ""
	}

	strs := MapToList(slice, func(t T) string { return strconv.FormatInt(int64(t), 10) })
	return strings.Join(strs, sep)
}

func SplitToInts[T int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64](str string, sep string) []T {
	if len(str) == 0 {
		return []T{}
	}

	var jsonStr string
	if sep == "," {
		jsonStr = "[" + str + "]"
	} else {
		jsonStr = "[" + strings.ReplaceAll(str, sep, ",") + "]"
	}

	return dgutils.MustConvertJsonStringToList[T](jsonStr)
}

func Shuffle[T any](slice []T) {
	if len(slice) == 0 {
		return
	}

	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
}

func Remove[T comparable](slice []T, elements []T) []T {
	if len(slice) == 0 {
		return []T{}
	}

	if len(elements) == 0 {
		return slice
	}

	return FilterList(slice, func(t T) bool {
		return !Contains(elements, t)
	})
}
