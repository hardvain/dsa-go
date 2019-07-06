package queue

import (
	"gotest.tools/assert"
	"testing"
)

func TestArrayQueueCreation(t *testing.T) {
	arrayStack := NewArrayQueue()
	assert.Assert(t, arrayStack.top == -1)
	assert.Assert(t, arrayStack.rear == 0)
	assert.Assert(t, arrayStack.Size() == 0)
}

func TestArrayStackEnqueue(t *testing.T) {
	arrayStack := NewArrayQueue()
	arrayStack.Enqueue(20)
	assert.Assert(t, arrayStack.top == 0)
	assert.Assert(t, arrayStack.rear == 0)
	assert.Assert(t, arrayStack.Size() == 1)
	arrayStack.Enqueue(30)
	arrayStack.Enqueue(40)
	assert.Assert(t, arrayStack.top == 2)
	assert.Assert(t, arrayStack.rear == 0)
	assert.Assert(t, arrayStack.Size() == 3)

}

func TestArrayStackDequeue(t *testing.T) {
	arrayStack := NewArrayQueue()
	arrayStack.Enqueue(20)
	arrayStack.Enqueue(30)
	arrayStack.Enqueue(40)
	item := arrayStack.Dequeue()
	assert.Assert(t, item == 20)
	assert.Assert(t, arrayStack.top == 2)
	assert.Assert(t, arrayStack.rear == 1)
	assert.Assert(t, arrayStack.Size() == 2)

}
