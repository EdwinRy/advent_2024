package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/EdwinRy/advent-2024/internal/day01"
	"github.com/EdwinRy/advent-2024/internal/day02"
	"github.com/EdwinRy/advent-2024/internal/day03"
	"github.com/EdwinRy/advent-2024/internal/day04"
	"github.com/EdwinRy/advent-2024/internal/day05"
	"github.com/EdwinRy/advent-2024/internal/day06"
	"github.com/EdwinRy/advent-2024/internal/day07"
)

var excercises = []func(){
	day01.Day01,
	day02.Day02,
	day03.Day03,
	day04.Day04,
	day05.Day05,
	day06.Day06,
	day07.Day07,
}

func main() {
	// Run a specific day
	if len(os.Args) > 1 {
		day, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("format: <program> <day number>")
			return
		}
		excercises[day-1]()
		return
	}

	// Run all days
	for _, excercise := range excercises {
		excercise()
	}
}
