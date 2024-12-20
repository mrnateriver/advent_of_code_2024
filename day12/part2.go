package day12

import (
	"mrnateriver.io/advent_of_code_2024/shared"
)

func CalcFencePriceWithSides() int {
	grid := readInput()
	lenX, lenY := len(grid[0]), len(grid)
	dp := make(map[pos]struct{})

	sum := 0
	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			s := scanSides(grid, lenX, lenY, pos{x, y}, dp)
			sum += s.area * s.perim
		}
	}

	return sum
}

func scanSides(grid [][]byte, lenX, lenY int, p pos, dp map[pos]struct{}) stats {
	c := shared.GridAt(grid, p)
	if _, ok := dp[p]; ok {
		return stats{}
	}

	dp[p] = struct{}{}

	area, sides := 1, countCorners(grid, lenX, lenY, p)
	for _, np := range shared.Neighbours(p, false) {
		if !shared.Point2dWithinBounds(np, lenX, lenY) {
			continue
		} else {
			nval := shared.GridAt(grid, np)
			if nval == c {
				s := scanSides(grid, lenX, lenY, np, dp)
				area += s.area
				sides += s.perim
			}
		}
	}

	return stats{area, sides}
}

func countCorners(grid [][]byte, lenX, lenY int, p shared.Point2d) (count int) {
	t := shared.GridAt(grid, p)

	for i := 0; i < len(dirs)-1; i++ {
		s1p, s2p, cp := shared.MoveInDir(p, dirs[i]), shared.MoveInDir(p, dirs[i+1]), shared.MoveInDir(shared.MoveInDir(p, dirs[i]), dirs[i+1])

		var s1, s2, c byte
		if shared.Point2dWithinBounds(s1p, lenX, lenY) {
			s1 = shared.GridAt(grid, s1p)
		}
		if shared.Point2dWithinBounds(s2p, lenX, lenY) {
			s2 = shared.GridAt(grid, s2p)
		}
		if shared.Point2dWithinBounds(cp, lenX, lenY) {
			c = shared.GridAt(grid, cp)
		}

		if (t != s1 && t != s2) || (s1 == t && s2 == t && c != t) {
			count++
		}
	}

	return
}

var dirs = []shared.Point2d{shared.DirUp, shared.DirRight, shared.DirDown, shared.DirLeft, shared.DirUp}
