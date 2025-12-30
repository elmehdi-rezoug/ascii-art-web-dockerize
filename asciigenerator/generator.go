package asciigenerator

import (
	"errors"
	"strings"
)

func GenerateAsciiArt(input string, banner string) (string, error) {
	//validate banner
	bannerPath := "banners/" + banner + ".txt"
	bannerText, err := readFile(bannerPath)

	if err != nil {
		return "", errors.New("500")
	}
	bannerText = strings.ReplaceAll(bannerText, "\r", "")
	bannerLines := strings.Split(bannerText, "\n")

	//validate input
	if input == "" {
		return "", nil
	}
	if len(input) > 10000 {
		return "", errors.New("400 Input is too large")
	}
	// handle literal newlines "\n"
	input = strings.ReplaceAll(input, "\r", "")
	if isOnlyNewline(input) {
		return input, nil
	}

	inputLines := strings.Split(input, "\n")
	var result string
	const (
		AsciiOffset = 32
		CharHeight  = 9
	)

	for _, line := range inputLines {
		isValidInput := isValidInput(line)
		if !isValidInput {
			return "", errors.New("400 Only ascii characters 32-126 are allowed")
		}
		if line == "" {
			result += "\n"
			continue
		}
		// loop for height - 8 lines per character
		for height := 1; height < 9; height++ {
			for _, char := range line {
				asciiStartLine := (int(char) - AsciiOffset) * CharHeight
				lineIndex := asciiStartLine + height
				result += bannerLines[lineIndex]
			}
			result += "\n"
		}
	}
	return result, nil
}
