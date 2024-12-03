package main

import (
	"advent_2024_go/utils"
	"fmt"
	"strconv"
	"strings"
)

func read2dNumArray(input string) ([][]int, error) {
	fileLines := strings.Split(input, "\n")
	itemCount := len(fileLines)
	list := make([][]int, itemCount)

	for i, line := range fileLines {
		nums := strings.Split(line, " ")
		for _, num := range nums {
			num, _ := strconv.Atoi(num)
			list[i] = append(list[i], num)
		}
	}

	fileLines = nil
	return list, nil
}

func isRowChangingSafely(row []int) (bool, int) {
	if len(row) < 2 {
		return true, 0
	}

	increasing := row[0] < row[1]

	for i := 0; i < len(row)-1; i++ {

		n1 := row[i]
		n2 := row[i+1]

		if increasing && n1 > n2 {
			return false, i
		}

		if !increasing && n1 < n2 {
			return false, i
		}

		if n1 == n2 {
			return false, i
		}

		if utils.AbsDiffInt(n1, n2) > 3 {
			return false, i
		}
	}
	return true, 0
}

func day_02_task1() {
	input, _ := utils.ReadFile("inputs/day_02/input_1.txt")
	rows, _ := read2dNumArray(input)

	safeRows := 0
	for _, row := range rows {
		safe, _ := isRowChangingSafely(row)
		if safe {
			safeRows++
		}
	}
	fmt.Println("Day 02, Task 1: ", safeRows)
}

func day_02_task2() {
	input, _ := utils.ReadFile("inputs/day_02/input_2.txt")
	rows, _ := read2dNumArray(input)

	safeRows := 0
	for _, row := range rows {
		safe, _ := isRowChangingSafely(row)

		if safe {
			safeRows++
			continue
		}

		// todo: refactor this
		succ := false
		for i := 0; i < len(row); i++ {
			row2 := utils.SliceRemove(row, i)
			safe, _ := isRowChangingSafely(row2)
			if safe {
				succ = true
				break
			}
		}
		if succ {
			safeRows++
			continue
		}
	}
	fmt.Println("Day 02, Task 2: ", safeRows)

}

func Day_02() {
	day_02_task1()
	day_02_task2()
}
