package main

import (
	"dsa/stack"
	"fmt"
)

func main() {
	arrayStack := stack.NewSliceStack()
	arrayStack.Push(20)
	arrayStack.Push(30)
	arrayStack.Push(40)
	arrayStack.Push(50)
	fmt.Println(arrayStack)
	fmt.Println(arrayStack.Top())
	fmt.Println(arrayStack)
	fmt.Println(arrayStack.Pop())
	fmt.Println(arrayStack.Pop())
	fmt.Println(arrayStack)
}

