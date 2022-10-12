package binarytree

import (
	"fmt"
	"time"
)

func Main() {
	tree := NewBinaryTree()
	fmt.Println("PreorderTraversal")
	tree.PreorderTraversal()
	fmt.Println("\n--------------------------------")
	time.Sleep(500 * time.Microsecond)

	fmt.Println("InorderTraversal")
	tree.InorderTraversal()
	fmt.Println("\n--------------------------------")
	time.Sleep(500 * time.Microsecond)

	fmt.Println("PostorderTraversal")
	tree.PostorderTraversal()
	fmt.Println("\n--------------------------------")
}
