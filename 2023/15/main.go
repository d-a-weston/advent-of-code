package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	lensBoxes := map[int][]map[string]string{}

	for scanner.Scan() {
		line := scanner.Text()

		instructions := strings.Split(line, ",")

		for _, instruction := range instructions {
			lensRegex, _ := regexp.Compile("([a-z]+)")
			lensLabel := lensRegex.FindString(instruction)
			boxNumber := hash(lensLabel, 17, 256)
			lensBoxes[boxNumber] = followInstruction(instruction, lensBoxes[boxNumber])
		}
	}

	total := 0

	for lensBoxIdx, lensBox := range lensBoxes {
		for lensIdx, lens := range lensBox {
			lensStrength, _ := strconv.Atoi(lens["strength"])
			total += (lensBoxIdx + 1) * (lensIdx + 1) * lensStrength
		}
	}

	fmt.Println(total)

	file.Close()
}

func hash(instruction string, convertValue int, bounded int) int {
	chars := []rune(instruction)

	hash := 0

	for _, char := range chars {
		hash += int(char)
		hash *= convertValue
		hash %= bounded
	}

	return hash
}

func followInstruction(instruction string, lensBox []map[string]string) []map[string]string {
	if strings.Contains(instruction, "-") {
		lensName := strings.Split(instruction, "-")[0]

		lensBox = slices.DeleteFunc(lensBox, func(lens map[string]string) bool {
			return lens["name"] == lensName
		})
	} else if strings.Contains(instruction, "=") {
		instructionParts := strings.Split(instruction, "=")
		lensName := instructionParts[0]
		lensStrength := instructionParts[1]

		lensIdx := slices.IndexFunc(lensBox, func(lens map[string]string) bool {
			return lens["name"] == lensName
		})

		if lensIdx == -1 {
			lensBox = append(lensBox, map[string]string{
				"name":     lensName,
				"strength": lensStrength,
			})
		} else {
			lensBox[lensIdx]["strength"] = lensStrength
		}
	}

	return lensBox
}
