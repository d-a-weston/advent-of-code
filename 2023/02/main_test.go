package main

import (
	"testing"
)

func TestColorMax(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{"Color max #1", "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", map[string]int{"blue": 6, "red": 4, "green": 2}},
		{"Color max #2", "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", map[string]int{"blue": 4, "red": 1, "green": 3}},
		{"Color max #3", "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", map[string]int{"blue": 6, "red": 20, "green": 13}},
		{"Color max #4", "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", map[string]int{"blue": 15, "red": 14, "green": 3}},
		{"Color max #5", "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", map[string]int{"blue": 2, "red": 6, "green": 3}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := colorMax(test.input)

			for color, value := range test.expected {
				if actual[color] != value {
					t.Errorf("Expected %d, got %d", value, actual[color])
				}
			}
		})
	}
}

func TestGameIsPossible(t *testing.T) {
	bagContents := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	var tests = []struct {
		name     string
		input    map[string]int
		expected bool
	}{
		{"Game is possible #1", map[string]int{"blue": 6, "red": 4, "green": 2}, true},
		{"Game is possible #2", map[string]int{"blue": 4, "red": 1, "green": 3}, true},
		{"Game is possible #3", map[string]int{"blue": 6, "red": 20, "green": 13}, false},
		{"Game is possible #4", map[string]int{"blue": 15, "red": 14, "green": 3}, false},
		{"Game is possible #5", map[string]int{"blue": 2, "red": 6, "green": 3}, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := gameIsPossible(bagContents, test.input)

			if actual != test.expected {
				t.Errorf("Expected %t, got %t", test.expected, actual)
			}
		})
	}
}
