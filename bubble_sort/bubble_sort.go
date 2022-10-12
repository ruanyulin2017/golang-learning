package bubblesort

type Slice[T any] []T

func BubbleSort[T any](s Slice[T], f func(a, b T) int8) {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			res := f(s[j], s[j+1])
			if res == 1 {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}
