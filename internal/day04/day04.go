package day04

import (
	"fmt"
	"strings"

	"github.com/EdwinRy/advent-2024/internal/utils"
)

type direction int

const (
	up direction = iota
	down
	left
	right
	diagonalUpRight
	diagonalUpLeft
	diagonalDownRight
	diagonalDownLeft
)

var directions = []direction{up, down, left, right, diagonalUpRight, diagonalUpLeft, diagonalDownRight, diagonalDownLeft}

func readInputCharRows() [][]string {
	// 	input := `MMMSXXMASM
	// MSAMXMSMSA
	// AMXSXMAAMM
	// MSAMASMSMX
	// XMASAMXAMM
	// XXAMMXXAMA
	// SMSMSASXSS
	// SAXAMASAAA
	// MAMMMXMMMM
	// MXMXAXMASX`
	input, _ := utils.ReadFile("inputs/day_04/input_1.txt")
	lines := strings.Split(input, "\n")
	charsLines := make([][]string, 0)
	for _, line := range lines {
		lineChars := strings.Split(line, "")
		charsLines = append(charsLines, lineChars)
	}
	return charsLines
}

func checkWordInDirection(lines *[][]string, word string, wordIdx int, dir direction, currI int, currJ int) bool {

	xmax := len((*lines)[0]) - 1
	ymax := len(*lines) - 1

	currChar := (*lines)[currI][currJ]
	if currChar != string(word[wordIdx]) {
		return false
	}

	if wordIdx == len(word)-1 {
		return true
	}

	switch dir {
	case up:
		if currI == 0 {
			return false
		}
		return checkWordInDirection(lines, word, wordIdx+1, dir, currI-1, currJ)
	case down:
		if currI == ymax {
			return false
		}
		return checkWordInDirection(lines, word, wordIdx+1, dir, currI+1, currJ)
	case left:
		if currJ == 0 {
			return false
		}
		return checkWordInDirection(lines, word, wordIdx+1, dir, currI, currJ-1)
	case right:
		if currJ == xmax {
			return false
		}
		return checkWordInDirection(lines, word, wordIdx+1, dir, currI, currJ+1)
	case diagonalUpRight:
		if currI == 0 || currJ == xmax {
			return false
		}
		return checkWordInDirection(lines, word, wordIdx+1, dir, currI-1, currJ+1)
	case diagonalUpLeft:
		if currI == 0 || currJ == 0 {
			return false
		}
		return checkWordInDirection(lines, word, wordIdx+1, dir, currI-1, currJ-1)
	case diagonalDownRight:
		if currI == ymax || currJ == xmax {
			return false
		}
		return checkWordInDirection(lines, word, wordIdx+1, dir, currI+1, currJ+1)
	case diagonalDownLeft:
		if currI == ymax || currJ == 0 {
			return false
		}
		return checkWordInDirection(lines, word, wordIdx+1, dir, currI+1, currJ-1)
	}

	return false
}

func task1() {
	lines := readInputCharRows()
	word := "XMAS"
	total := 0
	for i, _ := range lines {
		for j, _ := range lines[i] {
			for _, dir := range directions {
				if checkWordInDirection(&lines, word, 0, dir, i, j) {
					total++
				}
			}

		}
	}
	fmt.Println("Day 04 task 1: ", total)
}

func findIfQuadsOverlap(xQuad [][]string, yQuad [][]string, ignoreChar string) bool {
	for i, _ := range xQuad {
		for j, _ := range xQuad[i] {
			if xQuad[i][j] == ignoreChar {
				continue
			}
			if xQuad[i][j] != yQuad[i][j] {
				return false
			}
		}
	}
	return true
}

func rotate2dArrayClockwise(quad [][]string) [][]string {
	n := len(quad)
	newQuad := make([][]string, n)
	for i, _ := range quad {
		newQuad[i] = make([]string, n)
	}
	for i, _ := range quad {
		for j, _ := range quad[i] {
			newQuad[j][n-i-1] = quad[i][j]
		}
	}
	return newQuad
}

func quadFromLines(lines [][]string, lineI int, lineJ int, quadSize int) [][]string {
	quad := make([][]string, quadSize)
	for i := 0; i < quadSize; i++ {
		quad[i] = make([]string, quadSize)
		for j := 0; j < quadSize; j++ {
			quad[i][j] = lines[lineI+i][lineJ+j]
		}
	}
	return quad
}

func task2() {
	quad := [][]string{
		{"M", ".", "M"},
		{".", "A", "."},
		{"S", ".", "S"},
	}

	quads := make([][][]string, 4)
	quads[0] = quad
	for i := 1; i < 4; i++ {
		quads[i] = rotate2dArrayClockwise(quads[i-1])
	}

	lines := readInputCharRows()
	total := 0
	for i := 0; i < len(lines)-len(quad)+1; i++ {
		for j := 0; j < len(lines[0])-len(quad[0])+1; j++ {
			currQuad := quadFromLines(lines, i, j, len(quad))
			for _, q := range quads {
				if findIfQuadsOverlap(q, currQuad, ".") {
					total++
					break
				}
			}
		}
	}
	fmt.Println("Day 04 task 2: ", total)
}

func Day04() {
	task1()
	task2()
}
