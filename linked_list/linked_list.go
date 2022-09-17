package linkedlist

type link interface {
	Add(name string, age uint8)
	Remove(name string)
	Get(name string) (string, uint8, bool)
	Traverse()
	GetNodeNum() int
}
