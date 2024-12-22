package day15

import (
	"fmt"
	"mrnateriver.io/advent_of_code_2024/shared"
	"time"
)

func Move1DBoxesAndSumPositions() int {
	grid, dirs, robot := readInput()

	lenX, lenY := len(grid[0]), len(grid)
	for _, dir := range dirs {
		robot = move1DBoxes(grid, lenX, lenY, robot, dir, false)
	}

	return countCoordsOf1DBoxes(grid)
}

func move1DBoxes(grid [][]string, lenX, lenY int, robot shared.Point2d, dir shared.Direction, debug bool) shared.Point2d {
	if debug {
		shared.PrintGrid(grid)
		fmt.Printf("%c\n", dirChar(dir))
		shared.MoveCursorUp(1)
		time.Sleep(60 * time.Millisecond)
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
	}

	return robot
}

func countCoordsOf1DBoxes(grid [][]string) int {
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
