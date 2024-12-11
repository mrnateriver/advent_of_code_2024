package day03

import (
	"fmt"
	"regexp"
	"strconv"

	"mrnateriver.io/advent_of_code_2024/shared"
)

func AddUpMultiplicationInstructionsWithConditions() int {
	input := shared.ReadInput("day03/input", func(arg string) (string, error) {
		return arg, nil
	})

	rg, err := regexp.Compile("mul\\((\\d+),(\\d+)\\)|do\\(\\)|don't\\(\\)")
	if err != nil {
		panic(err)
	}

	sum := 0
	do := true

	for line := range input {
		matches := rg.FindAllStringSubmatch(line, -1)
		if matches != nil {
			for _, match := range matches {
				switch match[0] {
				case "do()":
					do = true

				case "don't()":
					do = false

				default:
					if !do {
						continue
					}

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
	}

	return sum
}
