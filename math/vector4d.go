package math

// Vector4D is a 3D vector.
type Vector4D struct {
	X float64
	Y float64
	Z float64
	W float64
}

// Hash for cache.
func (v Vector4D) Hash() interface{} {
	return v
}

// Neighbours if it's a coordinate.
func (v Vector4D) Neighbours() []Vector {
	vectors := []Vector{}
	checkRange := []float64{-1, 0, 1}

	//nolint: varnamelen // These are common names for coordinates.
	for _, x := range checkRange {
		for _, y := range checkRange {
			for _, z := range checkRange {
				for _, w := range checkRange {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}

					vectors = append(
						vectors,
						Vector4D{
							X: v.X + x,
							Y: v.Y + y,
							Z: v.Z + z,
							W: v.W + w,
						},
					)
				}
			}
		}
	}

	return vectors
}

// MinimizeFrom set x, y and z values to the minimum value from the two vectors.
func (v *Vector4D) MinimizeFrom(vector Vector4D) {
	if vector.X < v.X {
		v.X = vector.X
	}

	if vector.Y < v.Y {
		v.Y = vector.Y
	}

	if vector.Z < v.Z {
		v.Z = vector.Z
	}

	if vector.W < v.W {
		v.W = vector.W
	}
}

// MaximizeFrom set x, y and z values to the maximum value from the two vectors.
func (v *Vector4D) MaximizeFrom(vector Vector4D) {
	if vector.X > v.X {
		v.X = vector.X
	}

	if vector.Y > v.Y {
		v.Y = vector.Y
	}

	if vector.Z > v.Z {
		v.Z = vector.Z
	}

	if vector.W > v.W {
		v.W = vector.W
	}
}

// Values for the vector as plain float64 slice.
func (v Vector4D) Values() []float64 {
	return []float64{v.X, v.Y, v.Z, v.W}
}
