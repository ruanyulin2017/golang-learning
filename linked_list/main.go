package linkedlist

import "fmt"

func Main() {

	type lT = string
	datas := []lT{"aa", "bb", "cc"}
	link := NewDoublyLink[lT]()
	for _, i := range datas {
		link.Add(i)
	}
	// 遍历
	fmt.Println("遍历链表")
	link.Traverse()
	// 获取节点
	for _, i := range []lT{"aa", "bb", "dd"} {
		data, ok := link.Get(i)
		if ok {
			fmt.Printf("Get node: data=%v\n", data)
		} else {
			fmt.Printf("node %v not found\n", i)
		}
	}
	// 移除节点
	fmt.Println("移除节点")
	link.Remove("aa")
	link.Traverse()
	// 移除不存在的节点
	fmt.Println("移除不存在的节点")
	link.Remove("dd")
	link.Traverse()
	// 移除尾节点
	fmt.Println("移除尾节点")
	link.Remove("cc")
	link.Traverse()
}
