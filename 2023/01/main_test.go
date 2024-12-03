package main

import (
	"testing"
)

func TestFindNumInLine(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected int
	}{
		{"Numbers at first and last", "1abc2", 12},
		{"Numbers not at first and last", "pqr3stu8vwx", 38},
		{"More than 2 numbers", "a1b2c3d4e5f", 15},
		{"Only one number", "treb7uchet", 77},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := findNumInLine(test.input)

			if actual != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, actual)
			}
		})
	}
}

func TestReplaceTokens(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{"Replace one with 1", "1one1", "1on1e1"},
		{"Replace multiple numbers", "two1nine", "tw2o1nin9e"},
		{"Replace each valid number in the string", "eightwothree", "eigh8tw2othre3e"},
		{"Replace while preserving numbers", "xtwone3four", "xtw2on1e3fou4r"},
		{"Additional AoC test case #1", "4nineeightseven2", "4nin9eeigh8tseve7n2"},
		{"Additional AoC test case #2", "zoneight234", "zon1eigh8t234"},
		{"Additional AoC test case #3", "7pqrstsixteen", "7pqrstsi6xteen"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := replaceTokens(test.input)

			if actual != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, actual)
			}
		})
	}
}
