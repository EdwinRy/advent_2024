package utils

import "fmt"

type Point2d struct {
	X int
	Y int
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
	DiagonalUpRight
	DiagonalUpLeft
	DiagonalDownRight
	DiagonalDownLeft
)

type Position struct {
	Location  Point2d
	Direction Direction
}

var DirectionsOrtho = []Direction{Up, Down, Left, Right}
var Directions2d = []Direction{Up, Down, Left, Right, DiagonalUpRight, DiagonalUpLeft, DiagonalDownRight, DiagonalDownLeft}

func GetMapValue[T any](m [][]T, p Point2d) (*T, error) {

	if CheckPointOutOfBounds(m, p) {
		return nil, fmt.Errorf("out of bounds")
	}

	return &m[p.X][p.Y], nil
}

func CheckPointOutOfBounds[T any](m [][]T, p Point2d) bool {
	return p.X < 0 || p.X >= len(m) || p.Y < 0 || p.Y >= len(m[0])
}
