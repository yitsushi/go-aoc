package math

// Vector is a vector interface.
type Vector interface {
	Hash() interface{}
	Neighbours() []Vector
	Values() []float64
}
