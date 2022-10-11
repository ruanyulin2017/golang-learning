package linkedlist

import "fmt"

type node[T linkElemType] struct {
	data    T
	nodeNum int
	pre     *node[T]
	next    *node[T]
}

func NewDoublyLink[T linkElemType]() link[T] {
	return &node[T]{}
}

func (n *node[T]) Add(i T) {
	n1 := &node[T]{
		data: i,
	}
	p := n
	for p.next != nil {
		p = p.next
	}
	p.next = n1
	n1.pre = p
	n.nodeNum++
}

func (n *node[T]) Remove(i T) {
	no := n.get(i)

	if no == nil {
		return
	}

	no.pre.next = no.next
	if no.next != nil {
		no.next.pre = no.pre
	}

	n.nodeNum--
}

func (n node[T]) Get(i T) (T, bool) {
	no := n.get(i)
	if no == nil {
		var t T
		return t, false
	}
	return no.data, true
}

func (n node[T]) Traverse() {
	fmt.Printf("node num: %v\n", n.nodeNum)
	p := n.next
	for p != nil {
		fmt.Printf("data: %v\n", p.data)
		p = p.next
	}
}

func (n node[T]) GetNodeNum() int {
	return n.nodeNum
}

func (n *node[T]) get(i T) *node[T] {
	p := n
	for p.data != i {
		p = p.next
		if p == nil {
			return nil
		}
	}
	return p
}
