package main

import (
	"bufio"
	"fmt"
	"graph"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	cityGraph := graph.New()

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}

	file.Close()
}
