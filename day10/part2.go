package day10

import (
	"mrnateriver.io/advent_of_code_2024/shared"
)

func SumTrailheadsRatings() int {
	grid, lenX, lenY := readInput()

	dp := make(map[pos]int)

	sum := 0
	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			if grid[y][x] == '0' {
				rating := countRating(grid, lenX, lenY, pos{x, y}, dp)
				sum += rating
			}
		}
	}

	return sum
}

func countRating(grid [][]byte, lenX, lenY int, start pos, dp map[pos]int) (score int) {
	cell := grid[start.Y][start.X]
	if cell == '9' {
		return 1
	}

	if score, ok := dp[start]; ok {
		return score
	}

	next := cell + 1
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == y || x == -y {
				continue
			}

			nextPos := pos{start.X + x, start.Y + y}
			if !shared.Point2dWithinBounds(nextPos, lenX, lenY) {
				continue
			}

			nextCell := grid[nextPos.Y][nextPos.X]
			if nextCell != next {
				continue
			}

			score += countRating(grid, lenX, lenY, nextPos, dp)
		}
	}

	dp[start] = score
	return score
}
