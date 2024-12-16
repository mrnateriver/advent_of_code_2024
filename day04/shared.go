package day04

import (
	"fmt"

	"mrnateriver.io/advent_of_code_2024/shared"
)

func countPatterns(bufferLinesAhead int, counter func(lines []string, y int) int) int {
	input := shared.ReadInputLines("day04/input")

	lineLength := 0
	lines := make([]string, 0, 12)
	processed := 0
	cursor := 0
	foundWords := 0

	for line := range input {
		curLineLength := len(line)
		if lineLength == 0 {
			lineLength = curLineLength
		} else if lineLength != curLineLength {
			panic(fmt.Errorf("line %d has length %d, expected %d", processed, curLineLength, lineLength))
		}

		lines = append(lines, line)
		processed++
		if processed < bufferLinesAhead {
			continue
		}

		foundWords += counter(lines, cursor)
		cursor++
	}

	for ; cursor < processed; cursor++ {
		foundWords += counter(lines, cursor)
	}

	return foundWords
}
