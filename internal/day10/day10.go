package day10

import (
	"fmt"

	"github.com/EdwinRy/advent-2024/internal/utils"
)

type point2d = utils.Point2d

func getTrailLocations(topoMap [][]int, destinations map[point2d]bool, currPoint point2d, expectNum int) {

	if utils.CheckPointOutOfBounds(topoMap, currPoint) {
		return
	}

	currValPtr, err := utils.GetMapValue(topoMap, currPoint)
	if err != nil {
		return
	}
	currVal := *currValPtr

	if currVal != expectNum {
		return
	}

	if currVal == 9 {
		destinations[currPoint] = true
		return
	}

	getTrailLocations(topoMap, destinations, point2d{X: currPoint.X + 1, Y: currPoint.Y}, expectNum+1)
	getTrailLocations(topoMap, destinations, point2d{X: currPoint.X - 1, Y: currPoint.Y}, expectNum+1)
	getTrailLocations(topoMap, destinations, point2d{X: currPoint.X, Y: currPoint.Y + 1}, expectNum+1)
	getTrailLocations(topoMap, destinations, point2d{X: currPoint.X, Y: currPoint.Y - 1}, expectNum+1)
}

func task1(input string) (int, error) {
	topoMap := utils.ReadAsRowsOfDigits(input)

	totalScore := 0
	for i, row := range topoMap {
		for j, val := range row {
			if val == 0 {
				curr := point2d{X: i, Y: j}
				destinations := map[point2d]bool{}
				getTrailLocations(topoMap, destinations, curr, 0)
				trailScore := len(destinations)
				totalScore += trailScore
			}
		}
	}
	return totalScore, nil
}

func getTrailScore(topoMap [][]int, currPoint point2d, expectNum int) int {

	if utils.CheckPointOutOfBounds(topoMap, currPoint) {
		return 0
	}

	currValPtr, err := utils.GetMapValue(topoMap, currPoint)
	if err != nil {
		return 0
	}
	currVal := *currValPtr

	if currVal != expectNum {
		return 0
	}

	if currVal == 9 {
		return 1
	}

	score := 0
	score += getTrailScore(topoMap, point2d{X: currPoint.X + 1, Y: currPoint.Y}, expectNum+1)
	score += getTrailScore(topoMap, point2d{X: currPoint.X - 1, Y: currPoint.Y}, expectNum+1)
	score += getTrailScore(topoMap, point2d{X: currPoint.X, Y: currPoint.Y + 1}, expectNum+1)
	score += getTrailScore(topoMap, point2d{X: currPoint.X, Y: currPoint.Y - 1}, expectNum+1)
	return score
}

func task2(input string) (int, error) {
	topoMap := utils.ReadAsRowsOfDigits(input)

	totalScore := 0
	for i, row := range topoMap {
		for j, val := range row {
			if val == 0 {
				curr := point2d{X: i, Y: j}
				trailScore := getTrailScore(topoMap, curr, 0)
				totalScore += trailScore
			}
		}
	}
	return totalScore, nil
}

func Day10() {
	input, _ := utils.ReadFile("inputs/day10.txt")
	task1Result, _ := task1(input)
	fmt.Println("Day 10 task 1: ", task1Result)
	task2Result, _ := task2(input)
	fmt.Println("Day 10 task 2: ", task2Result)
}
