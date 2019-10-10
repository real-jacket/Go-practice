package main

import (
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// Point should be a potint
type Point struct{ X, Y int }

// Circle should be a  circle
type Circle struct {
	Point
	Radius int
}

// Wheel should be a  circle
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	a := []int{7, 23, 8, 5, 30}
	Sort(a)
	fmt.Println(a)
	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p == q)

	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
}
