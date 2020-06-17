package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	var i int
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}


func main() {
	var c1, c2 chan int
	var n int
	select {
	case n = <-c1:
		fmt.Println("Received from c1:", n)
	case n = <-c2:
		fmt.Println("Received from c2:", c2)
	default:
		fmt.Println("No value received")
	}
}
