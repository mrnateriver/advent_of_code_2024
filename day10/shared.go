package day10

import "mrnateriver.io/advent_of_code_2024/shared"

func readInput() (grid [][]byte, lenX, lenY int) {
	grid = make([][]byte, 0, 8)

	input := shared.ReadInput("day10/input", func(line string) ([]byte, error) {
		return []byte(line), nil
	})
	for line := range input {
		grid = append(grid, line)
	}

	lenX = len(grid[0])
	lenY = len(grid)

	return
}

type pos = shared.Point2d
