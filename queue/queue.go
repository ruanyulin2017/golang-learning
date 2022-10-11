package queue

import "errors"

type queueElemType interface{ any }

// type T = queueElemType

type Queue[T queueElemType] interface {
	Push(i T) error
	Pop() (T, error)
}

var ErrQueueFull = errors.New("QueueFullErr")
var ErrQueueEnpty = errors.New("QueueEnptyErr")
