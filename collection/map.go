package dgcoll

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
