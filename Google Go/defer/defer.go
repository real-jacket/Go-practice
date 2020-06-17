package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writeFile("fib.txt")
}

func tryDefer() {
	defer fmt.Println("a")
	fmt.Println("b")
	fmt.Println("c")
	panic("error")
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
}
