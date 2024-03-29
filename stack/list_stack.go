package stack

import (
	"dsa/list"
)

type ListStack struct {
	list list.List
}

func NewListStack() ListStack {
	var l list.List
	singlyLinkedList := list.NewSinglyLinkedList()
	l = &singlyLinkedList
	s := ListStack{list: l}
	return s
}

func (s *ListStack) Push(value int) {
	s.list.Append(value)
}

func (s *ListStack) Pop() int {
	i, _ := s.list.RemoveAt(s.list.Size() - 1)
	return i
}

func (s *ListStack) Size() int {
	return s.list.Size()
}

func (s *ListStack) Top() int {
	i, _ := s.list.ValueAt(s.list.Size() - 1)
	return i
}
