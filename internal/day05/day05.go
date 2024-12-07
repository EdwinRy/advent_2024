package day05

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/EdwinRy/advent-2024/internal/utils"
)

func processOrderingRules(input string) map[int]map[int]bool {
	orderingPartRows := strings.Split(input, "\n")
	ordering := make(map[int]map[int]bool)
	for _, row := range orderingPartRows {
		parts := strings.Split(row, "|")
		pre, _ := strconv.Atoi(parts[0])
		post, _ := strconv.Atoi(parts[1])
		if _, ok := ordering[pre]; !ok {
			ordering[pre] = make(map[int]bool)
		}
		ordering[pre][post] = true
	}
	return ordering
}

func processIntRows(input string) [][]int {
	rows := strings.Split(input, "\n")
	intRows := make([][]int, len(rows))
	for i, row := range rows {
		rowNums := strings.Split(row, ",")
		intRow := make([]int, len(rowNums))
		for j, num := range rowNums {
			intNum, _ := strconv.Atoi(num)
			intRow[j] = intNum
		}
		intRows[i] = intRow
	}

	return intRows
}

func numbersCorrectlyOrdered(ordering map[int]map[int]bool, intRow []int) bool {

	seen := make(map[int]bool)

	for i := 0; i < len(intRow); i++ {
		curr := intRow[i]
		numAfter, hasRule := ordering[curr]
		if !hasRule {
			seen[curr] = true
			continue
		}

		for k, _ := range numAfter {
			if seen[k] {
				return false
			}
		}

		seen[curr] = true
	}
	return true
}

func task1(input string) (int, error) {
	inputParts := strings.Split(input, "\n\n")
	ordering := processOrderingRules(inputParts[0])
	intRows := processIntRows(inputParts[1])

	correctRows := make([][]int, 0)
	for _, intRow := range intRows {
		if numbersCorrectlyOrdered(ordering, intRow) {
			correctRows = append(correctRows, intRow)
		}
	}

	totalMids := 0
	for _, row := range correctRows {
		middle := len(row) / 2
		middleNum := row[middle]
		totalMids += middleNum
	}

	return totalMids, nil
}

func orderRow(ordering map[int]map[int]bool, row []int) []int {
	rowCopy := make([]int, len(row))
	copy(rowCopy, row)

	numScores := make(map[int]int)
	for _, num := range rowCopy {
		for postNum, _ := range ordering[num] {
			numScores[postNum]++
		}
	}

	sort.Slice(rowCopy, func(i, j int) bool {
		return numScores[rowCopy[i]] < numScores[rowCopy[j]]
	})

	return rowCopy
}

func task2(input string) (int, error) {
	inputParts := strings.Split(input, "\n\n")
	ordering := processOrderingRules(inputParts[0])
	intRows := processIntRows(inputParts[1])

	incorrectRows := make([][]int, 0)
	for _, intRow := range intRows {
		if !numbersCorrectlyOrdered(ordering, intRow) {
			incorrectRows = append(incorrectRows, intRow)
		}
	}
	orderedRows := make([][]int, 0)
	for _, row := range incorrectRows {
		orderedRows = append(orderedRows, orderRow(ordering, row))
	}
	totalMids := 0
	for _, row := range orderedRows {
		middle := len(row) / 2
		middleNum := row[middle]
		totalMids += middleNum
	}

	return totalMids, nil
}

func Day05() {
	input, _ := utils.ReadFile("inputs/day_05/input_1.txt")
	task1Result, _ := task1(input)
	fmt.Println("Day 05 task 1: ", task1Result)
	task2Result, _ := task2(input)
	fmt.Println("Day 05 task 2: ", task2Result)
}
