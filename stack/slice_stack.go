package stack

type SliceStack struct {
	top, length int
	items []int
}


func NewSliceStack() SliceStack{
	stack := SliceStack{}
	stack.top = -1
	stack.length = 0
	stack.items = []int{}
	return stack
}


func (s *SliceStack) Push(value int) {
	s.top++
	s.length++
	s.items = append(s.items, value)
}

func (s *SliceStack) Pop() int {
	if s.top == -1 {
		panic("Empty stack")
	}
	item := s.items[s.top]
	s.items = s.items[:len(s.items) - 1]
	s.top--
	s.length--
	return item
}

func (s *SliceStack) Top() int {
	if s.top == -1 {
		panic("Empty stack")
	}
	item := s.items[s.top]
	return item
}

