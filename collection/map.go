package dgcoll

func TransMap2List[T comparable, K any, V any](mp map[T]K, mapFunc BiFunction[T, K, V]) []V {
	if len(mp) == 0 {
		return []V{}
	}

	list := make([]V, 0, len(mp))

	for t, k := range mp {
		v := mapFunc(t, k)
		list = append(list, v)
	}

	return list
}

func TransMap2Set[T comparable, K any, V comparable](mp map[T]K, mapFunc BiFunction[T, K, V]) []V {
	return DeDupToSet(TransMap2List(mp, mapFunc))
}

func Trans2NewMap[T comparable, K any, V any](mp map[T]K, mapFunc Function[K, V]) map[T]V {
	if len(mp) == 0 {
		return map[T]V{}
	}

	newMap := make(map[T]V, len(mp))

	for t, k := range mp {
		newMap[t] = mapFunc(k)
	}

	return newMap
}

func Trans2NewListMap[T comparable, K any, V any](mp map[T][]K, mapFunc Function[K, V]) map[T][]V {
	if len(mp) == 0 {
		return map[T][]V{}
	}

	newMap := make(map[T][]V, len(mp))

	for t, keys := range mp {
		newMap[t] = MapToList(keys, mapFunc)
	}

	return newMap
}

func Trans2NewSetMap[T comparable, K any, V comparable](mp map[T][]K, mapFunc Function[K, V]) map[T][]V {
	if len(mp) == 0 {
		return map[T][]V{}
	}

	newMap := make(map[T][]V, len(mp))

	for t, keys := range mp {
		newMap[t] = MapToSet(keys, mapFunc)
	}

	return newMap
}

func MapToSliceByKeysAndFilterEmpty[T comparable, V any](mp map[T]*V, keys []T) []*V {
	if len(mp) == 0 {
		return []*V{}
	}

	vs := MapToList(keys, func(t T) *V {
		return mp[t]
	})
	if len(vs) == 0 {
		return []*V{}
	}

	return FilterList(vs, func(v *V) bool {
		return v != nil
	})
}

func ExtractMapValues[K comparable, V any](mp map[K]V) []V {
	if len(mp) == 0 {
		return []V{}
	}

	var values []V
	for _, v := range mp {
		values = append(values, v)
	}

	return values
}

func ExtractMapValueSet[K comparable, V comparable](mp map[K]V) []V {
	return DeDupToSet(ExtractMapValues(mp))
}
