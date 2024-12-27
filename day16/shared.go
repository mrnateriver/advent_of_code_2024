package day16

import (
	"fmt"

	"mrnateriver.io/advent_of_code_2024/shared"
)

func readInput() (grid [][]string, sp, ep shared.Point2d) {
	grid = make([][]string, 0, 8)

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

	return
}

func justAWall(line string) bool {
	for i := range line {
		if line[i] != wall[0] {
			return false
		}
	}
	return true
}

const (
	empty = "."
	wall  = "#"
	start = "S"
	end   = "E"
)

type nav struct {
	start shared.Point2d
	dir   shared.Direction
}

type path struct {
	steps []shared.Point2d
	dir   shared.Direction
	cost  int
}

func (p path) end() shared.Point2d {
	if len(p.steps) == 0 {
		panic(fmt.Errorf("attempt to get end of an empty path"))
	}
	return p.steps[len(p.steps)-1]
}

func (p path) key() nav {
	return nav{p.end(), p.dir}
}

func (p path) moveInDir(dir shared.Direction, newCost int) path {
	newSteps := make([]shared.Point2d, len(p.steps), len(p.steps)+1)
	copy(newSteps, p.steps)

	return path{
		steps: append(newSteps, shared.MoveInDir(p.end(), dir)),
		dir:   dir,
		cost:  newCost,
	}
}

func (p path) rotate(clockwise bool, newCost int) path {
	return path{p.steps, shared.RotateDir(p.dir, clockwise), newCost}
}

func countRotations(d1, d2 shared.Direction) (rotations int) {
	if (d1 == shared.DirUp || d1 == shared.DirDown) && (d2 == shared.DirLeft || d2 == shared.DirRight) {
		return 1
	}
	if (d1 == shared.DirUp && d2 == shared.DirDown) || (d1 == shared.DirDown && d2 == shared.DirUp) {
		return 2
	}
	return
}
