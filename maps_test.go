package maps

import (
	"reflect"
	"testing"
)

func TestSortKeys(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	expected := []string{"a", "b", "c"}
	for i, k := range SortKeys(m) {
		if k != expected[i] {
			t.Errorf("expected %s, got %s", expected[i], k)
		}
	}
}

func TestMerge(t *testing.T) {
	r := Merge([]map[string]int{
		{"a": 1, "b": 2},
		{"c": 3, "d": 4},
		{"e": 5, "f": 6},
	}...)

	expected := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
	}
	if reflect.DeepEqual(r, expected) == false {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestSortByKeys(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
	}

	expected := []KeyValue[string, int]{
		{"a", 1},
		{"b", 2},
		{"c", 3},
		{"d", 4},
		{"e", 5},
		{"f", 6},
		{"g", 7},
		{"h", 8},
		{"i", 9},
		{"j", 10},
		{"k", 11},
		{"l", 12},
		{"m", 13},
		{"n", 14},
		{"o", 15},
		{"p", 16},
	}
	for i, k := range SortByKeys(m) {
		if k != expected[i] {
			t.Errorf("expected %v, got %v", expected[i], k)
		}
	}
}
