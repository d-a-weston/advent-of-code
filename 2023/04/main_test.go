package main

import (
	"testing"
)

func TestNumSharedElements(t *testing.T) {
	tests := []struct {
		name        string
		firstSlice  []string
		secondSlice []string
		expected    []string
	}{
		{"Shares one element", []string{"1", "2", "3", "4"}, []string{"4", "5", "6"}, []string{"4"}},
		{"Shares two elements", []string{"7", "8", "9"}, []string{"9", "10", "11", "8"}, []string{"8", "9"}},
		{"Shares no elements", []string{"12", "13", "14"}, []string{"15", "16", "17"}, []string{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := findCommonElements(test.firstSlice, test.secondSlice)

			if len(actual) != len(test.expected) {
				t.Errorf("Expected %s, got %s", test.expected, actual)
			}

			for i, v := range actual {
				if v != test.expected[i] {
					t.Errorf("Expected %s, got %s", test.expected, actual)
				}
			}
		})
	}
}
