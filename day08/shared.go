package day08

import "mrnateriver.io/advent_of_code_2024/shared"

func readInput() (freqs map[byte][]pos, grid [][]string) {
	freqs = make(map[byte][]pos)
	grid = make([][]string, 0, 32)

	y, x := 0, 0
	for line := range shared.ReadInputLines("day08/input") {
		row := make([]string, len(line))
		for x = range line {
			row[x] = string(line[x])
			if line[x] == '.' {
				continue
			}

			freqs[line[x]] = append(freqs[line[x]], pos{x, y})
		}
		grid = append(grid, row)
		y++
	}

	return
}

type pos = shared.Point2d
