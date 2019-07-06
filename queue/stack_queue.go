package queue

import "dsa/stack"

type StackQueue struct {
	main, aux stack.Stack
}

func NewStackQueue() StackQueue {
	s := StackQueue{}
	mainStack := stack.NewSliceStack()
	auxStack := stack.NewSliceStack()
	s.main = &mainStack
	s.aux = &auxStack
	return s
}

func (s *StackQueue) Enqueue(value int) {
	s.offloadToAux()
	s.main.Push(value)
	s.reloadFromAux()
}
func (s *StackQueue) offloadToAux() {
	for s.main.Size() != 0 {
		s.aux.Push(s.main.Pop())
	}
}
func (s *StackQueue) reloadFromAux() {
	for s.aux.Size() != 0 {
		s.main.Push(s.aux.Pop())
	}
}
func (s *StackQueue) Dequeue() int {
	return s.main.Pop()
}

func (s *StackQueue) Size() int {
	return s.main.Size()
}
