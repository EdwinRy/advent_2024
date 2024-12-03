package main

import (
	"fmt"
	"os"
	"strconv"
)

var excercises = []func(){
	Day_01,
	Day_02,
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
