package binarysearch

import (
	"errors"
	"fmt"
)

type slice[T binarySearchElemType] []T

func NewBinarySearch[T binarySearchElemType](cap int) BinarySearch[T] {
	s := make(slice[T], 0, cap)
	return &s
}

func (s slice[T]) bubbleSort(f func(a, b T) int8) {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			res := f(s[j], s[j+1])
			if res == 1 {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func (s *slice[T]) Add(i T, f func(a, b T) int8) {
	*s = append(*s, i)
	s.bubbleSort(f)
}

func (s *slice[T]) Adds(i slice[T], f func(a, b T) int8) {
	*s = append(*s, i...)
	s.bubbleSort(f)
}

func (s slice[T]) Search(i *T, f func(a, b T) int8) (err error) {
	l := len(s)

	// 头部或尾部匹配到直接返回
	if f(*i, s[0]) == 0 {
		*i = s[0]
		return
	}
	if f(*i, s[l-1]) == 0 {
		*i = s[l-1]
		return
	}

	hl := l / 2
	res := f(*i, s[hl])
	switch res {
	case -1:
		if f(*i, s[0]) == -1 {
			return errors.New("not found")
		}
		s[:hl].Search(i, f)
		return
	case 0:
		fmt.Printf("s, l: %v, %v\n", s, l)
		*i = s[hl]
		return
	case 1:

		if f(*i, s[l-1]) == 1 {
			return errors.New("not found")
		}
		// 左闭右开，s[hl] 可以不要
		s[hl+1:].Search(i, f)
		return
	default:
		return errors.New("not found")
	}
}
