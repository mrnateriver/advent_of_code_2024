package day20

import (
	"mrnateriver.io/advent_of_code_2024/shared"
)

func CountTopCheatsOnRacetrack() int {
	grid, s, e := readInput()
	track := trace(grid, s, e)
	ch := cheats(grid, s, e, track)
	best := bestCheats(ch, 100)

	cheatStats(best)

	return len(best)
}

func cheats(grid [][]string, s, e shared.Point2d, track map[shared.Point2d]int) map[cheat]struct{} {
	lenX, lenY := len(grid[0]), len(grid)

	res := make(map[cheat]struct{})

	p, prev := s, shared.Point2d{}
	for p != e {
		var nextStep shared.Point2d
		for dir, next := range shared.Neighbours(p, false) {
			if !shared.Point2dWithinBounds(next, lenX, lenY) || next == prev {
				continue
			}

			if shared.GridAt(grid, next) == wall {
				curLen := track[p]

				overTheWall := shared.MoveInDir(next, dir)
				if shared.Point2dWithinBounds(overTheWall, lenX, lenY) && shared.GridAt(grid, overTheWall) != wall {
					if remLen, ok := track[overTheWall]; ok && remLen < curLen {
						save := curLen - remLen - 2 /* 1 for the wall, 1 for the next cell */
						res[cheat{next, overTheWall, save}] = struct{}{}
					}
				}
			} else {
				nextStep = next
			}
		}
		prev = p
		p = nextStep
	}

	return res
}

type cheat struct {
	s, e shared.Point2d
	save int
}
