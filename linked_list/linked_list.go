package linkedlist

type linkElemType interface{ comparable }

type link[T linkElemType] interface {
	Add(i T)
	Remove(i T)
	Get(i T) (T, bool)
	Traverse()
	GetNodeNum() int
}
