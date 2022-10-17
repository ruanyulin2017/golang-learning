package binarysearch

type binarySearchElemType interface{ any }

type BinarySearch[T binarySearchElemType] interface {
	Add(T, func(a, b T) int8)
	Adds(slice[T], func(a, b T) int8)
	Search(*T, func(a, b T) int8) error
}
