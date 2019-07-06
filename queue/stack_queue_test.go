package queue

import (
	"gotest.tools/assert"
	"testing"
)

func TestStackQueueCreation(t *testing.T) {
	arrayStack := NewStackQueue()
	assert.Assert(t, arrayStack.Size() == 0)
}

func TestStackQueueEnqueue(t *testing.T) {
	arrayStack := NewStackQueue()
	arrayStack.Enqueue(20)
	assert.Assert(t, arrayStack.Size() == 1)
	arrayStack.Enqueue(30)
	arrayStack.Enqueue(40)
	assert.Assert(t, arrayStack.Size() == 3)

}

func TestStackQueueDequeue(t *testing.T) {
	arrayStack := NewStackQueue()
	arrayStack.Enqueue(20)
	arrayStack.Enqueue(30)
	arrayStack.Enqueue(40)
	item := arrayStack.Dequeue()
	assert.Assert(t, item == 20)
	assert.Assert(t, arrayStack.Size() == 2)
}
