package day06

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/EdwinRy/advent-2024/internal/utils"
)

type Vec2d struct {
	x, y int
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Position struct {
	location  Vec2d
	direction Direction
}

type MapInfo struct {
	obstacleMap   [][]string
	guardPosition Position
	obstacle      string
	free          string
	pastPositions map[Position]bool
	pastLocations map[Vec2d]bool
}

func findGuard(obstacleMap [][]string) (Position, error) {
	for y, row := range obstacleMap {
		for x, cell := range row {
			switch cell {
			case `^`:
				return Position{location: Vec2d{x, y}, direction: Up}, nil
			case `v`:
				return Position{location: Vec2d{x, y}, direction: Down}, nil
			case `<`:
				return Position{location: Vec2d{x, y}, direction: Left}, nil
			case `>`:
				return Position{location: Vec2d{x, y}, direction: Right}, nil
			}
		}
	}
	return Position{}, fmt.Errorf("guard not found")

}

func checkOutOfBounds(mapInfo MapInfo) bool {
	currentPos := mapInfo.guardPosition
	switch currentPos.direction {
	case Up:
		if currentPos.location.y == 0 {
			return true
		}
	case Down:
		if currentPos.location.y == len(mapInfo.obstacleMap)-1 {
			return true
		}
	case Left:
		if currentPos.location.x == 0 {
			return true
		}
	case Right:
		if currentPos.location.x == len(mapInfo.obstacleMap[0])-1 {
			return true
		}
	}
	return false
}

func checkInFrontOfObstacle(mapInfo MapInfo) bool {
	currentPos := mapInfo.guardPosition
	switch currentPos.direction {
	case Up:
		if mapInfo.obstacleMap[currentPos.location.y-1][currentPos.location.x] == mapInfo.obstacle {
			return true
		}
	case Down:
		if mapInfo.obstacleMap[currentPos.location.y+1][currentPos.location.x] == mapInfo.obstacle {
			return true
		}
	case Left:
		if mapInfo.obstacleMap[currentPos.location.y][currentPos.location.x-1] == mapInfo.obstacle {
			return true
		}
	case Right:
		if mapInfo.obstacleMap[currentPos.location.y][currentPos.location.x+1] == mapInfo.obstacle {
			return true
		}
	}
	return false
}

func turnRight(mapInfo *MapInfo) {
	currentPos := mapInfo.guardPosition
	switch currentPos.direction {
	case Up:
		mapInfo.guardPosition.direction = Right
	case Down:
		mapInfo.guardPosition.direction = Left
	case Left:
		mapInfo.guardPosition.direction = Up
	case Right:
		mapInfo.guardPosition.direction = Down
	}
}

func checkRepeatingPosition(mapInfo MapInfo) bool {
	if _, ok := mapInfo.pastPositions[mapInfo.guardPosition]; ok {
		return true
	}
	return false
}

func savePosition(mapInfo *MapInfo) {
	mapInfo.pastLocations[mapInfo.guardPosition.location] = true
	mapInfo.pastPositions[mapInfo.guardPosition] = true
	mapInfo.obstacleMap[mapInfo.guardPosition.location.y][mapInfo.guardPosition.location.x] = `^`
}

func step(mapInfo *MapInfo) (isFinished bool, circularError bool) {

	// check if reached end
	if checkOutOfBounds(*mapInfo) {
		return true, false
	}

	// turn around and check next available direction
	for i := 0; i < 3; i++ {
		if checkInFrontOfObstacle(*mapInfo) {
			turnRight(mapInfo)
			if checkRepeatingPosition(*mapInfo) {
				return false, true
			}
			mapInfo.pastPositions[mapInfo.guardPosition] = true
		} else {
			break
		}
	}

	// erase previous position
	mapInfo.obstacleMap[mapInfo.guardPosition.location.y][mapInfo.guardPosition.location.x] = `.`

	// take step in valid direction
	switch mapInfo.guardPosition.direction {
	case Up:
		mapInfo.guardPosition.location.y--
	case Down:
		mapInfo.guardPosition.location.y++
	case Left:
		mapInfo.guardPosition.location.x--
	case Right:
		mapInfo.guardPosition.location.x++
	}

	savePosition(mapInfo)
	return false, false
}

func task1(input string) (int, error) {
	obstacleMap := utils.ReadAsRowsOfChars(input)
	guardPosition, _ := findGuard(obstacleMap)

	mapInfo := MapInfo{
		obstacle:      `#`,
		free:          `.`,
		obstacleMap:   obstacleMap,
		guardPosition: guardPosition,
		pastPositions: make(map[Position]bool),
		pastLocations: make(map[Vec2d]bool),
	}

	mapInfo.pastPositions[mapInfo.guardPosition] = true
	mapInfo.pastLocations[mapInfo.guardPosition.location] = true

	circularError := false
	finished := false
	maxRuns := 100000
	for !circularError && !finished && maxRuns > 0 {
		maxRuns--
		finished, circularError = step(&mapInfo)
	}

	if maxRuns == 0 {
		fmt.Println("maxRuns reached")
		return 0, fmt.Errorf("maxRuns reached")
	}

	individualPositions := len(mapInfo.pastLocations)
	return individualPositions, nil
}

func createMapInfoWithNewObstacle(obstacleMap [][]string, guardPosition Position, obstacleLocation Vec2d) MapInfo {

	newObstacleMap := make([][]string, len(obstacleMap))
	for i, row := range obstacleMap {
		newObstacleMap[i] = make([]string, len(row))
		copy(newObstacleMap[i], row)
	}

	mapInfo := MapInfo{
		obstacle:      `#`,
		free:          `.`,
		obstacleMap:   newObstacleMap,
		guardPosition: guardPosition,
		pastPositions: make(map[Position]bool),
		pastLocations: make(map[Vec2d]bool),
	}

	mapInfo.obstacleMap[obstacleLocation.y][obstacleLocation.x] = mapInfo.obstacle
	mapInfo.pastPositions[mapInfo.guardPosition] = true
	mapInfo.pastLocations[mapInfo.guardPosition.location] = true
	return mapInfo
}

func checkMapResultsInCircularError(mapInfo *MapInfo) bool {
	circularError := false
	finished := false
	maxRuns := 1000000
	for !circularError && !finished && maxRuns > 0 {
		maxRuns--
		finished, circularError = step(mapInfo)
	}

	if maxRuns == 0 {
		fmt.Println("maxRuns reached")
		return false
	}

	if circularError {
		return true
	} else {
		return false
	}
}

func task2(input string) (int, error) {
	obstacleMap := utils.ReadAsRowsOfChars(input)
	guardPosition, _ := findGuard(obstacleMap)

	// create all possible maps with one obstacle added
	allMaps := make([]MapInfo, 0)
	for y, row := range obstacleMap {
		for x, cell := range row {
			// skip guard position and obstacles
			if cell == `^` || cell == `v` || cell == `<` || cell == `>` || cell == `#` {
				continue
			}

			newMap := createMapInfoWithNewObstacle(obstacleMap, guardPosition, Vec2d{x, y})
			allMaps = append(allMaps, newMap)
		}
	}

	var circularErrors atomic.Uint64
	var wg sync.WaitGroup
	for _, mapInfo := range allMaps {
		wg.Add(1)
		mapRef := &mapInfo
		go func(mapInfo *MapInfo) {
			defer wg.Done()
			if checkMapResultsInCircularError(mapInfo) {
				circularErrors.Add(1)
			}
		}(mapRef)
	}
	wg.Wait()

	return int(circularErrors.Load()), nil
}

func Day06() {
	input, _ := utils.ReadFile("inputs/day06/input.txt")
	task1Result, _ := task1(input)
	fmt.Println("Day 06 task 1: ", task1Result)
	task2Result, _ := task2(input)
	fmt.Println("Day 06 task 2: ", task2Result)
}
