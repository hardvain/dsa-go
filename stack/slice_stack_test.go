package stack

import (
	"gotest.tools/assert"
	"testing"
)

func TestSliceStackCreation(t *testing.T) {
	arrayStack := NewSliceStack()
	assert.Assert(t, arrayStack.top == -1)
}

func TestSliceStackPush(t *testing.T) {
	arrayStack := NewSliceStack()
	arrayStack.Push(20)
	assert.Assert(t, arrayStack.top == 0)
	assert.Assert(t, arrayStack.Top() == 20)
}

func TestSliceStackPop(t *testing.T) {
	arrayStack := NewSliceStack()
	arrayStack.Push(20)
	arrayStack.Push(30)
	item := arrayStack.Pop()
	assert.Assert(t, item == 30)
	assert.Assert(t, arrayStack.top == 0)
	assert.Assert(t, arrayStack.Top() == 20)
}

func TestSliceStackTop(t *testing.T) {
	arrayStack := NewSliceStack()
	arrayStack.Push(20)
	arrayStack.Push(30)
	item := arrayStack.Top()
	assert.Assert(t, item == 30)
	assert.Assert(t, arrayStack.top == 1)
}
