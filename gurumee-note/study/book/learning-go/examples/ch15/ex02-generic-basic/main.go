package main

import "fmt"

type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}

	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s *Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}

	return false
}

func main() {
	var s Stack[int]
	s.Push(10)
	fmt.Println(s)

	s.Push(20)
	fmt.Println(s)

	fmt.Println(s.Contains(10))
	fmt.Println(s.Contains(20))
	fmt.Println(s.Contains(30))

	v, ok := s.Pop()
	fmt.Println(v, ok)

	v, ok = s.Pop()
	fmt.Println(v, ok)

	v, ok = s.Pop()
	fmt.Println(v, ok)
}
