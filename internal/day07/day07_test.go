package day07

import (
	"strings"
	"testing"
)

const input string = `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func Test_task1_correctAnswer(t *testing.T) {
	input := strings.Trim(input, "\n")
	expected := 3749

	if result, _ := task1(input); result != expected {
		t.Errorf("expected task1() = %v, got %v", expected, result)
	}
}

func Test_task2_correctAnswer(t *testing.T) {
	input := strings.Trim(input, "\n")
	expected := 11387

	if result, _ := task2(input); result != expected {
		t.Errorf("expected task2() = %v, got %v", expected, result)
	}
}
