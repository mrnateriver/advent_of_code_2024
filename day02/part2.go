package day02

import (
	"fmt"
)

func CountSafeLevelsWithProblemDampener() int {
	input := readInput()

	safeLevels := 0

rowsLoop:
	for row := range input {
		if len(row) < 2 {
			panic(fmt.Errorf("invalid number of columns: %v", len(row)))
		}

		increasing := false

	colsLoop:
		for i := 0; i < len(row)-1; i++ {
			if i == 0 {
				increasing = row[i] < row[i+1]
			}

			cur, next := row[i], row[i+1]
			if !check(increasing, cur, next) {
				j := -1
				if i == 0 {
					j = 0
				}
				for ; j < 2; j++ {
					if checkRowIgnoringElementAt(row, i+j) {
						break colsLoop
					}
				}

				continue rowsLoop
			}
		}

		safeLevels++
	}

	return safeLevels
}

func checkRowIgnoringElementAt(row []int, ignoreIdx int) bool {
	if len(row) == 2 {
		return checkDiff(row[0], row[1])
	}

	setIncreasing := false
	increasing := false

	i, j := 0, 1
	for {
		if j == ignoreIdx {
			j++
		} else if i == ignoreIdx {
			i++
			j = i + 1
		}
		if j >= len(row) {
			break
		}

		if !setIncreasing {
			increasing = row[i] < row[j]
			setIncreasing = true
		}

		cur, next := row[i], row[j]
		if !check(increasing, cur, next) {
			return false
		}

		i = j
		j++
	}

	return true
}
