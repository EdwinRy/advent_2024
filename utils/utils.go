package utils

import (
	"io"
	"os"
	"path/filepath"
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

	return string(fileBuf), nil
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

func AbsInt(x int, y int) int {
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
