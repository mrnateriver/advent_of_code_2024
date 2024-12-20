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

func Point2dWithinBounds(p Point2d, lx, ly int) bool {
	return CoordsWithinBounds(p.X, p.Y, lx, ly)
}

func CoordsWithinBounds(x, y, lx, ly int) bool {
	return x >= 0 && x < lx && y >= 0 && y < ly
}
