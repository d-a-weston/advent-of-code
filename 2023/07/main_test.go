package main

import (
	"testing"
)

func TestGetHandType(t *testing.T) {
	tests := []struct {
		name     string
		hand     string
		expected int
	}{
		{"Gets value of high card", "32TJ5", 1},
		{"Gets value of one pair", "32T3K", 2},
		{"Gets value of two pairs", "32T32", 3},
		{"Gets value of three of a kind", "T55J5", 4},
		{"Gets value of full house", "555JJ", 5},
		{"Gets value of four of a kind", "5555J", 6},
		{"Gets value of five of a kind", "55555", 7},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := getHandType(test.hand)

			if actual != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, actual)
			}
		})
	}
}

func TestDoesHandWin(t *testing.T) {
	tests := []struct {
		name     string
		hand1    string
		hand2    string
		expected bool
	}{
		{"Returns true if hand one wins", "AKQJT", "23456", true},
		{"Returns false if hand two wins", "23456", "AKQJT", false},
		{"Return true if hand one wins with a pair", "AA234", "A2345", true},
		{"Return false if hand two wins with full house vs two pair", "AA233", "AAA44", false},
		{"Returns true if hand one wins with higher leading card", "A2364", "A2345", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := doesHandWin(test.hand1, test.hand2)

			if actual != test.expected {
				t.Errorf("Expected %t, got %t", test.expected, actual)
			}
		})
	}
}
