package queue

import (
	"sync"
)

type node struct {
	data int
	next *node
}

type linkQueue struct {
	len  int // 长度，包含元素个数
	cap  int // 容量，最多能存储的元素个数
	next *node
	lock sync.RWMutex
}

func NewLinkQueue(cap int) Queue {
	return &linkQueue{
		cap: cap,
	}
}

func (q *linkQueue) Push(i int) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.len == q.cap {
		return ErrQueueFull
	}

	n := &node{data: i}
	q.next = n
	q.len++
	return nil
}

func (q *linkQueue) Pop() (int, error) {
	q.lock.RLock()
	defer q.lock.RUnlock()

	if q.len == 0 {
		return 0, ErrQueueEnpty
	}
	p := q.next
	q.next = q.next.next
	q.len--
	return p.data, nil
}
