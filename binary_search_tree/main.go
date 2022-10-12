package binarysearchtree

import "fmt"

func comparef(a, b int) int8 {
	if a < b {
		return -1
	}
	if a == b {
		return 0
	}
	if a > b {
		return 1
	}
	return 0
}

func Main() {
	//                4
	//               / \
	//              /   \
	//             1     5
	//            / \     \
	//           /   \     \
	//          -1    3     8
	//               /
	//              /
	//             2
	tree := NewBinarySearchTree(4)
	for _, i := range []int{1, 3, 2, 5, 8, -1} {
		tree.Add(i, comparef)
	}
	// 4 1 -1 3 2 5 8
	tree.preorderTraversal()
	println()

	node := tree.Search(1, comparef)
	fmt.Printf("node: %v\n", node.GetData())
}
