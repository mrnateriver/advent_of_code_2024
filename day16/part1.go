package day16

import (
	"fmt"

	"mrnateriver.io/advent_of_code_2024/shared"
)

func FindPathWithLowestScore() int {
	grid := make([][]string, 0, 8)

	var sp, ep shared.Point2d

	y := 0
	input := shared.ReadInputLines("day16/input")
	for line := range input {
		if justAWall(line) {
			continue
		}

		row := make([]string, 0, len(line)-2)
		for x := 1; x < len(line)-1; x++ {
			cc := string(line[x])
			row = append(row, cc)
			if cc == start {
				sp = shared.Point2d{x - 1, y}
			} else if cc == end {
				ep = shared.Point2d{x - 1, y}
			}
		}

		grid = append(grid, row)
		y++
	}

	lenX, lenY := len(grid[0]), len(grid)

	minScore := findShortestPath(grid, lenX, lenY, sp, ep)

	return minScore
}

func findShortestPath(grid [][]string, lenX, lenY int, start, end shared.Point2d) int {
	seen := make(map[nav]int)
	queue := shared.MakePriorityQueue[path]()
	queue.PushEntry(path{nav: nav{start, shared.DirRight}, cost: 0}, 0)

	for queue.Len() > 0 {
		entry := queue.PollEntry()
		key := nav{entry.start, entry.dir}

		if entry.start == end {
			return entry.cost
		} else if seenCost, ok := seen[key]; !ok || seenCost > entry.cost {
			seen[key] = entry.cost

			next := shared.MoveInDir(entry.start, entry.dir)
			if shared.Point2dWithinBounds(next, lenX, lenY) && shared.GridAt(grid, next) != wall {
				newCost := entry.cost + 1
				queue.PushEntry(path{nav: nav{next, entry.dir}, cost: newCost}, newCost)
			}

			rotateCost := entry.cost + 1000
			cwDir := shared.RotateDir(entry.dir, true)
			acwDir := shared.RotateDir(entry.dir, false)
			queue.PushEntry(path{nav: nav{entry.start, cwDir}, cost: rotateCost}, rotateCost)
			queue.PushEntry(path{nav: nav{entry.start, acwDir}, cost: rotateCost}, rotateCost)
		}
	}

	panic(fmt.Errorf("no path found"))
}

type nav struct {
	start shared.Point2d
	dir   shared.Direction
}

type path struct {
	nav
	cost int
}
