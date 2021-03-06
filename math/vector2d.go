package math

import (
	"math"
)

// Vector2D is a 2D vector.
type Vector2D struct {
	X float64
	Y float64
}

// Rotate the vector.
func (v *Vector2D) Rotate(angle float64) {
	angle *= rad

	cos, sin := math.Round(math.Cos(angle)), math.Round(math.Sin(angle))

	x := (v.X * cos) - (v.Y * sin)
	y := (v.X * sin) + (v.Y * cos)

	v.X, v.Y = x, y
}

// Hash for cache.
func (v Vector2D) Hash() interface{} {
	return v
}

// Manhattan distance.
func (v Vector2D) Manhattan() float64 {
	return math.Abs(v.X) + math.Abs(v.Y)
}

// Neighbours if it's a coordinate.
func (v Vector2D) Neighbours() []Vector {
	vectors := []Vector{}
	checkRange := []float64{-1, 0, 1}

	//nolint: varnamelen // These are common names for coordinates.
	for _, x := range checkRange {
		for _, y := range checkRange {
			if x == 0 && y == 0 {
				continue
			}

			vectors = append(
				vectors,
				Vector2D{
					X: v.X + x,
					Y: v.Y + y,
				},
			)
		}
	}

	return vectors
}

// Values for the vector as plain float64 slice.
func (v Vector2D) Values() []float64 {
	return []float64{v.X, v.Y}
}

// Vector2DInt is a 2D vector.
type Vector2DInt struct {
	X int
	Y int
}
