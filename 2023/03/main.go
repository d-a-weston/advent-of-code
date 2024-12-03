package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var numbers = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)

	input := [][]string{}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		input = append(input, strings.Split(line, ""))
	}

	total := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == "*" {
				partPower := findGearParts(input, j, i)

				if partPower != -1 {
					total += partPower
				}
			}
		}
	}

	fmt.Println(total)
}

func checkSurroundingCells(input [][]string, start int, end int, lineNumber int) bool {
	ignore := []string{".", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for i := lineNumber - 1; i <= lineNumber+1; i++ {
		if i < 0 || i >= len(input) {
			continue
		}

		for j := start - 1; j <= end+1; j++ {
			if j < 0 || j >= len(input[i]) {
				continue
			}

			if slices.Contains(ignore, input[i][j]) {
				continue
			}

			return true
		}
	}

	return false
}

func findGearParts(input [][]string, index int, lineNumber int) int {
	potentialParts := map[int]bool{}

	for i := lineNumber - 1; i <= lineNumber+1; i++ {
		if i < 0 || i >= len(input) {
			continue
		}

		for j := index - 1; j <= index+1; j++ {
			if j < 0 || j >= len(input[i]) {
				continue
			}

			if slices.Contains(numbers, input[i][j]) {
				potentialParts[findNumber(input[i], j)] = true
			}
		}
	}

	if len(potentialParts) != 2 {
		return -1
	}

	value := 1

	for part := range potentialParts {
		value *= part
	}

	return value
}

func findNumber(input []string, index int) int {
	number := ""

	for i := index; i < len(input); i++ {
		if slices.Contains(numbers, input[i]) {
			number = number + input[i]
		} else {
			break
		}
	}

	for i := index - 1; i >= 0; i-- {
		if slices.Contains(numbers, input[i]) {
			number = input[i] + number
		} else {
			break
		}
	}

	numberInt, _ := strconv.Atoi(number)

	return numberInt
}
