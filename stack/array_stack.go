package stack

type ArrayStack struct {
	top   int
	items [1024]int
}

func NewArrayStack() ArrayStack {
	s := ArrayStack{}
	s.top = -1
	return s
}

func (s *ArrayStack) Push(value int) {
	s.top++
	s.items[s.top] = value
}

func (s *ArrayStack) Size() int {
	return s.top
}

func (s *ArrayStack) Pop() int {
	if s.top == -1 {
		panic("Empty stack")
	}
	item := s.items[s.top]
	s.items[s.top] = 0
	s.top--
	return item
}

func (s *ArrayStack) Top() int {
	if s.top == -1 {
		panic("Empty stack")
	}
	item := s.items[s.top]
	return item
}
