package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	board := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		board = append(board, strings.Split(line, ""))
	}

	board = rotateLeft(board)

	cycles := 1000000000
	tiltedBoard := board

	cachedBoards := map[string]int{}
	i := 0
	cycleStart := 0

	for ; i < cycles; i++ {
		tiltedBoard = spinBoard(tiltedBoard)

		boardKey := createBoardKey(tiltedBoard)
		val, ok := cachedBoards[boardKey]

		if ok {
			cycleStart = val
			break
		} else {
			cachedBoards[boardKey] = i
		}
	}

	cyclesRemaining := (cycles - i - 1) % (i - cycleStart)

	for j := 0; j < cyclesRemaining; j++ {
		tiltedBoard = spinBoard(tiltedBoard)
	}

	tiltedBoard = rotateRight(tiltedBoard)

	total := 0

	for i, row := range tiltedBoard {
		numStones := 0

		for _, col := range row {
			if col == "O" {
				numStones++
			}
		}

		total += numStones * (len(tiltedBoard) - i)
	}

	fmt.Println(total)

	file.Close()
}

func spinBoard(tiltedBoard [][]string) [][]string {
	tiltedBoard = tiltBoard(tiltedBoard)
	tiltedBoard = rotateRight(tiltedBoard)

	tiltedBoard = tiltBoard(tiltedBoard)
	tiltedBoard = rotateRight(tiltedBoard)

	tiltedBoard = tiltBoard(tiltedBoard)
	tiltedBoard = rotateRight(tiltedBoard)

	tiltedBoard = tiltBoard(tiltedBoard)
	tiltedBoard = rotateRight(tiltedBoard)

	return tiltedBoard
}

func tiltBoard(board [][]string) [][]string {
	tiltedBoard := [][]string{}

	for i := 0; i < len(board); i++ {
		tiltedBoard = append(tiltedBoard, []string{})
	}

	for i := 0; i < len(board); i++ {
		rowSlice := board[i]
		tiltedSlice := []string{}
		tiltedSlice = tiltSlice(rowSlice)
		tiltedBoard[i] = tiltedSlice
	}

	return tiltedBoard
}

func tiltSlice(slice []string) []string {
	i := 0
	j := 0

	for j < len(slice) {
		if slice[j] == "O" {
			slice[i] = "O"

			if i != j {
				slice[j] = "."
			}

			i++
		} else if slice[j] == "#" {
			i = j + 1
		}

		j++
	}

	return slice
}

func rotateRight(matrix [][]string) [][]string {
	n := len(matrix)
	if n == 0 || len(matrix[0]) == 0 {
		return matrix
	}

	result := make([][]string, len(matrix[0]))
	for i := range result {
		result[i] = make([]string, n)
		for j := range result[i] {
			result[i][j] = matrix[n-j-1][i]
		}
	}
	return result
}

func rotateLeft(matrix [][]string) [][]string {
	n := len(matrix)
	if n == 0 || len(matrix[0]) == 0 {
		return matrix
	}

	result := make([][]string, len(matrix[0]))
	for i := range result {
		result[i] = make([]string, n)
		for j := range result[i] {
			result[i][j] = matrix[j][n-i-1]
		}
	}
	return result
}

func createBoardKey(board [][]string) string {
	key := ""

	for _, row := range board {
		key += fmt.Sprintf("%v", row)
	}

	return key
}
