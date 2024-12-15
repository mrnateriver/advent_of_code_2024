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
			nextPos := moveInDir(p, d)
			if outOfBounds(nextPos, lenX, lenY) {
				break
			}

			next := gridAt(grid, nextPos)
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
	if d == DIR_UP {
		return DIR_RIGHT
	} else if d == DIR_RIGHT {
		return DIR_DOWN
	} else if d == DIR_DOWN {
		return DIR_LEFT
	} else if d == DIR_LEFT {
		return DIR_UP
	}

	return dir{}
}

func seen(pos pos, d dir, dp map[nav]struct{}) bool {
	_, ok := dp[nav{pos, d}]
	return ok
}

func moveInDir(p pos, d dir) pos {
	return pos{p.x + d.x, p.y + d.y}
}

func outOfBounds(p pos, lx, ly int) bool {
	return p.x < 0 || p.x >= lx || p.y < 0 || p.y >= ly
}

func gridAt(grid [][]byte, p pos) byte {
	return grid[p.y][p.x]
}

func setGridAt(grid [][]byte, p pos, value byte) {
	grid[p.y][p.x] = value
}

type dir struct {
	x, y int
}

type pos = dir

type nav struct {
	p pos
	d dir
}

var (
	DIR_UP    = dir{0, -1}
	DIR_RIGHT = dir{1, 0}
	DIR_DOWN  = dir{0, 1}
	DIR_LEFT  = dir{-1, 0}
)
