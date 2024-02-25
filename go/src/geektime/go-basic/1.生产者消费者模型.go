package main

import (
	"sync"
	"time"
)

type Queue struct {
	queue []string
	cond  *sync.Cond
}

func main() {
	q := Queue{
		queue: []string{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}

	go func() {
		time.Sleep(2 * time.Second)
		q.Enququq("1")
	}()

	for true {
		q.Dequeue()
		time.Sleep(2 * time.Second)
	}

}

func (q *Queue) Enququq(item string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.queue = append(q.queue, item)
	q.cond.Broadcast()
}

func (q *Queue) Dequeue() string {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	result := q.queue[0]
	q.queue = q.queue[1:]
	return result
}
