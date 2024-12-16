package day08

import (
	"mrnateriver.io/advent_of_code_2024/shared"
	"time"
)

func CountAntinodes() int {
	grid := make([][]string, 0, 32)

	freqs := make(map[byte][]pos)
	antinodes := make(map[pos]struct{})

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

	lenY := y
	lenX := x + 1

	shared.PrintGrid(grid)

	for freq, positions := range freqs {
		if len(positions) < 2 {
			continue
		}

		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				a := positions[i]
				b := positions[j]
				dist := shared.DistanceBetweenPoints(a, b)

				grid[a.Y][a.X] = shared.Colored(shared.ColorGreen, "1")
				grid[b.Y][b.X] = shared.Colored(shared.ColorYellow, "2")

				antinodeA := shared.PointAlongLineAfterB(b, a, dist)
				if shared.PointWithinBounds(antinodeA, lenX, lenY) {
					grid[antinodeA.Y][antinodeA.X] = shared.Colored(shared.ColorRed, "X")
					shared.PrintGrid(grid)
					time.Sleep(100 * time.Millisecond)

					antinodes[antinodeA] = struct{}{}
				}

				antinodeB := shared.PointAlongLineAfterB(a, b, dist)
				if shared.PointWithinBounds(antinodeB, lenX, lenY) {
					grid[antinodeB.Y][antinodeB.X] = shared.Colored(shared.ColorRed, "X")
					shared.PrintGrid(grid)
					time.Sleep(100 * time.Millisecond)

					antinodes[antinodeB] = struct{}{}
				}

				grid[a.Y][a.X] = string(freq)
				grid[b.Y][b.X] = string(freq)
			}
		}
	}

	count := 0
	for range antinodes {
		count++
	}

	return count
}

type pos = shared.Point2d
