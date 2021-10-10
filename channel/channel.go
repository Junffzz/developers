/*
@Time : 2020/11/30 11:30
@Author : ZhaoJunfeng
@File : channel
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done()
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
		workers[i] = createWorker(i,&wg)
	}

	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
	}

	wg.Wait()
	time.Sleep(time.Millisecond)

}

//func bufferedChannel() {
//	c := make(chan int, 2)
//	go worker(0, c)
//	c <- 1
//	c <- 2
//	c <- 3
//
//}

//func channelClose() {
//	c := make(chan int)
//	go worker(0, c)
//	c <- 'a'
//	c <- 'b'
//	c <- 'c'
//	c <- 'd'
//	close(c)
//	time.Sleep(time.Millisecond)
//}

func main() {
	chanDemo()

	//bufferedChannel()
}
