package queue

import (
	"gotest.tools/assert"
	"testing"
)

func TestStackQueueCreation(t *testing.T) {
	stackQueue := NewStackQueue()
	assert.Assert(t, stackQueue.Size() == 0)
}

func TestStackQueueEnqueue(t *testing.T) {
	stackQueue := NewStackQueue()
	stackQueue.Enqueue(20)
	assert.Assert(t, stackQueue.Size() == 1)
	stackQueue.Enqueue(30)
	stackQueue.Enqueue(40)
	assert.Assert(t, stackQueue.Size() == 3)

}

func TestStackQueueDequeue(t *testing.T) {
	stackQueue := NewStackQueue()
	stackQueue.Enqueue(20)
	stackQueue.Enqueue(30)
	stackQueue.Enqueue(40)
	item := stackQueue.Dequeue()
	assert.Assert(t, item == 20)
	assert.Assert(t, stackQueue.Size() == 2)
}
