package main

import (
	"testing"
)

func TestSequenceDiff(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"First example", []int{0, 3, 6, 9, 12, 15}, []int{3, 3, 3, 3, 3}},
		{"Second example", []int{3, 3, 3, 3, 3}, []int{0, 0, 0, 0}},
		{"Third example", []int{10, 13, 16, 21, 30, 45}, []int{3, 3, 5, 9, 15}},
		{"Fourth example", []int{3, 3, 5, 9, 15}, []int{0, 2, 4, 6}},
		{"Fifth example", []int{0, 2, 4, 6}, []int{2, 2, 2}},
		{"Sixth example", []int{2, 2, 2}, []int{0, 0}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := sequenceDiff(test.input)

			for i, v := range actual {
				if v != test.expected[i] {
					t.Errorf("Expected %v, got %v", test.expected, actual)
				}
			}
		})
	}
}

func TestFindNextValue(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"First example", []int{0, 3, 6, 9, 12, 15}, 18},
		{"Second example", []int{3, 3, 3, 3, 3}, 3},
		{"Third example", []int{10, 13, 16, 21, 30, 45}, 68},
		{"Fourth example", []int{3, 3, 5, 9, 15}, 23},
		{"Fifth example", []int{0, 2, 4, 6}, 8},
		{"Sixth example", []int{2, 2, 2}, 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := findNextValue(test.input)

			if actual != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestFindPreviousValue(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"First example", []int{0, 3, 6, 9, 12, 15}, -3},
		{"Second example", []int{3, 3, 3, 3, 3}, 3},
		{"Third example", []int{10, 13, 16, 21, 30, 45}, 5},
		{"Fourth example", []int{3, 3, 5, 9, 15}, 5},
		{"Fifth example", []int{0, 2, 4, 6}, -2},
		{"Sixth example", []int{2, 2, 2}, 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := findPreviousValue(test.input)

			if actual != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, actual)
			}
		})
	}
}
