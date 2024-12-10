package day07

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/EdwinRy/advent-2024/internal/utils"
)

type Row struct {
	total int
	parts []int
}

func lineToRow(input string) Row {
	totalAndParts := strings.Split(input, ":")
	total, _ := strconv.Atoi(totalAndParts[0])
	parts := strings.Split(strings.Trim(totalAndParts[1], " "), " ")
	intParts := make([]int, len(parts))
	for i, part := range parts {
		intPart, _ := strconv.Atoi(part)
		intParts[i] = intPart
	}
	return Row{total: total, parts: intParts}
}

func inputToRows(input string) []Row {
	rows := strings.Split(input, "\n")
	rowObjects := make([]Row, len(rows))
	for i, row := range rows {
		rowObjects[i] = lineToRow(row)
	}
	return rowObjects
}

func bigIntToSetLenBinaryStr(num big.Int, setLen int) string {
	binaryStr := num.Text(2)
	padded := fmt.Sprintf("%0*s", setLen, binaryStr)
	return padded
}

func performAddMulOperations(parts []int, operators string) int {
	total := parts[0]
	for i, part := range parts[1:] {
		switch operators[i] {
		case '0':
			total += part
		case '1':
			total *= part
		}
	}
	return total
}

func checkCanBeMadeWithParts1(row Row) bool {
	parts := row.parts
	operators := len(parts) - 1
	i := big.NewInt(0)
	maxIter := big.NewInt(2)
	maxIter.Exp(maxIter, big.NewInt(int64(operators)), nil)
	for i.Cmp(maxIter) < 0 {
		binaryStr := bigIntToSetLenBinaryStr(*i, operators)

		if performAddMulOperations(parts, binaryStr) == row.total {
			return true
		}

		i.Add(i, big.NewInt(1))
	}
	return false
}

func task1(input string) (int, error) {
	rows := inputToRows(input)

	validTotals := make([]int, 0)
	for _, row := range rows {
		if checkCanBeMadeWithParts1(row) {
			validTotals = append(validTotals, row.total)
			println(row.total)
		}
	}

	total := 0
	for _, validTotal := range validTotals {
		total += validTotal
	}

	return total, nil
}
func bigIntToSetLenTernaryStr(num big.Int, setLen int) string {
	binaryStr := num.Text(3)
	padded := fmt.Sprintf("%0*s", setLen, binaryStr)
	return padded
}

func performAddMulConcatOperations(parts []int, operators string) int {
	total := parts[0]
	for i, part := range parts[1:] {
		switch operators[i] {
		case '0':
			total += part
		case '1':
			total *= part
		case '2':
			totalStr := strconv.Itoa(total)
			partStr := strconv.Itoa(part)
			concatStr := totalStr + partStr
			total, _ = strconv.Atoi(concatStr)
		}
	}
	return total
}

func checkCanBeMadeWithParts2(row Row) bool {
	parts := row.parts
	operators := len(parts) - 1
	i := big.NewInt(0)
	maxIter := big.NewInt(3)
	maxIter.Exp(maxIter, big.NewInt(int64(operators)), nil)
	for i.Cmp(maxIter) < 0 {
		binaryStr := bigIntToSetLenTernaryStr(*i, operators)

		if performAddMulConcatOperations(parts, binaryStr) == row.total {
			return true
		}

		i.Add(i, big.NewInt(1))
	}
	return false
}

func task2(input string) (int, error) {
	rows := inputToRows(input)

	validTotals := make([]int, 0)
	for _, row := range rows {
		if checkCanBeMadeWithParts2(row) {
			validTotals = append(validTotals, row.total)
			// println(row.total)
		}
	}

	total := 0
	for _, validTotal := range validTotals {
		total += validTotal
	}

	return total, nil
}

func Day07() {
	input, _ := utils.ReadFile("inputs/day07/input.txt")
	task1Result, _ := task1(input)
	fmt.Println("Day 07 task 1: ", task1Result)
	task2Result, _ := task2(input)
	fmt.Println("Day 07 task 2: ", task2Result)
}
