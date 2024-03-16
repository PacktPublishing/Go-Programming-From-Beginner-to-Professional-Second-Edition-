package main

import (
	"fmt"
	"sync"
	"time"
)

type WorkQueue struct {
	cond      *sync.Cond
	maxSize   int
	workItems []string
}

func NewWorkQueue(maxSize int) *WorkQueue {
	return &WorkQueue{
		cond:      sync.NewCond(&sync.Mutex{}),
		maxSize:   maxSize,
		workItems: make([]string, 0),
	}
}

func (wq *WorkQueue) enqueue(item string) {
	wq.cond.L.Lock()
	defer wq.cond.L.Unlock()
	for len(wq.workItems) == wq.maxSize {
		wq.cond.Wait()
	}
	wq.workItems = append(wq.workItems, item)
	wq.cond.Signal()
}

func (wq *WorkQueue) dequeue() string {
	wq.cond.L.Lock()
	defer wq.cond.L.Unlock()
	for len(wq.workItems) == 0 {
		wq.cond.Wait()
	}
	item := wq.workItems[0]
	wq.workItems = wq.workItems[1:]
	wq.cond.Signal()
	return item
}

func main() {
	var wg sync.WaitGroup
	workQueue := NewWorkQueue(3)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			workItem := fmt.Sprintf("WorkItem %d", i)
			workQueue.enqueue(workItem)
			fmt.Printf("Enqueued: %s\n", workItem)
			time.Sleep(time.Second)
		}
	}()
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			workItem := workQueue.dequeue()
			fmt.Printf("Dequeued: %s\n", workItem)
			time.Sleep(2 * time.Second)
		}
	}()
	wg.Wait()
}
