package runtime

type Stack []any

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Peek() any {
	index := len(*s) - 1

	if index < 0 {
		return nil
	}

	return (*s)[index]
}

func (s *Stack) Pull() (value any, ok bool) {
	index := len(*s) - 1

	if index < 0 {
		return nil, false
	}

	value, ok = (*s)[index], true

	(*s)[index] = nil
	*s = (*s)[:index]

	return
}

func (s *Stack) Push(value any) {
	*s = append(*s, value)
}

func (s *Stack) Flip() *Stack {
	stack := NewStack()

	for {
		if value, ok := s.Pull(); ok {
			stack.Push(value)
		} else {
			return stack
		}
	}
}
