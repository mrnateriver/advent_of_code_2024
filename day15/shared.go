package day15

import "mrnateriver.io/advent_of_code_2024/shared"

func readInput() (grid [][]string, dirs []shared.Direction, robot shared.Point2d) {
	grid = make([][]string, 0, 8) // using []string for debug
	dirs = make([]shared.Direction, 0, 16)

	i := 0
	robotX, robotY := 0, 0
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

	robot = shared.Point2d{robotX, robotY}
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

const (
	empty    = "."
	wall     = "#"
	box      = "O"
	boxLeft  = "["
	boxRight = "]"
)
