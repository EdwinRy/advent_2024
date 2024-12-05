package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/EdwinRy/advent-2024/internal/day01"
	"github.com/EdwinRy/advent-2024/internal/day02"
	"github.com/EdwinRy/advent-2024/internal/day03"
	"github.com/EdwinRy/advent-2024/internal/day04"
)

var excercises = []func(){
	day01.Day01,
	day02.Day02,
	day03.Day03,
	day04.Day04,
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
