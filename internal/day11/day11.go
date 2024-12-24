package day11

import (
	"fmt"
	"strconv"

	"github.com/EdwinRy/advent-2024/internal/utils"
)

type stoneStep struct {
	stone int
	steps int
}

func stoneBlink(cache map[stoneStep]int, stone int, blinksLeft int) int {
	if blinksLeft == 0 {
		return 1
	}

	if _, ok := cache[stoneStep{stone, blinksLeft}]; ok {
		return cache[stoneStep{stone, blinksLeft}]
	}

	if stone == 0 {
		pathRes := stoneBlink(cache, 1, blinksLeft-1)
		cache[stoneStep{stone, blinksLeft}] = pathRes
		return pathRes
	}

	stoneNumStr := fmt.Sprintf("%d", stone)
	stoneNumLen := len(stoneNumStr)

	if stoneNumLen%2 == 0 {
		firstHalf := stoneNumStr[:stoneNumLen/2]
		secondHalf := stoneNumStr[stoneNumLen/2:]
		firstHalfInt, _ := strconv.Atoi(firstHalf)
		secondHalfInt, _ := strconv.Atoi(secondHalf)

		res1 := stoneBlink(cache, firstHalfInt, blinksLeft-1)
		res2 := stoneBlink(cache, secondHalfInt, blinksLeft-1)

		cache[stoneStep{stone, blinksLeft}] = res1 + res2
		return res1 + res2
	}

	res := stoneBlink(cache, stone*2024, blinksLeft-1)
	cache[stoneStep{stone, blinksLeft}] = res
	return res
}

func task(input string, count int) (int, error) {
	stones := utils.ReadAsListOfNums(input, " ")

	total := 0
	for _, stone := range stones {
		total += stoneBlink(map[stoneStep]int{}, stone, count)
	}

	return total, nil
}

func Day11() {
	input, _ := utils.ReadFile("inputs/day11.txt")
	task1Result, _ := task(input, 25)
	fmt.Println("Day 11 task 1: ", task1Result)
	task2Result, _ := task(input, 75)
	fmt.Println("Day 11 task 2: ", task2Result)
}
