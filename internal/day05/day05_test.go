package day05

import (
	"strings"
	"testing"
)

const input string = `
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func Test_task1_correctAnswer(t *testing.T) {
	input := strings.Trim(input, "\n")
	expected := 143

	if result, _ := task1(input); result != expected {
		t.Errorf("expected task1() = %v, got %v", expected, result)
	}
}

func Test_task2_correctAnswer(t *testing.T) {
	input := strings.Trim(input, "\n")
	expected := 123

	if result, _ := task2(input); result != expected {
		t.Errorf("expected task2() = %v, got %v", expected, result)
	}
}
