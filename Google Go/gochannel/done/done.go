package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d recieved %c\n", id, n)

		go func() {
			w.done()
		}()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)

	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		wg.Add(1)
	}
	//for _, worker := range workers {
	//	<-worker.done
	//}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		wg.Add(1)
	}
	// wait all of them
	//for _, worker := range workers {
	//	<-worker.done
	//}
	wg.Wait()
}

func main() {
	chanDemo()
}
