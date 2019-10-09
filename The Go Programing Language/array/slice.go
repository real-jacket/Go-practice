package main

import (
	"fmt"
)

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	data1 := []string{"one", "", "three"}
	data2 := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data1))
	fmt.Printf("%q\n", data1)
	fmt.Printf("%q\n", nonempty2(data2))
	fmt.Printf("%q\n", data2)
	data3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(remove(data3, 3))
	fmt.Println(data3)
}
