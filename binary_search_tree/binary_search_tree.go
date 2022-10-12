package binarysearchtree

type binarySearchTreeElemType interface{ any }

type BinarySearchTree[T binarySearchTreeElemType] interface {
	Add(T, func(a, b T) int8)
	Search(T, func(a, b T) int8) BinarySearchTree[T]
	preorderTraversal()
	GetData() T
}
