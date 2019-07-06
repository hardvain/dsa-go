package queue

import (
	"gotest.tools/assert"
	"testing"
)

func TestArrayQueueCreation(t *testing.T) {
	arrayQueue := NewArrayQueue()
	assert.Assert(t, arrayQueue.top == -1)
	assert.Assert(t, arrayQueue.rear == 0)
	assert.Assert(t, arrayQueue.Size() == 0)
}

func TestArrakQueueEnqueue(t *testing.T) {
	arrayQueue := NewArrayQueue()
	arrayQueue.Enqueue(20)
	assert.Assert(t, arrayQueue.top == 0)
	assert.Assert(t, arrayQueue.rear == 0)
	assert.Assert(t, arrayQueue.Size() == 1)
	arrayQueue.Enqueue(30)
	arrayQueue.Enqueue(40)
	assert.Assert(t, arrayQueue.top == 2)
	assert.Assert(t, arrayQueue.rear == 0)
	assert.Assert(t, arrayQueue.Size() == 3)

}

func TestArrayQueueDequeue(t *testing.T) {
	arrayQueue := NewArrayQueue()
	arrayQueue.Enqueue(20)
	arrayQueue.Enqueue(30)
	arrayQueue.Enqueue(40)
	item := arrayQueue.Dequeue()
	assert.Assert(t, item == 20)
	assert.Assert(t, arrayQueue.top == 2)
	assert.Assert(t, arrayQueue.rear == 1)
	assert.Assert(t, arrayQueue.Size() == 2)

}
