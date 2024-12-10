package day02

import (
	"fmt"
)

func CountSafeLevels() int {
	input := readInput()

	safeLevels := 0

rowsLoop:
	for row := range input {
		if len(row) < 2 {
			panic(fmt.Errorf("invalid number of columns: %v", len(row)))
		}

		increasing := false

		for i := 0; i < len(row)-1; i++ {
			if i == 0 {
				increasing = row[i] < row[i+1]
			}

			cur, next := row[i], row[i+1]
			if !check(increasing, cur, next) {
				continue rowsLoop
			}
		}

		safeLevels++
	}

	return safeLevels
}
