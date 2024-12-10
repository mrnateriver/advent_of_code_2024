package day01

import (
	"fmt"
	"strings"

	"mrnateriver.io/advent_of_code_2024/shared"
)

func readInput() <-chan [2]string {
	return shared.ReadInput("day01/input", func(line string) ([2]string, error) {
		cols := strings.Split(line, "   ")

		if len(cols) != 2 {
			return [2]string{}, (fmt.Errorf("invalid number of columns: %v", len(cols)))
		}

		return [2]string{cols[0], cols[1]}, nil
	})
}
