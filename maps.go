package maps

import (
	"sort"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string
}

// SortKeys sorts the keys of a map.
func SortKeys[K Ordered, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i += 1
	}
	sort.Sort(sortable[K](keys))
	return keys
}

// KeyValue is a Key Value pair.
type KeyValue[K Ordered, V any] struct {
	Key   K
	Value V
}
type KeyVals[K Ordered, V any] []KeyValue[K, V]

// Sort sorts a map by key, returning a slice of KV structs.
func SortByKeys[K Ordered, V any](m map[K]V) KeyVals[K, V] {
	sorted := make(KeyVals[K, V], len(m))
	keys := SortKeys(m)
	for i, k := range keys {
		sorted[i] = KeyValue[K, V]{k, m[k]}
	}
	return sorted
}

type sortable[T Ordered] []T

func (x sortable[T]) Len() int               { return len(x) }
func (x sortable[T]) Less(i int, j int) bool { return x[i] < x[j] }
func (x sortable[T]) Swap(i int, j int)      { x[i], x[j] = x[j], x[i] }

var _ sort.Interface = (sortable[uint8])(nil)
