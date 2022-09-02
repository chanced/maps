package maps

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// SortKeys sorts the keys of a map.
func SortKeys[K constraints.Ordered, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i += 1
	}
	sort.Sort(sortable[K](keys))
	return keys
}

// KeyVal is a Key Value pair.
type KeyVal[K constraints.Ordered, V any] struct {
	Key K
	Val V
}
type KeyVals[K constraints.Ordered, V any] []KeyVal[K, V]

// Sort sorts a map by key, returning a slice of KV structs.
func SortByKeys[K constraints.Ordered, V any](m map[K]V) KeyVals[K, V] {
	sorted := make(KeyVals[K, V], len(m))
	keys := SortKeys(m)
	for i, k := range keys {
		sorted[i] = KeyVal[K, V]{k, m[k]}
	}
	return sorted
}

type sortable[T constraints.Ordered] []T

func (x sortable[T]) Len() int               { return len(x) }
func (x sortable[T]) Less(i int, j int) bool { return x[i] < x[j] }
func (x sortable[T]) Swap(i int, j int)      { x[i], x[j] = x[j], x[i] }

var _ sort.Interface = (sortable[uint8])(nil)
