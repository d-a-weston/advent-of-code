package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	fileScanner := bufio.NewScanner(file)

	directions := ""
	startingPositions := []string{}
	maps := map[string]map[string]string{}
	lineNumber := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if lineNumber == 0 {
			directions = line
		}

		re := regexp.MustCompile(`[A-Z]{3}`)
		matches := re.FindAll([]byte(line), -1)

		if len(matches) == 3 {
			maps[string(matches[0])] = map[string]string{
				"L": string(matches[1]),
				"R": string(matches[2]),
			}

			if strings.HasSuffix(string(matches[0]), "A") {
				startingPositions = append(startingPositions, string(matches[0]))
			}
		}

		lineNumber++
	}

	stepsPerPath := []int{}

	for _, startingPosition := range startingPositions {
		stepsPerPath = append(stepsPerPath, getNumSteps(maps, directions, startingPosition))
	}

	steps := 1

	for _, step := range stepsPerPath {
		steps = leastCommonMultiple(steps, step)
	}

	fmt.Println(steps)

	file.Close()
}

func getNumSteps(maps map[string]map[string]string, directions string, startingPosition string) int {
	steps := 0
	currentStep := 0
	currentMap := startingPosition

	for !strings.HasSuffix(currentMap, "Z") {
		steps++

		currentMap = maps[currentMap][string(directions[currentStep])]

		if currentStep == len(directions)-1 {
			currentStep = 0
		} else {
			currentStep++
		}
	}

	return steps
}

func greatestCommonDivisor(a, b int) int {
	if b == 0 {
		return a
	}

	return greatestCommonDivisor(b, a%b)
}

func leastCommonMultiple(a, b int) int {
	return (a / greatestCommonDivisor(a, b)) * b
}
