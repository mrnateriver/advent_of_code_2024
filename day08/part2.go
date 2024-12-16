package day08

import (
	"fmt"
	"mrnateriver.io/advent_of_code_2024/shared"
	"time"
)

func CountMultiAntinodes() int {
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

				antinodes[a] = struct{}{}
				antinodes[b] = struct{}{}

				grid[a.Y][a.X] = shared.Colored(shared.ColorGreen, "1")
				grid[b.Y][b.X] = shared.Colored(shared.ColorYellow, "2")

				antinodeA := shared.PointAlongLineAfterB(b, a, dist)
				for m := 2; shared.PointWithinBounds(antinodeA, lenX, lenY); m++ {
					grid[antinodeA.Y][antinodeA.X] = shared.Colored(shared.ColorRed, "X")
					outputIteration(grid, antinodeA)

					antinodes[antinodeA] = struct{}{}
					antinodeA = shared.PointAlongLineAfterB(b, a, dist*float64(m))
				}

				antinodeB := shared.PointAlongLineAfterB(a, b, dist)
				for m := 2; shared.PointWithinBounds(antinodeB, lenX, lenY); m++ {
					grid[antinodeB.Y][antinodeB.X] = shared.Colored(shared.ColorRed, "X")
					outputIteration(grid, antinodeB)

					antinodes[antinodeB] = struct{}{}
					antinodeB = shared.PointAlongLineAfterB(a, b, dist*float64(m))
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

func outputIteration(grid [][]string, node pos) {
	shared.PrintGrid(grid)
	outputNode(node)
	time.Sleep(100 * time.Millisecond)
}

func outputNode(node pos) {
	fmt.Printf("                                            \n") // clear the line
	shared.MoveCursorUp(1)
	fmt.Printf("node: %v\n", node)
	shared.MoveCursorUp(1)
}
