package day06

import "mrnateriver.io/advent_of_code_2024/shared"

func CountPossibleObstructions() int {
	guard, grid := readGrid()

	obstructions := 0
	for p := range distinctSteps(grid, guard, shared.DirUp) {
		if p == guard {
			continue
		}

		shared.SetGridAt(grid, p, '#')
		if looped(grid, guard, shared.DirUp) {
			obstructions++
		}
		shared.SetGridAt(grid, p, '.')
	}

	return obstructions
}
