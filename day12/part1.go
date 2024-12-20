package day12

import (
	"mrnateriver.io/advent_of_code_2024/shared"
)

func CalcFencePrice() int {
	grid := readInput()

	lenX, lenY := len(grid[0]), len(grid)

	dp := make(map[pos]struct{})

	sum := 0
	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			s := scan(grid, lenX, lenY, pos{x, y}, dp)
			sum += s.area * s.perim
		}
	}

	return sum
}

func scan(grid [][]byte, lenX, lenY int, p pos, dp map[pos]struct{}) stats {
	c := shared.GridAt(grid, p)
	if _, ok := dp[p]; ok {
		return stats{}
	}

	dp[p] = struct{}{}

	area, perim := 1, 0
	for _, np := range shared.Neighbours(p, false) {
		if !shared.Point2dWithinBounds(np, lenX, lenY) {
			perim++
		} else {
			nval := shared.GridAt(grid, np)
			if nval != c {
				perim++
			} else {
				s := scan(grid, lenX, lenY, np, dp)
				area += s.area
				perim += s.perim
			}
		}
	}

	return stats{area, perim}
}
