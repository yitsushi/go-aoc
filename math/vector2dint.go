package math

import "math"

// Rotate the vector.
func (v *Vector2DInt) Rotate(angle float64) {
	angle *= rad

	cos, sin := math.Round(math.Cos(angle)), math.Round(math.Sin(angle))

	x := (float64(v.X) * cos) - (float64(v.Y) * sin)
	y := (float64(v.X) * sin) + (float64(v.Y) * cos)

	v.X, v.Y = int(x), int(y)
}

// Hash for cache.
func (v Vector2DInt) Hash() interface{} {
	return v
}

// Manhattan distance.
func (v Vector2DInt) Manhattan() int {
	return int(math.Abs(float64(v.X)) + math.Abs(float64(v.Y)))
}

// Neighbours if it's a coordinate.
func (v Vector2DInt) Neighbours() []Vector2DInt {
	vectors := []Vector2DInt{}
	checkRange := []int{-1, 0, 1}

	//nolint: varnamelen // These are common names for coordinates.
	for _, x := range checkRange {
		for _, y := range checkRange {
			if x == 0 && y == 0 {
				continue
			}

			vectors = append(
				vectors,
				Vector2DInt{
					X: v.X + x,
					Y: v.Y + y,
				},
			)
		}
	}

	return vectors
}

// Values for the vector as plain int slice.
func (v Vector2DInt) Values() []int {
	return []int{v.X, v.Y}
}

// Add two vectors together.
func (v Vector2DInt) Add(v2 Vector2DInt) Vector2DInt {
	return Vector2DInt{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}
