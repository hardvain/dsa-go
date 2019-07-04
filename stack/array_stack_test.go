package stack

import (
	"gotest.tools/assert"
	"testing"
)

func TestArrayStackCreation(t *testing.T) {
	arrayStack := NewArrayStack()
	assert.Assert(t, arrayStack.top == -1)
}

func TestArrayStackPush(t *testing.T) {
	arrayStack := NewArrayStack()
	arrayStack.Push(20)
	assert.Assert(t, arrayStack.top == 0)
	assert.Assert(t, arrayStack.Top() == 20)
}

func TestArrayStackPop(t *testing.T) {
	arrayStack := NewArrayStack()
	arrayStack.Push(20)
	arrayStack.Push(30)
	item := arrayStack.Pop()
	assert.Assert(t, item == 30)
	assert.Assert(t, arrayStack.top == 0)
	assert.Assert(t, arrayStack.Top() == 20)
}

func TestArrayStackTop(t *testing.T) {
	arrayStack := NewArrayStack()
	arrayStack.Push(20)
	arrayStack.Push(30)
	item := arrayStack.Top()
	assert.Assert(t, item == 30)
	assert.Assert(t, arrayStack.top == 1)
}
