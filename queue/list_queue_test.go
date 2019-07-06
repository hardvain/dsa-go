package queue

import (
	"gotest.tools/assert"
	"testing"
)

func TestListQueueCreation(t *testing.T) {
	listQueue := NewListQueue()
	assert.Assert(t, listQueue.Size() == 0)
}

func TestListQueueEnqueue(t *testing.T) {
	listQueue := NewListQueue()
	listQueue.Enqueue(20)
	assert.Assert(t, listQueue.Size() == 1)
	listQueue.Enqueue(30)
	listQueue.Enqueue(40)
	assert.Assert(t, listQueue.Size() == 3)

}

func TestListQueueDequeue(t *testing.T) {
	listQueue := NewListQueue()
	listQueue.Enqueue(20)
	listQueue.Enqueue(30)
	listQueue.Enqueue(40)
	item := listQueue.Dequeue()
	assert.Assert(t, item == 20)
	assert.Assert(t, listQueue.Size() == 2)

}
