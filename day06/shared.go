package day06

import (
	"iter"

	"mrnateriver.io/advent_of_code_2024/shared"
)

func readGrid() (guard pos, grid [][]byte) {
	guardX, guardY := -1, -1
	grid = make([][]byte, 0, 32)

	for line := range shared.ReadInputLines("day06/input") {
		grid = append(grid, []byte(line))
		if guardX < 0 {
			for x, char := range line {
				if char == '^' {
					guardX = x
					break
				}
			}
			guardY++
		}
	}

	guard = pos{guardX, guardY}

	return
}

func distinctSteps(grid [][]byte, from pos, d dir) iter.Seq[pos] {
	return func(yield func(pos) bool) {
		dp := make(map[pos]struct{})

		for p := range traverse(grid, from, d) {
			if _, ok := dp[p]; !ok {
				if !yield(p) {
					break
				}
				dp[p] = struct{}{}
			}
		}
	}
}

func looped(grid [][]byte, from pos, d dir) bool {
	dp := make(map[nav]struct{})

	for p, d := range traverse(grid, from, d) {
		if seen(p, d, dp) {
			return true
		}

		dp[nav{p, d}] = struct{}{}
	}

	return false
}

func traverse(grid [][]byte, p pos, d dir) iter.Seq2[pos, dir] {
	return func(yield func(pos, dir) bool) {
		lenX := len(grid[0])
		lenY := len(grid)

		for {
			nextPos := shared.MoveInDir(p, d)
			if !shared.Point2dWithinBounds(nextPos, lenX, lenY) {
				break
			}

			next := shared.GridAt(grid, nextPos)
			if next == '#' {
				d = rotateDirClockwise(d)
			} else {
				if !yield(nextPos, d) {
					return
				}

				p = nextPos
			}
		}
	}
}

func rotateDirClockwise(d dir) dir {
	if d == shared.DirUp {
		return shared.DirRight
	} else if d == shared.DirRight {
		return shared.DirDown
	} else if d == shared.DirDown {
		return shared.DirLeft
	} else if d == shared.DirLeft {
		return shared.DirUp
	}

	return dir{}
}

func seen(pos pos, d dir, dp map[nav]struct{}) bool {
	_, ok := dp[nav{pos, d}]
	return ok
}

type dir = shared.Point2d

type pos = dir

type nav struct {
	p pos
	d dir
}
