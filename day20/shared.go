package day20

import (
	"fmt"
	"mrnateriver.io/advent_of_code_2024/shared"
)

func readInput() (grid [][]string, s, e shared.Point2d) {
	grid = make([][]string, 0, 8)

	for line := range shared.ReadInputLines("day20/input") {
		row := make([]string, 0, len(line))
		for x, c := range line {
			char := string(c)
			row = append(row, char)
			if char == start {
				s = shared.Point2d{x, len(grid)}
			} else if char == end {
				e = shared.Point2d{x, len(grid)}
			}
		}
		grid = append(grid, row)
	}

	return
}

func trace(grid [][]string, s, e shared.Point2d) map[shared.Point2d]int {
	lenX, lenY := len(grid[0]), len(grid)
	dp := make(map[shared.Point2d]int)
	dp[s] = 1 + measureTrack(grid, lenX, lenY, shared.Point2d{}, s, e, dp)
	return dp
}

func measureTrack(grid [][]string, lenX, lenY int, prev, from, to shared.Point2d, dp map[shared.Point2d]int) (nextLen int) {
	if from == to {
		return
	}

	for _, p := range shared.Neighbours(from, false) {
		if p == prev || !shared.Point2dWithinBounds(p, lenX, lenY) || shared.GridAt(grid, p) == wall {
			continue
		}
		nextLen = 1 + measureTrack(grid, lenX, lenY, from, p, to, dp)
		dp[p] = nextLen
		return nextLen
	}

	return
}

func bestCheats(ch []cheat, minSave int) []cheat {
	sortedCheats := shared.MakeHeap[cheat](func(a, b cheat) bool {
		return a.save > b.save
	})
	for _, c := range ch {
		sortedCheats.PushEntry(c)
	}

	bestCheats := make([]cheat, 0, 8)
	for sortedCheats.Len() > 0 {
		entry := sortedCheats.PopEntry()
		if entry.save < minSave {
			break
		}
		bestCheats = append(bestCheats, entry)
	}

	return bestCheats
}

func cheatStats(best []cheat) {
	keys := make([]int, 0)
	grouped := make(map[int][]cheat)
	for _, c := range best {
		if _, ok := grouped[c.save]; !ok {
			grouped[c.save] = make([]cheat, 0, 4)
			keys = append(keys, c.save)
		}
		grouped[c.save] = append(grouped[c.save], c)
	}

	for _, saves := range keys {
		cheats := grouped[saves]
		fmt.Printf("%d cheats save %d\n", len(cheats), saves)
	}
}

func cheats(grid [][]string, s, e shared.Point2d, maxCheatLen int, track map[shared.Point2d]int) []cheat {
	lenX, lenY := len(grid[0]), len(grid)

	grouped := make(map[cheat]struct{})

	p, prev := s, shared.Point2d{}
	for p != e {
		var nextStep shared.Point2d
		for _, next := range shared.Neighbours(p, false) {
			if !shared.Point2dWithinBounds(next, lenX, lenY) || next == prev {
				continue
			}

			if shared.GridAt(grid, next) != wall {
				nextStep = next
			}
		}

		ch := candidateCheats(grid, p, maxCheatLen, track)
		for _, c := range ch {
			grouped[c] = struct{}{}
		}

		prev = p
		p = nextStep
	}

	res := make([]cheat, 0, len(grouped))
	for c := range grouped {
		res = append(res, c)
	}

	return res
}

func candidateCheats(grid [][]string, p shared.Point2d, maxCheatLen int, track map[shared.Point2d]int) []cheat {
	// Scan a circle with diameter of maxCheatLen around p, since the farthest we can get with a cheat
	// is equal to maxCheatLen if we go in a straight direction from p
	// In practice, we're still going to check the full square around p, since it'll be cheaper than doing trigonometry
	lenX, lenY := len(grid[0]), len(grid)

	candidates := make([]cheat, 0, 8)

	curRemLen, ok := track[p]
	if !ok {
		return candidates
	}

	for y := p.Y - maxCheatLen; y <= p.Y+maxCheatLen; y++ {
		for x := p.X - maxCheatLen; x <= p.X+maxCheatLen; x++ {
			n := shared.Point2d{x, y}

			if n == p {
				continue
			}
			if !shared.Point2dWithinBounds(n, lenX, lenY) {
				continue
			}
			if shared.GridAt(grid, n) == wall {
				continue
			}

			cheatLen := abs(n.X-p.X) + abs(n.Y-p.Y)
			if cheatLen > maxCheatLen {
				continue
			}

			remLen, ok := track[n]
			if !ok || remLen >= curRemLen {
				continue
			}

			save := curRemLen - remLen - cheatLen
			if save > 0 {
				candidates = append(candidates, cheat{p, n, save})
			}
		}
	}

	return candidates
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type cheat struct {
	s, e shared.Point2d
	save int
}

const (
	empty = "."
	wall  = "#"
	start = "S"
	end   = "E"
)
