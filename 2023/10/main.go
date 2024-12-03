package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	field := [][]string{}
	startX := 0
	startY := 0

	lineNum := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		containsStart := slices.Index(line, "S")

		if containsStart != -1 {
			startX = containsStart
			startY = lineNum
		}

		field = append(field, line)

		lineNum++
	}

	direction := "down"
	pipes := map[string]bool{}
	curX := startX
	curY := startY
	backToStart := false

	for !backToStart {
		if direction == "down" {
			curY++
		} else if direction == "up" {
			curY--
		} else if direction == "left" {
			curX--
		} else if direction == "right" {
			curX++
		}

		direction = setDirection(field[curY][curX], direction)

		pipes[fmt.Sprintf("%d,%d", curY, curX)] = true

		if field[curY][curX] == "S" {
			backToStart = true
		}
	}

	contained := 0
	walls := []string{"|", "L", "J", "S"}

	for y := range field {
		inside := false
		for x := range field[y] {
			if !pipes[fmt.Sprintf("%d,%d", y, x)] {
				if inside {
					contained++
				}
			} else if slices.Contains(walls, field[y][x]) {
				inside = !inside
			}
		}
	}

	fmt.Println(contained)

	file.Close()
}

func setDirection(pipe string, currDirection string) string {
	nextDirection := ""

	if pipe == "|" {
		nextDirection = currDirection
	} else if pipe == "-" {
		nextDirection = currDirection
	} else if pipe == "F" {
		if currDirection == "up" {
			nextDirection = "right"
		} else if currDirection == "left" {
			nextDirection = "down"
		}
	} else if pipe == "7" {
		if currDirection == "up" {
			nextDirection = "left"
		} else if currDirection == "right" {
			nextDirection = "down"
		}
	} else if pipe == "L" {
		if currDirection == "down" {
			nextDirection = "right"
		} else if currDirection == "left" {
			nextDirection = "up"
		}
	} else if pipe == "J" {
		if currDirection == "down" {
			nextDirection = "left"
		} else if currDirection == "right" {
			nextDirection = "up"
		}
	}

	return nextDirection
}
