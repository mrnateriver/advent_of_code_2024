package shared

import "math"

type Point2d struct {
	X, Y int
}

func DistanceBetweenPoints(a, b Point2d) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2))
}

func PointAlongLineAfterB(a, b Point2d, distance float64) Point2d {
	ab := Point2d{b.X - a.X, b.Y - a.Y}
	abLength := DistanceBetweenPoints(a, b)
	abUnitX, abUnitY := float64(ab.X)/abLength, float64(ab.Y)/abLength

	return Point2d{b.X + int(math.Round(abUnitX*distance)), b.Y + int(math.Round(abUnitY*distance))}
}

func PointWithinBounds(p Point2d, lx, ly int) bool {
	return p.X >= 0 && p.X < lx && p.Y >= 0 && p.Y < ly
}
