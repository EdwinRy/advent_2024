package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/EdwinRy/advent-2024/internal/utils"
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

		if utils.AbsDiffNum(n1, n2) > 3 {
			return false, i
		}
	}
	return true, 0
}

func task1(input string) (int, error) {
	rows, _ := read2dNumArray(input)

	safeRows := 0
	for _, row := range rows {
		safe, _ := isRowChangingSafely(row)
		if safe {
			safeRows++
		}
	}
	return safeRows, nil
}

func task2(input string) (int, error) {
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
	return safeRows, nil
}

func Day02() {
	input, _ := utils.ReadFile("inputs/day02.txt")
	task1Result, _ := task1(input)
	fmt.Println("Day 02 task 1: ", task1Result)
	task2Result, _ := task2(input)
	fmt.Println("Day 02 task 2: ", task2Result)
}
