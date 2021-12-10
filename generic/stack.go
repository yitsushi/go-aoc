package generic

// Stack is a basic stack.
type Stack struct {
	items []interface{}
}

// NewStack creates a new empty stack.
func NewStack() Stack {
	return Stack{items: []interface{}{}}
}

// Size of the stack.
func (s *Stack) Size() int {
	return len(s.items)
}

// IsEmpty tells if the stack is empty or not.
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

// Push an item onto the stack.
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop the top item.
func (s *Stack) Pop() interface{} {
	value := s.items[s.Size()-1]
	s.items = s.items[:s.Size()-1]

	return value
}

// Peek the top item.
func (s *Stack) Peek() interface{} {
	return s.items[s.Size()-1]
}
