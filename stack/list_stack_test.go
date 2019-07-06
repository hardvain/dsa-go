package stack

import (
	"gotest.tools/assert"
	"testing"
)

func TestListStackPush(t *testing.T) {
	listStack := NewListStack()
	listStack.Push(20)
	assert.Assert(t, listStack.Top() == 20)
}

func TestListStackPop(t *testing.T) {
	listStack := NewListStack()
	listStack.Push(20)
	listStack.Push(30)
	item := listStack.Pop()
	assert.Assert(t, item == 30)
	assert.Assert(t, listStack.Top() == 20)
}

func TestListStackTop(t *testing.T) {
	listStack := NewListStack()
	listStack.Push(20)
	listStack.Push(30)
	item := listStack.Top()
	assert.Assert(t, item == 30)
}
