package main

import (
	"fmt"
)

type person struct {
	age  int
	name string
}

func updatePerson(i int, s string, p person) {
	i *= 2
	s = "Goodbye"
	p.age = i
	p.name = s
	fmt.Println("func", i, s, p)
}

func updateMap(m map[int]string) {
	m[2] = "Hello"
	m[3] = "Goodbye"
	delete(m, 1)
}

func updateSlice(arr []int) {
	for k, v := range arr {
		arr[k] = v * 2
	}
	arr = append(arr, 10)
}
func main() {
	p := person{}
	i := 2
	s := "Hello"
	updatePerson(i, s, p)
	fmt.Println("main", i, s, p)

	m := map[int]string{
		1: "first",
		2: "second",
	}
	fmt.Println(m)
	updateMap(m)
	fmt.Println(m)

	arr := []int{1, 2, 3}
	fmt.Println(arr)
	updateSlice(arr)
	fmt.Println(arr)
}
