package generic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-aoc/generic"
)

func TestStack(t *testing.T) {
	stack := generic.NewStack()

	assert.Equal(t, 0, stack.Size())
	assert.True(t, stack.IsEmpty())

	stack.Push("value 1")
	assert.Equal(t, 1, stack.Size())
	assert.False(t, stack.IsEmpty())
	assert.Equal(t, "value 1", stack.Peek())

	stack.Push("value 2")
	assert.Equal(t, 2, stack.Size())
	assert.False(t, stack.IsEmpty())
	assert.Equal(t, "value 2", stack.Peek())

	value := stack.Pop()
	assert.Equal(t, 1, stack.Size())
	assert.False(t, stack.IsEmpty())
	assert.Equal(t, "value 2", value)
	assert.Equal(t, "value 1", stack.Peek())
}
