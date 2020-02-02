package container

// Stack defines a stack data structure providing LIFO semantics, where each
// item in the queue is an interface.
type Stack struct {
	items []interface{}
}

func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

// Size returns the number of items in the stack.
func (s *Stack) Size() int {
	return len(s.items)
}

// Push pushes an item onto the stack.
func (s *Stack) Push(i interface{}) {
	s.items = append(s.items, i)
}

// Pop removes an item from the stack. If the stack is empty, nil is returned.
func (s *Stack) Pop() interface{} {
	if s.Size() == 0 {
		return nil
	}

	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return v
}
