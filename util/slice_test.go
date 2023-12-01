package tests

import (
	"egghead/app/util"
	"reflect"
	"testing"
)

func TestContainsString(t *testing.T) {
	tests := []struct {
		slice    []string
		target   string
		expected bool
	}{
		{[]string{"apple", "banana", "orange"}, "banana", true},
		{[]string{"apple", "banana", "orange"}, "grape", false},
		{[]string{}, "test", false},
	}

	for _, test := range tests {
		result := util.ContainsString(test.slice, test.target)
		if result != test.expected {
			t.Errorf("ContainsString(%v, %s) returned %v, expected %v", test.slice, test.target, result, test.expected)
		}
	}
}

func TestUniqueString(t *testing.T) {
	tests := []struct {
		slice    []string
		expected []string
	}{
		{[]string{"apple", "banana", "banana", "orange"}, []string{"apple", "banana", "orange"}},
		{[]string{"apple", "orange", "grape"}, []string{"apple", "orange", "grape"}},
		{[]string{"apple", "apple", "apple"}, []string{"apple"}},
		{[]string{}, []string{}},
	}

	for _, test := range tests {
		result := util.UniqueString(test.slice)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("UniqueString(%v) returned %v, expected %v", test.slice, result, test.expected)
		}
	}
}
