package shared

import "fmt"

func FindShortestPaths[T comparable](grid [][]T, start, end Point2d, wall T) []Path {
	lenX, lenY := len(grid[0]), len(grid)

	finalLength := -1
	seen := make(map[Point2d]int)
	queue := MakePriorityQueue[Path]()
	queue.PushEntry(Path{Points: []Point2d{start}}, 0)

	res := make([]Path, 0)

	for queue.Len() > 0 {
		path := queue.PollEntry()
		pathLen := path.Len()
		pathEnd := path.End()

		if finalLength >= 0 && pathLen > finalLength {
			return res
		} else if pathEnd == end {
			finalLength = pathLen
			res = append(res, path)
		} else if seenLen, ok := seen[pathEnd]; !ok || seenLen > pathLen {
			seen[pathEnd] = pathLen

			for _, next := range Neighbours(pathEnd, false) {
				newLength := pathLen + 1
				if Point2dWithinBounds(next, lenX, lenY) && GridAt(grid, next) != wall {
					queue.PushEntry(path.Append(next), newLength)
				}
			}
		}
	}

	return res
}

func FindShortestPathLength[T comparable](grid [][]T, start, end Point2d, wall T) int {
	paths := FindShortestPaths(grid, start, end, wall)
	if len(paths) == 0 {
		return -1
	}
	return paths[0].Len()
}

type Path struct {
	Points []Point2d
}

func (p Path) End() Point2d {
	if len(p.Points) == 0 {
		panic(fmt.Errorf("attempt to get end of an empty path"))
	}
	return p.Points[len(p.Points)-1]
}

func (p Path) Len() int {
	return len(p.Points) - 1
}

func (p Path) MoveInDir(dir Direction) Path {
	return p.Append(MoveInDir(p.End(), dir))
}

func (p Path) Append(newStep Point2d) Path {
	l := len(p.Points)
	newSteps := make([]Point2d, l, l+1)
	copy(newSteps, p.Points)

	return Path{
		Points: append(newSteps, newStep),
	}
}
