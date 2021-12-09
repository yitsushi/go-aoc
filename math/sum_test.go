package math_test

import (
	"testing"

	"github.com/yitsushi/go-aoc/math"
)

func TestSumInt64(t *testing.T) {
	tests := []struct {
		name string
		list []int64
		want int64
	}{
		{name: "empty", list: []int64{}, want: 0},
		{name: "values", list: []int64{1, 2, 3}, want: 6},
		{name: "values", list: []int64{5, 4, 92, 13, 42, 52}, want: 208},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := math.SumInt64(tt.list...); got != tt.want {
				t.Errorf("SumInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
