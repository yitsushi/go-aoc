package slice_test

import (
	"testing"

	"github.com/yitsushi/go-aoc/slice"
)

type equalTest struct {
	a     []int64
	b     []int64
	equal bool
}

func TestEqualInt64(t *testing.T) {
	testCases := []equalTest{
		{
			a:     []int64{2, 3, 5, 6},
			b:     []int64{2, 3, 5, 6},
			equal: true,
		},
		{
			a:     []int64{2, 3, 5, 6},
			b:     []int64{3, 2, 5, 6},
			equal: false,
		},
		{
			a:     []int64{3, 2},
			b:     []int64{3, 2, 5, 6},
			equal: false,
		},
		{
			a:     []int64{3, 2, 5, 6},
			b:     []int64{3, 2},
			equal: false,
		},
	}

	for _, testCase := range testCases {
		got := slice.EqualInt64(testCase.a, testCase.b)
		if got != testCase.equal {
			t.Errorf("%v == %v should be %v; got = %v", testCase.a, testCase.b, testCase.equal, got)
		}
	}
}
