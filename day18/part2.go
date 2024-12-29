package day18

import (
	"fmt"
	"mrnateriver.io/advent_of_code_2024/shared"
)

func FindFirstBlockerOnPath() string {
	const gridLenX = 71
	const gridLenY = 71

	grid := make([][]string, gridLenY)
	for i := 0; i < gridLenY; i++ {
		grid[i] = make([]string, gridLenX)
		for j := 0; j < gridLenX; j++ {
			grid[i][j] = "."
		}
	}

	end := shared.Point2d{gridLenX - 1, gridLenY - 1}
	start := shared.Point2d{0, 0}

	directPaths := shared.FindShortestPaths(grid, start, end, wall)
	pathSteps := uniqueSteps(directPaths)

	for pt := range readInput() {
		grid[pt.Y][pt.X] = wall

		if _, ok := pathSteps[pt]; ok {
			directPaths = shared.FindShortestPaths(grid, start, end, wall)
			if len(directPaths) == 0 {
				println()
				return stringifyPoint(pt)
			}

			pathSteps = uniqueSteps(directPaths)
		}
	}

	panic(fmt.Errorf("no blocker found"))
}

func uniqueSteps(paths []shared.Path) map[shared.Point2d]struct{} {
	steps := make(map[shared.Point2d]struct{})
	for _, pth := range paths {
		for _, pt := range pth.Points {
			steps[pt] = struct{}{}
		}
	}
	return steps
}

func stringifyPoint(p shared.Point2d) string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func drawPaths(paths []shared.Path, grid [][]string) {
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] != wall {
				grid[row][col] = empty
			}
		}
	}
	for _, pth := range paths {
		for _, pt := range pth.Points {
			if grid[pt.Y][pt.X] == "." {
				grid[pt.Y][pt.X] = shared.Colored(shared.ColorRed, path)
			}
		}
	}
}
