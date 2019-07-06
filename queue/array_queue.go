package queue

type ArrayQueue struct {
	top, rear int
	items     [1024]int
}

func NewArrayQueue() ArrayQueue {
	s := ArrayQueue{}
	s.top = -1
	s.rear = 0
	return s
}

func (s *ArrayQueue) Enqueue(value int) {
	s.top++
	s.items[s.top] = value
}

func (s *ArrayQueue) Dequeue() int {
	if s.top == -1 {
		panic("Empty stack")
	}
	item := s.items[s.rear]
	s.items[s.rear] = 0
	s.rear++
	return item
}

func (s *ArrayQueue) Size() int {
	return s.top - s.rear + 1
}
