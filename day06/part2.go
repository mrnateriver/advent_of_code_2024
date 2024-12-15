package day06

func CountPossibleObstructions() int {
	guard, grid := readGrid()

	obstructions := 0
	for p := range distinctSteps(grid, guard, DirUp) {
		if p == guard {
			continue
		}

		setGridAt(grid, p, '#')
		if looped(grid, guard, DirUp) {
			obstructions++
		}
		setGridAt(grid, p, '.')
	}

	return obstructions
}
