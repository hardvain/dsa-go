package queue

import (
	"dsa/list"
)

type ListQueue struct {
	list list.List
}

func NewListQueue() ListQueue {
	linkedList := list.NewSinglyLinkedList()
	s := ListQueue{list: &linkedList}
	return s
}

func (s *ListQueue) Enqueue(value int) {
	s.list.Append(value)
}

func (s *ListQueue) Dequeue() int {
	i, _ := s.list.RemoveAt(0)
	return i
}

func (s *ListQueue) Size() int {
	return s.list.Size()
}
