package utils

func MergeMaps(a, b map[string]any) map[string]any {
	for k, v := range b {
		if vMap, ok := v.(map[string]any); ok {
			if aMap, ok := a[k].(map[string]any); ok {
				a[k] = MergeMaps(aMap, vMap)
				continue
			}
		}
		a[k] = v
	}
	return a
}
