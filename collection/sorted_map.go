package dgcoll

import (
	"encoding/json"
)

type SortedMap[K comparable, V any] struct {
	keys   []K
	values map[K]V
}

func NewSortedMap[K comparable, V any]() *SortedMap[K, V] {
	return &SortedMap[K, V]{
		keys:   make([]K, 0),
		values: make(map[K]V),
	}
}

func (sm *SortedMap[K, V]) Set(key K, value V) {
	if _, exists := sm.values[key]; !exists {
		sm.keys = append(sm.keys, key)
	}
	sm.values[key] = value
}

func (sm *SortedMap[K, V]) Get(key K) (V, bool) {
	value, exists := sm.values[key]
	return value, exists
}

func (sm *SortedMap[K, V]) Delete(key K) {
	if _, exists := sm.values[key]; exists {
		delete(sm.values, key)
		for i, k := range sm.keys {
			if k == key {
				sm.keys = append(sm.keys[:i], sm.keys[i+1:]...)
				break
			}
		}
	}
}

func (sm *SortedMap[K, V]) Keys() []K {
	return sm.keys
}

func (sm *SortedMap[K, V]) Values() []V {
	values := make([]V, len(sm.keys))
	for i, key := range sm.keys {
		values[i] = sm.values[key]
	}
	return values
}

func (sm *SortedMap[K, V]) Len() int {
	return len(sm.keys)
}

func (sm *SortedMap[K, V]) MarshalJSON() ([]byte, error) {
	type kv struct {
		Key   K
		Value V
	}

	kvs := make([]kv, len(sm.keys))
	for i, key := range sm.keys {
		kvs[i] = kv{Key: key, Value: sm.values[key]}
	}

	return json.Marshal(kvs)
}
