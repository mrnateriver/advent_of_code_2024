package shared

import (
	"iter"
	"math"
)

type Point2d struct {
	X, Y int
}

type Direction = Point2d

var (
	DirUp    = Direction{0, -1}
	DirRight = Direction{1, 0}
	DirDown  = Direction{0, 1}
	DirLeft  = Direction{-1, 0}
)

func DistanceBetweenPoints(a, b Point2d) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2))
}

func PointAlongLineAfterB(a, b Point2d, distance float64) Point2d {
	ab := Point2d{b.X - a.X, b.Y - a.Y}
	abLength := DistanceBetweenPoints(a, b)
	abUnitX, abUnitY := float64(ab.X)/abLength, float64(ab.Y)/abLength

	return Point2d{b.X + int(math.Round(abUnitX*distance)), b.Y + int(math.Round(abUnitY*distance))}
}

func Point2dWithinBounds(p Point2d, lx, ly int) bool {
	return p.X >= 0 && p.X < lx && p.Y >= 0 && p.Y < ly
}

func GetDirection(a, b Point2d) Direction {
	dx := b.X - a.X
	dy := b.Y - a.Y
	if dx != 0 {
		if dx < 0 {
			dx = -1
		} else {
			dx = 1
		}
	}
	if dy != 0 {
		if dy < 0 {
			dy = -1
		} else {
			dy = 1
		}
	}
	return Point2d{dx, dy}
}

func MoveInDir(p Point2d, d Direction) Point2d {
	return Point2d{p.X + d.X, p.Y + d.Y}
}

func GridAt[T any](grid [][]T, p Point2d) T {
	return grid[p.Y][p.X]
}

func GridInDirection[T any](grid [][]T, p Point2d, d Direction) (res T) {
	np := MoveInDir(p, d)
	if !Point2dWithinBounds(np, len(grid[0]), len(grid)) {
		return
	}
	res = GridAt(grid, np)
	return
}

func SetGridAt[T any](grid [][]T, p Point2d, value T) {
	grid[p.Y][p.X] = value
}

func Neighbours(p Point2d, incDiags bool) iter.Seq2[Direction, Point2d] {
	return func(yield func(Direction, Point2d) bool) {
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if (dx == 0 && dy == 0) || (!incDiags && (dx == dy || dx == -dy)) {
					continue
				}

				d := Direction{dx, dy}
				if !yield(d, MoveInDir(p, d)) {
					return
				}
			}
		}
	}
}
