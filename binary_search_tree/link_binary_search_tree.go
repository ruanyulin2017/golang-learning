package binarysearchtree

import (
	"fmt"
	"sync"
)

type node[T binarySearchTreeElemType] struct {
	data       T
	leftChild  *node[T]
	rightChild *node[T]
	lock       sync.RWMutex
}

func NewBinarySearchTree[T binarySearchTreeElemType](data T) BinarySearchTree[T] {
	return &node[T]{
		data: data,
	}
}

func (n *node[T]) Add(data T, f func(a, b T) int8) {
	n.lock.Lock()
	defer n.lock.Unlock()

	res := f(data, n.data)
	switch res {
	case -1:
		// data < n.data
		if n.leftChild == nil {
			n.leftChild = &node[T]{
				data: data,
			}
			return
		}
		n.leftChild.Add(data, f)
	case 0:
		// data == n.data
		return
	case 1:
		// data > n.data
		if n.rightChild == nil {
			n.rightChild = &node[T]{
				data: data,
			}
			return
		}
		n.rightChild.Add(data, f)
	}
}

func (n *node[T]) Search(data T, f func(a, b T) int8) BinarySearchTree[T] {
	if n == nil {
		return nil
	}

	n.lock.RLock()
	defer n.lock.RUnlock()

	res := f(data, n.data)
	switch res {
	case -1:
		return n.leftChild.Search(data, f)
	case 0:
		return n
	case 1:
		return n.rightChild.Search(data, f)
	}
	return nil
}

func (n *node[T]) preorderTraversal() {
	if n == nil {
		return
	}
	fmt.Printf("%v ", n.data)
	n.leftChild.preorderTraversal()
	n.rightChild.preorderTraversal()
}

func (n *node[T]) GetData() T {
	if n == nil {
		var t T
		return t
	}
	return n.data
}
