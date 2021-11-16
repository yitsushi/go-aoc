package puzzle_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-aoc/puzzle"
)

type testSolver struct{}

func (s testSolver) Part1() (string, error) {
	return "", nil
}

func (s testSolver) Part2() (string, error) {
	return "", nil
}

func (s testSolver) SetInput(value io.Reader) error {
	return nil
}

func TestDaySelector_valid(t *testing.T) {
	month := puzzle.NewMonth()

	month.Register(1, &testSolver{})

	day, err := month.Get(1)

	assert.NoError(t, err)
	assert.NotNil(t, day)
}

func TestDaySelector_notImplemented(t *testing.T) {
	month := puzzle.NewMonth()

	month.Register(1, &testSolver{})
	month.Register(2, nil)

	day, err := month.Get(2)

	assert.Error(t, err)
	assert.EqualError(t, err, (puzzle.NotImplementedError{}).Error())
	assert.Nil(t, day)
}

func TestDaySelector_notFound(t *testing.T) {
	month := puzzle.NewMonth()

	month.Register(1, &testSolver{})

	for _, v := range []int{-1, 0, 1000} {
		day, err := month.Get(v)

		assert.Error(t, err)
		assert.EqualError(t, err, (puzzle.UnkownDayError{Day: v}).Error())
		assert.Nil(t, day)
	}
}
