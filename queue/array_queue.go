package queue

import (
	"sync"
)

type arayQueue struct {
	buf   []int
	len   int // 长度，包含元素个数
	cap   int // 容量，最多能存储的元素个数
	pushx int // push 索引，指向下次 push 位置
	popx  int // pop 索引，执行下次 pop 位置
	lock  sync.RWMutex
}

func NewArrayQueue(cap int) Queue {
	return &arayQueue{
		buf: make([]int, cap),
		cap: cap,
	}
}

func (q *arayQueue) Push(i int) error {
	q.lock.Lock() // 其他 goroutine 不可读不可写
	defer q.lock.Unlock()

	if q.len == q.cap {
		// 容量已满
		return ErrQueueFull
	}
	q.buf[q.pushx] = i

	q.len++
	q.pushx++
	if q.pushx == q.cap {
		q.pushx = 0
	}
	return nil
}

func (q *arayQueue) Pop() (int, error) {
	q.lock.RLock() // 其他 goroutine 可读不可写
	defer q.lock.RUnlock()

	if q.len == 0 {
		return 0, ErrQueueEnpty
	}
	popData := q.buf[q.popx]

	q.len--
	q.popx++
	if q.popx == q.cap {
		q.popx = 0
	}
	return popData, nil
}
