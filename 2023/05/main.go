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

	maps := map[string][][]string{}
	mapNames := []string{}

	var mapName string

	var seeds []string

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if strings.HasPrefix(line, "seeds: ") {
			seeds = strings.Fields(strings.TrimPrefix(line, "seeds: "))

			continue
		}

		if strings.Contains(line, "-") {
			mapName = strings.Fields(line)[0]
			mapNames = append(mapNames, mapName)

			continue
		}

		if line != "" {
			splitLine := strings.Fields(line)
			maps[mapName] = append(maps[mapName], splitLine)
		}
	}

	closestSeed := 0
	validLocation := false

	for !validLocation {
		currentMap := "humidity-to-location"

		currentValue := closestSeed
		prevValue := closestSeed

		for currentMap != "" {
			prevValue = currentValue

			for _, slice := range maps[currentMap] {
				destinationCategory, _ := strconv.Atoi(slice[0])
				rangeValue, _ := strconv.Atoi(slice[2])

				lowerBound := destinationCategory
				upperBound := destinationCategory + rangeValue

				if currentValue >= lowerBound && currentValue < upperBound {
					sourceCategory, _ := strconv.Atoi(slice[1])

					destinationDiff := currentValue - destinationCategory

					currentValue = sourceCategory + destinationDiff
					currentMap = findNextMap(mapNames, currentMap)

					break
				}
			}

			if prevValue == currentValue {
				currentMap = findNextMap(mapNames, currentMap)
			}
		}

		for i := 0; i < len(seeds); i += 2 {
			seedStart, _ := strconv.Atoi(seeds[i])
			seedRange, _ := strconv.Atoi(seeds[i+1])

			if currentValue >= seedStart && currentValue <= seedStart+seedRange {
				validLocation = true
				break
			}
		}

		if !validLocation {
			closestSeed++
		}
	}

	fmt.Println(closestSeed)
}

func findNextMap(maps []string, currentMap string) string {
	nextMap := ""

	splitCurrentMap := strings.Split(currentMap, "-")

	mapPrefix := splitCurrentMap[0]

	for _, v := range maps {

		mapSplit := strings.Split(v, "-")

		if mapSplit[2] == mapPrefix {
			nextMap = v
		}
	}

	return nextMap
}
