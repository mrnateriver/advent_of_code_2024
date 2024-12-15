package day06

func CountDistinctGuardPositions() int {
	guard, grid := readGrid()

	return countDistinctSteps(grid, guard, DIR_UP)
}

func countDistinctSteps(grid [][]byte, from pos, d dir) int {
	steps := 0

	for range distinctSteps(grid, from, d) {
		steps++
	}

	return steps
}
