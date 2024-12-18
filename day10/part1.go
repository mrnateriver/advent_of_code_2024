package day10

import (
	"mrnateriver.io/advent_of_code_2024/shared"
)

func SumTrailheadsScores() int {
	grid, lenX, lenY := readInput()

	dp := make(map[pos][]pos)

	sum := 0
	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			if grid[y][x] == '0' {
				summits := countScore(grid, lenX, lenY, pos{x, y}, dp)
				sum += len(summits)
			}
		}
	}

	return sum
}

func countScore(grid [][]byte, lenX, lenY int, start pos, dp map[pos][]pos) []pos {
	cell := grid[start.Y][start.X]
	if cell == '9' {
		return []pos{start}
	}

	if positions, ok := dp[start]; ok {
		return positions
	}

	uniquePositions := make(map[pos]struct{})

	next := cell + 1
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == y || x == -y {
				continue
			}

			nextPos := pos{start.X + x, start.Y + y}
			if !shared.PointWithinBounds(nextPos, lenX, lenY) {
				continue
			}

			nextCell := grid[nextPos.Y][nextPos.X]
			if nextCell != next {
				continue
			}

			reachable := countScore(grid, lenX, lenY, nextPos, dp)
			for _, p := range reachable {
				uniquePositions[p] = struct{}{}
			}
		}
	}

	dpPositions := make([]pos, 0, len(uniquePositions))
	for p := range uniquePositions {
		dpPositions = append(dpPositions, p)
	}

	dp[start] = dpPositions
	return dpPositions
}
