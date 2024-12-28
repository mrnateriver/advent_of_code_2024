package day14

import (
	"fmt"
	"mrnateriver.io/advent_of_code_2024/shared"
	"regexp"
	"strconv"
)

func readInput() []robot {
	robotRegex := regexp.MustCompile("p=(\\d+),(\\d+) v=(-?\\d+),(-?\\d+)")

	robots := make([]robot, 0, 2)

	for line := range shared.ReadInputLines("day14/input") {
		matches := robotRegex.FindStringSubmatch(line)
		if matches == nil {
			continue
		}

		x, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(fmt.Errorf("failed to parse robot x: %w", err))
		}

		y, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(fmt.Errorf("failed to parse robot y: %w", err))
		}

		dx, err := strconv.Atoi(matches[3])
		if err != nil {
			panic(fmt.Errorf("failed to parse robot dx: %w", err))
		}

		dy, err := strconv.Atoi(matches[4])
		if err != nil {
			panic(fmt.Errorf("failed to parse robot dy: %w", err))
		}

		robots = append(robots, robot{x, y, dx, dy})
	}

	return robots
}

func (r *robot) move(lenX, lenY, steps int) {
	r.x = (r.x + r.dx*steps) % lenX
	r.y = (r.y + r.dy*steps) % lenY
	if r.x < 0 {
		r.x += lenX
	}
	if r.y < 0 {
		r.y += lenY
	}
}

func splitIntoQuadrants(robots []robot, lenX, lenY int) (q1, q2, q3, q4 []robot) {
	xEven := lenX%2 == 0
	yEven := lenY%2 == 0
	halfX := lenX / 2
	halfY := lenY / 2

	q1 = make([]robot, 0, 2)
	q2 = make([]robot, 0, 2)
	q3 = make([]robot, 0, 2)
	q4 = make([]robot, 0, 2)

	for _, r := range robots {
		if r.x < halfX {
			if r.y < halfY {
				q1 = append(q1, r)
			} else if yEven || r.y > halfY {
				q3 = append(q3, r)
			}
		} else if xEven || r.x > halfX {
			if r.y < halfY {
				q2 = append(q2, r)
			} else if yEven || r.y > halfY {
				q4 = append(q4, r)
			}
		}
	}

	return
}

func calcSafetyFactor(robots []robot, lenX, lenY int) int {
	q1, q2, q3, q4 := splitIntoQuadrants(robots, lenX, lenY)
	return len(q1) * len(q2) * len(q3) * len(q4)
}

func overlappingRobots(robots []robot) bool {
	m := make(map[shared.Point2d]struct{})
	for _, r := range robots {
		p := shared.Point2d{r.x, r.y}
		if _, ok := m[p]; ok {
			return true
		}
		m[p] = struct{}{}
	}
	return false
}

func printGrid(robots []robot, grid [][]string) {
	for y := range grid {
		for x := range grid[y] {
			grid[y][x] = "."
		}
	}
	for _, r := range robots {
		grid[r.y][r.x] = shared.Colored(shared.ColorRed, "#")
	}

	shared.PrintGrid(grid)
	shared.AwaitInput()
}

type robot struct {
	x, y, dx, dy int
}
