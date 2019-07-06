package stack

type Stack interface {
	Push(item int)
	Size() int
	Top() int
	Pop() int
}
