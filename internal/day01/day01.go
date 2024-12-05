package day01

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/EdwinRy/advent-2024/internal/utils"
)

func read2ColsNumbers(input string) ([]int, []int, error) {
	fileLines := strings.Split(input, "\n")
	itemCount := len(fileLines)
	list1 := make([]int, itemCount)
	list2 := make([]int, itemCount)

	for i, line := range fileLines {
		nums := strings.Split(line, "   ")
		if len(nums) != 2 {
			continue
		}
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println(err)
			return nil, nil, err
		}
		num2, err2 := strconv.Atoi(nums[1])
		if err2 != nil {
			fmt.Println(err2)
			return nil, nil, err2
		}

		list1[i] = num1
		list2[i] = num2
	}

	fileLines = nil
	input = ""
	return list1, list2, nil
}

func task1() {
	input, err := utils.ReadFile("inputs/day_01/input_1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	list1, list2, err := read2ColsNumbers(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	slices.Sort(list1)
	slices.Sort(list2)

	listLen := len(list1)

	distances := make([]int, listLen)
	for i := 0; i < listLen; i++ {
		distances[i] = utils.AbsDiffInt(list1[i], list2[i])
	}

	totalDistance := utils.SliceSumInt(distances)
	fmt.Println("Day 1 Task 1: ", totalDistance)
	utils.WriteStringToFile("solutions/day_01/task_1.txt", strconv.Itoa(totalDistance))
}

func task2() {
	input, err := utils.ReadFile("inputs/day_01/input_2.txt")
	if err != nil {
		fmt.Println(err)

		return
	}

	list1, list2, err := read2ColsNumbers(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	listLen := len(list1)

	slices.Sort(list1)
	slices.Sort(list2)

	list2NumCounts := make(map[int]int)
	for _, num := range list2 {
		list2NumCounts[num]++
	}

	distances := make([]int, len(list2))
	for i := 0; i < listLen; i++ {
		n1 := list1[i]
		similarity := n1 * list2NumCounts[n1]
		distances[i] = similarity
	}

	totalDistance := utils.SliceSumInt(distances)
	fmt.Println("Day 1 Task 2: ", totalDistance)
	utils.WriteStringToFile("solutions/day_01/task_2.txt", strconv.Itoa(totalDistance))
}

func Day01() {
	task1()
	task2()
}
