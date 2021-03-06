package math_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-aoc/math"
)

func TestVector3D_Hash(t *testing.T) {
	v := math.Vector3D{15, 33, 11}

	assert.Equal(t, math.Vector3D{X: 15, Y: 33, Z: 11}, v.Hash())
}

func TestVector3D_Values(t *testing.T) {
	v := math.Vector3D{15, 33, 11}

	assert.Equal(t, []float64{15, 33, 11}, v.Values())
}

func TestVector3D_Neighbours(t *testing.T) {
	v := math.Vector3D{15, 33, 11}

	neighbors := v.Neighbours()

	assert.Len(t, neighbors, 26)
}

func TestVector3D_MinimizeFrom(t *testing.T) {
	v := math.Vector3D{15, 33, 11}
	v2 := math.Vector3D{33, 12, 13}
	expected := math.Vector3D{15, 12, 11}

	v.MinimizeFrom(v2)

	assert.Equal(t, expected, v)
}

func TestVector3D_MaximizeFrom(t *testing.T) {
	v := math.Vector3D{15, 33, 11}
	v2 := math.Vector3D{33, 12, 13}
	expected := math.Vector3D{33, 33, 13}

	v.MaximizeFrom(v2)

	assert.Equal(t, expected, v)
}

func TestVector3D_MinimizeFrom_full(t *testing.T) {
	v := math.Vector3D{33, 44, 13}
	v2 := math.Vector3D{15, 33, 11}
	expected := math.Vector3D{15, 33, 11}

	v.MinimizeFrom(v2)

	assert.Equal(t, expected, v)
}

func TestVector3D_MaximizeFrom_full(t *testing.T) {
	v := math.Vector3D{10, 12, 8}
	v2 := math.Vector3D{15, 33, 11}
	expected := math.Vector3D{15, 33, 11}

	v.MaximizeFrom(v2)

	assert.Equal(t, expected, v)
}
