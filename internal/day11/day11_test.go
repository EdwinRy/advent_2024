package day11

import (
	"strings"
	"testing"
)

const input string = `125 17`

func Test_task1_correctAnswer(t *testing.T) {
	input := strings.Trim(input, "\n")
	expected := 55312

	if result, _ := task(input, 25); result != expected {
		t.Errorf("expected task1() = %v, got %v", expected, result)
	}
}
