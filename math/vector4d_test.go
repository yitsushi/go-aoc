package math_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-aoc/math"
)

func TestVector4D_Hash(t *testing.T) {
	v := math.Vector4D{15, 33, 11, 23}

	assert.Equal(t, math.Vector4D{X: 15, Y: 33, Z: 11, W: 23}, v.Hash())
}

func TestVector4D_Values(t *testing.T) {
	v := math.Vector4D{15, 33, 11, 23}

	assert.Equal(t, []float64{15, 33, 11, 23}, v.Values())
}

func TestVector4D_Neighbours(t *testing.T) {
	v := math.Vector4D{15, 33, 11, 23}

	neighbors := v.Neighbours()

	assert.Len(t, neighbors, 80)
}

func TestVector4D_MinimizeFrom(t *testing.T) {
	v := math.Vector4D{15, 33, 11, 23}
	v2 := math.Vector4D{33, 12, 13, 33}
	expected := math.Vector4D{15, 12, 11, 23}

	v.MinimizeFrom(v2)

	assert.Equal(t, expected, v)
}

func TestVector4D_MaximizeFrom(t *testing.T) {
	v := math.Vector4D{15, 33, 11, 23}
	v2 := math.Vector4D{33, 12, 13, 33}
	expected := math.Vector4D{33, 33, 13, 33}

	v.MaximizeFrom(v2)

	assert.Equal(t, expected, v)
}

func TestVector4D_MinimizeFrom_full(t *testing.T) {
	v := math.Vector4D{33, 44, 13, 23}
	v2 := math.Vector4D{15, 33, 11, 13}
	expected := math.Vector4D{15, 33, 11, 13}

	v.MinimizeFrom(v2)

	assert.Equal(t, expected, v)
}

func TestVector4D_MaximizeFrom_full(t *testing.T) {
	v := math.Vector4D{10, 12, 8, 9}
	v2 := math.Vector4D{15, 33, 11, 45}
	expected := math.Vector4D{15, 33, 11, 45}

	v.MaximizeFrom(v2)

	assert.Equal(t, expected, v)
}
