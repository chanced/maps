package maps

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type sortable[T constraints.Ordered] []T

// Len implements sort.Interface
func (x sortable[T]) Len() int               { return len(x) }
func (x sortable[T]) Less(i int, j int) bool { return x[i] < x[j] }
func (x sortable[T]) Swap(i int, j int)      { x[i], x[j] = x[j], x[i] }

var _ sort.Interface = sortable[uint8]{}

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

// KV is a Key Value pair.
type KV[K constraints.Ordered, V any] struct {
	K K
	V V
}

// Sort sorts a map by key, returning a slice of KV structs.
func Sort[K constraints.Ordered, V any](m map[K]V) []KV[K, V] {
	sorted := make([]KV[K, V], len(m))
	keys := SortKeys(m)
	for i, k := range keys {
		sorted[i] = KV[K, V]{k, m[k]}
	}
	return sorted
}
