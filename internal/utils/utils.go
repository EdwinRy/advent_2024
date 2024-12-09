package utils

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(path string) (string, error) {
	reader, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer reader.Close()

	fileBuf, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	outStr := strings.Trim(string(fileBuf), "\n")
	return outStr, nil
}

func WriteStringToFile(path string, content string) error {
	os.MkdirAll(filepath.Dir(path), os.ModePerm)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func AbsDiffInt(x int, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func SliceSumInt(slice []int) int {
	sum := 0
	for _, val := range slice {
		sum += val
	}
	return sum
}

func SliceRemove(slice []int, index int) []int {
	copiedSlice := make([]int, 0, len(slice)-1)
	copiedSlice = append(copiedSlice, slice[:index]...)
	copiedSlice = append(copiedSlice, slice[index+1:]...)
	return copiedSlice
}

func SliceContainsInt(slice []int, val int) bool {
	for _, sliceVal := range slice {
		if sliceVal == val {
			return true
		}
	}
	return false
}

func ReadAsRowsOfChars(input string) [][]string {
	lines := strings.Split(input, "\n")
	charsLines := make([][]string, 0)
	for _, line := range lines {
		lineChars := strings.Split(line, "")
		charsLines = append(charsLines, lineChars)
	}
	return charsLines
}
