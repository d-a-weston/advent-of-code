package main

import (
	"testing"
)

func TestCheckSurroundingCells(t *testing.T) {
	tests := []struct {
		name       string
		input      [][]string
		start      int
		finish     int
		lineNumber int
		expected   bool
	}{
		{"Find single digit number with symbol", [][]string{{".", ".", "."}, {".", "1", "."}, {"*", ".", "."}}, 1, 1, 1, true},
		{"Find two number digit with symbol", [][]string{{".", ".", ".", "*"}, {".", "1", "2", "."}, {".", ".", ".", "."}}, 1, 2, 1, true},
		{"Number on top row", [][]string{{".", "9", "0", "*"}, {".", "*", ".", "."}, {".", ".", ".", "."}}, 1, 2, 0, true},
		{"Invalid case", [][]string{{".", ".", ".", "*"}, {".", "3", ".", "."}, {".", ".", ".", "."}}, 1, 1, 1, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := checkSurroundingCells(test.input, test.start, test.finish, test.lineNumber)

			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
