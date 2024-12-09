package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/EdwinRy/advent-2024/internal/utils"
)

func task1(input string) (int, error) {
	r := regexp.MustCompile(`mul\(\d+\,\d+\)`)
	matches := r.FindAllString(input, -1)

	total := 0
	digits := regexp.MustCompile(`\d+\,\d+`)
	for _, match := range matches {
		nums := digits.FindAllString(match, -1)
		num1, _ := strconv.Atoi(strings.Split(nums[0], ",")[0])
		num2, _ := strconv.Atoi(strings.Split(nums[0], ",")[1])
		total += num1 * num2
	}

	return total, nil
}

func task2(input string) (int, error) {
	reg_mul := `mul\(\d+\,\d+\)`
	reg_dont := `don\'t\(\)`
	reg_do := `do\(\)`
	r := regexp.MustCompile(reg_mul + "|" + reg_dont + "|" + reg_do)
	matches := r.FindAllString(input, -1)

	enabled := true
	total := 0
	for _, match := range matches {
		if match == "don't()" {
			enabled = false
		} else if match == "do()" {
			enabled = true
		} else {
			if enabled {
				nums := regexp.MustCompile(`\d+\,\d+`).FindAllString(match, -1)
				num1, _ := strconv.Atoi(strings.Split(nums[0], ",")[0])
				num2, _ := strconv.Atoi(strings.Split(nums[0], ",")[1])
				total += num1 * num2
			}
		}
	}
	return total, nil
}

func Day03() {
	input, _ := utils.ReadFile("inputs/day03/input.txt")
	task1Result, _ := task1(input)
	fmt.Println("Day 03 task 1: ", task1Result)
	task2Result, _ := task2(input)
	fmt.Println("Day 03 task 2: ", task2Result)
}
