package asciigenerator

import (
	"errors"
	"os"
)

func readFile(filename string) (string, error) {
	textBytes, err := os.ReadFile(filename)
	if err != nil {
		return "", errors.New("error reading banner file")
	}
	text := string(textBytes)
	return text, nil
}

func isValidInput(input string) bool {
	for _, val := range input {
		if val < 32 || val > 126 {
			return false 
		}
	}
	return true
}

func isOnlyNewline(s string) bool {
	for _, char := range s {
		if char != '\n' {
			return false
		}
	}
	return true
}
