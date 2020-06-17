package main

import (
	"fmt"
	"time"
)

func worker(id int,c chan int)  {
	for{
		n,ok := <-c
		if ok {
			fmt.Printf("Worker %d recieved %c\n",id,n)
		}
	}
}

func createWorker(id int) chan<- int{
	c := make(chan int)
	go func() {
		for{
			//n:= <-c
			//fmt.Println(n)
			fmt.Printf("Worker %d recieved %c\n",id,<-c)
		}
	}()
	return c
}

func chanDemo()  {
	//var c chan int // c == nil
	var channels [10]chan<- int
	for i:=0;i<10 ;i++  {
		channels[i] = createWorker(i)
	}
	for i:=0;i<10 ;i++  {
		channels[i] <- 'a' + i
	}
	for i:=0;i<10 ;i++  {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel()  {
	c := make(chan int)
	c <- 1
}

func channelClose()  {
	c :=make(chan int,3)
	go worker(0,c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	time.Sleep(time.Millisecond)
}

func main()  {
	//chanDemo()
	//bufferedChannel()
	channelClose()
}
