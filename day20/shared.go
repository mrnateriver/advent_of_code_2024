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

func bestCheats(ch map[cheat]struct{}, minSave int) []cheat {
	sortedCheats := shared.MakeHeap[cheat](func(a, b cheat) bool {
		return a.save > b.save
	})
	for c := range ch {
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

const (
	empty = "."
	wall  = "#"
	start = "S"
	end   = "E"
)
