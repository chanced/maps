package maps_test

import (
	"testing"

	"github.com/chanced/maps"
)

func TestSortKeys(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	expected := []string{"a", "b", "c"}
	for i, k := range maps.SortKeys(m) {
		if k != expected[i] {
			t.Errorf("expected %s, got %s", expected[i], k)
		}
	}
}

func TestSortByKeys(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	expected := []maps.KeyVal[string, int]{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}
	for i, k := range maps.SortByKeys(m) {
		if k != expected[i] {
			t.Errorf("expected %v, got %v", expected[i], k)
		}
	}
}
