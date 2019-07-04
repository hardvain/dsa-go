package linkedlist

import (
	"gotest.tools/assert"
	"testing"
)

func TestSinglyLinedListCreation(t *testing.T) {
	list := NewSinglyLinkedList()
	assert.Assert(t, list.head == nil)
	assert.Assert(t, list.Size() == 0)
}

func TestSinglyLinedListAdd(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Append(20)
	assert.Assert(t, list.head != nil)
	assert.Assert(t, list.Size() == 1)
	list.Append(20)
	assert.Assert(t, list.Size() == 2)
}

func TestSinglyLinedListLength(t *testing.T) {
	list := NewSinglyLinkedList()
	assert.Assert(t, list.Size() == 0)
	list.Append(20)
	list.Append(30)
	assert.Assert(t, list.Size() == 2)
}

func TestSinglyLinedListContains(t *testing.T) {
	list := NewSinglyLinkedList()
	assert.Assert(t, list.Size() == 0)
	list.Append(20)
	list.Append(30)
	list.Append(40)
	list.Append(50)
	assert.Assert(t, list.Contains(20))
	assert.Assert(t, list.Contains(50))
	assert.Assert(t, list.Contains(40))
	assert.Assert(t, list.Contains(90) == false)
}

func TestSinglyLinedListInsertAt(t *testing.T) {
	list := NewSinglyLinkedList()
	assert.Assert(t, list.Size() == 0)
	list.Append(20)
	list.Append(30)
	list.Append(40)
	list.Append(50)
	list.InsertAt(10, 0)
	list.InsertAt(12, 3)
	list.InsertAt(42, list.Size())
	assert.Assert(t, list.Contains(10))
	assert.Assert(t, list.Contains(12))
	assert.Assert(t, list.Contains(42))
}
