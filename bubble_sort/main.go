package bubblesort

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
	s := Slice[int]{3, 2, 4, 8, 1, 3, 10}
	fmt.Printf("pre-s: %v\n", s)
	BubbleSort(s, comparef)
	fmt.Printf("post-s: %v\n", s)
}
