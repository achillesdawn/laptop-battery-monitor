package stats

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) ([]byte, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	return b, nil
}

func readString(path string) (string, error) {
	b, err := readFile(path)
	if err != nil {
		return "", err
	}

	s := strings.TrimSpace(string(b))

	return s, nil
}

func readInt(path string) (int, error) {
	s, err := readString(path)
	if err != nil {
		return 0, err
	}

	raw, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("could not convert to int: %w", err)
	}

	return raw, nil
}

func readFloat(path string) (float32, error) {
	i, err := readInt(path)
	if err != nil {
		return 0, err
	}

	result := float32(i) / 1_000_000

	return result, nil
}
