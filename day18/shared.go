package day18

import (
	"fmt"
	"mrnateriver.io/advent_of_code_2024/shared"
	"strconv"
	"strings"
)

func readInput() <-chan shared.Point2d {
	return shared.ReadInput("day18/input", func(line string) (shared.Point2d, error) {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return shared.Point2d{}, fmt.Errorf("invalid input line %v", line)
		}

		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return shared.Point2d{}, fmt.Errorf("failed to parse x: %w", err)
		}

		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return shared.Point2d{}, fmt.Errorf("failed to parse y: %w", err)
		}

		return shared.Point2d{x, y}, nil
	})
}

const (
	empty = "."
	wall  = "#"
	path  = "O"
)
