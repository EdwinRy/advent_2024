package day06

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/EdwinRy/advent-2024/internal/utils"
)

type point2d = utils.Point2d
type position = utils.Position
type direction = utils.Direction

type MapInfo struct {
	obstacleMap   [][]string
	guardPosition position
	obstacle      string
	free          string
	pastPositions map[position]bool
	pastLocations map[point2d]bool
}

func findGuard(obstacleMap [][]string) (position, error) {
	for y, row := range obstacleMap {
		for x, cell := range row {
			switch cell {
			case `^`:
				return position{Location: point2d{X: x, Y: y}, Direction: utils.Up}, nil
			case `v`:
				return position{Location: point2d{X: x, Y: y}, Direction: utils.Down}, nil
			case `<`:
				return position{Location: point2d{X: x, Y: y}, Direction: utils.Left}, nil
			case `>`:
				return position{Location: point2d{X: x, Y: y}, Direction: utils.Right}, nil
			}
		}
	}
	return position{}, fmt.Errorf("guard not found")

}

func checkOutOfBounds(mapInfo MapInfo) bool {
	currentPos := mapInfo.guardPosition
	switch currentPos.Direction {
	case utils.Up:
		if currentPos.Location.Y == 0 {
			return true
		}
	case utils.Down:
		if currentPos.Location.Y == len(mapInfo.obstacleMap)-1 {
			return true
		}
	case utils.Left:
		if currentPos.Location.X == 0 {
			return true
		}
	case utils.Right:
		if currentPos.Location.X == len(mapInfo.obstacleMap[0])-1 {
			return true
		}
	}
	return false
}

func checkInFrontOfObstacle(mapInfo MapInfo) bool {
	currentPos := mapInfo.guardPosition
	switch currentPos.Direction {
	case utils.Up:
		if mapInfo.obstacleMap[currentPos.Location.Y-1][currentPos.Location.X] == mapInfo.obstacle {
			return true
		}
	case utils.Down:
		if mapInfo.obstacleMap[currentPos.Location.Y+1][currentPos.Location.X] == mapInfo.obstacle {
			return true
		}
	case utils.Left:
		if mapInfo.obstacleMap[currentPos.Location.Y][currentPos.Location.X-1] == mapInfo.obstacle {
			return true
		}
	case utils.Right:
		if mapInfo.obstacleMap[currentPos.Location.Y][currentPos.Location.X+1] == mapInfo.obstacle {
			return true
		}
	}
	return false
}

func turnRight(mapInfo *MapInfo) {
	currentPos := mapInfo.guardPosition
	switch currentPos.Direction {
	case utils.Up:
		mapInfo.guardPosition.Direction = utils.Right
	case utils.Down:
		mapInfo.guardPosition.Direction = utils.Left
	case utils.Left:
		mapInfo.guardPosition.Direction = utils.Up
	case utils.Right:
		mapInfo.guardPosition.Direction = utils.Down
	}
}

func checkRepeatingPosition(mapInfo MapInfo) bool {
	if _, ok := mapInfo.pastPositions[mapInfo.guardPosition]; ok {
		return true
	}
	return false
}

func savePosition(mapInfo *MapInfo) {
	mapInfo.pastLocations[mapInfo.guardPosition.Location] = true
	mapInfo.pastPositions[mapInfo.guardPosition] = true
	mapInfo.obstacleMap[mapInfo.guardPosition.Location.Y][mapInfo.guardPosition.Location.X] = `^`
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
	mapInfo.obstacleMap[mapInfo.guardPosition.Location.Y][mapInfo.guardPosition.Location.X] = `.`

	// take step in valid direction
	switch mapInfo.guardPosition.Direction {
	case utils.Up:
		mapInfo.guardPosition.Location.Y--
	case utils.Down:
		mapInfo.guardPosition.Location.Y++
	case utils.Left:
		mapInfo.guardPosition.Location.X--
	case utils.Right:
		mapInfo.guardPosition.Location.X++
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
		pastPositions: make(map[position]bool),
		pastLocations: make(map[point2d]bool),
	}

	mapInfo.pastPositions[mapInfo.guardPosition] = true
	mapInfo.pastLocations[mapInfo.guardPosition.Location] = true

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

func createMapInfoWithNewObstacle(obstacleMap [][]string, guardPosition position, obstacleLocation point2d) MapInfo {

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
		pastPositions: make(map[position]bool),
		pastLocations: make(map[point2d]bool),
	}

	mapInfo.obstacleMap[obstacleLocation.Y][obstacleLocation.X] = mapInfo.obstacle
	mapInfo.pastPositions[mapInfo.guardPosition] = true
	mapInfo.pastLocations[mapInfo.guardPosition.Location] = true
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

			newMap := createMapInfoWithNewObstacle(obstacleMap, guardPosition, point2d{X: x, Y: y})
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
	input, _ := utils.ReadFile("inputs/day06.txt")
	task1Result, _ := task1(input)
	fmt.Println("Day 06 task 1: ", task1Result)
	task2Result, _ := task2(input)
	fmt.Println("Day 06 task 2: ", task2Result)
}
