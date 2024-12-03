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

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		config := parts[0]

		groups := []int{}

		for _, group := range strings.Split(parts[1], ",") {
			groupAsInt, _ := strconv.Atoi(group)
			groups = append(groups, groupAsInt)
		}

		unfoldedConfig := []string{}
		unfoldedGroups := []int{}

		numFolds := 5

		for i := 0; i < numFolds; i++ {
			unfoldedConfig = append(unfoldedConfig, config)
			unfoldedGroups = append(unfoldedGroups, groups...)
		}

		config = strings.Join(unfoldedConfig, "?")
		groups = unfoldedGroups

		var cache [][]int
		for i := 0; i < len(config); i++ {
			cache = append(cache, make([]int, len(groups)+1))
			for j := 0; j < len(groups)+1; j++ {
				cache[i][j] = -1
			}
		}

		total += numPossibleConfigurations(0, 0, config, groups, cache)
	}

	fmt.Println(total)

	file.Close()
}

func numPossibleConfigurations(i int, j int, config string, groups []int, cache [][]int) int {
	if i >= len(config) {
		if j < len(groups) {
			return 0
		}

		return 1
	}

	if cache[i][j] != -1 {
		return cache[i][j]
	}

	result := 0

	if config[i] == '.' {
		result = numPossibleConfigurations(i+1, j, config, groups, cache)
	} else {
		if config[i] == '?' {
			result += numPossibleConfigurations(i+1, j, config, groups, cache)
		}

		if j < len(groups) {
			count := 0

			for k := i; k < len(config); k++ {
				if count > groups[j] || config[k] == '.' || count == groups[j] && config[k] == '?' {
					break
				}

				count++
			}

			if count == groups[j] {
				if i+count < len(config) && config[i+count] != '#' {
					result += numPossibleConfigurations(i+count+1, j+1, config, groups, cache)
				} else {
					result += numPossibleConfigurations(i+count, j+1, config, groups, cache)
				}
			}
		}
	}

	cache[i][j] = result

	return result
}
