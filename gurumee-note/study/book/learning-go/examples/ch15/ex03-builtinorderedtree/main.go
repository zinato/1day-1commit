package main

import "fmt"

type BuildInOrdered interface {
	string | int | uint | float64 | uintptr
}

type Tree[T BuildInOrdered] struct {
	val         T
	left, right *Tree[T]
}

func (t *Tree[T]) Insert(val T) *Tree[T] {
	if t == nil {
		return &Tree[T]{val: val}
	}

	switch {
	case val < t.val:
		t.left = t.left.Insert(val)
	case val > t.val:
		t.right = t.right.Insert(val)
	}

	return t
}

func main() {
	var it *Tree[int]
	fmt.Println(it)

	it = it.Insert(5)
	fmt.Println(it)

	it = it.Insert(10)
	fmt.Println(it)

	it = it.Insert(3)
	fmt.Println(it)
}
