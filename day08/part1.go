package day08

import (
	"mrnateriver.io/advent_of_code_2024/shared"
	"time"
)

func CountAntinodes() int {
	freqs, grid := readInput()

	lenY := len(grid)
	lenX := len(grid[0])

	shared.PrintGrid(grid)

	antinodes := make(map[pos]struct{})
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
				if shared.Point2dWithinBounds(antinodeA, lenX, lenY) {
					grid[antinodeA.Y][antinodeA.X] = shared.Colored(shared.ColorRed, "X")
					shared.PrintGrid(grid)
					time.Sleep(100 * time.Millisecond)

					antinodes[antinodeA] = struct{}{}
				}

				antinodeB := shared.PointAlongLineAfterB(a, b, dist)
				if shared.Point2dWithinBounds(antinodeB, lenX, lenY) {
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
