package linkedlist

import "fmt"

func Main() {
	datas := map[string]int{
		"name1": 1,
		"name2": 2,
		"name3": 3,
	}
	link := NewDoublyLink()
	for k, v := range datas {
		link.Add(k, uint8(v))
	}
	// 遍历
	fmt.Println("遍历链表")
	link.Traverse()
	// 获取节点
	for _, name := range [2]string{"name2", "name4"} {
		name1, age, ok := link.Get(name)
		if ok {
			fmt.Printf("Get node: name=%s, age=%d\n", name1, age)
		} else {
			fmt.Printf("node %s not found\n", name)
		}
	}
	// 移除节点
	fmt.Println("移除节点")
	link.Remove("name1")
	link.Traverse()
	// 移除不存在的节点
	fmt.Println("移除不存在的节点")
	link.Remove("name11")
	link.Traverse()
	// 移除尾节点
	fmt.Println("移除尾节点")
	link.Remove("name3")
	link.Traverse()
}
