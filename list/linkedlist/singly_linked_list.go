package linkedlist

import "fmt"

type SinglyLinkedListNode struct {
	value int
	next  *SinglyLinkedListNode
}
type SinglyLinkedList struct {
	head   *SinglyLinkedListNode
	length int
}

func NewSinglyLinkedList() SinglyLinkedList {
	return SinglyLinkedList{}
}

func (list *SinglyLinkedList) Append(item int) {
	nodeToAdd := SinglyLinkedListNode{value: item}
	if list.head == nil {
		list.head = &nodeToAdd
	} else {
		var currentNode = list.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = &nodeToAdd
	}
	list.length++
}

func (list *SinglyLinkedList) Size() int {
	return list.length
}

func (list *SinglyLinkedList) Contains(item int) bool {
	if list.Size() == 0 {
		return false
	} else {
		var found bool
		var currentNode = list.head
		for currentNode != nil {
			if currentNode.value == item {
				found = true
				break
			}
			currentNode = currentNode.next
		}
		return found
	}
}

func (list *SinglyLinkedList) InsertAt(item, index int) {
	if list.Size() == 0 && index != 0 {
		panic("Cannot insert in an empty list at index other than 0")
	} else if index > list.Size() {
		panic("Cannot insert beyond the list")
	}
	if index == 0 {
		nodeToBeInserted := SinglyLinkedListNode{value: item}
		currentHead := list.head
		list.head = &nodeToBeInserted
		nodeToBeInserted.next = currentHead
	} else {
		currentNode := list.head
		var counter = 0
		for currentNode != nil && counter < index-1 {
			counter++
			currentNode = currentNode.next
		}
		currentNext := currentNode.next
		nodeToBeInserted := SinglyLinkedListNode{value: item}
		currentNode.next = &nodeToBeInserted
		nodeToBeInserted.next = currentNext
	}
	list.length++
}

func (list *SinglyLinkedList) RemoveAt(index int) (int, error) {
	size := list.Size()
	if size == 0 {
		panic("Cannot remove from an empty list")
	} else if index > size-1 {
		panic("Cannot remove beyond the list")
	}
	if index == 0 {
		list.head = list.head.next
		list.length--
		return list.head.value, nil
	} else {
		currentNode := list.head
		var counter = 1
		for currentNode != nil && counter < index {
			counter++
			currentNode = currentNode.next
		}
		list.length--
		value := currentNode.next.value
		currentNode.next = currentNode.next.next
		return value, nil
	}
}

func (list *SinglyLinkedList) ValueAt(index int) (int, error) {
	size := list.Size()
	if size == 0 {
		panic("Cannot remove from an empty list")
	} else if index > size-1 {
		panic("Cannot remove beyond the list")
	}
	if index == 0 {
		return list.head.value, nil
	} else {
		currentNode := list.head
		var counter = 1
		for currentNode != nil && counter < index {
			counter++
			currentNode = currentNode.next
		}
		return currentNode.next.value, nil
	}
}

func (list *SinglyLinkedList) Print() {
	var currentNode = list.head
	for currentNode != nil {
		fmt.Printf("%d, ", currentNode.value)
		currentNode = currentNode.next
	}
}
