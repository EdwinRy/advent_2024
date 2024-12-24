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
