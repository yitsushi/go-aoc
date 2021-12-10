package math_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-aoc/math"
)

func TestVector2DInt_Rotate(t *testing.T) {
	tests := []struct {
		name     string
		initial  math.Vector2DInt
		degree   float64
		expected math.Vector2DInt
	}{
		{
			name:     "Rotate 90",
			initial:  math.Vector2DInt{X: 1, Y: 0},
			degree:   90,
			expected: math.Vector2DInt{X: 0, Y: 1},
		},
		{
			name:     "Rotate -90",
			initial:  math.Vector2DInt{X: 1, Y: 0},
			degree:   -90,
			expected: math.Vector2DInt{X: 0, Y: -1},
		},
		{
			name:     "Rotate 180",
			initial:  math.Vector2DInt{X: 1, Y: 0},
			degree:   180,
			expected: math.Vector2DInt{X: -1, Y: 0},
		},
		{
			name:     "Rotate -180",
			initial:  math.Vector2DInt{X: 1, Y: 0},
			degree:   -180,
			expected: math.Vector2DInt{X: -1, Y: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &math.Vector2DInt{
				X: tt.initial.X,
				Y: tt.initial.Y,
			}

			v.Rotate(tt.degree)

			assert.Equal(t, tt.expected.X, v.X)
			assert.Equal(t, tt.expected.Y, v.Y)
		})
	}
}

func TestVector2DInt_Manhattan(t *testing.T) {
	v := math.Vector2DInt{15, 33}

	assert.Equal(t, int(48), v.Manhattan())
}

func TestVector2DInt_Hash(t *testing.T) {
	v := math.Vector2DInt{15, 33}

	assert.Equal(t, math.Vector2DInt{X: 15, Y: 33}, v.Hash())
}

func TestVector2DInt_Neighbours(t *testing.T) {
	v := math.Vector2DInt{15, 33}

	neighbors := v.Neighbours()

	assert.Len(t, neighbors, 8)
}

func TestVector2DInt_Values(t *testing.T) {
	v := math.Vector2DInt{15, 33}

	assert.Equal(t, []int{15, 33}, v.Values())
}

func TestVector2DInt_Add(t *testing.T) {
	v := math.Vector2DInt{15, 33}

	assert.Equal(t, math.Vector2DInt{26, 36}, v.Add(math.Vector2DInt{11, 3}))
}
