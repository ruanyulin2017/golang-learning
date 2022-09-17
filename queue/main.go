package queue

import (
	"errors"
	"fmt"
	"sync"
)

func Main() {
	n := 10
	gn := 3
	// que := NewArrayQueue(n)
	que := NewLinkQueue(n)
	finished := false

	wg := new(sync.WaitGroup)

	wg.Add(1)

	go func(q Queue, wg *sync.WaitGroup, finished *bool) {
		for i := 0; i < n; i++ {
			fmt.Printf("push [%d] to queue\n", i)
			if err := q.Push(i); err != nil {
				println(i)
				panic(err)
			}
		}
		wg.Done()
		*finished = true
	}(que, wg, &finished)

	for i := 0; i < gn; i++ {
		wg.Add(1)
		go func(q Queue, wg *sync.WaitGroup, finished *bool, gnum int) {
			for {
				data, err := q.Pop()
				if err != nil {
					if errors.Is(err, ErrQueueEnpty) {
						if *finished {
							fmt.Println("queue empty and send finished")
							break
						}
						continue
					}
				}
				// print(data, gnum, printL)
				fmt.Printf("goroutine [%d] pop data: %d\n", gnum, data)
			}
			wg.Done()
		}(que, wg, &finished, i)
	}

	wg.Wait()
}
