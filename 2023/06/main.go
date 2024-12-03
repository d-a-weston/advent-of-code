package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	fileScanner := bufio.NewScanner(file)

	var raceTime string
	var raceDistance string

	for fileScanner.Scan() {
		line := fileScanner.Text()

		lineSplit := strings.Fields(line)

		if lineSplit[0] == "Time:" {
			for _, time := range lineSplit[1:] {
				raceTime += time
			}
		}

		if lineSplit[0] == "Distance:" {
			for _, distance := range lineSplit[1:] {
				raceDistance += distance
			}
		}
	}

	file.Close()

	time, _ := strconv.Atoi(raceTime)
	distance, _ := strconv.Atoi(raceDistance)

	total := numValidSolutions(time, distance)

	fmt.Println(total)
}

func numValidSolutions(time int, distance int) int {
	totalSolutions := 0

	for i := 0; i <= time; i++ {
		velocity := i
		remainingTime := time - i
		distanceTravelled := velocity * remainingTime

		if distanceTravelled > distance {
			totalSolutions++
		}
	}

	return totalSolutions
}
