package queue

import (
	"sync"
)

type node[T queueElemType] struct {
	data T
	next *node[T]
}

type linkQueue[T queueElemType] struct {
	len  int // 长度，包含元素个数
	cap  int // 容量，最多能存储的元素个数
	next *node[T]
	lock sync.RWMutex
}

func NewLinkQueue[T queueElemType](cap int) Queue[T] {
	return &linkQueue[T]{
		cap: cap,
	}
}

func (q *linkQueue[T]) Push(i T) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.len == q.cap {
		return ErrQueueFull
	}

	n := &node[T]{data: i}
	q.next = n
	q.len++
	return nil
}

func (q *linkQueue[T]) Pop() (T, error) {
	q.lock.RLock()
	defer q.lock.RUnlock()

	if q.len == 0 {
		var t T
		return t, ErrQueueEnpty
	}
	p := q.next
	q.next = q.next.next
	q.len--
	return p.data, nil
}
