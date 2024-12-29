package day18

import (
	"mrnateriver.io/advent_of_code_2024/shared"
)

func MeasureShortestPathAfterBytefall() int {
	const gridLenX = 71
	const gridLenY = 71
	const iterations = 1024

	grid := make([][]string, gridLenY)
	for i := 0; i < gridLenY; i++ {
		grid[i] = make([]string, gridLenX)
	}

	i := 0
	for pt := range readInput() {
		grid[pt.Y][pt.X] = wall
		i++
		if i == iterations {
			break
		}
	}

	return shared.FindShortestPathLength(grid, shared.Point2d{0, 0}, shared.Point2d{gridLenX - 1, gridLenY - 1}, wall)
}
