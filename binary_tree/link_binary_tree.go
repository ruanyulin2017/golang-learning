package binarytree

import "fmt"

type node struct {
	data       int
	leftChild  *node
	rightChild *node
}

//	     root
//	      /\
//	     /  \
//	    /    \
//	 node1  node2
//	  /\       \
//	 /  \       \
//	/    \       \
//
// node3  node4   node5
func NewBinaryTree() BinaryTree {
	root := &node{
		data: 0,
	}

	node1 := &node{
		data: 1,
	}
	node2 := &node{
		data: 2,
	}
	node3 := &node{
		data: 3,
	}
	node4 := &node{
		data: 4,
	}
	node5 := &node{
		data: 5,
	}

	root.leftChild = node1
	root.rightChild = node2

	node1.leftChild = node3
	node1.rightChild = node4

	node2.rightChild = node5

	return root
}

// root -> node1 -> node3 -> node4 -> node2 -> node5
func (n *node) PreorderTraversal() {
	if n == nil {
		return
	}
	fmt.Printf("%v ", n.data)
	n.leftChild.PreorderTraversal()
	n.rightChild.PreorderTraversal()
}

// node3 -> node1 -> node4 -> root -> node2 -> node5
func (n *node) InorderTraversal() {
	if n == nil {
		return
	}
	n.leftChild.InorderTraversal()
	fmt.Printf("%v ", n.data)
	n.rightChild.InorderTraversal()
}

// node3 -> node4 -> node1 -> node5 -> node2 -> root
func (n *node) PostorderTraversal() {
	if n == nil {
		return
	}
	n.leftChild.PostorderTraversal()
	n.rightChild.PostorderTraversal()
	fmt.Printf("%v ", n.data)
}
