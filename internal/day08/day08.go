package day08

import (
	"fmt"

	"github.com/EdwinRy/advent-2024/internal/utils"
)

type Point struct {
	x, y int
}

func getFreqPositions(grid [][]string, emptyChars []string) map[string][]Point {
	freqMap := make(map[string][]Point)
	for y, row := range grid {
		for x, cell := range row {
			if !utils.SliceContains(emptyChars, cell) {
				freqMap[cell] = append(freqMap[cell], Point{x: x, y: y})
			}
		}
	}
	return freqMap
}

func markFreqPositions(grid [][]string, points []Point, emptySpace string) {
	for i, x := range points {
		for j, y := range points {
			if i == j {
				continue
			}

			dx := x.x - y.x
			dy := x.y - y.y

			newX := x.x + dx
			newY := x.y + dy

			notOutsideX := newX >= 0 && newX < len(grid)
			notOutsideY := newY >= 0 && newY < len(grid[0])

			if notOutsideX && notOutsideY {
				grid[newY][newX] = "#"
			}
		}
	}
}

// todo: this is probably wrong
func task1(input string) (int, error) {
	grid := utils.ReadAsRowsOfChars(input)
	freqMap := getFreqPositions(grid, []string{"#", "."})

	for _, v := range freqMap {
		markFreqPositions(grid, v, ".")
	}

	total := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == "#" {
				total++
			}
		}
	}

	return total, nil
}

func markRepeatFreqPositions(grid [][]string, points []Point, emptySpace string) {
	for i, x := range points {
		for j, y := range points {
			if i == j {
				continue
			}

			grid[x.y][x.x] = "#"
			grid[y.y][y.x] = "#"

			dx := x.x - y.x
			dy := x.y - y.y

			newX1 := x.x + dx
			newY1 := x.y + dy
			notOutsideX1 := newX1 >= 0 && newX1 < len(grid)
			notOutsideY1 := newY1 >= 0 && newY1 < len(grid[0])

			for notOutsideX1 && notOutsideY1 {

				grid[newY1][newX1] = "#"
				newX1 += dx
				newY1 += dy

				notOutsideX1 = newX1 >= 0 && newX1 < len(grid)
				notOutsideY1 = newY1 >= 0 && newY1 < len(grid[0])
			}

			newX2 := x.x + dx
			newY2 := x.y + dy
			notOutsideX2 := newX2 >= 0 && newX2 < len(grid)
			notOutsideY2 := newY2 >= 0 && newY2 < len(grid[0])

			for notOutsideX2 && notOutsideY2 {
				grid[newY2][newX2] = "#"
				newX2 -= dx
				newY2 -= dy

				notOutsideX2 = newX2 >= 0 && newX2 < len(grid)
				notOutsideY2 = newY2 >= 0 && newY2 < len(grid[0])
			}
		}
	}
}

func task2(input string) (int, error) {
	grid := utils.ReadAsRowsOfChars(input)
	freqMap := getFreqPositions(grid, []string{"#", "."})

	for _, v := range freqMap {
		markRepeatFreqPositions(grid, v, ".")
	}

	total := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == "#" {
				total++
			}
		}
	}

	return total, nil
}

func Day08() {
	input, _ := utils.ReadFile("inputs/day08/input.txt")
	task1Result, _ := task1(input)
	fmt.Println("Day 08 task 1: ", task1Result)
	task2Result, _ := task2(input)
	fmt.Println("Day 08 task 2: ", task2Result)
}
