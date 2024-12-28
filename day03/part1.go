package day03

import (
	"fmt"
	"regexp"
	"strconv"

	"mrnateriver.io/advent_of_code_2024/shared"
)

func AddUpMultiplicationInstructions() int {
	input := shared.ReadInputLines("day03/input")

	rg := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)|do\\(\\)|don't\\(\\)")

	sum := 0

	for line := range input {
		matches := rg.FindAllStringSubmatch(line, -1)
		if matches != nil {
			for _, match := range matches {
				a, err := strconv.Atoi(match[1])
				if err != nil {
					panic(fmt.Errorf("invalid number %v for match[1]: %w", match[1], err))
				}
				b, err := strconv.Atoi(match[2])
				if err != nil {
					panic(fmt.Errorf("invalid number %v for match[2]: %w", match[2], err))
				}

				sum += a * b
			}
		}
	}

	return sum
}
