package linkedlist

import "fmt"

type node struct {
	name    string
	age     uint8
	nodeNum int
	pre     *node
	next    *node
}

func NewDoublyLink() link {
	return &node{}
}

func (n *node) Add(name string, age uint8) {
	n1 := &node{
		name: name,
		age:  age,
	}
	p := n
	for p.next != nil {
		p = p.next
	}
	p.next = n1
	n1.pre = p
	n.nodeNum++
}

func (n *node) Remove(name string) {
	no := n.get(name)

	if no == nil {
		return
	}

	no.pre.next = no.next
	if no.next != nil {
		no.next.pre = no.pre
	}

	n.nodeNum--
}

func (n node) Get(name string) (string, uint8, bool) {
	no := n.get(name)
	if no == nil {
		return "", 0, false
	}
	return no.name, no.age, true
}

func (n node) Traverse() {
	fmt.Printf("node num: %d\n", n.nodeNum)
	p := n.next
	for p != nil {
		fmt.Printf("name: %s, age: %d\n", p.name, p.age)
		p = p.next
	}
}

func (n node) GetNodeNum() int {
	return n.nodeNum
}

func (n *node) get(name string) *node {
	p := n
	for p.name != name {
		p = p.next
		if p == nil {
			break
		}
	}
	return p
}
