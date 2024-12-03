package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	starfield := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		starfield = append(starfield, strings.Split(line, ""))
	}

	file.Close()

	emptyRows := []int{}
	emptyCols := []int{}

	// Expand Starfield
	for y := 0; y < len(starfield); y++ {
		if !slices.Contains(starfield[y], "#") {
			emptyRows = append(emptyRows, y)
		}
	}

	for x := 0; x < len(starfield[0]); x++ {
		hasGalaxy := false

		for y := 0; y < len(starfield); y++ {
			if starfield[y][x] == "#" {
				hasGalaxy = true

				break
			}
		}

		if !hasGalaxy {
			emptyCols = append(emptyCols, x)
		}

		hasGalaxy = false
	}

	planets := []struct{ x, y int }{}
	spaceDistortion := 1000000

	for y := 0; y < len(starfield); y++ {
		fmt.Println(starfield[y])

		for x := 0; x < len(starfield[y]); x++ {
			if starfield[y][x] == "#" {
				distortionX := 0
				distortionY := 0

				for _, v := range emptyRows {
					if y > v {
						distortionY += spaceDistortion - 1
					}
				}

				for _, v := range emptyCols {
					if x > v {
						distortionX += spaceDistortion - 1
					}
				}

				planets = append(planets, struct{ x, y int }{x + distortionX, y + distortionY})
			}
		}
	}

	totalSteps := 0

	for i := 0; i < len(planets); i++ {
		for j := i + 1; j < len(planets); j++ {
			totalSteps += int(math.Abs(float64(planets[i].x)-float64(planets[j].x)) + math.Abs(float64(planets[i].y)-float64(planets[j].y)))
		}
	}

	fmt.Println(totalSteps)
}
