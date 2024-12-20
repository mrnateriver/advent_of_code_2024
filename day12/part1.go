package day12

import (
	"mrnateriver.io/advent_of_code_2024/shared"
)

func CalcFencePrice() int {
	input := shared.ReadInputLines("day12/input")

	grid := make([][]byte, 0, 8)
	for line := range input {
		grid = append(grid, []byte(line))
	}

	lenX, lenY := len(grid[0]), len(grid)

	dp := make(map[cell]struct{})

	sum := 0
	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			s := scan(grid, x, y, lenX, lenY, dp)
			sum += s.area * s.perim
		}
	}

	return sum
}

func scan(grid [][]byte, x, y, lenX, lenY int, dp map[cell]struct{}) stats {
	c := grid[y][x]
	cl := cell{x, y, c}
	if _, ok := dp[cl]; ok {
		return stats{}
	}

	dp[cl] = struct{}{}

	area, perim := 1, 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == dy || dx == -dy {
				continue
			}

			nx, ny := x+dx, y+dy
			if !shared.CoordsWithinBounds(nx, ny, lenX, lenY) {
				perim++
				continue
			} else {
				nval := grid[ny][nx]
				if nval != c {
					perim++
				} else {
					s := scan(grid, nx, ny, lenX, lenY, dp)
					area += s.area
					perim += s.perim
				}
			}
		}
	}

	return stats{area, perim}
}

type cell struct {
	x, y int
	c    byte
}

type stats struct {
	area, perim int
}
