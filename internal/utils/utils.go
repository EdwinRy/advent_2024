package utils

import (
	"strconv"
	"strings"
)

func ReadAsRowsOfChars(input string) [][]string {
	lines := strings.Split(input, "\n")
	charsLines := make([][]string, 0)
	for _, line := range lines {
		lineChars := strings.Split(line, "")
		charsLines = append(charsLines, lineChars)
	}
	return charsLines
}

func ReadAsRowsOfDigits(input string) [][]int {
	lines := strings.Split(input, "\n")
	charsLines := make([][]int, 0)
	for _, line := range lines {
		lineChars := strings.Split(line, "")
		lineInts := make([]int, 0)
		for _, char := range lineChars {
			intVal, _ := strconv.Atoi(char)
			lineInts = append(lineInts, intVal)
		}
		charsLines = append(charsLines, lineInts)
	}
	return charsLines
}

func ReadAsListOfNums(input string, sep string) []int {
	strNums := strings.Split(input, sep)
	nums := make([]int, 0)
	for _, strNum := range strNums {
		num, _ := strconv.Atoi(strNum)
		nums = append(nums, num)
	}
	return nums
}
