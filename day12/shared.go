package day12

import "mrnateriver.io/advent_of_code_2024/shared"

func readInput() [][]byte {
	input := shared.ReadInputLines("day12/input")

	grid := make([][]byte, 0, 8)
	for line := range input {
		grid = append(grid, []byte(line))
	}

	return grid
}

type pos = shared.Point2d

type stats struct {
	area, perim int
}
