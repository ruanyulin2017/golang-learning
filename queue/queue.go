package queue

import "errors"

type Queue interface {
	Push(i int) error
	Pop() (int, error)
}

var ErrQueueFull = errors.New("QueueFullErr")
var ErrQueueEnpty = errors.New("QueueEnptyErr")
