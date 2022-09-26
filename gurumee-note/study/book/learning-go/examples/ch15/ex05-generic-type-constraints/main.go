package main

import "fmt"

type Rangeable[E comparable, T any] interface {
	[]T | map[E]T
}

func CopyValues[E comparable, T any, R Rangeable[E, T]](vals R) []T {
	var out []T
	for _, v := range vals {
		out = append(out, v)
	}

	return out
}

func main() {
	arr := []string{"a", "b", "c"}
	out := CopyValues[string, string, []string](arr)
	fmt.Println(out)

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	out2 := CopyValues[string, int, map[string]int](m)
	fmt.Println(out2)
}
