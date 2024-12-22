package day15

import (
	"fmt"
	"mrnateriver.io/advent_of_code_2024/shared"
	"time"
)

func Move2DBoxesAndSumPositions() int {
	grid, dirs, robot := readInput()
	grid, robot = expandGrid(grid)

	lenX, lenY := len(grid[0]), len(grid)
	for _, dir := range dirs {
		robot = move2DBoxes(grid, lenX, lenY, robot, dir, false)
	}

	return countCoordsOf2DBoxes(grid)
}

func expandGrid(grid [][]string) ([][]string, shared.Point2d) {
	robotX, robotY := 0, 0
	newGrid := make([][]string, 0, len(grid))
	for i := range grid {
		row := make([]string, 0, len(grid[i])*2)
		for _, c := range grid[i] {
			switch c {
			case wall:
				row = append(row, wall, wall)
			case box:
				row = append(row, boxLeft, boxRight)
			case empty:
				row = append(row, empty, empty)
			default:
				robotX, robotY = len(row), i
				row = append(row, "@", empty) // shared.Colored(shared.ColorRed, "@"), empty)
			}
		}
		newGrid = append(newGrid, row)
	}
	return newGrid, shared.Point2d{robotX, robotY}
}

func move2DBoxes(grid [][]string, lenX, lenY int, robot shared.Point2d, dir shared.Direction, debug bool) shared.Point2d {
	if debug {
		debugGrid(grid, dir)
	}

	moved := false
	if dir == shared.DirUp || dir == shared.DirDown {
		moved = moveVert(grid, lenY, robot, dir)
	} else {
		moved = moveHor(grid, lenX, robot, dir)
	}

	if moved {
		robotVal := shared.GridAt(grid, robot)
		adjacent := shared.MoveInDir(robot, dir)
		shared.SetGridAt(grid, robot, empty)
		shared.SetGridAt(grid, adjacent, robotVal)

		if debug {
			debugGrid(grid, dir)
		}

		return adjacent
	}

	if debug {
		debugGrid(grid, dir)
	}

	return robot
}

func debugGrid(grid [][]string, dir shared.Direction) {
	shared.PrintGrid(grid)
	fmt.Printf("%c\n", dirChar(dir))
	shared.MoveCursorUp(1)
	shared.AwaitInput()
	time.Sleep(30 * time.Millisecond)
}

func moveHor(grid [][]string, lenX int, robot shared.Point2d, dir shared.Direction) (moved bool) {
	if dir != shared.DirLeft && dir != shared.DirRight {
		return
	}

	found := false
	dx := dir.X
	x := robot.X + dx
	for x >= 0 && x < lenX {
		c := shared.GridAt(grid, shared.Point2d{x, robot.Y})
		if c == wall {
			return
		}
		if c == empty {
			found = true
			break
		}
		x += dx
	}

	if found {
		for i := x; i != robot.X; i -= dx {
			shared.SetGridAt(grid, shared.Point2d{i, robot.Y}, shared.GridAt(grid, shared.Point2d{i - dx, robot.Y}))
		}

		return true
	}

	return
}

func moveVert(grid [][]string, lenY int, robot shared.Point2d, dir shared.Direction) (moved bool) {
	if dir != shared.DirUp && dir != shared.DirDown {
		return
	}

	dy := dir.Y
	y := robot.Y + dy
	seen := map[int]struct{}{robot.X: {}}
	shift := make([]shared.Point2d, 0, 4)
	shiftSeen := make(map[shared.Point2d]struct{})
	for y >= 0 && y < lenY {
		found := true

		keys := shared.Keys(seen)
		row := make([]shared.Point2d, 0, 4)
		rowSeen := make(map[int]struct{})
		for _, i := range keys {
			c := shared.GridAt(grid, shared.Point2d{i, y})
			if c == wall {
				return // If we reached a wall after a sequence of boxes, we're done
			}
			if c == empty {
				continue
			}

			delta := boxDelta(c)

			p1 := shared.Point2d{i, y}
			p2 := shared.Point2d{i + delta, y}

			if _, ok := shiftSeen[p1]; !ok {
				row = append(row, p1)
				shiftSeen[p1] = struct{}{}
			}
			if _, ok := shiftSeen[p2]; !ok {
				row = append(row, p2)
				shiftSeen[p2] = struct{}{}
			}

			rowSeen[i] = struct{}{}
			rowSeen[i+delta] = struct{}{}

			found = false
		}
		seen = rowSeen

		for _, r := range row {
			shift = append(shift, r)
		}

		if found {
			if len(shift) > 0 {
				break
			} else {
				return true
			}
		}

		y += dy
	}

	if len(shift) > 0 && y >= 0 && y < lenY {
		for i := len(shift) - 1; i >= 0; i-- {
			p := shift[i]
			c := shared.GridAt(grid, p)
			m := shared.MoveInDir(p, dir)

			shared.SetGridAt(grid, m, c)
			shared.SetGridAt(grid, p, empty)
		}

		return true
	}

	return
}

func boxDelta(char string) int {
	if char == boxLeft {
		return 1
	}
	return -1
}

func countCoordsOf2DBoxes(grid [][]string) int {
	count := 0

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == boxLeft {
				count += 100*(y+1) + (x + 2)
			}
		}
	}

	return count
}
