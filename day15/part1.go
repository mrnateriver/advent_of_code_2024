package day15

import (
	"fmt"
	"mrnateriver.io/advent_of_code_2024/shared"
	"time"
)

func MoveBoxesAndSumPositions() int {
	grid := make([][]string, 0, 8) // using []string for debug
	dirs := make([]shared.Direction, 0, 16)

	i := 0
	robotX, robotY := 0, 0
	lenX, lenY := 0, 0
	parsedBoxes := false
	for line := range shared.ReadInputLines("day15/input") {
		if line == "" {
			parsedBoxes = true
			continue
		}

		if !parsedBoxes {
			if justAWall(line) {
				continue
			}

			row := make([]string, 0, 16)
			for ix := 1; ix < len(line)-1; ix++ {
				c := line[ix]
				if c == '@' {
					robotX, robotY = i, ix-1
					row = append(row, shared.Colored(shared.ColorRed, "@"))
				} else {
					row = append(row, string(c))
				}
			}
			grid = append(grid, row)
			i++
		} else {
			for ix := range line {
				dir := parseDir(line[ix])
				dirs = append(dirs, dir)
			}
		}
	}

	dp := make(map[state]struct{})
	robot := shared.Point2d{robotX, robotY}
	lenX, lenY = len(grid[0]), len(grid)
	for _, dir := range dirs {
		robot = moveBoxes(grid, lenX, lenY, robot, dir, dp, false)
	}

	return countCoordsOfBoxes(grid)
}

func justAWall(line string) bool {
	for i := range line {
		if line[i] != wall[0] {
			return false
		}
	}
	return true
}

func parseDir(c byte) shared.Direction {
	switch c {
	case 'v':
		return shared.DirDown
	case '>':
		return shared.DirRight
	case '<':
		return shared.DirLeft
	case '^':
	default:
		return shared.DirUp
	}
	return shared.DirUp
}

func dirChar(dir shared.Direction) byte {
	switch dir {
	case shared.DirDown:
		return 'v'
	case shared.DirRight:
		return '>'
	case shared.DirLeft:
		return '<'
	case shared.DirUp:
	default:
		return '^'
	}
	return '^'
}

func moveBoxes(grid [][]string, lenX, lenY int, robot shared.Point2d, dir shared.Direction, dp map[state]struct{}, debug bool) shared.Point2d {
	if debug {
		shared.PrintGrid(grid)
		fmt.Printf("%c\n", dirChar(dir))
		shared.MoveCursorUp(1)
		//shared.AwaitInput()
		time.Sleep(60 * time.Millisecond)
	}

	if _, ok := dp[state{robot, dir}]; ok {
		return robot
	}

	next := shared.MoveInDir(robot, dir)
	if !shared.Point2dWithinBounds(next, lenX, lenY) {
		return robot
	}
	nextVal := shared.GridAt(grid, next)

	found := false
	pos := shared.Point2d{}
	for nextVal != wall {
		if nextVal == empty {
			pos = next
			found = true
			break
		}

		next = shared.MoveInDir(next, dir)
		if !shared.Point2dWithinBounds(next, lenX, lenY) {
			break
		}
		nextVal = shared.GridAt(grid, next)
	}

	if found {
		robotVal := shared.GridAt(grid, robot)

		shared.SetGridAt(grid, pos, box)
		shared.SetGridAt(grid, robot, empty)

		adjacent := shared.MoveInDir(robot, dir)
		shared.SetGridAt(grid, adjacent, robotVal)

		return adjacent
	} else {
		//dp[state{robot, dir}] = struct{}{}
	}

	return robot
}

func countCoordsOfBoxes(grid [][]string) int {
	count := 0

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == box {
				count += 100*(y+1) + (x + 1)
			}
		}
	}

	return count
}

type state struct {
	from shared.Point2d
	dir  shared.Direction
}

const (
	empty = "."
	wall  = "#"
	box   = "O"
)
