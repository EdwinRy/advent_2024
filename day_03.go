package main

import (
	"advent_2024_go/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day_03_task1() {
	input, _ := utils.ReadFile("inputs/day_03/input_1.txt")
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

	fmt.Println("Day 03 task 1: ", total)
}

func day_03_task2() {
	input, _ := utils.ReadFile("inputs/day_03/input_1.txt")
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
	fmt.Println("Day 03 task 2: ", total)
}

func Day_03() {
	day_03_task1()
	day_03_task2()
}
