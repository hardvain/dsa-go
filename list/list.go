package list

type List interface {
	Append(item int)
	Size() int
	Contains(item int) bool
	InsertAt(item, index int)
	RemoveAt(index int) (int, error)
	ValueAt(index int) (int, error)
}
