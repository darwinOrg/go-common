package dgcoll

func Trans2NewMap[T comparable, K any, V any](mp map[T]K, mapFunc Function[K, V]) map[T]V {
	if len(mp) == 0 {
		return map[T]V{}
	}

	newMap := make(map[T]V, len(mp))

	for k, v := range mp {
		newMap[k] = mapFunc(v)
	}

	return newMap
}
