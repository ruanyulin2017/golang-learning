package binarysearch

import (
	"fmt"
)

type data struct {
	id    int
	name  string
	phone string
}

func compareNamef(a, b data) int8 {
	if a.id < b.id {
		return -1
	}
	if a.id == b.id {
		return 0
	}
	if a.id > b.id {
		return 1
	}
	return 0
}

func Main() {
	s := NewBinarySearch[data](10)
	datas := [][]any{
		{3, "name1", "11111"},
		{11, "name2", "22222"},
		{43, "name3", "33333"},
		{10, "name4", "44444"},
		{5, "name5", "55555"},
		{100, "name9", "99999"},
	}

	for _, d := range datas {
		data := data{
			id:    d[0].(int),
			name:  d[1].(string),
			phone: d[2].(string),
		}
		s.Add(data, compareNamef)
	}

	fmt.Printf("s: %v\n", s)

	searchData := data{id: 100}

	if err := s.Search(&searchData, compareNamef); err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("d: %v\n", searchData)
	}

}
