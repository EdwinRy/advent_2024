package day08

import (
	"strings"
	"testing"
)

const input string = `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`

func Test_task1_correctAnswer(t *testing.T) {
	input := strings.Trim(input, "\n")
	expected := 14

	if result, _ := task1(input); result != expected {
		t.Errorf("expected task1() = %v, got %v", expected, result)
	}
}

func Test_task2_correctAnswer(t *testing.T) {
	input := strings.Trim(input, "\n")
	expected := 34

	if result, _ := task2(input); result != expected {
		t.Errorf("expected task2() = %v, got %v", expected, result)
	}
}
