package queue

import (
	"gotest.tools/assert"
	"testing"
)

func TestListQueueCreation(t *testing.T) {
	arrayStack := NewListQueue()
	assert.Assert(t, arrayStack.Size() == 0)
}

func TestListQueueEnqueue(t *testing.T) {
	arrayStack := NewListQueue()
	arrayStack.Enqueue(20)
	assert.Assert(t, arrayStack.Size() == 1)
	arrayStack.Enqueue(30)
	arrayStack.Enqueue(40)
	assert.Assert(t, arrayStack.Size() == 3)

}

func TestListQueueDequeue(t *testing.T) {
	arrayStack := NewListQueue()
	arrayStack.Enqueue(20)
	arrayStack.Enqueue(30)
	arrayStack.Enqueue(40)
	item := arrayStack.Dequeue()
	assert.Assert(t, item == 20)
	assert.Assert(t, arrayStack.Size() == 2)

}
