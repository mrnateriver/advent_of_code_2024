package day06

func CountPossibleObstructions() int {
	guard, grid := readGrid()

	obstructions := 0
	for p := range distinctSteps(grid, guard, DIR_UP) {
		if p == guard {
			continue
		}

		setGridAt(grid, p, '#')
		if looped(grid, guard, DIR_UP) {
			obstructions++
		}
		setGridAt(grid, p, '.')
	}

	return obstructions
}
