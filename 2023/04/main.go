package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)

	total := 0
	gameNumber := 0
	copies := map[int]int{}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		gameNumber++

		removePrefix := strings.Split(line, ":")

		numbers := strings.Split(removePrefix[1], "|")
		winningNumbers := strings.Fields(numbers[0])
		playersNumbers := strings.Fields(numbers[1])

		commonElements := findCommonElements(winningNumbers, playersNumbers)

		numCommonElements := len(commonElements)

		for i := gameNumber + 1; i <= gameNumber+numCommonElements; i++ {
			copies[i] += 1 + copies[gameNumber]
		}

		total += 1 + copies[gameNumber]
	}

	fmt.Println(total)
}

func findCommonElements(firstSlice []string, secondSlice []string) []string {
	sharedElements := []string{}

	for _, v1 := range firstSlice {
		for _, v2 := range secondSlice {
			if v1 == v2 {
				sharedElements = append(sharedElements, v1)
			}
		}
	}

	return sharedElements
}
