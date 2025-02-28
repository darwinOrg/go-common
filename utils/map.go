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

func RemoveMapEmptyValues(m map[string]any) map[string]any {
	cleaned := make(map[string]any)

	for k, v := range m {
		switch value := v.(type) {
		case string:
			if value != "" {
				cleaned[k] = v
			}
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			if value != 0 {
				cleaned[k] = v
			}
		case float32, float64:
			if value != 0.0 {
				cleaned[k] = v
			}
		case bool:
			cleaned[k] = v
		case nil:
			continue
		default:
			if value != nil {
				cleaned[k] = v
			}
		}
	}

	return cleaned
}
