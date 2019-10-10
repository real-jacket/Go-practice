package main

import (
	"fmt"
	"sort"
)

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, vx := range x {
		if vy, ok := y[k]; !ok || vy != vx {
			return false
		}
	}
	return true
}

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
		"jacke":   44,
		"kathy":   12,
		"bob":     7,
	}
	ages2 := map[string]int{
		"alice":   31,
		"charlie": 34,
		"jacke":   44,
		"kathy":   12,
		"bob":     8,
	}
	// var names []string
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		if age, ok := ages[name]; ok {
			fmt.Printf("%s\t%d\n", name, age)
		}
	}

	isEqual := equal(ages, ages2)
	fmt.Println(isEqual)
}
