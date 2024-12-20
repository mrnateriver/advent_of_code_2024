package day06

import "mrnateriver.io/advent_of_code_2024/shared"

func CountDistinctGuardPositions() int {
	guard, grid := readGrid()

	return countDistinctSteps(grid, guard, shared.DirUp)
}

func countDistinctSteps(grid [][]byte, from pos, d dir) int {
	steps := 0

	for range distinctSteps(grid, from, d) {
		steps++
	}

	return steps
}
