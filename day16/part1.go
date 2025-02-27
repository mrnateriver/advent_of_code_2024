package day16

import (
	"fmt"

	"mrnateriver.io/advent_of_code_2024/shared"
)

func FindPathWithLowestScore() int {
	grid, sp, ep := readInput()

	lenX, lenY := len(grid[0]), len(grid)

	minScore := findShortestPath(grid, lenX, lenY, sp, ep)

	return minScore
}

func findShortestPath(grid [][]string, lenX, lenY int, start, end shared.Point2d) int {
	seen := make(map[nav]int)
	queue := shared.MakePriorityQueue[path]()
	queue.PushEntry(path{steps: []shared.Point2d{start}, dir: shared.DirRight, cost: 0}, 0)

	for queue.Len() > 0 {
		path := queue.PollEntry()
		key := path.key()

		if path.end() == end {
			return path.cost
		} else if seenCost, ok := seen[key]; !ok || seenCost > path.cost {
			seen[key] = path.cost

			newCost := path.cost + 1
			next := path.moveInDir(path.dir, newCost)
			if shared.Point2dWithinBounds(next.end(), lenX, lenY) && shared.GridAt(grid, next.end()) != wall {
				queue.PushEntry(next, newCost)
			}

			rotateCost := path.cost + 1000
			cwPath := path.rotate(true, rotateCost)
			ccwPath := path.rotate(false, rotateCost)
			queue.PushEntry(cwPath, rotateCost)
			queue.PushEntry(ccwPath, rotateCost)
		}
	}

	panic(fmt.Errorf("no path found"))
}
