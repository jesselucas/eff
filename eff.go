package eff

import (
	"math"
	"math/rand"
)

// Point container for 2d points
type Point struct {
	X int
	Y int
}

// Color container for argb colors
type Color struct {
	R int
	G int
	B int
	A int
}

// RandomColor genereate a random color struct.  The opacity is also random
func RandomColor() Color {
	return Color{
		R: rand.Intn(0xFF),
		G: rand.Intn(0xFF),
		B: rand.Intn(0xFF),
		A: rand.Intn(0xFF),
	}
}

// Rect container for rectangle
type Rect struct {
	X int
	Y int
	W int
	H int
}

// ColorRect container for rectange and color
type ColorRect struct {
	Rect
	Color
}

// ColorPoint container for point and color
type ColorPoint struct {
	Point
	Color
}

// Font describes a ttf font
type Font struct {
	Path string
}

// Equals test to see if two rectangles occupy the same location exactly
func (r *Rect) Equals(otherRect Rect) bool {
	return (r.X == otherRect.X &&
		r.Y == otherRect.Y &&
		r.W == otherRect.W &&
		r.H == otherRect.H)
}

// Intersects check to see if a rectangle is inside of this rectangle
func (r *Rect) Intersects(otherRect Rect) bool {
	return (int(math.Abs(float64(r.X-otherRect.X)))*2 < (r.W + otherRect.W)) &&
		(int(math.Abs(float64(r.Y-otherRect.Y)))*2 < (r.H + otherRect.H))
}