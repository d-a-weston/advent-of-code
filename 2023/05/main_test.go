package main

import (
	"testing"
)

func TestFindNextMap(t *testing.T) {
	maps := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
	}

	tests := []struct {
		name       string
		currentMap string
		expected   string
	}{
		{"Finds next map", "seed-to-soil", "soil-to-fertilizer"},
		{"Finds next map", "soil-to-fertilizer", "fertilizer-to-water"},
		{"Finds no map", "foo-to-bar", ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := findNextMap(maps, test.currentMap)

			if actual != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, actual)
			}
		})
	}
}
