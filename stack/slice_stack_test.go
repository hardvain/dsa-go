package stack

import (
	"gotest.tools/assert"
	"testing"
)

func TestSliceStackCreation(t *testing.T) {
	sliceStack := NewSliceStack()
	assert.Assert(t, sliceStack.top == -1)
}

func TestSliceStackPush(t *testing.T) {
	sliceStack := NewSliceStack()
	sliceStack.Push(20)
	assert.Assert(t, sliceStack.top == 0)
	assert.Assert(t, sliceStack.Top() == 20)
}

func TestSliceStackPop(t *testing.T) {
	sliceStack := NewSliceStack()
	sliceStack.Push(20)
	sliceStack.Push(30)
	item := sliceStack.Pop()
	assert.Assert(t, item == 30)
	assert.Assert(t, sliceStack.top == 0)
	assert.Assert(t, sliceStack.Top() == 20)
}

func TestSliceStackTop(t *testing.T) {
	sliceStack := NewSliceStack()
	sliceStack.Push(20)
	sliceStack.Push(30)
	item := sliceStack.Top()
	assert.Assert(t, item == 30)
	assert.Assert(t, sliceStack.top == 1)
}
