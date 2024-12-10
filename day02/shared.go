package day02

import (
	"fmt"
	"strconv"
	"strings"

	"mrnateriver.io/advent_of_code_2024/shared"
)

func readInput() <-chan []int {
	return shared.ReadInput("day02/input", func(line string) ([]int, error) {
		cols := strings.Split(line, " ")

		res := make([]int, len(cols))
		for i, col := range cols {
			val, err := strconv.Atoi(col)
			if err != nil {
				return []int{}, (fmt.Errorf("invalid column: %v; %w", col, err))
			}
			res[i] = val
		}

		return res, nil
	})
}

func checkDiff(prev, cur int) bool {
	diff := abs(prev - cur)
	if diff < 1 || diff > 3 {
		return false
	}

	return true
}

func check(increasing bool, prev, cur int) bool {
	if !checkDiff(prev, cur) {
		return false
	}

	if (prev > cur && increasing) || (prev < cur && !increasing) {
		return false
	}

	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
